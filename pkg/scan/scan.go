package scan

import (
	"strconv"
	"strings"
	"sync"

	"github.com/Karthika-Rajagopal/authz0/pkg/authz0"
	"github.com/Karthika-Rajagopal/authz0/pkg/logger"
	"github.com/Karthika-Rajagopal/authz0/pkg/models"
	"github.com/Karthika-Rajagopal/authz0/pkg/utils"
	"github.com/sirupsen/logrus"
)

type ScanArguments struct {
	RoleName     string
	Cookie       string
	Headers      []string
	Concurrency  int
	Delay        int
	Timeout      int
	ProxyAddress string
}

type Query struct {
	Index int
	Query models.URL
}

func Run(filename string, arguments ScanArguments, debug bool) []models.Result {
	var results []models.Result
	var wg sync.WaitGroup
	log := logger.GetLogger(debug)
	queries := make(chan Query)
	template := authz0.FileToTemplate(filename)
	log.Info("loaded testing template: " + template.Name)
	if arguments.RoleName != "" {
		log.Info("append credential to template [" + arguments.RoleName + " role]")
		tCred := models.Credential{
			Rolename: arguments.RoleName,
			Headers:  arguments.Headers,
		}
		template.Credentials = append(template.Credentials, tCred)
	}
	for i := 0; i < arguments.Concurrency; i++ {
		wg.Add(1)
		go func() {
			for query := range queries {
				reqURL := query.Query
				for _, cred := range template.Credentials {
					res, cl, err := sendReq(reqURL, arguments, template, cred.Headers)
					if err != nil {
						log.Debug("sendReq Error")
						log.Trace(err)
					}
					aar := false
					adr := false
					rlt := "O"
					rltValue := true
					check := checkAssert(res, template.Asserts, cl)
					if cred.Rolename != "" {
						if check {
							if len(reqURL.AllowRole) > 0 {
								if utils.ContainsFromArray(reqURL.AllowRole, cred.Rolename) {
									aar = true
									rltValue = rltValue && true
								} else {
									rltValue = rltValue && false
								}
							}
							if len(reqURL.DenyRole) > 0 {
								if utils.ContainsFromArray(reqURL.DenyRole, cred.Rolename) {
									adr = false
									rltValue = rltValue && false
								} else {
									rltValue = rltValue && true
								}
							}
						} else {
							if len(reqURL.AllowRole) > 0 {
								if utils.ContainsFromArray(reqURL.AllowRole, cred.Rolename) {
									aar = true
									rltValue = rltValue && false
								} else {
									rltValue = rltValue && true
								}
							}
							if len(reqURL.DenyRole) > 0 {
								if utils.ContainsFromArray(reqURL.DenyRole, cred.Rolename) {
									adr = true
									rltValue = rltValue && true
								}
							}
						}
						if !rltValue {
							rlt = "X"
						}
					}

					result := models.Result{
						URL:             reqURL.URL,
						Method:          reqURL.Method,
						RoleName:        cred.Rolename,
						AllowRole:       reqURL.AllowRole,
						DenyRole:        reqURL.DenyRole,
						Assert:          check,
						AssertAllowRole: aar,
						AssertDenyRole:  adr,
						StatusCode:      res.StatusCode,
						RespSize:        cl,
						Alias:           reqURL.Alias,
						Result:          rlt,
						Index:           "#" + strconv.Itoa(query.Index),
					}
					results = append(results, result)
					iLog := log.WithField("index", "#"+strconv.Itoa(query.Index))
					if result.Alias != "" {
						iLog.Info("check '" + result.Alias + "'")
					} else {
						iLog.Info("check '" + result.URL + "'")
					}
					uLog := iLog.WithFields(logrus.Fields{
						"url":  result.Method + " " + result.URL,
						"type": "assert",
					})
					uLog.Info("response code: " + strconv.Itoa(result.StatusCode))
					if result.Assert {
						uLog.Info("assertion: hit")
					} else {
						if cred.Rolename == "" {
							if check {
								uLog.Info("assertion: fail")
							} else {
								uLog.WithField("assertion", "assertion: fail").Warn("found assert fail")
							}
						} else {
							uLog.Info("assertion: fail")
						}
					}
					rLog := iLog.WithFields(logrus.Fields{
						"type": "role-test",
					})
					ar := strings.Join(result.AllowRole, ",")
					dr := strings.Join(result.DenyRole, ",")
					if ar == "" {
						ar = "<allow-all>"
					}
					if dr == "" {
						dr = "<not-deny>"
					}
					rLog.Info("allow-role: " + ar)
					rLog.Info("deny-role: " + dr)
					if cred.Rolename != "" {
						if !rltValue {
							rLog.WithFields(logrus.Fields{
								"role-match": "role-match: " + rlt,
								"role-name":  "role-name: " + result.RoleName,
							}).Warn("found role mismatch")
						} else {
							rLog.Info("role-match: " + rlt + " (" + result.RoleName + ")")
						}
					}

					if result.AssertAllowRole {
						rLog.Info("matched: allow")
					}
					if result.AssertDenyRole {
						rLog.Info("matched: deny")
					}

					log.WithFields(logrus.Fields{
						"status": result.StatusCode,
						"alias":  result.Alias,
						"aar":    aar,
						"adr":    adr,
						"size":   cl,
					}).Debug("")
				}
			}
			wg.Done()
		}()
	}
	log.Info("targets:  " + strconv.Itoa(len(template.URLs)) + " URLs")
	for index, endpoint := range template.URLs {
		q := Query{
			Index: index,
			Query: endpoint,
		}
		queries <- q
	}
	close(queries)
	wg.Wait()
	return results
}

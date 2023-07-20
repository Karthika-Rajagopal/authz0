package report

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/Karthika-Rajagopal/authz0/pkg/logger"
	"github.com/Karthika-Rajagopal/authz0/pkg/models"
	"github.com/olekukonko/tablewriter"
	"gopkg.in/yaml.v2"
)

func PrettyJSON(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func WriteYAMLReportToFile(data []models.Result, filename string) {
	yamlData, err := yaml.Marshal(&data)
	log := logger.GetLogger(false)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filename, yamlData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func PrintTableReport(data []models.Result, t string) {
	results := data
	sort.Slice(results, func(i, j int) bool {
		iIndex, _ := strconv.Atoi(strings.ReplaceAll(results[i].Index, "#", ""))
		jIndex, _ := strconv.Atoi(strings.ReplaceAll(results[j].Index, "#", ""))
		return iIndex < jIndex
	})

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "Alias", "Assert", "Role", "Allow-Role", "Deny-Role", "Allow", "Deny", "RLT"})
	if t == "markdown" {
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
	}
	table.SetHeaderColor(
		nil, nil,
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		nil,
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgCyanColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
		tablewriter.Colors{tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor},
	)

	issue := 0
	for _, v := range data {
		ar := strings.Join(v.AllowRole, ",")
		dr := strings.Join(v.DenyRole, ",")
		if ar == "" {
			ar = "<ALLOWED-ALL>"
		}
		if dr == "" {
			dr = "<NOT-DENIED>"
		}
		line := []string{
			v.Index,
			v.Alias,
			strconv.FormatBool(v.Assert),
			v.RoleName,
			ar,
			dr,
			strconv.FormatBool(v.AssertAllowRole),
			strconv.FormatBool(v.AssertDenyRole),
			" " + v.Result + " ",
		}
		if v.Result == "X" {
			issue = issue + 1
			table.Rich(line, []tablewriter.Colors{
				tablewriter.Colors{tablewriter.FgHiRedColor},
				tablewriter.Colors{tablewriter.FgHiRedColor},
				tablewriter.Colors{tablewriter.FgHiRedColor},
				tablewriter.Colors{tablewriter.FgHiRedColor},
				tablewriter.Colors{tablewriter.FgHiRedColor},
				tablewriter.Colors{tablewriter.FgHiRedColor},
				tablewriter.Colors{tablewriter.FgHiRedColor},
				tablewriter.Colors{tablewriter.FgHiRedColor},
				tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor},
			})
		} else {
			table.Append(line)
		}
	}
	table.Render()
}

func PrintTableURLs(data []models.Result, t string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "Method", "URL", "CODE"})
	if t == "markdown" {
		table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
		table.SetCenterSeparator("|")
	}

	issue := 0
	for _, v := range data {
		line := []string{
			v.Index,
			v.Method,
			v.URL,
			strconv.Itoa(v.StatusCode),
		}
		if v.Result == "X" {
			issue = issue + 1
			table.Rich(line, []tablewriter.Colors{
				tablewriter.Colors{},
				tablewriter.Colors{},
				tablewriter.Colors{},
				tablewriter.Colors{},
			})
		} else {
			table.Append(line)
		}

	}
	table.Render()
}

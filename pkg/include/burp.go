package include

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/Karthika-Rajagopal/authz0/pkg/logger"
)

type Items struct {
	XMLName     xml.Name `xml:"items"`
	Text        string   `xml:",chardata"`
	BurpVersion string   `xml:"burpVersion,attr"`
	ExportTime  string   `xml:"exportTime,attr"`
	Item        []struct {
		Text string `xml:",chardata"`
		Time string `xml:"time"`
		URL  string `xml:"url"`
		Host struct {
			Text string `xml:",chardata"`
			Ip   string `xml:"ip,attr"`
		} `xml:"host"`
		Port      string `xml:"port"`
		Protocol  string `xml:"protocol"`
		Method    string `xml:"method"`
		Path      string `xml:"path"`
		Extension string `xml:"extension"`
		Request   struct {
			Text   string `xml:",chardata"`
			Base64 string `xml:"base64,attr"`
		} `xml:"request"`
		Status         string `xml:"status"`
		Responselength string `xml:"responselength"`
		Mimetype       string `xml:"mimetype"`
		Response       struct {
			Text   string `xml:",chardata"`
			Base64 string `xml:"base64,attr"`
		} `xml:"response"`
		Comment string `xml:"comment"`
	} `xml:"item"`
}

func ImportBurpFormat(filename string) Items {
	var burpObject Items
	log := logger.GetLogger(false)
	burpFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("no such file or directory")
	}
	err = xml.Unmarshal(burpFile, &burpObject)
	if err != nil {
		log.Fatal("xml unmarshal error")
	}
	return burpObject
}

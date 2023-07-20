package include

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Karthika-Rajagopal/authz0/pkg/logger"
	volt "github.com/Karthika-Rajagopal/volt/format/har"
)

func ImportZAPFormat(filename string) volt.HARObject {
	var harObject volt.HARObject
	harFile, err := ioutil.ReadFile(filename)
	log := logger.GetLogger(false)
	if err != nil {
		log.Fatal("no such file or directory")
	}
	err = json.Unmarshal(harFile, &harObject)
	if err != nil {
		log.Fatal("json unmarshal error")
	}
	return harObject
}

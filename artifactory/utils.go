package artifactory

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jfrog/jfrog-client-go/artifactory/services/utils"
)

// AqlFromSpecFile parses and returns Aql from a spec file
func AqlFromSpecFile(filePath string) (utils.Aql, error) {
	var aql utils.Aql

	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return aql, err
	}

	err = json.Unmarshal(dat, &aql)
	if err != nil {
		return aql, err
	}

	return aql, nil
}

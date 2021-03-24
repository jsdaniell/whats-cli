package json_utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type PackageJson struct {
	Version string `json:"version"`
}

func GettingJSONVersion() (string) {
	file, err := os.Open("package.json")
	if err != nil {
		return err.Error()
	}

	defer file.Close()

	var version PackageJson

	byteValue, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(byteValue, &version)
	if err != nil {
		return err.Error()
	}

	return version.Version
}

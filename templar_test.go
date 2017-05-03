package templar

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"fmt"
	"os"
)

func TestTemplateVars(t *testing.T) {

	fmt.Println("Testing vars")

	params := make(map[string]string)
	params["elasticurl"] = "Chiptop.com:9200"

	out, _ := ParseTemplateFile("fixtures/test.tpl", params)
	assert.Contains(t, out, params["elasticurl"])
	fmt.Println("Test done successfully.")
}

func TestTemplateJson(t *testing.T) {
	fmt.Println("Test JSON data")
	file, _ := ioutil.ReadFile("fixtures/vars.json")

	var varsJSON interface{}
	json.Unmarshal(file, &varsJSON)

	outJSON, _ := ParseTemplateFile("fixtures/test.tpl", varsJSON)
	assert.Contains(t, outJSON, "localhost:9200")

	fmt.Println("Finished Test with JSON file...")

}

func TestTemplateErrorJson(t *testing.T) {
	fmt.Println("Running Testing throwing error...")
	file, _ := ioutil.ReadFile("fixtures/vars.json-should-not-exist")

	var varsJSON interface{}
	json.Unmarshal(file, &varsJSON)

	_, err := ParseTemplateFile("should-not-exist.tpl", varsJSON)
	assert.Error(t, err)
	fmt.Println("Test done")
}

func TestCreateFileByTemplate(t *testing.T) {

	fmt.Println("Testing vars")

	params := make(map[string]string)
	params["elasticurl"] = "Chiptop.com:9200"

	destFile := "fixtures/test.out"

	if _, err := os.Stat(destFile); os.IsExist(err) {
		_ = os.Remove(destFile)
	}

	err := CreateFileByTemplate("fixtures/test.tpl", destFile, params)
	if err != nil {
		t.Errorf("Error creating file by template %v", err)
	}

	fmt.Println("Test done successfully.")

}

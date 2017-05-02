package templar

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"

	"fmt"
)

func TestTemplateVars(t *testing.T) {

	fmt.Println("Running Test with vars...")

	params := make(map[string]string)
	params["package_name"] = "Blitzkrieg Bop"
	params["phrase1"] = "Hey ho, let's go"
	out, _ := ParseTemplateFile("fixtures/test.tpl", params)
	assert.Contains(t, out, "# Blitzkrieg Bop")
	assert.Contains(t, out, "Hey ho, let's go")
	fmt.Println(out)
	fmt.Println("Finished Test with vars...")

}

func TestTemplateJson(t *testing.T) {
	fmt.Println("Running Test with JSON file...")
	file, _ := ioutil.ReadFile("fixtures/vars.json")

	var varsJson interface{}
	json.Unmarshal(file, &varsJson)

	outJson, _ := ParseTemplateFile("fixtures/test.md.tpl", varsJson)
	//assert.Contains(t, out, []string{"# Blitzkrieg Bop","Hey ho, let's go"})
	assert.Contains(t, outJson, "Hey ho, let's go")

	fmt.Println(string(outJson))
	fmt.Println("Finished Test with JSON file...")

}
func TestTemplateErrorJson(t *testing.T) {
	fmt.Println("Running Testing throwing error...")
	file, _ := ioutil.ReadFile("fixtures/vars.json-should-not-exist")

	var varsJson interface{}
	json.Unmarshal(file, &varsJson)

	_, err := ParseTemplateFile("should-not-exist.tpl", varsJson)
	assert.Error(t, err)
	fmt.Println("Finished Testing throwing error...\n")
}

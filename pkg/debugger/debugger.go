package debugger

import (
	"encoding/json"
	"fmt"

	"github.com/ghodss/yaml"
)

func JsonPrint(res interface{}) {
	jsonres, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(jsonres))
}

func YamlPrint(res interface{}) {
	jsonres, _ := json.Marshal(res)
	yamlres, _ := yaml.JSONToYAML(jsonres)
	fmt.Println(string(yamlres))
}

func StructToYamlString(res interface{}) string {
	jsonres, _ := json.Marshal(res)
	yamlres, _ := yaml.JSONToYAML(jsonres)
	return string(yamlres)
}
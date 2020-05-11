package helpers

import (
    "encoding/json"
    yaml "gopkg.in/yaml.v3"
)

func IsJSON(s string) bool {
    var js map[string]interface{}
    return json.Unmarshal([]byte(s), &js) == nil
}

func IsYAML(s string) bool {
    var yml map[string]interface{}
    return yaml.Unmarshal([]byte(s), &yml) == nil
}

func IsJSONString(s string) bool {
    var js string
    return json.Unmarshal([]byte(s), &js) == nil
}

func IsYAMLString(s string) bool {
    var yml string
    return yaml.Unmarshal([]byte(s), &yml) == nil
}

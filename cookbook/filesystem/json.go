package filesystem

import (
	"encoding/json"
	"strings"
	"fmt"
)

var (
	tony = `
{
	"hero": {
		"name": "Tony Stack",
		"nickname": "Iron Man"
	},
	"called": ["花花公子", "百万富翁", "慈善家"],
	"status": "gone"
}
`
	steve = `
{
	"hero": {
		"name": "Steve Rogers",
		"nickname": "Captain America"
	},
	"called": ["五五开", "美国翘臀"],
	"status": "old"
}
`
)

func ParseHero(hero string) (string, error) {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(hero), &m)
	if err != nil {
		return "", err
	}

	//解析已知结构的JSON格式文本，依靠不断地断言
	people := m["hero"].(map[string]interface{})
	name := people["name"].(string)
	nickname := people["nickname"].(string)
	socialTmp := m["called"].([]interface{})
	var socialSlice  []string
	for _, s := range socialTmp {
		socialSlice = append(socialSlice, s.(string))
	}
	social := strings.Join(socialSlice, "|")
	status := m["status"].(string)
	return fmt.Sprintf("the hero %s, real name is %s, he's called %s. he is %s .", nickname, name, social, status), nil
}

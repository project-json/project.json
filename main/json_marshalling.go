/**
 * See also: https://github.com/tidwall/gjson
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const MapStringInterface = "map[string]interface {}"

type Component struct {
	Order int    `json:"-"`
	Name  string `json:"name"`
	Type  string `json:"-"`
}

type ComponentMap map[string]*Component

type Stack struct {
	Components ComponentMap
}

type Project struct {
	Stack Stack `json:"stack"`
}

func LoadProject() Project {
	raw, err := ioutil.ReadFile("./project.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var project Project
	err = json.Unmarshal(raw, &project)
	if err != nil {
		panic("Oh the pain!")
	}

	return project
}

func (s *Stack) UnmarshalJSON(b []byte) error {
	var raw interface{}

	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	s.Components = make(ComponentMap)

	mi := raw.(map[string]interface{})
	order := 0
	for k, v := range mi {
		order++
		val := fmt.Sprintf("%#v", v)
		var c Component
		if !strings.HasPrefix(val, MapStringInterface) {
			c = *new(Component)
			c.Name = strings.Trim(val, "\"")
			c.Type = k
			s.Components[k] = &c
		} else {
			val = val[len(MapStringInterface):]
			if err := json.Unmarshal([]byte(val), &c); err != nil {
				return err
			}
			c.Type = k
			s.Components[k] = &c
		}
		c.Order = order
	}

	return nil
}

func main() {
	project := LoadProject()
	for t, c := range project.Stack.Components {
		fmt.Printf("\n[%d] %-22v %v", c.Order, t+":", c.Name)
	}

}

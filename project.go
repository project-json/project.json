package projectcfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Projects []*Project

type Project struct {
	Schema      string                 `json:"schema,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Hostname    string                 `json:"hostname,omitempty"`
	IsEnabled   bool                   `json:"enabled"`
	Identity    string                 `json:"identity,omitempty"`
	Description string                 `json:"description,omitempty"`
	Namespace   string                 `json:"namespace,omitempty"`
	Slug        string                 `json:"slug,omitempty"`
	Prefix      string                 `json:"prefix,omitempty"`
	Type        ProjectType            `json:"type,omitempty"`
	Aliases     []string               `json:"aliases,omitempty"`
	Team        Team                   `json:"team,omitempty"`
	Stack       Stack                  `json:"stack,omitempty"`
	Dev         *Dev                   `json:"dev,omitempty"`
	Source      *Source                `json:"source,omitempty"`
	Deploy      *Deploy                `json:"deploy,omitempty"`
	Hosts       HostMap                `json:"hosts,omitempty"`
	Extra       map[string]interface{} `json:"extra,omitempty"`
	Root        *string                `json:"-"`
}

func Load(cfgPath string) *Project {
	b, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		log.Fatal(err)
	}
	p := Project{}
	err = json.Unmarshal(b, &p)
	if err != nil {
		log.Fatal(err)
	}
	return &p
}

func New(name string, root *string) *Project {
	domain := name
	if !strings.Contains(name, ".") {
		domain = fmt.Sprintf("%s.local", name)
	}
	pr := Project{
		Root:     root,
		Name:     name,
		Hostname: domain,
	}
	return &pr
}

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
	Schema      string      `cfg:"schema,omitempty"`
	Name        string      `cfg:"name,omitempty"`
	Hostname    string      `cfg:"hostname,omitempty"`
	IsEnabled   bool        `cfg:"enabled"`
	Identity    string      `cfg:"identity,omitempty"`
	Description string      `cfg:"description,omitempty"`
	Namespace   string      `cfg:"namespace,omitempty"`
	Slug        string      `cfg:"slug,omitempty"`
	Prefix      string      `cfg:"prefix,omitempty"`
	Type        ProjectType `cfg:"type,omitempty"`
	Aliases     []string    `cfg:"aliases,omitempty"`
	Dev         *Dev        `cfg:"dev,omitempty"`
	Stack       Stack       `cfg:"stack,omitempty"`
	Source      *Source     `cfg:"source,omitempty"`
	Deploy      *Deploy     `cfg:"deploy,omitempty"`
	Hosts       HostMap     `cfg:"hosts,omitempty"`
	Root        *string     `cfg:"-"`
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

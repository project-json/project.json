package projectcfg

import (
	"encoding/json"
	"fmt"
	"github.com/projectcfg/projectcfg/util"
	"io/ioutil"
	"log"
	"strings"
)

var Instance *ProjectCfg

func Initialize(filepath string) {
	Instance = Load(filepath)
}

type ProjectCfg struct {
	Schema           string                 `json:"schema,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Hostname         string                 `json:"hostname,omitempty"`
	Identity         string                 `json:"identity,omitempty"`
	Description      string                 `json:"description,omitempty"`
	Namespace        string                 `json:"namespace,omitempty"`
	Slug             string                 `json:"slug,omitempty"`
	Prefix           string                 `json:"prefix,omitempty"`
	Type             ProjectType            `json:"type,omitempty"`
	Aliases          []string               `json:"aliases,omitempty"`
	Team             Team                   `json:"team,omitempty"`
	Stack            Stack                  `json:"stack,omitempty"`
	Dev              *Dev                   `json:"dev,omitempty"`
	Source           *Source                `json:"source,omitempty"`
	Deploy           *Deploy                `json:"deploy,omitempty"`
	Hosts            HostMap                `json:"hosts,omitempty"`
	Extra            map[string]interface{} `json:"extra,omitempty"`
	Root             *string                `json:"-"`
	loadedFrameworks FrameworkMap
}

func Load(cfgPath string) *ProjectCfg {
	b, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		util.Error(err)
	}
	p := ProjectCfg{}
	err = json.Unmarshal(b, &p)
	if err != nil {
		log.Fatal(err)
	}
	return &p
}

func New(name string, root *string) *ProjectCfg {
	domain := name
	if !strings.Contains(name, ".") {
		domain = fmt.Sprintf("%s.local", name)
	}
	pr := ProjectCfg{
		Root:     root,
		Name:     name,
		Hostname: domain,
	}
	return &pr
}

func (me *ProjectCfg) String() string {
	j, err := json.MarshalIndent(me, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(j)
}

func (me *ProjectCfg) GetFramework(name FrameworkName) (fw Framework) {
	fw, _ = me.loadedFrameworks[name]
	return
}

//
//@TODO Needs to allow additonal framework sources defined in the config file
//      Sources need to be able to be downloaded, but need to be secure
//@TODO Waiting on a GoLang bug fix: https://youtrack.jetbrains.com/issue/GO-6289
//
//var pathTemplate = "%s/frameworks/*.pcfw"
//
//func (me *ProjectCfg) LoadFrameworks() {
//	path := fmt.Sprintf(pathTemplate,util.GetProjectDir())
//	files, err := filepath.Glob(path)
//	if err != nil {
//		log.Fatal(fmt.Sprintf("Directory listing of framework file in '%s' failed: %s",path,err))
//	}
//	for _,file := range files {
//		fwpi, err := plugin.Open(file)
//		if err != nil {
//			log.Fatal(fmt.Sprintf("Open framework file '%s' failed: %s",file,err))
//		}
//		loader, err := fwpi.Lookup("GetInstance")
//		if err != nil {
//			log.Fatal(fmt.Sprintf("Lookup GetInstance in framework file '%s' failed: %s",file,err))
//		}
//		getter, ok := loader.(func() interface{})
//		if !ok {
//			log.Fatal(fmt.Sprintf("Type assert for framework file '%s' failed: %s",file,err))
//		}
//		generic := getter()
//		framework, ok := generic.(Framework)
//		if !ok {
//			log.Fatal(fmt.Sprintf("Framework file '%s' in not a valid framework: %s",file,err))
//		}
//		me.loadedFrameworks[FrameworkName(filepath.Base(file))] = framework
//	}
//}

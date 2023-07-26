package generator

import (
	"os"
	"path"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/zrs01/sct/internal/config"
	"github.com/zrs01/sct/internal/ts"
	"github.com/zrs01/sct/internal/utils"

	"github.com/CloudyKit/jet/v6"
	"gopkg.in/yaml.v2"
)

type Option struct {
	Entity       string
	TemplateFile string
	DataFile     string // user-defined data, it will be passed to the template engine
}

type Context struct {
	Name   string
	Entity []utils.Entity
	Data   map[interface{}]interface{}
}

func Output(option Option) error {
	cfg := config.GetConfig()

	var csClasses []utils.Entity
	csFiles := flatenFile(option.Entity)
	for _, file := range csFiles {
		csClasses = append(csClasses, utils.ParseEntity(file))
	}

	// user-defined data
	data := make(map[interface{}]interface{})
	if option.DataFile != "" {
		fileContent, err := os.ReadFile(option.DataFile)
		if err != nil {
			panic(err)
		}
		if err := yaml.Unmarshal([]byte(fileContent), &data); err != nil {
			panic(err)
		}
	}

	context := Context{
		Entity: csClasses,
		Data:   data,
	}
	templateFile := option.TemplateFile
	if cfg.TemplatePath != "" {
		templateFile = path.Join(cfg.TemplatePath, option.TemplateFile)
	}

	// define global function
	vars := make(jet.VarMap)
	vars.SetFunc("typescriptType", func(a jet.Arguments) reflect.Value {
		// parameters
		dataType := a.Get(0)
		isCollection := a.Get(1)
		return reflect.ValueOf(ts.ToType(dataType.String(), isCollection.Bool()))
	})

	if err := utils.GetTemplate(templateFile).Execute(os.Stdout, vars, context); err != nil {
		panic(err)
	}
	return nil
}

func flatenFile(input string) []string {
	cfg := config.GetConfig()

	files := strings.Split(input, ",")
	if len(cfg.EntityPath) > 0 {
		for j := 0; j < len(cfg.EntityPath); j++ {
			for i := 0; i < len(files); i++ {
				files[i] = filepath.Join(cfg.EntityPath[j], files[i])
			}
		}
	}
	return utils.SearchExactGlobFiles(files)
}

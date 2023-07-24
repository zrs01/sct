package generator

import (
	"os"
	"path"
	"path/filepath"
	"sct/internal/config"
	"sct/internal/utils"
	"strings"

	"gopkg.in/yaml.v2"
)

type Option struct {
	Entity       string
	TemplateFile string
	ModuleName   string
	DataFile     string // user-defined data, it will be passed to the template engine
}

type Context struct {
	Name   string
	Entity []utils.Csharp
	Data   map[interface{}]interface{}
}

func Output(option Option) error {
	cfg := config.GetConfig()

	var csClasses []utils.Csharp
	csFiles := flatenFile(option.Entity)
	for _, file := range csFiles {
		csClass := utils.CsharpParse(file)
		csClasses = append(csClasses, csClass)
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
		Name:   option.ModuleName,
		Entity: csClasses,
		Data:   data,
	}
	templateFile := option.TemplateFile
	if cfg.TemplatePath != "" {
		templateFile = path.Join(cfg.TemplatePath, option.TemplateFile)
	}

	utils.MergeTemplate(templateFile, context, os.Stdout)
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

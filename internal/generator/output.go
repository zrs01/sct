package generator

import (
	"os"
	"path"
	"path/filepath"
	"srcode/internal/config"
	"srcode/internal/utils"
	"strings"

	"gopkg.in/yaml.v2"
)

type Option struct {
	CsharpFile   string
	TemplateFile string
	ModuleName   string
	Namespace    string
	DataFile     string // user-defined data, it will be passed to the template engine
}

type Context struct {
	Name string
	Cs   []utils.Csharp
	Ts   []utils.TypeScript
	Data map[interface{}]interface{}
}

func Output(option Option) error {
	cfg := config.GetConfig()
	if len(cfg.DaoPath) == 0 {
		panic("DAO path cannot be found in configuration")
	}

	var csClasses []utils.Csharp
	var tsClasses []utils.TypeScript
	csFiles := flatenFile(option.CsharpFile)
	for _, file := range csFiles {
		csClass := utils.CsharpParse(file)
		csClass.Namespace = option.Namespace
		csClasses = append(csClasses, csClass)
		tsClasses = append(tsClasses, utils.TypescriptParse(file))
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
		Name: option.ModuleName,
		Cs:   csClasses,
		Ts:   tsClasses,
		Data: data,
	}
	utils.MergeTemplate(path.Join(cfg.TemplatePath, option.TemplateFile), context, os.Stdout)
	return nil
}

func flatenFile(input string) []string {
	cfg := config.GetConfig()

	files := strings.Split(input, ",")
	for i := 0; i < len(files); i++ {
		files[i] = filepath.Join(cfg.DaoPath[0], files[i])
	}
	return utils.SearchExactGlobFiles(files)
}

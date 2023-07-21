package generator

import (
	"os"
	"path"
	"path/filepath"
	"srcode/internal/config"
	"srcode/internal/utils"
	"strings"
)

type Option struct {
	CsharpFile   string
	TemplateFile string
	ModuleName   string
	Namespace    string
}

type Context struct {
	Name string
	Cs   []utils.Csharp
	Ts   []utils.TypeScript
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

	context := Context{
		Name: option.ModuleName,
		Cs:   csClasses,
		Ts:   tsClasses,
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

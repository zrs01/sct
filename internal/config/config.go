package config

import (
	"github.com/jinzhu/configor"
	"github.com/shomali11/util/xstrings"
)

type Setting struct {
	LangType     string   `default:"dotnet" yaml:"lang-type"` // dotnet | java
	EntityPath   []string `default:"." yaml:"entity-path"`
	TemplatePath string   `yaml:"template-path"`
}

var cfile string

func InitConfig(f string) {
	cfile = f
}

func GetConfig() Setting {
	var setting Setting
	if xstrings.IsBlank(cfile) {
		cfile = "config.yml"
	}
	if err := configor.Load(&setting, cfile); err != nil {
		panic(err)
	}
	return setting
}

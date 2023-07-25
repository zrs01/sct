package utils

import (
	"path/filepath"

	"github.com/CloudyKit/jet/v6"
)

func GetTemplate(tpl string) *jet.Template {
	loader := jet.NewOSFileSystemLoader(filepath.Dir(tpl))

	views := jet.NewSet(loader, jet.InDevelopmentMode())
	view, err := views.GetTemplate(filepath.Base(tpl))
	if err != nil {
		panic(err)
	}
	return view
}

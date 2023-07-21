package utils

import (
	"os"
	"path/filepath"

	"github.com/CloudyKit/jet/v6"
)

func MergeTemplate(tpl string, data interface{}, out *os.File) {
	loader := jet.NewOSFileSystemLoader(filepath.Dir(tpl))

	views := jet.NewSet(loader, jet.InDevelopmentMode())
	view, err := views.GetTemplate(filepath.Base(tpl))
	if err != nil {
		panic(err)
	}
	if err := view.Execute(out, nil, data); err != nil {
		panic(err)
	}
}

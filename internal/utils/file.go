package utils

import (
	"fmt"
	"path/filepath"
	"runtime"
	"srcode/internal/config"
	"strings"
	"unicode"

	"github.com/rotisserie/eris"
	"github.com/thoas/go-funk"
)

func GetEntityFiles(entity string) []string {
	config := config.GetConfig()
	igfiles := strings.Split(entity, ",")
	for i := 0; i < len(igfiles); i++ {
		for j := 0; j < len(config.DaoPath); j++ {
			igfiles[i] = filepath.Join(config.DaoPath[j], igfiles[i])
		}
	}
	return igfiles
}

func SearchExactGlobFiles(sfiles []string) []string {
	rfiles, err := SearchGlobFiles(sfiles)
	if err != nil {
		panic(err)
	}
	return rfiles
}

func SearchGlobFiles(sfiles []string) ([]string, error) {
	rfiles := []string{}
	for _, sfile := range sfiles {
		sfile = strings.TrimSpace(sfile)
		files, err := filepath.Glob(insensitiveFilepath(sfile))
		// fmt.Printf("%+v\n", files)
		if err != nil {
			return files, eris.Wrapf(err, "failed to search file with %s", sfile)
		}
		rfiles = append(rfiles, files...)
	}
	return funk.Uniq(rfiles).([]string), nil
}

func insensitiveFilepath(path string) string {
	if runtime.GOOS == "windows" {
		return path
	}

	var sb strings.Builder
	for _, r := range path {
		if unicode.IsLetter(r) {
			sb.WriteString(fmt.Sprintf("[%c%c]", unicode.ToLower(r), unicode.ToUpper(r)))
		} else {
			sb.WriteString(string(r))
		}
	}
	return sb.String()
}

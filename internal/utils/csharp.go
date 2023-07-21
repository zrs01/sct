package utils

import (
	"os"
	"regexp"
	"strings"
)

type Csharp struct {
	Namespace string
	Name      string
	Members   []CsharpMember
}

type CsharpMember struct {
	Name           string
	Virtual        bool
	Optional       bool
	DataType       string
	CollectionType string
}

func CsharpParse(fullPathFileName string) Csharp {
	var csharp Csharp
	fileContent, err := os.ReadFile(fullPathFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(fileContent), "\n")

	rClass := regexp.MustCompile(`public\s+class\s+(\w+)`)
	rMember := regexp.MustCompile(`public\s+(virtual\s)?([0-9A-Za-z_<>\[\]]+)(\??)\s+(\w+)`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, "public partial", "public")

		matchClass := rClass.FindStringSubmatch(line)
		if matchClass != nil {
			csharp.Name = matchClass[1]
		} else {
			matchMember := rMember.FindStringSubmatch(line)
			if matchMember != nil {
				dotnetMember := CsharpMember{
					Name:     matchMember[4],
					Virtual:  len(matchMember[1]) > 0,
					Optional: len(matchMember[3]) > 0,
					DataType: matchMember[2],
				}
				if strings.HasPrefix(dotnetMember.DataType, "ICollection") {
					rCollection := regexp.MustCompile(`ICollection<(\w+)>`)
					matchCollection := rCollection.FindStringSubmatch(dotnetMember.DataType)
					if matchCollection != nil {
						dotnetMember.CollectionType = matchCollection[1]
					}
				}
				csharp.Members = append(csharp.Members, dotnetMember)
			}
		}
	}
	return csharp
}

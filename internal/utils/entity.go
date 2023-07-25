package utils

import (
	"os"
	"regexp"
	"strings"
)

type Entity struct {
	Name    string
	Members []EntityMember
}

type EntityMember struct {
	Name         string
	Virtual      bool
	Optional     bool
	DataType     string
	IsCollection bool
}

func ParseEntity(fullPathFileName string) Entity {
	var entity Entity
	fileContent, err := os.ReadFile(fullPathFileName)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(fileContent), "\n")

	regexClass := regexp.MustCompile(`public\s+class\s+(\w+)`)
	regexMember := regexp.MustCompile(`public\s+(virtual\s)?([0-9A-Za-z_<>\[\]]+)(\??)\s+(\w+)`)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, "public partial", "public")

		matchClass := regexClass.FindStringSubmatch(line)
		if matchClass != nil {
			entity.Name = matchClass[1]
		} else {
			matchMember := regexMember.FindStringSubmatch(line)
			if matchMember != nil {
				member := EntityMember{
					Name:         matchMember[4],
					Virtual:      len(matchMember[1]) > 0,
					Optional:     len(matchMember[3]) > 0,
					DataType:     matchMember[2],
					IsCollection: false,
				}
				if strings.HasPrefix(member.DataType, "ICollection") {
					regexCollection := regexp.MustCompile(`ICollection<(\w+)>`)
					matchCollection := regexCollection.FindStringSubmatch(member.DataType)
					if matchCollection != nil {
						member.DataType = matchCollection[1]
						member.IsCollection = true
					}
				}
				entity.Members = append(entity.Members, member)
			}
		}
	}
	return entity
}

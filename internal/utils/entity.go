package utils

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
	"github.com/zrs01/sct/internal/config"
)

type Entity struct {
	Name                 string
	Members              []EntityMember
	IsContainsVirtual    bool
	IsContainsCollection bool
}

type EntityMember struct {
	Name         string
	IsVirtual    bool
	IsOptional   bool
	DataType     string
	IsCollection bool
	MaxLength    int
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
					IsVirtual:    len(matchMember[1]) > 0,
					IsOptional:   len(matchMember[3]) > 0,
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

	entity.IsContainsVirtual = funk.Contains(entity.Members, func(m EntityMember) bool {
		return m.IsVirtual
	})
	entity.IsContainsCollection = funk.Contains(entity.Members, func(m EntityMember) bool {
		return m.IsCollection
	})

	// addition information
	updateEntity(&entity)
	// repr.Println(entity)
	return entity
}

func updateEntity(entity *Entity) {
	cfg := config.GetConfig()
	fileBytes, err := os.ReadFile(cfg.Dotnet.DbContext)
	if err != nil {
		panic(err)
	}
	fileContent := strings.ReplaceAll(string(fileBytes), "\n", "")

	match := regexp.MustCompile(fmt.Sprintf(`modelBuilder.Entity<%s>\(entity\s*=>\s*\{([^\}]*)\}\);`, entity.Name)).
		FindStringSubmatch(fileContent)
	if match == nil {
		return
	}
	modelBuilder := match[1]

	matchs := regexp.MustCompile(`\.Property\([a-z] => [a-z].([^\)]*)\)\.HasMaxLength\((\d+)\)`).
		FindAllStringSubmatch(modelBuilder, -1)
	if matchs != nil {
		funk.ForEach(matchs, func(x []string) {
			memberName := x[1]
			maxLength := mustAtoi(x[2])
			for i := 0; i < len(entity.Members); i++ {
				if entity.Members[i].Name == memberName {
					entity.Members[i].MaxLength = maxLength
				}
			}
		})
	}
}

func mustAtoi(s string) int {
	ans, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return ans
}

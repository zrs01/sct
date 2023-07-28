# Source Code Template

## Template Syntax
https://github.com/CloudyKit/jet/blob/master/docs/index.md

## Context data for template

```go
type Context struct {
  Data   map[interface{}]interface{}  // follow the structure of the user-defined file in YAML format
  Entity []utils.Entity               // dotnet entity class
}

type Entity struct {
	Name                 string         // name of the entity class (e.g. User)
	Members              []EntityMember // class members
	IsContainsVirtual    bool
	IsContainsCollection bool
}

type EntityMember struct {
	Name         string                 // name of member (e.g. Name)
	IsVirtual    bool
	IsOptional   bool
	DataType     string
	IsCollection bool
}

```
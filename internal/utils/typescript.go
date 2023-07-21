package utils

type TypeScript struct {
	Name    string
	Members []TypeScriptMember
}

type TypeScriptMember struct {
	Name     string
	Optional bool
	DataType string
	Value    string
}

func TypescriptParse(fullPathFileName string) TypeScript {
	var tsClass TypeScript
	csClass := CsharpParse(fullPathFileName)
	tsClass.Name = csClass.Name
	tsClass.Members = make([]TypeScriptMember, len(csClass.Members))
	for i, member := range csClass.Members {
		tsClass.Members[i].Name = member.Name
		tsClass.Members[i].DataType, tsClass.Members[i].Value = convertTypescirptType(member.DataType, member.Optional)
		if len(member.CollectionType) != 0 {
			tsClass.Members[i].DataType = member.CollectionType + "[]"
			tsClass.Members[i].Value = "null"
		}
		tsClass.Members[i].Optional = tsClass.Members[i].Value == "null"
	}
	return tsClass
}

func convertTypescirptType(itype string, optional bool) (string, string) {
	dtype := itype
	dvalue := "null"
	if itype == "int" || itype == "long" || itype == "decimal" {
		dtype = "number"
		dvalue = "0"
		if optional {
			dvalue = "null"
		}
	} else if itype == "string" {
		dvalue = "''"
		if optional {
			dvalue = "null"
		}
	} else if itype == "DateTime" {
		dtype = "Date"
		dvalue = "null"
	} else if itype == "byte[]" {
		dtype = "number[]"
		dvalue = "[]"
	} else {
		dvalue = "null"
	}
	return dtype, dvalue
}

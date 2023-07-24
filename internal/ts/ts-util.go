// typescript utility
package ts

import "sct/internal/utils"

func ToType(member utils.CsharpMember) (string, string) {
	dtype := member.DataType
	dvalue := "null"

	if len(member.CollectionType) != 0 {
		dtype = member.CollectionType + "[]"
		dvalue = "null"
	} else {
		if member.DataType == "int" || member.DataType == "long" || member.DataType == "decimal" {
			dtype = "number"
			dvalue = "0"
			if member.Optional {
				dvalue = "null"
			}
		} else if member.DataType == "string" {
			dvalue = "''"
			if member.Optional {
				dvalue = "null"
			}
		} else if member.DataType == "DateTime" {
			dtype = "Date"
			dvalue = "null"
		} else if member.DataType == "byte[]" {
			dtype = "number[]"
			dvalue = "[]"
		} else {
			dvalue = "null"
		}
	}

	return dtype, dvalue
}

// typescript utility
package ts

func ToType(dataType string, isCollection bool) string {
	if isCollection {
		return dataType + "[]"
	} else {
		if dataType == "int" || dataType == "long" || dataType == "decimal" {
			return "number"
		} else if dataType == "string" {
			return "string"
		} else if dataType == "DateTime" {
			return "Date"
		} else if dataType == "byte[]" {
			return "number[]"
		}
	}
	return dataType
}

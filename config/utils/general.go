package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"
	"unicode"
)

// WhiteSpaceTrimmer removes white spaces.
func WhiteSpaceTrimmer(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

// TimestampNow returns current timestamp.
// If format is not speciefied, it will return timestamp in RFC3339 format.
func TimestampNow(UTC bool, format string) (time.Time, string) {
	var timestampNow time.Time
	if UTC {
		timestampNow = time.Now().UTC()
	} else {
		timestampNow = time.Now().Local()
	}

	if format != "" {
		return timestampNow, timestampNow.Format(format)
	} else {
		return timestampNow, timestampNow.Format(time.RFC3339)
	}
}

// AnyToJsonStr converts any data type to JSON string if the data type itself supported
// to be converted to JSON.
func AnyToJsonStr(data interface{}) string {
	switch data := data.(type) {
	case []byte:
		return string(data)
	default:
		dataByte, err := json.Marshal(data)
		if err != nil {
			return fmt.Sprintf("%v", data)
		}
		return string(dataByte)
	}
}

// AnyToMapStringInterface converts any data type to map[string]interface{}
func AnyToMapStringInterface(data interface{}) map[string]interface{} {
	var result map[string]interface{}
	dataByte, _ := json.Marshal(data)
	json.Unmarshal(dataByte, &result)

	return result
}

// StructIterator iterates struct
func StructIterator(s interface{}) {
	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}

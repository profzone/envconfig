package envconfig

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	UnSupportTypeError = errors.New("un support type")
)

func ConvertToStr(v interface{}) (string, error) {
	if stringer, ok := v.(fmt.Stringer); ok {
		return stringer.String(), nil
	}
	rt := reflect.TypeOf(v)
	switch rt.Kind() {
	case reflect.String:
		return v.(string), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
		return fmt.Sprintf("%d", v), nil
	case reflect.Bool:
		return strconv.FormatBool(v.(bool)), nil
	case reflect.Float32:
		return strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32), nil
	case reflect.Float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64), nil
	default:
		return "", UnSupportTypeError
	}
}

func tagValueAndFlags(tagString string) (string, map[string]bool) {
	valueAndFlags := strings.Split(tagString, ",")
	v := valueAndFlags[0]
	tagFlags := map[string]bool{}
	if len(valueAndFlags) > 1 {
		for _, flag := range valueAndFlags[1:] {
			tagFlags[flag] = true
		}
	}
	return v, tagFlags
}

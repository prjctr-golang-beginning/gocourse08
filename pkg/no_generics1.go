package pkg

import (
	"fmt"
	"reflect"
)

func convertToInt(num any) (int, error) {
	switch v := num.(type) {
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		return int(v), nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	default:
		return 0, fmt.Errorf("unsupported type: %s", reflect.TypeOf(num))
	}
}

func CalcNum(a, b any) (int, error) {
	aInt, err := convertToInt(a)
	if err != nil {
		return 0, err
	}

	bInt, err := convertToInt(b)
	if err != nil {
		return 0, err
	}

	return aInt + bInt, nil
}

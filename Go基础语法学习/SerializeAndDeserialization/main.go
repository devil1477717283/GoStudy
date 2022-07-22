package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func Marshal(v interface{}) ([]byte, error) {
	value := reflect.ValueOf(v)
	typ := value.Type()
	switch typ.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Int8:
		return []byte(fmt.Sprintf("%v", value.Interface())), nil
	case reflect.String:
		return []byte(value.String()), nil
	case reflect.Bool:
		return []byte(fmt.Sprintf("%t", value.Bool())), nil
	case reflect.Float32, reflect.Float64:
		return []byte(fmt.Sprintf("%f", value.Float())), nil
	default:
		return nil, errors.New("暂不支持")
	}
}
func UnMarshal(v interface{}, data []byte) error {
	value := reflect.ValueOf(v)
	typ := value.Type()
	if typ.Kind() != reflect.Ptr {
		return errors.New("v isn't a pointer")
	}
	typ = typ.Elem()
	value = value.Elem()
	s := string(data)
	switch typ.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int8:
		i, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			value.SetInt(i)
			return nil
		}
		return err
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		i, err := strconv.ParseUint(s, 10, 64)
		if err == nil {
			value.SetUint(i)
			return nil
		}
		return err
	case reflect.String:
		value.SetString(s)
		return nil
	case reflect.Bool:
		b, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}
		value.SetBool(b)
		return nil
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return err
		}
		value.SetFloat(f)
		return nil
	default:
		return errors.New("暂不支持")
	}
}
func main() {

}

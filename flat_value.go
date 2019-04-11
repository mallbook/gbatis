package gbatis

import (
	"fmt"
	"reflect"
	"strconv"
)

func flatValue(i interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(i)
	// if v.Kind() != reflect.Struct && v.Kind() != reflect.Map {
	// 	return nil, fmt.Errorf("Expect a map or struct type, but %v", v.Kind())
	// }

	prefix := ""
	var err error
	o := make(map[string]interface{})

	switch v.Kind() {
	case reflect.Struct:
		err = flatStruct(o, prefix, i)
	case reflect.Map:
		err = flatMap(o, prefix, i)
	case reflect.Array, reflect.Slice:
		err = flatArray(o, prefix, i)
	case reflect.Ptr:
		err = flatPtr(o, prefix, i)
	default:
		if v.CanInterface() {
			o[prefix+"."] = v.Interface()
		}
	}
	return o, err
}

func flatStruct(o map[string]interface{}, prefix string, i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("Expect a struct type, but %v", v.Kind())
	}

	t := v.Type()
	n := v.NumField()
	for i := 0; i < n; i++ {
		sf := t.Field(i)

		p := prefix + "." + sf.Name

		f := v.Field(i)
		switch f.Kind() {
		case reflect.Struct:
			return flatStruct(o, p, f.Interface())
		case reflect.Map:
			return flatMap(o, p, f.Interface())
		case reflect.Array, reflect.Slice:
			return flatArray(o, p, f.Interface())
		case reflect.Ptr:
			return flatPtr(o, p, f.Interface())
		default:
			if f.CanInterface() {
				o[p] = f.Interface()
			}
		}
	}
	return nil
}

func flatMap(o map[string]interface{}, prefix string, i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Map {
		return fmt.Errorf("Expect a map type, but %v", v.Kind())
	}

	for _, k := range v.MapKeys() {
		p := prefix + "." + k.String()

		vv := v.MapIndex(k)
		switch vv.Kind() {
		case reflect.Map:
			return flatMap(o, p, vv.Interface())
		case reflect.Struct:
			return flatStruct(o, p, vv.Interface())
		case reflect.Array, reflect.Slice:
			return flatArray(o, p, vv.Interface())
		case reflect.Ptr:
			return flatPtr(o, p, vv.Interface())
		default:
			if vv.CanInterface() {
				o[p] = vv.Interface()
			}
		}
	}

	return nil
}

func flatArray(o map[string]interface{}, prefix string, i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return fmt.Errorf("Expect an array or a slice type, but %v", v.Kind())
	}

	l := v.Len()
	for ind := 0; ind < l; ind++ {
		item := v.Index(ind)
		p := prefix + "." + strconv.Itoa(ind)
		switch item.Kind() {
		case reflect.Struct:
			return flatArray(o, p, item.Interface())
		case reflect.Map:
			return flatMap(o, p, item.Interface())
		case reflect.Array, reflect.Slice:
			return flatArray(o, p, item.Interface())
		case reflect.Ptr:
			return flatPtr(o, p, item.Interface())
		default:
			if item.CanInterface() {
				o[p] = item.Interface()
			}
		}
	}
	return nil
}

func flatPtr(o map[string]interface{}, prefix string, i interface{}) error {
	v := reflect.ValueOf(i)
	if v.Kind() != reflect.Ptr {
		return fmt.Errorf("Expect a ptr type, but %s", v.Kind())
	}

	elem := v.Elem()

	switch elem.Kind() {
	case reflect.Struct:
		return flatStruct(o, prefix, elem.Interface())
	default:
		if elem.CanInterface() {
			o[prefix+"."] = elem.Interface()
		}
	}
	return nil
}

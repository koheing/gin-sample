package utils

import "reflect"

func GetTypeName(instance interface{}) string {
	t := reflect.TypeOf(instance)
	if t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	}
	return t.Name()
}

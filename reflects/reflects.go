package reflects

import "reflect"

func IsStruct(obj interface{}) bool {

	return reflect.TypeOf(obj).Kind() == reflect.Struct
}
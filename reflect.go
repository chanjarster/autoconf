package autoconf

import (
	"reflect"
)

// Initialize p's field. If field is exported and is a pointer,
// that field will be initialized to new(Type)
func initStruct(p interface{}) {
	initPtrValue(reflect.ValueOf(p), "")
}

// param path is just for debug
func initPtrValue(v reflect.Value, path string) {

	if v.Kind() == reflect.Ptr {
		if t := v.Type().Elem(); v.IsNil() && v.CanSet() {
			v.Set(reflect.New(t))
		}
		v = v.Elem()
	}

	if !v.CanSet() {
		return
	}

	if v.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < v.NumField(); i++ {
		fv := v.Field(i)
		f := v.Type().Field(i)

		if path == "" {
			initPtrValue(fv, f.Name)
		} else {
			initPtrValue(fv, path+"."+f.Name)
		}
	}

}

// Field visitor
//
// args:
//   path: dot separated field path
//   v: value of the field
type visitor func(path string, f reflect.StructField, v reflect.Value)

// Visit all export fields, support fields of type:
//   bool
//   time.Duration TODO
//   float64
//   int
//   int64
//   string
//   uint
//   uint64
//   pointer to above type
// If field is a struct or a pointer to a struct, it will be
// recursively visited
func visitExportedFields(p interface{}, fn visitor) {

	v := reflect.ValueOf(p)

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}

	if !v.CanSet() {
		return
	}
	if v.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		fv := v.Field(i)
		visitExportedFieldsPath(f, fv, f.Name, fn)
	}

}

func visitExportedFieldsPath(f reflect.StructField, v reflect.Value, path string, fn visitor) {

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}

	if !v.CanSet() {
		return
	}

	switch k := v.Kind(); k {

	case reflect.Bool, reflect.Float64, reflect.Int,
		reflect.Int64, reflect.String, reflect.Uint, reflect.Uint64:
		fn(path, f, v)

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			f := v.Type().Field(i)
			visitExportedFieldsPath(f, fv, path+"."+f.Name, fn)
		}

	default:

	}

}

package autoconf

import (
	"flag"
	"os"
	"reflect"
)

type flagResolver struct {
	args    []string
	flagSet *flag.FlagSet
}

func (r *flagResolver) init(p interface{}) {

	if r.args == nil {
		r.args = os.Args[1:]
	}

	if r.flagSet == nil {
		r.flagSet = flag.NewFlagSet("flag", flag.ContinueOnError)
	}

	flagSet := r.flagSet

	visitExportedFields(p, func(path string, f reflect.StructField, v reflect.Value) {

		flagName := flagStyle(path)

		switch k := v.Type().Kind(); k {
		case reflect.Bool:
			flagSet.BoolVar(v.Addr().Interface().(*bool), flagName, v.Bool(), path)
		case reflect.Float64:
			flagSet.Float64Var(v.Addr().Interface().(*float64), flagName, v.Float(), path)
		case reflect.Int:
			flagSet.IntVar(v.Addr().Interface().(*int), flagName, int(v.Int()), path)
		case reflect.Int64:
			flagSet.Int64Var(v.Addr().Interface().(*int64), flagName, v.Int(), path)
		case reflect.String:
			flagSet.StringVar(v.Addr().Interface().(*string), flagName, v.String(), path)
		case reflect.Uint:
			flagSet.UintVar(v.Addr().Interface().(*uint), flagName, uint(v.Uint()), path)
		case reflect.Uint64:
			flagSet.Uint64Var(v.Addr().Interface().(*uint64), flagName, v.Uint(), path)
		default:
		}
	})

}
func (r *flagResolver) Resolve(p interface{}) error {

	flagSet := r.flagSet
	return flagSet.Parse(r.args)
}
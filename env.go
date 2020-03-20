package autoconf

import (
	"flag"
	"fmt"
	"os"
	"reflect"
)

type envResolver struct {
	environ []string
	flagSet *flag.FlagSet
}

func (r *envResolver) init(p interface{}) {

	if r.environ == nil {
		r.environ = make([]string, 0)
	}

	r.flagSet = flag.NewFlagSet("environ", flag.ContinueOnError)

	flagSet := r.flagSet

	visitExportedFields(p, func(path string, f reflect.StructField, v reflect.Value) {

		env := envStyle(path)
		r.environ = append(r.environ, env)

		switch k := v.Type().Kind(); k {
		case reflect.Bool:
			flagSet.BoolVar(v.Addr().Interface().(*bool), env, v.Bool(), path)
		case reflect.Float64:
			flagSet.Float64Var(v.Addr().Interface().(*float64), env, v.Float(), path)
		case reflect.Int:
			flagSet.IntVar(v.Addr().Interface().(*int), env, int(v.Int()), path)
		case reflect.Int64:
			flagSet.Int64Var(v.Addr().Interface().(*int64), env, v.Int(), path)
		case reflect.String:
			flagSet.StringVar(v.Addr().Interface().(*string), env, v.String(), path)
		case reflect.Uint:
			flagSet.UintVar(v.Addr().Interface().(*uint), env, uint(v.Uint()), path)
		case reflect.Uint64:
			flagSet.Uint64Var(v.Addr().Interface().(*uint64), env, v.Uint(), path)
		default:
		}

	})

}

func (r *envResolver) Resolve(p interface{}) error {

	flagSet := r.flagSet
	for _, env := range r.environ {
		v, ok := os.LookupEnv(env)
		if !ok {
			continue
		}
		if err := flagSet.Set(env, v); err != nil {
			return r.failf("invalid value %q for env %s: %v", v, env, err)
		}
	}
	return nil
}

func (r *envResolver) failf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	fmt.Fprintln(r.flagSet.Output(), err)
	return err
}

package autoconf

import (
	"reflect"
	"testing"
)

func TestYamlFileResolver_Resolve(t *testing.T) {

	t.Run("exported fields", func(t *testing.T) {

		o := &outer{}
		initStruct(o)

		y := &yamlFileResolver{
			File: "outer.yaml",
		}

		y.Resolve(o)

		var e_b = true
		var e_f float64 = 3
		var e_i = 3
		var e_i64 int64 = 3
		var e_s = "foo"
		var e_uint uint = 3

		want := &outer{
			B:     e_b,
			Bp:    &e_b,
			F:     e_f,
			Fp:    &e_f,
			I:     e_i,
			Ip:    &e_i,
			I64:   e_i64,
			I64p:  &e_i64,
			S:     e_s,
			Sp:    &e_s,
			Uint:  e_uint,
			Uintp: &e_uint,
			Inner: inner{
				I:  e_i,
				Ip: &e_i,
			},
			Innerp: &inner{
				I:  e_i,
				Ip: &e_i,
			},
			Inner2:  struct {
				I  int
				Ip *int
			}{
				I:  e_i,
				Ip: &e_i,
			},
		}

		if !reflect.DeepEqual(o, want) {
			t.Errorf("o = %v, want %v", o, want)
		}

	})

	t.Run("unexported fields", func(t *testing.T) {

		f := &foo{}
		initStruct(f)

		y := &yamlFileResolver{
			File: "foo.yaml",
		}

		y.Resolve(f)

		i3 := 3
		want := &foo{
			i:  0,
			ip: nil,
			bar: bar{
				i:  0,
				ip: nil,
				I:  0,
				Ip: nil,
			},
			barP: nil,
			I:    3,
			Ip:   &i3,
			Bar: bar{
				i:  0,
				ip: nil,
				I:  3,
				Ip: &i3,
			},
			BarP: &bar{
				i:  0,
				ip: nil,
				I:  3,
				Ip: &i3,
			},
			baz: struct {
				i int
			}{
				i: 0,
			},
			Baz: struct {
				i int
			}{
				i: 0,
			},
		}

		if !reflect.DeepEqual(f, want) {
			t.Errorf("f = %v, want %v",f, want)
		}

	})

}
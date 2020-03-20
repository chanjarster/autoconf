package autoconf

//---
// struct with exported fields

type inner struct {
	I  int
	Ip *int
}

type outer struct {
	B  bool
	Bp *bool

	F  float64
	Fp *float64

	I  int
	Ip *int

	I64  int64
	I64p *int64

	S  string
	Sp *string

	Uint  uint
	Uintp *uint

	Inner  inner
	Innerp *inner

	Inner2 struct {
		I  int
		Ip *int
	}
}

//-----
// struct with exported and unexported fields
type bar struct {
	i  int
	ip *int

	I  int
	Ip *int
}

type foo struct {
	i    int
	ip   *int
	bar  bar
	barP *bar

	I    int
	Ip   *int
	Bar  bar
	BarP *bar

	baz struct {
		i int
	}

	Baz struct {
		i int
	}
}

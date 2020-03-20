package autoconf

// Foo.BarBaz -> foo-bar-baz
func flagStyle(path string) string {
	bs := []byte(path)
	rs := make([]byte, 0, len(bs))

	prevDot := true
	for i := 0; i < len(bs); i++ {
		b := bs[i]
		if isUpperASCII(b) {
			if !prevDot {
				rs = append(rs, '-')
			}
			rs = append(rs, lowerASCII(b))
			prevDot = false
		} else if b == '.' {
			rs = append(rs, '-')
			prevDot = true
		} else {
			rs = append(rs, b)
			prevDot = false
		}
	}

	return string(rs)
}

// Foo.BarBaz -> FOO_BAR_BAZ
func envStyle(path string) string {

	bs := []byte(path)
	rs := make([]byte, 0, len(bs))

	prevDot := true
	for i := 0; i < len(bs); i++ {
		b := bs[i]
		if isUpperASCII(b) {
			if !prevDot {
				rs = append(rs, '_')
			}
			rs = append(rs, b)
			prevDot = false
		} else if b == '.' {
			rs = append(rs, '_')
			prevDot = true
		} else {
			rs = append(rs, upperASCII(b))
			prevDot = false
		}
	}

	return string(rs)

}

func lowerASCII(b byte) byte {
	if isUpperASCII(b) {
		return b + ('a' - 'A')
	}
	return b
}

func upperASCII(b byte) byte {
	if isLowerASCII(b) {
		return b - ('a' - 'A')
	}
	return b
}

func isLowerASCII(b byte) bool {
	return 'a' <= b && b <= 'z'
}

func isUpperASCII(b byte) bool {
	return 'A' <= b && b <= 'Z'
}

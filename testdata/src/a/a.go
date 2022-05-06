package a

func a_() { // want "a_ is used under score. You should use mixedCap or MixedCap"
}

func b(a_a int) { // want "a_a is used under score. You should use mixedCap or MixedCap."
}

func c() (c_c int) { // want "c_c is used under score. You should use mixedCap or MixedCap."
	c_c = 1    // want "c_c is used under score. You should use mixedCap or MixedCap."
	return c_c // want "c_c is used under score. You should use mixedCap or MixedCap."
}

func d() {
	var d_d int // want "d_d is used under score. You should use mixedCap or MixedCap."
	_ = d_d     // It's never detected, because `_` is meaningful in Go and `d_d` is already detected.
}

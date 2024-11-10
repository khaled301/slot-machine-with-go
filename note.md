	name := "" // := is called "short variable declaration" and type is determined by value
	fmt.Println(name)

	var name2 string
	fmt.Println(name2)

	var age int8
	fmt.Println(age)

	var id int64
	fmt.Println(id)

	// unsigned | positive
	var id2 uint8
	fmt.Println(id2)

	// float
	var pi float32
	fmt.Println(pi)

	// boolean
	var isTrue bool
	fmt.Println(isTrue)

	// array
	var names [3]string
	fmt.Println(names)

	// complex | explanation = real + imaginary i | i = sqrt(-1)
	var complexNumber complex128
	fmt.Println(complexNumber)

	// byte | alias for uint8 | 0 to 255
	var b byte
	fmt.Println(b)

	// rune | alias for int32 | 0 to 2^32 - 1
	var r rune
	fmt.Println(r)

	// map | key => value | key must be unique and immutable | value can be any type
	var m map[string]int
	fmt.Println(m)

	// composite literal
	var s = []int{1, 2, 3} // slice of integers of length 3 with values 1, 2, 3 | slice is a reference type
	fmt.Println(s)

	var strs [3]string | fixed sized array
	var strSlice []string | Slice | resizable array | flexible version of fixed size array | can automatically resize
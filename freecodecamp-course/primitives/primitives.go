package primitives

import "fmt"

var aBoolean bool

// signed integers
var i int = 42
var i8 int8 = 42
var i16 int16 = 42
var i32 int32 = 42
var i64 int64 = 42

// unsigned integers
var ui uint = 42
var ui8 uint8 = 42
var ui16 uint16 = 42
var ui32 uint32 = 42
var ui64 uint64 = 42

func Primitives() {
	fmt.Printf("%v, %T", aBoolean, aBoolean)

	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", i8, i8)
	fmt.Printf("%v, %T\n", i16, i16)
	fmt.Printf("%v, %T\n", i32, i32)
	fmt.Printf("%v, %T\n", i64, i64)

	fmt.Printf("%v, %T\n", ui, ui)
	fmt.Printf("%v, %T\n", ui8, ui8)
	fmt.Printf("%v, %T\n", ui16, ui16)
	fmt.Printf("%v, %T\n", ui32, ui32)
	fmt.Printf("%v, %T\n", ui64, ui64)

	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

	// -- bit-shifting

	c := 8              // 2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0

	// floating point numbers
	n := 3.14
	n = 3.7e72
	n = 2.1e14
	fmt.Printf("%v, %T\n", n, n)

	// complex types
	var c64Declaration complex64 = complex(1, 2)
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 1 + 2i

	fmt.Printf("%v, %T\n", c64Declaration, c64Declaration)

	fmt.Printf("%v, %T\n", c64, c64)
	fmt.Printf("%v, %T\n", real(c64), real(c64))
	fmt.Printf("%v, %T\n", imag(c64), imag(c64))
	fmt.Printf("%v, %T\n", c128, c128)

	// Text -> String (any UTF-8 character)
	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])

	b2 := []byte(s)
	fmt.Printf("%v, %T\n", b2, b2)

	// Text -> Rune (UTF-32) (type alias int32)
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)

}

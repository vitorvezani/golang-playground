package constants

import "fmt"

func Constants() {

	naming()
	typed()
	untyped()
	enumeratedConstants()
	enumeratedExpressions()

}

func naming() {
	// NEEDs to be evaluated at compile time
	// shadowing is also available
	const myConst int = 42
	fmt.Printf("%v, %T", myConst, myConst)
}

func typed() {
	const a = 32
	var b int16 = 16
	fmt.Printf("%v, %T", a+b, a+b)
}

const (
	a = iota
	b = iota
	c = iota
)

func untyped() {
	fmt.Printf("%v, %T", a, a)
	fmt.Printf("%v, %T", b, b)
	fmt.Printf("%v, %T", c, c)
}

func enumeratedConstants() {

}

func enumeratedExpressions() {

}

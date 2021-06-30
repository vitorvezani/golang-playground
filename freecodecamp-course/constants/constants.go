package constants

import "fmt"

/*
* - Immutable but can be shadowed
* - Replaced by the compiler at compile time
* - Named like variables
* - Can be Type and Untyped (works like literal values)
* - enumerated constants (iota -> be careful of 0 though)
* - enumerated expressions (arithmetic, bitwise and bitshifting)
 */
func init() {
	fmt.Println("Init primitive file executed")
}

func Constants() {
	naming()
	typed()
	untyped()
	enumeratedConstants()
	enumeratedExpressions()
	checkPrivileges()
}

func naming() {
	const myConst int = 42
	fmt.Printf("%v, %T\n", myConst, myConst)
}

func typed() {
	const a = 32
	var b int16 = 16
	fmt.Printf("%v, %T\n", a+b, a+b)
}

const (
	isAdmin = 1 << iota
	isHQ
	canSeeFinancial
)

func checkPrivileges() {
	var roles byte = isAdmin | canSeeFinancial
	fmt.Printf("is Admin? %v\n", isAdmin&roles == isAdmin)
}

const (
	_ = iota
	a
	b
	c
)

func untyped() {
	fmt.Printf("a: %v, %T\n", a, a)
	fmt.Printf("b: %v, %T\n", b, b)
	fmt.Printf("c: %v, %T\n", c, c)
}

func enumeratedConstants() {

}

func enumeratedExpressions() {

}

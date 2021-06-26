package main

import "fmt"

func constants() {

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

}

func untyped() {

}

func enumeratedConstants() {

}

func enumeratedExpressions() {

}

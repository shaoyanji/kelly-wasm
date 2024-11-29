package main

func main() {
}

// This function is exported to JavaScript, so can be called using
// exports.kc() in JavaScript.
//
//export kc
func kc(p int, re int, ri int) int {
	return p - (100*ri-p*ri)/re
}

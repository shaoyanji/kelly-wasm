package main

func main() {
}

// This function is exported to JavaScript, so can be called using
// exports.kc() in JavaScript.
//export kc 
func kc(x, y int) int {
    return x - (100-x)/y
}

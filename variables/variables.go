package variables

var greeting string = "Hello, World!"
var number int = 33
var isTrue bool = true
var floatNumber float64 = 3.14
var a, b, c, d, e int = 1, 2, 3, 4, 5
const PI float64 = 3.14159
const (
	A int = 1
	B float32 = 3.14
	C string = "Hi!"
	D bool = true
	E = 3
	F = 4.555
	G = "Hello"
  )

  func Change() {
	greeting = "Hello, Universe!"
	println(greeting)
	  }
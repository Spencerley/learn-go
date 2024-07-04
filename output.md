# Output

- Go has three functions to output text:
    1. Print()
    2. Println()
    3. Printf()

## The Print() Function
- The Print() function prints its arguments with their default format.
- If we want to print the arguments in new lines, we need to use \n.
- It is also possible to only use one Print() for printing multiple variables.
- If we want to add a space between string arguments, we need to use " ".
- Print() inserts a space between the arguments if neither are strings.

## The Println() Function
- The Println() function is similar to Print() with the difference that a whitespace is added between the arguments, and a newline is added at the end.
  
## The Printf() function 
- The Printf() function first formats its argument based on the given formatting verb and then prints them.

Here we will use two formatting verbs:

- %v is used to print the value of the arguments
- %T is used to print the type of the arguments

`package main`
`import ("fmt")`

`func main() {`
  `var i string = "Hello"`
  `var j int = 15`

  `fmt.Printf("i has value: %v and type: %T\n",` `i, i)`
  `fmt.Printf("j has value: %v and type: %T", j,` `j)`
`}`
Result:
- i has value: Hello and type: string
- j has value: 15 and type: int

## Formatting Verbs for Printf()
- Go offers several formatting verbs that can be used with the Printf() function.
  
## General Formatting Verbs
The following verbs can be used with all data types:

| Verb	| Description |
|%v	    |  Prints the value in the default format |
|%#v    |	Prints the value in Go-syntax format |
|%T     |	Prints the type of the value |
|%%     |	Prints the % sign |

## Integer Formatting Verbs
The following verbs can be used with the integer data type:

|Verb	| Description |
| --- | --- |
|%b |	Base 2 |
|%d |	Base 10 |
|%+d | 	Base 10 and always show sign |
|%o |	Base 8 |
|%O |	Base 8, with leading 0o |
|%x |	Base 16, lowercase |
|%X |	Base 16, uppercase |
|%#x |	Base 16, with leading 0x |
|%4d |	Pad with spaces (width 4, right justified) |
|%-4d |	Pad with spaces (width 4, left justified) |
|%04d |	Pad with zeroes (width 4) |

## String Formatting Verbs
The following verbs can be used with the string data type:

|Verb |	Description |
| --- | --- |
|%s |	Prints the value as plain string |
|%q |	Prints the value as a double-quoted string |
|%8s |	Prints the value as plain string (width 8, right justified) |
|%-8s |	Prints the value as plain string (width 8, left justified) |
|%x |	Prints the value as hex dump of byte values |
|% x |	Prints the value as hex dump with spaces |

## Boolean Formatting Verbs
The following verb can be used with the boolean data type:

| Verb |	Description |
| --- | --- |
| %t | 	Value of the boolean operator in true or false format (same as using %v) |

## Float Formatting Verbs
The following verbs can be used with the float data type:

| Verb |	Description |
| --- | --- |
| %e |	Scientific notation with 'e' as exponent |
| %f |	Decimal point, no exponent |
| %.2f |	Default width, precision 2 |
| %6.2f |	Width 6, precision 2 |
| %g |	Exponent as needed, only necessary digits |
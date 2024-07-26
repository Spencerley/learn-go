 # GO Conditions

Conditional statements are used to perform different actions based on different conditions.

A condition can be either true or false.

Go supports the usual comparison operators from mathematics:

- Less than <
- Less than or equal <=
- Greater than >
- Greater than or equal >=
- Equal to ==
- Not equal to !=

Additionally, Go supports the usual logical operators:

- Logical AND &&
- Logical OR ||
- Logical NOT !

Go has the following conditional statements:

- Use if to specify a block of code to be executed, if a specified condition is true
- Use else to specify a block of code to be executed, if the same condition is false
- Use else if to specify a new condition to test, if the first condition is false
-  Use switch to specify many alternative blocks of code to be executed

## The if Statement
- Use the if statement to specify a block of Go code to be executed if a condition is true.
  `if condition {
  // code to be executed if condition is true
}`

## The else Statement
- Use the else statement to specify a block of code to be executed if the condition is false.
`if condition {
  // code to be executed if condition is true
} else {
  // code to be executed if condition is false
}`
- The brackets in the else statement should be like `} else {` otherwise it will fail.

## The else if Statement
- Use the else if statement to specify a new condition if the first condition is false.
`if condition1 {
   // code to be executed if condition1 is true
} else if condition2 {
   // code to be executed if condition1 is false and condition2 is true
} else {
   // code to be executed if condition1 and condition2 are both false
}`

## The Nested if Statement
- You can have if statements inside if statements, this is called a nested if.
`if condition1 {
   // code to be executed if condition1 is true
  if condition2 {
     // code to be executed if both condition1 and condition2 are true
  }
}`

`package main
import ("fmt")
func main() {
  num := 20
  if num >= 10 {
    fmt.Println("Num is more than 10.")
    if num > 15 {
      fmt.Println("Num is also more than 15.")
     }
  } else {
    fmt.Println("Num is less than 10.")
  }
}`
# Go Operators
- Operators are used to perform operations on variables and values.
- Go divides the operators into the following groups:
  - Arithmetic operators
  - Assignment operators
  - Comparison operators
  - Logical operators
  - Bitwise operators

## Arithmetic Operators
Arithmetic operators are used to perform common mathematical operations.

| Operator |	Name |	Description |	Example | 
| --- | --- | --- | --- | 
| + |	Addition |	Adds together two values |	x + y | 
| - |	Subtraction |	Subtracts one value from another |	x - y | 
| * |	Multiplication |	Multiplies two values |	x * y |	 
| / |	Division |	Divides one value by another |	x / y	| 
| % |	Modulus |	Returns the division remainder |	x % y	| 
| ++ |	Increment |	Increases the value of a variable by 1 |	x++	| 
| -- | 	Decrement | 	Decreases the value of a variable by 1 |	x-- | 

## Assignment Operators
Assignment operators are used to assign values to variables.
| Operator |	Example |	Same As | 
| --- | --- | --- |	
| = |	x = 5 |	x = 5 |
| += | 	x += 3 |	x = x + 3 |
| -= | 	x -= 3 |	x = x - 3 |
| *= | 	x *= 3 |	x = x * 3 |
| /= |	x /= 3 |	x = x / 3 |
| %= |	x %= 3 |	x = x % 3 |	
| &= |	x &= 3 |	x = x & 3 |	
| |= |	x |= 3 |	x = x | 3 | 
| ^= |	x ^= 3 |	x = x ^ 3 |	
| >>= |	x >>= 3 |	x = x >> 3 |
| <<= |	x <<= 3 |	x = x << 3 |

- blank line = `|= x |= 3 x = x| 3`

## Comparison Operators
- Comparison operators are used to compare two values.
- Note: The return value of a comparison is either true (1) or false (0).
| Operator |	Name |	Example |
| --- | --- | --- |
| == |	Equal to |	x == y |
| != |	Not equal |	x != y |
| > |	Greater than |	x > y |
| < |	Less than | 	x < y |
| >= |	Greater than or equal to | 	x >= y |
| <= |	Less than or equal to | 	x <= y |

## Logical Operators
Logical operators are used to determine the logic between variables or values:

| Operator | Name | Description | Example |
| --- | --- | --- | --- |	
| && |  	Logical and | 	Returns true if both statements are true | 	x < 5 &&  x < 10 |
| || | 	Logical or |  	Returns true if one of the statements is true | 	x < 5 || x < 4 |
| ! | 	Logical not |	Reverse the result, returns false if the result is true |	!(x < 5 && x < 10) |

- Blank line is: `|| 	Logical or	Returns true if one of the statements is true	x < 5 || x < 4`

## Bitwise Operators
Bitwise operators are used on (binary) numbers:
| Operator | Name | Description | Example |
| --- | --- | --- | --- |
| & |  	AND |	Sets each bit to 1 if both bits are 1 |	x & y |
| |	OR |	Sets each bit to 1 if one of two bits is 1 |	x | y |	
| ^ |	XOR |	Sets each bit to 1 if only one of two bits is 1 |	x ^ b |
| << |	Zero fill left shift |	Shift left by pushing zeros in from the right |	x << 2 |	
| >> |	Signed right shift |	Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits fall off |	x >> 2 |

- Blank line is: `|	OR	Sets each bit to 1 if one of two bits is 1	x | y`
# Arrays

## Go Arrays
- Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.
- The length specifies the number of elements to store in the array. In Go, arrays have a fixed length. The length of the array is either defined by a number or is inferred (means that the compiler decides the length of the array, based on the number of values).

## Array Initialization
- If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.
- Tip: The default value for int is 0, and the default value for string is "".
- Initialize Only Specific Elements:
    - It is possible to initialize only specific elements in an array.
    - [index number]:[value], i.e. 0:1

## Find the Length of an Array
- The len() function is used to find the length of an array: `fmt.Println(len(arr1))`


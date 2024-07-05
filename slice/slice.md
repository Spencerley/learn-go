# Go Slices
- Slices are similar to arrays, but are more powerful and flexible.

- Like arrays, slices are also used to store multiple values of the same type in a single variable.

- However, unlike arrays, the length of a slice can grow and shrink as you see fit.

- In Go, there are several ways to create a slice:
    - Using the []datatype{values} format
    - Create a slice from an array
    - Using the make() function

## Create a Slice With []datatype{values}
- Syntax: 
    - `slice_name := []datatype{values}`
    - `myslice := []int{1,2,3}`

- In Go, there are two functions that can be used to return the length and capacity of a slice:
    - len() function - returns the length of the slice (the number of elements in the slice)
    - cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

## Create a Slice From an Array
- You can create a slice by slicing an array:
- Syntax:
    - `var myarray = [length]datatype{values} // An array`
    - `myslice := myarray[start:end] // A slice made from the array`

## Create a Slice With The make() Function
- The make() function can also be used to create a slice.
- Syntax
    - `slice_name := make([]type, length, capacity)`
- Note: If the capacity parameter is not defined, it will be equal to length.

## Go Access, Change, Append and Copy Slices

### Access Elements of a Slice
- You can access a specific slice element by referring to the index number.
- In Go, indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.

### Change Elements of a Slice
- You can also change a specific slice element by referring to the index number.
- Syntax:
  - `prices := []int{10,20,30}`
  - `prices[2] = 50`

### Append Elements To a Slice
- You can append elements to the end of a slice using the append()function:
  - `slice_name = append(slice_name, element1, element2, ...)`
  - `slice_name = append(slice_name, 1, 2)`

### Append One Slice To Another Slice
- To append all the elements of one slice to another slice, use the append()function:
- Syntax
  - `slice3 = append(slice1, slice2...)`
  - Note: The '...' after slice2 is necessary when appending the elements of one slice to another.

### Change The Length of a Slice
- Unlike arrays, it is possible to change the length of a slice.
- `myslice1 = arr1[1:3] // Change length by re-slicing the array`
- ` myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items`

## Memory Efficiency
 - When using slices, Go loads all the underlying elements into the memory.
- If the array is large and you need only a few elements, it is better to copy those elements using the copy() function.
- The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program. 
- Syntax
    - `copy(dest, src)`
    - The copy() function takes in two slices dest and src, and copies data from src to dest. It returns the number of elements copied.
    - See ./memory.go
- The capacity of the new slice is now less than the capacity of the original slice because the new underlying array is smaller.


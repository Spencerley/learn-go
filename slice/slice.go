package slice
import "fmt"

var myslice = []int{1,2,3}

func Slice() {
	arr1 := [6]int{10, 11, 12, 13, 14,15}
	myslice1 := arr1[2:4]

	myslice3 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))
  
	// with omitted capacity
	myslice2 := make([]int, 5)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))
  
	fmt.Printf("myslice = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
	fmt.Println(myslice)
	fmt.Println(myslice[0])
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	myslice4 := append(myslice1, myslice2...)

	fmt.Printf("myslice3=%v\n", myslice4)
}
package main

import "fmt"

// for -> only for loop

func main() {

	// while loop
	// i := 0
	// for i < 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }

	// for loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// infinite loop
	// for {
	// 	fmt.Println("infinite loop")
	// }

	// break
	// for i := 0; i < 5; i++ {
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// continue
	// for i := 0; i < 5; i++ {
	// 	if i == 2 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// for range
	for i := range 5{
		fmt.Println(i)
	}

}

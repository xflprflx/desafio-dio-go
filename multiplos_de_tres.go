package main

import "fmt"

func main() {

	i := 1
	for i <= 100 {
		if i%3 == 0 {
			fmt.Println(i)
		}
		i++
	}

	// ou
	// for i := 0; i <= 100; i++ {
	// 	if i%3 == 0 {
	// 		fmt.Println(i)
	// 	}
	// 	i++
	// }
}

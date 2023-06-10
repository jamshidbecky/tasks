package main

import (
	"fmt"
)

// func main() {
// 	var number int
// 	var qoldiq, revNumber string
// 	fmt.Scanln(&number)

// 	for number > 0 {
// 		count := 0
// 		qoldiq1 := number % 10
// 		number = number / 10
// 		qoldiq = strconv.Itoa(qoldiq1)
// 		revNumber += qoldiq
// 		count += 1
// 	}

// 	revNumber1, _ := strconv.Atoi(revNumber)

// 	fmt.Println(revNumber1)
	
// }

func main() {
	var str1, str2 string

	fmt.Scanln(&str1)
	fmt.Scanln(&str2)

	res := str1[:len(str1) / 2] + str2 + str1[:len(str1) / 2]

	fmt.Println(res)

}

package array

import "fmt"

func Array() {
	// var num, i int
	// fmt.Print("Enter The Number : ")
	// fmt.Scanln(&num)

	// dupArr := make([]string, num)
	// fmt.Print("Enter the String Array Items  = ")

	// for i = 0; i < num; i++ {
	// 	fmt.Scanln(&dupArr[i])
	// }
	// fmt.Println(dupArr)

	/* ------------------------ */

	var num [5]string
	for i := 0; i < 5; i++ {
		fmt.Scanln(&num[i])
	}
	fmt.Println(num)
	/* ------------------------ */
	// var num int
	// start := "W"
	// end := "w!"
	// fmt.Scanln(&num)
	// allO := strings.Repeat("o", num)

	// fmt.Println(start + allO + end)
}

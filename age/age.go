package age

import "fmt"

func Age() {
	// var num1, num2 int
	// fmt.Print("Enter the First Number = ")
	// fmt.Scanln(&num1)

	// fmt.Print("Enter the Second Number = ")
	// fmt.Scanln(&num2)

	// fmt.Println("The Sum of num1 and num2  = ", num1+num2)

	// var eonum int

	// fmt.Print("Enter the Number to find Even or Odd = ")
	// fmt.Scanln(&eonum)

	// if eonum%2 == 0 {
	// 	fmt.Println("\nIt is an Even Number")
	// } else {
	// 	fmt.Println("\nIt is an Odd Number")
	// }

	// var num1 int
	// var num2 int
	// var operation string

	// fmt.Print("Enter the Number: ")
	// fmt.Scanln(&num1)

	// fmt.Print("Enter the Number: ")
	// fmt.Scanln(&num2)

	// fmt.Print("Enter the Number: ")
	// fmt.Scanln(&operation)

	// switch operation {

	// case "*":
	// 	fmt.Println(num1 * num2)
	// case "+":
	// 	fmt.Println(num1 + num2)
	// case "-":
	// 	fmt.Println(num1 - num2)
	// case "/":
	// 	fmt.Println(num1 / num2)
	// }

	// for i := 0; i < 100; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	var num, i, user, sum int

	fmt.Print("Enter the Number: ")
	fmt.Scanln(&num)

	for i = 0; i < num; i++ {
		fmt.Scanln(&user)
		sum = sum + user
	}
	fmt.Println(sum)
}

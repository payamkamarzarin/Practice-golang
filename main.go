package main

import (
	"fmt"
	"sync"
)

// func sumDigits(number int) int {
// 	remainder := 0
// 	sumResult := 0
// 	for number != 0 {
// 		remainder = number % 10
// 		sumResult += remainder
// 		number = number / 10
// 	}
// 	return sumResult
// }
/*--------------------------*/
// type Company struct {
// 	ID     int
// 	Name   string
// 	Salary int
// }

// type Employee struct {
// 	FirstName    string
// 	LastName     string
// 	Age          int
// 	NationalCode string
// 	Company      Company
// }
/*--------------------------*/
// func updatearray(funarr []int) {
// 	funarr[4] = 750
// }
/*--------------------------*/

func main() {

	// if 10 < 5 {
	// 	fmt.Println("Yes")
	// } else if 10 > 2 {
	// 	fmt.Println("No")
	// } else {
	// 	fmt.Println("Error")
	// }

	// // fmt.Println("Hello world")
	// // calculator.Sum()
	// // age.Age()
	// array.Array()

	/*---------------------------*/

	// var num int
	// start := "W"
	// end := "w!"
	// fmt.Scanln(&num)
	// allO := strings.Repeat("o", num)

	// fmt.Println(start + allO + end)

	/*--------------------------*/

	// var num int
	// fmt.Scanln(&num)
	// x := strings.Repeat("man khoshghlab hastam \n", num)
	// fmt.Println(x)

	/*--------------------------*/

	// var num int
	// fmt.Scanln(&num)

	// if num < 0 {
	// 	fmt.Println("Ice")
	// } else if num > 100 {
	// 	fmt.Println("Steam")
	// } else {
	// 	fmt.Println("Water")
	// }

	/*--------------------------*/

	// var number int
	// fmt.Print("Enter Number:")
	// fmt.Scanln(&number)
	// fmt.Println(sumDigits(number))

	/*--------------------------*/
	// var x int
	// var emp []Employee
	// var tmp Employee
	//var uuid = strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	// fmt.Scanln(&x)
	// for i := 0; i < x; i++ {
	// 	fmt.Scanln(&tmp.FirstName)
	// 	fmt.Scanln(&tmp.LastName)
	// 	fmt.Scanln(&tmp.Age)
	// 	fmt.Scanln(&tmp.NationalCode)
	// 	fmt.Scanln(&tmp.Company.ID)
	// 	fmt.Scanln(&tmp.Company.Name)
	// 	fmt.Scanln(&tmp.Company.Salary)

	// 	emp = append(emp, tmp)
	// }
	// fmt.Println(&emp)

	// var max int
	// var richestperson string
	// for _, Employee := range emp {
	// 	salary := Employee.Company.Salary
	// 	if salary > max {
	// 		max = salary
	// 		richestperson = Employee.FirstName
	// 	}
	// }
	// fmt.Println(richestperson, max)
	/*--------------------------*/
	// var number int

	// fmt.Scanln(&number)
	// fmt.Println(number)
	// fmt.Println(&number)
	/*--------------------------*/
	// arr := []int{78, 89, 45, 56, 14}
	// arr2 := make([]int, 5)
	// copy(arr2, arr)
	// updatearray(arr2)
	// fmt.Println(arr)
	/*--------------------------*/
	// 	var number1 int
	// 	var number2 []string
	// 	fmt.Scanln(&number1)
	// 	for i := 1; i <= number1; i++ {
	// 		fmt.Scanln(&number2)
	// 	}
	// 	fmt.Println(number2)
	/*--------------------------*/
	// fac.Factorial()
	/*--------------------------*/
	//file.File()
	/*--------------------------*/
	// type Person struct {
	// 	FirstName string
	// 	LastName  string
	// 	Age       int
	// }
	// person1 := func(firstname string, lastName string) Person {
	// 	fmt.Println("we are calling anynymous fuction")
	// 	return Person{
	// 		FirstName: firstname,
	// 		LastName:  lastName,
	// 		Age:       25,
	// 	}
	// }
	// person2 := person1("arash", "rahimi")
	// fmt.Println(person2.FirstName)
	/*--------------------------*/
	var wg sync.WaitGroup
	var lock sync.Mutex
	var i int
	var sum = 0
	for i = 0; i < 10000000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			lock.Lock()
			sum = sum + 1
			defer lock.Unlock()
		}(i)

	}
	wg.Wait()
	fmt.Println(sum)
	/*--------------------------*/
}

// func sumSequential()

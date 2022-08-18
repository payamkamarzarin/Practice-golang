package file

import (
	"fmt"
	"os"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func File() {
	persons := []Person{
		{
			FirstName: "arash",
			LastName:  "zarin",
			Age:       23,
		}, {
			FirstName: "payam",
			LastName:  "zarin",
			Age:       23,
		},
	}
	file, _ := os.Create("result.txt")
	defer file.Close()
	for _, Person := range persons {
		file.WriteString("first name: " + Person.FirstName + "\n")
		file.WriteString("last name: " + Person.LastName + "\n")
		file.WriteString(fmt.Sprintf("%s %d \n", "Age: ", Person.Age))
		file.WriteString("---------------------- \n")
	}
}

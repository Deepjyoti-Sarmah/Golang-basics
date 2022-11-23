package main

import "fmt"

// type Student struct {
// 	Name     string
// 	Rollno   int
// 	Subjects []string
// }

type Processor struct {
	processorName string
	cores         int
}

type Memory struct {
	memoyCapacity int
	memoryType    string
}

type Computer struct {
	Processor
	Memory
	price int
}

func main() {
	// student1 := Student{
	// 	Name:   "Deepjyoit",
	// 	Rollno: 68,
	// 	Subjects: []string{
	// 		"Maths",
	// 		"Physics",
	// 		"Chemistry",
	// 	},
	// }
	// // student1.Rollno = 16
	// student2 := student1
	// student2.Name = "Kunal"
	// fmt.Println(student1)
	// fmt.Println(student2)

	computer := Computer{}
	computer.price = 50000
	computer.processorName = "AMD 5500U"
	computer.cores = 6
	computer.memoyCapacity = 8
	computer.memoryType = "DDR4"

	computer1 := Computer {
		Processor: Processor{
			processorName: "Intel i5",
			cores: 6,
		},
		Memory: Memory{
			memoyCapacity: 8,
			memoryType: "DDR4",
		},
		price: 45000,
	}

	fmt.Println(computer.Processor)
	fmt.Println(computer1)
}

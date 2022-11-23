package main

import (
	"os"
)

func main() {
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)

	// f := createFile("defer.txt")
	// defer closeFile(f);
	// writeFile(f);

	// fmt.Println("Start")
	// panic("This is panic")
	// fmt.Println("End")

	panic("a problem")

	_, err := os.Create("newFile.txt")
	if err != nil {
		panic(err)
	}
}

// func createFile(p string) *os.File  {
// 	fmt.Println("creating")
// 	f,err := os.Create(p)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return f;
// }

// func writeFile(f *os.File)  {
// 	fmt.Println("writing")
// 	fmt.Fprintln(f, "data");
// }

// func closeFile(f *os.File)  {
// 	fmt.Println("closing")
// 	err := f.Close()

// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "error: %v\n", err)
// 		os.Exit(1)
// 	}

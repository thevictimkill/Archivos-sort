package main

import (
	"fmt"
	"sort"
	"os"
)

func main() {
	var n int
	var aux string
	var slice []string

	fmt.Println("Numero de strings: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Println("String: ")
		fmt.Scan(&aux)

		slice = append(slice,aux)
	}

	fmt.Println(slice)

	sort.Strings(slice)

	//fmt.Println(slice)

	file, err := os.Create("ascendente.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for i := 0; i < n; i++ {
		file.WriteString(slice[i])
		file.WriteString("\n")	
	}
	

	sort.Sort(sort.Reverse(sort.StringSlice(slice)))
	file1, err1 := os.Create("descendente.txt")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	defer file.Close()

	for i := 0; i < n; i++ {
		file1.WriteString(slice[i])
		file1.WriteString("\n")	
	}
	//fmt.Println(slice)


}
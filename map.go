

package main

import "fmt"

func main() {
	var a = map[string]string{"Name": "Sarojini", "Role": "Developer"}
	b := map[string]int{"ID": 101, "Sal": 20000}
	fmt.Println(a)
	fmt.Println(b)
	a["location"] = "Chennai" //adding the elements
	b["year"] = 2023
	fmt.Println("After adding the elements:")
	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
}

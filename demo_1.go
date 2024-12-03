package main

import (
	"fmt"
)

func main() {
	const name string = "Tom"
	fmt.Println(name)

	const age = 30
	fmt.Println(age)

	const name_1, name_2 string = "Tom", "Jay"
	fmt.Println(name_1, name_2)

	const name_3, age_1 = "Tom", 30
	fmt.Println(name_3, age_1)

	fmt.Print("输出到控制台不换行")
	fmt.Println("---")
	fmt.Println("输出到控制台并换行")
	fmt.Printf("name=%s,age=%d\n", "Tom", 30)
	fmt.Printf("name=%s,age=%d,height=%v\n", "Tom", 30, fmt.Sprintf("%.2f", 180.567))
}

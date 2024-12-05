package pkg1

import "fmt"

// lib1提供的API
func Lib1Test() {
	fmt.Println("lib1Test()...")
}

func init() {
	fmt.Println("lib1")
}

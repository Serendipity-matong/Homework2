package main

import "fmt"

func main() {
	//第一个复数
	var num1, num2 float64
	fmt.Print("Enter first real part: ")
	fmt.Scan(&num1)
	fmt.Print("Enter first img part: ")
	fmt.Scan(&num2)
	var complex1 = complex(num1, num2)
	//第二个复数
	var num3, num4 float64
	fmt.Print("Enter first real part: ")
	fmt.Scan(&num3)
	fmt.Print("Enter first img part: ")
	fmt.Scan(&num4)
	var complex2 = complex(num3, num4)
	//计算
	fmt.Printf("The answer after adding two complex numbers: %f\n", complex1+complex2)
	fmt.Printf("The answer after subtracting two complex numbers: %f\n", complex1-complex2)
	fmt.Printf("The answer after multiplying two complex numbers:%f\n", complex1*complex2)
	fmt.Printf("The answer after quotienting two complex numbers:%f\n", complex1/complex2)
	fmt.Printf("Eliminate too many zeros:\n")
	//去0
	fmt.Printf("%g\n", complex1+complex2)
	fmt.Printf("%g\n", complex1-complex2)
	fmt.Printf("%g\n", complex1*complex2)
	fmt.Printf("%g\n", complex1/complex2)
}

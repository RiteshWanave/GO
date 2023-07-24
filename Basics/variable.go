package main
import ("fmt")

func main () {
	var v1 string = "Anjali & "						 // Can be declared inside as well as outside
	fmt.Printf(v1)

	v2 := "Ram" 					// Can't Be Null // Can only be declared inside the function
	fmt.Println(v2)

	var a string
	var b int
	var c float64
	var d bool
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	var v3 = 30
	fmt.Println(v3)

	var a1, a2, a3, a4 int = 5, 6, 7, 8
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	var (
		b1 int
		b2 int = 77
		b3 string = "Neha"
	)
	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)

}
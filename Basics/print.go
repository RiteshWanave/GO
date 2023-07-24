package main
import ("fmt")

func main () {
	var i = "Anjali"
	var j = "Ram"
	fmt.Print(i, j)

	fmt.Print("\n")

	var a = 2
	var b = 10
	fmt.Print(a, b)

	fmt.Print("\n")

	fmt.Println(i, j)

	var k = "Ritesh"
	fmt.Printf("Value = %v and Type = %T \n", a, a)
	fmt.Printf("Value = %#v and Type = %T \n", k, k)
}
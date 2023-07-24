package main
import ("fmt")

func main () {
	const C = "this is constant"		// Constants are generally written in uppercase format // can be declared inside as well as outside
	fmt.Println(C)

	const (
		C1 = "Ritesh"
		C3 = "Neha"
		C2 = "Ram"
		C4= "Anjali"
	)
	fmt.Printf(C1)
	fmt.Printf(C2)
	fmt.Printf(C3)
	fmt.Printf(C4)
}
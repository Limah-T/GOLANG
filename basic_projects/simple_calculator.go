package main  // This indicates that the file belongs to the main package
import ("fmt")  // This is also known as "format"

func calculator(num1 int, num2 int, op string){
	if op == "+" {
		var result = num1 + num2
		fmt.Printf("Result is %v", 
		result) 
	} else if op == "-"{
		var result = num1 - num2
		fmt.Printf("Result is %v", result)
	} else if op == "*"{
		var result = num1 * num2
		fmt.Printf("Result is %v", result)
	} else if op == "/"{
		if num2 == 0 {
			fmt.Printf("%v cannot be divided by %v", num2, num1)
		} else {
			var result = num1 / num2
			fmt.Printf("Result is %v", result)
		}
	} else {
		fmt.Print("Invalid operator input!")
	}
}

// func main() {
// 	var(num1 int 
// 		num2 int
// 		op string)
// 	fmt.Print("Enter first num: ")
// 	fmt.Scan(&num1)
// 	fmt.Print("Enter the operator to use - +, -, /, *: ")
// 	fmt.Scan(&op)
// 	fmt.Print("Enter second num: ")
// 	fmt.Scan(&num2)
	
// 	calculator(num1, num2, op)
// }

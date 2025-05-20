package main
import ("fmt")
func simple_banking_operations() {
	var action string
	balance := 0.00
	var amount float64
	fmt.Print("Enter 'D' for Deposit, 'W' for Withdraw: ")
	fmt.Scan(&action)
	fmt.Print("Enter amount (Only numbers are allowed): ")
	fmt.Scan(&amount)
	if action == "D" {
		if amount <  1 {
			fmt.Print("Amount must be greater than 0")
		} else {
			balance += amount
			fmt.Printf("Your balance is %.2f", balance)
		}
	} else if action == "W" {
		if balance < amount {
			fmt.Print("Insufficient balance!")
		} else if balance >= amount {
			balance -= amount
			fmt.Printf("Withdrew %v\nBalance - %v", amount, balance)
		}
	} else {
		fmt.Print("Invalid Input")
	}
}
func main() {
	simple_banking_operations()
}
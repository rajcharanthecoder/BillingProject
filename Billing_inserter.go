package main //This is main package

// Language: go
// import is used to import the packages
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// bill is a struct and it is a bill and it has name, items and tip as fields of the struct bill and it is a slice of pointers to item struct

//getInput is a function which takes prompt and reader as input and returns the input and error as output and it is a function which is used to get the input from the user //

func getInput(prompt string, r *bufio.Reader) (string, error) { //r is reader to take input from the user
	//bufio is a package which is used to read the input from the user
	//reader is a struct which is used to read the input from the user

	fmt.Print(prompt) //fmt is a package which is used to print the output to the user

	//prompt is a string which is used to print the output to the user

	input, err := r.ReadString('\n') //ReadString is a function which is used to read the input from the user

	//err is a variable which is used to store the error if any error occurs while reading the input from the user

	return strings.TrimSpace(input), err //strings is a package which is used to trim the spaces from the input
}

func createBill() bill { //createBill is a function which is used to create the bill
	reader := bufio.NewReader(os.Stdin)

	//reader is a struct which is used to read the input from the user
	//bufio is a package which is used to read the input from the user
	//NewReader is a function which is used to create a new reader to read the input from the user and it takes os.Stdin as input and returns the reader as output and os.Stdin is a variable which is used to read the input from the user and it is a file descriptor which is used to read the input from the user

	name, _ := getInput("Create a new bill name: ", reader) //name,_ is a string which is used to store the name of the bill and it is a variable which is used to store the error if any error occurs while reading the input from the user

	b := newBill(name) //newBill is a function which is used to create a new bill and it takes name as input and returns the bill as output

	fmt.Println("Created the bill -", b.name) //This line prints the output to the user

	return b //returning the bill
}

func promptOptions(b bill) { //promptOptions is a function which is used to prompt the options to the user and it takes bill as input and returns nothing as output

	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	//opt,_ is a string which is used to store the option and it is a variable which is used to store the error if any error occurs while reading the input from the user

	//opt == "a" { //if the option is a then it will add the item to the bill
	//opt == "s" { //if the option is s then it will save the bill
	//opt == "t" { //if the option is t then it will add the tip to the bill
	//opt == "x" { //if the option is x then it will exit the program where x is anyOther Letter

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil { //This means if error is present then only execute the code inside the if block aand if the input is correct why error will come as nil and if the input is incorrect then error will come and it will execute the code inside the if block
			fmt.Println("The price must be a number...")
			promptOptions(b)
		}
		b.addItem(name, p)

		fmt.Println("item added -", name, price)
		promptOptions(b)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number...")
			promptOptions(b)
		}
		b.updateTip(t)

		fmt.Println("tip has been updated to", tip)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("bill has been saved as", b.name)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}
}

func main() {

	mybill := createBill() //createBill is a function which is used to create the bill
	promptOptions(mybill)  //promptOptions is a function which is used to prompt the options to the user and it takes bill as input and returns nothing as output

}

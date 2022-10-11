package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string             //name is a string
	items map[string]float64 //map is used to store key value pairs
	tip   float64            //tip is a float64
}

// make new bills

// this newBill Function return type is struct bill and it takes a string as a parameter
// and it returns a struct bill with name, items and tip
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// add item to bill

// This adItem function takes two parameters name and price and it returns nothing and it adds the item to the bill and this is call by reference so we need to use pointer to the bill and we can use the pointer to the bill to access the items map and we can add the item to the map by using the name as the key and the price as the value and we can use the pointer to the bill to access the tip and we can update the tip by adding the price to the tip and we can use the pointer to the bill to access the name and we can print the name of the bill and the name of the item and the price of the item
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price //This is Key Value Pair Update
}

// format the bill to string format to save it to the file system
func (b *bill) format() string {
	fs := "Bill breakdown:\n" //fs is a string and it is a format string
	var total float64 = 0     //total is a float64 and it is a total price of the bill

	// list items
	for k, v := range b.items { //k is a string and it is a key and v is a float64 and it is a value and we are looping through the items map
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v) //fs is a string and it is a format string and we are adding the key and the value to the format string
		//Sprintf is used to format the string and it takes a format string and the values to be formatted and it returns a string
		//Sprintf syntax is %v and it is a value and it is a default format and it is used to format the value and it is used to format the value to the default format

		total += v //total updates by adding the value to the total
	}

	// add tip
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip) //fs is a string and it is a format string and we are adding the tip to the format string to add Tip to the bill

	// add total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip) //fs is a string and it is a format string and we are adding the total to the format string to add Total to the bill

	return fs //fs is a string and it is a format string and we are returning the format string to the format function and we are returning the format string to the save function
}

// update tip
func (b *bill) updateTip(tip float64) { //tip is a float64 and it is a tip amount and we are updating the tip amount
	(*b).tip = tip
	// b.tip = tip
}

// save bill
func (b *bill) save() { //we are saving the bill
	data := []byte(b.format())                                        //data is a byte slice and it is a data to be saved and we are converting the format string to the byte slice
	err := os.WriteFile("billing_Details/"+b.name+".txt", data, 0644) //err is an error and it is an error and we are saving the data to the file system to save each person bill in a separate file

	//os package is used to interact with the operating system

	//WriteFile is used to write the data to the file system and it takes the file name, the data to be written and the permission and it returns an error

	//0644 is a permission and it is a permission to read and write the file

	//Syntax of os.writeFile is os.WriteFile("directory/file name" + "file_extension", data, permission)
	// handle error

	if err != nil {
		panic(err)
	}
	fmt.Println("Bill saved to file")
}

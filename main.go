package main

import (
	"fmt"
	"gostock/sina"
	"os"
)

func main() {
	os.Args = os.Args[1:] //Development mode
	if len(os.Args) != 3 {
		fmt.Printf("usage: stock get <name>\n Now:%v", os.Args)

		os.Exit(1)
	}
	if os.Args[1] == "list" {
		data, _ := sina.ListStock(os.Args[2])
		for _, v := range data {
			fmt.Printf("Name: %v Code: %v\n", v.Name, v.Code)
		}

	} else {
		data := sina.GetData(os.Args[2])
		fmt.Println("Name:  ", data.Name)
		fmt.Println("High:  ", data.High)
		fmt.Println("Low:   ", data.Low)
		fmt.Println("Open:  ", data.Open)
		fmt.Println("Close: ", data.Close)
		fmt.Println("Volume:", data.Volume)
	}

}

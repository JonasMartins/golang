package cmd

import (
	"errors"
	"fmt"

	//"os"
	"reflect"
)

func menu() error {
	//cmd := exec.Command("clear")
	//cmd.Stdout = os.Stdout
	//cmd.Run()
	fmt.Print("\n\noptions:\n")
	fmt.Print("1.	Generate random data\n")
	fmt.Print("2.	Labeling existent data\n")
	fmt.Print("3.	Exit\n")

	var input uint8
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Invalid option")
		return err
	}

	fmt.Println(reflect.TypeOf(input))

	switch input {
	case 1:
		fmt.Println("\n\tGenerating new data ...")
		return GenerateNewData()
	case 2:
		fmt.Println("\n\tLabeling existing data ...")
		return StartAnalysis()
	case 3:
		fmt.Println("\n\tExit")
	default:
		return errors.New("invalid input")
	}
	return nil
}

func Run() {
	fmt.Println("\n\tWelcome to Six Random Analyzer")
	for menu() != nil {
		menu()
	}
}

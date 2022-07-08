package cmd

import (
	"errors"
	"fmt"

	"os"
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

	switch input {
	case 1:
		fmt.Println("\n\tGenerating new data ...")
		return GenerateNewData()
	case 2:
		fmt.Println("\n\tLabeling existing data ...")
		return StartAnalysis()
	case 3:
		fmt.Println("\n\tExiting ...")
		os.Exit(0)
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

package main

import (
	"fmt"
	ui "github.com/Poleron402/Polinka/UI"
	"github.com/Poleron402/Polinka/database"
	"github.com/Poleron402/Polinka/modules"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	ui.AppUI()
	polinkadatabase.CreateDBIfNotExist()
	
	fmt.Println("\n\nYou are ready to learn! 🤓")
	

	selector.SelectAction()

	fmt.Println("\nProgram stopping.")
}
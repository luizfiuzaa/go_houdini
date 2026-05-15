package main

import (
	"fmt"
	"os"

	"go_houdini/src/functions"
	"go_houdini/src/structs"
)

func printHelp() {
	fmt.Println(`Usage: houdini <command>

Commands:
  create    Interactively scaffold a new .go file

Flags:
  -h, --help    Show this help message

Examples:
  houdini create
  houdini --help`)
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		fileData := functions.GetFileData()
		code := functions.GenerateCode(fileData.FileContent)
		file := structs.File{Name: fileData.FileName, Code: code}
		functions.CreateFile(&file)
	case "-h", "--help":
		printHelp()
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", os.Args[1])
		printHelp()
		os.Exit(1)
	}
}

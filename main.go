package main

import (
	"fmt"
	"os"

	"go_houdini/src/functions"
	"go_houdini/src/structs"
)

func printHelp() {
	fmt.Println(`Usage: houdini <command> [args]

Commands:
  create [filename]    Interactively scaffold a new .go file
                       Optionally pass the filename to skip the name prompt

Flags:
  -h, --help    Show this help message

Examples:
  houdini create
  houdini create hello_world
  houdini --help`)
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		var filename string
		if len(os.Args) >= 3 {
			filename = os.Args[2]
		}
		fileData := functions.GetFileData(filename)
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

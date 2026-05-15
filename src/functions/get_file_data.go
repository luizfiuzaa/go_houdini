package functions

import (
	"fmt"
	"log"

	"go_houdini/src/structs"
	"go_houdini/src/validator"
	"go_houdini/src/functions/utils"
)

func GetFileData(name string) structs.FileData {
	var file_package string
	var file_imports string
	var file_functions string

	fmt.Println()

	if name != "" {
		if validator.ValidateFileName(name) {
			log.Fatalf("Cannot create file '%s': file already exists or name is invalid.\n", name)
		}
	} else {
		for validator.ValidateFileName(name) {
			fmt.Print("📦 File name: ")
			name = utils.GetInfo()
		}
	}

	fmt.Print("📦 Package name (default 'main'): ")
	file_package = utils.GetInfo()

	fmt.Print("📚 Imports (space separated): ")
	file_imports = utils.GetInfo()

	fmt.Print("🔧 Functions (space separated): ")
	file_functions = utils.GetInfo()
	fmt.Println()

	file_code := structs.FileCode{
		FilePackage:   file_package,
		FileImports:   file_imports,
		FileFunctions: file_functions,
	}

	return structs.FileData{
		FileName:    name,
		FileContent: file_code,
	}
}



/*
  File created with go_houdini

  Check on: https://github.com/luizfiuzaa/go_houdini
*/

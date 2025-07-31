package functions

import (
  "fmt"
  "go_houdini/src/structs"
  "go_houdini/src/functions/utils"
)

func GetFileData() structs.FileData{
  var name string
	var file_package string
	var file_imports string
	var file_functions string

	fmt.Println("[HOUDINI WAS INVOKED!]")
	for utils.ValidateFileName(name) {
		fmt.Println("Enter a file name:")
		name = utils.GetInfo() 
		fmt.Println()
	}
	
	fmt.Println("[PACKAGE NAME]")
	fmt.Println("Want to define a package name? (by default it will be 'main'):")
	file_package = utils.GetInfo()
	fmt.Println()

	fmt.Println("[SPACE SEPARATED]")
	fmt.Println("What to import? (Enter to jump this part):")
	file_imports = utils.GetInfo() 
	fmt.Println()
	
	fmt.Println("[SPACE SEPARATED]")
	fmt.Println("Enter the functions name (Enter to jump this part):")
	file_functions = utils.GetInfo() 
	fmt.Println()

	fmt.Printf("Houdini will do some magic to create your '%s' file\n", utils.ValidateExtension(name))

	file_code := structs.FileCode{
		FilePackage: file_package,
		FileImports: file_imports,
		FileFunctions: file_functions,
	}

	return structs.FileData{
		FileName: name,
		FileContent: file_code,
	}
}



/*
  File created with go_houdini

  Check on: https://github.com/luizfiuzaa/go_houdini
*/
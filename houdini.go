package main

import "go_houdini/src/structs"

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)


func main(){
	var file_data []string = getData()
	writed_code := generateCode(file_data)

	var file = structs.File{
		FileName: file_data[0],
		Code: writed_code,
	}

	createFile(&file)
}

func getData() []string{
	var name string
	var imports string
	var functions string

	fmt.Printf("[HOUDINI WAS INVOKED!]\n")
	for validateFileName(name) {
		fmt.Println("Enter a file name:")
		name = getInfo() 
		fmt.Println()
	}

	fmt.Println("[SPACE SEPARATED]")
	fmt.Println("What you want to import? (Enter to jump this part)")
	imports = getInfo() 
	fmt.Println()
	
	fmt.Println("[SPACE SEPARATED]")
	fmt.Println("Enter the functions name (Enter to jump this part)")
	functions = getInfo() 
	fmt.Println()

	fmt.Printf("Houdini will do some magic to create your %s file\n", validateExtension(name))

	return []string{name, imports, functions}
}

func getInfo() string{
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input_value := strings.TrimSpace(input)

	return input_value
}

func validateFileName(file_name string) bool{
	already_exists := false
	name_is_empty := file_name == ""

	filepath.Walk("./", (func (path string, info os.FileInfo, err error) error{
		if(strings.Contains(info.Name(), validateExtension(file_name)) && !name_is_empty){
			already_exists = true
			fmt.Println("! [ERROR] - Already exists a file with this name")
			return nil
		}
		return nil
	}))

	return already_exists || name_is_empty
}

func validateExtension(file_name string) string{
	if(!strings.Contains(file_name, ".go")){
		file_name += ".go"
	}
	return file_name
}

func writeImports(imports string) string{
	if(imports != ""){
		imports_list := strings.Fields(imports)
		writed_code := "import (\n"
		for _, pkg := range imports_list {
			writed_code += " \"" + pkg + "\"\n"
		}
		writed_code += ")\n\n"
		return writed_code
	}

	return ""
}

func writeFunctions(writed_code string, functions string) string{
	writed_code += "func main(){\n\n}\n\n"
	
	if(functions != ""){	
		functions_list := strings.Fields(functions)
		for _, function := range functions_list {
			writed_code += "func " + function + "() {}\n\n"
		}
	}

	return writed_code 
}

func generateCode(file_data []string) string{
	var code string = "package main\n\n"

	code += writeImports(file_data[1])
	code = writeFunctions(code, file_data[2])
	code += "\n\n/*\n File created with go_houdini\n\n Check on: https://github.com/luizfiuzaa/go_houdini \n*/"

	return code
}

func createFile(file *structs.File) {
	file.FileName = validateExtension(file.FileName)
	err := os.WriteFile(file.FileName, []byte(file.Code), 0644)

	if err != nil {
		fmt.Println("ERROR CREATING FILE:\n", err)
		log.Fatal(err)
	} else{	
		fmt.Printf("\n\nHOUDINI HAS CREATED YOUR FILE!\n\nENJOY!\n")
		fmt.Println()
	}
}
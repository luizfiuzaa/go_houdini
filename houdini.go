package main

import (
	"go_houdini/src/structs"
	"go_houdini/src/functions"
)

func main(){
	file_data := functions.GetFileData()
	writed_code := functions.GenerateCode(file_data.FileContent)

	file := structs.File{
		Name: file_data.FileName,
		Code: writed_code,
	}

	functions.CreateFile(&file)
}
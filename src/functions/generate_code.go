package functions

import (
  "go_houdini/src/writer"
  "go_houdini/src/structs"
)

func GenerateCode(file_data structs.FileCode) string{
	var code string = writer.WritePackage(file_data.FilePackage)
	code += writer.WriteImports(file_data.FileImports)
	code = writer.WriteFunctions(code, file_data.FileFunctions)
	code += "\n\n/*\n  File created with go_houdini\n\n  Check on: https://github.com/luizfiuzaa/go_houdini\n*/"

	return code
}



/*
  File created with go_houdini

  Check on: https://github.com/luizfiuzaa/go_houdini
*/
package functions

import (
  "go_houdini/src/structs"
  "go_houdini/src/functions/utils"
)

func GenerateCode(file_data structs.FileCode) string{
	var code string = utils.WritePackage(file_data.FilePackage)
	code += utils.WriteImports(file_data.FileImports)
	code = utils.WriteFunctions(code, file_data.FileFunctions)
	code += "\n\n/*\n  File created with go_houdini\n\n  Check on: https://github.com/luizfiuzaa/go_houdini\n*/"

	return code
}



/*
  File created with go_houdini

  Check on: https://github.com/luizfiuzaa/go_houdini
*/
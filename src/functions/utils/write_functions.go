package utils

import (
  "strings"
)

func WriteFunctions(writed_code string, functions string) string{
	if(functions != ""){	
		functions_list := strings.Fields(functions)
		for _, function := range functions_list {
			writed_code += "func " + function + "() {}\n\n"
		}
	}

	return writed_code
}


/*
  File created with go_houdini

  Check on: https://github.com/luizfiuzaa/go_houdini
*/
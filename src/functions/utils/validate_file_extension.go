package utils

import (
  "strings"
)

func ValidateExtension(file_name string) string{
  if(!strings.Contains(file_name, ".go")){
		file_name += ".go"
	}

	return file_name
}



/*
  File created with go_houdini

  Check on: https://github.com/luizfiuzaa/go_houdini
*/
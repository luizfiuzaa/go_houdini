package validator

import (
  "fmt"
  "os"
  "strings"
  "path/filepath"
)

func ValidateFileName(file_name string) bool{
  already_exists := false
	name_is_empty := file_name == ""

	filepath.Walk("./", (func (path string, info os.FileInfo, err error) error{
		if(strings.Contains(info.Name(), ValidateExtension(file_name)) && !name_is_empty){
			already_exists = true
			fmt.Println("! [ERROR] - Already exists a file with this name")
			return nil
		}
		return nil
	}))

	return already_exists || name_is_empty
}



/*
  File created with go_houdini

  Check on: https://github.com/luizfiuzaa/go_houdini
*/
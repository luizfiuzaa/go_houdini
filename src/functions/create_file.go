package functions

import (
  "os"
  "fmt"
  "log"
  "go_houdini/src/structs"
  "go_houdini/src/validator"
  "go_houdini/src/functions/utils"
)

func CreateFile(file *structs.File) {
	file.Name = validator.ValidateExtension(file.Name)
	err := os.WriteFile(file.Name, []byte(file.Code), 0644)

	if err != nil {
		fmt.Println("ERROR CREATING FILE:\n", err)
		log.Fatal(err)
	} else{	
		utils.EndMessage()
	}
}



/*
  File created with go_houdini

  Check on: https://github.com/luizfiuzaa/go_houdini
*/
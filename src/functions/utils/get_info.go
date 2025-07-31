package utils

import (
  "os"
  "bufio"
  "strings"
)

func GetInfo() string{
  reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input_value := strings.TrimSpace(input)

	return input_value
}



/*
  File created with go_houdini

  Check on: https://github.com/luizfiuzaa/go_houdini
*/
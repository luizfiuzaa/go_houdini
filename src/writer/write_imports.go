package writer

import (
  "strings"
)

func WriteImports(imports string) string{
	if(imports != ""){
		imports_list := strings.Fields(imports)
		writed_code := "import(\n"
		for _, pkg := range imports_list {
			writed_code += "  \"" + pkg + "\"\n"
		}
		writed_code += ")\n\n"
		return writed_code
	}

	return ""
}


/*
  File created with go_houdini

  Check on: https://github.com/luizfiuzaa/go_houdini
*/
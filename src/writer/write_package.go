package writer

func WritePackage(package_name string) string{
	var package_define = "package "
	
	if package_name == "" {
		package_name = "main"
	}
	package_define += package_name + "\n\n"

	return package_define
}



/*
  File created with go_houdini

  Check on: https://github.com/luizfiuzaa/go_houdini
*/
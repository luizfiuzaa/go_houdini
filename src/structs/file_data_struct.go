package structs

type FileData struct{
	FileName string
	FileContent FileCode
}

type FileCode struct{
	FilePackage string
	FileImports string
	FileFunctions string
}
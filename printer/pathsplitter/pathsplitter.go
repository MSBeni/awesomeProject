package pathsplitter

import ( "fmt"; "path" )
 
// A function to print the directory and file name from file [path]
func Splitpath(pathstring string){
	var dir, file string
	dir, file = path.Split(pathstring)
	fmt.Println("the directory is: ", dir, "The file name is: ", file)
}
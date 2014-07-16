package main

import "fmt"
import "flag"
import "os"

var homeDir string = os.Getenv("HOME")
var listName string = "/.bcon/bcon_files" //eventually, let people config this
var listFile *os.File

func main() {

	flag.Parse()
	printArgs()

	//process verbs (add, remove, etc)
	switch flag.Arg(0) {
	case "add":
		//add command: takes a path to a file, a name and a tag (optional).
		fmt.Println("adding", flag.Arg(1))
		addEntry()
	case "help":
		fmt.Println("This is the help for bcon! TODO: help people.")
	default:
		fmt.Println("Not a valid command, try bcon help.")

	}
}

func addEntry() {

	var err error

	//open file list for writing (in append mode), if fail, create new file
	listFile, err = os.OpenFile(homeDir+listName, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		listFile, err = os.Create(homeDir + listName)
		checkError(err)
	}
	defer listFile.Close()

	if flag.Arg(1) != "" {
		listFile.WriteString(flag.Arg(1) + " ")
	} else {
		listFile.WriteString("stuff\n")
	}

}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}
func printArgs() {
	for i := 0; i < flag.NArg(); i++ {
		fmt.Println(i, ":", flag.Arg(i))
	}
}

package main

import "fmt"
import "flag"
import "os"

var homeDir string = os.Getenv("HOME")
var listName string = "/.bcon/bcon_files" //eventually, let people config this
var listFile *os.File

func main() {

	flag.Parse()

	//process verbs (add, remove, etc)
	switch flag.Arg(0) {
	case "add":
		//add command: takes a path to a file, a name and a tag (optional).
		err := addEntry()
		if err != nil {
			fmt.Println(err)
		}
	case "help":
		printHelp()
	default:
		fmt.Println("Not a valid command, try bcon help.")

	}
}

func addEntry() error {

	//check args. needs to be filename, list name, then optionally, a list of tags
	// var fileName = flag.Args(1)
	// var entryName = flag.Args(2)
	// var tags = flag.Args(3)
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
		return BconError{"No argument to add!"}
	}

	//all good, lets boogie.
	return nil
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

func printHelp() {
	fmt.Println("bcon commands:\n")
	fmt.Println("   add (filename, name, [tags])  Adds a file.")
	fmt.Println("   search (name or tag)          Search the filelist by name or tag.")
	fmt.Println("   remove (name or tag)          Remove a file from the file list.")
	fmt.Println("   help                          Show this text. ")
}

type BconError struct {
	what string
}

func (e BconError) Error() string {
	return e.what
}

package main

import "fmt"
import "flag"
import "os"

var homeDir string = os.Getenv("HOME")
var listName string = "/.bcon/bcon_files" //eventually, let people config this

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
	fileName := flag.Arg(1)
	//ensure file exists
	if _, err := os.Stat(fileName); err != nil {
    	return BconError{"Could not find file."}
	}

	entryName := flag.Arg(2)
	//ensure name TODO: here, also check if name is unique
	if entryName == "" {
		return BconError{"Specify a name for the new entry."}
	}
	
	//process tags. TODO: maximum number of tags is 10. Look into expanding this? 
	tags := make([]string, 10)
	for x := 0; flag.Arg(x + 3) != "" && x < len(tags); x++ {
		tags[x] = flag.Arg(x + 3)
	}

	//open file list for writing (in append mode), if fail, create new file
	listFile, err := os.OpenFile(homeDir+listName, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		listFile, err = os.Create(homeDir + listName)
		checkError(err)
	}
	defer listFile.Close()

	//write to file!
	listFile.WriteString(entryName + " " + fileName)
	for _, v := range(tags){
		if v != "" {
			listFile.WriteString(" " + v)
		}
	}
	listFile.WriteString("\n")

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
	fmt.Println("   rvhemove (name or tag)          Remove a file from the file list.")
	fmt.Println("   help                          Show this text. ")
}

type BconError struct {
	what string
}

func (e BconError) Error() string {
	return e.what
}

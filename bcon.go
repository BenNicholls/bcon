package main

import "fmt"
import "flag"
import "os"
import "github.com/bennicholls/bcon/entries"

var homeDir string = os.Getenv("HOME")
var filelistPath string = "/.bcon/bcon_files" //eventually, let people config this
var entrylist entries.BconEntrylist

func main() {

	flag.Parse()
	var err error

	//grab the file list.
	entrylist, err = entries.ParseFilelist(homeDir + filelistPath)
	if err != nil {
		fmt.Println(err)
	}

	//process verbs (add, remove, etc)
	switch flag.Arg(0) {
	case "add":
		//add command: takes a path to a file, a name and a tag (optional).
		err := addEntry()
		if err != nil {
			fmt.Println(err)
		}
	case "list":
		entrylist.Print()
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
	if f, err := os.Stat(fileName); err != nil || f.IsDir() {
		return BconError{"Could not find file."}
	}

	entryName := flag.Arg(2)
	//ensure name exists
	if entryName == "" {
		return BconError{"Specify a name for the new entry."}
	}

	//process tags. TODO: maximum number of tags is 10. Look into expanding this?
	tags := make([]string, 10)
	for x := 0; flag.Arg(x+3) != "" && x < len(tags); x++ {
		tags[x] = flag.Arg(x + 3)
	}

	//add the entry. NOTE: this looks ugly.
	if dup := entrylist.Add(entryName, fileName, tags); !dup {
		return BconError{"Entry with name " + entryName + " already exists."}
	}

	err := entries.WriteFilelist(homeDir+filelistPath, entrylist)
	if err != nil {
		return BconError{err.Error()}
	}

	//all good, lets boogie.
	return nil
}

func printHelp() {
	fmt.Println("bcon commands:\n")
	fmt.Println("   add (filename, name, [tags])  Adds a file.")
	fmt.Println("   search (name or tag)          Search the filelist by name or tag.")
	fmt.Println("   remove (name)                 Remove a file from the file list.")
	fmt.Println("   list                          List all recorded files. ")
	fmt.Println("   help                          Show this text. ")
}

type BconError struct {
	what string
}

func (e BconError) Error() string {
	return e.what
}

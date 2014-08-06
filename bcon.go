package main

import "fmt"
import "flag"
import "os"
import "os/exec"
import "os/user"
import "path"
import "path/filepath"
import "strings"
import "github.com/bennicholls/bcon/entries"
import "github.com/bennicholls/bcon/util"

var filelistPath string = "/.bcon/bcon_files" //eventually, let people config this
var entrylist entries.BconEntrylist
var bconUser *user.User

func main() {

	flag.Parse()
	err := initialize()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//process verbs (add, remove, etc)
	switch flag.Arg(0) {
	case "add":
		//add command: takes a path to a file, a name and a tag (optional).
		err := addEntry()
		if err != nil {
			fmt.Println(err)
		}
	case "remove":
		if flag.Arg(1) == "" {
			fmt.Println("Please provide an entry name to remove")
			break
		}
		err := entrylist.Remove(flag.Arg(1))
		if err != nil {
			fmt.Println(err)
		}
	case "list":
		entrylist.Print()
	case "help":
		printHelp()
	case "":
		fmt.Println("Try \"bcon help\" maybe!")

	//No command, attempt to launch!
	default:

		entry, err := entrylist.Get(flag.Arg(0))
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(entry.Path())
		cmd := exec.Command("nano", entry.Path())
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

	//cleanup
	if entrylist.IsDirty() {
		err = entries.WriteFilelist(bconUser.HomeDir + filelistPath, entrylist)
		if err != nil {
			fmt.Println("Could not write to file: " + err.Error())
		}
	}
}

func initialize() error {

	var err error
	sudo := false
	//find current user. if command was called from sudo, ensure we can find our files
	if sudoUser := os.Getenv("SUDO_USER"); sudoUser == "" {
		bconUser, err = user.Current()
	} else {
		bconUser, _ = user.Lookup(sudoUser)
		sudo = true
	}

	//check if entry file exists
	if f, err := os.Stat(bconUser.HomeDir + filelistPath); err != nil || f.IsDir() {

		//If directory doesn't exist, create it. 
		err = os.MkdirAll(path.Dir(bconUser.HomeDir + filelistPath), 0755)
		if err != nil {
			return err
		}	

		_, err = os.Create(bconUser.HomeDir + filelistPath)
		if err != nil {
			return err
		}

		//if bcon was called from sudo, ensure folder/file has right owner
		if sudo {
			cmd := exec.Command("chown", "-R", bconUser.Username + ":" + bconUser.Username, ".bcon")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err = cmd.Run()
			if err != nil {
				return err
			}
		}	
	}

	//grab the file list.
	entrylist, err = entries.ParseFilelist(bconUser.HomeDir + filelistPath)
	if err != nil {
		return err
	}

	return nil
}

func addEntry() error {

	//check args. needs to be filename, list name, then optionally, a list of tags
	argIndex := 1 //Tracks how many arguments the pathname took, used to offset other arguments
	fileName := flag.Arg(argIndex)

	for ; strings.HasSuffix(flag.Arg(argIndex), "\\"); argIndex++ {
		fileName += " " + flag.Arg(argIndex+1)
	}

	//ensure file exists
	if f, err := os.Stat(fileName); err != nil || f.IsDir() {
		return util.BconError{"Could not find file."}
	} else {
		fileName, err = filepath.Abs(fileName)
	}

	entryName := flag.Arg(argIndex + 1)
	//ensure name exists
	if entryName == "" {
		return util.BconError{"Specify a name for the new entry."}
	}

	//process tags. TODO: maximum number of tags is 10. Look into expanding this?
	tags := make([]string, 0, 10)
	for x := 0; flag.Arg(x+argIndex+2) != "" && x < cap(tags); x++ {
		tags = tags[0 : x+1]
		tags[x] = flag.Arg(x + argIndex + 2)
	}

	//add the entry.
	err := entrylist.Add(entryName, fileName, tags)
	if err != nil {
		return err
	}

	return nil
}

func printHelp() {
	fmt.Println("bcon commands:\n")
	fmt.Println("   add (filename, name, [tags])  Adds a file.")
	fmt.Println("   remove (name)                 Remove a file from the file list.")
	fmt.Println("   list                          List all recorded files. ")
	fmt.Println("   help                          Show this text. ")
}

package entries

import "fmt"
import "github.com/bennicholls/bcon/util"

type BconEntry struct {
	name     string
	fileName string
	tags     []string
}

//returns a string representing the entry.
func (e BconEntry) Output() string {
	out := e.name + " " + e.fileName
	for _, tag := range e.tags {
		out += " " + tag
	}
	return out
}

//returns the filename TODO: sanity checking maybe? Or do this during parsing?
func (e BconEntry) Path() string {
	return e.fileName
}

//Struct for holding the entry list.
type BconEntrylist struct {
	entries []BconEntry
}

//prints the entrylist to the console
func (list BconEntrylist) Print() {
	for _, e := range list.entries {
		fmt.Println(e.Output() + "\n")
	}
}

//add an entry to the entrylist.
func (list *BconEntrylist) Add(name string, path string, tags []string) error {

	//check for duplicates NOTE: should list.entries be a map? this would be easier.
	for _, e := range list.entries {
		if e.name == name {
			return util.BconError{"Entry with name " + name + " already exists."}
		}
	}

	//check for disallowed names (bcon commands, reserved keywords for bcon config files, etc.)
	disallowed := [...]string{"add", "remove", "bcon", "list", "help"}
	for _, v := range disallowed {
		if name == v {
			return util.BconError{"Name " + " not allowed (probably is a bcon command.)"}
		}
	}

	i := len(list.entries)
	list.entries = list.entries[0 : i+1]
	list.entries[i].name = name
	list.entries[i].fileName = path
	list.entries[i].tags = tags

	return nil
}

//Removes an entry. When the entrylist is written, it skips over name = ""
func (list *BconEntrylist) Remove(name string) error {
	for i, e := range list.entries {
		if e.name == name {
			list.entries[i].name = ""
			return nil
		}
	}

	return util.BconError{"No item called \"" + name + "\""}
}

func (list BconEntrylist) Get(name string) (BconEntry, error) {
	for _, e := range list.entries {
		if e.name == name {
			return e, nil
		}
	}

	return BconEntry{}, util.BconError{"No entry with name \"" + name +"\""}
}

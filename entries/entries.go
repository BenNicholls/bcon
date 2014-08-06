package entries

import "fmt"
import "github.com/bennicholls/bcon/util"

//Struct for an entry. Fields are exported so they can be marshaled (see go-yaml)
type BconEntry struct {
	Name     string
	Filename string
	Tags     []string
}

//returns a string representing the entry.
func (e BconEntry) Output() string {
	out := e.Name + " " + e.Filename + " ["
	for i, tag := range e.Tags {
		out += tag
		if i != len(e.Tags)-1 {
			out += ", "
		}
	}
	out += "]"
	return out
}

//returns the filename TODO: sanity checking maybe? Or do this during parsing?
func (e BconEntry) Path() string {
	return e.Filename
}

//Struct for holding the entry list.
type BconEntrylist struct {
	Entries []BconEntry
	dirty   bool
}

//prints the entrylist to the console
func (list BconEntrylist) Print() {
	for _, e := range list.Entries {
		fmt.Println(e.Output())
	}
}

//add an entry to the entrylist.
func (list *BconEntrylist) Add(name string, path string, tags []string) error {

	//check for duplicates NOTE: should list.Entries be a map? this would be easier.
	for _, e := range list.Entries {
		if e.Name == name {
			return util.BconError{"Entry with name " + name + " already exists."}
		}
	}

	//check for disallowed names (bcon commands, reserved keywords for bcon config files, etc.)
	disallowed := [...]string{"add", "remove", "bcon", "list", "help"}
	for _, v := range disallowed {
		if name == v {
			return util.BconError{"Name " + name + " not allowed (probably is a bcon command.)"}
		}
	}

	i := len(list.Entries)
	list.Entries = list.Entries[0 : i+1]
	list.Entries[i].Name = name
	list.Entries[i].Filename = path
	list.Entries[i].Tags = tags
	list.dirty = true

	return nil
}

//Removes an entry.
func (list *BconEntrylist) Remove(name string) error {
	for i, e := range list.Entries {
		if e.Name == name {
			list.Entries = append(list.Entries[:i], list.Entries[i+1:]...)
			list.dirty = true
			return nil
		}
	}

	return util.BconError{"No item called \"" + name + "\""}
}

func (list BconEntrylist) IsDirty() bool {
	return list.dirty
}

func (list BconEntrylist) Get(name string) (BconEntry, error) {
	for _, e := range list.Entries {
		if e.Name == name {
			return e, nil
		}
	}

	return BconEntry{}, util.BconError{"No entry with name \"" + name + "\""}
}

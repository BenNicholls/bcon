package entries

import "fmt"

type BconEntry struct {
	name string
	fileName string
	tags []string
}

//returns a string representing the entry.
func (e BconEntry) Output() string {
	out := e.name + " " + e.fileName
	for _, tag := range e.tags {
		out += " " + tag
	}
	return out
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

//add an entry to the entrylist TODO: check for duplicate names here.
func (list *BconEntrylist) Add(name string, path string, tags []string) {

	i := len(list.entries)
	list.entries = list.entries[0: i + 1]
	list.entries[i].name = name
	list.entries[i].fileName = path
	list.entries[i].tags = tags
}
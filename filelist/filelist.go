package filelist

import "os"
import "bufio"
import "strings"

type BconEntry struct {
	name string
	fileName string
	tags []string
}

//returns a string representing the entry.
func (e BconEntry) Output() string {
	out := e.name + " " + e.fileName
	for _, tag := range e.tags {
		out += " (" + tag + ")"
	}
	return out
}

//Parse the file list. If there is no filelist, it makes a blank one.
//TODO: actually throw some errors
func ParseFilelist(path string) ([]BconEntry, error) {

	//entries NOTE: is 50 too much as a default capacity? too small? who can say
	entries := make([]BconEntry, 0, 50)

	listFile, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		//NOTE: this just assumes the error indicates that the file doesn't
		//exist. Really, this should be checked I guess.
		listFile, err = os.Create(path)
		if err != nil {
			return entries, err
		}
	}
	defer listFile.Close()

	//parsing happens here. TODO: validate tokens, parse spaces in paths
	scanner := bufio.NewScanner(listFile)
	for line := 0; scanner.Scan(); line++ {
		tokens := strings.Split(scanner.Text(), " ")
		entries = entries[0:line + 1]
		entries[line].name = tokens[0]
		entries[line].fileName = tokens[1]
		entries[line].tags = tokens[2:]
	}

	return entries, nil
}
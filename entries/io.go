package entries

import "os"
import "bufio"
import "strings"

//Parse the file list. If there is no filelist, it makes a blank one.
//TODO: actually throw some errors
func ParseFilelist(path string) (BconEntrylist, error) {

	//entries NOTE: is 50 too much as a default capacity? too small? who can say
	list := BconEntrylist{make([]BconEntry, 0, 50)}

	listFile, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		//NOTE: this just assumes the error indicates that the file doesn't
		//exist. Really, this should be checked I guess.
		listFile, err = os.Create(path)
		if err != nil {
			return list, err
		}
	}
	defer listFile.Close()

	//parsing happens here. TODO: validate tokens, parse spaces in paths
	scanner := bufio.NewScanner(listFile)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		list.Add(tokens[0], tokens[1], tokens[2:])
	}

	return list, nil
}

//Writes the entries to the filelist. path is a full pathname.
func WriteFilelist(path string, list BconEntrylist) error {

	fileList, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fileList.Close()

	writer := bufio.NewWriter(fileList)

	for _, e := range list.entries {
		if e.name != "" {
			_, err := writer.WriteString(e.Output() + "\n")
			if err != nil {
				return err
			}
		}
	}

	writer.Flush()

	return nil
}

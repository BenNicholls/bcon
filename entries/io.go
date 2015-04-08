package entries

import "io/ioutil"
import "gopkg.in/yaml.v1"

//Parse the file list. If there is no filelist, it makes a blank one.
func ParseFilelist(filePath string) (BconEntrylist, error) {

	//entries NOTE: is 50 too much as a default capacity? too small? who can say
	list := BconEntrylist{make([]BconEntry, 0, 50), false}

	listBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return list, err
	}
	err = yaml.Unmarshal(listBytes, &list)
	if err != nil {
		return list, err
	}

	return list, nil
}

//Writes the entries to the filelist. path is a full pathname.
func WriteFilelist(path string, list BconEntrylist) error {

	listBytes, err := yaml.Marshal(&list)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, listBytes, 0660)
	if err != nil {
		return err
	}

	return nil
}

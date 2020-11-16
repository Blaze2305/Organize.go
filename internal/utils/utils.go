package utils

import (
	"fmt"
	"os"
)

// MoveFile -  moves a file to the given location
func MoveFile(from string, to string) {
	err := os.Rename(from, to)
	if err != nil {
		fmt.Println(err)
	}
}

// MakeInitFolders - makes all the initial folders for the code to fit files into
func MakeInitFolders() error {
	FOLDERS := [9]string{"Videos", "Images", "Executables", "Zips", "Documents", "Torrents", "Excels", "Random", "Folders"}

	for _, value := range FOLDERS {
		there, err := exists(value)
		if err != nil {
			return err
		}

		if there {
			continue
		} else {
			err = os.Mkdir(value, 0777)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// exists returns whether the given file or directory exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var folders []string = []string{"Videos", "Images", "Executables", "Zips", "Documents", "Torrents", "Excels", "Random", "Folders", "Music"}

// moveFile -  moves a file to the given location
func moveFile(from string, to string) {
	err := os.Rename(from, to)
	if err != nil {
		fmt.Println(err)
	}
}

// MakeInitFolders - makes all the initial folders for the code to fit files into
func MakeInitFolders() error {
	for _, value := range folders {
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

// OrganizeFolder - organizes the folders in the directory
func OrganizeFolder() error {
	items, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}
	for _, item := range items {
		fileName := item.Name()
		if presentInFoldersList(fileName) {
			continue
		}
		if item.IsDir() {
			newPath := fmt.Sprintf("Folders/%s", fileName)
			moveFile(fileName, newPath)
		} else {
			category := getCategory(filepath.Ext(fileName))
			newPath := fmt.Sprintf("%s/%s", category, fileName)
			moveFile(fileName, newPath)
		}
	}
	return nil
}

// getCategory - returns the category under which to store the file
func getCategory(extension string) string {
	extensionMap := map[string][]string{
		"Videos":      {".mp4", ".mkv", ".avi"},
		"Images":      {".png", ".jpg", ".jpeg", ".bmp", ".gif", ".jfif"},
		"Executables": {".msi", ".exe"},
		"Zips":        {".tar", ".gzip", ".7z", ".zip", ".rar", ".tar.gz"},
		"Documents":   {".pdf", ".docx", ".ppt", ".pptx"},
		"Torrents":    {".torrent"},
		"Excels":      {".xlsx", ".csv"},
		"Music":       {".mp3", ".wav", ".flac"},
	}

	for k := range extensionMap {
		for _, val := range extensionMap[k] {
			if val == extension {
				return (k)
			}
		}
	}
	return ("Random")
}

// presentInFoldersList - check if the folder is the same as the one mentioned in the folders const array
func presentInFoldersList(name string) bool {
	for _, val := range folders {
		if name == val {
			return true
		}
	}
	return false
}

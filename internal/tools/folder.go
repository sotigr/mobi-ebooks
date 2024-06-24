package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const JSON_PATH = "/mnt/media/folders.json"

type Folder struct {
	Name string `json:"name"`
}

func CheckFolder(name string) bool {
	return !strings.Contains(name, "/")
}

func EnsureFolderJsonExists() error {

	if _, err := os.Stat(JSON_PATH); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(JSON_PATH)
		if err != nil {
			return err
		}

		file.WriteString("[]")
		file.Close()
	}
	return nil
}

func AddFeaturedFolder(name string) error {
	EnsureFolderJsonExists()

	if FolderExists(name) {
		return errors.New("folder already exists")
	}

	folders := GetFeaturedFolders()
	folders = append(folders, Folder{Name: name})
	WriteFeaturedFolders(folders)
	return nil
}
func FolderExists(name string) bool {
	folders := GetFeaturedFolders()
	for _, folder := range folders {
		if folder.Name == name {
			return true
		}
	}
	return false
}

func RemoveFeaturedFolder(name string) error {
	EnsureFolderJsonExists()
	folders := GetFeaturedFolders()
	for i, folder := range folders {
		if folder.Name == name {
			folders = append(folders[:i], folders[i+1:]...)
			break
		}
	}
	WriteFeaturedFolders(folders)
	return nil
}

func GetFeaturedFolders() []Folder {
	EnsureFolderJsonExists()

	var folders []Folder

	file, err := os.Open(JSON_PATH)
	if err != nil {
		fmt.Println("Error openning JSON")
		return folders
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&folders)
	if err != nil {
		fmt.Println("Error decoding JSON")
		return folders
	}
	return folders
}

func WriteFeaturedFolders(folders []Folder) error {
	file, err := os.Create(JSON_PATH)
	if err != nil {
		return err
	}

	jsonb, err := json.Marshal(folders)
	if err != nil {
		return err
	}

	file.Write(jsonb)
	file.Close()

	return nil
}

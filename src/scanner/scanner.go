package scanner

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

// Scan a folder and return a string list including all paths of pdf files.
func ScanFolder(folderPath string) ([]string, error) {

	// Check folder path
	if _, err := os.Stat(folderPath); err != nil {
		if os.IsNotExist(err) {
			log.WithField("path", folderPath).Error("The folder path doesn't exist!")
			return nil, err
		} else {
			return nil, err
		}
	}

	var books []string

	err:= filepath.Walk(folderPath, func (path string, info os.FileInfo, err error) error{
		if err != nil {
			log.WithField("current_path", path).Error("Failed to access the path!")
			return err
		}
		if info.IsDir() { return nil }
		if info.Mode().IsRegular() && filepath.Ext(info.Name()) == ".pdf" {

			// If the file extension is in the default extension list, add the file
			// path to the path list.
			books = append(books, path)
		}
		return nil
	})

	if err != nil {
		log.WithField("path", folderPath).Error("Failed to walk through the folder!")
		return nil, err
	}

	return books, nil
}
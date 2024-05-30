package main


import (
	"os"
	"fmt"
	"path/filepath"
)
var filas []string
var Exts []string

func ext() ([]string ,string) {
	fmt.Println("Write the path to folder!!( ●ω●)♡")
	filep := repeat()
	filas = slicefiles(filep)
	for i:= 0; i < len(filas); i++{
		flext := filepath.Ext(filas[i])
		Exts = append(Exts, flext)
	}
	return Exts, filep
}
func org(){
	ex,f := ext()
	checkAndCreateFolders(f, ex)
	fl := slicefiles(f)
	for i:= 0; i < len(fl); i++{
		os.Rename(f + "/" + fl[i], f + "/" + ex[i] + "/" + fl[i])
		
	}
	
}
func checkAndCreateFolders(parentFolder string, folders []string) error {
	// Check if parent folder exists
	if _, err := os.Stat(parentFolder); os.IsNotExist(err) {
	return fmt.Errorf("Parent folder does not exist: %s", parentFolder)
	}
	
	
	// Check if each folder exists, create it if not
	for _, folder := range folders {
		folderPath := parentFolder + "/" + folder
		if _, err := os.Stat(folderPath); os.IsNotExist(err) {
			err := os.Mkdir(folderPath, 0755)
			if err != nil {
				return fmt.Errorf("Failed to create folder: %s", folderPath)
			}
			fmt.Printf("Folder created: %s\n", folderPath)
		} else {
			fmt.Printf("Folder already exists: %s\n", folderPath)
		}
	}
	return nil
	}
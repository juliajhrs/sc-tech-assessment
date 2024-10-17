package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// Your code here...
	var srcFolder *Folder
	var destFolder *Folder
	for i := range f.folders {
		if f.folders[i].Name == name {
			srcFolder = &f.folders[i] // located src
			break
		}
	}
	for i := range f.folders {
		if f.folders[i].Name == dst {
			destFolder = &f.folders[i] // located dest
			break
		}
	}

	// error checking
	if srcFolder == nil {
		return []Folder{}, errors.New("source folder does not exist")
	}
	if destFolder == nil {
		return []Folder{}, errors.New("destination folder does not exist")
	}
	if srcFolder.OrgId != destFolder.OrgId {
		return []Folder{}, errors.New("cannot move a folder to a different organization")
	}
	if srcFolder.Name == destFolder.Name {
		return []Folder{}, errors.New("cannot move a folder to itself")
	}
	if strings.HasPrefix(destFolder.Paths, srcFolder.Paths) {
		return []Folder{}, errors.New("cannot move a folder to a child of itself")
	}

	oldPath := srcFolder.Paths
	newPath := destFolder.Paths + "." + srcFolder.Name // constructs new path more src path
	for i := range f.folders {
		if strings.HasPrefix(f.folders[i].Paths, oldPath) {
			f.folders[i].Paths = strings.Replace(f.folders[i].Paths, oldPath, newPath, 1) // updates all other path that starts with the old path
		}
	}

	return f.folders, nil
}

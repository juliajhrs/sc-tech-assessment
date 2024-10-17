package folder

import (
	"errors"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res

}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	// Your code here...
	folders := f.folders
	res := []Folder{}

	checkCorrectOrg := false
	checkWrongOrg := false

	for _, f := range folders {
		folderHasChild := hasSubstringWithDot(f.Paths, name)
		if folderHasChild {
			res = append(res, f)
		}
		// validity checks
		if f.OrgId == orgID && f.Name == name {
			checkCorrectOrg = true
		} else if f.OrgId != orgID && f.Name == name {
			checkWrongOrg = true
		}
	}

	if !checkCorrectOrg {
		if checkWrongOrg {
			return nil, errors.New("Folder does not exist in the specified organization")
		}
		return nil, errors.New("Folder does not exist")
	}
	return res, nil
}

// Helper Functions

/*
Function to find if a string contains a dot after certain substring
Arg: (string, string) main string, specified substring
Return: (bool) true if it contains, false otherwise
...
Source: https://www.tutorialspoint.com/golang-program-to-check-a-string-contains-a-specified-substring-or-not
*/
func hasSubstringWithDot(str, substr string) bool {
	substrDot := substr + "."
	for i := 0; i < len(str)-len(substrDot)+1; i++ {
		if str[i:i+len(substrDot)] == substrDot {
			return true
		}
	}
	return false
}

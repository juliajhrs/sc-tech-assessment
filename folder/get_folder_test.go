package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	org1 := uuid.NewV4()
	org2 := uuid.NewV4()
	org3 := uuid.NewV4()

	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
		{
			name: ""
			orgID: org1
			folders: []folder.Folder{

			},
			want: []folder.Folder{

			},
		},
		{
			name: ""
			orgID: org1
			folders: []folder.Folder{

			},
			want: []folder.Folder{

			},
		},
		{
			name: ""
			orgID: org1
			folders: []folder.Folder{

			},
			want: []folder.Folder{

			},
		},
		{
			name: ""
			orgID: org1
			folders: []folder.Folder{

			},
			want: []folder.Folder{

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// f := folder.NewDriver(tt.folders)
			// get := f.GetFoldersByOrgID(tt.orgID)

		})
	}
}

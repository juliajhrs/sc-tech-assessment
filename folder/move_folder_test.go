package folder_test

import (
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()

	org1Str := "38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"
	org2Str := "c1556e17-b7c0-45a3-a6ae-9546248fb17a"
	org3Str := "9b4cdb0a-cfea-4f9d-8a68-24f038fae385"

	org1 := uuid.FromStringOrNil(org1Str)
	org2 := uuid.FromStringOrNil(org2Str)
	org3 := uuid.FromStringOrNil(org3Str)

	tests := [...]struct {
		name    string
		fname   string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
		err     error
	}{
		// TODO: your tests here
		{
			name:  "No child folders",
			fname: "creative-scalphunter",
			orgID: org1,
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: org1,
					Paths: "creative-scalphunter",
				},
			},
			want: []folder.Folder{},
			err:  nil,
		},
		{
			name:  "Multiple child folders",
			fname: "creative-scalphunter",
			orgID: org1,
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: org1,
					Paths: "creative-scalphunter",
				},
				{
					Name:  "clear-arclight",
					OrgId: org1,
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1,
					Paths: "creative-scalphunter.clear-arclight.topical-micromax",
				},
				{
					Name:  "close-layla-miller",
					OrgId: org1,
					Paths: "creative-scalphunter.close-layla-miller",
				},
				{
					Name:  "dashing-mirage",
					OrgId: org2,
					Paths: "noble-vixen.fast-watchmen.full-weapon-x.honest-greymalkin.dashing-mirage",
				},
				{
					Name:  "steady-insect",
					OrgId: org3,
					Paths: "steady-insect",
				},
			},
			want: []folder.Folder{
				{
					Name:  "clear-arclight",
					OrgId: org1,
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1,
					Paths: "creative-scalphunter.clear-arclight.topical-micromax",
				},
				{
					Name:  "close-layla-miller",
					OrgId: org1,
					Paths: "creative-scalphunter.close-layla-miller",
				},
			},
			err: nil,
		},
		{
			name:  "Folder does not exist in the specified organization",
			fname: "dashing-mirage",
			orgID: org1,
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: org1,
					Paths: "creative-scalphunter",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1,
					Paths: "creative-scalphunter.clear-arclight.topical-micromax",
				},
				{
					Name:  "dashing-mirage",
					OrgId: org2,
					Paths: "noble-vixen.fast-watchmen.full-weapon-x.honest-greymalkin.dashing-mirage",
				},
			},
			want: []folder.Folder{},
			err:  errors.New("Folder does not exist in the specified organization"),
		},
		{
			name:  "Folder does not exist at all",
			fname: "steady-insect",
			orgID: org1,
			folders: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: org1,
					Paths: "creative-scalphunter",
				},
				{
					Name:  "clear-arclight",
					OrgId: org1,
					Paths: "creative-scalphunter.clear-arclight",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1,
					Paths: "creative-scalphunter.clear-arclight.topical-micromax",
				},
				{
					Name:  "close-layla-miller",
					OrgId: org1,
					Paths: "creative-scalphunter.close-layla-miller",
				},
				{
					Name:  "dashing-mirage",
					OrgId: org2,
					Paths: "noble-vixen.fast-watchmen.full-weapon-x.honest-greymalkin.dashing-mirage",
				},
			},
			want: []folder.Folder{},
			err:  errors.New("Folder does not exist"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, get)
		})
	}
}

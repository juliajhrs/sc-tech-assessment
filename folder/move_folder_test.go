package folder_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
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
		src     string
		dest    string
		folders []folder.Folder
		want    []folder.Folder
		err     error
	}{
		// TODO: your tests here
		{
			name: "Moving folder within the same parent folder",
			src:  "clear-arclight",
			dest: "close-layla-miller",
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
					Name:  "creative-scalphunter",
					OrgId: org1,
					Paths: "creative-scalphunter",
				},
				{
					Name:  "clear-arclight",
					OrgId: org1,
					Paths: "creative-scalphunter.close-layla-miller.clear-arclight",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1,
					Paths: "creative-scalphunter.close-layla-miller.clear-arclight.topical-micromax",
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
			err: nil,
		},
		{
			name: "Moving folder under a different parent folder",
			src:  "clear-arclight",
			dest: "steady-insect",
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
					OrgId: org1,
					Paths: "steady-insect",
				},
			},
			want: []folder.Folder{
				{
					Name:  "creative-scalphunter",
					OrgId: org1,
					Paths: "creative-scalphunter",
				},
				{
					Name:  "clear-arclight",
					OrgId: org1,
					Paths: "steady-insect.clear-arclight",
				},
				{
					Name:  "topical-micromax",
					OrgId: org1,
					Paths: "steady-insect.clear-arclight.topical-micromax",
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
					OrgId: org1,
					Paths: "steady-insect",
				},
			},
			err: nil,
		},
		{
			name: "Moving folder to a different organization",
			src:  "clear-arclight",
			dest: "dashing-mirage",
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
			want: []folder.Folder{},
			err:  errors.New("Cannot move a folder to a different organization"),
		},
		{
			name: "Moving folder to itself",
			src:  "clear-arclight",
			dest: "clear-arclight",
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
			want: []folder.Folder{},
			err:  errors.New("Cannot move a folder to itself"),
		},
		{
			name: "Moving folder to a child of itself",
			src:  "clear-arclight",
			dest: "topical-micromax",
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
			want: []folder.Folder{},
			err:  errors.New("Cannot move a folder to a child of itself"),
		},
		{
			name: "Source folder does not exist",
			src:  "Random-folder",
			dest: "clear-arclight",
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
			},
			want: []folder.Folder{},
			err:  errors.New("Source folder does not exist"),
		},
		{
			name: "Destination folder does not exist",
			src:  "creative-scalphunter",
			dest: "Random-folder",
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
			},
			want: []folder.Folder{},
			err:  errors.New("Destination folder does not exist"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get, err := f.MoveFolder(tt.src, tt.dest)
			if err != nil && tt.err != nil && err.Error() != tt.err.Error() { // both nil errors
				t.Errorf("MoveFolder() error = %v, want %v", err, tt.err)
				return
			} else if (err == nil && tt.err != nil) || (err != nil && tt.err == nil) { // one nil, one isnt
				t.Errorf("MoveFolder() error = %v, want %v", err, tt.err)
				return
			}

			if !reflect.DeepEqual(get, tt.want) {
				t.Errorf("MoveFolder() = %v, want %v", get, tt.want)
			}
		})
	}
}

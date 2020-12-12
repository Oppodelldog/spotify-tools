//go:generate statics -i=assets\files -o=assets/files.go -pkg=assets -group=Assets

package assets

import (
	"os"

	"github.com/go-playground/statics/static"
)

// newStaticAssets initializes a new *static.Files instance for use
func newStaticAssets(config *static.Config) (*static.Files, error) {

	return static.New(config, &static.DirFile{
		Path:    "\\assets\\files",
		Name:    "files",
		Size:    4096,
		Mode:    os.FileMode(2147484159),
		ModTime: 1607785363,
		IsDir:   true,
		Compressed: `
`,
		Files: []*static.DirFile{{
			Path:    "\\assets\\files\\admin.html",
			Name:    "admin.html",
			Size:    1150,
			Mode:    os.FileMode(438),
			ModTime: 1607785363,
			IsDir:   false,
			Compressed: `
H4sIAAAAAAAA/5xUwY7TMBC9V+o/WDlwWSnOckF0k0ClIlEkVoiUMzLOJLbWsSPPZJcQ5d+R025p1ZbC
3sbP7z2/0YycKmoMM8LWWQT2+7ciyuezVIEo85Q0GciHId6EYhxTvkXmM8YYS422D8yDySKk3gAqAIqY
8lBlkUAEQi4RuXW+EUb/glgiRv8pNqJ3Hb1ESX3rXqLbXl5Qth6ksxbkXqmIWlxwXjlLGNdIgrSMpWuO
1WfJztUGRKsx8MPjr99VotGmz+4duZtCWLz59GXxVCt6f5skd2+S5O5tkrwqNbZG9Bk+iTY6aSnMj08D
nM/SH67sp4HeHs1R3ebz2TB4YWtgcQGI2lkcx13kUj/u0k/Hzhycdj3l69WCDUO8XgU/o88w7kUDEycU
F1kb3YCfaFMVf3Sdx3FUf5DP2nYEOI7N3ywYAjFBB04F0JIuvrvqtuFW3ZVsbClJP8KB8RoLuGxctI50
1ccb9wB2Uj0jSykBccKvqr9C5QHVqcnhxVWXDz9b7QHXxxZ79F9TQLmkcyECfuKR8v2+pHy7ScMAtgzb
lfLnheTh38l/BwAA///sudzCfgQAAA==
`,
			Files: []*static.DirFile{},
		},
			{
				Path:    "\\assets\\files\\callback.html",
				Name:    "callback.html",
				Size:    576,
				Mode:    os.FileMode(438),
				ModTime: 1607785317,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/5yST2syMRDG74LfIeTwXoTN6uWlurvtpZceSsH2XGIcN6HZTMiMSCp+9xLtH6Ttod7C
75nfhPCksTx44XXoWwnh+Wkpu/GosaDXXcOOPXT7ffVYDodDo05kPBJCiMa78CIS+FYSZw9kAVgKm2DT
Sk0ETMoQqYBp0N69QmWI5B9lrzNu+RKTc8RLvFP4ixkTGAwBzKdpmSPNldpgYKp6Ys3OVAaHc/vHYcTe
g46Oyny5fHa90YPzub1HxslSB5rcPcx3veWbaV0v/tf14qqu/60dRa9zSzsd5bcnlf7UscDxqFnhOh8L
nZ71aKdHOCtwuV198VnhxAlDX7LblDCV4B2V1R8rVfk53VsAAAD//6jri2lAAgAA
`,
				Files: []*static.DirFile{},
			},
			{
				Path:    "\\assets\\files\\index.html",
				Name:    "index.html",
				Size:    3841,
				Mode:    os.FileMode(438),
				ModTime: 1607782354,
				IsDir:   false,
				Compressed: `
H4sIAAAAAAAA/5xX227bRhO+D5B3mH8vfki1RVIpkLaWqB7iAnHRuEVk97ZZkyNxm+WusDu0oxB892JJ
SjwnaQQD3sPMN6dvhuQ6oVSC5GofMlR/32/Z5vmzdYI83jx/BgCwJkESN3nu3eEHst6d2xbF2q/OayEp
1HswKENm6SjRJojEIDG4Cxm3Fsn6kbW+0iblUnxEL7KW/UdlyY86o6/RpONBf41edTmheTAYaaUwOmsm
RAd75fs7rch6e0ucRORFOu1qjwprvZfID8I6eWf8xY87ngp5DG816YstV/bitz+vnvYJ/bQMgtV3QbD6
IQj+Hwt7kPwY2id+YIOQXC39upjrBx0f3f9YPEIkubUhixQBzyjRRnzEeJHn3o39+byHojh7nixHKJAs
6+s8B7EDpQn6ANV9CWHJaLVvYO5VY9qh1fctjZ02KfCIhFYhy3PwHPS9kc4xSJESHYdsX8UJrd9aqENG
QMcDhsxmD6kgdgrZ2WTwyGWGDvMUkt4eNInd8eQ8d0aLgvltf3zn0DlklBYnAry3aBRP8RNhiThkJFI0
C7dtwjloOxVPHUEs9oIsKxESVkdJ+IEYOJvu0N+s7YGrk0amBLFNsvbd4Rdjp2PY6Th2OsDuJKs8iMVj
3/ZDRqTPUAata0OtIimi9/V+Nmebt26x9ivpjo0uZpvaQkWLSA1SmedguNojeK+0IqOlbZdwwjGhopZb
QkUzR8a/HIegKOZs47a3PHW7ETdruqi4w5ZP+B4RH/O9LpXLa5c7tVqZvS7p+zzfIt05xR6zJ+oDI03o
vZLIzb2Rrgdrw5Uvwi4skhsipQnvxm6RmhHymRxHDtY1OScyt26StFDgXSwsf5AYO4pWK/auKCagq4TX
IZf+1kFPi1ecdrSPMzwH5tabxeKq/Bttn6ac57KP5LDfCzBGgHp/Hi8VX5rjtY2MOFCrKf1/+COvTk9Z
lkhQlgNCyItVcxhnCCEE3ZMbRWgeuYQQVCbl6mR9l6my4OC4XnJoDnnjamnAS+EirAi2aq6yQ8wJr6un
0mxe3xQD4Lq326i25uadvtVPZ9UR0ClMi1T25EzElzD0OtZRlqIib4/0q0S3/OV4E89EPPdKaeiGMzSw
bwx0kA1SZtTnDUwCn9+LygR00yJ2MDslfBPCy6Bz2y5HeF4t4GWwGhNKLi565w18Aht48e0AvaXcItDp
12qpYiq6XvnaBvqBrzp0qJLNEnZ58mBcID0LpNOME/ZGCRJcvkUV95NcV5ANhhf8LwyBkcmQTSL3iDvo
FNeK3awlVxBcdo/S4ZHtHhWTDtQvgdcZzr6I8eVcm3tCKTSv7978DiHEmSlfee70ZO8+CRXrJ08rqXkM
IczmEG76TB1kecCn6TYvwL1S9eXLwZXn3nWGRTFG6mF66xSfi/laZ8YWxeVQKr2CRuyNUBnhULBvtZ3u
/lVnplqk027WKF3CMgiC+ar9pCg+O0GHs3mkZJ3iZ9hpdlcdl8tNCMMRUj59G1+bKOYjxRgMAamj0hPP
oONGt6jNMtLKErjo3PMGn+CaE86CtrS79CzS1n1bxbZ0+BtYLLvZOs1bJ0z6ZvvHloxQ+9ncs9mDJTNb
Li/h+xaB1371lCy/h07fQb779N38GwAA//8qy65DAQ8AAA==
`,
				Files: []*static.DirFile{},
			},
		},
	})
}

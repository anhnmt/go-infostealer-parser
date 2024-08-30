package userinfo

import (
	"os"
	"testing"

	"github.com/anhnmt/go-infostealer-parser/parser/extract"
	"github.com/anhnmt/go-infostealer-parser/parser/userinfo/meta"
)

func TestExtract(t *testing.T) {
	outputDir := "./testdata/GODELESS CLOUD"

	// FOR EXAMPLE, WE REMOVE OLD DATA BEFORE EXTRACT.
	err := os.RemoveAll(outputDir) // BE CAREFUL!
	if err != nil {
		t.Errorf("RemoveAll() error = %v", err)
		return
	}

	files, err := extract.ExtractFile(
		"./testdata/GODELESS CLOUD.rar",
		outputDir,
	)
	if err != nil {
		t.Fatal(err)
	}

	// outputDir2 := "./testdata/@berserklogs - 345 LOGS JANUARY"
	//
	// // FOR EXAMPLE, WE REMOVE OLD DATA BEFORE EXTRACT.
	// err = os.RemoveAll(outputDir2) // BE CAREFUL!
	// if err != nil {
	// 	t.Errorf("RemoveAll() error = %v", err)
	// 	return
	// }
	//
	// files2, err := parser.ExtractFile(
	// 	"./testdata/@berserklogs - 345 LOGS JANUARY.rar",
	// 	outputDir,
	// )
	// if err != nil {
	// 	t.Fatal(err)
	// }

	type args struct {
		files     []string
		fn        Parser
		outputDir string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sample GODELESS CLOUD",
			args: args{
				files:     files,
				fn:        meta.ExtractUserInfo,
				outputDir: outputDir,
			},
			want: 538,
		},
		// {
		//     name: "Sample berserklogs",
		//     args: args{
		//         files:     files2,
		//         fn:        meta.ExtractUserInfo,
		//         outputDir: outputDir2,
		//     },
		//     want: 16,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Extract(tt.args.files, tt.args.fn); tt.want != len(got) {
				t.Errorf("Extract() = %d, want %d", len(got), tt.want)
			}
		})
	}
}
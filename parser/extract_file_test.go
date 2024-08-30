package parser

import (
	"os"
	"testing"
)

func TestExtractFile(t *testing.T) {
	type args struct {
		filePath  string
		outputDir string
		passwords []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Sample Extract",
			args: args{
				filePath:  "testdata/SAMPLE_EXTRACT.zip",
				outputDir: "testdata/SAMPLE_EXTRACT",
			},
			want:    2,
			wantErr: false,
		},
		{
			name: "GODELESS CLOUD",
			args: args{
				filePath:  "testdata/GODELESS CLOUD.rar",
				outputDir: "testdata/GODELESS CLOUD",
			},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// FOR EXAMPLE, WE REMOVE OLD DATA BEFORE EXTRACT.
			err := os.RemoveAll(tt.args.outputDir) // BE CAREFUL!
			if err != nil {
				t.Errorf("RemoveAll() error = %v", err)
				return
			}

			got, err := ExtractFile(tt.args.filePath, tt.args.outputDir, tt.args.passwords...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Extract() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			l := len(got)
			if l != tt.want {
				t.Errorf("Extract() got(%d) = %v, want %v", l, got, tt.want)
			}
		})
	}
}

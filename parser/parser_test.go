package parser

import (
	"os"
	"testing"
)

func TestParseCredentials(t *testing.T) {
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
			want:    3,
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

			got, err := ParseCredentials(tt.args.filePath, tt.args.outputDir, tt.args.passwords...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCredentials() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != len(got) {
				t.Errorf("ParseCredentials() got = %v, want %v", len(got), tt.want)
			}
		})
	}
}

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
		{
			name: "Sample GODELESS CLOUD",
			args: args{
				filePath:  "testdata/GODELESS CLOUD.rar",
				outputDir: "testdata/GODELESS CLOUD",
			},
			want: 29647,
		},
		{
			name: "Sample MANTICORECLOUD",
			args: args{
				filePath:  "testdata/@MANTICORECLOUD - 14.08 - 3800 PCS.rar",
				outputDir: "testdata/@MANTICORECLOUD - 14.08 - 3800 PCS",
			},
			want: 3783,
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

func TestParseUserInfo(t *testing.T) {
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
			name: "Sample GODELESS CLOUD",
			args: args{
				filePath:  "testdata/GODELESS CLOUD.rar",
				outputDir: "testdata/GODELESS CLOUD",
			},
			want: 538,
		},
		{
			name: "Sample MANTICORECLOUD",
			args: args{
				filePath:  "testdata/@MANTICORECLOUD - 06.09 - 3500 PCS.rar",
				outputDir: "testdata/@MANTICORECLOUD - 06.09 - 3500 PCS",
			},
			want: 247100,
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

			got, err := ParseUserInfo(tt.args.filePath, tt.args.outputDir, tt.args.passwords...)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != len(got) {
				t.Errorf("ParseUserInfo() got = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestParser(t *testing.T) {
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
		// {
		//     name: "Sample GODELESS CLOUD",
		//     args: args{
		//         filePath:  "testdata/GODELESS CLOUD.rar",
		//         outputDir: "testdata/GODELESS CLOUD",
		//     },
		//     want: 538,
		// },
		// {
		// 	name: "Sample MANTICORECLOUD",
		// 	args: args{
		// 		filePath:  "testdata/@MANTICORECLOUD - 06.09 - 3500 PCS.rar",
		// 		outputDir: "testdata/@MANTICORECLOUD - 06.09 - 3500 PCS",
		// 	},
		// 	want: 247100,
		// },
		{
			name: "Unknown @DumpsSlivCloud - RED PRIVATE LOGS1",
			args: args{
				filePath:  "testdata/@DumpsSlivCloud - RED PRIVATE LOGS1.rar",
				outputDir: "testdata/@DumpsSlivCloud - RED PRIVATE LOGS1",
			},
			want: 1124,
		},
		{
			name: "Unknown @DeathLogsCloud",
			args: args{
				filePath:  "testdata/@DeathLogsCloud.rar",
				outputDir: "testdata/@DeathLogsCloud",
			},
			want: 1197,
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

			got, err := Parser(tt.args.filePath, tt.args.outputDir, tt.args.passwords...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != got.Size() {
				t.Errorf("Parser() got = %v, want %v", got.Size(), tt.want)
			}
		})
	}
}

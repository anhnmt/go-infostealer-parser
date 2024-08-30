package meta

import (
	"os"
	"testing"
)

func TestExtractCredentials(t *testing.T) {
	filePath := "testdata/SAMPLE_EXTRACT/Passwords.txt"

	// get file content
	body, err := os.ReadFile(filePath)
	if err != nil {
		t.Error(err)
		return
	}

	filePath2 := "testdata/GODELESS CLOUD/AEX22N56YL9HOLK9FJDKZIVOIDC5DFUT2_2024_08_10T13_68_80_585685]/Passwords.txt"

	// get file content
	body2, err := os.ReadFile(filePath2)
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		filePath string
		body     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Sample Credentials",
			args: args{
				filePath: filePath,
				body:     string(body),
			},
		},
		{
			name: "Sample Credentials 2",
			args: args{
				filePath: filePath2,
				body:     string(body2),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractCredentials(tt.args.filePath, tt.args.body); tt.want && len(got) > 0 {
				t.Errorf("ExtractCredentials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractCredentialsPasswordFile(t *testing.T) {
	filePath := "testdata/SAMPLE/GODELESS CLOUD.txt"

	// get file content
	body, err := os.ReadFile(filePath)
	if err != nil {
		t.Error(err)
		return
	}

	type args struct {
		filePath string
		body     string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sample Credentials",
			args: args{
				filePath: filePath,
				body:     string(body),
			},
			want: 52,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractCredentialsPasswordFile(tt.args.filePath, tt.args.body); tt.want != len(got) {
				t.Errorf("ExtractCredentialsPasswordFile() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

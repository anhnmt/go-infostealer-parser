package meta

import (
	"os"
	"testing"
)

func TestExtractUserInfo(t *testing.T) {
	filePath := "testdata/SAMPLE_EXTRACT/UserInformation.txt"

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
		want bool
	}{
		{
			name: "Sample Meta",
			args: args{
				filePath: filePath,
				body:     string(body),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractUserInfo(tt.args.filePath, tt.args.body); tt.want != got.Valid() {
				t.Errorf("ExtractUserInfo() = %v", got)
			}
		})
	}
}

package unknown

import (
	"os"
	"testing"
)

func TestExtractUserInfo(t *testing.T) {
	filePath := "testdata/@DumpsSlivCloud - RED PRIVATE LOGS1/PRIVATE/AE5SP7UYEM4YG9YQW2U2M7J8QG0YZGSSQ_2024_08_08T16_82_42_204634/UserInformation.txt"

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
			name: "Sample Unknown",
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

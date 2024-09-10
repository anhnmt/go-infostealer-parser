package unknown

import (
	"os"
	"testing"
)

func TestExtractCredentials(t *testing.T) {
	filePath := "testdata/@DumpsSlivCloud - RED PRIVATE LOGS1/PRIVATE/AE5SP7UYEM4YG9YQW2U2M7J8QG0YZGSSQ_2024_08_08T16_82_42_204634/Passwords.txt"

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
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractCredentials(tt.args.filePath, tt.args.body); tt.want != len(got) {
				t.Errorf("ExtractCredentials() = %v, want %v", len(got), tt.want)
			}
		})
	}
}

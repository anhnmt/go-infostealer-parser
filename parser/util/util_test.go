package util

import (
	"reflect"
	"testing"
)

func TestGetMatchSubString(t *testing.T) {
	type args struct {
		pattern string
		line    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Example https",
			args: args{
				pattern: "(.+):(.+):(.+)",
				line:    "https://www.instagram.com/accounts/login/:fffacu_arcee:ffacu_arcee0",
			},
			want: []string{
				"https://www.instagram.com/accounts/login/:fffacu_arcee:ffacu_arcee0",
				"https://www.instagram.com/accounts/login/",
				"fffacu_arcee",
				"ffacu_arcee0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMatchSubString(tt.args.pattern, tt.args.line); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMatchAllString() = %v, want %v", got, tt.want)
			}
		})
	}
}

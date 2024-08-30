package util

import (
	"testing"
)

func TestGetHostFromUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example.com",
			args: args{
				url: "https://example.com",
			},
			want: "example.com",
		},
		{
			name: "my.account.sony.com",
			args: args{
				url: "https://my.account.sony.com/sonyacct/signin/",
			},
			want: "my.account.sony.com",
		},
		{
			name: "com.roblox.client",
			args: args{
				url: "android://RGlUtI9NY0ps7eW1mdYoROkaZ3iIqThRr1OIJOwe5lqdRX93aUt2TxUUz13PLlTFN5B1C0mMDPyM4BsBic8Fmg==@com.roblox.client/",
			},
			want: "client.roblox.com",
		},
		{
			name: "LegacyGeneric:target=Microsoft OneDrive Generic Data - Personal Vault VHD Info",
			args: args{
				url: "LegacyGeneric:target=Microsoft OneDrive Generic Data - Personal Vault VHD Info",
			},
			want: "LegacyGeneric:target=Microsoft OneDrive Generic Data - Personal Vault VHD Info",
		},
		{
			name: "iforgot.apple.com",
			args: args{
				url: "https://iforgot.apple.com/password/reset",
			},
			want: "iforgot.apple.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHostFromUrl(tt.args.url); got != tt.want {
				t.Errorf("GetHostFromUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

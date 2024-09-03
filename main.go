// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"os"

	"github.com/anhnmt/go-infostealer-parser/parser/util"
)

func main() {
	// body := `adgroup\rhuerta`
	//
	// check := util.ValidString(body)
	// fmt.Println(check)
	//
	// body = strings.ToValidUTF8(body, "")
	// body = strings.Map(func(r rune) rune {
	//     if unicode.IsPrint(r) {
	//         return r
	//     }
	//     return -1
	// }, body)
	// body = strings.TrimSpace(body)

	// content, err := os.ReadFile("testdata/GODELESS CLOUD/AEX22N56YL9HOLK9FJDKZIVOIDC5DFUT2_2024_08_10T13_68_80_585685]/UserInformation.txt")
	// content, err := os.ReadFile("testdata/GODELESS CLOUD/CZ[ZJ2WSECMWMSUSZPGASFWNWR6V6FR2NNE] [2024_08_10T07_20_02]/UserInformation.txt")
	content, err := os.ReadFile("testdata/@MANTICORECLOUD - 14.08 - 3800 PCS/@MANTICORECLOUD - 14.08 - 3800 PCS/AE[8FZLH5LQVKSY5YEE4S5P6QQ4YEKEKNN751] [2024-08-13T01_59_37.8361402]/UserInformation.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(util.GetMatchStealerHeader(util.ManticoreHeader, string(content)))
}

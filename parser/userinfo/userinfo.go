package userinfo

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"

	"github.com/anhnmt/go-infostealer-parser/parser/model"
)

var UserInformation = []string{
	"UserInformation.txt",
}

type Parser func(filePath, body string) *model.UserInformation

func Extract(files []string, fn Parser) []*model.UserInformation {
	if len(files) == 0 {
		return nil
	}

	results := make([]*model.UserInformation, 0)
	lo.ForEach(files, func(file string, _ int) {
		fileName := filepath.Base(file)

		lo.ForEach(UserInformation, func(item string, _ int) {
			if !strings.EqualFold(fileName, item) {
				return
			}

			// get file content
			body, err := os.ReadFile(file)
			if err != nil {
				return
			}

			userInfo := fn(file, string(body))
			if userInfo != nil {
				results = append(results, userInfo)
			}

		})
	})

	return results
}

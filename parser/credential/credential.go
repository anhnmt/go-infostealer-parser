package credential

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/samber/lo"

	"github.com/anhnmt/go-infostealer-parser/parser/model"
)

var Passwords = []string{
	"Passwords.txt",
}

type Parser func(filePath, body string) []*model.Credential

func Extract(files []string, fn Parser) []*model.Credential {
	if len(files) == 0 {
		return nil
	}

	results := make([]*model.Credential, 0)
	lo.ForEach(files, func(file string, _ int) {
		fileName := filepath.Base(file)

		lo.ForEach(Passwords, func(item string, _ int) {
			if !strings.EqualFold(fileName, item) {
				return
			}

			// get file content
			body, err := os.ReadFile(file)
			if err != nil {
				return
			}

			credentials := fn(file, string(body))
			if len(credentials) > 0 {
				results = append(results, credentials...)
			}

		})
	})

	return results
}

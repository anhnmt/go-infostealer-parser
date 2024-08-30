package meta

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/samber/lo"

	"github.com/anhnmt/go-infostealer-parser/parser/model"
	"github.com/anhnmt/go-infostealer-parser/parser/util"
)

// ExtractCredentials extracts Credentials pattern from body
func ExtractCredentials(filePath, body string) []*model.Credential {
	entries := strings.Split(body, "===============")
	if len(entries) == 0 {
		return ExtractCredentialsPasswordFile(filePath, body)
	}

	credentials := make([]*model.Credential, 0)
	lo.ForEach(entries, func(entry string, _ int) {
		// Split the entry into individual lines
		lines := strings.Split(entry, "\n")
		if len(lines) == 0 {
			return
		}

		credential := &model.Credential{
			OutputDir: filepath.Dir(filePath),
		}

		lo.ForEach(lines, func(line string, _ int) {
			line = strings.TrimSpace(line)
			if len(line) == 0 ||
				strings.HasPrefix(line, "*") ||
				strings.HasPrefix(line, "http") {
				return
			}

			// URL
			if val := util.GetMatchString(URL, line); val != "" {
				credential.URL = val
				credential.Host = util.GetHostFromUrl(val)
				return
			}

			// Username
			if val := util.GetMatchString(Username, line); val != "" {
				if val == "UNKNOWN" {
					return
				}

				credential.Username = val
				return
			}

			// Password
			if val := util.GetMatchString(Password, line); val != "" {
				credential.Password = val
				return
			}

			// Application
			if val := util.GetMatchString(Application, line); val != "" {
				credential.Application = val
				return
			}
		})

		// Validate
		if !credential.Valid() {
			return
		}

		credentials = append(credentials, credential)

		fmt.Printf(
			"host: %s\nurl: %s\nusername: %s\npassword: %s\napplication: %s\n\n",
			credential.Host,
			credential.URL,
			credential.Username,
			credential.Password,
			credential.Application,
		)
	})

	return credentials
}

func ExtractCredentialsPasswordFile(filePath, body string) []*model.Credential {
	credentials := make([]*model.Credential, 0)
	lines := strings.Split(body, "\n")
	if len(lines) == 0 {
		return nil
	}

	credential := &model.Credential{
		OutputDir: filepath.Dir(filePath),
	}

	lo.ForEach(lines, func(line string, _ int) {
		line = strings.TrimSpace(line)
		if len(line) == 0 || strings.HasPrefix(line, "*") {
			return
		}

		// Password File
		matches := util.GetMatchSubString(PasswordFile, line)
		if len(matches) != 4 {
			return
		}

		url := util.GetGroupValue(matches, 1)
		credential.URL = url
		credential.Host = util.GetHostFromUrl(url)
		credential.Username = util.GetGroupValue(matches, 2)
		credential.Password = util.GetGroupValue(matches, 3)

		// Validate
		if !credential.Valid() {
			return
		}

		credentials = append(credentials, credential)

		fmt.Printf(
			"host: %s\nurl: %s\nusername: %s\npassword: %s\n\n",
			credential.Host,
			credential.URL,
			credential.Username,
			credential.Password,
		)
	})

	return credentials
}
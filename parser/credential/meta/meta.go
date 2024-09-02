package meta

import (
	"path/filepath"
	"strings"

	"github.com/samber/lo"

	"github.com/anhnmt/go-infostealer-parser/parser/model"
	"github.com/anhnmt/go-infostealer-parser/parser/util"
)

const (
	URL         = "(URL|url):"
	Username    = "(Username|username):"
	Password    = "(Password|password):"
	Application = "(Application|application):"
)

// ExtractCredentials extracts Credentials pattern from body
func ExtractCredentials(filePath, body string) []*model.Credential {
	entries := strings.Split(body, "===============")
	if len(entries) == 0 {
		return nil
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
			line = util.TrimString(line)
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

		// fmt.Printf(
		//     "host: %s\nurl: %s\nusername: %s\npassword: %s\napplication: %s\n\n",
		//     credential.Host,
		//     credential.URL,
		//     credential.Username,
		//     credential.Password,
		//     credential.Application,
		// )
	})

	return credentials
}

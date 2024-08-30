package util

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/samber/lo"
)

// FilterWhitelistFiles returns files that match whitelist
func FilterWhitelistFiles(files []string) []string {
	return lo.Filter(files, func(file string, _ int) bool {
		fileName := filepath.Base(file)

		return lo.ContainsBy(WhitelistFiles, func(item string) bool {
			return strings.EqualFold(item, fileName)
		})
	})
}

// GetMatchString returns matched string
func GetMatchString(pattern, line string) string {
	re := regexp.MustCompile(pattern)
	if re.MatchString(line) {
		return strings.TrimSpace(re.ReplaceAllString(line, ""))
	}

	return ""
}

// GetMatchSubString returns matched sub string
func GetMatchSubString(pattern, line string) []string {
	re := regexp.MustCompile(pattern)
	if !re.MatchString(line) {
		return nil
	}

	return re.FindStringSubmatch(line)
}

func GetGroupValue(group []string, index int) string {
	if index <= len(group)-1 {
		return strings.TrimSpace(group[index])
	}

	return ""
}

func GetMatchStealerHeader(header, body string) bool {
	flag := true
	for _, line := range strings.Split(strings.TrimSpace(header), "\n") {
		re := regexp.MustCompile(regexp.QuoteMeta(strings.TrimSpace(line)))
		flag = flag && re.MatchString(body)
	}
	return flag
}

func HandlerExtract(files []string, whitelist []string, fn func(string, string)) {
	if len(files) == 0 {
		return
	}

	lo.ForEach(files, func(file string, _ int) {
		fileName := filepath.Base(file)

		lo.ForEach(whitelist, func(item string, _ int) {
			if !strings.EqualFold(fileName, item) {
				return
			}

			// get file content
			contents, err := os.ReadFile(file)
			if err != nil {
				return
			}

			fn(file, string(contents))
		})
	})
}

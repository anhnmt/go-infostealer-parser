package util

import (
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

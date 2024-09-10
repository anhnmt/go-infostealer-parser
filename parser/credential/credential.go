package credential

import (
	"github.com/anhnmt/go-infostealer-parser/parser/credential/meta"
	"github.com/anhnmt/go-infostealer-parser/parser/credential/unknown"
	"github.com/anhnmt/go-infostealer-parser/parser/model"
	"github.com/anhnmt/go-infostealer-parser/parser/util"
)

type Parser func(filePath, body string) []*model.Credential

func Extract(files []string, fn Parser) []*model.Credential {
	if len(files) == 0 {
		return nil
	}

	results := make([]*model.Credential, 0)
	util.HandlerExtract(files, util.Passwords, func(file string, body string) {
		credentials := fn(file, body)
		if len(credentials) > 0 {
			results = append(results, credentials...)
		}
	})

	if len(files) == 0 {
		return nil
	}

	return results
}

func DetectStealer(files []string) []*model.Credential {
	if len(files) == 0 {
		return nil
	}

	results := make([]*model.Credential, 0)
	util.HandlerExtract(files, util.Passwords, func(file string, body string) {
		if util.GetMatchStealerHeader(util.MetaHeader, body) ||
			util.GetMatchStealerHeader(util.RedlineHeader, body) ||
			util.GetMatchStealerHeader(util.BradMaxHeader, body) ||
			util.GetMatchStealerHeader(util.ManticoreHeader, body) {
			credentials := meta.ExtractCredentials(file, body)
			if len(credentials) > 0 {
				results = append(results, credentials...)
			}
			return
		}

		// Unknown stealer
		credentials := unknown.ExtractCredentials(file, body)
		if len(credentials) > 0 {
			results = append(results, credentials...)
		}
	})

	return results
}

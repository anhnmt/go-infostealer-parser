package userinfo

import (
	"github.com/anhnmt/go-infostealer-parser/parser/model"
	"github.com/anhnmt/go-infostealer-parser/parser/userinfo/meta"
	"github.com/anhnmt/go-infostealer-parser/parser/util"
)

type Parser func(filePath, body string) *model.UserInformation

func Extract(files []string, fn Parser) []*model.UserInformation {
	if len(files) == 0 {
		return nil
	}

	results := make([]*model.UserInformation, 0)
	util.HandlerExtract(files, util.UserInformation, func(file string, body string) {
		userInfo := fn(file, body)
		if userInfo != nil {
			results = append(results, userInfo)
		}
	})

	return results
}

func DetectStealer(files []string) []*model.UserInformation {
	if len(files) == 0 {
		return nil
	}

	results := make([]*model.UserInformation, 0)
	util.HandlerExtract(files, util.UserInformation, func(file string, body string) {
		if util.GetMatchStealerHeader(util.MetaHeader, body) ||
			util.GetMatchStealerHeader(util.RedlineHeader, body) ||
			util.GetMatchStealerHeader(util.BradMaxHeader, body) {
			userInfo := meta.ExtractUserInfo(file, body)
			if userInfo != nil {
				results = append(results, userInfo)
			}
			return
		}
	})

	return results
}

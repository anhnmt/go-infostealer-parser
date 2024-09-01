package parser

import (
	"github.com/samber/lo"

	"github.com/anhnmt/go-infostealer-parser/parser/credential"
	"github.com/anhnmt/go-infostealer-parser/parser/extract"
	"github.com/anhnmt/go-infostealer-parser/parser/model"
	"github.com/anhnmt/go-infostealer-parser/parser/userinfo"
)

type Stealer struct {
	UserInfo    *model.UserInformation
	Credentials []*model.Credential
}

func Parse(filePath, outputDir string, passwords ...string) (map[string]*Stealer, error) {
	files, err := extract.ExtractFile(filePath, outputDir, passwords...)
	if err != nil {
		return nil, err
	}

	results := make(map[string]*Stealer)

	userInfos := userinfo.DetectStealer(files)
	if len(userInfos) > 0 {
		groupUserInfos := lo.GroupBy(userInfos, func(userInfo *model.UserInformation) string {
			return userInfo.OutputDir
		})

		for group, slice := range groupUserInfos {
			if len(slice) == 0 {
				continue
			}

			val, ok := results[group]
			if !ok {
				val = &Stealer{
					UserInfo: slice[0],
				}
				results[group] = val
			} else {
				val.UserInfo = slice[0]
				results[group] = val
			}
		}
	}

	credentials := credential.DetectStealer(files)
	if len(credentials) > 0 {
		groupCredentials := lo.GroupBy(credentials, func(credential *model.Credential) string {
			return credential.OutputDir
		})

		for group, slice := range groupCredentials {
			if len(slice) == 0 {
				continue
			}

			val, ok := results[group]
			if !ok {
				val = &Stealer{
					Credentials: slice,
				}
				results[group] = val
			} else {
				val.Credentials = append(val.Credentials, slice...)
				results[group] = val
			}
		}
	}

	return results, nil
}

func ParseCredentials(filePath, outputDir string, passwords ...string) ([]*model.Credential, error) {
	files, err := extract.ExtractFile(filePath, outputDir, passwords...)
	if err != nil {
		return nil, err
	}

	return credential.DetectStealer(files), nil
}

func ParseUserInfo(filePath, outputDir string, passwords ...string) ([]*model.UserInformation, error) {
	files, err := extract.ExtractFile(filePath, outputDir, passwords...)
	if err != nil {
		return nil, err
	}

	return userinfo.DetectStealer(files), nil
}

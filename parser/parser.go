package parser

import (
	"github.com/anhnmt/go-infostealer-parser/parser/credential"
	"github.com/anhnmt/go-infostealer-parser/parser/extract"
	"github.com/anhnmt/go-infostealer-parser/parser/model"
	"github.com/anhnmt/go-infostealer-parser/parser/userinfo"
)

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

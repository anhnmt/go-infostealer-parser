package parser

import (
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"golang.org/x/sync/errgroup"

	"github.com/anhnmt/go-infostealer-parser/parser/credential"
	"github.com/anhnmt/go-infostealer-parser/parser/extract"
	"github.com/anhnmt/go-infostealer-parser/parser/model"
	"github.com/anhnmt/go-infostealer-parser/parser/userinfo"
)

type InfoStealer struct {
	UserInfo    *model.UserInformation
	Credentials []*model.Credential
}

func Parse(filePath, outputDir string, passwords ...string) (map[string]*InfoStealer, error) {
	files, err := extract.ExtractFile(filePath, outputDir, passwords...)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, nil
	}

	results := make(map[string]*InfoStealer)
	chunkCh := make(chan []string, 100)
	maxWorkers := 10
	g := errgroup.Group{}
	g.SetLimit(maxWorkers)

	for range maxWorkers {
		g.Go(func() error {
			for chunks := range chunkCh {
				userInfos := userinfo.DetectStealer(chunks)
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
							val = &InfoStealer{
								UserInfo: slice[0],
							}
							results[group] = val
						} else {
							val.UserInfo = slice[0]
							results[group] = val
						}
					}
				}

				credentials := credential.DetectStealer(chunks)
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
							val = &InfoStealer{
								Credentials: slice,
							}
							results[group] = val
						} else {
							val.Credentials = append(val.Credentials, slice...)
							results[group] = val
						}
					}
				}
			}

			return nil
		})
	}

	chunks := lo.Chunk(files, 100)
	lop.ForEach(chunks, func(chunk []string, _ int) {
		chunkCh <- chunk
	})

	close(chunkCh)
	g.Wait()

	return results, nil
}

func ParseCredentials(filePath, outputDir string, passwords ...string) ([]*model.Credential, error) {
	files, err := extract.ExtractFile(filePath, outputDir, passwords...)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, nil
	}

	return ParseCredentialsFromFiles(files...)
}

func ParseCredentialsFromFiles(files ...string) ([]*model.Credential, error) {
	results := make([]*model.Credential, 0)
	chunkCh := make(chan []string, 100)
	maxWorkers := 10
	g := errgroup.Group{}
	g.SetLimit(maxWorkers)

	for range maxWorkers {
		g.Go(func() error {
			for chunks := range chunkCh {
				credentials := credential.DetectStealer(chunks)
				if len(credentials) == 0 {
					continue
				}

				results = append(results, credentials...)
			}

			return nil
		})
	}

	chunks := lo.Chunk(files, 100)
	lop.ForEach(chunks, func(chunk []string, _ int) {
		chunkCh <- chunk
	})

	close(chunkCh)
	g.Wait()

	return results, nil
}

func ParseUserInfo(filePath, outputDir string, passwords ...string) ([]*model.UserInformation, error) {
	files, err := extract.ExtractFile(filePath, outputDir, passwords...)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, nil
	}

	return ParseUserInfoFromFiles(files...)
}

func ParseUserInfoFromFiles(files ...string) ([]*model.UserInformation, error) {
	results := make([]*model.UserInformation, 0)
	chunkCh := make(chan []string, 100)
	maxWorkers := 10
	g := errgroup.Group{}
	g.SetLimit(maxWorkers)

	for range maxWorkers {
		g.Go(func() error {
			for chunks := range chunkCh {
				userInfos := userinfo.DetectStealer(chunks)
				if len(userInfos) == 0 {
					continue
				}

				results = append(results, userInfos...)
			}

			return nil
		})
	}

	chunks := lo.Chunk(files, 100)
	lop.ForEach(chunks, func(chunk []string, _ int) {
		chunkCh <- chunk
	})

	close(chunkCh)
	g.Wait()

	return results, nil
}

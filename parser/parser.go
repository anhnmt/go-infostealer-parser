package parser

import (
	"sync/atomic"

	"github.com/puzpuzpuz/xsync/v3"
	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
	"golang.org/x/sync/errgroup"

	"github.com/anhnmt/go-infostealer-parser/parser/credential"
	"github.com/anhnmt/go-infostealer-parser/parser/extract"
	"github.com/anhnmt/go-infostealer-parser/parser/model"
	"github.com/anhnmt/go-infostealer-parser/parser/userinfo"
)

const ChunkSize = 100

var MaxWorkers atomic.Int32

func init() {
	MaxWorkers.Store(5)
}

type InfoStealer struct {
	UserInfo    *model.UserInformation
	Credentials []*model.Credential
}

func Parse(filePath, outputDir string, passwords ...string) (*xsync.MapOf[string, *InfoStealer], error) {
	files, err := extract.ExtractFile(filePath, outputDir, passwords...)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, nil
	}

	results := xsync.NewMapOf[string, *InfoStealer]()
	chunkCh := make(chan []string, ChunkSize)
	g := errgroup.Group{}
	maxWorkers := int(MaxWorkers.Load())
	g.SetLimit(maxWorkers)

	for range maxWorkers {
		g.Go(func() error {
			for chunks := range chunkCh {
				userInfos := userinfo.DetectStealer(chunks)
				if len(userInfos) > 0 {
					for _, userInfo := range userInfos {
						group := userInfo.OutputDir
						val, ok := results.Load(group)
						if !ok {
							val = &InfoStealer{
								UserInfo: userInfo,
							}
						} else {
							val.UserInfo = userInfo
						}
						results.Store(group, val)
					}
				}

				credentials := credential.DetectStealer(chunks)
				if len(credentials) > 0 {
					for _, credential := range credentials {
						group := credential.OutputDir
						val, ok := results.Load(group)
						if !ok {
							val = &InfoStealer{
								Credentials: []*model.Credential{
									credential,
								},
							}
						} else {
							val.Credentials = append(val.Credentials, credential)
						}
						results.Store(group, val)
					}
				}
			}

			return nil
		})
	}

	lop.ForEach(lo.Chunk(files, ChunkSize), func(chunk []string, _ int) {
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
	chunkCh := make(chan []string, ChunkSize)
	g := errgroup.Group{}
	maxWorkers := int(MaxWorkers.Load())
	g.SetLimit(maxWorkers)

	for range maxWorkers {
		g.Go(func() error {
			for chunks := range chunkCh {
				credentials := credential.DetectStealer(chunks)
				if len(credentials) > 0 {
					results = append(results, credentials...)
				}
			}

			return nil
		})
	}

	lop.ForEach(lo.Chunk(files, ChunkSize), func(chunk []string, _ int) {
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
	chunkCh := make(chan []string, ChunkSize)
	g := errgroup.Group{}
	maxWorkers := int(MaxWorkers.Load())
	g.SetLimit(maxWorkers)

	for range maxWorkers {
		g.Go(func() error {
			for chunks := range chunkCh {
				userInfos := userinfo.DetectStealer(chunks)
				if len(userInfos) > 0 {
					results = append(results, userInfos...)
				}
			}

			return nil
		})
	}

	lop.ForEach(lo.Chunk(files, ChunkSize), func(chunk []string, _ int) {
		chunkCh <- chunk
	})

	close(chunkCh)
	g.Wait()

	return results, nil
}

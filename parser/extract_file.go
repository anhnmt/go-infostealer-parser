package parser

import (
	"path/filepath"

	"golift.io/xtractr"

	"github.com/anhnmt/go-infostealer-parser/parser/util"
)

func ExtractFile(filePath, outputDir string, passwords ...string) ([]string, error) {
	files, err := ExtractFileWithoutFilter(filePath, outputDir, passwords...)
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		return nil, nil
	}

	return util.FilterWhitelistFiles(files), err
}

func ExtractFileWithoutFilter(filePath, outputDir string, passwords ...string) ([]string, error) {
	x := &xtractr.XFile{
		FilePath:  filepath.Clean(filePath),
		OutputDir: filepath.Clean(outputDir), // do not forget this.
		Passwords: passwords,
	}

	// size is how many bytes were written.
	// files may be nil, but will contain any files written (even with an error).
	_, files, _, err := xtractr.ExtractFile(x)
	if err != nil {
		return nil, err
	}

	return files, err
}

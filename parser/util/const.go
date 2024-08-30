package util

var WhitelistFiles = append(
	Passwords,
	UserInformation...,
)

var Passwords = []string{
	"Passwords.txt",
}

var UserInformation = []string{
	"UserInformation.txt",
}

const (
	MetaHeader = `
*               _   _   _   _                 *
*              / \ / \ / \ / \                *
*             ( M | E | T | A )               *
*              \_/ \_/ \_/ \_/                *
`

	RedlineHeader = `
*   ____  _____ ____  _     ___ _   _ _____   *
*  |  _ \| ____|  _ \| |   |_ _| \ | | ____|  *
*  | |_) |  _| | | | | |    | ||  \| |  _|    *
*  |  _ <| |___| |_| | |___ | || |\  | |___   *
*  |_| \_|_____|____/|_____|___|_| \_|_____|  *
`
)

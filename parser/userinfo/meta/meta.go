package meta

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/samber/lo"

	"github.com/anhnmt/go-infostealer-parser/parser/model"
	"github.com/anhnmt/go-infostealer-parser/parser/util"
)

const (
	IP           string = "(IP|ip):"
	FileLocation        = "(FileLocation):"
	UserName            = "(UserName):"
	MachineName         = "(MachineName):"
	MachineID           = "(MachineID):"
	Country             = "(Country):"
	Location            = "(Location):"
	HWID                = "(HWID):"
	OS                  = "(Operation System):"
	LogDate             = "(Log date):"
)

// ExtractUserInfo extracts UserInformation pattern from body
func ExtractUserInfo(filePath, body string) *model.UserInformation {
	lines := strings.Split(body, "\n")
	if len(lines) == 0 {
		return nil
	}

	userInfo := &model.UserInformation{
		OutputDir: filepath.Dir(filePath),
	}

	lo.ForEach(lines, func(line string, _ int) {
		line = strings.TrimSpace(line)
		if len(line) == 0 ||
			strings.HasPrefix(line, "*") ||
			strings.HasPrefix(line, "http") {
			return
		}

		// IP
		if val := util.GetMatchString(IP, line); val != "" {
			userInfo.IP = val
			return
		}

		// FileLocation
		if val := util.GetMatchString(FileLocation, line); val != "" {
			userInfo.FileLocation = val
			return
		}

		// Location
		// Should check location after file location
		if val := util.GetMatchString(Location, line); val != "" {
			userInfo.Location = val
			return
		}

		// UserName
		if val := util.GetMatchString(UserName, line); val != "" {
			userInfo.UserName = val
			return
		}

		// MachineName
		if val := util.GetMatchString(MachineName, line); val != "" {
			userInfo.MachineName = val
			return
		}

		// MachineID
		if val := util.GetMatchString(MachineID, line); val != "" {
			userInfo.MachineID = val
			return
		}

		// Country
		if val := util.GetMatchString(Country, line); val != "" {
			userInfo.Country = val
			return
		}

		// HWID
		if val := util.GetMatchString(HWID, line); val != "" {
			userInfo.HWID = val
			return
		}

		// OS
		if val := util.GetMatchString(OS, line); val != "" {
			userInfo.OS = val
			return
		}

		// LogDate
		if val := util.GetMatchString(LogDate, line); val != "" {
			logDate, err := time.Parse("01/02/2006 15:04:05", val)
			if err == nil {
				userInfo.LogDate = &logDate
				return
			}

			logDate, err = time.Parse("1/02/2006 15:04:05 PM", val)
			if err == nil {
				userInfo.LogDate = &logDate
				return
			}

			logDate, err = time.Parse("1/2/2006 15:04:05 PM", val)
			if err == nil {
				userInfo.LogDate = &logDate
				return
			}

			logDate, err = time.Parse(time.ANSIC, val)
			if err == nil {
				userInfo.LogDate = &logDate
				return
			}

			return
		}
	})

	// Validate
	if !userInfo.Valid() {
		return nil
	}

	fmt.Printf(
		"ip: %s\nfile_location: %s\nuser_name: %s\nmachine_name: %s\nmachine_id: %s\ncountry: %s\nlocation: %s\nhwid: %s\nos: %s\nlog_date: %s\n\n",
		userInfo.IP,
		userInfo.FileLocation,
		userInfo.UserName,
		userInfo.MachineName,
		userInfo.MachineID,
		userInfo.Country,
		userInfo.Location,
		userInfo.HWID,
		userInfo.OS,
		userInfo.LogDate,
	)

	return userInfo
}

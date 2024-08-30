package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// UserInformation is a model for user_information.txt file
type UserInformation struct {
	IP           string     `validate:"required" json:"ip_address"`
	OS           string     `json:"operating_system"`
	FileLocation string     `json:"file_location"`
	Country      string     `json:"country"`
	Location     string     `json:"location"`
	HWID         string     `json:"hwid"`
	LogDate      *time.Time `json:"log_date"`
	UserName     string     `json:"user_name"`
	MachineName  string     `json:"machine_name"`
	MachineID    string     `json:"machine_id"`
	OutputDir    string     `json:"output_dir"`
}

// Valid returns true if credential is valid
func (u *UserInformation) Valid() bool {
	if u == nil {
		return false
	}

	// Validate
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(u) == nil
}

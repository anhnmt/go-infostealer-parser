package model

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// Credential is a model for password.txt file
type Credential struct {
	Host        string `json:"host"`
	URL         string `validate:"required" json:"url"`
	Username    string `validate:"required" json:"username"`
	Password    string `validate:"required" json:"password"`
	Application string `json:"application"`
	OutputDir   string `json:"output_dir"`
}

// Valid returns true if credential is valid
func (c *Credential) Valid() bool {
	if c == nil {
		return false
	}

	// Trim space
	c.TrimSpace()

	// Validate
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(c) == nil
}

func (c *Credential) TrimSpace() {
	if c == nil {
		return
	}

	c.Host = strings.TrimSpace(c.Host)
	c.URL = strings.TrimSpace(c.URL)
	c.Username = strings.TrimSpace(c.Username)
	c.Password = strings.TrimSpace(c.Password)
	c.Application = strings.TrimSpace(c.Application)
}

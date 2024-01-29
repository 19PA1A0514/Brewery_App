package model

import (
	"errors"
	"regexp"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string    `validate:"required,min=4,max=32" json:"firstname"`
	Email       string    `validate:"required,email" json:"email"`
	Password    string    `validate:"required" json:"password"`
	LastName    string    `validate:"required,min=4,max=20" json:"lastname"`
	DateOfBirth time.Time `validate:"required" json:"dateofbirth"`
}

var (
	regexContainsOnlyDigitsUnderscoreDashesDotsAlpha = regexp.MustCompile(`^[0-9a-zA-Z_\.\-]*$`)
	regexBeginsWithUnderscoreOrAlpha                 = regexp.MustCompile(`^[a-zA-Z_].*$`)
)

func ValidateUser(user *User) error {
	if !regexContainsOnlyDigitsUnderscoreDashesDotsAlpha.MatchString(user.FirstName) {
		return errors.New("username must contain only digits, underscores, dashes, dots and alphabetical letters")
	}
	if !regexBeginsWithUnderscoreOrAlpha.MatchString(user.FirstName) {
		return errors.New("username must begin with either underscore or an alphabetical letter")
	}
	return nil
}

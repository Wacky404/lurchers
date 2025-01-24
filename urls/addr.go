package url

import (
	"errors"
	"strings"
)

type Websites map[string]string

type userParams struct {
	websites Websites
	timeout  int
	retry    bool
}

func NewWebsite(w string) (string, error) {
	if strings.HasPrefix(w, "http://") || strings.HasPrefix(w, "https://") {
		return w, nil
	}

	return w, errors.New("url entered is not an http or https url")
}

func (u *userParams) Getwebsites() (map[string]string, error) {
	if u.websites == nil {
		return nil, errors.New("no websites given in user params")
	}

	return u.websites, nil
}

package kmatch

import (
	"regexp"
)

func MatchPhone(phone string) bool {
	phoneRes := regexp.MustCompile(`^(13[0-9]|14[5|7]|15[0|1|2|3|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$`)
	return phoneRes.MatchString(phone)
}

func MatchEmail(email string) bool {
	emailRes := regexp.MustCompile(`^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$`)
	return emailRes.MatchString(email)
}
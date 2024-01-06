package lib

import "regexp"

func IsUrlNotValid(url string) bool {
	matched, err := regexp.MatchString(`https:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()!@:%_\+.~#?&\/\/=]*)`, url)

	if err != nil {
		return false
	}

	return !matched
}

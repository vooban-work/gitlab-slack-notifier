package main

import "strings"

func formatFullnameToUserEmail(username string) (string, error) {
	removedAccents, err := deburr(strings.TrimSpace(username))
	if err != nil {
		return "", err
	}
	withSpaceReplaced := strings.ReplaceAll(strings.ToLower(removedAccents), " ", USER_EMAIL_SPACE_REPLACER)
	return stripNonEmailChars(withSpaceReplaced), nil
}

func stripNonEmailChars(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch {
		case r >= 'a' && r <= 'z',
			r >= '0' && r <= '9',
			r == '.', r == '-', r == '_':
			b.WriteRune(r)
		}
	}
	return b.String()
}

package slug

import "strings"

func Make(title string) string {
	slug := strings.TrimSpace(title)
	slug = strings.ToLower(slug)
	slug = strings.Replace(slug, " ", "-", -1)
	return slug
}

package slugs

import "strings"

func Make(title string) string {
	title = strings.TrimSpace(title)
	title = strings.ToLower(title)
	title = strings.Replace(title, " ", "-", -1)
	return title
}

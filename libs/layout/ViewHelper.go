// Author: James Mallon <jamesmallondev@gmail.com>
// layout package -
package layout

import (
	"strings"
	"time"
)

// Struct type ViewHelper -
type ViewHelper struct{}

// FormatDate method - format a date
func (this *ViewHelper) FormatDate(t time.Time) string {
	layout := "2006-01-01"
	return t.Format(layout)
}

// ToLower method -
func (this *ViewHelper) ToLower(text string) string {
	return strings.ToLower(text)
}

// ToUpper method -
func (this *ViewHelper) ToUpper(text string) string {
	return strings.ToUpper(text)
}

// UCFirst method -
func (this *ViewHelper) UCFirst(text string) string {
	return strings.Title(text)
}

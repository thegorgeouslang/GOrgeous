// Author: James Mallon <jamesmallondev@gmail.com>
// layout package -
package layout

import (
	"time"
)

// Struct type ViewHelper -
type ViewHelper struct{}

// FormatDate method - format a date
func (this *ViewHelper) FormatDate(t time.Time) string {
	layout := "2006-01-01"
	return t.Format(layout)
}

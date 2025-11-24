package formatting

import (
	"bytes"
	"fmt"
	"strings"
)

// OutlookHtmlProvider wraps HtmlProvider to ensure Windows Outlook compatibility
// by replacing empty table cells with &nbsp; to prevent broken icon rendering
type OutlookHtmlProvider struct {
	*HtmlProvider
}

// Table overrides HtmlProvider.Table to ensure empty cells contain &nbsp; and use Outlook-compatible CSS
func (outlook *OutlookHtmlProvider) Table(rows [][]string) string {
	table := make([]string, 0)
	// Outlook-specific styles: mso-table-lspace/rspace for proper spacing, explicit borders
	table = append(table, "<TABLE border='1' style='width: 100%; border-collapse: collapse; border-spacing: 0; mso-table-lspace: 0pt; mso-table-rspace: 0pt;'>")
	for i, r := range rows {
		var tag string
		if i == 0 {
			tag = "TH"
		} else {
			tag = "TD"
		}
		table = append(table, "<TR>")
		var rowBuilder bytes.Buffer
		for _, field := range r {
			// Replace empty strings with &nbsp; for Outlook compatibility
			if strings.TrimSpace(field) == "" {
				field = "&nbsp;"
			}
			// Outlook-specific cell styles: explicit borders with mso-border-alt for proper rendering
			rowBuilder.WriteString(fmt.Sprintf("<%s style='padding: 5px; border: 1px solid #cccccc; mso-border-alt: solid #cccccc .5pt;'>%s</%s>", tag, field, tag))
		}
		table = append(table, rowBuilder.String())
		table = append(table, "</TR>")
	}

	table = append(table, "</TABLE>\n")
	return strings.Join(table, "\n")
}

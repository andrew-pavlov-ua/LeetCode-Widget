// Code generated by qtc from "badge_no_user_found.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line internal/templates/v1/badge_no_user_found.qtpl:1
package v1

//line internal/templates/v1/badge_no_user_found.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line internal/templates/v1/badge_no_user_found.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line internal/templates/v1/badge_no_user_found.qtpl:1
func StreamBadgeNoUserFound(qw422016 *qt422016.Writer) {
//line internal/templates/v1/badge_no_user_found.qtpl:1
	qw422016.N().S(`
<svg width="500" height="200" viewBox="0 0 500 200" fill="none" xmlns="http://www.w3.org/2000/svg">
  <!-- Background -->
  <rect width="500" height="200" rx="12" fill="#1E1E1E"></rect>

  <!-- LeetCode Logo -->
  <g>
    <image src="/LeetCodeLogo.png" x="20" y="10" width="30" height="30" />
  </g>

  <!-- No User Info Text -->
  <g fill="#FFF" font-family="Verdana, sans-serif">
    <text x="250" y="100" font-size="24" font-weight="bold" text-anchor="middle" alignment-baseline="middle">No User Found</text>
  </g>

  <!-- Placeholder for Circle and Progress Bars -->
  <g>
    <circle cx="80" cy="140" r="45" fill="none" stroke="#FFA116" stroke-width="5"></circle>
    <circle cx="80" cy="140" r="40" fill="#444"></circle>
    <text x="80" y="145" font-size="28" fill="#FFF" font-family="Verdana, sans-serif" text-anchor="middle" alignment-baseline="middle" dominant-baseline="middle">N/A</text>
  </g>

  <!-- Progress Bars Placeholder -->
  <g font-family="Verdana, sans-serif">
    <text x="150" y="120" font-size="18" fill="#FFF">Easy</text>
    <rect x="240" y="110" width="220" height="10" fill="#333" rx="5"></rect>
    <rect x="240" y="110" width="0" height="10" fill="#4CAF50" rx="5"></rect>
    <text x="470" y="120" font-size="14" fill="#FFF" text-anchor="end">N/A / N/A</text>

    <text x="150" y="145" font-size="18" fill="#FFF">Medium</text>
    <rect x="240" y="135" width="220" height="10" fill="#333" rx="5"></rect>
    <rect x="240" y="135" width="0" height="10" fill="#FF9800" rx="5"></rect>
    <text x="470" y="145" font-size="14" fill="#FFF" text-anchor="end">N/A / N/A</text>

    <text x="150" y="170" font-size="18" fill="#FFF">Hard</text>
    <rect x="240" y="160" width="220" height="10" fill="#333" rx="5"></rect>
    <rect x="240" y="160" width="0" height="10" fill="#F44336" rx="5"></rect>
    <text x="470" y="170" font-size="14" fill="#FFF" text-anchor="end">N/A / N/A</text>
  </g>
</svg>
`)
//line internal/templates/v1/badge_no_user_found.qtpl:41
}

//line internal/templates/v1/badge_no_user_found.qtpl:41
func WriteBadgeNoUserFound(qq422016 qtio422016.Writer) {
//line internal/templates/v1/badge_no_user_found.qtpl:41
	qw422016 := qt422016.AcquireWriter(qq422016)
//line internal/templates/v1/badge_no_user_found.qtpl:41
	StreamBadgeNoUserFound(qw422016)
//line internal/templates/v1/badge_no_user_found.qtpl:41
	qt422016.ReleaseWriter(qw422016)
//line internal/templates/v1/badge_no_user_found.qtpl:41
}

//line internal/templates/v1/badge_no_user_found.qtpl:41
func BadgeNoUserFound() string {
//line internal/templates/v1/badge_no_user_found.qtpl:41
	qb422016 := qt422016.AcquireByteBuffer()
//line internal/templates/v1/badge_no_user_found.qtpl:41
	WriteBadgeNoUserFound(qb422016)
//line internal/templates/v1/badge_no_user_found.qtpl:41
	qs422016 := string(qb422016.B)
//line internal/templates/v1/badge_no_user_found.qtpl:41
	qt422016.ReleaseByteBuffer(qb422016)
//line internal/templates/v1/badge_no_user_found.qtpl:41
	return qs422016
//line internal/templates/v1/badge_no_user_found.qtpl:41
}

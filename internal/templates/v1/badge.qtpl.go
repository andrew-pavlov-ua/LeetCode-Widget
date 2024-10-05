// Code generated by qtc from "badge.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line internal/templates/v1/badge.qtpl:1
package v1

//line internal/templates/v1/badge.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line internal/templates/v1/badge.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line internal/templates/v1/badge.qtpl:1
func StreamBadge(qw422016 *qt422016.Writer, stats LcUserData, barsWidth BarsWidth) {
//line internal/templates/v1/badge.qtpl:1
	qw422016.N().S(`
<a href="https://github.com/andrew-pavlov-ua/LeetCode-Widget" target="_blank">
  <svg width="500" height="200" viewBox="0 0 500 200" fill="none" xmlns="http://www.w3.org/2000/svg">
    <!-- Background -->
    <rect width="500" height="200" rx="12" fill="#1E1E1E"></rect>

    <!-- LeetCode Logo -->
    <g>
      <image src="/LeetCodeLogo.png" x="20" y="10" width="30" height="30" />
    </g>

    <!-- User Info -->
    <g fill="#FFF" font-family="Verdana, sans-serif">
      <!-- Username next to the logo -->
      <text x="60" y="35" font-size="24" font-weight="bold">`)
//line internal/templates/v1/badge.qtpl:15
	qw422016.E().S(stats.UserSlug)
//line internal/templates/v1/badge.qtpl:15
	qw422016.N().S(`</text>

      <!-- Real Name under the Username -->
      <text x="60" y="60" font-size="18" fill="#888">`)
//line internal/templates/v1/badge.qtpl:18
	qw422016.E().S(stats.Username)
//line internal/templates/v1/badge.qtpl:18
	qw422016.N().S(`</text>

      <!-- Rank on the top right corner -->
      <text x="450" y="35" font-size="18" fill="#888" text-anchor="end">#`)
//line internal/templates/v1/badge.qtpl:21
	qw422016.N().F(stats.Rank)
//line internal/templates/v1/badge.qtpl:21
	qw422016.N().S(`</text>
    </g>

    <!-- Circle with number and border -->
    <g>
      <circle cx="80" cy="140" r="45" fill="none" stroke="#FFA116" stroke-width="5"></circle>
      <circle cx="80" cy="140" r="40" fill="#444"></circle>
      <text x="80" y="145" font-size="28" fill="#FFF" font-family="Verdana, sans-serif" text-anchor="middle" alignment-baseline="middle" dominant-baseline="middle">`)
//line internal/templates/v1/badge.qtpl:28
	qw422016.N().DL(stats.TotalCount)
//line internal/templates/v1/badge.qtpl:28
	qw422016.N().S(`</text>
    </g>

    <!-- Progress Bars for Problem Difficulty -->
    <g font-family="Verdana, sans-serif">
      <!-- Easy -->
      <text x="150" y="120" font-size="18" fill="#FFF">Easy</text>
      <rect x="240" y="110" width="220" height="10" fill="#333" rx="5"></rect>
      <rect x="240" y="110" width="`)
//line internal/templates/v1/badge.qtpl:36
	qw422016.N().F(barsWidth.EasyWidth)
//line internal/templates/v1/badge.qtpl:36
	qw422016.N().S(`" height="10" fill="#4CAF50" rx="5"></rect>
      <text x="470" y="120" font-size="14" fill="#FFF" text-anchor="end">`)
//line internal/templates/v1/badge.qtpl:37
	qw422016.N().DL(stats.EasyCount)
//line internal/templates/v1/badge.qtpl:37
	qw422016.N().S(` / `)
//line internal/templates/v1/badge.qtpl:37
	qw422016.N().DL(EasyMaxValue)
//line internal/templates/v1/badge.qtpl:37
	qw422016.N().S(`</text>

      <!-- Medium -->
      <text x="150" y="145" font-size="18" fill="#FFF">Medium</text>
      <rect x="240" y="135" width="220" height="10" fill="#333" rx="5"></rect>
      <rect x="240" y="135" width="`)
//line internal/templates/v1/badge.qtpl:42
	qw422016.N().F(barsWidth.MediumWidth)
//line internal/templates/v1/badge.qtpl:42
	qw422016.N().S(`" height="10" fill="#FF9800" rx="5"></rect>
      <text x="470" y="145" font-size="14" fill="#FFF" text-anchor="end">`)
//line internal/templates/v1/badge.qtpl:43
	qw422016.N().DL(stats.MediumCount)
//line internal/templates/v1/badge.qtpl:43
	qw422016.N().S(` / `)
//line internal/templates/v1/badge.qtpl:43
	qw422016.N().DL(MediumMaxValue)
//line internal/templates/v1/badge.qtpl:43
	qw422016.N().S(`</text>

      <!-- Hard -->
      <text x="150" y="170" font-size="18" fill="#FFF">Hard</text>
      <rect x="240" y="160" width="220" height="10" fill="#333" rx="5"></rect>
      <rect x="240" y="160" width="`)
//line internal/templates/v1/badge.qtpl:48
	qw422016.N().F(barsWidth.HardWidth)
//line internal/templates/v1/badge.qtpl:48
	qw422016.N().S(`" height="10" fill="#F44336" rx="5"></rect>
      <text x="470" y="170" font-size="14" fill="#FFF" text-anchor="end">`)
//line internal/templates/v1/badge.qtpl:49
	qw422016.N().DL(stats.HardCount)
//line internal/templates/v1/badge.qtpl:49
	qw422016.N().S(` / `)
//line internal/templates/v1/badge.qtpl:49
	qw422016.N().DL(HardMaxValue)
//line internal/templates/v1/badge.qtpl:49
	qw422016.N().S(`</text>
    </g>
  </svg>
</a>
`)
//line internal/templates/v1/badge.qtpl:53
}

//line internal/templates/v1/badge.qtpl:53
func WriteBadge(qq422016 qtio422016.Writer, stats LcUserData, barsWidth BarsWidth) {
//line internal/templates/v1/badge.qtpl:53
	qw422016 := qt422016.AcquireWriter(qq422016)
//line internal/templates/v1/badge.qtpl:53
	StreamBadge(qw422016, stats, barsWidth)
//line internal/templates/v1/badge.qtpl:53
	qt422016.ReleaseWriter(qw422016)
//line internal/templates/v1/badge.qtpl:53
}

//line internal/templates/v1/badge.qtpl:53
func Badge(stats LcUserData, barsWidth BarsWidth) string {
//line internal/templates/v1/badge.qtpl:53
	qb422016 := qt422016.AcquireByteBuffer()
//line internal/templates/v1/badge.qtpl:53
	WriteBadge(qb422016, stats, barsWidth)
//line internal/templates/v1/badge.qtpl:53
	qs422016 := string(qb422016.B)
//line internal/templates/v1/badge.qtpl:53
	qt422016.ReleaseByteBuffer(qb422016)
//line internal/templates/v1/badge.qtpl:53
	return qs422016
//line internal/templates/v1/badge.qtpl:53
}

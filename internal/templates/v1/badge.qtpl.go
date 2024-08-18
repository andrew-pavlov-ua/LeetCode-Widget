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
func StreamBadge(qw422016 *qt422016.Writer, stats LcStats, barsWidth BarsWidth) {
//line internal/templates/v1/badge.qtpl:1
	qw422016.N().S(`
<svg width="600" height="150" viewBox="0 0 600 150" fill="none" xmlns="http://www.w3.org/2000/svg">
  <!-- Background -->
  <rect width="600" height="150" rx="12" fill="#F0F0F0"></rect>

  <!-- User Info -->
  <g fill="#444" font-family="Verdana, sans-serif">
    <text x="20" y="35" font-size="18" font-weight="bold">`)
//line internal/templates/v1/badge.qtpl:8
	qw422016.E().S(stats.Username)
//line internal/templates/v1/badge.qtpl:8
	qw422016.N().S(`</text>
    <text x="20" y="60" font-size="14">Rank: `)
//line internal/templates/v1/badge.qtpl:9
	qw422016.N().DL(stats.Rank)
//line internal/templates/v1/badge.qtpl:9
	qw422016.N().S(`</text>
    <text x="20" y="80" font-size="14">Experience: `)
//line internal/templates/v1/badge.qtpl:10
	qw422016.N().DL(stats.Lvl)
//line internal/templates/v1/badge.qtpl:10
	qw422016.N().S(`</text>
    <text x="20" y="100" font-size="14">Total Problems: `)
//line internal/templates/v1/badge.qtpl:11
	qw422016.N().DL(stats.TotalCount)
//line internal/templates/v1/badge.qtpl:11
	qw422016.N().S(`</text>
  </g>

  <!-- Progress Bars for Problem Difficulty -->
  <g font-family="Verdana, sans-serif">
    <!-- Easy -->
    <text x="250" y="35" font-size="14" fill="#444">Easy</text>
    <rect x="310" y="25" width="`)
//line internal/templates/v1/badge.qtpl:18
	qw422016.N().DL(BarWidthValue)
//line internal/templates/v1/badge.qtpl:18
	qw422016.N().S(`" height="12" fill="#A9A9A9" rx="6"></rect>
    <rect x="310" y="25" width="`)
//line internal/templates/v1/badge.qtpl:19
	qw422016.N().F(barsWidth.EasyWidth)
//line internal/templates/v1/badge.qtpl:19
	qw422016.N().S(`" height="12" fill="#4CAF50" rx="6"></rect>
    <text x="400" y="35" font-size="12" fill="#444">`)
//line internal/templates/v1/badge.qtpl:20
	qw422016.N().DL(stats.EasyCount)
//line internal/templates/v1/badge.qtpl:20
	qw422016.N().S(` / `)
//line internal/templates/v1/badge.qtpl:20
	qw422016.N().DL(EasyMaxValue)
//line internal/templates/v1/badge.qtpl:20
	qw422016.N().S(`</text>

    <!-- Medium -->
    <text x="250" y="70" font-size="14" fill="#444">Medium</text>
    <rect x="310" y="60" width="`)
//line internal/templates/v1/badge.qtpl:24
	qw422016.N().DL(BarWidthValue)
//line internal/templates/v1/badge.qtpl:24
	qw422016.N().S(`" height="12" fill="#A9A9A9" rx="6"></rect>
    <rect x="310" y="60" width="`)
//line internal/templates/v1/badge.qtpl:25
	qw422016.N().F(barsWidth.MediumWidth)
//line internal/templates/v1/badge.qtpl:25
	qw422016.N().S(`" height="12" fill="#FF9800" rx="6"></rect>
    <text x="400" y="70" font-size="12" fill="#444">`)
//line internal/templates/v1/badge.qtpl:26
	qw422016.N().DL(stats.MediumCount)
//line internal/templates/v1/badge.qtpl:26
	qw422016.N().S(` / `)
//line internal/templates/v1/badge.qtpl:26
	qw422016.N().DL(MediumMaxValue)
//line internal/templates/v1/badge.qtpl:26
	qw422016.N().S(`</text>

    <!-- Hard -->
    <text x="250" y="105" font-size="14" fill="#444">Hard</text>
    <rect x="310" y="95" width="`)
//line internal/templates/v1/badge.qtpl:30
	qw422016.N().DL(BarWidthValue)
//line internal/templates/v1/badge.qtpl:30
	qw422016.N().S(`" height="12" fill="#A9A9A9" rx="6"></rect>
    <rect x="310" y="95" width="`)
//line internal/templates/v1/badge.qtpl:31
	qw422016.N().F(barsWidth.HardWidth)
//line internal/templates/v1/badge.qtpl:31
	qw422016.N().S(`" height="12" fill="#F44336" rx="6"></rect>
    <text x="400" y="105" font-size="12" fill="#444">`)
//line internal/templates/v1/badge.qtpl:32
	qw422016.N().DL(stats.HardCount)
//line internal/templates/v1/badge.qtpl:32
	qw422016.N().S(` / `)
//line internal/templates/v1/badge.qtpl:32
	qw422016.N().DL(HardMaxValue)
//line internal/templates/v1/badge.qtpl:32
	qw422016.N().S(`</text>
  </g>
</svg>
`)
//line internal/templates/v1/badge.qtpl:35
}

//line internal/templates/v1/badge.qtpl:35
func WriteBadge(qq422016 qtio422016.Writer, stats LcStats, barsWidth BarsWidth) {
//line internal/templates/v1/badge.qtpl:35
	qw422016 := qt422016.AcquireWriter(qq422016)
//line internal/templates/v1/badge.qtpl:35
	StreamBadge(qw422016, stats, barsWidth)
//line internal/templates/v1/badge.qtpl:35
	qt422016.ReleaseWriter(qw422016)
//line internal/templates/v1/badge.qtpl:35
}

//line internal/templates/v1/badge.qtpl:35
func Badge(stats LcStats, barsWidth BarsWidth) string {
//line internal/templates/v1/badge.qtpl:35
	qb422016 := qt422016.AcquireByteBuffer()
//line internal/templates/v1/badge.qtpl:35
	WriteBadge(qb422016, stats, barsWidth)
//line internal/templates/v1/badge.qtpl:35
	qs422016 := string(qb422016.B)
//line internal/templates/v1/badge.qtpl:35
	qt422016.ReleaseByteBuffer(qb422016)
//line internal/templates/v1/badge.qtpl:35
	return qs422016
//line internal/templates/v1/badge.qtpl:35
}

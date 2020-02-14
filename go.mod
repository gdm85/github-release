module github.com/gdm85/github-release

go 1.13

require (
	github.com/dustin/go-humanize v0.0.0-20171012181109-77ed807830b4
	github.com/tomnomnom/linkheader v0.0.0-20170505194411-6c03f819bd09
	github.com/voxelbrain/goptions v0.0.0-20151102231003-26cb8b046923
)

replace github.com/voxelbrain/goptions => ./vendor/github.com/gdm85/goptions

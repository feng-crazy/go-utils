package regex

import (
	"regexp"
)

var (
	NumericRegexp              = New(`^(\d+)$`)
	AlphaNumericRegexp         = New("^([0-9A-Za-z]+)$")
	AlphaRegexp                = New("^([A-Za-z]+)$")
	AlphaCapsOnlyRegexp        = New("^([A-Z]+)$")
	AlphaNumericCapsOnlyRegexp = New("^([0-9A-Z]+)$")
	UrlRegexp                  = New(`^((http?|https?|ftps?):\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`)
	EmailRegexp                = New(`^(.+@([\da-z\.-]+)\.([a-z\.]{2,6}))$`)
	HashtagHexRegexp           = New(`^#([a-f0-9]{6}|[a-f0-9]{3})$`)
	ZeroXHexRegexp             = New(`^0x([a-f0-9]+|[A-F0-9]+)$`)
	IPv4Regexp                 = New(`^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	IPv6Regexp                 = New(`^([0-9A-Fa-f]{0,4}:){2,7}([0-9A-Fa-f]{1,4}$|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})$`)
)

// Represents the General pattern type
type Pattern struct {
	pat string
	reg *regexp.Regexp
}

func New(pat string) *Pattern {
	p := &Pattern{pat: pat}
	p.reg = regexp.MustCompile(p.pat)
	return p
}

// Returns true if s matches the compiled regex pattern
func (p *Pattern) Match(s string) bool {
	return p.reg.MatchString(s)
}

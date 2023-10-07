package Structure

import (
	"Silver/Silver/Utils"
	"regexp"
)

type Controller func(request Utils.Request, response Utils.Response)
type RequestMethod string

type Route struct {
	Method     RequestMethod
	Path       *regexp.Regexp
	FilePath   string
	Controller Controller
	Static     bool
}

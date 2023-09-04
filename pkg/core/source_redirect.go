package core

import (
	"github.com/bluenviron/mediamtx/pkg/logger"
)

// sourceRedirect is a source that redirects to another one.
type sourceRedirect struct{}

func (*sourceRedirect) Log(logger.Level, string, ...interface{}) {
}

// apiSourceDescribe implements source.
func (*sourceRedirect) apiSourceDescribe() pathAPISourceOrReader {
	return pathAPISourceOrReader{
		Type: "redirect",
		ID:   "",
	}
}

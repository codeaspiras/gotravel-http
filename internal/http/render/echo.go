package render

import (
	"github.com/codeaspiras/gotravel/stdio"
)

func echo(io stdio.IO, template string, args ...interface{}) {
	if io == nil {
		io = stdio.New()
	}

	io.Echo(template, args...)
}

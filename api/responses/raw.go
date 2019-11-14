package responses

import (
	"fmt"
	"strings"

	"github.com/valyala/fasthttp"
)

const (
	separator = ` `
)

// Strings reply from []strings
func Strings(ctx *fasthttp.RequestCtx, s []string) (answered bool) {

	var response string
	response = strings.Join(s, separator)

	if response != "" {
		answered = true
	}

	fmt.Fprintf(ctx, "%s", response)
	return
}

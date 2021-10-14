package api

import (
	"bytes"
	"testing"
)

const expect = `<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="go.beyondstorage.io git https://github.com/beyondstorage/go-storage">
<meta name="go-source" content="go.beyondstorage.io https://github.com/beyondstorage/go-storage https://github.com/beyondstorage/go-storage/tree/master{/dir} https://github.com/beyondstorage/go-storage/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://pkg.go.dev/go.beyondstorage.io/services/s3">
</head>
<body>
Nothing to see here; <a href="https://pkg.go.dev/go.beyondstorage.io/services/s3">see the package on pkg.go.dev</a>.
</body>
</html>`

func TestHandle(t *testing.T) {
	var w bytes.Buffer
	err := handle(&w,
		"services/s3",
		"",
		"go-storage")
	if err != nil {
		t.Errorf("Handle content: %v", err)
		return
	}

	if w.String() != expect {
		t.Errorf("content is not match, expect:\n%s\nacutal:\n%s", expect, w.String())
	}
}

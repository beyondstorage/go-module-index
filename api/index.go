package api

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/")

	var err error
	defer func() {
		if err != nil {
			log.Printf("handle path %s: %v", path, err)
			w.WriteHeader(500)
		} else {
			w.WriteHeader(200)
		}
	}()

	// Handle go-storage and it's submodules
	err = handle(w, path, "", "go-storage")

	// TODO: we will need to handle repo like beyond-ctl and so on.
	return
}

type module struct {
	Name string // The full module name.
	Root string
	Repo string // The corresponding repo name.
}

var moduleTmpl = template.Must(template.New("module").Parse(`<!DOCTYPE html>
<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8"/>
<meta name="go-import" content="go.beyondstorage.io{{ .Root }} git https://github.com/beyondstorage/{{ .Repo }}">
<meta name="go-source" content="go.beyondstorage.io{{ .Root }} https://github.com/beyondstorage/{{ .Repo }} https://github.com/beyondstorage/{{ .Repo }}/tree/master{/dir} https://github.com/beyondstorage/{{ .Repo }}/blob/master{/dir}/{file}#L{line}">
<meta http-equiv="refresh" content="0; url=https://pkg.go.dev/go.beyondstorage.io{{ .Name }}">
</head>
<body>
Nothing to see here; <a href="https://pkg.go.dev/go.beyondstorage.io{{ .Name }}">see the package on pkg.go.dev</a>.
</body>
</html>`))

func handle(w io.Writer, name, root, repo string) (err error) {
	m := module{
		Name: name,
		Root: root,
		Repo: repo,
	}

	// We will always add "go.beyondstorage.io" as prefix of our module name.
	//
	// If the real module name is "go.beyondstorage.io", user will input "", so
	// we don't need to do anything.
	// If the real module name is "go.beyondstorage.io/services/s3", user will
	// input "services/s3", so let's add a "/" before the name.
	if m.Name != "" {
		m.Name = "/" + m.Name
	}

	err = moduleTmpl.Execute(w, m)
	if err != nil {
		return fmt.Errorf("module template generate: %v", err)
	}
	return nil
}

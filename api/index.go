package api

import (
	"log"
	"net/http"
	"strings"

	. "go.beyondstorage.io/module-index"
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

	// Handle all service packages
	if strings.HasPrefix(path, "services/") {
		err = Handle(w, path, "", "go-storage")
		return
	}

	// Handle all pkg packages
	if strings.HasPrefix(path, "pkg/") {
		err = Handle(w, path, "", "go-storage")
		return
	}

	// TODO: we will need to handle repo like beyond-ctl and so on.

	// go.beyondstorage.io => go-storage
	err = Handle(w, "", "", "go-storage")
	return
}

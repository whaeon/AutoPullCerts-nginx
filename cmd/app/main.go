package main

import (
	"AutoPullCerts/internal/app"
	"flag"
)

func main() {
	domain := flag.String("D", "", "Domain")
	download := flag.Bool("download", false, "download certificates contained in nginx.conf")
	upload := flag.Bool("upload", false, "upload certs to dnspod")
	list := flag.Bool("list", false, "list certificates")
	display := flag.Bool("d", false, "list sub options, display certificate information (default false)")
	offset := flag.Uint64("o", 0, "list sub options, paging offset, starting at 0")
	limit := flag.Uint64("l", 20, "list sub options, limit records per page")
	Upload := flag.Int64("U", 1, "list sub options, filter managed certificate, 1 means filter, 0 means not filter")

	flag.Parse()

	cert := app.NewRequestParams(*display, *domain, *offset, *limit, *Upload)

	if *list {
		cert.List()
	} else if *upload {
		cert.Upload()
	} else if *download {
		cert.Download()
	}
	// app.File()
}

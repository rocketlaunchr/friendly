// Copyright 2020 PJ Engineering and Business Solutions Pty. Ltd. All rights reserved.

package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/pkg/browser"
	"github.com/rocketlaunchr/https-go"
	"github.com/spf13/cobra"
)

var noCacheHeaders = map[string]string{
	"Expires":         time.Unix(0, 0).Format(time.RFC1123),
	"Cache-Control":   "no-cache, no-store, must-revalidate, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}

var lastReqTime *time.Time

type wrapHandler struct {
	fs    http.Handler
	quiet bool
}

func (h *wrapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !h.quiet {
		start := time.Now()
		defer func() {
			end := time.Now()
			d := end.Sub(start)

			if lastReqTime != nil && time.Since(*lastReqTime) > 1*time.Second {
				fmt.Println("==========================")
			}

			lastReqTime = &start

			magenta1 := color.New(color.FgWhite, color.BgMagenta, color.Underline).SprintfFunc()
			magenta2 := color.New(color.FgWhite, color.BgCyan).SprintfFunc()

			html := color.New(color.Bold, color.FgBlue).SprintFunc()
			css := color.New(color.Bold, color.FgGreen).SprintFunc()
			js := color.New(color.Bold, color.FgRed).SprintFunc()
			img := color.New(color.Bold, color.FgYellow).SprintFunc()
			other := color.New(color.Bold, color.FgBlack).SprintFunc()

			switch filepath.Ext(r.URL.Path) {
			case ".html":
				fmt.Printf("[%s#%s] %s %s\n", magenta1("%s", start.Local().Format("15:04:05.000")), magenta2("%s", d.String()), strings.ToUpper(r.Method), html(r.URL.Path))
			case ".js":
				fmt.Printf("[%s#%s] %s %s\n", magenta1("%s", start.Local().Format("15:04:05.000")), magenta2("%s", d.String()), strings.ToUpper(r.Method), js(r.URL.Path))
			case ".css":
				fmt.Printf("[%s#%s] %s %s\n", magenta1("%s", start.Local().Format("15:04:05.000")), magenta2("%s", d.String()), strings.ToUpper(r.Method), css(r.URL.Path))
			case ".png", ".jpg", ".jpeg", ".ico", ".svg":
				fmt.Printf("[%s#%s] %s %s\n", magenta1("%s", start.Local().Format("15:04:05.000")), magenta2("%s", d.String()), strings.ToUpper(r.Method), img(r.URL.Path))
			default:
				fmt.Printf("[%s#%s] %s %s\n", magenta1("%s", start.Local().Format("15:04:05.000")), magenta2("%s", d.String()), strings.ToUpper(r.Method), other(r.URL.Path))
			}
		}()
	}

	// Prevent caching
	for k, v := range noCacheHeaders {
		w.Header().Set(k, v)
	}

	h.fs.ServeHTTP(w, r)
}

func runCmd(cmd *cobra.Command, args []string) {

	quietMode, _ := cmd.Flags().GetBool("quiet")

	port, _ := cmd.Flags().GetString("port")
	path, _ := cmd.Flags().GetString("path")
	httpsOn, _ := cmd.Flags().GetBool("https")
	save, _ := cmd.Flags().GetBool("save")
	browserOpen, _ := cmd.Flags().GetBool("browser")
	removeCerts, _ := cmd.Flags().GetBool("remove")

	if port == "" {
		if httpsOn {
			port = "4430"
		} else {
			port = "8080"
		}
	}

	http.Handle("/", &wrapHandler{http.FileServer(http.Dir(path)), quietMode})

	var homeURL string
	if httpsOn {
		homeURL = "https://localhost:" + port
	} else {
		homeURL = "http://localhost:" + port
	}

	if !quietMode {
		if httpsOn {
			fmt.Printf("running server on https://localhost:%s\n", port)
		} else {
			fmt.Printf("running server on http://localhost:%s\n", port)
		}
	}

	if browserOpen {
		go func() {
			time.Sleep(1250 * time.Millisecond)
			browser.OpenURL(homeURL)
		}()
	}

	if removeCerts {
		err := deleteCerts()
		if err != nil {
			if httpsOn && !strings.Contains(err.Error(), "no such file or directory") {
				log.Fatal(err)
			}
		}
	}

	if !httpsOn {
		log.Fatal(http.ListenAndServe(":"+port, nil))
	} else {

		pub, priv, err := getCerts(!save)
		if err != nil {
			pub, priv, err = https.GenerateKeys(https.GenerateOptions{Host: "localhost"})
			if err != nil {
				// could not generate keys
				log.Fatal(err)
			}
			if save {
				saveCerts(pub, priv)
			}
		}

		cert, err := tls.X509KeyPair(pub, priv)
		if err != nil {
			log.Fatal(err)
		}

		httpServer := &http.Server{
			Addr:      ":" + port,
			TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}},
			ErrorLog:  log.New(&ThrowAway{}, "", 0),
		}

		log.Fatal(httpServer.ListenAndServeTLS("", ""))
	}

}

package server

import (
	"fmt"
	"log"

	"cmd/internal/env"
	"crypto/tls"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

func Run(handler http.Handler) {
	environment := env.Must("ENVIRONMENT")

	switch environment {
	case "prod":
		runProduction(handler)
	case "dev":
		runDevelopment(handler)
	default:
		panic(fmt.Sprintf("unknown environment %q", environment))
	}
}

func runProduction(handler http.Handler) {
	certManager := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist(env.Must("HOST")),
		Cache:      autocert.DirCache(env.Must("TLS_CERTIFICATES_DIR")),
	}

	server := &http.Server{
		Addr:    ":https",
		Handler: handler,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			MinVersion:     tls.VersionTLS12, // improves cert reputation score at https://www.ssllabs.com/ssltest/
		},
	}

	var g errgroup.Group

	g.Go(func() error {
		return http.ListenAndServe(":http", certManager.HTTPHandler(nil))
	})

	g.Go(func() error {
		return server.ListenAndServeTLS("", "") // Key and cert are coming from Let's Encrypt
	})

	log.Fatal(g.Wait())
}

func runDevelopment(handler http.Handler) {
	fmt.Println("Started dev")
	log.Fatal(http.ListenAndServe(":http", handler))
}

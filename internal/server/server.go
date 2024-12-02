package server

import (
	"fmt"
	"log"
	"strings"
	"time"

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
		HostPolicy: autocert.HostWhitelist(strings.Split(env.Must("HOSTS"), ",")...),
		Cache:      autocert.DirCache(env.Must("TLS_CERTIFICATES_DIR")),
	}

	server := &http.Server{
		Addr:    ":https",
		Handler: handler,
		TLSConfig: &tls.Config{
			GetCertificate: certManager.GetCertificate,
			MinVersion:     tls.VersionTLS12, // improves cert reputation score at https://www.ssllabs.com/ssltest/
		},
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	var g errgroup.Group

	g.Go(func() error {
		return http.ListenAndServe(":http", certManager.HTTPHandler(nil))
	})

	g.Go(func() error {
		return server.ListenAndServeTLS("", "") // Key and cert are coming from Let's Encrypt
	})

	fmt.Print(g.Wait())
}

func runDevelopment(handler http.Handler) {
	fmt.Println("Started dev")
	server := &http.Server{
		Addr:         ":http",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("Started dev")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("runDevelopment: ", err)
	}

}

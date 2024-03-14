package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"

	"github.com/newrelic/go-agent/v3/newrelic"
)

var (
	nrAPMAppName string
	nrLicenseKey string
)

func init() {
	envVars := map[string]*string{
		"NR_APM_APP_NAME": &nrAPMAppName,
		"NR_LICENSE_KEY":  &nrLicenseKey,
	}

	for key, val := range envVars {
		if v, ok := os.LookupEnv(key); ok {
			*val = v
		} else {
			log.Fatalf("%s environment variables is not set", key)
		}
	}
}

func main() {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(nrAPMAppName),
		newrelic.ConfigLicense(nrLicenseKey),
	)
	if err != nil {
		log.Fatalf("fail to initialize newrelic APM: %v", err)
	}
	fmt.Println("initialize newrelic APM successfuly")

	http.HandleFunc(newrelic.WrapHandleFunc(app, "/foo", func(w http.ResponseWriter, r *http.Request) {
		rand.NewSource(time.Now().UnixNano())
		n := rand.Intn(200)
		fmt.Printf("Sleeping %d micro seconds...\n", n)
		time.Sleep(time.Duration(n) * time.Microsecond)
		fmt.Println("Done")

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

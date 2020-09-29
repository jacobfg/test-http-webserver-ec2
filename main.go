package main

import (
    "crypto/tls"
    "log"
    "fmt"
    "net/http"
    "io/ioutil"
    "flag"
)

const (
    addr = "http://169.254.169.254/latest/meta-data/"
)

var (
    metadata = map[string]string{
        "instance-id":       "instance-id",
        "local-ipv4":        "local-ipv4",
        "public-ipv4":       "public-ipv4",
        "availability-zone": "placement/availability-zone",
        "instance-type":     "instance-type",
        "ami-id":            "ami-id",
    }
    certfile string
)

// type T struct {
//     A string
//     B struct {
//             RenamedC int   `yaml:"c"`
//             D        []int `yaml:",flow"`
//     }
// }

func serveHTTP(mux *http.ServeMux, errs chan<- error) {
    errs <- http.ListenAndServe(":8080", mux)
}
  
func serveHTTPS(mux *http.ServeMux, certfile string, errs chan<- error) {
    cfg := &tls.Config{
        MinVersion:               tls.VersionTLS12,
        CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
        PreferServerCipherSuites: true,
        CipherSuites: []uint16{
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
            tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_RSA_WITH_AES_256_CBC_SHA,
        },
    }
    srv := &http.Server{
        Addr:         ":8443",
        Handler:      mux,
        TLSConfig:    cfg,
        TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
    }
    errs <- srv.ListenAndServeTLS(certfile, certfile)
}

func get(part string) (string, error) {
	resp, err := http.Get(addr + part)
	if err != nil {
		return "", err
    }
    if resp.StatusCode == 200 {
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return "", err
        }
        return string(body), nil
    } else {
        return "null", nil
    }
}

func AWSInfo(w http.ResponseWriter, r *http.Request) {
    for k, v := range metadata {
        value, err := get(v)
        if err == nil {
            fmt.Fprintf(w, "%s: %s\n", k, value)
        }
    }
}

func main() {
    log.SetFlags(0)

    flag.StringVar(&certfile, "c", "localhost.crt", "a string var")
    flag.Parse()
    fmt.Println("certfile:", certfile)

    mux := http.NewServeMux()
    // setup routes for mux     // define your endpoints
    mux.HandleFunc("/", AWSInfo)

    errs := make(chan error, 1)         // a channel for errors
    go serveHTTP(mux, errs)             // start the http server in a thread
    go serveHTTPS(mux, certfile, errs) // start the https server in a thread
    log.Fatal(<-errs)                   // block until one of the servers writes an error
}

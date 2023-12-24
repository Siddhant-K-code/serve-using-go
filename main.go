package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

func main() {
	// Simplify flag setup using a struct.
	cfg := struct {
		Port      int
		Public    bool
		CorsAllow string
		Open      bool
	}{
		Port: 8000,
	}

	flag.IntVar(&cfg.Port, "port", cfg.Port, "the port of the HTTP file server")
	flag.BoolVar(&cfg.Public, "public", cfg.Public, "listen on all interfaces")
	flag.StringVar(&cfg.CorsAllow, "cors-allow", cfg.CorsAllow, "origins to permit via CORS")
	flag.BoolVar(&cfg.Open, "open", cfg.Open, "open in web browser")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [opts] [directory]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	host := "127.0.0.1"
	if cfg.Public {
		host = "0.0.0.0"
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: %v", err)
	}

	done := make(chan struct{})
	go func() {
		if err := startServer(host, cfg.Port, cfg.CorsAllow, wd); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
		close(done)
	}()

	if cfg.Open {
		time.Sleep(100 * time.Millisecond)
		openBrowser(cfg.Port)
	}

	<-done
}

func openBrowser(port int) {
	url := fmt.Sprintf("http://127.0.0.1:%d", port)
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Printf("Failed to open browser: %v", err)
	}
}

type fileServer struct {
	root      string
	corsAllow string
}

func (s *fileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL.String())

	path := filepath.Join(s.root, r.URL.Path)
	if s.corsAllow != "" {
		w.Header().Add("Access-Control-Allow-Origin", s.corsAllow)
	}
	http.ServeFile(w, r, path)
}

func startServer(host string, port int, corsAllow, wd string) error {
	for {
		bindAddr := fmt.Sprintf("%s:%d", host, port)
		listener, err := net.Listen("tcp", bindAddr)
		if err != nil {
			if bindConflict(err) {
				log.Printf("Could not bind to %s, trying next port", bindAddr)
				port++
				continue
			}
			return err
		}

		log.Printf("🚀  Listening on http://%s/", bindAddr)
		srv := &fileServer{root: wd, corsAllow: corsAllow}
		return http.Serve(listener, srv)
	}
}

func bindConflict(err error) bool {
	// Use errors.As for error type assertion.
	var opErr *net.OpError
	if errors.As(err, &opErr) {
		var syscallErr *os.SyscallError
		if errors.As(opErr.Err, &syscallErr) {
			return syscallErr.Err.Error() == "address already in use"
		}
	}
	return false
}

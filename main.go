package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"passport-date/commands"
	c "passport-date/cron"
	rprof "runtime/pprof"
	"strings"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
)

func pprofs(w http.ResponseWriter, r *http.Request) {
	switch strings.TrimPrefix(r.URL.Path, "/") {
	case "", "pprof":
		pprof.Index(w, r)
	case "trace":
		pprof.Trace(w, r)
	case "profile":
		pprof.Profile(w, r)
	case "cmdline":
		pprof.Cmdline(w, r)
	case "symbol":
		pprof.Symbol(w, r)
	default:
		name := strings.ToLower(strings.TrimPrefix(r.URL.Path, "/"))
		if rprof.Lookup(name) == nil {
			pprof.Index(w, r)
			return
		}
		fmt.Println("name:", name)
		pprof.Handler(name).ServeHTTP(w, r)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cron := cron.New()
	cron.AddFunc("@every 1s", func() {
		fmt.Println("cron job")
		is, err := c.IsDateAvailable()
		if err != nil {
			fmt.Println("err checking date", err)
		}
		if is {
			// Mailer()
		}
	})
	cron.Start()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", r.URL.Path)
	})
	http.HandleFunc("/debug", pprofs)
	http.ListenAndServe(":8080", nil)
}

func MainCommand() {
	var rootCmd = &cobra.Command{Use: "passport-date"}
	rootCmd.AddCommand(commands.DateCmd)
	rootCmd.Execute()
}

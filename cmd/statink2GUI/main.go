package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/cad-san/statink2"
)

type config struct {
	APIKey string
}

type options struct {
	configPath string
}

type appHandler struct {
	client *statink2.Client
}

func parseFlag() options {
	var opt options
	flag.StringVar(&opt.configPath, "config", "config.json", "file path")
	flag.StringVar(&opt.configPath, "c", "config.json", "file path")
	flag.Parse()
	return opt
}

func (a appHandler) battleForm(w http.ResponseWriter, r *http.Request) {
}

func (a appHandler) postBattle(w http.ResponseWriter, r *http.Request) {
}

func (a appHandler) index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		a.battleForm(w, r)
	} else if r.Method == "POST" {
		a.postBattle(w, r)
	}
}

func main() {
	opt := parseFlag()
	if len(opt.configPath) == 0 {
		opt.configPath = "config.json"
	}
	file, err := os.Open(opt.configPath)
	if err != nil {
		fmt.Printf("cannot read config file:%v", err)
		os.Exit(1)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	var conf config
	if err = decoder.Decode(&conf); err != nil {
		fmt.Printf("connt read config file:%v", err)
		os.Exit(1)
	}

	logger := log.New(os.Stderr, "", log.LstdFlags|log.Llongfile)
	client, err := statink2.NewClient(conf.APIKey, logger)
	if err != nil {
		os.Exit(1)
	}
	app := appHandler{
		client: client,
	}
	http.HandleFunc("/", app.index)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("server failed. err:%v", err)
		os.Exit(1)
	}
}

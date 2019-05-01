package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Moekr/gopkg/logs"
	"github.com/Moekr/lightning/article"
	"github.com/Moekr/lightning/http"
)

func main() {
	var dataPath, logsPath, address string
	flag.StringVar(&dataPath, "data", "post", "Path of article files, default post")
	flag.StringVar(&logsPath, "logs", "", "Path of logs file, default stdout")
	flag.StringVar(&address, "bind", "127.0.0.1:8080", "Bind address, default 127.0.0.1:8080")
	flag.Parse()
	logs.InitLogs(logsPath)
	article.LoadArticles(dataPath)
	http.StartHTTPService(address)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	for s := range ch {
		if s == syscall.SIGHUP {
			logs.Info("[Lightning] received signal %v, reload data...", s)
			article.LoadArticles(dataPath)
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "Received signal %v, exit...\n", s)
			break
		}
	}
}

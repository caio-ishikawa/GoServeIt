package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)
const RED = "\033[31m"
const RESET = "\033[0m"

const LOGO = `
   ______     _____                      ______ 
  / ____/___ / ___/___  ______   _____  /  _/ /_
 / / __/ __ \\__ \/ _ \/ ___/ | / / _ \ / // __/
/ /_/ / /_/ /__/ /  __/ /   | |/ /  __// // /_  
\____/\____/____/\___/_/    |___/\___/___/\__/  v1.0.0
`

// Server configuration representing the data in the command line flags
type Config struct {
	filePath string
	portNum  string
	verbose  bool
}

// Wrapper around logger for custom formatting
type ReqLogger struct{}

func main() {
	config := newConfig()
	logger := initLogger()

	file, err := os.Open(config.filePath)
	if err != nil {
		fmt.Printf(
			"%sFailed to open file. Does it exist?%s\n",
			RED,
			RESET,
		)
		return
	}

	displayStartScreen(config.filePath, config.portNum)

	http.HandleFunc("/pwn", func(w http.ResponseWriter, r *http.Request) {
		err := logRequest(logger, r, config.verbose)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "text/plain")


		fileContent, err := io.ReadAll(file)
		if err != nil {
			logger.Printf("%s%s%s", RED, err, RESET)
		}

		_, err = w.Write(fileContent)
		if err != nil {
			logger.Printf("%s%s%s", RED, err, RESET)
		}

		logger.Printf("File transfer complete")
	})

	portNum := ":" + config.portNum
	if err := http.ListenAndServe(portNum, nil); err != nil {
		logger.Printf("%s%s%s", RED, err, RESET)
	}
}

// Returns a new Config object based on the command line flags
func newConfig() Config {
	filePtr := flag.String("f", "", "Path of file to serve")
	portPtr := flag.String("p", "8080", "Port")
	verbosePtr := flag.Bool("v", false, "If set, it displays the request header")
	flag.Parse()

	return Config{
		filePath: *filePtr,
		portNum:  *portPtr,
		verbose:  *verbosePtr,
	}
}

// Returns custom RequestLogger
func initLogger() *ReqLogger {
	log.SetOutput(os.Stdout)
	log.SetFlags(0)
	log.SetPrefix("")

	logger := &ReqLogger{}

	return logger
}

// Overrides log.Printf to use custom formatting.
func (rl *ReqLogger) Printf(str string, v ...interface{}) {
	timestamp := time.Now().Format(time.RFC3339)

	blue := "\033[1;34m"
	reset := "\033[0m"
	formattedLogEntry := fmt.Sprintf(
		"%s[%s]%s %s\n",
		blue,
		timestamp,
		reset,
		fmt.Sprintf(str, v...),
	)

	err := log.Output(2, formattedLogEntry)
	if err != nil {
		panic(err)
	}
}

// Prints ascii title and starting message
func displayStartScreen(filePath string, portNum string) {
	fmt.Printf("%s%s%s\n", RED, LOGO, RESET)
	fmt.Printf(
		"Serving %s%s%s on port %s%s%s\n",
		RED,
		filePath,
		RESET,
		RED,
		portNum,
		RESET,
	)
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
}

// Logs request. TODO: add debug mode in Config, and print reqDump depending on that
func logRequest(logger *ReqLogger, r *http.Request, verbose bool) error {
	logger.Printf("%s %s", r.Method, r.URL)

	if verbose {
		reqDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			return err
		}

		fmt.Printf("\n%s", string(reqDump))
	}

	return nil
}

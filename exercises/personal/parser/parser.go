/*
# Parse logfile

Mar-18-09:33:48 local_host_name amagent 42: result async=false sync failed
Mar-18-09:33:48 local_host_name Apache_tomcat 52: checkin uptime=391638 ver=1.0-20 UTC-0700 commit failed
Mar-18-09:33:48 local_host_name spark 40: Executing command: [Error, host nonexistent.
Mar-18-09:33:48 local_host_name amagent 42: fetch failed Bad URL
Mar-18-09:33:48 local_host_name Apache_tomcat 52: Error 404
Mar-18-09:33:48 local_host_name spark 40: Received execCmd command bad connection
Mar-18-09:33:48 local_host_name spark 40: Result: error in reading memory
Mar-18-09:33:49 local_host_name spark 40: Received execCmd command bad connection
Mar-18-09:33:49 local_host_name spark 40: Result: error in reading memory

# Desired output sorted by date

Mar-18-09:33:48 amaagent: 2
Mar-18-09:33:48 Apache_tomcat: 2
Mar-18-09:33:48 spark: 3

	:: TODO ::
	Parametrize log fields to parse
*/
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type logApp struct {
	Time      string `json:"time"`
	App       string `json:"application"`
	Incidents int    `json:"incidents"`
}

func (l *logApp) incrementCounter() {
	l.Incidents++
}

func main() {

	if len(os.Args) < 4 {
		log.Fatalln("Usage: parser [/path/to/logfile string] [separator string] [json/text]")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalln("error closing file")
		}
	}()

	switch os.Args[3] {
	case "text":
		fmt.Println(generateLog(scannerFile(f, os.Args[2])))
	case "json":
		enc := json.NewEncoder(os.Stdout)
		enc.Encode(scannerFile(f, os.Args[2]))
	default:
		fmt.Println(fmt.Errorf("output format invalid: received %s, wanted text or json", os.Args[3]))
	}
}

// ScannerFile scans the log file given a separator
// and populates the struct slice.
func scannerFile(f *os.File, sep string) []logApp {
	s := bufio.NewScanner(f)
	var lApp []logApp
	var counter int

	// Loop over every line
	for s.Scan() {
		// Divide lines by sep and create a []string
		lines := strings.Split(s.Text(), sep)

		// if the application already exist int the struct slice
		// just increment its incidence counter.
		// Otherwise add it as a new app incident.
		if len(lines) > 2 {
			if logExist(lines, lApp) {
				increment(lines[2], lApp)
				counter++
			} else {
				lApp = append(lApp, logApp{lines[0], lines[2], 1})
			}

		}
	}
	return lApp
}

// logExist checks if a given application has reported previously an error
func logExist(lines []string, lApp []logApp) bool {
	for k := range lApp {
		if lines[2] == lApp[k].App {
			return true
		}
	}
	return false
}

// increment increments the incident number of a given app
func increment(name string, lApp []logApp) {
	for k := range lApp {
		if lApp[k].App == name {
			lApp[k].incrementCounter()
		}
	}
}

// generateLog parses the final log
func generateLog(lApp []logApp) string {
	s := ""
	for app, line := range lApp {
		s = fmt.Sprintf("%s\n%s %s: %v", s, line.Time, line.App, lApp[app].Incidents)
	}
	return s
}

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

*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type logApp struct {
	time string
	app  string
}

var file string = "example.log"

func main() {

	// open file
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	// defer close to not leak files
	defer f.Close()

	// create data structures to hold logs and re-ocurrences
	var lApp []logApp
	m := make(map[string]int)
	var res string

	// scanner the file
	s := bufio.NewScanner(f)
	for s.Scan() {
		// divide by lines
		lines := strings.Split(s.Text(), " ")
		if len(lines) > 2 {
			lApp = append(lApp, logApp{lines[0], lines[2]})
		}
	}

	// count incidents
	for _, line := range lApp {
		m[line.app]++
	}

	// parse and print output
	exists := make(map[string]bool)
	for _, line := range lApp {
		if !exists[line.app] {
			res = fmt.Sprintf("%s\n%s %s: %v", res, line.time, line.app, m[line.app])
			exists[line.app] = true
		}
	}
	fmt.Println(res)
}

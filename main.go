package main

import (
	"flag"
	"log"
	"strconv"
	"sync"
)

// main function
func main() {

	// get command line options
	concurrency := flag.Uint("c", 5, "Concurrency; how many parallel connections to make")
	checksFile := flag.String("f", "checks.json", "Name of the checks JSON file")
	serversFile := flag.String("s", "servers.json", "Name of the servers JSON file")
	flag.Parse()

	// create checks manager and load checks file
	cm := ChecksManager{}
	cm.Load(*checksFile)

	// create servers file object and load the file
	sf := ServersFile{}
	sf.Load(*serversFile)

	log.Printf("Number of servers: %s", strconv.Itoa(len(sf.Servers)))

	// create and prepare wait sync group and semaphore
	var wg sync.WaitGroup
	wg.Add(len(sf.Servers))
	s := NewSemaphore(*concurrency)

	// go through all the servers and do the checks
	for i, server := range sf.Servers {
		s.Lock()
		go func(i int, server *Server) {
			defer wg.Done()
			defer s.Unlock()
			log.Printf("Checking server (#%s): %s", strconv.Itoa(i), server.Host)
			err := cm.ProcessChecks(server)
			if err != nil {
				log.Printf("%s", err)
			}
		}(i, server)
	}

	// wait for all goroutines to finish
	wg.Wait()
}

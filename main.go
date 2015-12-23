package main

import (
	"flag"
	"log"
	"strconv"
	"sync"
)

func main() {
	concurrency := flag.Uint("c", 5, "Concurrency; how many parallel connections to make")
	checksFile := flag.String("f", "checks.json", "Name of the checks JSON file")
	flag.Parse()

	cm := ChecksManager{}
	cm.Load(*checksFile)

	sf := ServersFile{}
	sf.Load()

	log.Printf("Number of servers: %s", strconv.Itoa(len(sf.Servers)))

	var wg sync.WaitGroup
	wg.Add(len(sf.Servers))

	s := NewSemaphore(*concurrency)

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

	wg.Wait()
}

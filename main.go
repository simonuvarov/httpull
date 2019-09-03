package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/lair-framework/go-nmap"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// read data from stdin
	data, err := ioutil.ReadAll(os.Stdin)
	check(err)

	// parse nmap output
	run, err := nmap.Parse(data)
	check(err)

	// go through hosts
	for _, host := range run.Hosts {

		// check that the host is up
		if host.Status.State != "up" {
			continue
		}

		for _, p := range host.Ports {

			var schema string

			if p.State.State != "open" {
				continue
			}

			if !strings.Contains(p.Service.Name, "http") {
				continue
			}

			if p.Service.Tunnel == "ssl" || strings.Contains(p.Service.Name, "https") {
				schema = "https"
			} else {
				schema = "http"
			}

			if p.PortId != 80 && p.PortId != 443 {
				fmt.Printf("%s://%s:%d\n", schema, host.Addresses[0].Addr, p.PortId)
			} else {
				fmt.Printf("%s://%s\n", schema, host.Addresses[0].Addr)
			}

		}
	}
}

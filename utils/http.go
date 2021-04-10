package utils

import (
	"log"
	"net"
)

func GetIPByDomain(Domain string) (ip string) {
	iprecords, err := net.LookupHost(Domain)
	if err != nil {
		log.Fatal(err, " :do not get domain ip")
	}
	ip = iprecords[len(iprecords)-1]

	return ip
}

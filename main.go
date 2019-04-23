package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	sqlite "github.com/gwenn/gosqlite"
)

var (
	// Display debugging info to the console
	debug = false
)

func main() {
	if debug {
		fmt.Printf("Command line arguments found: %v\n", os.Args[1:])
	}

	// Make sure a single command line argument was given
	if len(os.Args[1:]) != 1 {
		log.Fatal("Expecting IP address as command line argument, but no arguments provided")
	}
	ipString := os.Args[1:][0]

	// Open the Geo-IP database
	db, err := sqlite.Open("Geo-IP.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	// Automatically close the SQLite database when this function finishes
	defer func() {
		db.Close()
	}()

	// Break the IPv4 address into octets
	var part1, part2, part3, part4 int
	ip := strings.Split(ipString, ".")
	if len(ip) != 4 {
		log.Fatalf("Unknown IPv4 address string format")
	}
	part1, err = strconv.Atoi(ip[0])
	if err != nil {
		log.Fatal(err)
	}
	part2, err = strconv.Atoi(ip[1])
	if err != nil {
		log.Fatal(err)
	}
	part3, err = strconv.Atoi(ip[2])
	if err != nil {
		log.Fatal(err)
	}
	part4, err = strconv.Atoi(ip[3])
	if err != nil {
		log.Fatal(err)
	}

	// Convert the IP address pieces into the correct lookup value
	ipVal := part4 + (part3 * 256) + (part2 * 256 * 256) + (part1 * 256 * 256 * 256)

	if debug {
		fmt.Printf("Lookup value for '%s' is: %v\n", ipString, ipVal)
	}

	// Look up the country code for the IP address
	var country string
	dbQuery := `
		SELECT cntry
		FROM ipv4
		WHERE ipfrom < ?
			AND ipto > ?`
	stmt, err := db.Prepare(dbQuery)
	if err != nil {
		log.Fatalf("Error when preparing statement for database: %s\n", err)
	}
	defer stmt.Finalize()
	err = stmt.Select(func(s *sqlite.Stmt) (innerErr error) {
		innerErr = s.Scan(&country)
		return
	}, ipVal, ipVal)
	if err != nil {
		log.Fatal(err)
	}

	if debug {
		fmt.Printf("Country lookup result is '%s'\n", country)
	}

	// Print the country code
	fmt.Println(country)
}

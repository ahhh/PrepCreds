package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

var (
	serverLog *os.File
	method    = flag.String("method", "user", "the auth mechanism to use for the authentication attempt (user, spray)")

	creds = []string{}

	users     = []string{}
	user      = flag.String("user", "", "indicate user to brute force")
	userList  = flag.String("userList", "", "indicate wordlist file that has users on each line")
	passwords = []string{}
	pass      = flag.String("pass", "", "indicate password to use to brute force")
	passList  = flag.String("passList", "", "indicate wordlist file that has passwords on each line")

	outFile = flag.String("out", "creds.txt", "indicate output file that will have cred combo per line (username:password)")
)

func main() {
	flag.Parse()
	if paramCheck() {
		of, err := os.OpenFile(*outFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer of.Close()

		// debug
		// log.Printf("%v", users)
		// log.Printf("%v", passwords)
		switch *method {
		case "user":
			for _, us := range users {
				for _, ps := range passwords {
					of.WriteString(us + ":" + ps + "\n")
				}
			}
		case "spray":
			for _, ps := range passwords {
				for _, us := range users {
					of.WriteString(us + ":" + ps + "\n")
				}
			}
		}
	}
}

func paramCheck() bool {
	canRun := true
	// Make sure a user is set
	if (*user != "") || (*userList != "") {
		if *userList != "" {
			listUsers := readLines(*userList)
			users = append(users, listUsers...)
		}
		if *user != "" {
			users = append(users, *user)
		}
	} else {
		log.Println("No user or userList provided!")
		canRun = false
	}
	// Make sure Cred or CredList is set
	if (*pass != "") || (*passList != "") {
		if *passList != "" {
			listPass := readLines(*passList)
			passwords = append(passwords, listPass...)
		}
		if *pass != "" {
			passwords = append(passwords, *pass)
		}
	} else {
		log.Println("No pass or passList provided!")
		canRun = false
	}
	if len(users) == 0 || len(passwords) == 0 {
		log.Println("Error preping credentials, check list files")
		canRun = false
	}

	if !canRun {
		log.Println("Missing mandatory paramaters. use -h for the help menu.")
		return false
	}
	return true

}

func readLines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

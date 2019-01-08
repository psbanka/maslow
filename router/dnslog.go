package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func parseLog(s string) (string, string, error) {
	splitstr := strings.Split(s, ": ")
	if len(splitstr) != 2 {
		return "", "", errors.New("Unreadable line")
	}
	dataStr := splitstr[1]
	splitStr2 := strings.Split(dataStr, " ")
	if splitStr2[0] != "reply" {
		return "", "", nil
	}
	// TODO: Deal properly with CNAMEs?
	if splitStr2[3] == "<CNAME>" {
		return "", "", nil
	}
	name := strings.Replace(splitStr2[1], "\"", "", -1)
	iPAddress := splitStr2[3]
	return name, iPAddress, nil
}

func readAndParseDNS(filename string) (map[string][]string, error) {
	g, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return nil, err
	}

	s := bufio.NewScanner(g)
	output := make(map[string][]string)
	for s.Scan() {
		logLine := s.Text()
		key, value, err := parseLog(logLine)
		if err == nil {
			if key != "" {
				output[key] = append(output[key], value)
			}
		}
	}
	g.Close()
	return output, nil
}
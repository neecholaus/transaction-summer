package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sources := transactionSources()

	dict := make(map[string]float64)

	for _, v := range sources {
		processSource(v, &dict)
	}

	// sort keys
	keys := make([]string, 0, len(dict))
	for key := range dict {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(key, dict[key])
	}
}

func transactionSources() []string {
	files, err := ioutil.ReadDir("./records")
	if err != nil {
		return []string{}
	}

	var sources []string
	for _, v := range files {
		if strings.Contains(v.Name(), ".csv") {
			sources = append(sources, v.Name())
		}
	}

	return sources
}

func processSource(source string, dict *map[string]float64) {
	file, err := os.Open("./records/" + source)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		cleaned := line[1 : len(line)-1]
		split := strings.Split(cleaned, "\",\"")

		category := split[8]
		amount, _ := strconv.ParseFloat(split[4], 64)

		if category != "" {
			(*dict)[category] += (amount * -1)
		} else if amount < 0 {
			(*dict)["Unknown"] += amount
		}
	}
}

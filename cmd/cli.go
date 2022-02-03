package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/satheesh1997/sache"
)

var cacheSize int

func main() {
	cache := sache.New(cacheSize)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")

		command, _ := reader.ReadString('\n')
		command = strings.Replace(command, "\n", "", -1)
		commandSplit := strings.Split(command, " ")
		method := strings.ToLower(commandSplit[0])

		switch method {
		case "get":
			if len(commandSplit) < 2 {
				fmt.Println("Usage:")
				fmt.Println("  GET <key>")
			} else {
				value := cache.Get(commandSplit[1])
				if value != "" {
					fmt.Println(value)
				}
			}
		case "set":
			if len(commandSplit) < 3 {
				fmt.Println("Usage:")
				fmt.Println("  SET <key> <value>")
			} else {
				cache.Put(commandSplit[1], commandSplit[2])
			}
		case "exit":
			os.Exit(0)
		case "exit()":
			os.Exit(0)
		default:
			fmt.Println("Usage:")
			fmt.Println("  SET <key> <value> - to add data to cache")
			fmt.Println("  GET <key>         - to get data from cache")
		}
	}
}

func init() {
	size := os.Getenv("SACHE_SIZE")
	if size == "" {
		cacheSize = 100
	} else {
		var err error
		cacheSize, err = strconv.Atoi(size)
		if err != nil {
			cacheSize = 100
		}
	}
}

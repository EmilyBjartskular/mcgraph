package main

import (
	"flag"
	"log"
	"os"
)

var (
	// Flags
	debug     bool
	verbose   bool
	modLoader string
)

func init() {
	flag.BoolVar(&debug, "debug", false, "Enable debug mode.")
	flag.BoolVar(&verbose, "verbose", false, "Enable verbose logging.")

	flag.StringVar(&modLoader, "loader", "fabric", "Set modloader.")
}

func main() {
	flag.Parse()

	// Verify directory argument
	args := flag.Args()
	if len(args) == 0 {
		log.Fatalln("missing directory argument")
	}
	fileInfo, err := os.Stat(args[len(args)-1])
	if err != nil {
		log.Fatalln(err)
	}
	if !fileInfo.IsDir() {
		log.Fatalln("not a directory")
	}

	mods, err := ReadMods()
	if err != nil {
		log.Fatalln(err)
	}

	if verbose {
		log.Printf("mods: %v\n", mods)
	}

	graph := GenerateGraph(mods)
	err = graph.Render()
	if err != nil {
		log.Fatalln(err)
	}
}

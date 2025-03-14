package main

import "flag"

func parseArgs() map[string]any {
	filename := flag.String("filename", "generatedGUI", "the name of the executable without the extension")

	flag.Parse()

	return map[string]any{
		"filename": filename,
	}
}

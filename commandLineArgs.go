package main

import "flag"

func parseArgs() map[string]any {
	filename := flag.String("filename", "generatedGUI", "the name of the executable without the extension")
	deleteBuildFiles := flag.Bool("deleteBuildFiles", true, "wether to remove the temp files of the build")
	devMode := flag.Bool("devMode", false, "if true then the GUI won't be compilated and the build files won't be deleted")

	flag.Parse()

	return map[string]any{
		"filename":         *filename,
		"deleteBuildFiles": *deleteBuildFiles,
		"devMode":          *devMode,
	}
}

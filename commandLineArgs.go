package main

import "flag"

func parseArgs() map[string]any {
	filename := flag.String("filename", "generatedGUI", "the name of the executable without the extension")
	command := flag.String("command", "", "the command to be executed after the gui is set up")
	deleteBuildFiles := flag.Bool("deleteBuildFiles", true, "wether to remove the temp files of the build")
	devMode := flag.Bool("devMode", false, "if true then the GUI won't be compilated and the build files won't be deleted")

	flag.Parse()

	return map[string]any{
		"filename":         *filename,
		"command":          *command,
		"deleteBuildFiles": *deleteBuildFiles,
		"devMode":          *devMode,
	}
}

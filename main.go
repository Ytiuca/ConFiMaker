package main

func main() {
	args := parseArgs()
	filename := (args["filename"].(string))
	command := (args["command"].(string))
	deleteBuildFiles := (args["deleteBuildFiles"].(bool))
	devMode := (args["devMode"].(bool))
	ps := newPythonSkeleton(command, filename)
	ps.GetFields()
	ps.CreatePythonFile()
	if !devMode {
		ps.Build(filename, deleteBuildFiles)
	}
}

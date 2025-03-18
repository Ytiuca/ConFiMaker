package main

func main() {
	args := parseArgs()
	filename := (args["filename"].(string))
	command := (args["command"].(string))
	deleteBuildFiles := (args["deleteBuildFiles"].(bool))
	noCompil := (args["noCompil"].(bool))
	ps := newPythonSkeleton(command, filename)
	ps.GetFields()
	ps.CreatePythonFile()
	if !noCompil {
		ps.Build(filename, deleteBuildFiles)
	}
}

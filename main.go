package main

func main() {
	args := parseArgs()
	filename := (args["filename"].(string))
	deleteBuildFiles := (args["deleteBuildFiles"].(bool))
	devMode := (args["devMode"].(bool))
	ps := newPythonSkeleton()
	ps.GetFields()
	ps.CreatePythonFile(filename)
	if !devMode {
		ps.Build(filename, deleteBuildFiles)
	}
}

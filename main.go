package main

func main() {
	args := parseArgs()
	filename := (args["filename"].(*string))
	ps := newPythonSkeleton()
	ps.GetFields()
	ps.CreatePythonFile(*filename)
	ps.Build(*filename)
}

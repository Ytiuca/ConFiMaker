package main

func main() {
	ps := newPythonSkeleton()
	ps.GetFields()
	ps.CreatePythonFile()
	ps.Build()
}

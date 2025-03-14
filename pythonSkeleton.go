package main

import (
	"log"
	"os"
)

type PythonSkeleton struct {
	beginning string
	ending    string
}

func newPythonSkeleton() *PythonSkeleton {
	return &PythonSkeleton{
		"import customtkinter as tk" +
			"\n" +
			"class Skeleton(tk.CTk):" +
			"\n" +
			"\tdef __init__(self, fg_color=None, **kwargs):" +
			"\n" +
			"\t\tsuper().__init__(fg_color, **kwargs)" +
			"\n",
		"\t\tself.mainloop()",
	}
}

func (ps *PythonSkeleton) createPythonFile() {
	path := "./generatedPythonFile.py"
	err := os.Remove(path)

	if err != nil {
		log.Fatalf("Erreur lors de la supression du fichier %s", err)
	}

	file, err := os.Create(path)

	if err != nil {
		log.Fatalf("Erreur lors de la création du fichier %s", err)
	}

	defer file.Close()

	file.WriteString(ps.beginning + ps.ending)
}

func main() {
	ps := newPythonSkeleton()
	ps.createPythonFile()
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type PythonSkeleton struct {
	beginning       string
	fields          []string
	ending          string
	correspondances map[string]string
}

func newPythonSkeleton() *PythonSkeleton {
	return &PythonSkeleton{
		"from customtkinter import *" +
			"\n" +
			"class Skeleton(CTk):" +
			"\n" +
			"\tdef __init__(self, fg_color=None, **kwargs):" +
			"\n" +
			"\t\tsuper().__init__(fg_color, **kwargs)" +
			"\n",
		[]string{},
		"\t\tself.mainloop()\n" +
			"Skeleton()",
		map[string]string{
			"checkbox": "CTkCheckBox",
			"entry":    "CTkEntry",
			"input":    "CTkEntry",
			"button":   "CTkButton",
			"slider":   "CTkSlider",
		},
	}
}

func (ps *PythonSkeleton) CreatePythonFile() {
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

	file.WriteString(ps.beginning)

	for _, field := range ps.fields {
		f, exists := ps.correspondances[field]

		if exists {
			file.WriteString("\t\t" + f + "(self).pack()\n")
		} else {
			log.Printf("erreur lors de la conversion de %s en un composant customTkinter", field)
		}
	}

	file.WriteString(ps.ending)
}

func (ps *PythonSkeleton) GetFields() {
	path := "./description.conFM"

	file, err := os.Open(path)

	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture du fichier de description %s", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ps.fields = append(ps.fields, scanner.Text())
	}
}

func main() {
	ps := newPythonSkeleton()
	ps.GetFields()
	fmt.Println(ps.fields)
	ps.CreatePythonFile()
}

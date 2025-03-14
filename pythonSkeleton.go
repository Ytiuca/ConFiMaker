package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
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
		"\t\tCTkButton(self,text=\"Run\",command=run).pack()\n" +
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

func (ps *PythonSkeleton) CreatePythonFile(filename string) {
	path := "./" + filename + ".py"
	os.Remove(path)

	file, err := os.Create(path)

	if err != nil {
		log.Fatalf("Erreur lors de la création du fichier %s. %s", path, err)
	}

	defer file.Close()

	file.WriteString(ps.beginning)

	for _, field := range ps.fields {
		f, exists := ps.correspondances[field]

		if exists {
			file.WriteString("\t\t" + f + "(self).pack()\n")
		} else {
			log.Printf("erreur lors de la conversion de '%s' en un composant customTkinter", field)
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

func (ps *PythonSkeleton) Build(filename string, deleteBuildFiles bool) {
	cmd := exec.Command("pyinstaller", "--onefile", "--noconsole", "./"+filename+".py")
	_, err := cmd.Output()

	if err != nil {
		log.Fatalf("Erreur lors de la compilation de l'interface graphique %s", err)
	}

	err = os.Rename("dist/"+filename+".exe", "./"+filename+".exe")

	if err != nil {
		log.Fatalf("Erreur lors du déplacement de l'interface graphique compilée dans le dossier local %s", err)
	}

	if deleteBuildFiles {
		removeFileAndDir([]string{"build", "dist", "./" + filename + ".spec", "./" + filename + ".py"})
	}

}

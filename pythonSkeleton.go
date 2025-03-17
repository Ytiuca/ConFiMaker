package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type PythonSkeleton struct {
	widgets         []Widget
	command         string
	correspondances map[string]string
}

func newPythonSkeleton(command string) *PythonSkeleton {
	return &PythonSkeleton{
		[]Widget{},
		command,
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

	file.WriteString(ps.ToPython())
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
		attributes := strings.Split(scanner.Text(), " ")
		widget := ToWidget(attributes)
		ps.widgets = append(ps.widgets, widget)
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

func ToWidget(attributes []string) Widget {
	widgetType := attributes[0]
	widgetLabel := attributes[1]
	widgetOptions := attributes[2:]

	switch widgetType {
	case "checkbox":
		isChecked, err := strconv.ParseBool(widgetOptions[0])
		if err != nil {
			isChecked = false
		}
		return newCheckbox(widgetLabel, isChecked)
	default:
		log.Fatalf("Erreur. Ce type de widget n'est pas encore pris en charge: %s", widgetType)
		return nil
	}
}

func (ps *PythonSkeleton) ToPython() string {
	return ps.CreateImport() +
		ps.CreateClass() +
		ps.CreateInitFunc() +
		ps.CreateRunFunc() +
		"Skeleton()"
}

func (ps *PythonSkeleton) CreateRunFunc() string {
	splitedCommand := strings.Split(ps.command, " ")
	command := "'" + splitedCommand[0] + "'"
	if len(splitedCommand) > 1 {
		for _, arg := range splitedCommand[1:] {
			command += "," + "'" + arg + "'"
		}
	}
	return INDENT + "def run(self):" +
		NEWLINE +
		ps.WidgetsToGetters() +
		DOUBLE_INDENT + "try:" +
		NEWLINE +
		TRIPLE_INDENT + "subprocess.run([" + command + ps.WidgetToArgs() + "],capture_output=True,text=True,check=True)" +
		NEWLINE +
		GestionExceptions("subprocess.CalledProcessError", "e.stderr") +
		GestionExceptions("Exception", "e")
}

func (ps *PythonSkeleton) CreateImport() string {
	return "from customtkinter import *" +
		NEWLINE +
		"import subprocess" +
		NEWLINE
}

func (ps *PythonSkeleton) CreateClass() string {
	return "class Skeleton(CTk):" +
		NEWLINE
}

func (ps *PythonSkeleton) CreateInitFunc() string {
	return INDENT + "def __init__(self, fg_color=None, **kwargs):" +
		NEWLINE +
		DOUBLE_INDENT + "super().__init__(fg_color, **kwargs)" +
		NEWLINE +
		ps.CreateWidgets() +
		DOUBLE_INDENT + "CTkButton(self,text=\"Run\",command=self.run).pack()" +
		NEWLINE +
		DOUBLE_INDENT + "self.mainloop()" +
		NEWLINE
}

func (ps *PythonSkeleton) CreateWidgets() string {
	retour := ""
	for _, widget := range ps.widgets {
		retour += widget.ToPython()
	}
	return retour
}

func (ps *PythonSkeleton) WidgetsToGetters() string {
	retour := ""

	for _, widget := range ps.widgets {
		retour += widget.ToGetter()
	}

	return retour
}

func (ps *PythonSkeleton) WidgetToArgs() string {
	retour := ""

	for _, widget := range ps.widgets {
		retour += widget.ToArg()
	}

	return retour
}

func GestionExceptions(exception string, errorVar string) string {
	return DOUBLE_INDENT + "except " + exception + " as e:" +
		NEWLINE +
		TRIPLE_INDENT + "popup = CTkToplevel(self)" +
		NEWLINE +
		TRIPLE_INDENT + "CTkLabel(popup, text=" + errorVar + ").pack()" +
		NEWLINE +
		TRIPLE_INDENT + "popup.focus" +
		NEWLINE
}

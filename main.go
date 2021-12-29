package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/Leonardo-Antonio/launchers__app/entity"
)

func main() {
	start := "\tCREADOR DE LANZADORES DE APLICACIONES\t"
	fmt.Println(start)

	launcher := &entity.Launcher{}

	fmt.Print("Tipo de aplicación? (Application) ")
	fmt.Scanln(&launcher.Type)
	if launcher.Type == "" {
		launcher.Type = "Application"
	}

	fmt.Print("Nombre de aplicación? ")
	fmt.Scanln(&launcher.Name)

	fmt.Print("Categoria de aplicación? (Utility) ")
	fmt.Scanln(&launcher.Category)
	if launcher.Category == "" {
		launcher.Category = "Utility"
	}

	fmt.Print("Icono de aplicación? [path] ")
	fmt.Scanln(&launcher.IconPath)

	fmt.Print("Ejecutable de aplicación? [path] ")
	fmt.Scanln(&launcher.ExecPath)

	dataFile := `
	[Desktop Entry]
	Type=%s
	Categories=%s
	Name=%s
	Icon=%s
	Exec=%s
	`

	f := fmt.Sprintf(dataFile, launcher.Type, launcher.Category, launcher.Name, launcher.IconPath, launcher.ExecPath)

	if err := ioutil.WriteFile(fmt.Sprintf("%s/.local/share/applications/%s.desktop", os.Getenv("HOME"), launcher.Name), []byte(f), 0777); err != nil {
		log.Fatalln(err)
	}

	cmd := exec.Command("sh", "update.sh")
	_, err := cmd.Output()
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Se creo el lanzador correctamente")
}

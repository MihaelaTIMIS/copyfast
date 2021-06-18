package fssync

import (
	"fmt"
	"log"
	"os"

	goftp "github.com/martinr92/goftp"
)

func ScanRecursive(origin string, target string, ftp string) {
	fmt.Println(origin, target)

	if _, err := os.Stat(origin); os.IsNotExist(err) {
		os.Exit(1)
	}

	if _, err := os.Stat(target); os.IsNotExist(err) {
		os.MkdirAll(target, 0777)
	}

	if ftp == "1" {
		fmt.Println("copie vers FTP")
		copyFolderFTP(origin)
	} else {

		nbMission := scan(origin, target)
		Run()
		for nbMission != TotalEnCour() {

		}
	}

	fmt.Println("Finish")
}

func copyFolderFTP(origin string) {
	fmt.Println(origin, "récuperer dossier FTP")

	var error error

	ftpClient, error := goftp.NewFtp("127.0.0.1:21")

	if error != nil {
		panic(error)
	}
	defer ftpClient.Close()

	files_origin, error := os.ReadDir(origin)
	for _, item := range files_origin {
		if item.IsDir() {
			ftpClient.CreateDirectory(item.Name())
			copyFolderFTP(origin + item.Name() + "/")
		} else {
			fmt.Println(item.Name(), "...envoie vers FTP")

			if error = ftpClient.Upload(origin+item.Name(), item.Name()); error != nil {
				panic(error)
			}

			fmt.Println(item.Name(), "dossier copié sur FTP")
		}
	}
}

func scan(origin string, target string) int {
	var nbMission int = 0
	files, err := os.ReadDir(origin)
	if err != nil {
		log.Fatal(err)
	} else {
		for _, file := range files {

			if file.IsDir() {
				nbMission += scan(origin+"/"+file.Name(), target+"/"+file.Name())
			} else {
				go AddMission(&Mission{
					origin: origin,
					target: target,
					file:   file,
				})
				nbMission++
			}

		}
	}

	return nbMission
}

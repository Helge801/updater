package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	writeScript()
}

func writeScript() {
	var data map[string]string
	content, err := ioutil.ReadFile("./config.JSON")
	checkErr(err)
	err = json.Unmarshal(content, &data)
	checkErr(err)
	path, exists := data["path"]
	if !exists {
		log.Fatal("ERROR: cannot find path attribute in config")
		os.Exit(1)
	}
	file, exists := data["file"]
	if !exists {
		log.Fatal("ERROR: cannot find file attribute in config")
		os.Exit(1)
	}
	script := fmt.Sprintf("cd %s\ngit pull origin master\ngo run %s & disown", path, file)
	ioutil.WriteFile("getter.sh", []byte(script), 0665)
	command := exec.Command("chmod", "+x", "./getter.sh")
	command.Run()
	cmd := exec.Command("/bin/sh", "~/go/src/github.com/updater/getter.sh")
	cmd.Run()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

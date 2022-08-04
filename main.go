package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func readFile(fileName string) ([]byte, error) {

	_, err := os.Stat(fileName)

	if os.IsNotExist(err) {
		fmt.Printf("err: %v\n", err)
		return nil, err
	}

	fileContent, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	return fileContent, nil
}

func main() {

	appsettingsFileName := flag.String("an", "", "path to the settings file")
	appsettingsFileType := flag.String("at", "", "settings file type: dotnet-json")
	envFileName := flag.String("en", "", "path to the environment variables file")
	h := flag.Bool("h", false, "help")

	flag.Parse()

	if *h {
		flag.PrintDefaults()
		return
	}

	if len(*appsettingsFileName) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(*envFileName) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(*appsettingsFileType) == 0 {
		*appsettingsFileType = "dotnet-json"
	}

	// read settings file
	// if file not exists, close programm with sucessuful code
	appsettings, err := readFile(*appsettingsFileName)
	if os.IsNotExist(err) {
		return
	}
	// read enviroment variables file
	envContent, err := readFile(*envFileName)
	if os.IsNotExist(err) {
		os.Exit(1)
	}

	var remover Remover

	// init dotnet app setting remover
	if *appsettingsFileType == "dotnet-json" {

		dotnetRemover := DotnetAppSettingsRemover{
			EnvironmentData: string(envContent),
		}

		err = json.Unmarshal(appsettings, &dotnetRemover.AppSettings)

		if err != nil {
			panic(err)
		}

		remover = dotnetRemover
	} else {
		panic("flag -at must be dotnet-json")
	}

	res := remover.RemoveEnvVariable() //removeEnvVariable("", m, string(envContent))
	ioutil.WriteFile(*envFileName, []byte(res), 0644)
}

type Remover interface {
	RemoveEnvVariable() string
}

package gui

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/rivo/tview"
	// "github.com/g-e-e-z/gogetr/config"
)

type Environment map[string]string

type Environments map[string]*Environment

func LoadEnvironments(directory string) (*Environments, error) {
    environments := make(Environments)

	// Read all files in the specified path
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

    for _, file := range files {
        if !file.IsDir() && filepath.Ext(file.Name()) == ".env" {
            var environment Environment
            environment, err  := godotenv.Read(filepath.Join(directory, file.Name()))
            if err != nil {
                return nil, err
            }
            environments[file.Name()] = &environment

        }
    }
    return &environments, nil

}

func NewEnvSelector() *tview.TextView {
	// TODO: This is bad
	pwd, _ := os.Getwd()
	requestsDir := filepath.Join(pwd, "requests_dir")
    fmt.Println("TESTSET", requestsDir)
	environments, err := LoadEnvironments(requestsDir)
	if err != nil {
		log.Panic(err)
	}
    environment := tview.NewTextView()//.SetBorder(true).SetTitle("Env and Build Info")

    for k, v := range *environments {
        fmt.Fprint(environment, k, v,"\n")
    }
    // input := tview.NewInputField().
    //     SetLabel("Select Environment: ").
    //     SetPlaceholder("dev.env, prod.env").
    //     SetAcceptanceFunc(func(text string, key rune) bool {
    //     // SetAcceptanceFunc(func(text string, key tview.Key) bool {
    //         // Logic to load the selected environment file
    //         // config.LoadEnvFile(text) // hypothetical function
    //         return true
    //     })

    return environment
}


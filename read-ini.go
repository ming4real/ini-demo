package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("sample.ini")
    if err != nil {
        fmt.Printf("Fail to read file: %v", err)
        os.Exit(1)
    }

	// Default sections are an empty string
	fmt.Println("App Mode:", cfg.Section("").Key("app_mode").String())
	fmt.Printf("Port Number: %[1]d\n", cfg.Section("server").Key("http_port").MustInt(9999))
	fmt.Println("Data Path:", cfg.Section("paths").Key("data").String())
}
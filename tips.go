package main

import (
    "fmt"
    "os"
    "os/user"
    "flag"
    "bytes"
    "io/ioutil"
    "path/filepath"
    "github.com/BurntSushi/toml"
)

type Config struct {
    Repository string
    Branch string
}

const DefaultRepository = "https://github.com/Te-k/commands-for-sec"
const DefaultBranch = "master"

func getconfiguration() (conf Config) {
    usr, _ := user.Current()
    config_path := filepath.Join(usr.HomeDir, ".tips")

    if _, err := toml.DecodeFile(config_path, &conf); err != nil {
        // Bad configuration, update conf file with default values
        fmt.Printf("Bad configuration, updated with default values\n")
        conf.Repository = DefaultRepository
        conf.Branch = DefaultBranch

        buf := new(bytes.Buffer)
        if err := toml.NewEncoder(buf).Encode(conf); err != nil {
            fmt.Println("Error with toml format, quitting")
            os.Exit(1)
        }
        if err := ioutil.WriteFile(config_path, buf.Bytes(), 0600); err != nil {
            fmt.Println("Error with config path: %s", err)
            os.Exit(1)
        }
    }
    return
}

func main() {
    fmt.Printf("Tips !\n")

    // Command line arguments
    flag.Usage = func() {
        fmt.Printf("Usage of %s:\n", os.Args[0])
        fmt.Printf("tips COMMAND\n")
        flag.PrintDefaults()
    }
    optionConfig := flag.Bool("c", false, "List configuration")
    flag.Parse()

    // Configuration
    var conf Config
    conf = getconfiguration()

    // Display configuration
    if *optionConfig {
        fmt.Printf("Configuration:\n")
        fmt.Printf("-Repository: %s\n", conf.Repository)
        fmt.Printf("-Branch: %s\n", conf.Branch)
        os.Exit(0)
    }
    if flag.NArg() != 1 {
        flag.Usage()
        os.Exit(1)
    }
}

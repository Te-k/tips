package main

import (
    "fmt"
    "os"
    "os/user"
    "os/exec"
    "flag"
    "bytes"
    "io/ioutil"
    "path/filepath"
    "github.com/BurntSushi/toml"
)

type Config struct {
    Repository string
    Branch string
    ConfigPath string
    RepositoryPath string
}

type ConfigFile struct {
    Repository string
    Branch string
}

const DefaultRepository = "https://github.com/Te-k/commands-for-sec.git"
const DefaultBranch = "master"

func getconfiguration() (conf Config) {
    usr, _ := user.Current()
    main_dir := filepath.Join(usr.HomeDir, ".config/tips")
    conf.ConfigPath = filepath.Join(main_dir, "config")
    conf.RepositoryPath = filepath.Join(main_dir, "tips")

    // Check if directory exist
    if _, err := os.Stat(main_dir); err != nil {
        if os.IsNotExist(err) {
            os.Mkdir(main_dir, 0740)
        } else {
            fmt.Printf("Weird error: %s", err)
            os.Exit(1)
        }
    }

    // Read configuration file
    var conf_file ConfigFile
    if _, err := toml.DecodeFile(conf.ConfigPath, &conf_file); err != nil {
        // Bad configuration, update conf file with default values
        fmt.Printf("Bad configuration, updated with default values\n")
        conf_file.Repository = DefaultRepository
        conf_file.Branch = DefaultBranch

        buf := new(bytes.Buffer)
        if err := toml.NewEncoder(buf).Encode(conf_file); err != nil {
            fmt.Println("Error with toml format, quitting")
            os.Exit(1)
        }
        if err := ioutil.WriteFile(conf.ConfigPath, buf.Bytes(), 0600); err != nil {
            fmt.Println("Error with config path: %s", err)
            os.Exit(1)
        }
    }
    conf.Repository = conf_file.Repository
    conf.Branch = conf_file.Branch
    return
}

func pull(conf Config) {
    // Check if the repository exists
    if _, err := os.Stat(conf.RepositoryPath); err != nil {
        if os.IsNotExist(err) {
            // If not clone it
            cmdName := "git"
            cmdArgs := []string{"clone", conf.Repository, conf.RepositoryPath}
            out, err := exec.Command(cmdName, cmdArgs...).Output()
            if err != nil {
                fmt.Printf("Error: %s", err)
            }
            fmt.Printf("%s", out)
        } else {
            fmt.Printf("Weird error: %s", err)
            os.Exit(1)
        }
    } else {
        // Just pull the repository
        cmdName := "git"
        cmdArgs := []string{"-C", conf.RepositoryPath, "pull"}
        out, err := exec.Command(cmdName, cmdArgs...).Output()
        if err != nil {
            fmt.Printf("Error: %s", err)
        }
        fmt.Printf("%s", out)
    }
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
    pullConfig := flag.Bool("pull", false, "Pull new directory")
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
    if *pullConfig {
        fmt.Printf("Pull data from %s\n", conf.Repository)
        pull(conf)
        os.Exit(0)
    }
    if flag.NArg() != 1 {
        flag.Usage()
        os.Exit(1)
    }
    target_file := filepath.Join(conf.RepositoryPath, fmt.Sprintf("%s.md", flag.Args()[0]))
    if _, err := os.Stat(target_file); os.IsNotExist(err) {
        fmt.Printf("No tips for command %s\n", flag.Args()[0])
    } else {
        b, err := ioutil.ReadFile(target_file)
        if err != nil {
            fmt.Print(err)
        }
        str := string(b)
        fmt.Println(str)
    }
}

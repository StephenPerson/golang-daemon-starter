package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
    "os/exec"
    "os/signal"
    "strconv"
    "strings"
    "syscall"
)

var PIDFile = "/tmp/go-test.pid"

func savePID(pid int) {
    file, err := os.Create(PIDFile)
    if err != nil {
        log.Printf("Unable to create pid file : %v\n", err)
        os.Exit(1)
    }

    defer file.Close()

    _, err = file.WriteString(strconv.Itoa(pid))

    if err != nil {
        log.Printf("Unable to create pid file : %v\n", err)
        os.Exit(1)
    }
    file.Sync() // flush to disk
}

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
    html := "Hello World"
    w.Write([]byte(html))
}

func console() {
    ch := make(chan os.Signal, 1)
    signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGTERM)

    go func() {
        signalType := <-ch
        signal.Stop(ch)
        fmt.Println("Exit command received. Exiting...")
        fmt.Println("Received signal type : ", signalType)
        os.Remove(PIDFile)
        os.Exit(0)
    }()
    
    log.Println("Starting server.")
    mux := http.NewServeMux()
    mux.HandleFunc("/", sayHelloWorld)
    log.Fatalln(http.ListenAndServe(":8080",mux))
}

func start() {
    if _, err := os.Stat(PIDFile); err == nil {
        fmt.Printf("Already running or %s file exists.\n", PIDFile)
        os.Exit(1)
    }

    cmd := exec.Command(os.Args[0], "console")
    cmd.Start()
    fmt.Println("Daemon process ID is : ", cmd.Process.Pid)
    savePID(cmd.Process.Pid)
    os.Exit(0)
}

func stop() {
    if _, err := os.Stat(PIDFile); err == nil {
        data, err := ioutil.ReadFile(PIDFile)
        if err != nil {
            fmt.Println("Not running")
            os.Exit(1)
        }
        ProcessID, err := strconv.Atoi(string(data))
        if err != nil {
            fmt.Println("Unable to read and parse process id found in ", PIDFile)
            os.Exit(1)
        }
        process, err := os.FindProcess(ProcessID)
        if err != nil {
            fmt.Printf("Unable to find process ID [%v] with error %v\n", ProcessID, err)
            os.Exit(1)
        }
        // remove PID
        os.Remove(PIDFile)
        err = process.Kill()
        if err != nil {
            fmt.Printf("Unable to kill process ID [%v] with error %v\n", ProcessID, err)
            os.Exit(1)
        } else {
            fmt.Printf("Killed process ID [%v]\n", ProcessID)
            os.Exit(0)
        }
    } else {
        fmt.Println("Not running.")
        os.Exit(1)
    }
}

func status() {
    if _, err := os.Stat(PIDFile); err == nil {
        data, err := ioutil.ReadFile(PIDFile)
        if err != nil {
            fmt.Println("Not running")
            os.Exit(1)
        }
        ProcessID, err := strconv.Atoi(string(data))
        if err != nil {
            fmt.Println("Unable to read and parse process id found in ", PIDFile)
            os.Exit(1)
        }
        fmt.Printf("Running with process ID [%v]\n", ProcessID)
    } else {
        fmt.Printf("Not running.\n")
    }
    os.Exit(0)
}

func main() {
    if len(os.Args) != 2 {
        fmt.Printf("Usage: %s [start|stop]\n", os.Args[0])
        os.Exit(1)
    }

    switch strings.ToLower(os.Args[1]) {
        case "status":
            status()
        case "console":
            console()
        case "start":
            start()
        case "stop":
            stop()
        default:
            fmt.Printf("Unknown command : %v\n", os.Args[1])
            fmt.Printf("Usage: %s [console|start|stop|status]\n", os.Args[0])
            os.Exit(1)
    }
}
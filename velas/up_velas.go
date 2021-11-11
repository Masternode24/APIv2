package main

import (
        "log"
        "os/exec"
)

func UpService() string {
        out, err := exec.Command("bash", "-c", "velas validators | grep 8rmw3sav8hL2a72qg7otPuFYh5r8bt92si9AZjgZidGN | cut -d '%' -f2 | awk '{print $NF}' | tr -d '\n'").Output()
        if err != nil {
                log.Fatal(err)
        }
        verString := string(out[:])
        return verString
}

package main

import (
        "log"
        "os/exec"
)

func UpService() string {
        out, err := exec.Command("bash", "-c", "solana validators | grep 6h9jyRgfpmgXNyaWpbDpbxbCoF56WEbzsruhMwDn2om4 | cut -d '%' -f2 | awk '{print $NF}' | tr -d '\n'").Output()
        if err != nil {
                log.Fatal(err)
        }
        verString := string(out[:])
        return verString
}

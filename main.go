package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
    signature, running := os.LookupEnv("HYPRLAND_INSTANCE_SIGNATURE")

    if !running {
        panic("Hyprland is not running")
    }
    
    runtimeDir, found := os.LookupEnv("XDG_RUNTIME_DIR")

    if !found {
        panic("Can't find runtime directory (XDG_RUNTIME_DIR is unset)")
    }

    socketPath := fmt.Sprintf("%s/hypr/%s/.socket2.sock", runtimeDir, signature)

    conn, err := net.Dial("unix", socketPath)
    if err != nil {
        panic(err)
    }

    for { 
        buff := make([]byte, 256)
        length, err := conn.Read(buff)
        if err != nil {
            panic(err)
        }

        rawMessage := string(buff[:length])
        var builder strings.Builder
        for _, char := range rawMessage {
            if char == '\n' {
                break
            }
            builder.WriteRune(char)
        }

        parsedMessage := builder.String()

        fmt.Println(parsedMessage)
    }
}

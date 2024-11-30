package main

import (
    "fmt"
    "os"

    pb "protobuf-in-go/proto"
	"google.golang.org/protobuf/proto"
)

func Read() {
    in, err := os.ReadFile("user")
    if err != nil {
        fmt.Println("Error reading file:", err)
        return
    }

    user := &pb.User{}
    if err := proto.Unmarshal(in, user); err != nil {
        fmt.Println("Failed to parse user:", err)
        return
    }

    fmt.Println(user)
}

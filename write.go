package main

import (
	"fmt"
    "os"
    pb "protobuf-in-go/proto"

	"google.golang.org/protobuf/proto"
)

func Write() {
    // Create a new user
    user := pb.User{
        Id: 1,
        Name: "John Doe",
        Email: "jdoe@example.com",
        Type: pb.User_DIAMOND,
        Addresses: []*pb.User_Address{
            {Country: "USA", State: "CA"},
        },
    }

    // Encode the user to binary format
    data, err := proto.Marshal(&user)
    if err != nil {
        fmt.Println("marshal error:", err)
    }
    if err := os.WriteFile("user", data, 0644); err != nil {
        fmt.Println("write error:", err)
    }
    fmt.Println("user has been written to file")
}

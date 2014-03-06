package main

import (
  "fmt"
  "github.com/conformal/btcjson"
)

func main() {
  // Hello
  fmt.Println("Hello world")
  msg, err := btcjson.CreateMessage("getinfo")
  if err != nil {
    // Handle some error
  }
  reply, err := btcjson.RpcCommand("username", "password", "localhost:8332", msg)
  fmt.Println(reply)
}

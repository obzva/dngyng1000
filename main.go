package main

import (
	"log"

	"github.com/obzva/dngyng1000/server"
)

func main() {
	// postMap, err := server.NewPostMap(os.DirFS("posts"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// post, err := postMap.Get("this-is-the-title-2")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", post)
	log.Fatal(server.Run())
}

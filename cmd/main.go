package main

import (
	"log"

	server "github.com/obzva/dngyng1000"
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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rwynn/gtm"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	mgoHosts := flag.String("mgo", "127.0.0.1", "MongoDB hosts, comma separated")
	flag.Parse()

	session, err := mgo.Dial(*mgoHosts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	ctx := gtm.Start(session, nil)
	for {
		select {
		case err := <-ctx.ErrC:
			fmt.Println(err)
		case op := <-ctx.OpC:
			fmt.Println(op)
		}
	}
}

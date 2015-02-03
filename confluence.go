
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
        versionFlag = flag.Bool("version", false, "Print version info and exit.")

	version   string
	commit    string
	buildTime string
	sdkInfo   string
)

func init() {
        flag.Parse()
        if *versionFlag {
        	log.Printf("Version: %s, CommitID: %s, build time: %s, SDK Info: %s\n", version, commit, buildTime, sdkInfo)
                os.Exit(0)
        }
}

func main() {
	fmt.Printf("hello, world\n")
}

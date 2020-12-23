package main

import (
	"fmt"
	// "time"
	vlc_wrapper "github.com/0187773933/VLCTelnetWrapper/vlc"
)

func main() {
	vlc := vlc_wrapper.Wrapper{}
	vlc.Connect( "127.0.0.1:4212" )

	fmt.Println( vlc.Help() )
	fmt.Println( vlc.Status() )
	fmt.Println( vlc.Info() )
	fmt.Println( vlc.GetTime() )

	vlc.Disconnect()
}
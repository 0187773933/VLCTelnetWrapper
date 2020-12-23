# VLC Telnet Wrapper

```
package main

import (
	"fmt"
	vlc_wrapper "github.com/0187773933/VLCTelnetWrapper/vlc"
)

func main() {
	vlc := vlc_wrapper.Wrapper{}
	vlc.Connect( "127.0.0.1:4212" )

	fmt.Println( vlc.Help() )
	fmt.Println( vlc.Add( "/media/morphs/14TB/MEDIA_MANAGER/TVShows/DrakeAndJosh/001 - Drake and Josh - S01E01 - Pilot.mp4" ) )
	fmt.Println( vlc.Status() )
	fmt.Println( vlc.Info() )
	fmt.Println( vlc.GetTime() )
	fmt.Println( vlc.GetLength() )
	fmt.Println( vlc.Stats() )

	vlc.Disconnect()
}
```
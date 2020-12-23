package vlc

import (
	"fmt"
	"time"
	"strings"
	telnet "github.com/ziutek/telnet"
)

// https://github.com/DerMitch/py-vlcclient
// https://github.com/48723247842/VLCController/blob/master/vlc_controller/__init__.py


const timeout = 10 * time.Second
func expect( t *telnet.Conn , d ...string ) {
	error := t.SetReadDeadline( time.Now().Add( timeout ) )
	if error != nil { panic( error ) }
	error = t.SkipUntil( d... )
	if error != nil { panic( error ) }
}

func send_line( t *telnet.Conn , input string ) {
	error := t.SetWriteDeadline( time.Now().Add( timeout ) )
	if error != nil { panic( error ) }
	buf := make( []byte , ( len( input ) + 1 ) )
	copy( buf , input )
	buf[ len( input ) ] = '\n'
	_ , err := t.Write( buf )
	if err != nil { panic( err ) }
}

type Wrapper struct {
	Telnet *telnet.Conn
}

func ( vlc *Wrapper ) Connect( server_address string ) {
	var err error
	vlc.Telnet , err = telnet.Dial( "tcp" , server_address )
	if err != nil { panic( err ) }
	vlc.Telnet.SetUnixWriteMode( true )
	vlc.Telnet.SkipUntil( "Password:" )
	vlc.Telnet.Write( []byte( "admin\n" ) )
	// var data []byte
	// data , err = vlc.Telnet.ReadUntil( ">" )
	// fmt.Println( string( data[:] ) )
	vlc.Telnet.ReadUntil( ">" )
}

func ( vlc *Wrapper ) Disconnect() {
	vlc.Telnet.Close()
}

func ( vlc *Wrapper ) ReadResult() ( result string ) {
	result = "failed"
	var data []byte
	var err error
	data , err = vlc.Telnet.ReadUntil( ">" )
	if err != nil { panic( err ) }
	result = string( data[:] )
	lines := strings.Split( result , "\n" )
	all_but_last_line := lines[ 0 : ( len( lines ) - 1 ) ]
	result = strings.Join( all_but_last_line , "\n" )
	return
}

func ( vlc *Wrapper ) SendCommand( command string ) ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( fmt.Sprintf( "%s\n" , command ) ) )
	result = vlc.ReadResult()
	return
}

// func ( vlc *Wrapper ) Raw( commands ...string ) ( result string )  {
// 	result = "failed"
// 	vlc.Telnet.Write( []byte( fmt.Sprintf( "%s\n" ) , command ) )
// 	result = vlc.ReadResult()
// 	return
// }

func ( vlc *Wrapper ) Help() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "help\n" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Status() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "status\n" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Info() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "info\n" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) GetTime() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "get_time\n" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) GetLength() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "get_length\n" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Stats() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "stats\n" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) FullscreenOn() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "fullscreen on\n" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) FullscreenOff() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "fullscreen off\n" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) VolumeGet() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "volume\n" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) VolumeSet( volume_level string ) ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( fmt.Sprintf( "volume %s\n" , volume_level ) ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) VolumeUp( number_of_steps int ) ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( fmt.Sprintf( "volup %d\n" , number_of_steps ) ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) VolumeDown( number_of_steps int ) ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( fmt.Sprintf( "volup %d\n" , number_of_steps ) ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Add( file_path string ) ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( fmt.Sprintf( "add %s\n" , file_path ) ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Enqueue( file_path string ) ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( fmt.Sprintf( "enqueue %s\n" , file_path ) ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Seek( seconds int64 ) ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( fmt.Sprintf( "seek %d\n" , seconds ) ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Play() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "play" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Pause() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "pause" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Stop() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "stop" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Rewind() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "rewind" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Next() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "next" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Previous() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "previous" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Clear() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "clear" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Loop() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "loop" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Repeat() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "repeat" ) )
	result = vlc.ReadResult()
	return
}

func ( vlc *Wrapper ) Random() ( result string )  {
	result = "failed"
	vlc.Telnet.Write( []byte( "random" ) )
	result = vlc.ReadResult()
	return
}
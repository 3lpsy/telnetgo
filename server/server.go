package main


import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
    "flag"
	"io"
	"time"
    "fmt"
    "bytes"
    "strconv"
)


func fiveHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	oi.LongWriteString(stdout, "The number FIVE looks like this: 5\r\n")
	return nil
}

func fiveProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	return telsh.PromoteHandlerFunc(fiveHandler)
}

func danceHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	for i:=0; i<20; i++ {
		oi.LongWriteString(stdout, "\r⠋")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠙")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠹")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠸")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠼")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠴")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠦")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠧")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠇")
		time.Sleep(50*time.Millisecond)

		oi.LongWriteString(stdout, "\r⠏")
		time.Sleep(50*time.Millisecond)
	}
	oi.LongWriteString(stdout, "\r \r\n")

	return nil
}

func danceProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	return telsh.PromoteHandlerFunc(danceHandler)
}


func main() {
	shellHandler := telsh.NewShellHandler()

	shellHandler.WelcomeMessage = `
Welcome to Dancing Telnet Server.
 __          __ ______  _        _____   ____   __  __  ______
 \ \        / /|  ____|| |      / ____| / __ \ |  \/  ||  ____|
  \ \  /\  / / | |__   | |     | |     | |  | || \  / || |__
   \ \/  \/ /  |  __|  | |     | |     | |  | || |\/| ||  __|
    \  /\  /   | |____ | |____ | |____ | |__| || |  | || |____
     \/  \/    |______||______| \_____| \____/ |_|  |_||______|

`

    var portPtr = flag.Int("port", 9001, "Port to listen on")
    var addrPtr = flag.String("address", "0.0.0.0", "Address to listen on")

    flag.Parse()

    var port = *portPtr
    var addr = *addrPtr

    var listenerBuffer bytes.Buffer


    listenerBuffer.WriteString(addr)
    listenerBuffer.WriteString(":")
    listenerBuffer.WriteString(strconv.Itoa(port))

    listener := listenerBuffer.String()

    var messageBuffer bytes.Buffer
    messageBuffer.WriteString("Listening on address: ")
    messageBuffer.WriteString(listener)

    fmt.Println(messageBuffer.String())


	// Register the "five" command.
	var fileCommandName = "five"
	var fiveCommandProducer = telsh.ProducerFunc(fiveProducer)
	shellHandler.Register(fileCommandName, fiveCommandProducer)

	// Register the "dance" command.
	var danceCommandName = "dance"
	var danceCommandProducer = telsh.ProducerFunc(danceProducer)

	shellHandler.Register(danceCommandName, danceCommandProducer)
	shellHandler.Register("dance", telsh.ProducerFunc(danceCommandProducer))

	if err := telnet.ListenAndServe(listener, shellHandler); nil != err {
		panic(err)
	}
}

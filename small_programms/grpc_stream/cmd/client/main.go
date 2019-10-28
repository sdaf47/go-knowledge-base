package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/marcusolsson/tui-go"
	"github.com/sdaf47/go-knowledge-base/small_programms/grpc_stream/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"os"
	"time"
)

func main() {
	//conn, err := grpc.Dial("[::]:1024", grpc.WithInsecure())
	conn, err := grpc.Dial(":9981", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := stream.NewMessageBrokerClient(conn)
	username := read("Your name: ")
	password := read("Password: ")

	status, err := client.Subscribe(context.Background(), &stream.Logon{
		Username: username,
		Password: password,
	})
	if err != nil {
		panic(err)
	}
	if status.Error != "" {
		panic(status.Error)
	}

	streamClient, err := client.OpenStream(metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{
		"authorization": status.Token,
	})))
	if err != nil {
		panic(err)
	}

	history := tui.NewVBox()
	historyScroll := tui.NewScrollArea(history)
	historyScroll.SetAutoscrollToBottom(true)

	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)

	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)

	chat := tui.NewVBox(historyBox, inputBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)

	root := tui.NewHBox(chat)
	ui, err := tui.New(root)
	if err != nil {
		panic(err)
	}
	go func() {
		panic(ui.Run())
	}()

	input.OnSubmit(func(entry *tui.Entry) {
		defer entry.SetText("")
		err = streamClient.Send(&stream.Request{
			Message: entry.Text(),
		})
		if err != nil {
			panic(err)
		}
	})

	for {
		// chat messages
		msg, err := streamClient.Recv()
		if err != nil {
			panic(err)
		}

		history.Append(tui.NewHBox(
			tui.NewLabel(time.Now().Format("15:04:05")),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("%14s", msg.Username))),
			tui.NewLabel(msg.Message),
			tui.NewSpacer(),
		))

		time.Sleep(time.Millisecond)
		ui.Repaint()
	}
}

func read(q string) (a string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(q)
	a, _ = reader.ReadString('\n')
	fmt.Print("\r$i")
	return a[:len(a)-1]
}

package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/tarm/serial"
	tree "github.com/twstokes/treedivided"
	"github.com/twstokes/treedivided/pkg/controller"
	"github.com/twstokes/treedivided/pkg/fetcher"
	"github.com/twstokes/treedivided/pkg/player"
)

var (
	serialPort = flag.String("serialPort", "/dev/ttyUSB0", "serial port") // default port on RPi
	baudRate   = flag.Int("baud", 115200, "serial baud rate")
)

func main() {
	flag.Parse()

	// connect to the MCU
	sConfig := &serial.Config{Name: *serialPort, Baud: *baudRate}
	s, err := serial.OpenPort(sConfig)
	if err != nil {
		panic("Failed to connect to the MCU")
	}

	defer s.Close()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	stopChan := make(chan bool, 1)

	go func() {
		<-sig
		stopChan <- true
	}()

	clemson := tree.Team{
		ID:   tree.TeamA,
		Name: "Clemson Tigers",
		Colors: tree.TeamColors{
			Primary:   tree.NewColor(246, 103, 51), // orange
			Secondary: tree.NewColor(82, 45, 128),  // purple
		},
		SongPath: "cu.mp3",
	}

	usc := tree.Team{
		ID:   tree.TeamB,
		Name: "Carolina Gamecocks",
		Colors: tree.TeamColors{
			// swap their colors because orange and garnet look terrible at the same time
			Primary:   tree.NewColor(255, 255, 255), // white
			Secondary: tree.NewColor(115, 0, 10),    // garnet
		},
		SongPath: "usc.mp3",
	}

	i := controller.NewController(s)
	f := fetcher.NewLocalFetcher("scores.json")
	p := player.NewLocalPlayer()

	c := &tree.Config{
		TeamA:       clemson,
		TeamB:       usc,
		Illuminator: i,
		Fetcher:     f,
		Player:      p,
	}

	t := tree.NewTree(c)
	t.Run(stopChan)
}

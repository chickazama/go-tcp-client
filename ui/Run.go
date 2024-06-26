package ui

import (
	"errors"
	"log"
	"os"
	"os/exec"

	"github.com/awesome-gocui/gocui"
)

func Run() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
	// printGreeting()
	// Set up GUI
	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer g.Close()
	g.SetManagerFunc(setLayout)
	err = setKeyBindings(g)
	if err != nil {
		log.Fatal(err.Error())
	}
	// Spawn a go-routine to update the view
	go updateOutputView(g)
	if err := g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		log.Fatal(err.Error())
	}
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

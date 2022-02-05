package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"mrogalski.eu/go/pulseaudio"
)

func main() {
	flag.Parse()

	c, e := pulseaudio.NewClient()
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}

	if os.Args[0] == "puvol-cont" {
		cont(c)
	} else {
		single(c)
	}
}

func printvol(c *pulseaudio.Client) float32 {
	muted, e := c.Mute()
	if muted {
		fmt.Println("🔇")
		return 0
	}

	v, e := c.Volume()
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Printf("%d%%\n", int(100*v))
	}
	return v
}

func cont(c *pulseaudio.Client) {

	ch, e := c.Updates()
	if e != nil {
		log.Println(e)
		os.Exit(2)
	}

	for {
		printvol(c)
		_ = <-ch
	}
}

func single(c *pulseaudio.Client) {
	if len(os.Args) == 1 {
		printvol(c)
		return
	}

	switch a := os.Args[1]; a {
	case "toggle":
		c.ToggleMute()
		printvol(c)
	case "mute", "unmute":
		c.SetMute(a == "mute")
		printvol(c)
	case "inc", "dec":
		v, e := c.Volume()
		if e != nil {
			log.Println(e)
			os.Exit(2)
		}

		if a == "inc" {
			v += 0.1
		} else if a == "dec" {
			v -= 0.1
		}

		if v <= 0 {
			v = 0
		}

		c.SetVolume(float32(v))
		printvol(c)

	case "set":
		if len(os.Args) < 3 {
			log.Println("missing volume")
			os.Exit(3)
		}
		v, e := strconv.ParseFloat(os.Args[2], 32)
		if e != nil {
			log.Println(e)
			os.Exit(4)
		}
		c.SetVolume(float32(v))
		printvol(c)
	default:
		printvol(c)
	}
}

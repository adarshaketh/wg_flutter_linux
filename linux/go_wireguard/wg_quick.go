package main

/*
#cgo CFLAGS: -I./
#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread
#include <stdlib.h>
*/
import "C"
import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	wgquick "github.com/nmiculinic/wg-quick-go"
	"github.com/sirupsen/logrus"
)

func executeCommand(action, configPath string) (string, error) {
	verbose := flag.Bool("v", false, "verbose")
	protocol := flag.Int("route-protocol", 0, "route protocol to use for our routes")
	metric := flag.Int("route-metric", 0, "route metric to use for our routes")
	iface := flag.String("iface", "", "interface")
	flag.Parse()

	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	log := logrus.WithField("iface", *iface)

	_, err := os.Stat(configPath)
	switch {
	case err == nil:
	case os.IsNotExist(err):
		if *iface == "" {
			*iface = configPath
			log = logrus.WithField("iface", *iface)
		}
		configPath = "/etc/wireguard/" + configPath + ".conf"
		_, err = os.Stat(configPath)
		if err != nil {
			log.WithError(err).Errorln("cannot find config file")
			return "", fmt.Errorf("cannot find config file: %w", err)
		}
	default:
		logrus.WithError(err).Errorln("error while reading config file")
		return "", fmt.Errorf("error while reading config file: %w", err)
	}

	b, err := ioutil.ReadFile(configPath)
	if err != nil {
		logrus.WithError(err).Fatalln("cannot read file")
		return "", fmt.Errorf("cannot read file: %w", err)
	}
	c := &wgquick.Config{}
	if err := c.UnmarshalText(b); err != nil {
		logrus.WithError(err).Fatalln("cannot parse config file")
		return "", fmt.Errorf("cannot parse config file: %w", err)
	}

	c.RouteProtocol = *protocol
	c.RouteMetric = *metric

	switch action {
	case "up":
		if err := wgquick.Up(c, *iface, log); err != nil {
			logrus.WithError(err).Errorln("cannot up interface")
			return "", fmt.Errorf("cannot up interface: %w", err)
		}
	case "down":
		if err := wgquick.Down(c, *iface, log); err != nil {
			logrus.WithError(err).Errorln("cannot down interface")
			return "", fmt.Errorf("cannot down interface: %w", err)
		}
	case "sync":
		if err := wgquick.Sync(c, *iface, log); err != nil {
			logrus.WithError(err).Errorln("cannot sync interface")
			return "", fmt.Errorf("cannot sync interface: %w", err)
		}
	default:
		return "", fmt.Errorf("invalid action: %s", action)
	}

	return "success", nil
}

//export StartWireGuard
func StartWireGuard(configPath *C.char) *C.char {
	output, err := executeCommand("up", C.GoString(configPath))
	if err != nil {
		return C.CString(err.Error() + ": " + output)
	}
	return C.CString(output)
}

//export StopWireGuard
func StopWireGuard(configPath *C.char) *C.char {
	output, err := executeCommand("down", C.GoString(configPath))
	if err != nil {
		return C.CString(err.Error() + ": " + output)
	}
	return C.CString(output)
}

//export SyncWireGuard
func SyncWireGuard(configPath *C.char) *C.char {
	output, err := executeCommand("sync", C.GoString(configPath))
	if err != nil {
		return C.CString(err.Error() + ": " + output)
	}
	return C.CString(output)
}

func printHelp() {
	fmt.Print("wg-quick [flags] [ up | down | sync ] [ config_file | interface ]\n\n")
	flag.Usage()
	os.Exit(1)
}

func main() {
	flag.String("iface", "", "interface")
	verbose := flag.Bool("v", false, "verbose")
	protocol := flag.Int("route-protocol", 0, "route protocol to use for our routes")
	metric := flag.Int("route-metric", 0, "route metric to use for our routes")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		printHelp()
	}

	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
	}

	iface := flag.Lookup("iface").Value.String()
	log := logrus.WithField("iface", iface)

	cfg := args[1]

	_, err := os.Stat(cfg)
	switch {
	case err == nil:
	case os.IsNotExist(err):
		if iface == "" {
			iface = cfg
			log = logrus.WithField("iface", iface)
		}
		cfg = "/etc/wireguard/" + cfg + ".conf"
		_, err = os.Stat(cfg)
		if err != nil {
			log.WithError(err).Errorln("cannot find config file")
			printHelp()
		}
	default:
		logrus.WithError(err).Errorln("error while reading config file")
		printHelp()
	}

	b, err := ioutil.ReadFile(cfg)
	if err != nil {
		logrus.WithError(err).Fatalln("cannot read file")
	}
	c := &wgquick.Config{}
	if err := c.UnmarshalText(b); err != nil {
		logrus.WithError(err).Fatalln("cannot parse config file")
	}

	c.RouteProtocol = *protocol
	c.RouteMetric = *metric

	switch args[0] {
	case "up":
		if err := wgquick.Up(c, iface, log); err != nil {
			logrus.WithError(err).Errorln("cannot up interface")
		}
	case "down":
		if err := wgquick.Down(c, iface, log); err != nil {
			logrus.WithError(err).Errorln("cannot down interface")
		}
	case "sync":
		if err := wgquick.Sync(c, iface, log); err != nil {
			logrus.WithError(err).Errorln("cannot sync interface")
		}
	default:
		printHelp()
	}
}

package main

import (
	"flag"
	"time"

	graylog "github.com/gemnasium/logrus-graylog-hook/v3"
	log "github.com/sirupsen/logrus"
)

var (
	addr = flag.String("addr", "", "udp address")
	msg  = flag.String("msg", "empty", "message")
)

func main() {
	flag.Parse()

	hook := graylog.NewGraylogHook(*addr, map[string]interface{}{"app": "test-graylog"})
	defer hook.Flush()
	log.AddHook(hook)
	log.WithFields(log.Fields{
		"foo":  "foo" + time.Now().String(),
		"addr": *addr,
	})
	log.Info(*msg)
}

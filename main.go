package main

import (
	"fmt"
	"time"

	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Print("\n--- nested-logrus-formatter ---\n\n")
	printDemo(&formatter.Formatter{
		HideKeys:        false,
		FieldsOrder:     []string{"component", "category", "req"},
		CallerFirst:     true,
		TimestampFormat: time.RFC3339,
	}, "nested-logrus-formatter")
}

func printDemo(f logrus.Formatter, title string) {
	l := logrus.New()

	l.SetLevel(logrus.DebugLevel)
	l.SetReportCaller(true)

	if f != nil {
		l.SetFormatter(f)
	}

	// enable/disable file/function name
	// l.SetReportCaller(false)

	l.Infof("this is %v demo", title)

	lWebServer := l.WithField("component", "web-server")
	lWebServer.Info("starting...")

	lWebServerReq := lWebServer.WithFields(logrus.Fields{
		"req":   "GET /api/stats",
		"reqId": "#1",
	})

	lWebServerReq.Info("params: startYear=2048")
	lWebServerReq.Error("response: 400 Bad Request")

	lDbConnector := l.WithField("category", "db-connector")
	lDbConnector.Info("connecting to db on 10.10.10.13...")
	lDbConnector.Warn("connection took 10s")

	l.Info("demo end.")
}

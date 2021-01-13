package mymodule

import (
	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
	"time"
)

var log *logrus.Logger

// AperLog : Log entry of aper
var AperLog *logrus.Entry

func init() {
	log = logrus.New()
	log.SetReportCaller(false)

	log.Formatter = &formatter.Formatter{
		TimestampFormat: time.RFC3339,
		TrimMessages:    true,
		NoFieldsSpace:   true,
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
	}

	AperLog = log.WithFields(logrus.Fields{"component": "MOD", "category": "TEST"})
}

func main() {
	AperLog.Infof("Information !!")
}

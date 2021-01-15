package mymodule

import (
	"fmt"
	formatter "github.com/antonfisher/nested-logrus-formatter"
	"github.com/chenht1998/testA"
	"github.com/jordan0210/releaseGoModTest"
	"github.com/sirupsen/logrus"
	"time"
)

var log *logrus.Logger

var Log *logrus.Entry

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

	Log = log.WithFields(logrus.Fields{"component": "MOD", "category": "TEST"})
}

func Test() {
	myGoMod.Main()
	fmt.Println("------------------")
	testA.Exec()
	Log.Infof("Test !!")
}

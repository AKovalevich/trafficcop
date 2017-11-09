package trafficcop

import (
	"runtime"
	"os"

	log "github.com/AKovalevich/trafficcop/pkg/log/logrus"
	"github.com/AKovalevich/trafficcop/pkg/config"
	"github.com/AKovalevich/trafficcop/pkg/server"
	"github.com/containous/flaeg"
)

func Run(args []string) int {
	runtime.GOMAXPROCS(runtime.NumCPU())

	trafficcopConfiguration := config.NewTrafficcopConfiguration()
	trafficcopPointersConfiguration := config.NewTrafficcopDefaultConfiguration()

	trafficcopCmd := &flaeg.Command{
		Name:					"Trafficcop",
		Description:			`Traffic cop API Gateway`,
		Config:					trafficcopConfiguration,
		DefaultPointersConfig:	trafficcopPointersConfiguration,
		Run: func() error {
			trafficcopConfiguration.Reload()
			start(trafficcopConfiguration)
			return nil
		},
		Metadata: map[string]string{
			"parseAllSources": "true",
		},
	}

	healthCheckCmd := &flaeg.Command{
		Name:					"healthcheck",
		Description:			`Calls trafficcop /ping to check health`,
		Config:					struct{}{},
		DefaultPointersConfig:	struct{}{},
		Run: func() error {
			os.Exit(0)
			return nil
		},
		Metadata: map[string]string {
			"parseAllSources": "true",
		},
	}

	f := flaeg.New(trafficcopCmd, args)
	f.AddCommand(healthCheckCmd)
	f.AddCommand(newVersionCmd())

	// Add custom parsers to flaeg
	parserType, parser := config.CustomEntryPointParsers()
	f.AddParser(parserType, parser)

	if err := f.Run(); err != nil {
		log.Do.Error("Running error: ", err.Error())
	}

	return 1
}

// Start trafficcop application
func start(config *config.TrafficcopConfiguration) {
	log.Do.Infof("trafficcop started")
	log.Do.Debugf("PID: %d\n", os.Getpid())
	s := server.NewServer(config)
	s.Serve()
}

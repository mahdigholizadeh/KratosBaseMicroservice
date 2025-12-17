package main

import (
	"flag"
	"os"

	"KratosBaseMicroservice/internal/conf"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	_ "go.uber.org/automaxprocs"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	kratoszap "github.com/go-kratos/kratos/contrib/log/zap/v2"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
    // Define log level (you can make this configurable via env var later)
	logLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, gs *grpc.Server, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
			hs,
		),
	)
}

func main() {
	flag.Parse() // read flags and update global variables if needed Parses command-line flags and Reads -conf value into flagconf
	// Production config + set level
	zapConfig := zap.NewProductionConfig()
	zapConfig.Level = logLevel
	// Optional: make output more readable in development
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    zapConfig.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder // INFO, ERROR, etc.
    z, err := zapConfig.Build(zap.AddCallerSkip(1)) // AddCallerSkip helps with correct caller info
    if err != nil {
    panic(err)
    }
	// Create base logger (returns concrete type, but we store as interface)
	baseLogger := kratoszap.NewLogger(z)
	// Now use the interface type — this is the key fix!
	var logger log.Logger = baseLogger
	// Add your contextual fields (these will appear in every log entry)
	logger = log.With(logger,
		"ts", log.DefaultTimestamp,     // timestamp
		"caller", log.DefaultCaller,    // file:line
		"service.id", id,
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),  // OpenTelemetry trace ID
		"span.id", tracing.SpanID(),    // OpenTelemetry span ID
    ) 
	// Optional: set as global logger so log.Info(), log.Error() etc. work everywhere
    log.SetLogger(logger)
	// load configuration
	c := config.New( //Creates a configuration manager
		config.WithSource( //Reads config from files
			file.NewSource(flagconf), //flagconf is the config directory or file
		),
	)
	///
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

package rendergt

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

const (
	DefaultFileExt    = "tmpl"
	DefaultValuesFile = "values.yaml"
	DefaultOutputDir  = "stdout"
	DefaultDelims     = "{{}}"
)

var (
	globalConfig = &config{}
)

type Config struct {
	OutputDir string
	FileExt   string
	Values    []string
	InputDirs []string
	Delims    string
}

type config struct {
	outputDir  string
	fileExt    string
	values     []string
	inputDirs  []string
	fileChan   chan *string
	exit       chan os.Signal
	createdDir map[string]*struct{}
	leftDelim  string
	rightDelim string
}

func Run(conf *Config) error {
	globalConfig.outputDir = conf.OutputDir
	globalConfig.fileExt = conf.FileExt
	globalConfig.values = conf.Values
	globalConfig.inputDirs = conf.InputDirs
	globalConfig.fileChan = make(chan *string, 1)
	globalConfig.exit = make(chan os.Signal, 1)
	globalConfig.createdDir = map[string]*struct{}{}
	globalConfig.leftDelim = conf.Delims[:len(conf.Delims)/2]
	globalConfig.rightDelim = conf.Delims[len(conf.Delims)/2:]

	return runRenderer()
}

func runRenderer() error {
	read := make(chan struct{}, 1)
	done := make(chan struct{}, 1)
	signal.Notify(globalConfig.exit, os.Interrupt, os.Kill)
	g, ctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	g.Go(func() error {
		defer close(read)
		return getFileList(ctx)
	})

	g.Go(func() error {
		defer close(done)
		return renderTemplates(ctx)
	})

	for {
		select {
		case <-read:
			<-done
			return g.Wait()
		case <-done:
			log.Debugf("Aborting...")
			cancel()
			return g.Wait()
		case <-globalConfig.exit:
			log.Info("Exit signal detected...Closing...")
			cancel()
			select {
			case <-done:
				return fmt.Errorf("Exit signal detected...Closing")
			}
		}
	}
}

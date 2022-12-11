package rendergt

import (
	"context"
	"fmt"
	"io/fs"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func getFileList(ctx context.Context) error {
	done := make(chan error, 1)
	go func() {
		defer close(globalConfig.fileChan)
		defer close(done)
		for _, file := range globalConfig.inputDirs {
			err := filepath.WalkDir(file, func(path string, info fs.DirEntry, err error) error {
				if err != nil {
					return fmt.Errorf("Adding template file %s: %v", path, err)
				}
				if !info.IsDir() {
					if filepath.Ext(path) == globalConfig.fileExt {
						log.Debugf("Adding template file %s", path)
						globalConfig.fileChan <- &path
					}
				}
				return nil
			})
			if err != nil {
				done <- fmt.Errorf("Adding template files: %v", err)
			}
		}
	}()
	for {
		select {
		case err := <-done:
			return err
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

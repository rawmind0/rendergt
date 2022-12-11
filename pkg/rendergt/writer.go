package rendergt

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func getValues() (map[string]interface{}, error) {
	valuesPrefix := "Values"
	values := map[string]interface{}{}
	for _, file := range globalConfig.values {
		log.Debugf("Reading values: %s", file)
		if len(file) == 0 {
			continue
		}
		readed, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("Reading values: %v", err)
		}
		new := map[string]interface{}{}
		err = yaml.Unmarshal(readed, &new)
		if err != nil {
			return nil, fmt.Errorf("Parsing values: %v", err)
		}
		mergeMaps(values, new)
	}
	return map[string]interface{}{valuesPrefix: values}, nil
}

func renderTemplates(ctx context.Context) error {
	values, err := getValues()
	if err != nil {
		return fmt.Errorf("Getting values: %v", err)
	}
	for {
		select {
		case in := <-globalConfig.fileChan:
			if in == nil {
				return nil
			}
			err := renderTemplate(*in, values)
			if err != nil {
				return fmt.Errorf("Rendering templates %s: %v", *in, err)
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func renderTemplate(path string, values map[string]interface{}) error {
	t, err := template.ParseGlob(path)
	if err != nil {
		return fmt.Errorf("parsing template: %v", err)
	}
	output := &bytes.Buffer{}
	log.Debugf("Rendering template %s", path)
	err = t.Execute(output, values)
	if err != nil {
		return fmt.Errorf("rendering template: %v", err)
	}
	tmplName := t.Name()
	outFile := filepath.Join(globalConfig.outputDir, filepath.Dir(path), tmplName[:len(tmplName)-len(filepath.Ext(tmplName))])
	err = writeOutput(outFile, output)
	if err != nil {
		return fmt.Errorf("writing ouput: %v", err)
	}
	return nil
}

func writeOutput(path string, output *bytes.Buffer) error {
	outDir := filepath.Dir(path)
	if globalConfig.outputDir == DefaultOutputDir {
		os.Stdout.Write(output.Bytes())
		return nil
	}
	if globalConfig.createdDir[outDir] == nil {
		globalConfig.createdDir[outDir] = &struct{}{}
		log.Debugf("Reading dir %s", outDir)
		if _, err := os.ReadDir(outDir); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return fmt.Errorf("Reading dir %s: %v", outDir, err)
			}
			log.Debugf("Creating dir %s", outDir)
			err = os.MkdirAll(outDir, 0755)
			if err != nil {
				return fmt.Errorf("Creating dir %s: %v", outDir, err)
			}
		}
	}
	log.Debugf("Writing file %s", path)
	err := os.WriteFile(path, output.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("Writing file: %v", err)
	}
	return nil
}
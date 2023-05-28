package helm

import (
	"os"
	"path/filepath"
)

// FindCharts locates all "Chart.yaml" files in the current directory, and all sub-directories
func FindCharts(chartSearchDir, chartDirectory string) ([]string, error) {
	fileList := []string{}
	var err error
	if chartDirectory != "" {
		err = filepath.Walk(chartDirectory, func(path string, f os.FileInfo, err error) error {
			fileName := filepath.Base(path)
			if fileName == "Chart.yaml" {
				fileList = append(fileList, path)
			}
			return nil
		})
	} else {
		err = filepath.Walk(chartSearchDir, func(path string, f os.FileInfo, err error) error {
			fileName := filepath.Base(path)
			if fileName == "Chart.yaml" {
				fileList = append(fileList, path)
			}
			return nil
		})
	}

	return fileList, err
}

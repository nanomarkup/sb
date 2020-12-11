package golang

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/sapplications/sbuilder/src/smod"
)

func check(configuration string, config *smod.ConfigFile) (string, error) {
	// read the main item
	var main = config.Items["main"]
	if main == nil {
		return "", fmt.Errorf("The main item is not found")
	}
	// read the current configuration if it is not specified
	if configuration == "" {
		if _, err := os.Stat(configFileName); err == nil {
			configuration, _ = readConfiguration(configFileName)
		} else if os.IsNotExist(err) {
			// check the number of configurations
			// if it is 1 then select it
			if len(main) != 1 {
				return "", fmt.Errorf("The configuration is not specified")
			}
			// select the existing configuration
			for key := range main {
				configuration = key
				break
			}
		} else {
			return "", err
		}
	}
	// check the configuration is exist
	if _, found := main[configuration]; !found {
		return "", fmt.Errorf("The selected \"%s\" configuration is not found", configuration)
	}
	return configuration, nil
}

func goBuild(src, dst string) error {
	args := []string{"build"}
	if dst != "" {
		args = append(args, "-o", dst)
	}
	args = append(args, src)
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func goClean(src string) error {
	cmd := exec.Command("go", "clean", src)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func readConfiguration(filePath string) (string, error) {
	// read the generated configuration golang file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	constConfig := "const Configuration ="
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return "", err
		}
		line = strings.Trim(line, "\t \n \r")
		if strings.HasPrefix(line, constConfig) {
			// read the configuration
			return strings.Trim(strings.Replace(strings.Replace(line, constConfig, "", 1), "\"", "", 2), " "), nil
		}
	}
	return "", fmt.Errorf("Failed to read the configuration from \"%s\" file", filePath)
}

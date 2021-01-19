package golang

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func checkApplication(application string) error {
	if application == "" {
		return fmt.Errorf("The application is not specified")
	}
	return nil
}

func readMain(items map[string]map[string]string) (map[string]string, error) {
	if main, found := items["main"]; found {
		return main, nil
	}
	return nil, fmt.Errorf("The main item is not found")
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

func isDirEmpty(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true, nil
	}
	return false, nil
}

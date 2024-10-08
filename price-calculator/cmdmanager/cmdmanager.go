package cmdmanager

import "fmt"

type CMDManager struct {
}

func (cmd CMDManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter the prices values")
	var lines []string

	for {
		var line string
		fmt.Scanln(&line)
		if line == "" {
			break
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func (cmd CMDManager) WriteResult(data any) error {
	fmt.Println(data)
	return nil
}

func New() CMDManager {
	return CMDManager{}
}

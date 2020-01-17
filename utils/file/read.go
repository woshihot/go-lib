package file

import (
	"bufio"
	"fmt"
	"github.com/woshihot/go-lib/utils/log"
	"io"
	"os"
	"strings"
)

func ReadFileLine(fileName string) []string {
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open file error!", err)
		return nil
	}
	defer file.Close()

	var result []string
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		result = append(result, line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.EF(TAG_ERROR, "Read file error -> %s\n", err.Error())
				return nil
			}
		}
	}
	return result
}

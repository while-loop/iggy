package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sync"
)

func GetIgnores(client *http.Client, urls []string) ([]string, error) {
	lines := make([]string, 0)
	mu := sync.Mutex{}
	var wg sync.WaitGroup

	for _, u := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()

			resp, err := client.Get(u)
			if err != nil || resp.StatusCode != http.StatusOK {
				fmt.Fprintf(os.Stderr, "Unable to download ignore file %s: %v", u, err)
				return
			}
			defer resp.Body.Close()

			buf := bufio.NewReader(resp.Body)
			mu.Lock()
			defer mu.Unlock()
			lines = append(lines, "", "## "+path.Base(u)+" ##")
			for buf.Size() > 0 {
				line, _, err := buf.ReadLine()
				if err == io.EOF {
					break
				} else if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to read from gitignore file %s: %v", path.Base(u), err)
					break
				}
				lines = append(lines, string(line))
			}
		}(u)
	}

	wg.Wait()
	return lines, nil
}

func AddIgnores(w io.Writer, lines []string) {
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
}

package main

import (
	"bufio"
	"io"
	"os"
	"path"
	"sync"
	"time"
)

var seenFile *os.File
var seen = make(map[string]struct{})
var mutex = sync.RWMutex{}

// TODO: Make this a module?

func init() {
	var err error

	seenFile, err = os.OpenFile(path.Join(*savePath, ".md5"), os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(seenFile)

	for scanner.Scan() {
		seen[scanner.Text()] = struct{}{}
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)

			_, err := seenFile.Seek(0, io.SeekStart)
			if err != nil {
				panic(err)
			}

			var keys string
			mutex.RLock()
			for k := range seen {
				keys += k + "\n"
			}
			mutex.RUnlock()

			_, err = seenFile.WriteString(keys)
			if err != nil {
				panic(err)
			}
		}
	}()
}

func SeenMD5(str string) bool {
	mutex.RLock()
	_, exists := seen[str]
	mutex.RUnlock()
	return exists
}

func AddMD5(str string) {
	mutex.Lock()
	seen[str] = struct{}{}
	mutex.Unlock()
}

// Final Commit Done

package filelistener

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

var (
	myChannel = make(chan bool, 1)
	wg        sync.WaitGroup
)

func doSomething(filePath string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating watcher:", err)
		return
	}
	defer watcher.Close()

	err = watcher.Add(filePath)
	if err != nil {
		fmt.Println("Error adding file to watcher:", err)
		return
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Printf("File %s was modified\n", event.Name)
			}
			if event.Op&fsnotify.Create == fsnotify.Create {
				fmt.Printf("File %s was created\n", event.Name)
			}
			if event.Op&fsnotify.Remove == fsnotify.Remove {
				fmt.Printf("File %s was deleted\n", event.Name)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("Error watching file:", err)
		}
	}
}

func myServiceCode(filePath string) {
	defer wg.Done()
	for {
		select {
		case <-myChannel:
			fmt.Println("Signal Came, Stopping Service")
			return
		default:
			doSomething(filePath)
			time.Sleep(time.Second * 2)
		}
	}
}

func ListenToFile() {
	fmt.Println("File Listener Started")
	file, err := os.Open("23_File_Listener/file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	wg.Add(1)
	go myServiceCode("23_File_Listener/file.txt")
	wg.Wait()
}

package robot

import (
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)

var paths []string

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func Listen(path string, exec func(message XCAutoLog)) {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.Add(path)
	if err != nil {
		log.Fatal(err)
	}

	sc := make(chan string)

	go func() {
		s := path + "\\" + time.Now().Format("2006-01-02")
		log.Println(s)
		if _, err2 := os.Stat(s); err2 == nil {
			sc <- s
		}
		for {
			select {

			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Create) && !contains(paths, event.Name) {
					paths = append(paths, event.Name)
					sc <- event.Name
					log.Println(paths)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	go func(sc chan string, exec func(message XCAutoLog)) {
		_watcher, _err := fsnotify.NewWatcher()
		if _err != nil {
			log.Fatal(_err)
		}
		go func() {
			for {
				select {
				case event, ok := <-_watcher.Events:
					if !ok {
						return
					}
					if event.Op&fsnotify.Write == fsnotify.Write {
						exec(Parse(event.Name))
					}
				case _err, ok := <-_watcher.Errors:
					if !ok {
						return
					}
					log.Println("error:", _err)
				}
			}
		}()
		for s := range sc {
			_err = _watcher.Add(s)
			if _err != nil {
				log.Fatal(_err)
			}
		}

	}(sc, exec)
}

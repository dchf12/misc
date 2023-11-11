package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer w.Close()

	go func() {
		for {
			select {
			case event, ok := <-w.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-w.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	if err := w.Add("./"); err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}

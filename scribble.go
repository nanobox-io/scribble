// Package scribble is a tiny JSON database
package scribble

import (
	"os"
	"sync"
)

// Version is the current version of the project
const Version = "1.1"

// stat checks for dir, if path isn't a directory check to see if it's a file
func stat(path string) (fi os.FileInfo, err error) {
	if fi, err = os.Stat(path); os.IsNotExist(err) {
		fi, err = os.Stat(path + ".json")
	}
	return
}

// getOrCreateMutex creates a new collection specific mutex any time a collection
// is being modfied to avoid unsafe operations
func (d *Driver) getOrCreateMutex(collection string) *sync.Mutex {
	// create mutex
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// if the mutex doesn't exist make it
	m, ok := d.mutexes[collection]
	if !ok {
		m = &sync.Mutex{}
		d.mutexes[collection] = m
	}

	return m
}

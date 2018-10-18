// Package scribble is a tiny JSON database
package scribble

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/jcelliott/lumber"
)

// Version is the current version of the project
const Version = "1.1"

// New creates a new scribble database at the desired directory location, and
// returns a *Driver to then use for interacting with the database
func New(dir string, options *Options) (*Driver, error) {

	//
	dir = filepath.Clean(dir)

	// create default options
	opts := Options{}

	// if options are passed in, use those
	if options != nil {
		opts = *options
	}

	// if no logger is provided, create a default
	if opts.Logger == nil {
		opts.Logger = lumber.NewConsoleLogger(lumber.INFO)
	}

	//
	driver := Driver{
		dir:     dir,
		mutexes: make(map[string]*sync.Mutex),
		log:     opts.Logger,
	}

	// if the database already exists, just use it
	if _, err := os.Stat(dir); err == nil {
		opts.Logger.Debug("Using '%s' (database already exists)\n", dir)
		return &driver, nil
	}

	// if the database doesn't exist create it
	opts.Logger.Debug("Creating scribble database at '%s'...\n", dir)
	return &driver, os.MkdirAll(dir, 0755)
}

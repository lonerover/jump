package cmd

import (
	"os"
	"path/filepath"

	"github.com/gsamokovarov/jump/cli"
	"github.com/gsamokovarov/jump/config"
)

func forgetCmd(args cli.Args, conf config.Config) error {
	dir, err := os.Getwd()
	if len(args) == 0 && err != nil {
		return err
	} else {
		if dir, err = filepath.Abs(args.CommandName()); err != nil {
			return err
		}
	}

	entries, err := conf.ReadEntries()
	if err != nil {
		return err
	}

	if entry, found := entries.Find(dir); found {
		cli.Outf("Cleaning %s\n", entry.Path)
		entries.Remove(entry.Path)

		if err := conf.WriteEntries(entries); err != nil {
			return err
		}
	}

	return nil
}

func init() {
	cli.RegisterCommand("forget", "Removes the current directory from the database.", forgetCmd)
}

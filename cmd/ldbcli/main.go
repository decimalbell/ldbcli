package main

import (
	"fmt"
	"log"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
	"github.com/urfave/cli/v2"
)

func main() {
	var (
		dbpath string
		prefix string
	)

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "dbpath",
				Usage:       "dbpath",
				Destination: &dbpath,
				Required:    true,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "get",
				Usage: "get",
				Action: func(cCtx *cli.Context) error {
					opts := &opt.Options{
						ReadOnly: true,
					}
					db, err := leveldb.OpenFile(dbpath, opts)
					if err != nil {
						return err
					}

					for _, key := range cCtx.Args().Slice() {
						value, err := db.Get([]byte(key), nil)
						if err != nil {
							continue
						}
						fmt.Printf("%s %s\n", key, string(value))
					}
					return nil
				},
			},
			{
				Name:  "ls",
				Usage: "ls",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "prefix",
						Usage:       "prefix",
						Destination: &prefix,
					},
				},
				Action: func(cCtx *cli.Context) error {
					opts := &opt.Options{
						ReadOnly: true,
					}
					db, err := leveldb.OpenFile(dbpath, opts)
					if err != nil {
						return err
					}

					var slice *util.Range
					if prefix != "" {
						slice = util.BytesPrefix([]byte(prefix))
					}

					iter := db.NewIterator(slice, nil)
					for iter.Next() {
						key := iter.Key()
						value := iter.Value()
						fmt.Printf("%s %s\n", key, string(value))
					}
					iter.Release()
					return iter.Error()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

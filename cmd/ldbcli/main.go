package main

import (
	"fmt"
	"log"
	"os"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/urfave/cli/v2"
)

func main() {
	var dbpath string

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
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

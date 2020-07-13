package main

import (
	"fmt"
	"context"
	

	"github.com/urfave/cli/v2"
	"cloud.google.com/go/storage"
)

func GcpEntrypoint(c *cli.Context) {
	path := c.Args().Get(0)
	bucketName := c.Args().Get(1)
	projectID := c.Args().Get(2)
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		PrintError(err.Error(), 1)
	}
	defer client.Close()
	bkt := client.Bucket(bucketName)
	if err := bkt.Create(ctx, projectID, nil); err != nil {
		PrintWarning("Bucket already exists")
	}

	if path == "" {
		PrintError("Must specify a path.", 64)
	}
	for _, contents := range GetFilesInDir(path) {
		fmt.Println(contents)
	}
}

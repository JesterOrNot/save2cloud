package main

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/storage"
	"github.com/urfave/cli/v2"
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
		PrintError(err.Error(), 1)
	}

	if path == "" {
		PrintError("Must specify a path.", 64)
	}
	contents, paths := GetFilesInDir(path)
	numOfFiles := len(paths)
	var wg sync.WaitGroup
	wg.Add(numOfFiles)
	for i, content := range contents {
		go uploadFile(ctx, bkt, paths[i], content, &wg)
	}
	wg.Wait()
}

func uploadFile(ctx context.Context, bkt *storage.BucketHandle, path, content string, wg *sync.WaitGroup) {
	PrintInfo(fmt.Sprint("Uploading: ", path))
	obj := bkt.Object(path)
	writer := obj.NewWriter(ctx)
	defer writer.Close()
	defer wg.Done()
	fmt.Fprintf(writer, content)
}

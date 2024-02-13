package main

import (
	"context"
	"github.com/docker/docker/api/types/container"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/elasticsearch"
	"log"
	"os"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	elasticsearchContainer, err := elasticsearch.RunContainer(
		ctx,
		testcontainers.WithImage("docker.elastic.co/elasticsearch/elasticsearch:8.12.1"),
		testcontainers.WithHostConfigModifier(func(hostConfig *container.HostConfig) {
			hostConfig.Binds = append(hostConfig.Binds, cwd+"/data:/usr/share/elasticsearch/data")
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := elasticsearchContainer.Terminate(ctx); err != nil {
			log.Printf("failed to terminate elasticsearch container: %v\n", err)
		}
	}()
	print(elasticsearchContainer.Settings.Address)
}

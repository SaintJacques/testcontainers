package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const port = "5432/tcp"

var (
	errEmptyDBName = errors.New("database name must not be empty")
	env            = map[string]string{
		"POSTGRES_USER":     "postgres",
		"POSTGRES_PASSWORD": "password",
	}
)

func NewTestDB(ctx context.Context, dbName string) (string, error) {
	if len(dbName) == 0 {
		return "", errEmptyDBName
	}

	env["POSTGRES_DB"] = dbName
	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "postgres",
			ExposedPorts: []string{port},
			Cmd:          []string{"postgres", "-c", "fsync=off"},
			Env:          env,
			Privileged:   true,
			SkipReaper:   true,
			WaitingFor:   wait.ForSQL(nat.Port(port), "postgres", buildDSN).Timeout(time.Second * 10),
		},
		Started: true,
	}

	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return "", err
	}

	mappedPorts, err := container.MappedPort(ctx, nat.Port(port))
	if err != nil {
		return "", err
	}

	return buildDSN(mappedPorts), nil
}

func buildDSN(port nat.Port) string {
	return fmt.Sprintf("postgres://%s:%s@localhost:%s/%s?sslmode=disable", env["POSTGRES_USER"], env["POSTGRES_PASSWORD"], port.Port(), env["POSTGRES_DB"])
}

package mongorepo_test

import (
	"context"
	"fmt"
	"testing"
	"xm-challenge/internal/domain"
	"xm-challenge/internal/infra/mongorepo"

	"github.com/d5/tengo/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	mongoContainer testcontainers.Container
	companyRepo    *mongorepo.CompanyRepo
)

func setupContainer(t *testing.T) (func(), error) {
	ctx := context.Background()

	// Define the Redis container configuration
	req := testcontainers.ContainerRequest{
		Image:        "mongo:6",
		ExposedPorts: []string{"27017"},
		WaitingFor:   wait.ForLog("waiting for connections on port"),
	}

	// Create and start the Redis container
	mongoC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start Redis container: %w", err)
	}

	// Get the Redis container's host and port
	mongoHost, err := mongoC.Host(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get Redis container host: %w", err)
	}
	mongoPort, err := mongoC.MappedPort(ctx, "27017")
	if err != nil {
		return nil, fmt.Errorf("failed to get Redis container port: %w", err)
	}

	// Create the RedisPortRepository with the container's host and port
	mongoURL := fmt.Sprintf("mongodb://%s:%s/0", mongoHost, mongoPort.Port())
	repo, err := mongorepo.NewRepo(mongoURL)
	if err != nil {
		return nil, fmt.Errorf("failed to create Redis repository: %w", err)
	}

	// Set the global variables for the container and repository
	mongoContainer = mongoC
	companyRepo = repo

	// Create a cleanup function to terminate the container after the tests
	cleanup := func() {
		err := mongoContainer.Terminate(ctx)
		if err != nil {
			t.Errorf("failed to terminate Redis container: %v", err)
		}
	}

	return cleanup, nil
}

func TestMongoRepo(t *testing.T) {
	clean, err := setupContainer(t)
	defer clean()
	ctx := context.Background()
	company := domain.Company{
		ID:          "AEAUH",
		Name:        "Aadsfasdfasdfasdfasdfas",
		Description: "asfdasdf",
		Employees:   5,
		Registered:  true,
		CompanyType: "coorporation",
	}
	_, err = companyRepo.CreateOrganization(ctx, company)
	if err != nil {
		fmt.Printf(err.Error())
	}
	assert.NoError(t, err)

	_, p, err := companyRepo.GetOrgByName(ctx, company.Name)
	fmt.Println(p[0].Name)

}

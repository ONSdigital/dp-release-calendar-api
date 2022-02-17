package api

import (
	"context"

	"github.com/ONSdigital/dp-api-clients-go/v2/zebedee"
	"github.com/ONSdigital/dp-healthcheck/healthcheck"
)

// To mock interfaces in this file
//go:generate moq -out mock_clients.go -pkg api . ZebedeeClient

// ZebedeeClient defines the required methods to talk to Zebedee
type ZebedeeClient interface {
	GetRelease(ctx context.Context, userAccessToken, collectionID, lang, uri string) (zebedee.Release, error)
	Checker(ctx context.Context, check *healthcheck.CheckState) error
}

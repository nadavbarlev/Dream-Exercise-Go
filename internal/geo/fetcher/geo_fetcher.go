package fetcher

import (
	"github.com/nadavbarlev/dream-exercise-go/internal/models"
)

type GeoFetcher interface {
	FetchGeoLocation(request models.LocationRequest) (*models.LocationResponse, error)
}

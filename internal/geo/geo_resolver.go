package geo

import (
	"github.com/nadavbarlev/dream-exercise-go/internal/geo/fetcher"
	"github.com/nadavbarlev/dream-exercise-go/internal/models"
)

type GeoResolver struct {
	fetcher fetcher.GeoFetcher
	cache   map[string]*models.LocationResponse
}

func NewGeoResolver(geoFetcher fetcher.GeoFetcher) *GeoResolver {
	return &GeoResolver{
		fetcher: geoFetcher,
		cache: make(map[string]*models.LocationResponse),
	}
}

func (resolver *GeoResolver) Resolve(request models.LocationRequest) (*models.LocationResponse, error) {

	if resolver.cache[request.IP] != nil {
		return resolver.cache[request.IP], nil
	}
	
	res, err := resolver.fetcher.FetchGeoLocation(request)
	if err != nil {
		return nil, err
	}

	resolver.cache[request.IP] = res
	return res, nil
}


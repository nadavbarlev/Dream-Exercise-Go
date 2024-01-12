package fetcher

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nadavbarlev/dream-exercise-go/internal/models"
)


type IpApiGeoFetcher struct {
	baseURL string
}

func NewIpApiGeoFetcher() *IpApiGeoFetcher {
	return &IpApiGeoFetcher{
		baseURL: "http://ip-api.com/json",
	}
}

func (fetcher *IpApiGeoFetcher) FetchGeoLocation(request models.LocationRequest) (*models.LocationResponse, error) {

	// Perform API call to the external service
    apiUrl := fmt.Sprintf("%s/%s", fetcher.baseURL, request.IP)
	res, err := http.Get(apiUrl)
    if err != nil {
        return nil, err
    }

	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println(body)

	// Use a map[string]interface{} to capture any extra fields in the JSON response
	var rawData map[string]interface{}
	err = json.Unmarshal(body, &rawData)
	if err != nil {
		return nil, err
	}

	fmt.Println(rawData)

	// Decode the response body
	var locationResponse models.LocationResponse
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return nil, err
	}

	fmt.Println(locationResponse)

	return &locationResponse, nil
}

package internal

import (
	"fmt"
	"math"
)

func ST_GEOGPOINT(longitude float64, latitude float64) (Value, error) {
	if latitude < -90 || latitude > 90 {
		return nil, fmt.Errorf("ST_GEOGPOINT: invalid latitude: %f", latitude)
	}

	return &GeographyValue{
		g: &geographyPoint{
			longitude: geographyNormalizeLongitude(longitude),
			latitude:  latitude,
		},
	}, nil
}

func geographyNormalizeLongitude(longitude float64) float64 {
	longitude = math.Mod(longitude, 360)

	if longitude > 180 {
		return math.Mod(longitude, 180) - 180
	}
	if longitude < -180 {
		return 180 - math.Mod(longitude, 180)
	}

	return longitude
}

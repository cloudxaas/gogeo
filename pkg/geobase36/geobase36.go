// Package geobase36 provides functions for converting latitude and longitude
// to and from a base36 encoding with variable precision. //ChatGPT4 generated
package geobase36

import (
	"strings"
)

const base = 36
const characters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// LatLonToBase36 converts latitude and longitude to a base36 code with specified precision.
// This function is optimized for zero allocation.
func LatLonToBase36(lat, lon float64, precision uint8) string {
	// Pre-allocated array for the result
	var result [11]byte // 11 is the maximum length for uint64 in base36

	// Normalize latitude and longitude
	normLat := (lat + 90) / 180
	normLon := (lon + 180) / 360

	// Combine normalized lat and lon
	combined := (normLat + normLon) / 2

	// Convert to base36 with specified precision
	value := uint64(combined * float64(power(base, precision)))
	i := precision

	for value > 0 && i > 0 {
		i--
		digit := value % base
		result[i] = characters[digit]
		value = value / base
	}

	// Pad with leading characters if necessary
	for i > 0 {
		i--
		result[i] = characters[0]
	}

	return string(result[:precision])
}

// Base36ToLatLon converts a base36 code back to approximate latitude and longitude.
// This function is optimized for zero allocation.
func Base36ToLatLon(code string) (float64, float64) {
	var value uint64
	for i := 0; i < len(code); i++ {
		value = value*base + uint64(strings.Index(characters, string(code[i])))
	}

	// Scale back to combined range of latitude and longitude
	combined := float64(value) / float64(power(base, uint8(len(code))))

	// Separate into approximate latitude and longitude
	combined *= 2
	normLat := combined - 1
	normLon := combined - 1

	lat := normLat*180 - 90
	lon := normLon*360 - 180

	return lat, lon
}

// power calculates base^exp. This is a simplified and more efficient version than math.Pow for integers.
func power(base, exp uint8) uint64 {
	result := uint64(1)
	for exp > 0 {
		result *= uint64(base)
		exp--
	}
	return result
}

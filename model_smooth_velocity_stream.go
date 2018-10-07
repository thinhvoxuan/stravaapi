/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stravaapi

type SmoothVelocityStream struct {
	// The number of data points in this stream
	OriginalSize int32 `json:"original_size,omitempty"`
	// The level of detail (sampling) in which this stream was returned
	Resolution string `json:"resolution,omitempty"`
	// The base series used in the case the stream was downsampled
	SeriesType string `json:"series_type,omitempty"`
	// The sequence of velocity values for this stream, in meters per second
	Data []float32 `json:"data,omitempty"`
}

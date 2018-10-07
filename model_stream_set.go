/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stravaapi

type StreamSet struct {
	Time *TimeStream `json:"time,omitempty"`
	Distance *DistanceStream `json:"distance,omitempty"`
	Latlng *LatLngStream `json:"latlng,omitempty"`
	Altitude *AltitudeStream `json:"altitude,omitempty"`
	VelocitySmooth *SmoothVelocityStream `json:"velocity_smooth,omitempty"`
	Heartrate *HeartrateStream `json:"heartrate,omitempty"`
	Cadence *CadenceStream `json:"cadence,omitempty"`
	Watts *PowerStream `json:"watts,omitempty"`
	Temp *TemperatureStream `json:"temp,omitempty"`
	Moving *MovingStream `json:"moving,omitempty"`
	GradeSmooth *SmoothGradeStream `json:"grade_smooth,omitempty"`
}

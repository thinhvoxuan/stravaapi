/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stravaapi

import (
	"time"
)

type SummaryActivity struct {
	// The unique identifier of the activity
	Id int64 `json:"id,omitempty"`
	// The identifier provided at upload time
	ExternalId string `json:"external_id,omitempty"`
	// The identifier of the upload that resulted in this activity
	UploadId int64 `json:"upload_id,omitempty"`
	Athlete *SummaryAthlete `json:"athlete,omitempty"`
	// The name of the activity
	Name string `json:"name,omitempty"`
	// The activity's distance, in meters
	Distance float32 `json:"distance,omitempty"`
	// The activity's moving time, in seconds
	MovingTime int32 `json:"moving_time,omitempty"`
	// The activity's elapsed time, in seconds
	ElapsedTime int32 `json:"elapsed_time,omitempty"`
	// The activity's total elevation gain.
	TotalElevationGain float32 `json:"total_elevation_gain,omitempty"`
	// The activity's highest elevation, in meters
	ElevHigh float32 `json:"elev_high,omitempty"`
	// The activity's lowest elevation, in meters
	ElevLow float32 `json:"elev_low,omitempty"`
	Type_ *ActivityType `json:"type,omitempty"`
	// The time at which the activity was started.
	StartDate time.Time `json:"start_date,omitempty"`
	// The time at which the activity was started in the local timezone.
	StartDateLocal time.Time `json:"start_date_local,omitempty"`
	// The timezone of the activity
	Timezone string `json:"timezone,omitempty"`
	StartLatlng *LatLng `json:"start_latlng,omitempty"`
	EndLatlng *LatLng `json:"end_latlng,omitempty"`
	// The number of achievements gained during this activity
	AchievementCount int32 `json:"achievement_count,omitempty"`
	// The number of kudos given for this activity
	KudosCount int32 `json:"kudos_count,omitempty"`
	// The number of comments for this activity
	CommentCount int32 `json:"comment_count,omitempty"`
	// The number of athletes for taking part in a group activity
	AthleteCount int32 `json:"athlete_count,omitempty"`
	// The number of Instagram photos for this activity
	PhotoCount int32 `json:"photo_count,omitempty"`
	// The number of Instagram and Strava photos for this activity
	TotalPhotoCount int32 `json:"total_photo_count,omitempty"`
	Map_ *PolylineMap `json:"map,omitempty"`
	// Whether this activity was recorded on a training machine
	Trainer bool `json:"trainer,omitempty"`
	// Whether this activity is a commute
	Commute bool `json:"commute,omitempty"`
	// Whether this activity was created manually
	Manual bool `json:"manual,omitempty"`
	// Whether this activity is private
	Private bool `json:"private,omitempty"`
	// Whether this activity is flagged
	Flagged bool `json:"flagged,omitempty"`
	// The activity's workout type
	WorkoutType int32 `json:"workout_type,omitempty"`
	// The activity's average speed, in meters per second
	AverageSpeed float32 `json:"average_speed,omitempty"`
	// The activity's max speed, in meters per second
	MaxSpeed float32 `json:"max_speed,omitempty"`
	// Whether the logged-in athlete has kudoed this activity
	HasKudoed bool `json:"has_kudoed,omitempty"`
	// The id of the gear for the activity
	GearId string `json:"gear_id,omitempty"`
	// The total work done in kilojoules during this activity. Rides only
	Kilojoules float32 `json:"kilojoules,omitempty"`
	// Average power output in watts during this activity. Rides only
	AverageWatts float32 `json:"average_watts,omitempty"`
	// Whether the watts are from a power meter, false if estimated
	DeviceWatts bool `json:"device_watts,omitempty"`
	// Rides with power meter data only
	MaxWatts int32 `json:"max_watts,omitempty"`
	// Similar to Normalized Power. Rides with power meter data only
	WeightedAverageWatts int32 `json:"weighted_average_watts,omitempty"`
}

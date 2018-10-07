/*
 * Strava API v3
 *
 * Strava API
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package stravaapi

type ExplorerSegment struct {
	// The unique identifier of this segment
	Id int64 `json:"id,omitempty"`
	// The name of this segment
	Name string `json:"name,omitempty"`
	// The category of the climb
	ClimbCategory int32 `json:"climb_category,omitempty"`
	// The description for the category of the climb
	ClimbCategoryDesc string `json:"climb_category_desc,omitempty"`
	// The segment's average grade, in percents
	AvgGrade float32 `json:"avg_grade,omitempty"`
	StartLatlng *LatLng `json:"start_latlng,omitempty"`
	EndLatlng *LatLng `json:"end_latlng,omitempty"`
	// The segments's evelation difference, in meters
	ElevDifference float32 `json:"elev_difference,omitempty"`
	// The segment's distance, in meters
	Distance float32 `json:"distance,omitempty"`
	// The polyline of the segment
	Points string `json:"points,omitempty"`
}

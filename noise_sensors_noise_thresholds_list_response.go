// This file was auto-generated by Fern from our API Definition.

package api

type NoiseSensorsNoiseThresholdsListResponse struct {
	NoiseThresholds []*NoiseThreshold `json:"noise_thresholds,omitempty"`
	Ok              bool              `json:"ok"`
}

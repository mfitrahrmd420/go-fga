package message

import "time"

// annotation -> OMITEMPTY -> menghapus json property yang
// value null, atau kosong
type Response struct {
	Status    string     `json:"status"`
	Message   string     `json:"message,omitempty"`    // nullable
	Data      any        `json:"data,omitempty"`       // nullable
	StartTime *time.Time `json:"start_time,omitempty"` // nullable
}

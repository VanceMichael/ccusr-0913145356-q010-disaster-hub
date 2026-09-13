package domain

import "time"

// RecordID 是对外暴露的稳定业务标识。
type RecordID string

// EvidenceRef 指向一条不可变的业务证据。
type EvidenceRef struct {
	ID        RecordID  `json:"id"`
	Digest    string    `json:"digest"`
	CreatedAt time.Time `json:"created_at"`
}

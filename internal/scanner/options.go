package scanner

import "time"

type ScanOptions struct {
	Timeout time.Duration
	Host    string
}

package scanner

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type PortScanner struct {
	o ScanOptions
}

func NewPortScanner(o *ScanOptions) *PortScanner {
	return &PortScanner{
		o: *o,
	}
}

func (ps *PortScanner) ScanPort(port int) bool {
	target := fmt.Sprintf("%s:%d", ps.o.Host, port)

	conn, err := net.DialTimeout("tcp", target, ps.o.Timeout)
	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			time.Sleep(ps.o.Timeout)
			return ps.ScanPort(port)
		}
		return false
	}

	conn.Close()
	return true
}

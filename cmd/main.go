package main

import (
	"context"
	"errors"
	"flag"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"example.com/m/internal/presenter"
	"example.com/m/internal/scanner"
	"golang.org/x/sync/semaphore"
)

func getPortsRange(ports string) (int, int) {
	portsSlice := strings.Split(ports, "-")
	rangeStart, _ := strconv.Atoi(portsSlice[0])
	rangeFinish, _ := strconv.Atoi(portsSlice[1])
	return rangeStart, rangeFinish
}

func parsePorts(ports string) (*int, *int, error) {
	regex := `\b(0|[1-9][0-9]{0,4}|[1-5][0-9]{5}):([0-5]?\d{0,4}|[1-5][0-9]{5})\b|\b(0|[1-9][0-9]{0,4}|[1-5][0-9]{5})\b`
	match, _ := regexp.MatchString(regex, ports)
	if strings.Contains(ports, "-") {
		s, f := getPortsRange(ports)
		if s < f && match {
			return &s, &f, nil
		}
		return nil, nil, errors.New("ports range is invalid")

	}

	singlePort, err := strconv.Atoi(ports)
	if err != nil {
		return nil, nil, errors.New("ports range is invalid")
	}
	return &singlePort, &singlePort, nil
}

func main() {
	host := flag.String("h", "0.0.0.0", "Target hostname (can be IPv4 or domain address), for ex. 'google.com'")
	timeoutMS := flag.Uint("t", 500, "Timeout duration (in milliseconds)")
	portsRange := flag.String("p", "0-65535", "Ports range or single port (from 0 to 65535), for ex. '80:100' or '443'")
	flag.Parse()

	start, finish, err := parsePorts(*portsRange)
	if err != nil {
		panic(err)
	}

	presenter.LogAtStart(*host, *timeoutMS, *portsRange)

	ps := scanner.NewPortScanner(
		&scanner.ScanOptions{
			Host:    *host,
			Timeout: time.Millisecond * time.Duration(*timeoutMS),
		},
	)

	lock := semaphore.NewWeighted(65535)
	wg := sync.WaitGroup{}
	defer wg.Wait()

	for p := *start; p <= *finish; p++ {

		lock.Acquire(context.TODO(), 1)
		wg.Add(1)

		go func(port int) {
			defer lock.Release(1)
			defer wg.Done()

			isOpen := ps.ScanPort(port)
			if isOpen {
				presenter.LogOpenedPort(*host, port)
			}
		}(p)

	}

}

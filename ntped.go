package ntped

import (
	"github.com/beevik/ntp"
	"time"
	"sort"
	"fmt"
	"sync"
)

type int64s []int64
func (is int64s) Len() int {
	return len(is)
}
func (is int64s) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}
func (is int64s) Less(i, j int) bool {
	return is[i] < is[j]
}

var	offsets int64s
var median  int64

// Returns NTP offset applied time.
// If NTP() is not called, returns system time.
func Now() time.Time {
 	return time.Now().Add(time.Duration(median))
}

func UnixMilli() int64 {
	return Now().UnixNano() / 1e6
}

func Sync(maxNtpRetry int, timeoutInMs int) error {
	// NTP pool
	ntpPool := [...]string{
		"0.asia.ntp.org",
		"1.asia.ntp.org",
		"2.asia.ntp.org",
		"3.asia.ntp.org",
		"0.europe.pool.ntp.org",
		"1.europe.pool.ntp.org",
		"2.europe.pool.ntp.org",
		"3.europe.pool.ntp.org",
		"0.north-america.pool.ntp.org",
		"1.north-america.pool.ntp.org",
		"2.north-america.pool.ntp.org",
		"3.north-america.pool.ntp.org",
		"0.oceania.pool.ntp.org",
		"1.oceania.pool.ntp.org",
		"2.oceania.pool.ntp.org",
		"3.oceania.pool.ntp.org",
	}

	// Get offset from servers
	offsets = make(int64s, 0)
	var ntpTry int
	for ntpTry = 0; ntpTry <= maxNtpRetry; ntpTry++ {
		// Linear timeout increment
		timeout := time.Duration(int64(ntpTry+1) * int64(timeoutInMs) * time.Millisecond.Nanoseconds())
		options := ntp.QueryOptions{Timeout: timeout}
		joinNtps := sync.WaitGroup{}
		respCh := make(chan *ntp.Response)
		//respCh := make(chan *ntp.Response, len(ntpPool))
		for _, ntpAddr := range ntpPool {
			joinNtps.Add(1)
			go func(out chan<- *ntp.Response) {
				defer joinNtps.Done()
				resp, err := ntp.QueryWithOptions(ntpAddr, options)
				if err == nil {
					out <- resp
				}
			}(respCh)
		}

		go func() {
			joinNtps.Wait()
			close(respCh)
		}()

		joined := false
		for !joined {
			select {
			case resp := <- respCh:
				if resp == nil {
					joined = true
				} else {
					offsets = append(offsets, int64(resp.ClockOffset))
				}
			}
		}
		if offsets.Len() != 0 {
			break
		}
	}
	if ntpTry > maxNtpRetry {
		return fmt.Errorf("NTP Failure")
	}

	// Pick median
	sort.Sort(offsets)
	i := offsets.Len() / 2
	if offsets.Len() % 2 == 0 {
		sum := 0
		for offset := range offsets {
			sum += offset
		}
		avg := float32(sum) / float32(offsets.Len())
		if float32(offsets[i]) - avg > float32(offsets[i-1]) - avg {
			median = offsets[i-1]
		} else {
			median = offsets[i]
		}
	} else {
		median = offsets[i]
	}
	return nil
}


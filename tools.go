package rtools

import "time"

func RetryFor(f func() error, times, seconds int) error {
	var err error
	for i := 0; i < times; i++ {
		err = f()
		if err == nil {
			break
		}
		time.Sleep(time.Duration(seconds) * time.Second)
	}
	return err
}

package hcron

import (
	"context"
	"errors"

	"github.com/robfig/cron"
)

// CronTab crontab struct
type CronTab struct {
	Cron         *cron.Cron
	CronTabParam []CronTabParam
}

// CronTabParam crontab param
type CronTabParam struct {
	Spec        string
	CronTabFunc func()
}

// Start
func (c *CronTab) Start(ctx context.Context) error {
	// make error chan
	errChan := make(chan error)
	// goroutine
	func() {
		// new
		c.Cron = cron.New()
		// judgement crontab param length
		if len(c.CronTabParam) == 0 {
			errChan <- errors.New("craotab param length is empty")
		}
		// add func
		for _, v := range c.CronTabParam {
			err := c.Cron.AddFunc(v.Spec, v.CronTabFunc)
			if err != nil {
				errChan <- err
			}
		}
		// start
		c.Cron.Start()
	}()
	// select
	select {
	case <-ctx.Done():
		return ctx.Err()
	case e := <-errChan:
		return e
	}
} // end func Start

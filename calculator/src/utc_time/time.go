package utc_time

import "time"

type nowFuncTime func() time.Time

var nowFunc nowFuncTime = time.Now

func Reset(f func() time.Time) {
	nowFunc = f
}

func Now() time.Time {
	return nowFunc().UTC()
}

package job

type JobType float64

const (
	_                  = iota
	DEPLOYMENT JobType = 1 << (10 * iota)
)

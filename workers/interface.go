package workers

type Worker interface {
	IsThere() bool

	Do()
	Say()
	Work()
	ProgressRate() int
	GoHome()
}

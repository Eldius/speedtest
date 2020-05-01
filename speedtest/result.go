package speedtest

/*
Result is the speed test result
*/
type Result struct {
	Speed    float64
	Seconds  float64
	Size     float64
	File     string
	Latency  float64
	Sender   string // Address of the sender part
	Receiver string // Address of the receiver part
}

/*
NewResult creates a new result instance
*/
func NewResult(size, seconds float64) *Result {
	/*
	* Gets a size in bytes and a time in seconds and returns
	* speed in mbps. All values are float64
	 */
	r := Result{Size: size, Seconds: seconds, Speed: size * 8 / seconds}

	return &r
}

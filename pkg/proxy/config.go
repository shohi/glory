package proxy

type Config struct {
	// A rule is a pattern/address pair, which are delimited by `|`.
	// Multiple rules are separated by comma.
	Rules string

	// Listen port
	Port int
}

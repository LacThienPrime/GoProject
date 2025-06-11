package config

var (
	UnixSocketPerm string
	Address        string
	Port           int
	BaseURL        string
	BasePath       string
	BaseHost       string
)

type Configuration struct {
	UnixSocketPerm string
	Address        string
	Port           int
}

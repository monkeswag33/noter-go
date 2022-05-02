package types

type LogLevelParams struct {
	LogLevel     string
	GormLogLevel string
}

type HashParams struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

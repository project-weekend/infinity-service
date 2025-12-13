package config

// Config is the main configuration structure for qms-engine usecase
type Config struct {
	Name         string       `json:"name"`
	Prefork      bool         `json:"prefork"`
	ServiceName  string       `json:"serviceName"`
	Env          string       `json:"env"`
	Host         string       `json:"host"`
	Port         int          `json:"port"`
	OwnerInfo    OwnerInfo    `json:"ownerInfo"`
	Database     Database     `json:"database"`
	Statsd       Statsd       `json:"statsd"`
	Trace        Trace        `json:"trace"`
	Logger       Logger       `json:"logger"`
	ValkeyConfig ValkeyConfig `json:"valkeyConfig"`
	Kafka        KafkaConfig  `json:"kafka"`
	Security     Security     `json:"security"`
}

// OwnerInfo contains information about the usecase owner
type OwnerInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	URL   string `json:"url"`
}

type Database struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	Pool     struct {
		Idle     int `json:"idle"`
		Max      int `json:"max"`
		Lifetime int `json:"lifetime"`
	} `json:"pool"`
}

// CircuitBreaker contains circuit breaker configuration
type CircuitBreaker struct {
	TimeoutInMs            int `json:"timeoutInMs"`
	MaxConcurrentReq       int `json:"maxConcurrentReq"`
	VolumePercentThreshold int `json:"volumePercentThreshold"`
}

// Statsd contains StatsD configuration
type Statsd struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// Trace contains tracing configuration
type Trace struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Disable bool   `json:"disable"`
}

// Logger contains logging configuration
type Logger struct {
	WorkerCount     int    `json:"workerCount"`
	BufferSize      int    `json:"bufferSize"`
	LogLevel        int    `json:"logLevel"`
	StacktraceLevel int    `json:"stacktraceLevel"`
	LogFormat       string `json:"logFormat"`
}

// KafkaConfig contains Kafka configuration
type KafkaConfig struct {
	BootstrapServers string `json:"bootstrap.servers"`
	GroupID          string `json:"group.id"`
	AutoOffsetReset  string `json:"auto.offset.reset"`
	ProducerEnabled  bool   `json:"producer.enabled"`
}

// Valkey contains Valkey cache configuration
type ValkeyConfig struct {
	Enabled            bool   `json:"enabled"`
	Host               string `json:"host"`
	Port               int    `json:"port"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	Database           int    `json:"database"`
	PoolSize           int    `json:"poolSize"`
	MinIdleConns       int    `json:"minIdleConns"`
	MaxRetries         int    `json:"maxRetries"`
	ConnectTimeoutInMs int    `json:"connectTimeoutInMs"`
	ReadTimeoutInMs    int    `json:"readTimeoutInMs"`
	WriteTimeoutInMs   int    `json:"writeTimeoutInMs"`
	TLSEnabled         bool   `json:"tlsEnabled"`
	TTLInMinutes       int    `json:"ttlInMinutes"`
}

// Security contains security configuration
type Security struct {
	SecretKey string `json:"secretKey"`
}

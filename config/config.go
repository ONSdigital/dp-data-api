package config

import (
	"encoding/json"
	"time"

	"github.com/kelseyhightower/envconfig"
)

// Config represents service configuration for dp-data-api
type Config struct {
	BindAddr                   string        `envconfig:"BIND_ADDR"`
	GracefulShutdownTimeout    time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckInterval        time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
	HealthCheckCriticalTimeout time.Duration `envconfig:"HEALTHCHECK_CRITICAL_TIMEOUT"`
	OTBatchTimeout             time.Duration `encconfig:"OTEL_BATCH_TIMEOUT"`
	OTExporterOTLPEndpoint     string        `envconfig:"OTEL_EXPORTER_OTLP_ENDPOINT"`
	OTServiceName              string        `envconfig:"OTEL_SERVICE_NAME"`
	OtelEnabled                bool          `envconfig:"OTEL_ENABLED"`
	CodeListAPIURL             string        `envconfig:"CODE_LIST_API_URL"`
	DatasetAPIURL              string        `envconfig:"DATASET_API_URL"`
	WebsiteURL                 string        `envconfig:"WEBSITE_URL"`
	ZebedeeURL                 string        `envconfig:"ZEBEDEE_URL"`
	DownloadServiceSecretKey   string        `envconfig:"DOWNLOAD_SERVICE_SECRET_KEY"      json:"-"`
	ServiceAuthToken           string        `envconfig:"SERVICE_AUTH_TOKEN"               json:"-"`
	EnablePrivateEndpoints     bool          `envconfig:"ENABLE_PRIVATE_ENDPOINTS"`
	EnableDetachDataset        bool          `envconfig:"ENABLE_DETACH_DATASET"`
	EnablePermissionsAuth      bool          `envconfig:"ENABLE_PERMISSIONS_AUTH"`
	EnableObservationEndpoint  bool          `envconfig:"ENABLE_OBSERVATION_ENDPOINT"`
	DisableGraphDBDependency   bool          `envconfig:"DISABLE_GRAPH_DB_DEPENDENCY"`
	KafkaVersion               string        `envconfig:"KAFKA_VERSION"`
	DefaultMaxLimit            int           `envconfig:"DEFAULT_MAXIMUM_LIMIT"`
	DefaultLimit               int           `envconfig:"DEFAULT_LIMIT"`
	DefaultOffset              int           `envconfig:"DEFAULT_OFFSET"`
	MaxRequestOptions          int           `envconfig:"MAX_REQUEST_OPTIONS"`
	EncryptionDisabled         bool          `envconfig:"ENCRYPTION_DISABLED"`
	ComponentTestUseLogFile    bool          `envconfig:"COMPONENT_TEST_USE_LOG_FILE"`
}

var cfg *Config

// Get returns the default config with any modifications through environment
// variables
func Get() (*Config, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Config{
		BindAddr:                   "localhost:29300",
		GracefulShutdownTimeout:    5 * time.Second,
		HealthCheckInterval:        30 * time.Second,
		HealthCheckCriticalTimeout: 90 * time.Second,
		OTBatchTimeout:             5 * time.Second,
		OTExporterOTLPEndpoint:     "localhost:4317",
		OTServiceName:              "dp-data-api",
		OtelEnabled:                false,
		CodeListAPIURL:             "http://localhost:22400",
		DatasetAPIURL:              "http://localhost:22000",
		WebsiteURL:                 "http://localhost:20000",
		ZebedeeURL:                 "http://localhost:8082",
		ServiceAuthToken:           "FD0108EA-825D-411C-9B1D-41EF7727F465",
		DownloadServiceSecretKey:   "QB0108EZ-825D-412C-9B1D-41EF7747F462",
		EnablePrivateEndpoints:     false,
		EnableDetachDataset:        false,
		EnablePermissionsAuth:      false,
		EnableObservationEndpoint:  false,
		KafkaVersion:               "1.0.2",
		DefaultMaxLimit:            1000,
		DefaultLimit:               20,
		DefaultOffset:              0,
		MaxRequestOptions:          100, // Maximum number of options acceptable in an incoming Patch request. Compromise between one option per call (inefficient) and an order of 100k options per call, for census data (memory and computationally expensive)
	}

	return cfg, envconfig.Process("", cfg)

}

// String is implemented to prevent sensitive fields being logged.
// The config is returned as JSON with sensitive fields omitted.
func (config Config) String() string {
	b, _ := json.Marshal(config)
	return string(b)
}

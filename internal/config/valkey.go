package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"log/slog"
	"time"

	"github.com/valkey-io/valkey-go"

	"github.com/infinity/infinity-service/server/config"
)

// NewCache initializes and returns a Valkey client with full configuration support
func NewCache(appCfg *config.Config, slogger *slog.Logger) valkey.Client {
	slogger.Info("Initializing Valkey cache connection...")

	// Extract configuration
	host := appCfg.Valkey.Host
	port := appCfg.Valkey.Port
	username := appCfg.Valkey.Username
	password := appCfg.Valkey.Password
	database := appCfg.Valkey.Database
	poolSize := appCfg.Valkey.PoolSize
	minIdleConns := appCfg.Valkey.MinIdleConns
	maxRetries := appCfg.Valkey.MaxRetries
	connectTimeout := time.Duration(appCfg.Valkey.ConnectTimeoutInMs) * time.Millisecond
	readTimeout := time.Duration(appCfg.Valkey.ReadTimeoutInMs) * time.Millisecond
	writeTimeout := time.Duration(appCfg.Valkey.WriteTimeoutInMs) * time.Millisecond
	tlsEnabled := appCfg.Valkey.TLSEnabled

	// Build address
	address := fmt.Sprintf("%s:%d", host, port)

	// Set defaults for optional fields
	if poolSize <= 0 {
		poolSize = 10
		slogger.Info("Using default pool size", "poolSize", poolSize)
	}
	if minIdleConns <= 0 {
		minIdleConns = 2
	}
	if maxRetries <= 0 {
		maxRetries = 3
	}
	if connectTimeout == 0 {
		connectTimeout = 5 * time.Second
	}
	if readTimeout == 0 {
		readTimeout = 3 * time.Second
	}
	if writeTimeout == 0 {
		writeTimeout = 3 * time.Second
	}

	// Build client options
	clientOpt := valkey.ClientOption{
		InitAddress:      []string{address},
		ConnWriteTimeout: writeTimeout,
		MaxFlushDelay:    100 * time.Millisecond,
	}

	// Add authentication if provided
	if username != "" || password != "" {
		clientOpt.Username = username
		clientOpt.Password = password
		slogger.Info("Valkey authentication configured")
	}

	// Add database selection if specified
	if database > 0 {
		clientOpt.SelectDB = database
		slogger.Info("Valkey database selected", "database", database)
	}

	// Configure TLS if enabled
	if tlsEnabled {
		clientOpt.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
		slogger.Info("Valkey TLS enabled")
	}

	// Create client
	client, err := valkey.NewClient(clientOpt)
	if err != nil {
		slogger.Error("Failed to create Valkey client", "error", err)
		log.Fatalf("failed to create valkey client: %v", err)
	}

	// Test connection with ping
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	pingCmd := client.B().Ping().Build()
	resp := client.Do(ctx, pingCmd)
	if resp.Error() != nil {
		slogger.Error("Failed to ping Valkey server", "error", resp.Error())
		log.Fatalf("failed to ping valkey server: %v", resp.Error())
	}

	pongResult, err := resp.ToString()
	if err != nil {
		slogger.Error("Failed to get ping response", "error", err)
		log.Fatalf("failed to get ping response: %v", err)
	}

	slogger.Info("Valkey cache connection established successfully",
		"host", host,
		"port", port,
		"database", database,
		"poolSize", poolSize,
		"minIdleConns", minIdleConns,
		"maxRetries", maxRetries,
		"tlsEnabled", tlsEnabled,
		"pingResponse", pongResult,
	)

	return client
}

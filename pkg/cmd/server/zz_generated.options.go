// Code generated by github.com/ecordell/optgen. DO NOT EDIT.
package server

import (
	dispatch "github.com/authzed/spicedb/internal/dispatch"
	graph "github.com/authzed/spicedb/internal/dispatch/graph"
	datastore "github.com/authzed/spicedb/pkg/cmd/datastore"
	util "github.com/authzed/spicedb/pkg/cmd/util"
	datastore1 "github.com/authzed/spicedb/pkg/datastore"
	auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	grpc "google.golang.org/grpc"
	"time"
)

type ConfigOption func(c *Config)

// NewConfigWithOptions creates a new Config with the passed in options set
func NewConfigWithOptions(opts ...ConfigOption) *Config {
	c := &Config{}
	for _, o := range opts {
		o(c)
	}
	return c
}

// ToOption returns a new ConfigOption that sets the values from the passed in Config
func (c *Config) ToOption() ConfigOption {
	return func(to *Config) {
		to.GRPCServer = c.GRPCServer
		to.GRPCAuthFunc = c.GRPCAuthFunc
		to.PresharedKey = c.PresharedKey
		to.ShutdownGracePeriod = c.ShutdownGracePeriod
		to.DisableVersionResponse = c.DisableVersionResponse
		to.HTTPGateway = c.HTTPGateway
		to.HTTPGatewayUpstreamAddr = c.HTTPGatewayUpstreamAddr
		to.HTTPGatewayUpstreamTLSCertPath = c.HTTPGatewayUpstreamTLSCertPath
		to.HTTPGatewayCorsEnabled = c.HTTPGatewayCorsEnabled
		to.HTTPGatewayCorsAllowedOrigins = c.HTTPGatewayCorsAllowedOrigins
		to.DatastoreConfig = c.DatastoreConfig
		to.Datastore = c.Datastore
		to.NamespaceCacheConfig = c.NamespaceCacheConfig
		to.SchemaPrefixesRequired = c.SchemaPrefixesRequired
		to.DispatchServer = c.DispatchServer
		to.DispatchMaxDepth = c.DispatchMaxDepth
		to.GlobalDispatchConcurrencyLimit = c.GlobalDispatchConcurrencyLimit
		to.DispatchConcurrencyLimits = c.DispatchConcurrencyLimits
		to.DispatchUpstreamAddr = c.DispatchUpstreamAddr
		to.DispatchUpstreamCAPath = c.DispatchUpstreamCAPath
		to.DispatchClientMetricsPrefix = c.DispatchClientMetricsPrefix
		to.DispatchClusterMetricsPrefix = c.DispatchClusterMetricsPrefix
		to.DispatchUpstreamTimeout = c.DispatchUpstreamTimeout
		to.Dispatcher = c.Dispatcher
		to.DispatchCacheConfig = c.DispatchCacheConfig
		to.ClusterDispatchCacheConfig = c.ClusterDispatchCacheConfig
		to.DisableV1SchemaAPI = c.DisableV1SchemaAPI
		to.V1SchemaAdditiveOnly = c.V1SchemaAdditiveOnly
		to.MaximumUpdatesPerWrite = c.MaximumUpdatesPerWrite
		to.MaximumPreconditionCount = c.MaximumPreconditionCount
		to.ExperimentalCaveatsEnabled = c.ExperimentalCaveatsEnabled
		to.DashboardAPI = c.DashboardAPI
		to.MetricsAPI = c.MetricsAPI
		to.MiddlewareModification = c.MiddlewareModification
		to.DispatchUnaryMiddleware = c.DispatchUnaryMiddleware
		to.DispatchStreamingMiddleware = c.DispatchStreamingMiddleware
		to.SilentlyDisableTelemetry = c.SilentlyDisableTelemetry
		to.TelemetryCAOverridePath = c.TelemetryCAOverridePath
		to.TelemetryEndpoint = c.TelemetryEndpoint
		to.TelemetryInterval = c.TelemetryInterval
	}
}

// ConfigWithOptions configures an existing Config with the passed in options set
func ConfigWithOptions(c *Config, opts ...ConfigOption) *Config {
	for _, o := range opts {
		o(c)
	}
	return c
}

// WithGRPCServer returns an option that can set GRPCServer on a Config
func WithGRPCServer(gRPCServer util.GRPCServerConfig) ConfigOption {
	return func(c *Config) {
		c.GRPCServer = gRPCServer
	}
}

// WithGRPCAuthFunc returns an option that can set GRPCAuthFunc on a Config
func WithGRPCAuthFunc(gRPCAuthFunc auth.AuthFunc) ConfigOption {
	return func(c *Config) {
		c.GRPCAuthFunc = gRPCAuthFunc
	}
}

// WithPresharedKey returns an option that can append PresharedKeys to Config.PresharedKey
func WithPresharedKey(presharedKey string) ConfigOption {
	return func(c *Config) {
		c.PresharedKey = append(c.PresharedKey, presharedKey)
	}
}

// SetPresharedKey returns an option that can set PresharedKey on a Config
func SetPresharedKey(presharedKey []string) ConfigOption {
	return func(c *Config) {
		c.PresharedKey = presharedKey
	}
}

// WithShutdownGracePeriod returns an option that can set ShutdownGracePeriod on a Config
func WithShutdownGracePeriod(shutdownGracePeriod time.Duration) ConfigOption {
	return func(c *Config) {
		c.ShutdownGracePeriod = shutdownGracePeriod
	}
}

// WithDisableVersionResponse returns an option that can set DisableVersionResponse on a Config
func WithDisableVersionResponse(disableVersionResponse bool) ConfigOption {
	return func(c *Config) {
		c.DisableVersionResponse = disableVersionResponse
	}
}

// WithHTTPGateway returns an option that can set HTTPGateway on a Config
func WithHTTPGateway(hTTPGateway util.HTTPServerConfig) ConfigOption {
	return func(c *Config) {
		c.HTTPGateway = hTTPGateway
	}
}

// WithHTTPGatewayUpstreamAddr returns an option that can set HTTPGatewayUpstreamAddr on a Config
func WithHTTPGatewayUpstreamAddr(hTTPGatewayUpstreamAddr string) ConfigOption {
	return func(c *Config) {
		c.HTTPGatewayUpstreamAddr = hTTPGatewayUpstreamAddr
	}
}

// WithHTTPGatewayUpstreamTLSCertPath returns an option that can set HTTPGatewayUpstreamTLSCertPath on a Config
func WithHTTPGatewayUpstreamTLSCertPath(hTTPGatewayUpstreamTLSCertPath string) ConfigOption {
	return func(c *Config) {
		c.HTTPGatewayUpstreamTLSCertPath = hTTPGatewayUpstreamTLSCertPath
	}
}

// WithHTTPGatewayCorsEnabled returns an option that can set HTTPGatewayCorsEnabled on a Config
func WithHTTPGatewayCorsEnabled(hTTPGatewayCorsEnabled bool) ConfigOption {
	return func(c *Config) {
		c.HTTPGatewayCorsEnabled = hTTPGatewayCorsEnabled
	}
}

// WithHTTPGatewayCorsAllowedOrigins returns an option that can append HTTPGatewayCorsAllowedOriginss to Config.HTTPGatewayCorsAllowedOrigins
func WithHTTPGatewayCorsAllowedOrigins(hTTPGatewayCorsAllowedOrigins string) ConfigOption {
	return func(c *Config) {
		c.HTTPGatewayCorsAllowedOrigins = append(c.HTTPGatewayCorsAllowedOrigins, hTTPGatewayCorsAllowedOrigins)
	}
}

// SetHTTPGatewayCorsAllowedOrigins returns an option that can set HTTPGatewayCorsAllowedOrigins on a Config
func SetHTTPGatewayCorsAllowedOrigins(hTTPGatewayCorsAllowedOrigins []string) ConfigOption {
	return func(c *Config) {
		c.HTTPGatewayCorsAllowedOrigins = hTTPGatewayCorsAllowedOrigins
	}
}

// WithDatastoreConfig returns an option that can set DatastoreConfig on a Config
func WithDatastoreConfig(datastoreConfig datastore.Config) ConfigOption {
	return func(c *Config) {
		c.DatastoreConfig = datastoreConfig
	}
}

// WithDatastore returns an option that can set Datastore on a Config
func WithDatastore(datastore datastore1.Datastore) ConfigOption {
	return func(c *Config) {
		c.Datastore = datastore
	}
}

// WithNamespaceCacheConfig returns an option that can set NamespaceCacheConfig on a Config
func WithNamespaceCacheConfig(namespaceCacheConfig CacheConfig) ConfigOption {
	return func(c *Config) {
		c.NamespaceCacheConfig = namespaceCacheConfig
	}
}

// WithSchemaPrefixesRequired returns an option that can set SchemaPrefixesRequired on a Config
func WithSchemaPrefixesRequired(schemaPrefixesRequired bool) ConfigOption {
	return func(c *Config) {
		c.SchemaPrefixesRequired = schemaPrefixesRequired
	}
}

// WithDispatchServer returns an option that can set DispatchServer on a Config
func WithDispatchServer(dispatchServer util.GRPCServerConfig) ConfigOption {
	return func(c *Config) {
		c.DispatchServer = dispatchServer
	}
}

// WithDispatchMaxDepth returns an option that can set DispatchMaxDepth on a Config
func WithDispatchMaxDepth(dispatchMaxDepth uint32) ConfigOption {
	return func(c *Config) {
		c.DispatchMaxDepth = dispatchMaxDepth
	}
}

// WithGlobalDispatchConcurrencyLimit returns an option that can set GlobalDispatchConcurrencyLimit on a Config
func WithGlobalDispatchConcurrencyLimit(globalDispatchConcurrencyLimit uint16) ConfigOption {
	return func(c *Config) {
		c.GlobalDispatchConcurrencyLimit = globalDispatchConcurrencyLimit
	}
}

// WithDispatchConcurrencyLimits returns an option that can set DispatchConcurrencyLimits on a Config
func WithDispatchConcurrencyLimits(dispatchConcurrencyLimits graph.ConcurrencyLimits) ConfigOption {
	return func(c *Config) {
		c.DispatchConcurrencyLimits = dispatchConcurrencyLimits
	}
}

// WithDispatchUpstreamAddr returns an option that can set DispatchUpstreamAddr on a Config
func WithDispatchUpstreamAddr(dispatchUpstreamAddr string) ConfigOption {
	return func(c *Config) {
		c.DispatchUpstreamAddr = dispatchUpstreamAddr
	}
}

// WithDispatchUpstreamCAPath returns an option that can set DispatchUpstreamCAPath on a Config
func WithDispatchUpstreamCAPath(dispatchUpstreamCAPath string) ConfigOption {
	return func(c *Config) {
		c.DispatchUpstreamCAPath = dispatchUpstreamCAPath
	}
}

// WithDispatchClientMetricsPrefix returns an option that can set DispatchClientMetricsPrefix on a Config
func WithDispatchClientMetricsPrefix(dispatchClientMetricsPrefix string) ConfigOption {
	return func(c *Config) {
		c.DispatchClientMetricsPrefix = dispatchClientMetricsPrefix
	}
}

// WithDispatchClusterMetricsPrefix returns an option that can set DispatchClusterMetricsPrefix on a Config
func WithDispatchClusterMetricsPrefix(dispatchClusterMetricsPrefix string) ConfigOption {
	return func(c *Config) {
		c.DispatchClusterMetricsPrefix = dispatchClusterMetricsPrefix
	}
}

// WithDispatchUpstreamTimeout returns an option that can set DispatchUpstreamTimeout on a Config
func WithDispatchUpstreamTimeout(dispatchUpstreamTimeout time.Duration) ConfigOption {
	return func(c *Config) {
		c.DispatchUpstreamTimeout = dispatchUpstreamTimeout
	}
}

// WithDispatcher returns an option that can set Dispatcher on a Config
func WithDispatcher(dispatcher dispatch.Dispatcher) ConfigOption {
	return func(c *Config) {
		c.Dispatcher = dispatcher
	}
}

// WithDispatchCacheConfig returns an option that can set DispatchCacheConfig on a Config
func WithDispatchCacheConfig(dispatchCacheConfig CacheConfig) ConfigOption {
	return func(c *Config) {
		c.DispatchCacheConfig = dispatchCacheConfig
	}
}

// WithClusterDispatchCacheConfig returns an option that can set ClusterDispatchCacheConfig on a Config
func WithClusterDispatchCacheConfig(clusterDispatchCacheConfig CacheConfig) ConfigOption {
	return func(c *Config) {
		c.ClusterDispatchCacheConfig = clusterDispatchCacheConfig
	}
}

// WithDisableV1SchemaAPI returns an option that can set DisableV1SchemaAPI on a Config
func WithDisableV1SchemaAPI(disableV1SchemaAPI bool) ConfigOption {
	return func(c *Config) {
		c.DisableV1SchemaAPI = disableV1SchemaAPI
	}
}

// WithV1SchemaAdditiveOnly returns an option that can set V1SchemaAdditiveOnly on a Config
func WithV1SchemaAdditiveOnly(v1SchemaAdditiveOnly bool) ConfigOption {
	return func(c *Config) {
		c.V1SchemaAdditiveOnly = v1SchemaAdditiveOnly
	}
}

// WithMaximumUpdatesPerWrite returns an option that can set MaximumUpdatesPerWrite on a Config
func WithMaximumUpdatesPerWrite(maximumUpdatesPerWrite uint16) ConfigOption {
	return func(c *Config) {
		c.MaximumUpdatesPerWrite = maximumUpdatesPerWrite
	}
}

// WithMaximumPreconditionCount returns an option that can set MaximumPreconditionCount on a Config
func WithMaximumPreconditionCount(maximumPreconditionCount uint16) ConfigOption {
	return func(c *Config) {
		c.MaximumPreconditionCount = maximumPreconditionCount
	}
}

// WithExperimentalCaveatsEnabled returns an option that can set ExperimentalCaveatsEnabled on a Config
func WithExperimentalCaveatsEnabled(experimentalCaveatsEnabled bool) ConfigOption {
	return func(c *Config) {
		c.ExperimentalCaveatsEnabled = experimentalCaveatsEnabled
	}
}

// WithDashboardAPI returns an option that can set DashboardAPI on a Config
func WithDashboardAPI(dashboardAPI util.HTTPServerConfig) ConfigOption {
	return func(c *Config) {
		c.DashboardAPI = dashboardAPI
	}
}

// WithMetricsAPI returns an option that can set MetricsAPI on a Config
func WithMetricsAPI(metricsAPI util.HTTPServerConfig) ConfigOption {
	return func(c *Config) {
		c.MetricsAPI = metricsAPI
	}
}

// WithMiddlewareModification returns an option that can append MiddlewareModifications to Config.MiddlewareModification
func WithMiddlewareModification(middlewareModification MiddlewareModification) ConfigOption {
	return func(c *Config) {
		c.MiddlewareModification = append(c.MiddlewareModification, middlewareModification)
	}
}

// SetMiddlewareModification returns an option that can set MiddlewareModification on a Config
func SetMiddlewareModification(middlewareModification []MiddlewareModification) ConfigOption {
	return func(c *Config) {
		c.MiddlewareModification = middlewareModification
	}
}

// WithDispatchUnaryMiddleware returns an option that can append DispatchUnaryMiddlewares to Config.DispatchUnaryMiddleware
func WithDispatchUnaryMiddleware(dispatchUnaryMiddleware grpc.UnaryServerInterceptor) ConfigOption {
	return func(c *Config) {
		c.DispatchUnaryMiddleware = append(c.DispatchUnaryMiddleware, dispatchUnaryMiddleware)
	}
}

// SetDispatchUnaryMiddleware returns an option that can set DispatchUnaryMiddleware on a Config
func SetDispatchUnaryMiddleware(dispatchUnaryMiddleware []grpc.UnaryServerInterceptor) ConfigOption {
	return func(c *Config) {
		c.DispatchUnaryMiddleware = dispatchUnaryMiddleware
	}
}

// WithDispatchStreamingMiddleware returns an option that can append DispatchStreamingMiddlewares to Config.DispatchStreamingMiddleware
func WithDispatchStreamingMiddleware(dispatchStreamingMiddleware grpc.StreamServerInterceptor) ConfigOption {
	return func(c *Config) {
		c.DispatchStreamingMiddleware = append(c.DispatchStreamingMiddleware, dispatchStreamingMiddleware)
	}
}

// SetDispatchStreamingMiddleware returns an option that can set DispatchStreamingMiddleware on a Config
func SetDispatchStreamingMiddleware(dispatchStreamingMiddleware []grpc.StreamServerInterceptor) ConfigOption {
	return func(c *Config) {
		c.DispatchStreamingMiddleware = dispatchStreamingMiddleware
	}
}

// WithSilentlyDisableTelemetry returns an option that can set SilentlyDisableTelemetry on a Config
func WithSilentlyDisableTelemetry(silentlyDisableTelemetry bool) ConfigOption {
	return func(c *Config) {
		c.SilentlyDisableTelemetry = silentlyDisableTelemetry
	}
}

// WithTelemetryCAOverridePath returns an option that can set TelemetryCAOverridePath on a Config
func WithTelemetryCAOverridePath(telemetryCAOverridePath string) ConfigOption {
	return func(c *Config) {
		c.TelemetryCAOverridePath = telemetryCAOverridePath
	}
}

// WithTelemetryEndpoint returns an option that can set TelemetryEndpoint on a Config
func WithTelemetryEndpoint(telemetryEndpoint string) ConfigOption {
	return func(c *Config) {
		c.TelemetryEndpoint = telemetryEndpoint
	}
}

// WithTelemetryInterval returns an option that can set TelemetryInterval on a Config
func WithTelemetryInterval(telemetryInterval time.Duration) ConfigOption {
	return func(c *Config) {
		c.TelemetryInterval = telemetryInterval
	}
}

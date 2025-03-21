// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// GcsFaultToleranceOptionsApplyConfiguration represents a declarative configuration of the GcsFaultToleranceOptions type for use
// with apply.
type GcsFaultToleranceOptionsApplyConfiguration struct {
	RedisUsername            *RedisCredentialApplyConfiguration `json:"redisUsername,omitempty"`
	RedisPassword            *RedisCredentialApplyConfiguration `json:"redisPassword,omitempty"`
	ExternalStorageNamespace *string                            `json:"externalStorageNamespace,omitempty"`
	RedisAddress             *string                            `json:"redisAddress,omitempty"`
}

// GcsFaultToleranceOptionsApplyConfiguration constructs a declarative configuration of the GcsFaultToleranceOptions type for use with
// apply.
func GcsFaultToleranceOptions() *GcsFaultToleranceOptionsApplyConfiguration {
	return &GcsFaultToleranceOptionsApplyConfiguration{}
}

// WithRedisUsername sets the RedisUsername field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RedisUsername field is set to the value of the last call.
func (b *GcsFaultToleranceOptionsApplyConfiguration) WithRedisUsername(value *RedisCredentialApplyConfiguration) *GcsFaultToleranceOptionsApplyConfiguration {
	b.RedisUsername = value
	return b
}

// WithRedisPassword sets the RedisPassword field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RedisPassword field is set to the value of the last call.
func (b *GcsFaultToleranceOptionsApplyConfiguration) WithRedisPassword(value *RedisCredentialApplyConfiguration) *GcsFaultToleranceOptionsApplyConfiguration {
	b.RedisPassword = value
	return b
}

// WithExternalStorageNamespace sets the ExternalStorageNamespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExternalStorageNamespace field is set to the value of the last call.
func (b *GcsFaultToleranceOptionsApplyConfiguration) WithExternalStorageNamespace(value string) *GcsFaultToleranceOptionsApplyConfiguration {
	b.ExternalStorageNamespace = &value
	return b
}

// WithRedisAddress sets the RedisAddress field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RedisAddress field is set to the value of the last call.
func (b *GcsFaultToleranceOptionsApplyConfiguration) WithRedisAddress(value string) *GcsFaultToleranceOptionsApplyConfiguration {
	b.RedisAddress = &value
	return b
}

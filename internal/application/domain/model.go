package domain

// Discoverable defines the interface for all types with self discovery
type Discoverable interface {
	Discover() Discoverable
}

// ResponseSerializer defines the interface for all types that serialize to JSON response
type ResponseSerializer interface {
	SerializeResponse(any, bool) (JSONResponse, error)
}

// Settable defines the interface for all types with self setter
type Settable interface {
	Set(Settable) Settable
}

// DomainModel defines the interface for all domain models
type DomainModel interface {
	Discoverable
	ResponseSerializer
	Settable
}

// domainRegistry defines a domain registry (constants)
type domainRegistry struct {
	Episode string
	Season  string
}

// DomainType exposes constants for all domain types
var DomainType = domainRegistry{
	Episode: "episode",
	Season:  "season",
}

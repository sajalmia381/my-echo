package enums

type ENVIRONMENT string

const (
	PRODUCTION  = ENVIRONMENT("PRODUCTION")
	DEVELOPMENT = ENVIRONMENT("DEVELOPMENT")
)

// Database

type DATABASE string

const (
	MONGO = DATABASE("MONGO")
)

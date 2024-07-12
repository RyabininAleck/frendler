package constants

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleUser   Role = "user"
	RoleTester Role = "tester"
)

// Status represents the status of a user
type Status string

const (
	StatusActive  Status = "active"
	StatusBlock   Status = "block"
	StatusDeleted Status = "deleted"
)

// Platform represents the platform used by the user
type Platform string

const (
	PlatformVK       Platform = "vk"
	PlatformGoogle   Platform = "google"
	PlatformTelegram Platform = "telegram"
	PlatformContact  Platform = "contact"
)

// Theme represents the theme of the application
type Theme string

const (
	ThemeLight  Theme = "light"
	ThemeDark   Theme = "dark"
	ThemeSystem Theme = "system"
)

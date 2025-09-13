package constants

// API related constants
// These constants are used to define the base API group and specific API paths for the application.
const (
	BASE_API_GROUP            = "/api"
	API_SWAGGER_ACCEPT        = "/swagger/*any"
	API_SWAGGER_URL           = "/swagger/index.html"
	API_PATH_TRANSACTION      = "/transactions"
	API_PATH_TYPE             = "/types"
	API_PATH_TYPE_GROUP       = "/type-groups"
	API_PATH_USER             = "/users"
	API_PATH_CATEGORY         = "/categories"
	API_PATH_ACCOUNT          = "/accounts"
	API_PATH_WORKSPACE        = "/workspaces"
	API_PATH_WORKSPACE_MEMBER = "/workspace-members"
	AUTH_HEADER               = "Authorization"
	AUTH_USER_ID_KEY          = "userID"
	AUTH_EMAIL_KEY            = "email"
	AUTH_DISPLAY_NAME_KEY     = "displayName"
)

package auth

import (
	"sodnix/apps/server/src/common/constants"
	"sodnix/apps/server/src/common/response"
	"sodnix/apps/server/src/common/validator"
	"sodnix/apps/server/src/database"
	"sodnix/apps/server/src/modules/users"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// @title           Auth Service API
// @version         1.0
// @description     This is a sample authentication service.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Type "Bearer" + space + JWT token.

// Login godoc
// @Summary Authenticate user and get JWT token
// @Description Authenticates a user with email and password and returns a JWT token.
// @Tags auth
// @Accept  json
// @Produce  json
// @Param   request body LoginRequest true "Login credentials"
// @Success 200 {object} response.GetDataSuccess[LoginResponse] "Login successful"
// @Failure 400 {object} response.BadRequestError "Invalid credentials or bad request"
// @Failure 500 {object} response.InternalServerError "Internal server error"
// @Router /auth/login [post]
func Login(c *gin.Context) {
	var req LoginRequest

	if err := validator.ValidateRequestBody(c, &req); err != "" {
		response.Response(c, response.BadRequestResponse[string](err))
		return
	}

	userRepo := users.NewUserRepository(database.DB)
	user, err := userRepo.FindByEmail(req.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		response.Response(c, response.BadRequestResponse[string]("Invalid credentials"))
		return
	}

	// Set session
	SetSession(c, user.ID)

	// Generate JWT
	accessToken, refreshToken, err := GenerateJWT(user.ID, user.Email, user.DisplayName)
	if err != nil {
		response.Response(c, response.InternalServerErrorResponse[string]("Could not generate token"))
		return
	}

	response.Response(c, response.GetDataSuccessResponse(LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken, Message: "Login successful"}))
}

// Logout godoc
// @Summary Logout user
// @Description Clears the user session.
// @Tags auth
// @Produce  json
// @Success 200 {object} response.GetDataSuccess[MessageResponse] "Logged out"
// @Failure 400 {object} response.BadRequestError "Invalid credentials or bad request"
// @Failure 500 {object} response.InternalServerError "Internal server error"
// @Router /auth/logout [post]
func Logout(c *gin.Context) {
	ClearSession(c)
	response.Response(c, response.GetDataSuccessResponse(MessageResponse{Message: "Logged out"}))
}

// Profile godoc
// @Summary Get user profile
// @Description Retrieves the profile information for the authenticated user.
// @Tags auth
// @Produce  json
// @Security BearerAuth
// @Success 200 {object} response.GetDataSuccess[ProfileResponse] "User profile data"
// @Failure 401 {object} response.UnauthorizedError "Unauthorized"
// @Router /auth/profile [get]
func Profile(c *gin.Context) {
	userID, exists := c.Get(constants.AUTH_USER_ID_KEY)
	if !exists {
		response.Response(c, response.UnauthorizedResponse[string]("User ID not found in context"))
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		response.Response(c, response.InternalServerErrorResponse[string]("Invalid user ID type in context"))
		return
	}

	parsedUserID, err := uuid.Parse(userIDStr)
	if err != nil {
		response.Response(c, response.InternalServerErrorResponse[string]("Failed to parse user ID"))
		return
	}

	response.Response(c, response.GetDataSuccessResponse(ProfileResponse{Message: "Welcome", UserID: parsedUserID}))
}

package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"yogaflow.ai/database"
	"yogaflow.ai/models"
	"yogaflow.ai/services"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

// Register creates a new user account
func Register(c *gin.Context) {
	var req RegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user already exists
	var existingUser models.User
	err = database.Db.QueryRow("SELECT id FROM users WHERE email = $1", req.Email).Scan(&existingUser.ID)
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create user
	newUser := models.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Role:         "user",
		UserType:     "student",
		Tier:         "free",
		IsActive:     true,
	}

	user, err := services.CreateUser(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token
	token, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Clear password hash before sending response
	user.PasswordHash = ""

	c.JSON(http.StatusCreated, AuthResponse{
		Token: token,
		User:  user,
	})
}

// Login authenticates a user and returns a JWT token
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user models.User
	err := database.Db.QueryRow(
		"SELECT id, username, email, password_hash, first_name, last_name, bio, avatar_url, role, user_type, tier, is_active FROM users WHERE email = $1",
		req.Email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Bio,
		&user.AvatarURL,
		&user.Role,
		&user.UserType,
		&user.Tier,
		&user.IsActive,
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check if user is active
	if !user.IsActive {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Account is deactivated"})
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT token
	token, err := generateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Clear password hash before sending response
	user.PasswordHash = ""

	c.JSON(http.StatusOK, AuthResponse{
		Token: token,
		User:  user,
	})
}

// GetMe returns the currently authenticated user
func GetMe(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	var user models.User
	err := database.Db.QueryRow(
		"SELECT id, username, email, first_name, last_name, bio, avatar_url, role, user_type, tier, is_active FROM users WHERE id = $1",
		userID,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Bio,
		&user.AvatarURL,
		&user.Role,
		&user.UserType,
		&user.Tier,
		&user.IsActive,
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func generateToken(user models.User) (string, error) {
	claims := models.Claims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Tier:     user.Tier,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"pgd-server.com/helpers"
	"pgd-server.com/src/entities"
)

type Claims struct {
	UserId string `json:"userId"`
	RoleId string `json:"roleId"`
	jwt.StandardClaims
}

type LoginDto struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthController struct{}

func (ac *AuthController) GenerateToken(user *entities.User) (string, string, time.Time, error) {
	accessToken, expiredAt, _ := ac.generateAccessToken(user)
	refreshToken, err := ac.generateRefreshToken(user)
	if err != nil {
		return "", "", time.Now(), err
	}

	return accessToken, refreshToken, expiredAt, err
}

func (ac *AuthController) generateAccessToken(user *entities.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	return ac.generateToken(user, expirationTime, []byte(authMiddleware.GetJWTSecret()))
}

func (ac *AuthController) generateToken(user *entities.User, expirationTime time.Time, secret []byte) (string, time.Time, error) {

	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func (ac *AuthController) generateRefreshToken(user *entities.User) (string, error) {

	refreshToken := jwt.New(jwt.SigningMethodHS384)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["userId"] = user.ID

	rtClaims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()

	rt, err := refreshToken.SignedString([]byte(authMiddleware.GetJWTSecret()))
	if err != nil {
		return "error", err
	}

	return rt, err
}

func (ac *AuthController) JWTErrorChecker(err error, c *gin.Context) {
	// redirect to signing form
	response := new(helpers.GeneralResponse)
	response.Status = false
	response.Message = "Akses tidak sah"
	c.JSON(http.StatusUnauthorized, response)
}

// GET: /me
func (ac *AuthController) Me(c *gin.Context) {
	user, err := userService.FindOneUser(c.GetString("userId"), c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GeneralResponse{
			Message: "User tidak ditemukan",
			Status:  false,
		})
		return
	}

	c.JSON(http.StatusOK, helpers.GeneralResponse{
		Message: "me",
		Data:    user,
		Status:  true,
	})
}

// GET: /check-token
func (ac *AuthController) CheckToken(c *gin.Context) {
	response := new(helpers.GeneralResponse)
	success, claims, _ := authMiddleware.TokenCheckMiddleware(c)
	if !success {
		c.AbortWithStatusJSON(http.StatusBadRequest, helpers.GeneralResponse{
			Message: "Cek token tidak berhasil",
			Status:  false,
		})
		return
	}

	response.Message = "authorized"
	response.Status = true
	response.Data = claims
	c.JSON(http.StatusOK, response)
}

// POST :/login
func (ac *AuthController) SignIn(c *gin.Context) {
	loginDto := LoginDto{}

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, helpers.GeneralResponse{
			Status:  false,
			Message: "Gagal",
			Error:   "Data input tidak valid",
		})
		return
	}

	// find user first
	user, userError := userService.FindOneByEmailOrPhone(loginDto.Username)
	if userError != nil {
		c.JSON(http.StatusBadRequest, helpers.GeneralResponse{
			Status:  false,
			Message: "Gagal",
			Error:   "Email atau Nomor HP tidak ditemukan",
		})
		return
	}

	if user != nil {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, helpers.GeneralResponse{
				Status:  false,
				Error:   err.Error(),
				Message: "Kata sandi salah",
			})
			return
		} else {

			// if password is correct, generate tokens
			accessToken, refreshToken, expiredAt, err := ac.GenerateToken(&entities.User{
				ID:        user.ID,
				Email:     user.Email,
				Phone:     user.Phone,
				Name:      user.Name,
				Address:   user.Address,
				BirthDate: user.BirthDate,
				Password:  user.Password,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			})

			if err != nil {
				c.JSON(http.StatusUnauthorized, helpers.GeneralResponse{
					Message: "Token salah",
					Status:  false,
				})
				return
			}

			// store access token to db
			at := new(entities.AccessToken)
			at.AccessToken = accessToken
			at.UserId = user.ID
			at.Expired = false
			at.ExpiredAt = expiredAt
			_ = accessTokenService.StoreAccessToken(at)

			c.JSON(http.StatusOK, helpers.GeneralResponse{
				Data: gin.H{
					"accessToken":  accessToken,
					"refreshToken": refreshToken,
				},
				Message: "authorized",
				Status:  true,
			})
			return

			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, helpers.GeneralResponse{
			Message: "user tidak ditemukan",
			Status:  false,
		})
		return
	}
}

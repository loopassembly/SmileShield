package controllers

import (
	"cypher-server/initializers"
	"cypher-server/models"
	"cypher-server/utils"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(c *fiber.Ctx) error {
	var payload *models.SignUpInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})

	}

	if payload.Password != payload.PasswordConfirm {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	newUser := models.User{
		Name:     payload.Name,
		Email:    strings.ToLower(payload.Email),
		Password: string(hashedPassword),
		Photo:    &payload.Photo,
	}

	result := initializers.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User with that email already exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	// email
	// The code below is from your initial code
	config, _ := initializers.LoadConfig(".")
	code := randstr.String(20)
	verificationCode := utils.Encode(code)
	newUser.VerificationCode = verificationCode
	initializers.DB.Save(&newUser)

	var firstName = newUser.Name
	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	emailData := utils.EmailData{
		URL:       config.ClientOrigin + "api/auth/verifyemail/" + code,
		// URL:       "192.168.64.202:3000/api/auth/verifyemail/" + code,
		FirstName: firstName,
		Subject:   "Your account verification code",
	}

	utils.SendEmail(&newUser, &emailData, "verificationCode.html")

	message := "We sent an email with a verification code to " + newUser.Email

	// return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"user": models.FilterUserRecord(&newUser)}})
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "message": message})
}



func VerifyEmail(c *fiber.Ctx) error {
	code := c.Params("verificationCode")
	verificationCode := utils.Encode(code)
	// fmt.Println(verificationCode)
	var updatedUser models.User
	result := initializers.DB.First(&updatedUser, "verification_code = ?", verificationCode)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid verification code or user doesn't exist"})
	}

	if *updatedUser.Verified {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "User already verified"})
	}

	updatedUser.VerificationCode = ""
	*updatedUser.Verified = true
	initializers.DB.Save(&updatedUser)

	// Render the email confirmation template
	if err := renderConfirmationTemplate(c); err != nil {
		log.Printf("Error rendering confirmation template: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal Server Error"})
	}

	return nil
}


func ForgotPassword(c *fiber.Ctx) error {
	var payload *models.ForgotPasswordInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	message := "You will receive a reset email if the user with that email exists"

	var user models.User
	result := initializers.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}

	if !*user.Verified {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Account not verified"})
	}

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config", err)
	}

	// Generate Verification Code
	resetToken := randstr.String(20)

	passwordResetToken := utils.Encode(resetToken)
	user.PasswordResetToken = passwordResetToken
	user.PasswordResetAt = time.Now().Add(time.Minute * 15)
	initializers.DB.Save(&user)

	firstName := user.Name
	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}

	// Send Email
	emailData := utils.EmailData{
		URL:       config.ClientOrigin + "api/auth/resetpassword/" + resetToken,
		FirstName: firstName,
		Subject:   "Your password reset token (valid for 10min)",
	}

	utils.SendEmail(&user, &emailData, "resetPassword.html")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": message})
}

func ResetPassword(c *fiber.Ctx) error {
	var payload *models.ResetPasswordInput
	resetToken := c.Params("resetToken")

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if payload.Password != payload.PasswordConfirm {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})
	}

	hashedPassword, _ := utils.HashPassword(payload.Password)

	passwordResetToken := utils.Encode(resetToken)

	var updatedUser models.User
	result := initializers.DB.First(&updatedUser, "password_reset_token = ? AND password_reset_at > ?", passwordResetToken, time.Now())
	if result.Error != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "The reset token is invalid or has expired"})
	}

	updatedUser.Password = hashedPassword
	updatedUser.PasswordResetToken = ""
	initializers.DB.Save(&updatedUser)

	// Assuming you want to clear a token in the response
	c.ClearCookie("token")

	return c.Status(http.StatusOK).JSON(fiber.Map{"status": "success", "message": "Password data updated successfully"})
}



func renderConfirmationTemplate(c *fiber.Ctx) error {
	

	// Render the template and send the output to the client
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
	
}

func SignInUser(c *fiber.Ctx) error {
	var payload *models.SignInInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	var user models.User
	result := initializers.DB.First(&user, "email = ?", strings.ToLower(payload.Email))
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}
	if !*user.Verified {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Email not verified"})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid email or Password"})
	}

	config, _ := initializers.LoadConfig(".")

	tokenByte := jwt.New(jwt.SigningMethodHS256)

	now := time.Now().UTC()
	claims := tokenByte.Claims.(jwt.MapClaims)

	claims["sub"] = user.ID
	claims["exp"] = now.Add(config.JwtExpiresIn).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	tokenString, err := tokenByte.SignedString([]byte(config.JwtSecret))

	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": fmt.Sprintf("generating JWT Token failed: %v", err)})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   config.JwtMaxAge * 60,
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": tokenString})
}

func LogoutUser(c *fiber.Ctx) error {
	expired := time.Now().Add(-time.Hour * 24)
	c.Cookie(&fiber.Cookie{
		Name:    "token",
		Value:   "",
		Expires: expired,
	})
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success"})
}

func GoogleOAuth(c *fiber.Ctx) error {
	code := c.Query("code")
	// var pathURL string = "/"

	if c.Query("state") != "" {
		c.Query("state")
	}

	if code == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Authorization code not provided!"})
	}

	tokenRes, err := utils.GetGoogleOauthToken(code)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	googleUser, err := utils.GetGoogleUser(tokenRes.Access_token, tokenRes.Id_token)
	if err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	now := time.Now()
	email := strings.ToLower(googleUser.Email)
	googleProvider := "Google"
	roleuser := "user"
	verified := true
	userData := models.User{
		Name:      googleUser.Name,
		Email:     email,
		Password:  "",
		Photo:     &googleUser.Picture,
		Provider:  &googleProvider,
		Role:      &roleuser,
		Verified:  &verified,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if initializers.DB.Model(&userData).Where("email = ?", email).Updates(&userData).RowsAffected == 0 {
		initializers.DB.Create(&userData)
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", email)

	config, _ := initializers.LoadConfig(".")

	token, err := utils.GenerateToken(config.JwtExpiresIn,user.ID, config.JwtSecret)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   int(config.JwtMaxAge * 60),
		Secure:   false,
		HTTPOnly: true,
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": token})

}

func GitHubOAuth(c *fiber.Ctx) error {
	code := c.Query("code")
	

	if c.Query("state") != "" {
		 c.Query("state")
	}

	if code == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Authorization code not provided!"})
	}

	tokenRes, err := utils.GetGitHubOauthToken(code)

	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	githubUser, err := utils.GetGitHubUser(tokenRes.Access_token)

	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	now := time.Now()
	email := strings.ToLower(githubUser.Email)
	provider := "GitHub"
	roleuser := "user"
	verified := true
	userData := models.User{
		Name:      githubUser.Name,
		Email:     email,
		Password:  "",
		Photo:     &githubUser.Photo,
		Provider:  &provider,
		Role:      &roleuser,
		Verified:  &verified,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if initializers.DB.Model(&userData).Where("email = ?", email).Updates(&userData).RowsAffected == 0 {
		initializers.DB.Create(&userData)
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", email)

	config, _ := initializers.LoadConfig(".")

	token, err := utils.GenerateToken(config.JwtExpiresIn, user.ID, config.JwtSecret)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Minute * time.Duration(config.JwtMaxAge)),
		HTTPOnly: true,
		SameSite: "Lax",
		Secure:   false, // Set to true if using HTTPS
		Path:     "/",
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": token})
}


func StackoverflowOAuth(c *fiber.Ctx) error {
	code := c.Query("code")
	

	if c.Query("state") != "" {
		 c.Query("state")
	}

	if code == "" {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Authorization code not provided!"})
	}

	tokenRes, err := utils.GetStackOverflowOauthToken(code)

	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	stackoverflowUser, err := utils.GetStackOverflowUser(tokenRes.Access_token)

	if err != nil {
		return c.Status(http.StatusBadGateway).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	now := time.Now()
	email := strings.ToLower(stackoverflowUser.Email)
	provider := "Stackoverflow"
	roleuser := "user"
	verified := true
	userData := models.User{
		Name:      stackoverflowUser.DisplayName,
		Email:     email,
		Password:  "",
		Photo:     &stackoverflowUser.ProfileURL,
		Provider:  &provider,
		Role:      &roleuser,
		Verified:  &verified,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if initializers.DB.Model(&userData).Where("email = ?", email).Updates(&userData).RowsAffected == 0 {
		initializers.DB.Create(&userData)
	}

	var user models.User
	initializers.DB.First(&user, "email = ?", email)

	config, _ := initializers.LoadConfig(".")

	token, err := utils.GenerateToken(config.JwtExpiresIn, user.ID, config.JwtSecret)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token, //jwt
		Expires:  time.Now().Add(time.Minute * time.Duration(config.JwtMaxAge)),
		HTTPOnly: true,
		SameSite: "Lax",
		Secure:   false, // Set to true if using HTTPS
		Path:     "/", // path 
		Domain:   "localhost",
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "token": token})
}

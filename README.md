# SmileShield - Go Fiber OAuth and Mail Authentication

Welcome to SmileShield, a Go Fiber project for OAuth and mail authentication! SmileShield provides a secure and convenient way to handle user authentication in your applications.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Features](#features)
- [Endpoints](#endpoints)
- [Example](#example)
- [Contributing](#contributing)
- [License](#license)

## Installation

 1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/smileshield.git
    cd smileshield
    ```

2. Install dependencies:

    ```bash
    go get -u .
    ```

3. Setup '.env':

    ```bash
    # POSTGRES_HOST=127.0.0.1
   # POSTGRES_USER=postgres
   # POSTGRES_PASSWORD=password123
   # POSTGRES_DB=golang-gorm
   # POSTGRES_PORT=6500

   PORT=8000
   CLIENT_ORIGIN='http://localhost:3000/'


   EMAIL_FROM='enter_email'
   SMTP_HOST='enter_host_name'
   SMTP_USER='enter_email_if_u_don't_know'
   SMTP_PASS='*******'
   SMTP_PORT='according_to_ur_provider'

   TOKEN_EXPIRED_IN=60m
   TOKEN_MAXAGE=60

   TOKEN_SECRET='my-ultra-secure-json-web-token-string'

   JWT_SECRET='my_ultra_secure_secret'
   JWT_EXPIRED_IN=60m
   JWT_MAXAGE=60

   #? google auth
   GOOGLE_OAUTH_CLIENT_ID='xyz...'
   GOOGLE_OAUTH_CLIENT_SECRET='secret....'
   GOOGLE_OAUTH_REDIRECT_URL='http://localhost:3000/api/auth/sessions/oauth/google'

   #? github auth
   GITHUB_OAUTH_CLIENT_ID='xyz...'
   GITHUB_OAUTH_CLIENT_SECRET='secret....'
   GITHUB_OAUTH_REDIRECT_URL='http://localhost:3000/api/auth/sessions/oauth/github'

   # ?stackoverflow auth
   STACKOVERFLOW_OAUTH_CLIENT_ID='xyz...'
   STACKOVERFLOW_CLIENT_SECRET='tsecret....'
   STACKOVERFLOW_REDIRECT_URL='http://localhost:3000/api/auth/sessions/oauth/stackoverflow' ```

4.  Run:

    ```bash
    go run main.go
    ```

## Usage

Key dependencies and libraries powering SmileShield:

> **Validator ([`go-playground/validator/v10`](https://github.com/go-playground/validator/v10) v10.16.0):**
   > - Validates input data against specified criteria.
>
> **Fiber ([`gofiber/fiber/v2`](https://github.com/gofiber/fiber/v2) v2.50.0):**
  > - Core web framework for handling HTTP requests and responses.
>
> **Fiber HTML Template Engine ([`gofiber/template/html/v2`](https://github.com/gofiber/template/html/v2) v2.0.5):**
  > - Renders HTML views in the web application.
>
> **JWT Library ([`golang-jwt/jwt`](https://github.com/golang-jwt/jwt) v3.2.2+incompatible, [`golang-jwt/jwt/v5`](https://github.com/golang-jwt/jwt/v5) v5.0.0):**
   > - Manages JSON Web Tokens for user authentication.
>
> **UUID ([`google/uuid`](https://github.com/google/uuid) v1.4.0):**
  > - Generates and works with universally unique identifiers (UUIDs).
>
> **HTML to Text Converter ([`k3a/html2text`](https://github.com/k3a/html2text) v1.2.1):**
  > - Converts HTML to plain text.
>
> **Viper ([`spf13/viper`](https://github.com/spf13/viper) v1.17.0):**
   > - Popular configuration management library for reading configuration files.
>
> **Random String Generator ([`thanhpk/randstr`](https://github.com/thanhpk/randstr) v1.0.6):**
   > - Generates random strings for various purposes.
>
> **Crypto ([`golang.org/x/crypto`](https://pkg.go.dev/golang.org/x/crypto) v0.14.0):**
  >  - Part of the Go standard library, includes cryptographic primitives.
>
> **Email ([`gomail.v2`](https://pkg.go.dev/gopkg.in/gomail.v2) v2.0.0-20160411212932-81ebce5c23df):**
   >  - Library for sending emails, crucial for applications involving email functionality.
>
> **GORM SQLite Driver ([`gorm.io/driver/sqlite`](https://gorm.io/docs/sqlite.html) v1.5.4):**
   >  - SQLite driver for GORM, a powerful Object-Relational Mapping (ORM) library.
>
> **GORM ([`gorm.io/gorm`](https://gorm.io/docs/index.html) v1.25.5):**
   > - Provides a flexible way to interact with databases.

### Note:

- Understand the licenses associated with these dependencies.
- Keep dependencies up to date for security and feature improvements.
- Refer to the project's documentation for specific use cases or configurations.


## Features

 **OAuth Authentication:**
   - Integrate seamlessly with popular OAuth providers for secure and efficient user authentication.

 **Mail Authentication:**
   - Simplify user registration by incorporating email confirmation as part of the authentication process.

 **Flexible Configuration:**
   - Easily customize settings to adapt SmileShield to your specific environment and requirements.



## Endpoints

- [/auth/login](https://github.com/loopassembly/SmileShield/blob/main/routes/auth.routes.go): Endpoint for user login with email and password.
- [/auth/register](https://github.com/loopassembly/SmileShield/blob/main/routes/auth.routes.go): Endpoint for user registration with email confirmation.
- [auth/sessions/oauth/google](https://github.com/loopassembly/SmileShield/blob/main/routes/auth.routes.go): Endpoint for user login using Google OAuth.
- [/auth/sessions/oauth/google](https://github.com/loopassembly/SmileShield/blob/main/routes/auth.routes.go): Endpoint for user login using Google OAuth.
- [/auth/verifyemail/:verificationCode](https://github.com/loopassembly/SmileShield/blob/main/routes/auth.routes.go): Endpoint for user email verification with verification code.
- [/auth/forgotpassword](https://github.com/loopassembly/SmileShield/blob/main/routes/auth.routes.go): Endpoint for forgot password.
- [/auth/resetpassword/:resetToken](https://github.com/loopassembly/SmileShield/blob/main/routes/auth.routes.go): Endpoint for reset password.


## Example

### User Registration

> **Navigate to the Registration Page:**

   [![Registration Page](https://cdn.discordapp.com/attachments/1102161138625564673/1193158525183078481/image.png?ex=65abb278&is=65993d78&hm=700ca6e6743220ef080c6a7d04da997e796b7b4cb0d6a47a2d7f795a050b894a&)](registration_page_url)



### Making a Call with Authorization Token

1. **Copy the JWT Token from the Login Response.**

   ![JWT Token](https://cdn.discordapp.com/attachments/1102161138625564673/1193162177922732123/image.png?ex=65abb5df&is=659940df&hm=efb361121bf3c5e2e9d0ca87fbc15e01156c72a16eaf788d6f37a48b551d3a41&)

2. **Make an Authenticated API Call:**

   ```bash
   curl -X GET \
     https://api.example.com/user-profile \
     -H "Authorization: Bearer YOUR_JWT_TOKEN"
      
![JWT Token](https://cdn.discordapp.com/attachments/1102161138625564673/1193167247657680936/image.png?ex=65abba97&is=65994597&hm=23d4065589c8aeb78870b316a8cd6d3000110dbe77c82c35b1278466c5d1728a&)
### Google OAuth Configuration Steps

1. **Create OAuth Credentials:**

   - Go to the [Google Cloud Console](https://console.cloud.google.com/).
   - Navigate to the "APIs & Services" > "Credentials" section.
   - Click on "Create Credentials" and choose "OAuth client ID."
   - Configure the consent screen and application type (Web application for this example).
   - Set the authorized redirect URI(s) for your application.

2. **Retrieve Client ID and Client Secret:**

   - Once the OAuth client is created, note the generated Client ID and Client Secret.

3. **Update SmileShield Configuration:**

   - Open your `.env` file.
   - Add the Google OAuth Client ID and Client Secret:

     ```env
     GOOGLE_OAUTH_CLIENT_ID=your_google_oauth_client_id
     GOOGLE_OAUTH_CLIENT_SECRET=your_google_oauth_client_secret
     ```


![Google OAuth](https://imgur.com/0XgvnpN.gif)

Replace placeholders like `your_google_oauth_client_id` and `your_google_oauth_client_secret` with your actual Google OAuth credentials. The animated GIF shows the process of updating the configuration in your `.env` file for Google OAuth.

### GitHub OAuth Configuration Steps

1. **Create OAuth Credentials:**

   - Go to [GitHub Developer Settings](https://github.com/settings/developers).
   - Create a new OAuth App.
   - Configure the application details and set the authorization callback URL(s) for your application.

2. **Retrieve Client ID and Client Secret:**

   - Once the OAuth App is created, note the generated Client ID and Client Secret.

3. **Update SmileShield Configuration:**

   - Open your `.env` file.
   - Add the GitHub OAuth Client ID and Client Secret:

     ```env
     GITHUB_OAUTH_CLIENT_ID=your_github_oauth_client_id
     GITHUB_OAUTH_CLIENT_SECRET=your_github_oauth_client_secret
     ```

![GitHub OAuth](https://i.imgur.com/utlpi0k.gif)

Replace placeholders like `your_github_oauth_client_id` and `your_github_oauth_client_secret` with your actual GitHub OAuth credentials. The animated GIF shows the process of updating the configuration in your `.env` file for GitHub OAuth.


## Contributing

Contributions are welcome! We appreciate your interest in improving SmileShield. Whether you want to report a bug, suggest a feature, or contribute code, here's how you can get involved:

### Bug Reports and Feature Requests

If you encounter a bug or have a feature request, please open an issue on the [GitHub Issues](https://github.com/your-username/smileshield/issues) page. Provide as much detail as possible, including steps to reproduce the issue or a clear description of the new feature.

### Pull Requests

1. Fork the repository and create your branch from `main`.

    ```bash
    git checkout -b your-branch-name
    ```

2. Make your changes, test thoroughly, and ensure your code adheres to the project's coding standards.

3. Commit your changes with descriptive commit messages.

    ```bash
    git commit -m "Your descriptive commit message"
    ```

4. Push your changes to your fork.

    ```bash
    git push origin your-branch-name
    ```

5. Open a pull request to the `main` branch of the original repository.

### Stack Overflow OAuth

We're currently working on adding Stack Overflow OAuth to SmileShield! If you're interested in contributing to this specific feature, follow the general contribution steps mentioned above. When working on the Stack Overflow OAuth integration, please ensure that your changes align with the project's goals and coding standards.

### Getting Help

If you have questions or need help, feel free to reach out on our [Discussions](https://github.com/loopassembly/smileshield/discussions) page.

Thank you for your contribution! 🚀

## License

MIT License
Copyright (c) 2024 Ashutosh Anand.

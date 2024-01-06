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
    PORT=3000
    OAUTH_CLIENT_ID=your_oauth_client_id
    OAUTH_CLIENT_SECRET=your_oauth_client_secret
    MAIL_USERNAME=your_mail_username
    MAIL_PASSWORD=your_mail_password
    ```

4.  Usage:

    ```bash
    go run main.go
    ```

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

### User Login

> **Navigate to the Login Page:**

   [![Login Page](https://cdn.discordapp.com/attachments/1102161138625564673/1193162177922732123/image.png?ex=65abb5df&is=659940df&hm=efb361121bf3c5e2e9d0ca87fbc15e01156c72a16eaf788d6f37a48b551d3a41&)](login_page_url)

## Contributing

Explain how others can contribute to your project.

## License

Specify the license for your project.

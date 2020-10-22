# Send an email using Go Using OAUTH2

## GMAIL API
Before using this program, you should get an access token and a refresh token from Google API.
Please refer to this document for more details : [https://developers.google.com/identity/protocols/oauth2](https://developers.google.com/identity/protocols/oauth2)

## Required Environment Variables
Set these env vars before running the program:

``` bash
GOOGLE_CLIENT_ID=""
GOOGLE_CLIENT_SECRET=""
GOOGLE_REFRESH_TOKEN=""
GOOGLE_ACCESS_TOKEN=""
```

## Send an email
``` bash
go run main.go -to=recipient_email_address@domain.com
```
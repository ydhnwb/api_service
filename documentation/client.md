# Client Documentation

This `client` terms means it is a collections of api documentation that used by client (reader). For backoffice and creator api documentation, please see others page.

## Auth

### Login

Authenticate user using email and password.

```
http://localhost:8080/v1/api/login
```

Payload

```json
{
  "email": "johndoe@mail.com", //required
  "password": "johndoe123" //required
}
```

Success response (User object)

```json
{
  "id": 1,
  "name": "John Doe",
  "bio": "...",
  "birth_of_date": "07-04-1998",
  "is_active": true,
  "email": "johndoe@mail.com",
  "token": ".....",
  "refresh_token": "....",
  "created_at": "TIMESTAMP",
  "updated_at": "TIMESTAMP"
}
```

Error

```json
{
  "error": "Example of error message"
}
```

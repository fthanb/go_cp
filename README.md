## Get User Information

### Description
This endpoint retrieves information about a specific user.

### Endpoint
- **HTTP Method:** GET
- **URL:** `/api/users/{userId}`

### Request Headers
- `Authorization: Bearer <token>` (required)

### Request Parameters
- **Path Parameters:**
  - `userId` (string, required): The ID of the user to retrieve.

### Request Example

```sh
curl -X GET "https://api.example.com/api/users/12345" \
     -H "Authorization: Bearer your-token"

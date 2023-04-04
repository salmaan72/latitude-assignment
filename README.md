# latitude-assignment

## Requirements
- Users in this system has their own ledger. Let's say it holds current balance amount of the user.
- Any third party app or admin has access to the current balance of all the users. In this case, let's say admin has a dashboard and all the users current balance is shown.
- When a user is authorized, admin should not be able to access current balance of that particular user.(ringfence, reserve that account info of the user).
- User session expires in 100 seconds.
- Once user session expires, admin should again be able to access that particular user's current balance.

## Solution
- This is achieved using JWT tokens.
- when a user logs in, access token is set with an expiration of 100 seconds. This token is stored on the server.
- when admin logs in and want to check dashboard, the loggedin users info is not accessible by admin. "Resource already in use" message is shown. When admin tries to access user info, backend fetches that particular user's accesstoken and checks if it's expired. If the token is expired, then admin can access user info. If not throw "Resource already in use" message.
- Redis is used in the backend to store user accesstoken.
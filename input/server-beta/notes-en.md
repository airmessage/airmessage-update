This release improves the stability of AirMessage Server 4.

- Implemented recovery logic for when the latest.log is deleted during a session
- Implemented recovery logic for when the user rejects Keychain access
- Fixed a crash when encountering a communications error with Firebase while signing in with an account
- Fixed a crash when dismissing the account sign-in sheet using the "Cancel" button while the authentication prompt was being displayed
- 
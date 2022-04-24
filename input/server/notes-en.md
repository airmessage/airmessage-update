This update improves the stability, compatibility, and security of AirMessage Server.

- Fixed a bug where update checks would not download the latest update data for short periods of time
- Fixed a bug where stickers and tapbacks would not show up when loading past messages on AirMessage for web
- Fixed a crash when generating a FaceTime link if the link was not copied correctly
- Fixed a crash that could occur in certain cases when pinging connected clients
- Added logging for client connections and disconnections in error reports
- Improved the security of advanced message fetch requests
- Upgraded to OpenSSL 3.0.2

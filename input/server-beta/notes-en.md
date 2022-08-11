AirMessage Server 4.1.2 fixes a bug when uploading or downloading attachment files over AirMessage Cloud.

AirMessage Server 4.1.1 has these changes:

- Added an error message to inform the user when an update fails because the app needs to be moved
- Improved lost connection detection over AirMessage Cloud
- Fixed a bug where previously past messages would be duplicated on AirMessage devices in certain cases
- Fixed a bug where the server would be unable to reconnect to AirMessage Cloud after getting disconnected
- Fixed a performance issue when checking the status of read receipts (thanks to Geczy for troubleshooting)
- Fixed a bug where setting a blank password would revert after restarting the app

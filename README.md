# Push to Google Sheets Example
- Secure server-to-server via API keys
## Manual Auth Setup
In GCP project:
1. Service Accounts -> Create New
    - Add credentials JSON (with spaces removed) to secure env var `GCP_CREDENTIALS_JSON`
2. Create & share spreadsheet with service account email address (or add Drive permissions)
    - Add sheet ID to env var `GCP_SHEET_ID`
3. Create a sheet tab called "PushSheet" or set the env var `GCP_SHEET_NAME` to the sheet tab name.
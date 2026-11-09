# Backend source

- `information_serve/`: Go API service and admin API.
- `information-app_adminui/`: Vue administration UI.

This directory records the existing backend source. Importing it does not mean
that its build, deployment or business flows have been verified.

## Local files excluded from Git

Private keys and certificates, runtime environment files, service settings YAML,
SQL exports, SQLite databases, uploaded files, logs, IDE state and compiled
artifacts stay on the local machine. The existing SMTP integration test also
stays local because it contains account credentials.

The service's `config/settings*.yml` files and the admin UI's `.env*` files must
be supplied separately for a deployment. Database setup and third-party service
credentials must also be supplied separately; this source import includes no
database export or production credentials.

Dependency lock files are included for both projects. Refer to each project's
README and build scripts for its original setup instructions.

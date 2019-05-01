# GOCOO

This application shows how to use Nulab Cacoo APIs in Go.
Some API calls are made in parallel to show how to use this Go feature.

### Configuration:

The application reads configuration from a __config.toml__ file placed in the root folder of the project. The file mut have the following structure:

```toml
user_id = "your-user-ID"
api_key = "your-api-key"
redis_host = "your-redis-server-url" # only in case you are not running the app from the docker container.
```

### Running the application:

Cd into the project folder and run: ```docker-compose up```

### Available routes:

- >/api/user
    >
    >Returns basic user information.

- >/api/info
    >
    >Returns information on user activity: numbers of diagrams and folders.
    
- >/api/folder/{folder-ID}
    >
    >Returns information for a given folder ID.

- >/api/diagram/{diagram-ID}
    >
    >Returns information for a given diagram ID,
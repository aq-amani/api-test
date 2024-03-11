# api-test

A simple backend service that can store and retrieve audio recordings associated with user IDs and practice phrases

## prerequisites
- docker

## Running the service

1. Create a `.env` file in the project root and configure the following environment variables:

    ```env
    POSTGRES_USER=my_user
    POSTGRES_PASSWORD=my_password
    ```

2. Build and run the service using Docker Compose:

    ```bash
    docker-compose up -d
    ```

## Usage
- Store audio
```
curl --request POST 'http://localhost/audio/user/1/phrase/1' --form 'audio_file=@"./test.m4a"'
```
- Retrieve audio
```
curl --request GET 'http://localhost/audio/user/1/phrase/1/m4a' -o './out.m4a'
```

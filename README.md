# api-test

A very simple backend service that can store and retrieve audio recordings associated with user IDs and practice phrases

The following enhancements can be done for a production scenario:
- Running the project on a cloud service and saving audio files in object storage like S3 buckets
- Make apps obtain secret connection parameters from an encrypted secret manager (ex. KMS) instead of environment variables
- Implementing user authentication - deny request to store/retrieve files unless authenticated with the same userID the operation is to be done on.
- Use SSL and run the app through https instead of http

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

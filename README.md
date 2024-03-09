# api-test
## requirements
```
go get -u github.com/gin-gonic/gin
go get -u github.com/lib/pq
```

## Test
```
docker-compose up -d
```

```
curl --request POST 'http://localhost/audio/user/1/phrase/2' --form 'audio_file=@"./test.m4a"'
curl --request GET 'http://localhost/audio/user/1/phrase/1/m4a' -o './out.m4a'
```

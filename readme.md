## Building Docker image
```
docker build -f Dockerfile.ubuntu2 . -t cosmid_i
docker create --name cosmid -i -v $(pwd):/cosmid -w /cosmid -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 cosmid_i
docker start cosmid
```

## Ignite Scaffold
```
docker exec -it cosmid bash
ignite scaffold chain github.com/serkanmulayim/cosmid 
cd cosmid
ignite chain serve
```

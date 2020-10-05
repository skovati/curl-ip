# curl-ip
Simple Gin web server written in Go that returns the client's IP address when cURLed. Currently live [here](https://ip.skovati.com)
### usage
```
curl -L ip.skovati.com
```
### installation
```
git clone https://github.com/skovati/curl-ip
cd curl-ip
./server
```
##### gophers can install from source
```
go build -o server main.go
```
##### also available on Docker Hub
```
docker pull skovati/curl-ip
docker run -p 8000:8000 -d skovati/curl-ip
```

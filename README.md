# curl-ip
Simple web server written in Go that returns the client's IP address when cURLed. Currently live [here](https://ip.skovati.com) running on GCP
### usage
```
curl -L ip.skovati.com
```
### installation
```
git clone https://github.com/skovati/curl-ip
cd curl-ip
go build -o server main.go
```
##### also available as a container image
```
docker run -p 8080:8080 -d ghcr.io/skovati/curl-ip
```

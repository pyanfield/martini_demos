## http://0value.com/build-a-restful-API-with-Martini

### run the server

	go run api.go data.go errors.go server.go encoding.go

### create certificate files

	go run /path/to/goroot/src/pkg/crypto/tls/generate_cert.go --host="localhost"

#### The `-k` option is required if you use a self-signed certificate. The `-u` option specifies the user:password, which in our case is simply token: (empty password). The `-i` option prints the whole response, including headers.

### get all albums

	$ curl -i -k -u token: "https://localhost:8001/albums"

### get album

	$ curl -i -k -u token: "https://localhost:8001/albums.text?band=Slayer"

### add album
	
	$ curl -i -k -u token: -X POST --data "band=Carcass&title=Heartwork&year=1994" "https://localhost:8001/albums"

	$ curl -i -k -u token: -X POST --data "band=Carcass&title=Heartwork&year=1994" "https://localhost:8001/albums.xml"

### update album

	$ curl -i -k -u token: -X PUT --data "band=Carcass&title=Heartwork&year=1993" "https://localhost:8001/albums/4"

### delete album

	$ curl -i -k -u token: -X DELETE "https://localhost:8001/albums/1"

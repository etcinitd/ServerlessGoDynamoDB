build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/device/addDevice device/addDevice.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/device/getDevice device/getDevice.go

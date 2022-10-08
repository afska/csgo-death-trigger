build:
	go build -trimpath -ldflags "-s -w" .

compress:
	upx -9 csgo-death-trigger.exe

run:
	go run .

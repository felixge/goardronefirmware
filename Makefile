deploy: arm ftp-upload telnet-run

arm: main.go
	GOOS=linux GOARCH=arm go build -o $@ $^

ftp-upload:
	echo put arm | ftp anonymous:anonymous@192.168.1.1

telnet-run:
	( sleep 0.2; echo "cd /data/video\nchmod +x arm\nclear\n./arm\nexit\n"; sleep 999 ) | telnet 192.168.1.1

.PHONY: ftp-upload telnet-run

DRONE_IP := 192.168.1.1

export GOOS := linux
export GOARCH := arm
export CGO_ENABLED=0

define run
curl -T $1 ftp://@$(DRONE_IP)/upload
(\
echo spawn telnet $(DRONE_IP);\
echo expect -re .*#;\
echo send \"cd /data/video\\\r\";\
echo expect -re .*#;\
echo send \"killall $1\\\r\";\
echo expect -re .*#;\
echo send \"rm $1\\\r\";\
echo expect -re .*#;\
echo send \"mv upload $1\\\r\";\
echo expect -re .*#;\
echo send \"chmod +x $1\\\r\";\
echo expect -re .*#;\
echo send \"./$1\\\r\";\
echo set timeout -1;\
echo expect -re .*#;\
) | expect
endef


debug-navdata: navdata/debug/server.go
	go build -o $@ $^
	@$(call run,$@)

.PHONY: debug-navdata

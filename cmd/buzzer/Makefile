clean: clean

clean:
	rm -rf buzzer

build:
	GOARM=7 GOARCH=arm GOOS=linux go build buzzer.go

upload:
	 scp buzzer pi@192.168.1.104:/home/pi/

publish: clean build upload

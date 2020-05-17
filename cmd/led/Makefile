clean: led_clean

led_clean:
	rm -rf led

led_build:
	GOARM=7 GOARCH=arm GOOS=linux go build led.go

led_upload:
	 scp led pi@192.168.1.104:/home/pi/

led_publish: led_clean led_build led_upload

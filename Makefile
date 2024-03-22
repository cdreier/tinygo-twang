startDebugServer:
	sudo openocd -f interface/cmsis-dap.cfg -f target/rp2040.cfg -c "adapter speed 5000"

picoflash:
	tinygo flash -target=pico . 
	
zeroflash:
	tinygo flash -target=waveshare-rp2040-zero .

monitor:
	tinygo monitor

emu:
	go run emulator/cmd/main.go
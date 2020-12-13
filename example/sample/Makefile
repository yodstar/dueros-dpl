TARGET := sample

# export CGO_ENABLED=0

.PHONY  : all clean clean rebuild install

all     : $(TARGET)

build   : all

rebuild : clean all

test    :

clean   :
	rm -fv ./build/$(TARGET)

deploy  :
	install -d -m0755 /usr/local/dueros/logs
	install -m0755 ./build/$(TARGET) /usr/local/dueros/$(TARGET)
	install -m0600 ./build/$(TARGET).conf /usr/local/dueros/$(TARGET).conf
	install -m0755 ./$(TARGET).service /usr/lib/systemd/system/$(TARGET).service

install :
	install -d -m0755 /usr/local/dueros/logs
	install -m0755 ./build/$(TARGET) /usr/local/dueros/$(TARGET)
	install -m0600 ./$(TARGET).conf /usr/local/dueros/$(TARGET).conf
	install -m0755 ./$(TARGET).service /usr/lib/systemd/system/$(TARGET).service

$(TARGET) :
	# go build -o ./build/$(TARGET) -race -ldflags "-w -s" $(TARGET)
	go build -o ./build/$(TARGET) -ldflags "-w -s" $(TARGET)

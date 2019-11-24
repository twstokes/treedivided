import serial

ser = serial.Serial('/dev/ttyUSB0', 115200)

# set colors for teams
payload = bytes([4, 1, 246, 103, 51, 82, 45, 128])
ser.write(payload)

payload = bytes([4, 2, 255, 255, 255, 115, 0, 10])
ser.write(payload)

# update scores
payload = bytes([2, 13, 7, 0, 0, 0, 0, 0])
ser.write(payload)

ser.close()

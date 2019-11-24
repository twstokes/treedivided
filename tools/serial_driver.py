import serial

ser = serial.Serial('/dev/ttyUSB0', 115200)

# set colors for teams
payload = bytes([4, 1, 246, 103, 51, 82, 45, 128])
ser.write(payload)

payload = bytes([4, 2, 115, 0, 10, 255, 255, 255])
ser.write(payload)

# show fanfare for team a
payload = bytes([1, 1, 0, 0, 0, 0, 0, 0])
ser.write(payload)

# update scores
payload = bytes([2, 13, 7, 0, 0, 0, 0, 0])
ser.write(payload)

# set winner
payload = bytes([3, 1, 0, 0, 0, 0, 0, 0])
ser.write(payload)

ser.close()

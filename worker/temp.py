import os
import glob
import time
import datetime
import json
import urllib2

os.system('modprobe w1-gpio')
os.system('modprobe w1-therm')

base_dir = '/sys/bus/w1/devices/'
device_folder = glob.glob(base_dir + '28*')[0]
device_file = device_folder + '/w1_slave'

def read_temp_raw():
    f = open(device_file, 'r')
    lines = f.readlines()
    f.close()
    return lines

def read_temp():
    lines = read_temp_raw()
    while lines[0].strip()[-3:] != 'YES':
        time.sleep(1)
        lines = read_temp_raw()
    equals_pos = lines[1].find('t=')
    if equals_pos != -1:
        temp_string = lines[1][equals_pos+2:]
        temp = float(temp_string) / 1000.0
        return temp

def send(data):
    req = urllib2.Request(os.getenv('TEMP_HOST') + '/temp')
    req.add_header('Content-Type', 'application/json')
    urllib2.urlopen(req, json.dumps(data))

while True:
    data = {
        "value": read_temp(),
        "time": datetime.datetime.now()
    }
    print(data)
    send(data)
    time.sleep(10)

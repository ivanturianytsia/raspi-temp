import os
import glob
import time
import datetime
import urllib.request
import json
import rfc3339

class Thermometer:
    def __init__(self):
        os.system('modprobe w1-gpio')
        os.system('modprobe w1-therm')

        base_dir = '/sys/bus/w1/devices/'
        device_folder = glob.glob(base_dir + '28*')[0]
        self.device_file = device_folder + '/w1_slave'

    def read_temp_raw(self):
        f = open(self.device_file, 'r')
        lines = f.readlines()
        f.close()
        return lines

    def read_temp(self):
        lines = self.read_temp_raw()
        while lines[0].strip()[-3:] != 'YES':
            time.sleep(1)
            lines = self.read_temp_raw()
        equals_pos = lines[1].find('t=')
        if equals_pos != -1:
            temp_string = lines[1][equals_pos+2:]
            temp = float(temp_string) / 1000.0
            return temp

def send(data):
    req = urllib.request.Request(os.getenv('TEMP_HOST') + '/temp')
    req.add_header('Content-Type', 'application/json; charset=utf-8')
    jsondata = json.dumps(data)
    jsondataasbytes = jsondata.encode('utf-8')
    req.add_header('Content-Length', len(jsondataasbytes))
    try:
        response = urllib.request.urlopen(req, jsondataasbytes)
    except urllib2.HTTPError as err:
        print(err)

therm = Thermometer()

sleepTime = int(os.getenv("FREQ"))

while True:
    data = {
        "measurement": "temp",
        "unit": "°C",
        "data": [{
            "value": therm.read_temp(),
            "datetime": rfc3339.rfc3339(datetime.datetime.now())
        }]
    }
    send(data)
    time.sleep(sleepTime)

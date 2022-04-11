package dbcon

import (
	"fmt"
	"smartpi/sensor"
	//influxdb "github.com/influxdata/influxdb-client-go.git"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)	

const dbhost = "http://localhost:8086"
const token = "qQQ5xDcMqauaLn3uBA8ZoiUXNEakqud7Db_k7KLzUAvVGcUxkO4msosfaOQYgoToiLeFwrCav_UA5atcW-c1tA=="

func Insert(s sensor.Sensor) {
	unit := s.Unit
	stype := s.Type
	measurement := s.Value
	client := influxdb2.NewClient(dbhost, token)
	defer client.Close()

	writeAPI := client.WriteAPI("smartpi", "smartpi-sensors")
	writeAPI.WriteRecord(fmt.Sprintf("stat,unit=%s %s=%f", unit, stype, measurement))
	writeAPI.Flush()
	client.Close()

}

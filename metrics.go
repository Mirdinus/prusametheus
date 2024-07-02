package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	tempBed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_temp_bed",
		Help: "Temperature of the printer bed",
	})
	tempNozzle = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_temp_nozzle",
		Help: "Temperature of the printer nozzle",
	})
	printSpeed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_print_speed",
		Help: "Print speed of the printer",
	})
	jobProgress = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_job_progress",
		Help: "Progress of the current print job",
	})
	axisZ = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_axis_z",
		Help: "Z axis position of the printer",
	})
	fanHotend = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_fan_hotend",
		Help: "Hotend fan speed",
	})
	fanPrint = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_fan_print",
		Help: "Print fan speed",
	})
	flow = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_flow",
		Help: "Flow rate",
	})
	speed = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_speed",
		Help: "Speed of the printer",
	})
	targetBedTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_target_bed_temp",
		Help: "Target temperature of the printer bed",
	})
	targetNozzleTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_target_nozzle_temp",
		Help: "Target temperature of the printer nozzle",
	})
	timePrinting = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_time_printing",
		Help: "Time spent printing",
	})
	timeRemaining = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_time_remaining",
		Help: "Time remaining for the print job",
	})
	jobName = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "printer_job_name",
		Help: "Name of the current print job",
	}, []string{"name"})
	nozzleDiameter = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "printer_nozzle_diameter",
		Help: "Diameter of the printer nozzle",
	})
	printerState = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "printer_state",
		Help: "State of the printer",
	}, []string{"state"})
)

func init() {
	prometheus.MustRegister(tempBed)
	prometheus.MustRegister(tempNozzle)
	prometheus.MustRegister(printSpeed)
	prometheus.MustRegister(jobProgress)
	prometheus.MustRegister(axisZ)
	prometheus.MustRegister(fanHotend)
	prometheus.MustRegister(fanPrint)
	prometheus.MustRegister(flow)
	prometheus.MustRegister(speed)
	prometheus.MustRegister(targetBedTemp)
	prometheus.MustRegister(targetNozzleTemp)
	prometheus.MustRegister(timePrinting)
	prometheus.MustRegister(timeRemaining)
	prometheus.MustRegister(jobName)
	prometheus.MustRegister(nozzleDiameter)
	prometheus.MustRegister(printerState)
}

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/icholy/digest"
)

var (
	printerEP = "/api/printer"
	statusEP  = "/api/v1/status"
	jobEP     = "/api/v1/job"
	infoEP    = "/api/v1/info"
)

func fetchData() {
	endpoints := []string{printerEP, statusEP, jobEP, infoEP}
	client := &http.Client{
		Transport: &digest.Transport{
			Username: AppConfig.Printer.Username,
			Password: AppConfig.Printer.Password,
		},
	}

	for _, endpoint := range endpoints {
		url := fmt.Sprintf("http://%s:%s%s", AppConfig.Printer.IP, AppConfig.Printer.Port, endpoint)
		resp, err := client.Get(url)
		if err != nil {
			log.Printf("Failed to fetch data from %s: %v", url, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Error response from %s: %s", url, resp.Status)
			continue
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body from %s: %v", url, err)
			continue
		}

		switch endpoint {
		case printerEP:
			var result PrinterResponse
			if err := json.Unmarshal(body, &result); err != nil {
				log.Printf("Failed to unmarshal JSON from %s: %v", url, err)
				continue
			}
			updatePrinterMetrics(&result)
		case statusEP:
			var result StatusResponse
			if err := json.Unmarshal(body, &result); err != nil {
				log.Printf("Failed to unmarshal JSON from %s: %v", url, err)
				continue
			}
			updateStatusMetrics(&result)
		case jobEP:
			var result JobResponse
			if err := json.Unmarshal(body, &result); err != nil {
				log.Printf("Failed to unmarshal JSON from %s: %v", url, err)
				continue
			}
			updateJobMetrics(&result)
		case infoEP:
			var result InfoResponse
			if err := json.Unmarshal(body, &result); err != nil {
				log.Printf("Failed to unmarshal JSON from %s: %v", url, err)
				continue
			}
			updateInfoMetrics(&result)
		}
	}
}

func updatePrinterMetrics(result *PrinterResponse) {
	tempBed.Set(result.Telemetry.TempBed)
	tempNozzle.Set(result.Telemetry.TempNozzle)
	printSpeed.Set(float64(result.Telemetry.PrintSpeed))
	printerState.WithLabelValues(result.State.Text).Set(1)
}

func updateStatusMetrics(result *StatusResponse) {
	tempBed.Set(result.Printer.TempBed)
	tempNozzle.Set(result.Printer.TempNozzle)
	axisZ.Set(result.Printer.AxisZ)
	fanHotend.Set(float64(result.Printer.FanHotend))
	fanPrint.Set(float64(result.Printer.FanPrint))
	flow.Set(float64(result.Printer.Flow))
	speed.Set(float64(result.Printer.Speed))
	targetBedTemp.Set(result.Printer.TargetBed)
	targetNozzleTemp.Set(result.Printer.TargetNozzle)
	timePrinting.Set(float64(result.Job.TimePrinting))
	timeRemaining.Set(float64(result.Job.TimeRemaining))
	jobProgress.Set(result.Job.Progress)
}

func updateJobMetrics(result *JobResponse) {
	jobName.WithLabelValues(result.File.DisplayName).Set(1)
	jobProgress.Set(result.Progress)
	timePrinting.Set(float64(result.TimePrinting))
	timeRemaining.Set(float64(result.TimeRemaining))
}

func updateInfoMetrics(result *InfoResponse) {
	nozzleDiameter.Set(result.NozzleDiameter)
}

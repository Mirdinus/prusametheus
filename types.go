package main

type PrinterResponse struct {
	State struct {
		Flags struct {
			Busy          bool   `json:"busy"`
			Cancelling    bool   `json:"cancelling"`
			ClosedOnError bool   `json:"closedOnError"`
			Error         bool   `json:"error"`
			LinkState     string `json:"link_state"`
			Operational   bool   `json:"operational"`
			Paused        bool   `json:"paused"`
			Pausing       bool   `json:"pausing"`
			Printing      bool   `json:"printing"`
			Ready         bool   `json:"ready"`
			SdReady       bool   `json:"sdReady"`
		} `json:"flags"`
		Text string `json:"text"`
	} `json:"state"`
	Telemetry struct {
		Material   string  `json:"material"`
		PrintSpeed int     `json:"print-speed"`
		TempBed    float64 `json:"temp-bed"`
		TempNozzle float64 `json:"temp-nozzle"`
		ZHeight    float64 `json:"z-height"`
	} `json:"telemetry"`
	Temperature struct {
		Bed struct {
			Actual float64 `json:"actual"`
			Offset float64 `json:"offset"`
			Target float64 `json:"target"`
		} `json:"bed"`
		Tool0 struct {
			Actual  float64 `json:"actual"`
			Display float64 `json:"display"`
			Offset  float64 `json:"offset"`
			Target  float64 `json:"target"`
		} `json:"tool0"`
	} `json:"temperature"`
}

type StatusResponse struct {
	Job struct {
		ID            int     `json:"id"`
		Progress      float64 `json:"progress"`
		TimePrinting  int     `json:"time_printing"`
		TimeRemaining int     `json:"time_remaining"`
	} `json:"job"`
	Printer struct {
		AxisZ        float64 `json:"axis_z"`
		FanHotend    int     `json:"fan_hotend"`
		FanPrint     int     `json:"fan_print"`
		Flow         int     `json:"flow"`
		Speed        int     `json:"speed"`
		State        string  `json:"state"`
		TargetBed    float64 `json:"target_bed"`
		TargetNozzle float64 `json:"target_nozzle"`
		TempBed      float64 `json:"temp_bed"`
		TempNozzle   float64 `json:"temp_nozzle"`
	} `json:"printer"`
	Storage struct {
		Name     string `json:"name"`
		Path     string `json:"path"`
		ReadOnly bool   `json:"read_only"`
	} `json:"storage"`
}

type JobResponse struct {
	File struct {
		DisplayName string  `json:"display_name"`
		MTimestamp  float64 `json:"m_timestamp"`
		Name        string  `json:"name"`
		Path        string  `json:"path"`
		Refs        struct {
			Download  string `json:"download"`
			Icon      string `json:"icon"`
			Thumbnail string `json:"thumbnail"`
		} `json:"refs"`
		Size float64 `json:"size"`
	} `json:"file"`
	ID            int     `json:"id"`
	Progress      float64 `json:"progress"`
	State         string  `json:"state"`
	TimePrinting  int     `json:"time_printing"`
	TimeRemaining int     `json:"time_remaining"`
}

type InfoResponse struct {
	Hostname         string  `json:"hostname"`
	MinExtrusionTemp float64 `json:"min_extrusion_temp"`
	Mmu              bool    `json:"mmu"`
	NozzleDiameter   float64 `json:"nozzle_diameter"`
	Serial           string  `json:"serial"`
}

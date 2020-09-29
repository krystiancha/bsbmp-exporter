package collectors

import (
	"log"
	"sync"
        "github.com/prometheus/client_golang/prometheus"
        client "github.com/david-igou/bsbmp-exporter/services"
)

var sensor client.Sensor

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type bsbmpCollector struct {

	mutex sync.RWMutex

	TemperatureC *prometheus.Desc
	HumidityRH *prometheus.Desc
	PressurePa *prometheus.Desc
	PressureMmHg *prometheus.Desc
	AltitudeM *prometheus.Desc
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func NewBsbmpCollector(c client.Sensor) *bsbmpCollector {
	sensor = c
	return &bsbmpCollector{
		TemperatureC: prometheus.NewDesc("TemperatureC",
			"The temperature in Celcius",
			nil, nil,
		),
		PressurePa: prometheus.NewDesc("PressurePa",
			"Pressure in Pascals",
			nil, nil,
		),
		PressureMmHg: prometheus.NewDesc("PressureMmHg",
			"Pressure in MmHg",
			nil, nil,
		),
		HumidityRH: prometheus.NewDesc("HumidityRH",
			"Relative humidity (Percent)",
			nil, nil,
		),
		AltitudeM: prometheus.NewDesc("AltitudeM",
			"Altitude in Meters above sea level",
			nil, nil,
		),
	}
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *bsbmpCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.TemperatureC
	ch <- collector.PressurePa
	ch <- collector.PressureMmHg
	ch <- collector.HumidityRH
	ch <- collector.AltitudeM
}

//Collect implements required collect function for all promehteus collectors
func (collector *bsbmpCollector) Collect(ch chan<- prometheus.Metric) {
//todo add mutex
	collector.mutex.Lock()
	defer collector.mutex.Unlock()
	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	resp, err := sensor.Poll() // I want the metrics
	if err != nil {
		log.Fatal(err)
	}

	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	ch <- prometheus.MustNewConstMetric(collector.TemperatureC, prometheus.CounterValue, float64(resp.TemperatureC))
	ch <- prometheus.MustNewConstMetric(collector.PressurePa, prometheus.CounterValue, float64(resp.PressurePa))
	ch <- prometheus.MustNewConstMetric(collector.PressureMmHg, prometheus.CounterValue, float64(resp.PressureMmHg))
	ch <- prometheus.MustNewConstMetric(collector.HumidityRH, prometheus.CounterValue, float64(resp.HumidityRH))
	ch <- prometheus.MustNewConstMetric(collector.AltitudeM, prometheus.CounterValue, float64(resp.AltitudeM))
}

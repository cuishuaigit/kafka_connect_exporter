package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

const nameSpace = "kafka_connect"

var (
	version       = "dev"
	versionUrl    = "https://github.com/wakeful/kafka_connect_exporter"
	pod           = os.Getenv("POD_NAME")
	svc           = os.Getenv("SVC")
	domain        = os.Getenv("DOMAIN")
	schema        = os.Getenv("SCHEMA")
	data          = []string{pod, svc, domain}
	couri         = schema + strings.Join(data, ".")
	showVersion   = flag.Bool("version", false, "show version and exit")
	listenAddress = flag.String("listen-address", ":8080", "Address on which to expose metrics.")
	metricsPath   = flag.String("telemetry-path", "/metrics", "Path under which to expose metrics.")
	scrapeURI     = flag.String("scrape-uri", couri, "URI on which to scrape kafka connect.")

	isConnectorRunning = prometheus.NewDesc(
		prometheus.BuildFQName(nameSpace, "connector", "state_running"),
		"is the connector running?",
		[]string{"connector", "state", "worker"}, nil)
	isConnecttorFailed = prometheus.NewDesc(
		prometheus.BuildFQName(nameSpace, "connector", "state_failed"),
		"is the connector failed?",
		[]string{"connector", "state", "worker"}, nil)
	isConnecttorPaused = prometheus.NewDesc(
		prometheus.BuildFQName(nameSpace, "connector", "state_paused"),
		"is the connector paused?",
		[]string{"connector", "state", "worker"}, nil)
	isConnectorUnassingned = prometheus.NewDesc(
		prometheus.BuildFQName(nameSpace, "conncetor", "state_unassingned"),
		"is the connector unassingned?",
		[]string{"connector", "state", "worker"}, nil)
	areConnectorTasksRunning = prometheus.NewDesc(
		prometheus.BuildFQName(nameSpace, "connector", "tasks_state_running"),
		"are connector tasks running?",
		[]string{"connector", "state", "worker_id", "id"}, nil)
	areConnectorTasksFailed = prometheus.NewDesc(
		prometheus.BuildFQName(nameSpace, "connector", "tasks_state_failed"),
		"are connector tasks failed?",
		[]string{"connector", "state", "worker_id", "id"}, nil)
	areConnectorTasksPaused = prometheus.NewDesc(
		prometheus.BuildFQName(nameSpace, "connector", "tasks_state_paused"),
		"are connector tasks paused?",
		[]string{"connector", "state", "worker_id", "id"}, nil)
	areConnectorTasksUnassingned = prometheus.NewDesc(
		prometheus.BuildFQName(nameSpace, "connector", "tasks_state_unassingned"),
		"are connector tasks unassingned?",
		[]string{"connector", "state", "worker_id", "id"}, nil)
)

type connectors []string

type status struct {
	Name      string    `json:"name"`
	Connector connector `json:"connector"`
	Tasks     []task    `json:"tasks"`
}

type connector struct {
	State    string `json:"state"`
	WorkerId string `json:"worker_id"`
}

type task struct {
	State    string  `json:"state"`
	Id       float64 `json:"id"`
	WorkerId string  `json:"worker_id"`
}

type Exporter struct {
	URI             string
	up              prometheus.Gauge
	connectorsCount prometheus.Gauge
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.up.Describe(ch)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {

	client := http.Client{
		Timeout: 3 * time.Second,
	}
	e.up.Set(0)

	response, err := client.Get(e.URI + "/connectors")
	if err != nil {
		log.Errorf("Can't scrape kafka connect: %v", err)
		return
	}
	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Errorf("Can't close connection to kafka connect: %v", err)
			return
		}
	}()

	output, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("Can't scrape kafka connect: %v", err)
		return
	}

	var connectorsList connectors
	if err := json.Unmarshal(output, &connectorsList); err != nil {
		log.Errorf("Can't scrape kafka connect: %v", err)
		return
	}

	e.up.Set(1)
	e.connectorsCount.Set(float64(len(connectorsList)))

	ch <- e.up
	ch <- e.connectorsCount

	for _, connector := range connectorsList {

		connectorStatusResponse, err := client.Get(e.URI + "/connectors/" + connector + "/status")
		if err != nil {
			log.Errorf("Can't get /status for: %v", err)
			continue
		}

		connectorStatusOutput, err := ioutil.ReadAll(connectorStatusResponse.Body)
		if err != nil {
			log.Errorf("Can't read Body for: %v", err)
			continue
		}

		var connectorStatus status
		if err := json.Unmarshal(connectorStatusOutput, &connectorStatus); err != nil {
			log.Errorf("Can't decode response for: %v", err)
			continue
		}

		var isRunning float64 = 0
		if strings.ToLower(connectorStatus.Connector.State) == "running" {
			isRunning = 1
		}

		ch <- prometheus.MustNewConstMetric(
			isConnectorRunning, prometheus.GaugeValue, isRunning,
			connectorStatus.Name, strings.ToLower(connectorStatus.Connector.State), connectorStatus.Connector.WorkerId,
		)

		var isFaild float64 = 0
		if strings.ToLower(connectorStatus.Connector.State) == "failed" {
			isFaild = 1
		}

		ch <- prometheus.MustNewConstMetric(
			isConnecttorFailed, prometheus.GaugeValue, isFaild,
			connectorStatus.Name, strings.ToLower(connectorStatus.Connector.State), connectorStatus.Connector.WorkerId,
		)

		var isPaused float64 = 0
		if strings.ToLower(connectorStatus.Connector.State) == "paused" {
			isPaused = 1
		}
		ch <- prometheus.MustNewConstMetric(
			isConnecttorPaused, prometheus.GaugeValue, isPaused,
			connectorStatus.Name, strings.ToLower(connectorStatus.Connector.State), connectorStatus.Connector.WorkerId,
		)

		var isUnassingned float64 = 0
		if strings.ToLower(connectorStatus.Connector.State) == "unassingned" {
			isUnassingned = 1
		}
		ch <- prometheus.MustNewConstMetric(
			isConnectorUnassingned, prometheus.GaugeValue, isUnassingned,
			connectorStatus.Name, strings.ToLower(connectorStatus.Connector.State), connectorStatus.Connector.WorkerId,
		)

		for _, connectorTask := range connectorStatus.Tasks {

			var isTaskRunning float64 = 0
			if strings.ToLower(connectorTask.State) == "running" {
				isTaskRunning = 1
			}

			ch <- prometheus.MustNewConstMetric(
				areConnectorTasksRunning, prometheus.GaugeValue, isTaskRunning,
				connectorStatus.Name, strings.ToLower(connectorTask.State), connectorTask.WorkerId, fmt.Sprintf("%d", int(connectorTask.Id)),
			)

			var isTaskFailed float64 = 0
			if strings.ToLower(connectorTask.State) == "failed" {
				isTaskFailed = 1
			}
			ch <- prometheus.MustNewConstMetric(
				areConnectorTasksFailed, prometheus.GaugeValue, isTaskFailed,
				connectorStatus.Name, strings.ToLower(connectorTask.State), connectorTask.WorkerId, fmt.Sprintf("%d", int(connectorTask.Id)),
			)

			var isTaskPaused float64 = 0
			if strings.ToLower(connectorTask.State) == "paused" {
				isTaskPaused = 1
			}
			ch <- prometheus.MustNewConstMetric(
				areConnectorTasksPaused, prometheus.GaugeValue, isTaskPaused,
				connectorStatus.Name, strings.ToLower(connectorTask.State), connectorTask.WorkerId, fmt.Sprintf("%d", int(connectorTask.Id)),
			)

			var isTaskUnassingned float64 = 0
			if strings.ToLower(connectorTask.State) == "unassingned" {
				isTaskUnassingned = 1
			}
			ch <- prometheus.MustNewConstMetric(
				areConnectorTasksUnassingned, prometheus.GaugeValue, isTaskUnassingned,
				connectorStatus.Name, strings.ToLower(connectorTask.State), connectorTask.WorkerId, fmt.Sprintf("%d", int(connectorTask.Id)),
			)
		}

		err = connectorStatusResponse.Body.Close()
		if err != nil {
			log.Errorf("Can't close connection to connector: %v", err)
		}
	}

	return
}

func NewExporter(uri string) *Exporter {
	log.Infoln("Collecting data from:", uri)

	return &Exporter{
		URI: uri,
		up: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: nameSpace,
			Name:      "up",
			Help:      "was the last scrape of kafka connect successful?",
		}),
		connectorsCount: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: nameSpace,
			Subsystem: "connectors",
			Name:      "count",
			Help:      "number of deployed connectors",
		}),
	}

}

var supportedSchema = map[string]bool{
	"http":  true,
	"https": true,
}

func main() {
	flag.Parse()

	if *showVersion {
		fmt.Printf("kafka_connect_exporter\n url: %s\n version: %s\n", versionUrl, version)
		os.Exit(2)
	}

	parseURI, err := url.Parse(*scrapeURI)
	if err != nil {
		log.Errorf("%v", err)
		os.Exit(1)
	}
	if !supportedSchema[parseURI.Scheme] {
		log.Error("schema not supported")
		os.Exit(1)
	}

	log.Infoln("Starting kafka_connect_exporter")

	prometheus.Unregister(prometheus.NewGoCollector())
	prometheus.Unregister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
	prometheus.MustRegister(NewExporter(*scrapeURI))

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, *metricsPath, http.StatusMovedPermanently)
	})

	log.Fatal(http.ListenAndServe(*listenAddress, nil))

}

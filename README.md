# Kafka connect exporter

A [Prometheus](https://prometheus.io/) exporter that collects [Kafka connect](https://docs.confluent.io/current/connect/index.html) metrics.

### Usage

```sh
$ ./kafka_connect_exporter -h
Usage of ./kafka_connect_exporter:
  -listen-address string
        Address on which to expose metrics. (default ":8080")
  -scrape-uri string
        URI on which to scrape kafka connect. (default "..")
  -telemetry-path string
        Path under which to expose metrics. (default "/metrics")
  -version
        show version and exit
```

## Metrics
```
# HELP kafka_connect_conncetor_state_unassingned is the connector unassingned?
# TYPE kafka_connect_conncetor_state_unassingned gauge
kafka_connect_conncetor_state_unassingned{connector="mysql-sale-binlog-hdfs-sink-v1",state="running",worker="10.18.4.240:8083"} 0
kafka_connect_conncetor_state_unassingned{connector="mysql-sale-system-bzj-binlog-hdfs-sink-v1",state="paused",worker="10.18.4.241:8083"} 0
kafka_connect_conncetor_state_unassingned{connector="source-mysql-sale-sale",state="running",worker="10.18.9.144:8083"} 0
kafka_connect_conncetor_state_unassingned{connector="source-mysql-sale-sale_uri_to_id",state="running",worker="10.18.4.240:8083"} 0
# HELP kafka_connect_connector_state_failed is the connector failed?
# TYPE kafka_connect_connector_state_failed gauge
kafka_connect_connector_state_failed{connector="mysql-sale-binlog-hdfs-sink-v1",state="running",worker="10.18.4.240:8083"} 0
kafka_connect_connector_state_failed{connector="mysql-sale-system-bzj-binlog-hdfs-sink-v1",state="paused",worker="10.18.4.241:8083"} 0
kafka_connect_connector_state_failed{connector="source-mysql-sale-sale",state="running",worker="10.18.9.144:8083"} 0
kafka_connect_connector_state_failed{connector="source-mysql-sale-sale_uri_to_id",state="running",worker="10.18.4.240:8083"} 0
# HELP kafka_connect_connector_state_paused is the connector paused?
# TYPE kafka_connect_connector_state_paused gauge
kafka_connect_connector_state_paused{connector="mysql-sale-binlog-hdfs-sink-v1",state="running",worker="10.18.4.240:8083"} 0
kafka_connect_connector_state_paused{connector="mysql-sale-system-bzj-binlog-hdfs-sink-v1",state="paused",worker="10.18.4.241:8083"} 1
kafka_connect_connector_state_paused{connector="source-mysql-sale-sale",state="running",worker="10.18.9.144:8083"} 0
kafka_connect_connector_state_paused{connector="source-mysql-sale-sale_uri_to_id",state="running",worker="10.18.4.240:8083"} 0
# HELP kafka_connect_connector_state_running is the connector running?
# TYPE kafka_connect_connector_state_running gauge
kafka_connect_connector_state_running{connector="mysql-sale-binlog-hdfs-sink-v1",state="running",worker="10.18.4.240:8083"} 1
kafka_connect_connector_state_running{connector="mysql-sale-system-bzj-binlog-hdfs-sink-v1",state="paused",worker="10.18.4.241:8083"} 0
kafka_connect_connector_state_running{connector="source-mysql-sale-sale",state="running",worker="10.18.9.144:8083"} 1
kafka_connect_connector_state_running{connector="source-mysql-sale-sale_uri_to_id",state="running",worker="10.18.4.240:8083"} 1
# HELP kafka_connect_connector_tasks_state_failed are connector tasks failed?
# TYPE kafka_connect_connector_tasks_state_failed gauge
kafka_connect_connector_tasks_state_failed{connector="mysql-sale-binlog-hdfs-sink-v1",id="0",state="running",worker_id="10.18.4.241:8083"} 0
kafka_connect_connector_tasks_state_failed{connector="mysql-sale-system-bzj-binlog-hdfs-sink-v1",id="0",state="paused",worker_id="10.18.9.144:8083"} 0
kafka_connect_connector_tasks_state_failed{connector="source-mysql-sale-sale",id="0",state="running",worker_id="10.18.4.240:8083"} 0
kafka_connect_connector_tasks_state_failed{connector="source-mysql-sale-sale_uri_to_id",id="0",state="running",worker_id="10.18.4.241:8083"} 0
# HELP kafka_connect_connector_tasks_state_paused are connector tasks paused?
# TYPE kafka_connect_connector_tasks_state_paused gauge
kafka_connect_connector_tasks_state_paused{connector="mysql-sale-binlog-hdfs-sink-v1",id="0",state="running",worker_id="10.18.4.241:8083"} 0
kafka_connect_connector_tasks_state_paused{connector="mysql-sale-system-bzj-binlog-hdfs-sink-v1",id="0",state="paused",worker_id="10.18.9.144:8083"} 1
kafka_connect_connector_tasks_state_paused{connector="source-mysql-sale-sale",id="0",state="running",worker_id="10.18.4.240:8083"} 0
kafka_connect_connector_tasks_state_paused{connector="source-mysql-sale-sale_uri_to_id",id="0",state="running",worker_id="10.18.4.241:8083"} 0
# HELP kafka_connect_connector_tasks_state_running are connector tasks running?
# TYPE kafka_connect_connector_tasks_state_running gauge
kafka_connect_connector_tasks_state_running{connector="mysql-sale-binlog-hdfs-sink-v1",id="0",state="running",worker_id="10.18.4.241:8083"} 1
kafka_connect_connector_tasks_state_running{connector="mysql-sale-system-bzj-binlog-hdfs-sink-v1",id="0",state="paused",worker_id="10.18.9.144:8083"} 0
kafka_connect_connector_tasks_state_running{connector="source-mysql-sale-sale",id="0",state="running",worker_id="10.18.4.240:8083"} 1
kafka_connect_connector_tasks_state_running{connector="source-mysql-sale-sale_uri_to_id",id="0",state="running",worker_id="10.18.4.241:8083"} 1
# HELP kafka_connect_connector_tasks_state_unassingned are connector tasks unassingned?
# TYPE kafka_connect_connector_tasks_state_unassingned gauge
kafka_connect_connector_tasks_state_unassingned{connector="mysql-sale-binlog-hdfs-sink-v1",id="0",state="running",worker_id="10.18.4.241:8083"} 0
kafka_connect_connector_tasks_state_unassingned{connector="mysql-sale-system-bzj-binlog-hdfs-sink-v1",id="0",state="paused",worker_id="10.18.9.144:8083"} 0
kafka_connect_connector_tasks_state_unassingned{connector="source-mysql-sale-sale",id="0",state="running",worker_id="10.18.4.240:8083"} 0
kafka_connect_connector_tasks_state_unassingned{connector="source-mysql-sale-sale_uri_to_id",id="0",state="running",worker_id="10.18.4.241:8083"} 0
# HELP kafka_connect_connectors_count number of deployed connectors
# TYPE kafka_connect_connectors_count gauge
kafka_connect_connectors_count 4
# HELP kafka_connect_up was the last scrape of kafka connect successful?
# TYPE kafka_connect_up gauge
kafka_connect_up 1
# HELP promhttp_metric_handler_requests_in_flight Current number of scrapes being served.
# TYPE promhttp_metric_handler_requests_in_flight gauge
promhttp_metric_handler_requests_in_flight 1
# HELP promhttp_metric_handler_requests_total Total number of scrapes by HTTP status code.
# TYPE promhttp_metric_handler_requests_total counter
promhttp_metric_handler_requests_total{code="200"} 0
promhttp_metric_handler_requests_total{code="500"} 0
promhttp_metric_handler_requests_total{code="503"} 0
```

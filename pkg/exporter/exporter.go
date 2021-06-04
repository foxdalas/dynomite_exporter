// Copyright 2020 Maxim Pogozhiy
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exporter

import (
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"time"
)

const (
	Namespace = "dynomite"
)

// Exporter collects metrics from a dynomite server.
type Exporter struct {
	address string
	timeout time.Duration
	logger  log.Logger

	up     *prometheus.Desc
	uptime *prometheus.Desc

	latency *prometheus.Desc

	payload_size *prometheus.Desc

	cross_region_rtt *prometheus.Desc

	cross_zone_latency *prometheus.Desc
	server_latency     *prometheus.Desc
	server_queue_wait  *prometheus.Desc

	cross_region_queue_wait *prometheus.Desc

	client_out_queue *prometheus.Desc

	server_in_queue  *prometheus.Desc
	server_out_queue *prometheus.Desc

	dnode_client_out_queue *prometheus.Desc

	peer_in_queue  *prometheus.Desc
	peer_out_queue *prometheus.Desc

	remote_peer_out_queue *prometheus.Desc
	remote_peer_in_queue  *prometheus.Desc

	alloc_msgs  *prometheus.Desc
	free_msgs   *prometheus.Desc
	alloc_mbufs *prometheus.Desc
	free_mbufs  *prometheus.Desc
	dyn_memory  *prometheus.Desc
}

// New returns an initialized exporter.
func New(server string, timeout time.Duration, logger log.Logger) *Exporter {
	return &Exporter{
		address: server,
		timeout: timeout,
		logger:  logger,
		up: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "up"),
			"Could the qynomite server be reached.",
			nil,
			nil,
		),
		uptime: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "uptime_seconds"),
			"Number of seconds since the server started.",
			[]string{"rack"},
			nil,
		),
		latency: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "latency"),
			"Server latency.",
			[]string{"rack","type"},
			nil,
		),
		payload_size: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "payload_size"),
			"Payload size.",
			[]string{"rack","type"},
			nil,
		),
		cross_region_rtt: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "cross_region_rtt"),
			"Cross region RTT.",
			[]string{"rack","type"},
			nil,
		),
		cross_zone_latency: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "cross_zone_latency"),
			"Cross region latency.",
			[]string{"rack","type"},
			nil,
		),
		server_latency: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "server_latency"),
			"Server latency.",
			[]string{"rack","type"},
			nil,
		),
		server_queue_wait: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "server_queue_wait"),
			"Server queue wait.",
			[]string{"rack","type"},
			nil,
		),
		cross_region_queue_wait: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "cross_region_queue_wait"),
			"Cross region queue wait.",
			[]string{"rack","type"},
			nil,
		),
		client_out_queue: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "client_out_queue"),
			"Client out queue.",
			[]string{"rack","type"},
			nil,
		),
		server_in_queue: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "server_in_queue"),
			"Server in queue.",
			[]string{"rack","type"},
			nil,
		),
		server_out_queue: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "server_out_queue"),
			"Server out queue.",
			[]string{"rack","type"},
			nil,
		),
		dnode_client_out_queue: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "dnode_client_out_queue"),
			"Dnode client out queue.",
			[]string{"rack","type"},
			nil,
		),
		peer_in_queue: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "peer_in_queue"),
			"Peer in queue.",
			[]string{"rack","type"},
			nil,
		),
		peer_out_queue: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "peer_out_queue"),
			"Peer out queue.",
			[]string{"rack","type"},
			nil,
		),
		remote_peer_in_queue: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "remote_peer_in_queue"),
			"Remote peer in queue.",
			[]string{"rack","type"},
			nil,
		),
		remote_peer_out_queue: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "remote_peer_out_queue"),
			"Remote peer out queue.",
			[]string{"rack","type"},
			nil,
		),
		alloc_msgs: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "alloc_msgs"),
			"The number of currently allocated messages.",
			[]string{"rack"},
			nil,
		),
		free_msgs: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "free_msgs"),
			"The number of currently free messages.",
			[]string{"rack"},
			nil,
		),
		alloc_mbufs: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "alloc_mbufs"),
			"The number of allocated mbufs.",
			[]string{"rack"},
			nil,
		),
		free_mbufs: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "free_mbufs"),
			"The number of free mbufs.",
			[]string{"rack"},
			nil,
		),
		dyn_memory: prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "dyn_memory"),
			"Dynomite memory usage.",
			[]string{"rack"},
			nil,
		),
	}
}

// Describe describes all the metrics exported by the dynomite exporter. It
// implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.up
	ch <- e.uptime
	ch <- e.latency
	ch <- e.payload_size
	ch <- e.cross_region_rtt
	ch <- e.cross_zone_latency
	ch <- e.server_latency
	ch <- e.server_queue_wait
	ch <- e.cross_region_queue_wait
	ch <- e.client_out_queue
	ch <- e.server_in_queue
	ch <- e.server_out_queue
	ch <- e.dnode_client_out_queue
	ch <- e.peer_in_queue
	ch <- e.peer_out_queue
	ch <- e.remote_peer_in_queue
	ch <- e.remote_peer_out_queue
	ch <- e.alloc_msgs
	ch <- e.free_msgs
	ch <- e.alloc_mbufs
	ch <- e.free_mbufs
	ch <- e.dyn_memory
}

// Collect fetches the statistics from the configured dynomite server, and
// delivers them as Prometheus metrics. It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	stats, err := GetMetrics(e.address)
	if err != nil {
		ch <- prometheus.MustNewConstMetric(e.up, prometheus.GaugeValue, 0)
		level.Error(e.logger).Log("msg", "Failed to connect to dynomite", "err", err)
		return
	}
	//TIMEOUT

	up := float64(1)

	if err := e.parseStats(ch, stats); err != nil {
		up = 0
	}

	ch <- prometheus.MustNewConstMetric(e.up, prometheus.GaugeValue, up)
}

func (e *Exporter) parseStats(ch chan<- prometheus.Metric, stats DynomiteMetrics) error {
	var parseError error

	ch <- prometheus.MustNewConstMetric(e.uptime, prometheus.CounterValue, float64(stats.Uptime), stats.Rack)

	ch <- prometheus.MustNewConstMetric(e.latency, prometheus.CounterValue, float64(stats.LatencyMax), stats.Rack, "max")
	ch <- prometheus.MustNewConstMetric(e.latency, prometheus.CounterValue, float64(stats.Latency999Th), stats.Rack, "999")
	ch <- prometheus.MustNewConstMetric(e.latency, prometheus.CounterValue, float64(stats.Latency99Th), stats.Rack, "99")
	ch <- prometheus.MustNewConstMetric(e.latency, prometheus.CounterValue, float64(stats.Latency95Th), stats.Rack, "95")
	ch <- prometheus.MustNewConstMetric(e.latency, prometheus.CounterValue, float64(stats.LatencyMean), stats.Rack, "50")

	ch <- prometheus.MustNewConstMetric(e.payload_size, prometheus.GaugeValue, float64(stats.PayloadSizeMax), stats.Rack, "max")
	ch <- prometheus.MustNewConstMetric(e.payload_size, prometheus.GaugeValue, float64(stats.PayloadSize999Th), stats.Rack, "999")
	ch <- prometheus.MustNewConstMetric(e.payload_size, prometheus.GaugeValue, float64(stats.PayloadSize99Th), stats.Rack, "99")
	ch <- prometheus.MustNewConstMetric(e.payload_size, prometheus.GaugeValue, float64(stats.PayloadSize95Th), stats.Rack, "95")
	ch <- prometheus.MustNewConstMetric(e.payload_size, prometheus.GaugeValue, float64(stats.PayloadSizeMean), stats.Rack, "50")

	ch <- prometheus.MustNewConstMetric(e.cross_region_rtt, prometheus.GaugeValue, float64(stats.Nine9CrossRegionRtt), stats.Rack, "99")
	ch <- prometheus.MustNewConstMetric(e.cross_region_rtt, prometheus.GaugeValue, float64(stats.AverageCrossRegionRtt), stats.Rack, "50")

	ch <- prometheus.MustNewConstMetric(e.cross_zone_latency, prometheus.GaugeValue, float64(stats.Nine9CrossZoneLatency), stats.Rack, "99")
	ch <- prometheus.MustNewConstMetric(e.cross_zone_latency, prometheus.GaugeValue, float64(stats.AverageCrossZoneLatency), stats.Rack, "50")

	ch <- prometheus.MustNewConstMetric(e.server_latency, prometheus.GaugeValue, float64(stats.Nine9ServerLatency), stats.Rack, "99")
	ch <- prometheus.MustNewConstMetric(e.server_latency, prometheus.GaugeValue, float64(stats.AverageServerLatency), stats.Rack, "50")

	ch <- prometheus.MustNewConstMetric(e.server_queue_wait, prometheus.GaugeValue, float64(stats.Nine9ServerQueueWait), stats.Rack, "99")
	ch <- prometheus.MustNewConstMetric(e.server_queue_wait, prometheus.GaugeValue, float64(stats.AverageServerQueueWait), stats.Rack, "50")

	ch <- prometheus.MustNewConstMetric(e.cross_region_queue_wait, prometheus.GaugeValue, float64(stats.Nine9CrossRegionQueueWait), stats.Rack, "99")
	ch <- prometheus.MustNewConstMetric(e.cross_region_queue_wait, prometheus.GaugeValue, float64(stats.AverageCrossRegionQueueWait), stats.Rack, "50")

	ch <- prometheus.MustNewConstMetric(e.client_out_queue, prometheus.GaugeValue, float64(stats.ClientOutQueue99), stats.Rack, "99")

	ch <- prometheus.MustNewConstMetric(e.server_in_queue, prometheus.GaugeValue, float64(stats.ServerInQueue99), stats.Rack, "99")
	ch <- prometheus.MustNewConstMetric(e.server_out_queue, prometheus.GaugeValue, float64(stats.ServerOutQueue99), stats.Rack, "99")

	ch <- prometheus.MustNewConstMetric(e.dnode_client_out_queue, prometheus.GaugeValue, float64(stats.DnodeClientOutQueue99), stats.Rack, "99")

	ch <- prometheus.MustNewConstMetric(e.peer_in_queue, prometheus.GaugeValue, float64(stats.PeerInQueue99), stats.Rack, "99")
	ch <- prometheus.MustNewConstMetric(e.peer_out_queue, prometheus.GaugeValue, float64(stats.PeerOutQueue99), stats.Rack, "99")

	ch <- prometheus.MustNewConstMetric(e.remote_peer_in_queue, prometheus.GaugeValue, float64(stats.RemotePeerInQueue99), stats.Rack, "99")
	ch <- prometheus.MustNewConstMetric(e.remote_peer_out_queue, prometheus.GaugeValue, float64(stats.RemotePeerOutQueue99), stats.Rack, "99")

	ch <- prometheus.MustNewConstMetric(e.alloc_msgs, prometheus.GaugeValue, float64(stats.AllocMsgs), stats.Rack)
	ch <- prometheus.MustNewConstMetric(e.free_msgs, prometheus.GaugeValue, float64(stats.FreeMsgs), stats.Rack)

	ch <- prometheus.MustNewConstMetric(e.alloc_mbufs, prometheus.GaugeValue, float64(stats.AllocMbufs), stats.Rack)
	ch <- prometheus.MustNewConstMetric(e.free_mbufs, prometheus.GaugeValue, float64(stats.FreeMbufs), stats.Rack)

	ch <- prometheus.MustNewConstMetric(e.dyn_memory, prometheus.GaugeValue, float64(stats.DynMemory), stats.Rack)


	return parseError
}

func GetMetrics(url string) (DynomiteMetrics, error) {
	var metrics DynomiteMetrics

	res, err := http.Get(url)
	if err != nil {
		return metrics, err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&metrics)
	return metrics, err
}

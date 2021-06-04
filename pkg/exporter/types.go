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

type DynomiteMetrics struct {
	Service                     string `json:"service"`
	Source                      string `json:"source"`
	Version                     string `json:"version"`
	Uptime                      int    `json:"uptime"`
	Timestamp                   int    `json:"timestamp"`
	Rack                        string `json:"rack"`
	Dc                          string `json:"dc"`
	LatencyMax                  int    `json:"latency_max"`
	Latency999Th                int    `json:"latency_999th"`
	Latency99Th                 int    `json:"latency_99th"`
	Latency95Th                 int    `json:"latency_95th"`
	LatencyMean                 int    `json:"latency_mean"`
	PayloadSizeMax              int    `json:"payload_size_max"`
	PayloadSize999Th            int    `json:"payload_size_999th"`
	PayloadSize99Th             int    `json:"payload_size_99th"`
	PayloadSize95Th             int    `json:"payload_size_95th"`
	PayloadSizeMean             int    `json:"payload_size_mean"`
	AverageCrossRegionRtt       int    `json:"average_cross_region_rtt"`
	Nine9CrossRegionRtt         int    `json:"99_cross_region_rtt"`
	AverageCrossZoneLatency     int    `json:"average_cross_zone_latency"`
	Nine9CrossZoneLatency       int    `json:"99_cross_zone_latency"`
	AverageServerLatency        int    `json:"average_server_latency"`
	Nine9ServerLatency          int    `json:"99_server_latency"`
	AverageCrossRegionQueueWait int    `json:"average_cross_region_queue_wait"`
	Nine9CrossRegionQueueWait   int    `json:"99_cross_region_queue_wait"`
	AverageCrossZoneQueueWait   int    `json:"average_cross_zone_queue_wait"`
	Nine9CrossZoneQueueWait     int    `json:"99_cross_zone_queue_wait"`
	AverageServerQueueWait      int    `json:"average_server_queue_wait"`
	Nine9ServerQueueWait        int    `json:"99_server_queue_wait"`
	ClientOutQueue99            int    `json:"client_out_queue_99"`
	ServerInQueue99             int    `json:"server_in_queue_99"`
	ServerOutQueue99            int    `json:"server_out_queue_99"`
	DnodeClientOutQueue99       int    `json:"dnode_client_out_queue_99"`
	PeerInQueue99               int    `json:"peer_in_queue_99"`
	PeerOutQueue99              int    `json:"peer_out_queue_99"`
	RemotePeerOutQueue99        int    `json:"remote_peer_out_queue_99"`
	RemotePeerInQueue99         int    `json:"remote_peer_in_queue_99"`
	AllocMsgs                   int    `json:"alloc_msgs"`
	FreeMsgs                    int    `json:"free_msgs"`
	AllocMbufs                  int    `json:"alloc_mbufs"`
	FreeMbufs                   int    `json:"free_mbufs"`
	DynMemory                   int    `json:"dyn_memory"`
	DynOMite                    struct {
		ClientEOF                  int   `json:"client_eof"`
		ClientErr                  int   `json:"client_err"`
		ClientConnections          int   `json:"client_connections"`
		ClientReadRequests         int   `json:"client_read_requests"`
		ClientWriteRequests        int   `json:"client_write_requests"`
		ClientDroppedRequests      int   `json:"client_dropped_requests"`
		ClientNonQuorumWResponses  int   `json:"client_non_quorum_w_responses"`
		ClientNonQuorumRResponses  int   `json:"client_non_quorum_r_responses"`
		ServerEjects               int   `json:"server_ejects"`
		DnodeClientEOF             int   `json:"dnode_client_eof"`
		DnodeClientErr             int   `json:"dnode_client_err"`
		DnodeClientConnections     int   `json:"dnode_client_connections"`
		DnodeClientInQueue         int   `json:"dnode_client_in_queue"`
		DnodeClientInQueueBytes    int   `json:"dnode_client_in_queue_bytes"`
		DnodeClientOutQueue        int   `json:"dnode_client_out_queue"`
		DnodeClientOutQueueBytes   int   `json:"dnode_client_out_queue_bytes"`
		PeerDroppedRequests        int   `json:"peer_dropped_requests"`
		PeerTimedoutRequests       int   `json:"peer_timedout_requests"`
		RemotePeerDroppedRequests  int   `json:"remote_peer_dropped_requests"`
		RemotePeerTimedoutRequests int   `json:"remote_peer_timedout_requests"`
		RemotePeerFailoverRequests int   `json:"remote_peer_failover_requests"`
		PeerEOF                    int   `json:"peer_eof"`
		PeerErr                    int   `json:"peer_err"`
		PeerTimedout               int   `json:"peer_timedout"`
		RemotePeerTimedout         int   `json:"remote_peer_timedout"`
		PeerConnections            int   `json:"peer_connections"`
		PeerForwardError           int   `json:"peer_forward_error"`
		PeerRequests               int   `json:"peer_requests"`
		PeerRequestBytes           int64 `json:"peer_request_bytes"`
		PeerResponses              int   `json:"peer_responses"`
		PeerResponseBytes          int   `json:"peer_response_bytes"`
		PeerEjectedAt              int64 `json:"peer_ejected_at"`
		PeerEjects                 int   `json:"peer_ejects"`
		PeerInQueue                int   `json:"peer_in_queue"`
		RemotePeerInQueue          int   `json:"remote_peer_in_queue"`
		PeerInQueueBytes           int   `json:"peer_in_queue_bytes"`
		RemotePeerInQueueBytes     int   `json:"remote_peer_in_queue_bytes"`
		PeerOutQueue               int   `json:"peer_out_queue"`
		RemotePeerOutQueue         int   `json:"remote_peer_out_queue"`
		PeerOutQueueBytes          int   `json:"peer_out_queue_bytes"`
		RemotePeerOutQueueBytes    int   `json:"remote_peer_out_queue_bytes"`
		PeerMismatchRequests       int   `json:"peer_mismatch_requests"`
		ForwardError               int   `json:"forward_error"`
		Fragments                  int   `json:"fragments"`
		StatsCount                 int   `json:"stats_count"`
	} `json:"dyn_o_mite"`
}

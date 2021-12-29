package kohaku

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

// TODO(v): mockDB を導入する

var (
	collectorTypeCodecJSON = `{
    "channel_id": "sora",
    "client_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "connection_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "id": "W8B607ZBG92PD9JTMS19BSTE18",
    "label": "WebRTC SFU Sora",
    "multistream": true,
    "node_name": "sora@127.0.0.1",
    "role": "sendrecv",
    "session_id": "JTYG1KGGPH2DKF86Y5B0GMWFSM",
    "simulcast": false,
    "spotlight": false,
    "stats": [
      {
        "channels": 2,
        "id": "RTCCodec_audio_NB1bb0_Inbound_109",
        "timestamp": 1640225763760.085,
        "type": "codec",
        "clockRate": 48000,
        "mimeType": "audio/opus",
        "payloadType": 109,
        "sdpFmtpLine": "minptime=10;useinbandfec=1",
        "transportId": "RTCTransport_data_1"
      }
    ],
    "timestamp": "2021-12-23T02:16:03.775161Z",
    "type": "connection.user-agent",
    "version": "2021.2.0"
  }`

	collectorTypeOutboundRTPJSON = `{
    "channel_id": "sora",
    "client_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "connection_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "id": "W8B607ZBG92PD9JTMS19BSTE18",
    "label": "WebRTC SFU Sora",
    "multistream": true,
    "node_name": "sora@127.0.0.1",
    "role": "sendrecv",
    "session_id": "JTYG1KGGPH2DKF86Y5B0GMWFSM",
    "simulcast": false,
    "spotlight": false,
    "stats": [
      {
        "framesEncoded": 892,
        "totalPacketSendDelay": 19.477,
        "mediaSourceId": "RTCVideoSource_10",
        "headerBytesSent": 26760,
        "transportId": "RTCTransport_data_1",
        "framesPerSecond": 31,
        "framesSent": 892,
        "id": "RTCOutboundRTPVideoStream_148236668",
        "totalEncodeTime": 1.532,
        "retransmittedBytesSent": 0,
        "keyFramesEncoded": 1,
        "frameWidth": 240,
        "qualityLimitationDurations": {
          "cpu": 0,
          "none": 30083,
          "other": 0,
          "bandwidth": 0
        },
        "packetsSent": 971,
        "nackCount": 0,
        "encoderImplementation": "libvpx",
        "trackId": "RTCMediaStreamTrack_sender_10",
        "qualityLimitationReason": "none",
        "type": "outbound-rtp",
        "firCount": 0,
        "codecId": "RTCCodec_video_WvsPAp_Outbound_120",
        "totalEncodedBytesTarget": 0,
        "kind": "video",
        "frameHeight": 160,
        "hugeFramesSent": 0,
        "pliCount": 0,
        "qpSum": 8808,
        "bytesSent": 722767,
        "timestamp": 1640225763760.085,
        "ssrc": 148236668,
        "remoteId": "RTCRemoteInboundRtpVideoStream_148236668",
        "retransmittedPacketsSent": 0,
        "mediaType": "video",
        "qualityLimitationResolutionChanges": 0
      }
    ],
    "timestamp": "2021-12-23T02:16:03.775161Z",
    "type": "connection.user-agent",
    "version": "2021.2.0"
  }`

	collectorTypeMediaSourceJSON = `{
    "channel_id": "sora",
    "client_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "connection_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "id": "W8B607ZBG92PD9JTMS19BSTE18",
    "label": "WebRTC SFU Sora",
    "multistream": true,
    "node_name": "sora@127.0.0.1",
    "role": "sendrecv",
    "session_id": "JTYG1KGGPH2DKF86Y5B0GMWFSM",
    "simulcast": false,
    "spotlight": false,
    "stats": [
      {
        "id": "RTCAudioSource_9",
        "kind": "audio",
        "timestamp": 1640225763760.085,
        "type": "media-source",
        "audioLevel": 0,
        "totalAudioEnergy": 0,
        "totalSamplesDuration": 30.090000000001904,
        "trackIdentifier": "9b36135b-f15f-4779-9aa2-d00609839d2d"
      },
      {
        "height": 160,
        "id": "RTCVideoSource_10",
        "kind": "video",
        "timestamp": 1640225763760.085,
        "type": "media-source",
        "width": 240,
        "frames": 894,
        "framesPerSecond": 31,
        "trackIdentifier": "425bc57b-5f59-4263-bcc5-579deb8c4d83"
      }
    ],
    "timestamp": "2021-12-23T02:16:03.775161Z",
    "type": "connection.user-agent",
    "version": "2021.2.0"
  }
`

	collectorTypeDataChannelJSON = `{
    "channel_id": "sora",
    "client_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "connection_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "id": "W8B607ZBG92PD9JTMS19BSTE18",
    "label": "WebRTC SFU Sora",
    "multistream": true,
    "node_name": "sora@127.0.0.1",
    "role": "sendrecv",
    "session_id": "JTYG1KGGPH2DKF86Y5B0GMWFSM",
    "simulcast": false,
    "spotlight": false,
    "stats": [
      {
        "id": "RTCDataChannel_17",
        "label": "signaling",
        "protocol": "",
        "state": "open",
        "timestamp": 1640225763760.085,
        "type": "data-channel",
        "bytesReceived": 0,
        "bytesSent": 0,
        "dataChannelIdentifier": 0,
        "messagesReceived": 0,
        "messagesSent": 0
      },
      {
        "id": "RTCDataChannel_18",
        "label": "notify",
        "protocol": "",
        "state": "open",
        "timestamp": 1640225763760.085,
        "type": "data-channel",
        "bytesReceived": 192,
        "bytesSent": 0,
        "dataChannelIdentifier": 2,
        "messagesReceived": 3,
        "messagesSent": 0
      },
      {
        "id": "RTCDataChannel_19",
        "label": "push",
        "protocol": "",
        "state": "open",
        "timestamp": 1640225763760.085,
        "type": "data-channel",
        "bytesReceived": 0,
        "bytesSent": 0,
        "dataChannelIdentifier": 4,
        "messagesReceived": 0,
        "messagesSent": 0
      },
      {
        "id": "RTCDataChannel_20",
        "label": "stats",
        "protocol": "",
        "state": "open",
        "timestamp": 1640225763760.085,
        "type": "data-channel",
        "bytesReceived": 28,
        "bytesSent": 0,
        "dataChannelIdentifier": 6,
        "messagesReceived": 1,
        "messagesSent": 0
      }
    ],
    "timestamp": "2021-12-23T02:16:03.775161Z",
    "type": "connection.user-agent",
    "version": "2021.2.0"
  }
`

	collectorTypeCandidatePairJSON = `{
    "channel_id": "sora",
    "client_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "connection_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "id": "W8B607ZBG92PD9JTMS19BSTE18",
    "label": "WebRTC SFU Sora",
    "multistream": true,
    "node_name": "sora@127.0.0.1",
    "role": "sendrecv",
    "session_id": "JTYG1KGGPH2DKF86Y5B0GMWFSM",
    "simulcast": false,
    "spotlight": false,
    "stats": [
      {
        "id": "RTCIceCandidatePair_eRplCBvi_JXPaEzOA",
        "priority": 179616219446525440,
        "state": "succeeded",
        "timestamp": 1640225763760.085,
        "type": "candidate-pair",
        "writable": true,
        "availableOutgoingBitrate": 1000000,
        "bytesDiscardedOnSend": 0,
        "bytesReceived": 5490,
        "bytesSent": 833847,
        "consentRequestsSent": 15,
        "currentRoundTripTime": 0.001,
        "localCandidateId": "RTCIceCandidate_eRplCBvi",
        "nominated": true,
        "packetsDiscardedOnSend": 0,
        "packetsReceived": 60,
        "packetsSent": 2520,
        "remoteCandidateId": "RTCIceCandidate_JXPaEzOA",
        "requestsReceived": 14,
        "requestsSent": 1,
        "responsesReceived": 16,
        "responsesSent": 14,
        "totalRoundTripTime": 0.032,
        "transportId": "RTCTransport_data_1"
      }
    ],
    "timestamp": "2021-12-23T02:16:03.775161Z",
    "type": "connection.user-agent",
    "version": "2021.2.0"
  }
`

	collectorTypeRemoteInboundRTPJSON = `{
    "channel_id": "sora",
    "client_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "connection_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "id": "W8B607ZBG92PD9JTMS19BSTE18",
    "label": "WebRTC SFU Sora",
    "multistream": true,
    "node_name": "sora@127.0.0.1",
    "role": "sendrecv",
    "session_id": "JTYG1KGGPH2DKF86Y5B0GMWFSM",
    "simulcast": false,
    "spotlight": false,
    "stats": [
      {
        "fractionLost": 0,
        "id": "RTCRemoteInboundRtpAudioStream_962078423",
        "kind": "audio",
        "ssrc": 962078423,
        "timestamp": 1640225763758.615,
        "type": "remote-inbound-rtp",
        "codecId": "RTCCodec_audio_NB1bb0_Outbound_109",
        "jitter": 0.0021041666666666665,
        "localId": "RTCOutboundRTPAudioStream_962078423",
        "packetsLost": 0,
        "roundTripTime": 0.002,
        "roundTripTimeMeasurements": 6,
        "totalRoundTripTime": 0.009,
        "transportId": "RTCTransport_data_1"
      },
      {
        "fractionLost": 0,
        "id": "RTCRemoteInboundRtpVideoStream_148236668",
        "kind": "video",
        "ssrc": 148236668,
        "timestamp": 1640225763393.525,
        "type": "remote-inbound-rtp",
        "codecId": "RTCCodec_video_WvsPAp_Outbound_120",
        "jitter": 0.0017111111111111112,
        "localId": "RTCOutboundRTPVideoStream_148236668",
        "packetsLost": 0,
        "roundTripTime": 0.003,
        "roundTripTimeMeasurements": 37,
        "totalRoundTripTime": 0.059,
        "transportId": "RTCTransport_data_1"
      }
    ],
    "timestamp": "2021-12-23T02:16:03.775161Z",
    "type": "connection.user-agent",
    "version": "2021.2.0"
  }
`

	collectorTypeTransportJSON = `{
    "channel_id": "sora",
    "client_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "connection_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "id": "W8B607ZBG92PD9JTMS19BSTE18",
    "label": "WebRTC SFU Sora",
    "multistream": true,
    "node_name": "sora@127.0.0.1",
    "role": "sendrecv",
    "session_id": "JTYG1KGGPH2DKF86Y5B0GMWFSM",
    "simulcast": false,
    "spotlight": false,
    "stats": [
      {
        "id": "RTCTransport_data_1",
        "timestamp": 1640225763760.085,
        "type": "transport",
        "bytesReceived": 5490,
        "bytesSent": 833847,
        "dtlsCipher": "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
        "dtlsState": "connected",
        "localCertificateId": "RTCCertificate_66:F6:14:8E:B3:3E:C1:44:D0:DB:3C:2B:1C:35:7E:F4:4B:3A:6C:87:AD:E2:09:06:7C:EB:5B:DD:62:6F:36:40",
        "packetsReceived": 60,
        "packetsSent": 2520,
        "remoteCertificateId": "RTCCertificate_A9:4A:03:B1:A9:66:46:EC:AD:03:73:D8:1E:99:46:06:5C:56:E9:00:AC:A5:F9:7C:50:8C:28:16:2A:E5:BF:07",
        "selectedCandidatePairChanges": 1,
        "selectedCandidatePairId": "RTCIceCandidatePair_eRplCBvi_JXPaEzOA",
        "srtpCipher": "AEAD_AES_128_GCM",
        "tlsVersion": "FEFD"
      }
    ],
    "timestamp": "2021-12-23T02:16:03.775161Z",
    "type": "connection.user-agent",
    "version": "2021.2.0"
  }
`

	collectorTypeErlangVMMemoryJSON = `{
    "id": "NXXP9WDGCH72SC7SJAZ2A28318",
    "label": "WebRTC SFU Sora",
    "node_name": "sora@127.0.0.1",
    "stats": [
      {
        "atom": 1270065,
        "atom_used": 1243247,
        "binary": 340376,
        "code": 33580550,
        "ets": 2661680,
        "processes": 18702696,
        "processes_used": 18702696,
        "system": 51684512,
        "total": 70387208,
        "type": "erlang-vm-memory"
      }
    ],
    "timestamp": "2021-12-23T02:25:07.471546Z",
    "type": "node.erlang-vm",
    "version": "2021.2.0"
  }
`
)

var (
	invalidConnectionIDLengthJSON = `{
    "channel_id": "sora",
    "client_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "connection_id": "QJ253E85SH1C170WQSPYJGFHCR=",
    "id": "W8B607ZBG92PD9JTMS19BSTE18",
    "label": "WebRTC SFU Sora",
    "multistream": true,
    "node_name": "sora@127.0.0.1",
    "role": "sendrecv",
    "session_id": "JTYG1KGGPH2DKF86Y5B0GMWFSM",
    "simulcast": false,
    "spotlight": false,
    "stats": [
      {
        "channels": 2,
        "id": "RTCCodec_audio_NB1bb0_Inbound_109",
        "timestamp": 1640225763760.085,
        "type": "codec",
        "clockRate": 48000,
        "mimeType": "audio/opus",
        "payloadType": 109,
        "sdpFmtpLine": "minptime=10;useinbandfec=1",
        "transportId": "RTCTransport_data_1"
      }
    ],
    "timestamp": "2021-12-23T02:16:03.775161Z",
    "type": "connection.user-agent",
    "version": "2021.2.0"
  }`

	unexpectedTypeJSON = `{
    "channel_id": "sora",
    "client_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "connection_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "id": "W8B607ZBG92PD9JTMS19BSTE18",
    "label": "WebRTC SFU Sora",
    "multistream": true,
    "node_name": "sora@127.0.0.1",
    "role": "sendrecv",
    "session_id": "JTYG1KGGPH2DKF86Y5B0GMWFSM",
    "simulcast": false,
    "spotlight": false,
    "stats": [
      {
        "channels": 2,
        "id": "RTCCodec_audio_NB1bb0_Inbound_109",
        "timestamp": 1640225763760.085,
        "type": "codec",
        "clockRate": 48000,
        "mimeType": "audio/opus",
        "payloadType": 109,
        "sdpFmtpLine": "minptime=10;useinbandfec=1",
        "transportId": "RTCTransport_data_1"
      }
    ],
    "timestamp": "2021-12-23T02:16:03.775161Z",
    "type": "connection.unexpected_type",
    "version": "2021.2.0"
  }`

	missingTimestampJSON = `{
    "channel_id": "sora",
    "client_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "connection_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "id": "W8B607ZBG92PD9JTMS19BSTE18",
    "label": "WebRTC SFU Sora",
    "multistream": true,
    "node_name": "sora@127.0.0.1",
    "role": "sendrecv",
    "session_id": "JTYG1KGGPH2DKF86Y5B0GMWFSM",
    "simulcast": false,
    "spotlight": false,
    "stats": [
      {
        "channels": 2,
        "id": "RTCCodec_audio_NB1bb0_Inbound_109",
        "timestamp": 1640225763760.085,
        "type": "codec",
        "clockRate": 48000,
        "mimeType": "audio/opus",
        "payloadType": 109,
        "sdpFmtpLine": "minptime=10;useinbandfec=1",
        "transportId": "RTCTransport_data_1"
      }
    ],
    "type": "connection.user-agent",
    "version": "2021.2.0"
  }`

	invalidChannelIDLengthJSON = `{
    "channel_id": "2QB23E50YD6FKEFG9GW2TX86RC2QB23E50YD6FKEFG9GW2TX86RC2QB23E50YD6FKEFG9GW2TX86RC2QB23E50YD6FKEFG9GW2TX86RC2QB23E50YD6FKEFG9GW2TX86RC2QB23E50YD6FKEFG9GW2TX86RC2QB23E50YD6FKEFG9GW2TX86RC2QB23E50YD6FKEFG9GW2TX86RC2QB23E50YD6FKEFG9GW2TX86RC2QB23E50YD6FKEFG9GW2TX",
    "client_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "connection_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "id": "W8B607ZBG92PD9JTMS19BSTE18",
    "label": "WebRTC SFU Sora",
    "multistream": true,
    "node_name": "sora@127.0.0.1",
    "role": "sendrecv",
    "session_id": "JTYG1KGGPH2DKF86Y5B0GMWFSM",
    "simulcast": false,
    "spotlight": false,
    "stats": [
      {
        "channels": 2,
        "id": "RTCCodec_audio_NB1bb0_Inbound_109",
        "timestamp": 1640225763760.085,
        "type": "codec",
        "clockRate": 48000,
        "mimeType": "audio/opus",
        "payloadType": 109,
        "sdpFmtpLine": "minptime=10;useinbandfec=1",
        "transportId": "RTCTransport_data_1"
      }
    ],
    "timestamp": "2021-12-23T02:16:03.775161Z",
    "type": "connection.user-agent",
    "version": "2021.2.0"
  }`

	missingMultistreamJSON = `{
    "channel_id": "sora",
    "client_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "connection_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "id": "W8B607ZBG92PD9JTMS19BSTE18",
    "label": "WebRTC SFU Sora",
    "node_name": "sora@127.0.0.1",
    "role": "sendrecv",
    "session_id": "JTYG1KGGPH2DKF86Y5B0GMWFSM",
    "simulcast": false,
    "spotlight": false,
    "stats": [
      {
        "channels": 2,
        "id": "RTCCodec_audio_NB1bb0_Inbound_109",
        "timestamp": 1640225763760.085,
        "type": "codec",
        "clockRate": 48000,
        "mimeType": "audio/opus",
        "payloadType": 109,
        "sdpFmtpLine": "minptime=10;useinbandfec=1",
        "transportId": "RTCTransport_data_1"
      }
    ],
    "timestamp": "2021-12-23T02:16:03.775161Z",
    "type": "connection.user-agent",
    "version": "2021.2.0"
  }`

	unexpectedErlangVMTypeJSON = `{
    "id": "NXXP9WDGCH72SC7SJAZ2A28318",
    "label": "WebRTC SFU Sora",
    "node_name": "sora@127.0.0.1",
    "stats": [
      {
        "atom": 1270065,
        "atom_used": 1243247,
        "binary": 340376,
        "code": 33580550,
        "ets": 2661680,
        "processes": 18702696,
        "processes_used": 18702696,
        "system": 51684512,
        "total": 70387208,
        "type": "unexpected_type"
      }
    ],
    "timestamp": "2021-12-23T02:25:07.471546Z",
    "type": "node.erlang-vm",
    "version": "2021.2.0"
  }
`
	unexpectedStatsTypeJSON = `{
    "channel_id": "sora",
    "client_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "connection_id": "QJ253E85SH1C170WQSPYJGFHCR",
    "id": "W8B607ZBG92PD9JTMS19BSTE18",
    "label": "WebRTC SFU Sora",
    "multistream": true,
    "node_name": "sora@127.0.0.1",
    "role": "sendrecv",
    "session_id": "JTYG1KGGPH2DKF86Y5B0GMWFSM",
    "simulcast": false,
    "spotlight": false,
    "stats": [
      {
        "channels": 2,
        "id": "RTCCodec_audio_NB1bb0_Inbound_109",
        "timestamp": 1640225763760.085,
        "type": "unexpected_type",
        "clockRate": 48000,
        "mimeType": "audio/opus",
        "payloadType": 109,
        "sdpFmtpLine": "minptime=10;useinbandfec=1",
        "transportId": "RTCTransport_data_1"
      }
    ],
    "timestamp": "2021-12-23T02:16:03.775161Z",
    "type": "connection.user-agent",
    "version": "2021.2.0"
  }`
)

const (
	connStr          = "postgres://%s:%s@%s/%s?sslmode=disable"
	postgresUser     = "postgres"
	postgresPassword = "password"
	postgresDB       = "kohakutest"

	channelID    = "sora"
	connectionID = "QJ253E85SH1C170WQSPYJGFHCR"
	clientID     = "QJ253E85SH1C170WQSPYJGFHCR"
)

var (
	pgPool *pgxpool.Pool
	server *Server
)

func getStatsType(table, connectionID string) (*string, error) {
	selectSQL := fmt.Sprintf("SELECT stats_type FROM %s WHERE sora_connection_id=$1", table)
	row := pgPool.QueryRow(context.Background(), selectSQL, connectionID)

	var statsType string
	if err := row.Scan(&statsType); err != nil {
		return nil, err
	}

	return &statsType, nil
}

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		panic(err)
	}

	pwd, _ := os.Getwd()

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "timescale/timescaledb",
		Tag:        "2.5.1-pg14",
		Env: []string{
			"POSTGRES_PASSWORD=" + postgresPassword,
			"POSTGRES_USER=" + postgresUser,
			"POSTGRES_DB=" + postgresDB,
			"listen_addresses = '*'",
		},
		Mounts: []string{
			pwd + "/script/schema.sql:/docker-entrypoint-initdb.d/schema.sql",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	if err != nil {
		panic(err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	kohakuDBURL := fmt.Sprintf(connStr, postgresUser, postgresPassword, hostAndPort, postgresDB)

	resource.Expire(60)
	pool.MaxWait = 60 * time.Second
	if err = pool.Retry(func() error {
		config, err := pgxpool.ParseConfig(kohakuDBURL)
		if err != nil {
			return err
		}
		pgPool, err = pgxpool.ConnectConfig(context.Background(), config)
		if err != nil {
			return err
		}

		return pgPool.Ping(context.Background())
	}); err != nil {
		panic(err)
	}

	server = NewServer(&KohakuConfig{}, pgPool)

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		panic(err)
	}

	os.Exit(code)
}

func TestTypeOutboundRTPCollector(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(collectorTypeOutboundRTPJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "connection.user-agent")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	assert.Equal(t, http.StatusNoContent, c.Writer.Status())

	statsType, err := getStatsType("rtc_outbound_rtp_stream_stats", connectionID)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "outbound-rtp", *statsType)
}

func TestTypeCodecCollector(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(collectorTypeCodecJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "connection.user-agent")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	assert.Equal(t, http.StatusNoContent, c.Writer.Status())

	statsType, err := getStatsType("rtc_codec_stats", connectionID)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "codec", *statsType)
}

func TestTypeMediaSourceCollector(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(collectorTypeMediaSourceJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "connection.user-agent")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	assert.Equal(t, http.StatusNoContent, c.Writer.Status())

	statsType, err := getStatsType("rtc_audio_source_stats", connectionID)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "media-source", *statsType)
}

func TestTypeDataChannelCollector(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(collectorTypeDataChannelJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "connection.user-agent")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	assert.Equal(t, http.StatusNoContent, c.Writer.Status())

	statsType, err := getStatsType("rtc_data_channel_stats", connectionID)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "data-channel", *statsType)
}

func TestTypeCandidatePairCollector(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(collectorTypeCandidatePairJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "connection.user-agent")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	assert.Equal(t, http.StatusNoContent, c.Writer.Status())

	statsType, err := getStatsType("rtc_ice_candidate_pair_stats", connectionID)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "candidate-pair", *statsType)
}

func TestTypeRemoteInboundRTPCollector(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(collectorTypeRemoteInboundRTPJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "connection.user-agent")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	assert.Equal(t, http.StatusNoContent, c.Writer.Status())

	statsType, err := getStatsType("rtc_remote_inbound_rtp_stream_stats", connectionID)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "remote-inbound-rtp", *statsType)
}

func TestTypeTransportCollector(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(collectorTypeTransportJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "connection.user-agent")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	assert.Equal(t, http.StatusNoContent, c.Writer.Status())

	statsType, err := getStatsType("rtc_transport_stats", connectionID)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "transport", *statsType)
}

func TestInvalidConnectionIDLength(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(invalidConnectionIDLengthJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "connection.user-agent")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	resp := rec.Result()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	assert.NotEmpty(t, body)
}

func TestUnexpectedType(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(unexpectedTypeJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "connection.unexpected_type")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	resp := rec.Result()

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	assert.Empty(t, body)
}

func TestMissingTimestamp(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(missingTimestampJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "connection.user-agent")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	resp := rec.Result()

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, `{"error":"Key: 'soraConnectionStats.soraStats.Timestamp' Error:Field validation for 'Timestamp' failed on the 'required' tag"}`, string(body))
}

func TestInvalidChannelIDLength(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(invalidChannelIDLengthJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "connection.user-agent")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	resp := rec.Result()

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, `{"error":"Key: 'soraConnectionStats.ChannelID' Error:Field validation for 'ChannelID' failed on the 'maxb' tag"}`, string(body))
}

func TestMissingMultistream(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(missingMultistreamJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "connection.user-agent")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	resp := rec.Result()

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, `{"error":"Key: 'soraConnectionStats.Multistream' Error:Field validation for 'Multistream' failed on the 'required' tag"}`, string(body))
}

func TestUnexpectedStatsType(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(unexpectedStatsTypeJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "connection.user-agent")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	resp := rec.Result()

	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, `{"error":"unexpected rtcStats.Type: unexpected_type"}`, string(body))
}

func TestTypeErlangVMMemoryCollector(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(collectorTypeErlangVMMemoryJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "node.erlang-vm")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	assert.Equal(t, http.StatusNoContent, c.Writer.Status())

	// TODO: 関数化
	selectSQL := "SELECT stats_type FROM erlang_vm_memory_stats WHERE sora_label=$1"
	row := pgPool.QueryRow(context.Background(), selectSQL, "WebRTC SFU Sora")
	var statsType string
	if err := row.Scan(&statsType); err != nil {
		panic(err)
	}

	assert.Equal(t, "erlang-vm-memory", statsType)
}

func TestUnexpectedErlangVMType(t *testing.T) {
	// Setup
	req := httptest.NewRequest(http.MethodPost, "/collector", strings.NewReader(unexpectedErlangVMTypeJSON))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-sora-stats-exporter-type", "node.erlang-vm")
	req.Proto = "HTTP/2.0"
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req

	// Assertions
	server.collector(c)
	assert.Equal(t, http.StatusBadRequest, c.Writer.Status())
}

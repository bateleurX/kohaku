package kohaku

import "time"

// 使ってない
// type soraNode struct {
// 	Timestamp time.Time `db:"timestamp"`
//
// 	Label    string `db:"label"`
// 	Version  string `db:"version"`
// 	NodeName string `db:"node_name"`
// }

// 使ってない
// type soraConnection struct {
// 	soraNode
//
// 	Multistream bool `db:"multistream"`
// 	Simulcast   bool `db:"simulcast"`
// 	Spotlight   bool `db:"spotlight"`
//
// 	Role         string `db:"role"`
// 	ChannelID    string `db:"channel_id"`
// 	SessionID    string `db:"session_id"`
// 	ClientID     string `db:"client_id"`
// 	ConnectionID string `db:"connection_id"`
// }

type rtc struct {
	Time *time.Time `db:"time"`

	ConnectionID string `db:"sora_connection_id"`
}

type rtcCodec struct {
	rtc
	rtcCodecStats
}

type rtcInboundRTPStream struct {
	rtc
	rtcInboundRTPStreamStats
}

type rtcRemoteInboundRTPStream struct {
	rtc
	rtcRemoteInboundRTPStreamStats
}

type rtcOutboundRTPStream struct {
	rtc
	rtcOutboundRTPStreamStats
}

type rtcRemoteOutboundRTPStream struct {
	rtc
	rtcRemoteOutboundRTPStreamStats
}

type rtcAuidoSource struct {
	rtc
	rtcAudioSourceStats
}

type rtcVideoSource struct {
	rtc
	rtcVideoSourceStats
}

type rtcDataChannel struct {
	rtc
	rtcDataChannelStats
}

type rtcTransport struct {
	rtc
	rtcTransportStats
}

type rtcIceCandidate struct {
	rtc
	rtcIceCandidateStats
}

type rtcIceCandidatePair struct {
	rtc
	rtcIceCandidatePairStats
}

package kohaku

// https://www.w3.org/TR/webrtc-stats/

// https://www.w3.org/TR/webrtc-stats/#dom-rtcstats
type rtcStats struct {
	Timestamp *float64 `json:"timestamp" validate:"required" db:"stats_timestamp"` // required DOMHighResTimeStamp timestamp;
	Type      *string  `json:"type" validate:"required" db:"stats_type"`           // required RTCStatsType        type;
	ID        *string  `json:"id" validate:"required" db:"stats_id"`               // required DOMString           id;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcrtpstreamstats
type rtcRTPStreamStats struct {
	rtcStats

	SSRC        *uint32 `json:"ssrc" validate:"required" db:"ssrc"` // required unsigned long       ssrc;
	Kind        *string `json:"kind" validate:"required" db:"kind"` // required DOMString           kind;
	TransportID *string `json:"transportId" db:"transport_id"`      // DOMString           transportId;
	CodecID     *string `json:"codecId" db:"codec_id"`              // DOMString           codecId;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtccodecstats
type rtcCodecStats struct {
	rtcStats

	PayloadType *uint32 `json:"payloadType" validate:"required" db:"payload_type"` // required unsigned long payloadType;
	CodecType   *string `json:"codecType" db:"codec_type"`                         // RTCCodecType  codecType;
	TransportID *string `json:"transportId" validate:"required" db:"transport_id"` // required DOMString     transportId;
	MimeType    *string `json:"mimeType" validate:"required" db:"mime_type"`       // required DOMString     mimeType;
	ClockRate   *uint32 `json:"clockRate" db:"clock_rate"`                         // unsigned long clockRate;
	Channels    *uint32 `json:"channels" db:"channels"`                            // unsigned long channels;
	SdpFmtpLine *string `json:"sdpFmtpLine" db:"sdp_fmtp_line"`                    // DOMString     sdpFmtpLine;
}

// https://www.w3.org/TR/webrtc-stats/#receivedrtpstats-dict*
type rtcReceivedRTPStreamStats struct {
	rtcRTPStreamStats

	PacketsReceived       *uint64  `json:"packetsReceived" db:"packets_received"`              // unsigned long long   packetsReceived;
	PacketsLost           *int64   `json:"packetsLost" db:"packets_lost"`                      // long long            packetsLost;
	Jitter                *float64 `json:"jitter" db:"jitter"`                                 // double               jitter;
	PacketsDiscarded      *uint64  `json:"packetsDiscarded" db:"packets_discarded"`            // unsigned long long   packetsDiscarded;
	PacketsRepaired       *uint64  `json:"packetsRepaired" db:"packets_repaired"`              // unsigned long long   packetsRepaired;
	BurstPacketsLost      *uint64  `json:"burstPacketsLost" db:"burst_packets_lost"`           // unsigned long long   burstPacketsLost;
	BurstPacketsDiscarded *uint64  `json:"burstPacketsDiscarded" db:"burst_packets_discarded"` // unsigned long long   burstPacketsDiscarded;
	BurstLossCount        *uint32  `json:"burstLossCount" db:"burst_loss_count"`               // unsigned long        burstLossCount;
	BurstDiscardCount     *uint32  `json:"burstDiscardCount" db:"burst_discard_count"`         // unsigned long        burstDiscardCount;
	BurstLossRate         *float64 `json:"burstLossRate" db:"burst_loss_rate"`                 // double               burstLossRate;
	BurstDiscardRate      *float64 `json:"burstDiscardRate" db:"burst_discard_rate"`           // double               burstDiscardRate;
	GapLossRate           *float64 `json:"gapLossRate" db:"gap_loss_rate"`                     // double               gapLossRate;
	GapDiscardRate        *float64 `json:"gapDiscardRate" db:"gap_discard_rate"`               // double               gapDiscardRate;
	FramesDropped         *uint32  `json:"framesDropped" db:"frames_dropped"`                  // unsigned long        framesDropped;
	PartialFramesLost     *uint32  `json:"partialFramesLost" db:"partial_frames_lost"`         // unsigned long        partialFramesLost;
	FullFramesLost        *uint32  `json:"fullFramesLost" db:"full_frames_lost"`               // unsigned long        fullFramesLost;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcinboundrtpstreamstats
type rtcInboundRTPStreamStats struct {
	rtcReceivedRTPStreamStats

	ReceiverID                  *string  `json:"receiverId" validate:"required" db:"receiver_id"`                  // required DOMString   receiverId;
	RemoteID                    *string  `json:"remoteId" db:"remote_id"`                                          // DOMString            remoteId;
	FramesDecoded               *uint32  `json:"framesDecoded" db:"frames_decoded"`                                // unsigned long        framesDecoded;
	KeyFramesDecoded            *uint32  `json:"keyFramesDecoded" db:"key_frames_decoded"`                         // unsigned long        keyFramesDecoded;
	FrameWidth                  *uint32  `json:"frameWidth" db:"frame_width"`                                      // unsigned long        frameWidth;
	FrameHeight                 *uint32  `json:"frameHeight" db:"frame_height"`                                    // unsigned long        frameHeight;
	FrameBitDepth               *uint32  `json:"frameBitDepth" db:"frame_bit_depth"`                               // unsigned long        frameBitDepth;
	FramesPerSecond             *float64 `json:"framesPerSecond" db:"frames_per_second"`                           // double               framesPerSecond;
	QpSum                       *uint64  `json:"qpSum" db:"qp_sum"`                                                // unsigned long long   qpSum;
	TotalDecodeTime             *float64 `json:"totalDecodeTime" db:"total_decode_time"`                           // double               totalDecodeTime;
	TotalInterFrameDelay        *float64 `json:"totalInterFrameDelay" db:"total_inter_frame_delay"`                // double               totalInterFrameDelay;
	TotalSquaredInterFrameDelay *float64 `json:"totalSquaredInterFrameDelay" db:"total_squared_inter_frame_delay"` // double               totalSquaredInterFrameDelay;
	VoiceActivityFlag           *bool    `json:"voiceActivityFlag" db:"voice_activity_flag"`                       // boolean              voiceActivityFlag;
	LastPacketReceivedTimestamp *float64 `json:"lastPacketReceivedTimestamp" db:"last_packet_received_timestamp"`  // DOMHighResTimeStamp  lastPacketReceivedTimestamp;
	AverageRtcpInterval         *float64 `json:"averageRtcpInterval" db:"average_rtcp_interval"`                   // double               averageRtcpInterval;
	HeaderBytesReceived         *uint64  `json:"headerBytesReceived" db:"header_bytes_received"`                   // unsigned long long   headerBytesReceived;
	FecPacketsReceived          *uint64  `json:"fecPacketsReceived" db:"fec_packets_received"`                     // unsigned long long   fecPacketsReceived;
	FecPacketsDiscarded         *uint64  `json:"fecPacketsDiscarded" db:"fec_packets_discarded"`                   // unsigned long long   fecPacketsDiscarded;
	BytesReceived               *uint64  `json:"bytesReceived" db:"bytes_received"`                                // unsigned long long   bytesReceived;
	PacketsFailedDecryption     *uint64  `json:"packetsFailedDecryption" db:"packets_failed_decryption"`           // unsigned long long   packetsFailedDecryption;
	PacketsDuplicated           *uint64  `json:"packetsDuplicated" db:"packets_duplicated"`                        // unsigned long long   packetsDuplicated;
	// TODO(v): 特殊型
	PerDscpPacketsReceived         interface{} `json:"perDscpPacketsReceived" db:"per_dscp_packets_received"`                 // record<USVString, unsigned long long> perDscpPacketsReceived;
	NackCount                      *uint32     `json:"nackCount" db:"nack_count"`                                             // unsigned long        nackCount;
	FirCount                       *uint32     `json:"firCount" db:"fir_count"`                                               // unsigned long        firCount;
	PliCount                       *uint32     `json:"pliCount" db:"pli_count"`                                               // unsigned long        pliCount;
	SliCount                       *uint32     `json:"sliCount" db:"sli_count"`                                               // unsigned long        sliCount;
	TotalProcessingDelay           *float64    `json:"totalProcessingDelay" db:"total_processing_delay"`                      // double               totalProcessingDelay;
	EstimatedPlayoutTimestamp      *float64    `json:"estimatedPlayoutTimestamp" db:"estimated_playout_timestamp"`            // DOMHighResTimeStamp  estimatedPlayoutTimestamp;
	JitterBufferDelay              *float64    `json:"jitterBufferDelay" db:"jitter_buffer_delay"`                            // double               jitterBufferDelay;
	JitterBufferEmittedCount       *uint64     `json:"jitterBufferEmittedCount" db:"jitter_buffer_emitted_count"`             // unsigned long long   jitterBufferEmittedCount;
	TotalSamplesReceived           *uint64     `json:"totalSamplesReceived" db:"total_samples_received"`                      // unsigned long long   totalSamplesReceived;
	TotalSamplesDecoded            *uint64     `json:"totalSamplesDecoded" db:"total_samples_decoded"`                        // unsigned long long   totalSamplesDecoded;
	SamplesDecodedWithSilk         *uint64     `json:"samplesDecodedWithSilk" db:"samples_decoded_with_silk"`                 // unsigned long long   samplesDecodedWithSilk;
	SamplesDecodedWithCelt         *uint64     `json:"samplesDecodedWithCelt" db:"samples_decoded_with_celt"`                 // unsigned long long   samplesDecodedWithCelt;
	ConcealedSamples               *uint64     `json:"concealedSamples" db:"concealed_samples"`                               // unsigned long long   concealedSamples;
	SilentConcealedSamples         *uint64     `json:"silentConcealedSamples" db:"silent_concealed_samples"`                  // unsigned long long   silentConcealedSamples;
	ConcealmentEvents              *uint64     `json:"concealmentEvents" db:"concealment_events"`                             // unsigned long long   concealmentEvents;
	InsertedSamplesForDeceleration *uint64     `json:"insertedSamplesForDeceleration" db:"inserted_samples_for_deceleration"` // unsigned long long   insertedSamplesForDeceleration;
	RemovedSamplesForAcceleration  *uint64     `json:"removedSamplesForAcceleration" db:"removed_samples_for_acceleration"`   // unsigned long long   removedSamplesForAcceleration;
	AudioLevel                     *float64    `json:"audioLevel" db:"audio_level"`                                           // double               audioLevel;
	TotalAudioEnergy               *float64    `json:"totalAudioEnergy" db:"total_audio_energy"`                              // double               totalAudioEnergy;
	TotalSamplesDuration           *float64    `json:"totalSamplesDuration" db:"total_samples_duration"`                      // double               totalSamplesDuration;
	FramesReceived                 *uint32     `json:"framesReceived" db:"frames_received"`                                   // unsigned long        framesReceived;
	DecoderImplementation          *string     `json:"decoderImplementation" db:"decoder_implementation"`                     // DOMString            decoderImplementation;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcremoteinboundrtpstreamstats
type rtcRemoteInboundRTPStreamStats struct {
	rtcReceivedRTPStreamStats

	LocalID                   *string  `json:"localId" db:"local_id"`                                       // DOMString            localId;
	RoundTripTime             *float64 `json:"roundTripTime" db:"round_trip_time"`                          // double               roundTripTime;
	TotalRoundTripTime        *float64 `json:"totalRoundTripTime" db:"total_round_trip_time"`               // double               totalRoundTripTime;
	FractionLost              *float64 `json:"fractionLost" db:"fraction_lost"`                             // double               fractionLost;
	ReportsReceived           *uint64  `json:"reportsReceived" db:"reports_received"`                       // unsigned long long   reportsReceived;
	RoundTripTimeMeasurements *uint64  `json:"roundTripTimeMeasurements" db:"round_trip_time_measurements"` // unsigned long long   roundTripTimeMeasurements;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcsentrtpstreamstats
type rtcSentRTPStreamStats struct {
	rtcRTPStreamStats

	PacketsSent *uint32 `json:"packetsSent" db:"packets_sent"` // unsigned long      packetsSent;
	BytesSent   *uint64 `json:"bytesSent" db:"bytes_sent"`     // unsigned long long bytesSent;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcoutboundrtpstreamstats
type rtcOutboundRTPStreamStats struct {
	rtcSentRTPStreamStats

	RtxSSRC                            *uint32     `json:"rtxSsrc" db:"rtx_ssrc"`                                                         // unsigned long        rtxSsrc;
	MediaSourceID                      *string     `json:"mediaSourceId" db:"media_source_id"`                                            // DOMString            mediaSourceId;
	SenderID                           *string     `json:"senderId" db:"sender_id"`                                                       // DOMString            senderId;
	RemoteID                           *string     `json:"remoteId" db:"remote_id"`                                                       // DOMString            remoteId;
	Rid                                *string     `json:"rid" db:"rid"`                                                                  // DOMString            rid;
	LastPacketSentTimestamp            *float64    `json:"lastPacketSentTimestamp" db:"last_packet_sent_timestamp"`                       // DOMHighResTimeStamp  lastPacketSentTimestamp;
	HeaderBytesSent                    *uint64     `json:"headerBytesSent" db:"header_bytes_sent"`                                        // unsigned long long   headerBytesSent;
	PacketsDiscardedOnSend             *uint32     `json:"packetsDiscardedOnSend" db:"packets_discarded_on_send"`                         // unsigned long        packetsDiscardedOnSend;
	BytesDiscardedOnSend               *uint64     `json:"bytesDiscardedOnSend" db:"bytes_discarded_on_send"`                             // unsigned long long   bytesDiscardedOnSend;
	FecPacketsSent                     *uint32     `json:"fecPacketsSent" db:"fec_packets_sent"`                                          // unsigned long        fecPacketsSent;
	RetransmittedPacketsSent           *uint64     `json:"retransmittedPacketsSent" db:"retransmitted_packets_sent"`                      // unsigned long long   retransmittedPacketsSent;
	RetransmittedBytesSent             *uint64     `json:"retransmittedBytesSent" db:"retransmitted_bytes_sent"`                          // unsigned long long   retransmittedBytesSent;
	TargetBitrate                      *float64    `json:"targetBitrate" db:"target_bitrate"`                                             // double               targetBitrate;
	TotalEncodedBytesTarget            *uint64     `json:"totalEncodedBytesTarget" db:"total_encoded_bytes_target"`                       // unsigned long long   totalEncodedBytesTarget;
	FrameWidth                         *uint32     `json:"frameWidth" db:"frame_width"`                                                   // unsigned long        frameWidth;
	FrameHeight                        *uint32     `json:"frameHeight" db:"frame_height"`                                                 // unsigned long        frameHeight;
	FrameBitDepth                      *uint32     `json:"frameBitDepth" db:"frame_bit_depth"`                                            // unsigned long        frameBitDepth;
	FramesPerSecond                    *float64    `json:"framesPerSecond" db:"frames_per_second"`                                        // double               framesPerSecond;
	FramesSent                         *uint32     `json:"framesSent" db:"frames_sent"`                                                   // unsigned long        framesSent;
	HugeFramesSent                     *uint32     `json:"hugeFramesSent" db:"huge_frames_sent"`                                          // unsigned long        hugeFramesSent;
	FramesEncoded                      *uint32     `json:"framesEncoded" db:"frames_encoded"`                                             // unsigned long        framesEncoded;
	KeyFramesEncoded                   *uint32     `json:"keyFramesEncoded" db:"key_frames_encoded"`                                      // unsigned long        keyFramesEncoded;
	FramesDiscardedOnSend              *uint32     `json:"framesDiscardedOnSend" db:"frames_discarded_on_send"`                           // unsigned long        framesDiscardedOnSend;
	QpSum                              *uint64     `json:"qpSum" db:"qp_sum"`                                                             // unsigned long long   qpSum;
	TotalSamplesSent                   *uint64     `json:"totalSamplesSent" db:"total_samples_sent"`                                      // unsigned long long   totalSamplesSent;
	SamplesEncodedWithSilk             *uint64     `json:"samplesEncodedWithSilk" db:"samples_encoded_with_silk"`                         // unsigned long long   samplesEncodedWithSilk;
	SamplesEncodedWithCelt             *uint64     `json:"samplesEncodedWithCelt" db:"samples_encoded_with_celt"`                         // unsigned long long   samplesEncodedWithCelt;
	VoiceActivityFlag                  *bool       `json:"voiceActivityFlag" db:"voice_activity_flag"`                                    // boolean              voiceActivityFlag;
	TotalEncodeTime                    *float64    `json:"totalEncodeTime" db:"total_encode_time"`                                        // double               totalEncodeTime;
	TotalPacketSendDelay               *float64    `json:"totalPacketSendDelay" db:"total_packet_send_delay"`                             // double               totalPacketSendDelay;
	AverageRtcpInterval                *float64    `json:"averageRtcpInterval" db:"average_rtcp_interval"`                                // double               averageRtcpInterval;
	QualityLimitationReason            *string     `json:"qualityLimitationReason" db:"quality_limitation_reason"`                        // RTCQualityLimitationReason                 qualityLimitationReason;
	QualityLimitationDurations         interface{} `json:"qualityLimitationDurations" db:"quality_limitation_durations"`                  // record<DOMString, double> qualityLimitationDurations;
	QualityLimitationResolutionChanges *uint32     `json:"qualityLimitationResolutionChanges" db:"quality_limitation_resolution_changes"` // unsigned long        qualityLimitationResolutionChanges;
	PerDscpPacketsSent                 interface{} `json:"perDscpPacketsSent" db:"per_dscp_packets_sent"`                                 // record<USVString, unsigned long long> perDscpPacketsSent;
	NackCount                          *uint32     `json:"nackCount" db:"nack_count"`                                                     // unsigned long        nackCount;
	FirCount                           *uint32     `json:"firCount" db:"fir_count"`                                                       // unsigned long        firCount;
	PliCount                           *uint32     `json:"pliCount" db:"pli_count"`                                                       // unsigned long        pliCount;
	SliCount                           *uint32     `json:"sliCount" db:"sli_count"`                                                       // unsigned long        sliCount;
	EncoderImplementation              *string     `json:"encoderImplementation" db:"encoder_implementation"`                             // DOMString            encoderImplementation;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcremoteoutboundrtpstreamstats
type rtcRemoteOutboundRTPStreamStats struct {
	rtcSentRTPStreamStats

	LocalID                   *string  `json:"localId" db:"local_id"`                                       // DOMString           localId;
	RemoteTimestamp           *float64 `json:"remoteTimestamp" db:"remote_timestamp"`                       // DOMHighResTimeStamp remoteTimestamp;
	ReportsSent               *uint64  `json:"reportsSent" db:"reports_sent"`                               // unsigned long long  reportsSent;
	RoundTripTime             *float64 `json:"roundTripTime" db:"round_trip_time"`                          // double              roundTripTime;
	TotalRoundTripTime        *float64 `json:"totalRoundTripTime" db:"total_round_trip_time"`               // double              totalRoundTripTime;
	RoundTripTimeMeasurements *uint64  `json:"roundTripTimeMeasurements" db:"round_trip_time_measurements"` // unsigned long long  roundTripTimeMeasurements
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcmediasourcestats
type rtcMediaSourceStats struct {
	rtcStats

	TrackIdentifier *string `json:"trackIdentifier" validate:"required" db:"track_identifier"` // required DOMString       trackIdentifier;
	Kind            *string `json:"kind" validate:"required" db:"kind"`                        // required DOMString       kind;
	RelayedSource   *bool   `json:"relayedSource" db:"relayed_source"`                         // boolean         relayedSource;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcmediasourcestats
type rtcAudioSourceStats struct {
	rtcMediaSourceStats

	AudioLevel                *float64 `json:"audioLevel" db:"audio_level"`                                 // double              audioLevel;
	TotalAudioEnergy          *float64 `json:"totalAudioEnergy" db:"total_audio_energy"`                    // double              totalAudioEnergy;
	TotalSamplesDuration      *float64 `json:"totalSamplesDuration" db:"total_samples_duration"`            // double              totalSamplesDuration;
	EchoReturnLoss            *float64 `json:"echoReturnLoss" db:"echo_return_loss"`                        // double              echoReturnLoss;
	EchoReturnLossEnhancement *float64 `json:"echoReturnLossEnhancement" db:"echo_return_loss_enhancement"` // double              echoReturnLossEnhancement;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcvideosourcestats
type rtcVideoSourceStats struct {
	rtcMediaSourceStats

	Width           *uint32  `json:"width" db:"width"`                       // unsigned long   width;
	Height          *uint32  `json:"height" db:"height"`                     // unsigned long   height;
	BitDepth        *uint32  `json:"bitDepth" db:"bit_depth"`                // unsigned long   bitDepth;
	Frames          *uint32  `json:"frames" db:"frames"`                     // unsigned long   frames;
	FramesPerSecond *float64 `json:"framesPerSecond" db:"frames_per_second"` // double          framesPerSecond;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcrtpcontributingsourcestats
type rtcRTPContributingSourceStats struct {
	rtcStats

	ContributorSsrc      *uint32  `json:"contributorSsrc" validate:"required" db:"contributor_ssrc"`         // required unsigned long contributorSsrc;
	InboundRTPStreamID   *string  `json:"inboundRTPStreamId" validate:"required" db:"inbound_rtp_stream_id"` // required DOMString     inboundRTPStreamId;
	PacketsContributedTo *uint32  `json:"packetsContributedTo" db:"packets_contributed_to"`                  // unsigned long packetsContributedTo;
	AudioLevel           *float64 `json:"audioLevel" db:"audio_level"`                                       // double        audioLevel;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcpeerconnectionstats
type rtcPeerConnectionStats struct {
	rtcStats

	DataChannelsOpened    *uint32 `json:"dataChannelsOpened" db:"data_channels_opened"`       // unsigned long dataChannelsOpened;
	DataChannelsClosed    *uint32 `json:"dataChannelsClosed" db:"data_channels_closed"`       // unsigned long dataChannelsClosed;
	DataChannelsRequested *uint32 `json:"dataChannelsRequested" db:"data_channels_requested"` // unsigned long dataChannelsRequested;
	DataChannelsAccepted  *uint32 `json:"dataChannelsAccepted" db:"data_channels_accepted"`   // unsigned long dataChannelsAccepted;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcrtptransceiverstats
type rtcRTPTransceiverStats struct {
	rtcStats

	SenderID   *string `json:"senderId" validate:"required" db:"sender_id"`     // required DOMString senderId;
	ReceiverID *string `json:"receiverId" validate:"required" db:"receiver_id"` // required DOMString receiverId;
	Mid        *string `json:"mid" db:"mid"`                                    // DOMString mid;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcmediahandlerstats
type rtcMediaHandlerStats struct {
	rtcStats

	TrackIdentifier *string `json:"trackIdentifier" db:"track_identifier"` // DOMString           trackIdentifier;
	Ended           *bool   `json:"ended" db:"ended"`                      // boolean             ended;
	Kind            *string `json:"kind" validate:"required" db:"kind"`    // required DOMString  kind;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcvideohandlerstats
type rtcVideoHandlerStats struct {
	rtcMediaHandlerStats
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcvideosenderstats
type rtcVideoSenderStats struct {
	rtcVideoHandlerStats
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcvideoreceiverstats
type rtcVideoReceiverStats struct {
	rtcVideoHandlerStats
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcaudiohandlerstats
type rtcAudioHandlerStats struct {
	rtcMediaHandlerStats
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcaudiosenderstats
type rtcAudioSenderStats struct {
	rtcAudioHandlerStats

	MediaSourceID *string `json:"mediaSourceId" db:"media_source_id"` // DOMString           mediaSourceId;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcaudioreceiverstats
type rtcAudioReceiverStats struct {
	rtcAudioHandlerStats
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcdatachannelstats
type rtcDataChannelStats struct {
	rtcStats

	Label                 *string `json:"label" db:"label"`                                   // DOMString           label;
	Protocol              *string `json:"protocol" db:"protocol"`                             // DOMString           protocol;
	DataChannelIdentifier *uint16 `json:"dataChannelIdentifier" db:"data_channel_identifier"` // unsigned short      dataChannelIdentifier;
	State                 *string `json:"state" validate:"required" db:"state"`               // required RTCDataChannelState state;
	MessagesSent          *uint32 `json:"messagesSent" db:"messages_sent"`                    // unsigned long       messagesSent;
	BytesSent             *uint64 `json:"bytesSent" db:"bytes_sent"`                          // unsigned long long  bytesSent;
	MessagesReceived      *uint32 `json:"messagesReceived" db:"messages_received"`            // unsigned long       messagesReceived;
	BytesReceived         *uint64 `json:"bytesReceived" db:"bytes_received"`                  // unsigned long long  bytesReceived;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtctransportstats
type rtcTransportStats struct {
	rtcStats

	PacketsSent                  *uint64 `json:"packetsSent" db:"packets_sent"`                                     // unsigned long long    packetsSent;
	PacketsReceived              *uint64 `json:"packetsReceived" db:"packets_received"`                             // unsigned long long    packetsReceived;
	BytesSent                    *uint64 `json:"bytesSent" db:"bytes_sent"`                                         // unsigned long long    bytesSent;
	BytesReceived                *uint64 `json:"bytesReceived" db:"bytes_received"`                                 // unsigned long long    bytesReceived;
	RtcpTransportStatsID         *string `json:"rtcpTransportStatsId" db:"rtcp_transport_stats_id"`                 // DOMString             rtcpTransportStatsId;
	IceRole                      *string `json:"iceRole" db:"ice_role"`                                             // RTCIceRole            iceRole;
	IceLocalUsernameFragment     *string `json:"iceLocalUsernameFragment" db:"ice_local_username_fragment"`         // DOMString             iceLocalUsernameFragment;
	DtlsState                    *string `json:"dtlsState" validate:"required" db:"dtls_state"`                     // required RTCDtlsTransportState dtlsState;
	IceState                     *string `json:"iceState" db:"ice_state"`                                           // RTCIceTransportState  iceState;
	SelectedCandidatePairID      *string `json:"selectedCandidatePairId" db:"selected_candidate_pair_id"`           // DOMString             selectedCandidatePairId;
	LocalCertificateID           *string `json:"localCertificateId" db:"local_certificate_id"`                      // DOMString             localCertificateId;
	RemoteCertificateID          *string `json:"remoteCertificateId" db:"remote_certificate_id"`                    // DOMString             remoteCertificateId;
	TLSVersion                   *string `json:"tlsVersion" db:"tls_version"`                                       // DOMString             tlsVersion;
	DtlsCipher                   *string `json:"dtlsCipher" db:"dtls_cipher"`                                       // DOMString             dtlsCipher;
	SrtpCipher                   *string `json:"srtpCipher" db:"srtp_cipher"`                                       // DOMString             srtpCipher;
	TLSGroup                     *string `json:"tlsGroup" db:"tls_group"`                                           // DOMString             tlsGroup;
	SelectedCandidatePairChanges *uint32 `json:"selectedCandidatePairChanges" db:"selected_candidate_pair_changes"` // unsigned long         selectedCandidatePairChanges;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcsctptransportstats
type rtcSctpTransportStats struct {
	rtcStats

	TransportID           *string  `json:"transportId" db:"transport_id"`                       // DOMString transportId;
	SmoothedRoundTripTime *float64 `json:"smoothedRoundTripTime" db:"smoothed_round_trip_time"` // double smoothedRoundTripTime;
	CongestionWindow      *uint32  `json:"congestionWindow" db:"congestion_window"`             // unsigned long congestionWindow;
	ReceiverWindow        *uint32  `json:"receiverWindow" db:"receiver_window"`                 // unsigned long receiverWindow;
	Mtu                   *uint32  `json:"mtu" db:"mtu"`                                        // unsigned long mtu;
	UnackData             *uint32  `json:"unackData" db:"unack_data"`                           // unsigned long unackData;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcicecandidatestats
type rtcIceCandidateStats struct {
	rtcStats

	TransportID   *string `json:"transportId" validate:"required" db:"transport_id"`     // required DOMString       transportId;
	Address       *string `json:"address" db:"address"`                                  // DOMString?               address;
	Port          *int32  `json:"port" db:"port"`                                        // long                     port;
	Protocol      *string `json:"protocol" db:"protocol"`                                // DOMString                protocol;
	CandidateType *string `json:"candidateType" validate:"required" db:"candidate_type"` // required RTCIceCandidateType candidateType;
	Priority      *int32  `json:"priority" db:"priority"`                                // long                     priority;
	URL           *string `json:"url" db:"url"`                                          // DOMString                url;
	RelayProtocol *string `json:"relayProtocol" db:"relay_protocol"`                     // DOMString                relayProtocol;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtcicecandidatepairstats
type rtcIceCandidatePairStats struct {
	rtcStats

	TransportID                 *string  `json:"transportId" validate:"required" db:"transport_id"`               // required DOMString            transportId;
	LocalCandidateID            *string  `json:"localCandidateId" validate:"required" db:"local_candidate_id"`    // required DOMString            localCandidateId;
	RemoteCandidateID           *string  `json:"remoteCandidateId" validate:"required" db:"remote_candidate_id"`  // required DOMString            remoteCandidateId;
	State                       *string  `json:"state" validate:"required" db:"state"`                            // required RTCStatsIceCandidatePairState state;
	Nominated                   *bool    `json:"nominated" db:"nominated"`                                        // boolean                       nominated;
	PacketsSent                 *uint64  `json:"packetsSent" db:"packets_sent"`                                   // unsigned long long            packetsSent;
	PacketsReceived             *uint64  `json:"packetsReceived" db:"packets_received"`                           // unsigned long long            packetsReceived;
	BytesSent                   *uint64  `json:"bytesSent" db:"bytes_sent"`                                       // unsigned long long            bytesSent;
	BytesReceived               *uint64  `json:"bytesReceived" db:"bytes_received"`                               // unsigned long long            bytesReceived;
	LastPacketSentTimestamp     *float64 `json:"lastPacketSentTimestamp" db:"last_packet_sent_timestamp"`         // DOMHighResTimeStamp           lastPacketSentTimestamp;
	LastPacketReceivedTimestamp *float64 `json:"lastPacketReceivedTimestamp" db:"last_packet_received_timestamp"` // DOMHighResTimeStamp           lastPacketReceivedTimestamp;
	FirstRequestTimestamp       *float64 `json:"firstRequestTimestamp" db:"first_request_timestamp"`              // DOMHighResTimeStamp           firstRequestTimestamp;
	LastRequestTimestamp        *float64 `json:"lastRequestTimestamp" db:"last_request_timestamp"`                // DOMHighResTimeStamp           lastRequestTimestamp;
	LastResponseTimestamp       *float64 `json:"lastResponseTimestamp" db:"last_response_timestamp"`              // DOMHighResTimeStamp           lastResponseTimestamp;
	TotalRoundTripTime          *float64 `json:"totalRoundTripTime" db:"total_round_trip_time"`                   // double                        totalRoundTripTime;
	CurrentRoundTripTime        *float64 `json:"currentRoundTripTime" db:"current_round_trip_time"`               // double                        currentRoundTripTime;
	AvailableOutgoingBitrate    *float64 `json:"availableOutgoingBitrate" db:"available_outgoing_bitrate"`        // double                        availableOutgoingBitrate;
	AvailableIncomingBitrate    *float64 `json:"availableIncomingBitrate" db:"available_incoming_bitrate"`        // double                        availableIncomingBitrate;
	CircuitBreakerTriggerCount  *uint32  `json:"circuitBreakerTriggerCount" db:"circuit_breaker_trigger_count"`   // unsigned long                 circuitBreakerTriggerCount;
	RequestsReceived            *uint64  `json:"requestsReceived" db:"requests_received"`                         // unsigned long long            requestsReceived;
	RequestsSent                *uint64  `json:"requestsSent" db:"requests_sent"`                                 // unsigned long long            requestsSent;
	ResponsesReceived           *uint64  `json:"responsesReceived" db:"responses_received"`                       // unsigned long long            responsesReceived;
	ResponsesSent               *uint64  `json:"responsesSent" db:"responses_sent"`                               // unsigned long long            responsesSent;
	RetransmissionsReceived     *uint64  `json:"retransmissionsReceived" db:"retransmissions_received"`           // unsigned long long            retransmissionsReceived;
	RetransmissionsSent         *uint64  `json:"retransmissionsSent" db:"retransmissions_sent"`                   // unsigned long long            retransmissionsSent;
	ConsentRequestsSent         *uint64  `json:"consentRequestsSent" db:"consent_requests_sent"`                  // unsigned long long            consentRequestsSent;
	ConsentExpiredTimestamp     *float64 `json:"consentExpiredTimestamp" db:"consent_expired_timestamp"`          // DOMHighResTimeStamp           consentExpiredTimestamp;
	PacketsDiscardedOnSend      *uint32  `json:"packetsDiscardedOnSend" db:"packets_discarded_on_send"`           // unsigned long                 packetsDiscardedOnSend;
	BytesDiscardedOnSend        *uint64  `json:"bytesDiscardedOnSend" db:"bytes_discarded_on_send"`               // unsigned long long            bytesDiscardedOnSend;
	RequestBytesSent            *uint64  `json:"requestBytesSent" db:"request_bytes_sent"`                        // unsigned long long            requestBytesSent;
	ConsentRequestBytesSent     *uint64  `json:"consentRequestBytesSent" db:"consent_request_bytes_sent"`         // unsigned long long            consentRequestBytesSent;
	ResponseBytesSent           *uint64  `json:"responseBytesSent" db:"response_bytes_sent"`                      // unsigned long long            responseBytesSent;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtccertificatestats
type rtcCertificateStats struct {
	rtcStats

	Fingerprint          *string `json:"fingerprint" validate:"required" db:"fingerprint"`                    // required DOMString fingerprint;
	FingerprintAlgorithm *string `json:"fingerprintAlgorithm" validate:"required" db:"fingerprint_algorithm"` // required DOMString fingerprintAlgorithm;
	Base64Certificate    *string `json:"base64Certificate" validate:"required" db:"base64_certificate"`       // required DOMString base64Certificate;
	IssuerCertificateID  *string `json:"issuerCertificateId" db:"issuer_certificate_id"`                      // DOMString issuerCertificateId;
}

// https://www.w3.org/TR/webrtc-stats/#dom-rtciceserverstats
type rtcIceServerStats struct {
	rtcStats

	URL                    *string  `json:"url" validate:"required" db:"url"`                     // required DOMString url;
	Port                   *int32   `json:"port" db:"port"`                                       // long port;
	RelayProtocol          *string  `json:"relayProtocol" db:"relay_protocol"`                    // DOMString relayProtocol;
	TotalRequestsSent      *uint32  `json:"totalRequestsSent" db:"total_requests_sent"`           // unsigned long totalRequestsSent;
	TotalResponsesReceived *uint32  `json:"totalResponsesReceived" db:"total_responses_received"` // unsigned long totalResponsesReceived;
	TotalRoundTripTime     *float64 `json:"totalRoundTripTime" db:"total_round_trip_time"`        // double totalRoundTripTime;
}

// Obsolate Stats

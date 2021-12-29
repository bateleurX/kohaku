# 注意事項


## 統計情報に含まれない項目

- 動作を確認した現時点では、次の項目はブラウザから送信されてきませんでしたので注意してください

### Google Chrome 95


#### rtc_codec_stats

- codec_type


#### rtc_video_source_stats

- relayed_source
- bit_depth


#### rtc_audio_source_stats

- relayed_source


#### rtc_ice_candidate_pair_stats

- pakcets_sent
- pakcets_received
- last_packet_sent_timestamp
- last_packet_received_timestamp
- first_request_timestamp
- last_request_timestamp
- last_response_timestamp
- circuit_breaker_trigger_count
- retransmissions_received
- retransmissions_sent
- consent_expired_timestamp
- packets_discarded_on_send
- bytes_discarded_on_send
- request_bytes_sent
- consent_request_bytes_sent
- response_bytes_sent


#### rtc_ice_candidate_stats

- url


#### rtc_inbound_rtp_stream_stats

- packets_repaired
- burst_packets_lost
- burst_packets_discarded
- burst_loss_count
- burst_discard_count
- burst_lostt_rate
- burst_discard_rate
- gap_loss_rate
- gap_discard_rate
- partial_frames_lost
- full_frames_lost
- receiver_id
- frame_bit_depth
- qp_sum
- voice_activity_flag
- average_rtcp_interval
- packets_failed_decryption
- packets_duplicated
- per_dscp_packets_received
- sli_count
- total_processing_delay
- total_samples_decoded
- samples_decoded_with_silk
- samples_decoded_with_celt


#### rtc_outbound_rtp_stream_stats

- rtx_ssrc
- sender_id
- rid
- last_packet_sent_timestamp
- packets_discarded_on_send
- bytes_discarded_on_send
- fec_packets_sent
- target_bitrate
- frame_bit_depth
- frames_discarded_on_send
- total_samples_sent
- samples_encoded_with_silk
- samples_encoded_with_celt
- voice_activity_flag
- average_rtcp_interval
- per_dscp_packets_sent
- sli_count


#### rtc_remote_inbound_rtp_stream_stats

- packets_received
- packets_discarded
- packets_repaired
- burst_packets_lost
- burst_packets_discarded
- burst_loss_count
- burst_discarde_count
- burst_loss_rate
- burst_discard_rate
- gap_loss_rage
- gap_discard_rate
- frames_dropped
- partial_frames_lost
- full_frames_lost
- reports_received


#### rtc_remote_outbound_rtp_stream_stats

- round_trip_time


#### rtc_transport_stats

- rtcp_transport_stats_id
- ice_role
- ice_local_username_fragment
- ice_state
- tls_group


### Safari 15.1


#### rtc_codec_stats

- codec_type
- transport_id


#### rtc_video_source_stats

- relayed_source
- width
- height
- bit_depth
- frames
- frames_per_second


#### rtc_audio_source_stats

- relayed_source


#### rtc_ice_candidate_pair_stats

- pakcets_sent
- pakcets_received
- last_packet_sent_timestamp
- last_packet_received_timestamp
- first_request_timestamp
- last_request_timestamp
- last_response_timestamp
- available_incoming_bitrate
- circuit_breaker_trigger_count
- retransmissions_received
- retransmissions_sent
- consent_requests_sent
- consent_expired_timestamp
- packets_discarded_on_send
- bytes_discarded_on_send
- request_bytes_sent
- consent_request_bytes_sent
- response_bytes_sent


#### rtc_ice_candidate_stats

- url
- relay_protocol


#### rtc_inbound_rtp_stream_stats

- packets_discarded
- packets_repaired
- burst_packets_lost
- burst_packets_discarded
- burst_loss_count
- burst_discard_count
- burst_lostt_rate
- burst_discard_rate
- gap_loss_rate
- gap_discard_rate
- frames_dropped
- partial_frames_lost
- full_frames_lost
- receiver_id
- remote_id
- frame_bit_depth
- qp_sum
- voice_activity_flag
- average_rtcp_interval
- packets_failed_decryption
- packets_duplicated
- per_dscp_packets_received
- sli_count
- total_processing_delay
- total_samples_decoded
- samples_decoded_with_silk
- samples_decoded_with_celt
- decoder_implementation


#### rtc_outbound_rtp_stream_stats

- rtx_ssrc
- sender_id
- rid
- last_packet_sent_timestamp
- packets_discarded_on_send
- bytes_discarded_on_send
- fec_packets_sent
- target_bitrate
- frame_bit_depth
- frames_discarded_on_send
- total_samples_sent
- samples_encoded_with_silk
- samples_encoded_with_celt
- voice_activity_flag
- average_rtcp_interval
- quality_limitation_reason
- per_dscp_packets_sent
- sli_count
- encoder_implementation


#### rtc_remote_inbound_rtp_stream_stats

- packets_received
- packets_lost
- jitter
- packets_discarded
- packets_repaired
- burst_packets_lost
- burst_packets_discarded
- burst_loss_count
- burst_discarde_count
- burst_loss_rate
- burst_discard_rate
- gap_loss_rage
- gap_discard_rate
- frames_dropped
- partial_frames_lost
- full_frames_lost
- total_round_trip_time
- fraction_lost
- reports_received
- round_trip_time_measurements


#### rtc_remote_outbound_rtp_stream_stats

- remote-outbound-rtp は送信されてきていない


#### rtc_transport_stats

- packets_sent
- packets_received
- rtcp_transport_stats_id
- ice_role
- ice_local_username_fragment
- ice_state
- tls_group
- selected_candidate_pair_changes


#### rtc_data_channel_stats

- data_channel_identifier

    - 代わりに data_channel_id が送信されてきている

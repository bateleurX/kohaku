-- name: InsertSoraNode :exec
INSERT INTO sora_node (
  timestamp,
  label, version, node_name
)
SELECT
  @timestamp,
  @label, @version, @node_name
WHERE
  NOT EXISTS (
    SELECT id
    FROM sora_node
    WHERE (
      (label = @label::varchar(255)) AND
      (version = @version::varchar(255)) AND
      (node_name = @node_name::varchar(255))
    )
);

-- name: InsertSoraConnection :exec
INSERT INTO sora_connection (
  timestamp,
  label, version, node_name,
  multistream, simulcast, spotlight,
  role, channel_id, session_id, client_id, connection_id
)
SELECT
  @timestamp,
  @label, @version, @node_name,
  @multistream, @simulcast, @spotlight,
  @role, @channel_id, @session_id, @client_id, @connection_id
WHERE
  NOT EXISTS (
    SELECT id
    FROM sora_connection
    WHERE (
      (channel_id = @channel_id::varchar(255)) AND
      (session_id = @session_id::char(26)) AND
      (client_id = @client_id::varchar(255)) AND
      (connection_id = @connection_id::char(26))
    )
);

-- name: InsertErlangVMMemoryStats :exec
INSERT INTO erlang_vm_memory_stats (
  time,
  sora_version, sora_label, sora_node_name,
  stats_type,
  type_total, type_processes, type_processes_used, type_system,
  type_atom, type_atom_used, type_binary, type_code, type_ets
) VALUES (
  @time,
  @sora_version, @sora_label, @sora_node_name,
  @stats_type,
  @type_total, @type_processes, @type_processes_used, @type_system,
  @type_atom, @type_atom_used, @type_binary, @type_code, @type_ets
);

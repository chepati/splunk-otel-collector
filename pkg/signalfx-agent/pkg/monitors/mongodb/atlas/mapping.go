package atlas

// metricsMap maps Atlas measurement names to metrics.
var metricsMap = map[string]string{
	"ASSERT_REGULAR":                               assertsRegular,
	"ASSERT_WARNING":                               assertsWarning,
	"ASSERT_MSG":                                   assertsMsg,
	"ASSERT_USER":                                  assertsUser,
	"CACHE_BYTES_READ_INTO":                        cacheBytesReadInto,
	"CACHE_BYTES_WRITTEN_FROM":                     cacheBytesWrittenFrom,
	"CACHE_USED_BYTES":                             cacheUsedBytes,
	"CACHE_DIRTY_BYTES":                            cacheDirtyBytes,
	"OPCOUNTER_CMD":                                opcounterCommand,
	"OPCOUNTER_DELETE":                             opcounterDelete,
	"OPCOUNTER_GETMORE":                            opcounterGetmore,
	"OPCOUNTER_INSERT":                             opcounterInsert,
	"OPCOUNTER_QUERY":                              opcounterQuery,
	"OPCOUNTER_UPDATE":                             opcounterUpdate,
	"OPCOUNTER_REPL_CMD":                           opcounterReplCommand,
	"OPCOUNTER_REPL_DELETE":                        opcounterReplDelete,
	"OPCOUNTER_REPL_INSERT":                        opcounterReplInsert,
	"OPCOUNTER_REPL_UPDATE":                        opcounterReplUpdate,
	"CONNECTIONS":                                  connectionsCurrent,
	"CURSORS_TOTAL_OPEN":                           cursorsTotalOpen,
	"CURSORS_TOTAL_TIMED_OUT":                      cursorsTimedOut,
	"DB_STORAGE_TOTAL":                             storageSize,
	"DB_DATA_SIZE_TOTAL":                           dataSize,
	"DB_INDEX_SIZE_TOTAL":                          indexSize,
	"DOCUMENT_METRICS_RETURNED":                    documentMetricsReturned,
	"DOCUMENT_METRICS_INSERTED":                    documentMetricsInserted,
	"DOCUMENT_METRICS_UPDATED":                     documentMetricsUpdated,
	"DOCUMENT_METRICS_DELETED":                     documentMetricsDeleted,
	"MEMORY_RESIDENT":                              memResident,
	"MEMORY_VIRTUAL":                               memVirtual,
	"MEMORY_MAPPED":                                memMapped,
	"NETWORK_BYTES_IN":                             networkBytesIn,
	"NETWORK_BYTES_OUT":                            networkBytesOut,
	"NETWORK_NUM_REQUESTS":                         networkNumRequests,
	"OP_EXECUTION_TIME_READS":                      opExecutionTimeReads,
	"OP_EXECUTION_TIME_WRITES":                     opExecutionTimeWrites,
	"OP_EXECUTION_TIME_COMMANDS":                   opExecutionTimeCommands,
	"OPLOG_RATE_GB_PER_HOUR":                       oplogRate,
	"OPLOG_MASTER_LAG_TIME_DIFF":                   oplogMasterLagTimeDiff,
	"OPLOG_SLAVE_LAG_MASTER_TIME":                  oplogSlaveLagMasterTime,
	"OPLOG_MASTER_TIME":                            oplogMasterTime,
	"EXTRA_INFO_PAGE_FAULTS":                       extraInfoPageFaults,
	"GLOBAL_LOCK_CURRENT_QUEUE_TOTAL":              globalLockCurrentQueueTotal,
	"GLOBAL_LOCK_CURRENT_QUEUE_READERS":            globalLockCurrentQueueReaders,
	"GLOBAL_LOCK_CURRENT_QUEUE_WRITERS":            globalLockCurrentQueueWriters,
	"QUERY_EXECUTOR_SCANNED":                       queryExecutorScanned,
	"QUERY_EXECUTOR_SCANNED_OBJECTS":               queryExecutorScannedObjects,
	"QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED": queryTargetingScannedObjectsPerReturned,
	"QUERY_TARGETING_SCANNED_PER_RETURNED":         queryTargetingScannedPerReturned,
	"TICKETS_AVAILABLE_READS":                      ticketsAvailableReads,
	"TICKETS_AVAILABLE_WRITE":                      ticketsAvailableWrite,
	"DISK_PARTITION_IOPS_READ":                     diskPartitionIopsRead,
	"DISK_PARTITION_IOPS_WRITE":                    diskPartitionIopsWrite,
	"DISK_PARTITION_IOPS_TOTAL":                    diskPartitionIopsTotal,
	"DISK_PARTITION_LATENCY_READ":                  diskPartitionLatencyRead,
	"DISK_PARTITION_LATENCY_WRITE":                 diskPartitionLatencyWrite,
	"DISK_PARTITION_SPACE_FREE":                    diskPartitionSpaceFree,
	"DISK_PARTITION_SPACE_PERCENT_FREE":            diskPartitionSpacePercentFree,
	"DISK_PARTITION_SPACE_USED":                    diskPartitionSpaceUsed,
	"DISK_PARTITION_SPACE_PERCENT_USED":            diskPartitionSpacePercentUsed,
	"DISK_PARTITION_UTILIZATION":                   diskPartitionUtilization,
	"PROCESS_CPU_USER":                             processCPUUser,
	"PROCESS_CPU_KERNEL":                           processCPUKernel,
	"PROCESS_NORMALIZED_CPU_USER":                  processNormalizedCPUUser,
	"PROCESS_NORMALIZED_CPU_KERNEL":                processNormalizedCPUKernel,
	"PROCESS_NORMALIZED_CPU_CHILDREN_USER":         processNormalizedCPUChildrenUser,
	"PROCESS_NORMALIZED_CPU_CHILDREN_KERNEL":       processNormalizedCPUChildrenKernel,
	"SYSTEM_CPU_USER":                              systemCPUUser,
	"SYSTEM_CPU_KERNEL":                            systemCPUKernel,
	"SYSTEM_CPU_NICE":                              systemCPUNice,
	"SYSTEM_CPU_IOWAIT":                            systemCPUIowait,
	"SYSTEM_CPU_IRQ":                               systemCPUIrq,
	"SYSTEM_CPU_SOFTIRQ":                           systemCPUSoftirq,
	"SYSTEM_CPU_GUEST":                             systemCPUGuest,
	"SYSTEM_CPU_STEAL":                             systemCPUSteal,
	"SYSTEM_NORMALIZED_CPU_USER":                   systemNormalizedCPUUser,
	"SYSTEM_NORMALIZED_CPU_KERNEL":                 systemNormalizedCPUKernel,
	"SYSTEM_NORMALIZED_CPU_NICE":                   systemNormalizedCPUNice,
	"SYSTEM_NORMALIZED_CPU_IOWAIT":                 systemNormalizedCPUIowait,
	"SYSTEM_NORMALIZED_CPU_IRQ":                    systemNormalizedCPUIrq,
	"SYSTEM_NORMALIZED_CPU_SOFTIRQ":                systemNormalizedCPUSoftirq,
	"SYSTEM_NORMALIZED_CPU_GUEST":                  systemNormalizedCPUGuest,
	"SYSTEM_NORMALIZED_CPU_STEAL":                  systemNormalizedCPUSteal,
	"OPERATIONS_SCAN_AND_ORDER":                    operationsScanAndOrder,
	"BACKGROUND_FLUSH_AVG":                         backgroundFlushAvg,
}
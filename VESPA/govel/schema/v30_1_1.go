package schema

const v3011 = `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "title": "VES Event Listener Common Event Format",
  "type": "object",
  "properties": {
    "event": {
      "$ref": "#/definitions/event"
    },
    "eventList": {
      "$ref": "#/definitions/eventList"
    }
  },
  "definitions": {
    "schemaHeaderBlock": {
      "description": "schema date, version, author and associated API",
      "type": "object",
      "properties": {
        "associatedApi": {
          "description": "VES Event Listener",
          "type": "string"
        },
        "lastUpdatedBy": {
          "description": "tl2972",
          "type": "string"
        },
        "schemaDate": {
          "description": "January 28, 2020",
          "type": "string"
        },
        "schemaVersion": {
          "description": "30.1.1",
          "type": "number"
        }
      }
    },
    "schemaLicenseAndCopyrightNotice": {
      "description": "Copyright (c) 2020, AT&T Intellectual Property.  All rights reserved",
      "type": "object",
      "properties": {
        "apacheLicense2.0": {
          "description": "Licensed under the Apache License, Version 2.0 (the 'License'); you may not use this file except in compliance with the License. You may obtain a copy of the License at:",
          "type": "string"
        },
        "licenseUrl": {
          "description": "http://www.apache.org/licenses/LICENSE-2.0",
          "type": "string"
        },
        "asIsClause": {
          "description": "Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an 'AS IS' BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.",
          "type": "string"
        },
        "permissionsAndLimitations": {
          "description": "See the License for the specific language governing permissions and limitations under the License.",
          "type": "string"
        }
      }
    },
    "arrayOfJsonObject": {
      "description": "array of json objects described by name, schema and other meta-information",
      "type": "array",
      "items": {
        "$ref": "#/definitions/jsonObject"
      }
    },
    "arrayOfNamedHashMap": {
      "description": "array of named hashMaps",
      "type": "array",
      "items": {
        "$ref": "#/definitions/namedHashMap"
      }
    },
    "codecsInUse": {
      "description": "number of times an identified codec was used over the measurementInterval",
      "type": "object",
      "properties": {
        "codecIdentifier": {
          "type": "string"
        },
        "numberInUse": {
          "type": "integer"
        }
      },
      "additionalProperties": false,
      "required": [
        "codecIdentifier",
        "numberInUse"
      ]
    },
    "commonEventHeader": {
      "description": "fields common to all events",
      "type": "object",
      "properties": {
        "domain": {
          "description": "the eventing domain associated with the event",
          "type": "string",
          "enum": [
            "fault",
            "heartbeat",
            "measurement",
            "mobileFlow",
            "notification",
            "other",
            "perf3gpp",
            "pnfRegistration",
            "sipSignaling",
            "stateChange",
            "syslog",
            "thresholdCrossingAlert",
            "voiceQuality"
          ]
        },
        "eventId": {
          "description": "event key that is unique to the event source",
          "type": "string"
        },
        "eventName": {
          "description": "unique event name",
          "type": "string"
        },
        "eventType": {
          "description": "for example - applicationNf, guestOS, hostOS, platform",
          "type": "string"
        },
        "internalHeaderFields": {
          "$ref": "#/definitions/internalHeaderFields"
        },
        "lastEpochMicrosec": {
          "description": "the latest unix time aka epoch time associated with the event from any component--as microseconds elapsed since 1 Jan 1970 not including leap seconds",
          "type": "number"
        },
        "nfcNamingCode": {
          "description": "3 character network function component type, aligned with vfc naming standards",
          "type": "string"
        },
        "nfNamingCode": {
          "description": "4 character network function type, aligned with nf naming standards",
          "type": "string"
        },
        "nfVendorName": {
          "description": "network function vendor name",
          "type": "string"
        },
        "priority": {
          "description": "processing priority",
          "type": "string",
          "enum": [
            "High",
            "Medium",
            "Normal",
            "Low"
          ]
        },
        "reportingEntityId": {
          "description": "UUID identifying the entity reporting the event, for example an OAM VM; must be populated by the ATT enrichment process",
          "type": "string"
        },
        "reportingEntityName": {
          "description": "name of the entity reporting the event, for example, an EMS name; may be the same as sourceName",
          "type": "string"
        },
        "sequence": {
          "description": "ordering of events communicated by an event source instance or 0 if not needed",
          "type": "integer"
        },
        "sourceId": {
          "description": "UUID identifying the entity experiencing the event issue; must be populated by the ATT enrichment process",
          "type": "string"
        },
        "sourceName": {
          "description": "name of the entity experiencing the event issue",
          "type": "string"
        },
        "startEpochMicrosec": {
          "description": "the earliest unix time aka epoch time associated with the event from any component--as microseconds elapsed since 1 Jan 1970 not including leap seconds",
          "type": "number"
        },
        "timeZoneOffset": {
          "description": "UTC offset for the local time zone of the device as UTC+/-hh.mm",
          "type": "string"
        },
        "version": {
          "description": "version of the event header",
          "type": "string",
          "enum": [
            "4.0",
            "4.0.1",
            "4.1"
          ]
        },
        "vesEventListenerVersion": {
          "description": "version of the VES Event Listener API",
          "type": "string",
          "enum": [
            "7.0",
            "7.0.1",
            "7.1",
            "7.1.1"
          ]
        }
      },
      "additionalProperties": false,
      "required": [
        "domain",
        "eventId",
        "eventName",
        "lastEpochMicrosec",
        "priority",
        "reportingEntityName",
        "sequence",
        "sourceName",
        "startEpochMicrosec",
        "version",
        "vesEventListenerVersion"
      ]
    },
    "counter": {
      "description": "performance counter",
      "type": "object",
      "properties": {
        "criticality": {
          "type": "string",
          "enum": [
            "CRIT",
            "MAJ"
          ]
        },
        "hashMap": {
          "$ref": "#/definitions/hashMap"
        },
        "thresholdCrossed": {
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "criticality",
        "hashMap",
        "thresholdCrossed"
      ]
    },
    "cpuUsage": {
      "description": "usage of an identified CPU",
      "type": "object",
      "properties": {
        "cpuCapacityContention": {
          "description": "the amount of time the CPU cannot run due to contention, in milliseconds over the measurementInterval",
          "type": "number"
        },
        "cpuDemandAvg": {
          "description": "the total CPU time that the NF/NFC/VM could use if there was no contention, in milliseconds over the measurementInterval",
          "type": "number"
        },
        "cpuDemandMhz": {
          "description": "CPU demand in megahertz",
          "type": "number"
        },
        "cpuDemandPct": {
          "description": "CPU demand as a percentage of the provisioned capacity",
          "type": "number"
        },
        "cpuIdentifier": {
          "description": "cpu identifer",
          "type": "string"
        },
        "cpuIdle": {
          "description": "percentage of CPU time spent in the idle task",
          "type": "number"
        },
        "cpuLatencyAvg": {
          "description": "percentage of time the VM is unable to run because it is contending for access to the physical CPUs",
          "type": "number"
        },
        "cpuOverheadAvg": {
          "description": "the overhead demand above available allocations and reservations, in milliseconds over the measurementInterval",
          "type": "number"
        },
        "cpuSwapWaitTime": {
          "description": "swap wait time. in milliseconds over the measurementInterval",
          "type": "number"
        },
        "cpuUsageInterrupt": {
          "description": "percentage of time spent servicing interrupts",
          "type": "number"
        },
        "cpuUsageNice": {
          "description": "percentage of time spent running user space processes that have been niced",
          "type": "number"
        },
        "cpuUsageSoftIrq": {
          "description": "percentage of time spent handling soft irq interrupts",
          "type": "number"
        },
        "cpuUsageSteal": {
          "description": "percentage of time spent in involuntary wait which is neither user, system or idle time and is effectively time that went missing",
          "type": "number"
        },
        "cpuUsageSystem": {
          "description": "percentage of time spent on system tasks running the kernel",
          "type": "number"
        },
        "cpuUsageUser": {
          "description": "percentage of time spent running un-niced user space processes",
          "type": "number"
        },
        "cpuWait": {
          "description": "percentage of CPU time spent waiting for I/O operations to complete",
          "type": "number"
        },
        "percentUsage": {
          "description": "aggregate cpu usage of the virtual machine on which the xNFC reporting the event is running",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "cpuIdentifier",
        "percentUsage"
      ]
    },
    "diskUsage": {
      "description": "usage of an identified disk",
      "type": "object",
      "properties": {
        "diskBusResets": {
          "description": "number of bus resets over the measurementInterval",
          "type": "number"
        },
        "diskCommandsAborted": {
          "description": "number of disk commands aborted over the measurementInterval",
          "type": "number"
        },
        "diskCommandsAvg": {
          "description": "average number of commands per second over the measurementInterval",
          "type": "number"
        },
        "diskFlushRequests": {
          "description": "total flush requests of the disk cache over the measurementInterval",
          "type": "number"
        },
        "diskFlushTime": {
          "description": "milliseconds spent on disk cache flushing over the measurementInterval",
          "type": "number"
        },
        "diskIdentifier": {
          "description": "disk identifier",
          "type": "string"
        },
        "diskIoTimeAvg": {
          "description": "milliseconds spent doing input/output operations over 1 sec; treat this metric as a device load percentage where 1000ms  matches 100% load; provide the average over the measurement interval",
          "type": "number"
        },
        "diskIoTimeLast": {
          "description": "milliseconds spent doing input/output operations over 1 sec; treat this metric as a device load percentage where 1000ms  matches 100% load; provide the last value measurement within the measurement interval",
          "type": "number"
        },
        "diskIoTimeMax": {
          "description": "milliseconds spent doing input/output operations over 1 sec; treat this metric as a device load percentage where 1000ms  matches 100% load; provide the maximum value measurement within the measurement interval",
          "type": "number"
        },
        "diskIoTimeMin": {
          "description": "milliseconds spent doing input/output operations over 1 sec; treat this metric as a device load percentage where 1000ms  matches 100% load; provide the minimum value measurement within the measurement interval",
          "type": "number"
        },
        "diskMergedReadAvg": {
          "description": "number of logical read operations that were merged into physical read operations, e.g., two logical reads were served by one physical disk access; provide the average measurement within the measurement interval",
          "type": "number"
        },
        "diskMergedReadLast": {
          "description": "number of logical read operations that were merged into physical read operations, e.g., two logical reads were served by one physical disk access; provide the last value measurement within the measurement interval",
          "type": "number"
        },
        "diskMergedReadMax": {
          "description": "number of logical read operations that were merged into physical read operations, e.g., two logical reads were served by one physical disk access; provide the maximum value measurement within the measurement interval",
          "type": "number"
        },
        "diskMergedReadMin": {
          "description": "number of logical read operations that were merged into physical read operations, e.g., two logical reads were served by one physical disk access; provide the minimum value measurement within the measurement interval",
          "type": "number"
        },
        "diskMergedWriteAvg": {
          "description": "number of logical write operations that were merged into physical write operations, e.g., two logical writes were served by one physical disk access; provide the average measurement within the measurement interval",
          "type": "number"
        },
        "diskMergedWriteLast": {
          "description": "number of logical write operations that were merged into physical write operations, e.g., two logical writes were served by one physical disk access; provide the last value measurement within the measurement interval",
          "type": "number"
        },
        "diskMergedWriteMax": {
          "description": "number of logical write operations that were merged into physical write operations, e.g., two logical writes were served by one physical disk access; provide the maximum value measurement within the measurement interval",
          "type": "number"
        },
        "diskMergedWriteMin": {
          "description": "number of logical write operations that were merged into physical write operations, e.g., two logical writes were served by one physical disk access; provide the minimum value measurement within the measurement interval",
          "type": "number"
        },
        "diskOctetsReadAvg": {
          "description": "number of octets per second read from a disk or partition; provide the average measurement within the measurement interval",
          "type": "number"
        },
        "diskOctetsReadLast": {
          "description": "number of octets per second read from a disk or partition; provide the last measurement within the measurement interval",
          "type": "number"
        },
        "diskOctetsReadMax": {
          "description": "number of octets per second read from a disk or partition; provide the maximum measurement within the measurement interval",
          "type": "number"
        },
        "diskOctetsReadMin": {
          "description": "number of octets per second read from a disk or partition; provide the minimum measurement within the measurement interval",
          "type": "number"
        },
        "diskOctetsWriteAvg": {
          "description": "number of octets per second written to a disk or partition; provide the average measurement within the measurement interval",
          "type": "number"
        },
        "diskOctetsWriteLast": {
          "description": "number of octets per second written to a disk or partition; provide the last measurement within the measurement interval",
          "type": "number"
        },
        "diskOctetsWriteMax": {
          "description": "number of octets per second written to a disk or partition; provide the maximum measurement within the measurement interval",
          "type": "number"
        },
        "diskOctetsWriteMin": {
          "description": "number of octets per second written to a disk or partition; provide the minimum measurement within the measurement interval",
          "type": "number"
        },
        "diskOpsReadAvg": {
          "description": "number of read operations per second issued to the disk; provide the average measurement within the measurement interval",
          "type": "number"
        },
        "diskOpsReadLast": {
          "description": "number of read operations per second issued to the disk; provide the last measurement within the measurement interval",
          "type": "number"
        },
        "diskOpsReadMax": {
          "description": "number of read operations per second issued to the disk; provide the maximum measurement within the measurement interval",
          "type": "number"
        },
        "diskOpsReadMin": {
          "description": "number of read operations per second issued to the disk; provide the minimum measurement within the measurement interval",
          "type": "number"
        },
        "diskOpsWriteAvg": {
          "description": "number of write operations per second issued to the disk; provide the average measurement within the measurement interval",
          "type": "number"
        },
        "diskOpsWriteLast": {
          "description": "number of write operations per second issued to the disk; provide the last measurement within the measurement interval",
          "type": "number"
        },
        "diskOpsWriteMax": {
          "description": "number of write operations per second issued to the disk; provide the maximum measurement within the measurement interval",
          "type": "number"
        },
        "diskOpsWriteMin": {
          "description": "number of write operations per second issued to the disk; provide the minimum measurement within the measurement interval",
          "type": "number"
        },
        "diskPendingOperationsAvg": {
          "description": "queue size of pending I/O operations per second; provide the average measurement within the measurement interval",
          "type": "number"
        },
        "diskPendingOperationsLast": {
          "description": "queue size of pending I/O operations per second; provide the last measurement within the measurement interval",
          "type": "number"
        },
        "diskPendingOperationsMax": {
          "description": "queue size of pending I/O operations per second; provide the maximum measurement within the measurement interval",
          "type": "number"
        },
        "diskPendingOperationsMin": {
          "description": "queue size of pending I/O operations per second; provide the minimum measurement within the measurement interval",
          "type": "number"
        },
        "diskReadCommandsAvg": {
          "description": "average number of read commands issued per second to the disk over the measurementInterval",
          "type": "number"
        },
        "diskTime": {
          "description": "nanoseconds spent on disk cache reads/writes within the measurement interval",
          "type": "number"
        },
        "diskTimeReadAvg": {
          "description": "milliseconds a read operation took to complete; provide the average measurement within the measurement interval",
          "type": "number"
        },
        "diskTimeReadLast": {
          "description": "milliseconds a read operation took to complete; provide the last measurement within the measurement interval",
          "type": "number"
        },
        "diskTimeReadMax": {
          "description": "milliseconds a read operation took to complete; provide the maximum measurement within the measurement interval",
          "type": "number"
        },
        "diskTimeReadMin": {
          "description": "milliseconds a read operation took to complete; provide the minimum measurement within the measurement interval",
          "type": "number"
        },
        "diskTimeWriteAvg": {
          "description": "milliseconds a write operation took to complete; provide the average measurement within the measurement interval",
          "type": "number"
        },
        "diskTimeWriteLast": {
          "description": "milliseconds a write operation took to complete; provide the last measurement within the measurement interval",
          "type": "number"
        },
        "diskTimeWriteMax": {
          "description": "milliseconds a write operation took to complete; provide the maximum measurement within the measurement interval",
          "type": "number"
        },
        "diskTimeWriteMin": {
          "description": "milliseconds a write operation took to complete; provide the minimum measurement within the measurement interval",
          "type": "number"
        },
        "diskTotalReadLatencyAvg": {
          "description": "average read time from the perspective of a Guest OS: sum of the Kernel Read Latency and Physical Device Read Latency in milliseconds over the measurement interval",
          "type": "number"
        },
        "diskTotalWriteLatencyAvg": {
          "description": "average write time from the perspective of a Guest OS: sum of the Kernel Write Latency and Physical Device Write Latency in milliseconds over the measurement interval",
          "type": "number"
        },
        "diskWeightedIoTimeAvg": {
          "description": "measure in ms over 1 sec of both I/O completion time and the backlog that may be accumulating; value is the average within the collection interval",
          "type": "number"
        },
        "diskWeightedIoTimeLast": {
          "description": "measure in ms over 1 sec of both I/O completion time and the backlog that may be accumulating; value is the last within the collection interval",
          "type": "number"
        },
        "diskWeightedIoTimeMax": {
          "description": "measure in ms over 1 sec of both I/O completion time and the backlog that may be accumulating; value is the maximum within the collection interval",
          "type": "number"
        },
        "diskWeightedIoTimeMin": {
          "description": "measure in ms over 1 sec of both I/O completion time and the backlog that may be accumulating; value is the minimum within the collection interval",
          "type": "number"
        },
        "diskWriteCommandsAvg": {
          "description": "average number of write commands issued per second to the disk over the measurementInterval",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "diskIdentifier"
      ]
    },
    "endOfCallVqmSummaries": {
      "description": "provides end of call voice quality metrics",
      "type": "object",
      "properties": {
        "adjacencyName": {
          "description": " adjacency name",
          "type": "string"
        },
        "endpointAverageJitter": {
          "description": "endpoint average jitter",
          "type": "number"
        },
        "endpointDescription": {
          "description": "either Caller or Callee",
          "type": "string",
          "enum": [
            "Caller",
            "Callee"
          ]
        },
        "endpointMaxJitter": {
          "description": "endpoint maximum jitter",
          "type": "number"
        },
        "endpointRtpOctetsDiscarded": {
          "description": "",
          "type": "number"
        },
        "endpointRtpOctetsLost": {
          "description": "endpoint RTP octets lost",
          "type": "number"
        },
        "endpointRtpOctetsReceived": {
          "description": "",
          "type": "number"
        },
        "endpointRtpOctetsSent": {
          "description": "",
          "type": "number"
        },
        "endpointRtpPacketsDiscarded": {
          "description": "",
          "type": "number"
        },
        "endpointRtpPacketsLost": {
          "description": "endpoint RTP packets lost",
          "type": "number"
        },
        "endpointRtpPacketsReceived": {
          "description": "",
          "type": "number"
        },
        "endpointRtpPacketsSent": {
          "description": "",
          "type": "number"
        },
        "localAverageJitter": {
          "description": "Local average jitter",
          "type": "number"
        },
        "localAverageJitterBufferDelay": {
          "description": "Local average jitter delay",
          "type": "number"
        },
        "localMaxJitter": {
          "description": "Local maximum jitter",
          "type": "number"
        },
        "localMaxJitterBufferDelay": {
          "description": "Local maximum jitter delay",
          "type": "number"
        },
        "localRtpOctetsDiscarded": {
          "description": "",
          "type": "number"
        },
        "localRtpOctetsLost": {
          "description": "Local RTP octets lost",
          "type": "number"
        },
        "localRtpOctetsReceived": {
          "description": "",
          "type": "number"
        },
        "localRtpOctetsSent": {
          "description": "",
          "type": "number"
        },
        "localRtpPacketsDiscarded": {
          "description": "",
          "type": "number"
        },
        "localRtpPacketsLost": {
          "description": "Local RTP packets lost",
          "type": "number"
        },
        "localRtpPacketsReceived": {
          "description": "",
          "type": "number"
        },
        "localRtpPacketsSent": {
          "description": "",
          "type": "number"
        },
        "mosCqe": {
          "description": "1-5 1dp",
          "type": "number"
        },
        "oneWayDelay": {
          "description": "one-way path delay in milliseconds",
          "type": "number"
        },
        "packetLossPercent": {
          "description": "Calculated percentage packet loss based on Endpoint RTP packets lost (as reported in RTCP) and Local RTP packets sent. Direction is based on Endpoint description (Caller, Callee). Decimal (2 dp)",
          "type": "number"
        },
        "rFactor": {
          "description": "0-100",
          "type": "number"
        },
        "roundTripDelay": {
          "description": "millisecs",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "adjacencyName",
        "endpointDescription"
      ]
    },
    "event": {
      "description": "the root level of the common event format",
      "type": "object",
      "properties": {
        "commonEventHeader": {
          "$ref": "#/definitions/commonEventHeader"
        },
        "faultFields": {
          "$ref": "#/definitions/faultFields"
        },
        "heartbeatFields": {
          "$ref": "#/definitions/heartbeatFields"
        },
        "measurementFields": {
          "$ref": "#/definitions/measurementFields"
        },
        "mobileFlowFields": {
          "$ref": "#/definitions/mobileFlowFields"
        },
        "notificationFields": {
          "$ref": "#/definitions/notificationFields"
        },
        "otherFields": {
          "$ref": "#/definitions/otherFields"
        },
        "perf3gppFields": {
          "$ref": "#/definitions/perf3gppFields"
        },
        "pnfRegistrationFields": {
          "$ref": "#/definitions/pnfRegistrationFields"
        },
        "sipSignalingFields": {
          "$ref": "#/definitions/sipSignalingFields"
        },
        "stateChangeFields": {
          "$ref": "#/definitions/stateChangeFields"
        },
        "syslogFields": {
          "$ref": "#/definitions/syslogFields"
        },
        "thresholdCrossingAlertFields": {
          "$ref": "#/definitions/thresholdCrossingAlertFields"
        },
        "voiceQualityFields": {
          "$ref": "#/definitions/voiceQualityFields"
        }
      },
      "additionalProperties": false,
      "required": [
        "commonEventHeader"
      ]
    },
    "eventList": {
      "description": "array of events",
      "type": "array",
      "items": {
        "$ref": "#/definitions/event"
      }
    },
    "faultFields": {
      "description": "fields specific to fault events",
      "type": "object",
      "properties": {
        "alarmAdditionalInformation": {
          "$ref": "#/definitions/hashMap"
        },
        "alarmCondition": {
          "description": "alarm condition reported by the device",
          "type": "string"
        },
        "alarmInterfaceA": {
          "description": "card, port, channel or interface name of the device generating the alarm",
          "type": "string"
        },
        "eventCategory": {
          "description": "Event category, for example: license, link, routing, security, signaling",
          "type": "string"
        },
        "eventSeverity": {
          "description": "event severity",
          "type": "string",
          "enum": [
            "CRITICAL",
            "MAJOR",
            "MINOR",
            "WARNING",
            "NORMAL"
          ]
        },
        "eventSourceType": {
          "description": "type of event source; examples: card, host, other, port, portThreshold, router, slotThreshold, switch, virtualMachine, virtualNetworkFunction",
          "type": "string"
        },
        "faultFieldsVersion": {
          "description": "version of the faultFields block",
          "type": "string",
          "enum": [
            "4.0"
          ]
        },
        "specificProblem": {
          "description": "short description of the alarm or problem",
          "type": "string"
        },
        "vfStatus": {
          "description": "virtual function status enumeration",
          "type": "string",
          "enum": [
            "Active",
            "Idle",
            "Preparing to terminate",
            "Ready to terminate",
            "Requesting termination"
          ]
        }
      },
      "additionalProperties": false,
      "required": [
        "alarmCondition",
        "eventSeverity",
        "eventSourceType",
        "faultFieldsVersion",
        "specificProblem",
        "vfStatus"
      ]
    },
    "filesystemUsage": {
      "description": "disk usage of an identified virtual machine in gigabytes and/or gigabytes per second",
      "type": "object",
      "properties": {
        "blockConfigured": {
          "type": "number"
        },
        "blockIops": {
          "type": "number"
        },
        "blockUsed": {
          "type": "number"
        },
        "ephemeralConfigured": {
          "type": "number"
        },
        "ephemeralIops": {
          "type": "number"
        },
        "ephemeralUsed": {
          "type": "number"
        },
        "filesystemName": {
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "blockConfigured",
        "blockIops",
        "blockUsed",
        "ephemeralConfigured",
        "ephemeralIops",
        "ephemeralUsed",
        "filesystemName"
      ]
    },
    "gtpPerFlowMetrics": {
      "description": "Mobility GTP Protocol per flow metrics",
      "type": "object",
      "properties": {
        "avgBitErrorRate": {
          "description": "average bit error rate",
          "type": "number"
        },
        "avgPacketDelayVariation": {
          "description": "Average packet delay variation or jitter in milliseconds for received packets: Average difference between the packet timestamp and time received for all pairs of consecutive packets",
          "type": "number"
        },
        "avgPacketLatency": {
          "description": "average delivery latency",
          "type": "number"
        },
        "avgReceiveThroughput": {
          "description": "average receive throughput",
          "type": "number"
        },
        "avgTransmitThroughput": {
          "description": "average transmit throughput",
          "type": "number"
        },
        "durConnectionFailedStatus": {
          "description": "duration of failed state in milliseconds, computed as the cumulative time between a failed echo request and the next following successful error request, over this reporting interval",
          "type": "number"
        },
        "durTunnelFailedStatus": {
          "description": "Duration of errored state, computed as the cumulative time between a tunnel error indicator and the next following non-errored indicator, over this reporting interval",
          "type": "number"
        },
        "flowActivatedBy": {
          "description": "Endpoint activating the flow",
          "type": "string"
        },
        "flowActivationEpoch": {
          "description": "Time the connection is activated in the flow (connection) being reported on, or transmission time of the first packet if activation time is not available",
          "type": "number"
        },
        "flowActivationMicrosec": {
          "description": "Integer microseconds for the start of the flow connection",
          "type": "number"
        },
        "flowActivationTime": {
          "description": "time the connection is activated in the flow being reported on, or transmission time of the first packet if activation time is not available; with RFC 2822 compliant format: Sat, 13 Mar 2010 11:29:05 -0800",
          "type": "string"
        },
        "flowDeactivatedBy": {
          "description": "Endpoint deactivating the flow",
          "type": "string"
        },
        "flowDeactivationEpoch": {
          "description": "Time for the start of the flow connection, in integer UTC epoch time aka UNIX time",
          "type": "number"
        },
        "flowDeactivationMicrosec": {
          "description": "Integer microseconds for the start of the flow connection",
          "type": "number"
        },
        "flowDeactivationTime": {
          "description": "Transmission time of the first packet in the flow connection being reported on; with RFC 2822 compliant format: Sat, 13 Mar 2010 11:29:05 -0800",
          "type": "string"
        },
        "flowStatus": {
          "description": "connection status at reporting time as a working / inactive / failed indicator value",
          "type": "string"
        },
        "gtpConnectionStatus": {
          "description": "Current connection state at reporting time",
          "type": "string"
        },
        "gtpTunnelStatus": {
          "description": "Current tunnel state  at reporting time",
          "type": "string"
        },
        "ipTosCountList": {
          "$ref": "#/definitions/hashMap"
        },
        "ipTosList": {
          "description": "Array of unique IP Type-of-Service values observed in the flow where values range from '0' to '255'",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "largePacketRtt": {
          "description": "large packet round trip time",
          "type": "number"
        },
        "largePacketThreshold": {
          "description": "large packet threshold being applied",
          "type": "number"
        },
        "maxPacketDelayVariation": {
          "description": "Maximum packet delay variation or jitter in milliseconds for received packets: Maximum of the difference between the packet timestamp and time received for all pairs of consecutive packets",
          "type": "number"
        },
        "maxReceiveBitRate": {
          "description": "maximum receive bit rate",
          "type": "number"
        },
        "maxTransmitBitRate": {
          "description": "maximum transmit bit rate",
          "type": "number"
        },
        "mobileQciCosCountList": {
          "$ref": "#/definitions/hashMap"
        },
        "mobileQciCosList": {
          "description": "Array of unique LTE QCI or UMTS class-of-service values observed in the flow",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "numActivationFailures": {
          "description": "Number of failed activation requests, as observed by the reporting node",
          "type": "number"
        },
        "numBitErrors": {
          "description": "number of errored bits",
          "type": "number"
        },
        "numBytesReceived": {
          "description": "number of bytes received, including retransmissions",
          "type": "number"
        },
        "numBytesTransmitted": {
          "description": "number of bytes transmitted, including retransmissions",
          "type": "number"
        },
        "numDroppedPackets": {
          "description": "number of received packets dropped due to errors per virtual interface",
          "type": "number"
        },
        "numGtpEchoFailures": {
          "description": "Number of Echo request path failures where failed paths are defined in 3GPP TS 29.281 sec 7.2.1 and 3GPP TS 29.060 sec. 11.2",
          "type": "number"
        },
        "numGtpTunnelErrors": {
          "description": "Number of tunnel error indications where errors are defined in 3GPP TS 29.281 sec 7.3.1 and 3GPP TS 29.060 sec. 11.1",
          "type": "number"
        },
        "numHttpErrors": {
          "description": "Http error count",
          "type": "number"
        },
        "numL7BytesReceived": {
          "description": "number of tunneled layer 7 bytes received, including retransmissions",
          "type": "number"
        },
        "numL7BytesTransmitted": {
          "description": "number of tunneled layer 7 bytes transmitted, excluding retransmissions",
          "type": "number"
        },
        "numLostPackets": {
          "description": "number of lost packets",
          "type": "number"
        },
        "numOutOfOrderPackets": {
          "description": "number of out-of-order packets",
          "type": "number"
        },
        "numPacketErrors": {
          "description": "number of errored packets",
          "type": "number"
        },
        "numPacketsReceivedExclRetrans": {
          "description": "number of packets received, excluding retransmission",
          "type": "number"
        },
        "numPacketsReceivedInclRetrans": {
          "description": "number of packets received, including retransmission",
          "type": "number"
        },
        "numPacketsTransmittedInclRetrans": {
          "description": "number of packets transmitted, including retransmissions",
          "type": "number"
        },
        "numRetries": {
          "description": "number of packet retries",
          "type": "number"
        },
        "numTimeouts": {
          "description": "number of packet timeouts",
          "type": "number"
        },
        "numTunneledL7BytesReceived": {
          "description": "number of tunneled layer 7 bytes received, excluding retransmissions",
          "type": "number"
        },
        "roundTripTime": {
          "description": "round trip time",
          "type": "number"
        },
        "tcpFlagCountList": {
          "$ref": "#/definitions/hashMap"
        },
        "tcpFlagList": {
          "description": "Array of unique TCP Flags observed in the flow",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "timeToFirstByte": {
          "description": "Time in milliseconds between the connection activation and first byte received",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "avgBitErrorRate",
        "avgPacketDelayVariation",
        "avgPacketLatency",
        "avgReceiveThroughput",
        "avgTransmitThroughput",
        "flowActivationEpoch",
        "flowActivationMicrosec",
        "flowDeactivationEpoch",
        "flowDeactivationMicrosec",
        "flowDeactivationTime",
        "flowStatus",
        "maxPacketDelayVariation",
        "numActivationFailures",
        "numBitErrors",
        "numBytesReceived",
        "numBytesTransmitted",
        "numDroppedPackets",
        "numL7BytesReceived",
        "numL7BytesTransmitted",
        "numLostPackets",
        "numOutOfOrderPackets",
        "numPacketErrors",
        "numPacketsReceivedExclRetrans",
        "numPacketsReceivedInclRetrans",
        "numPacketsTransmittedInclRetrans",
        "numRetries",
        "numTimeouts",
        "numTunneledL7BytesReceived",
        "roundTripTime",
        "timeToFirstByte"
      ]
    },
    "hashMap": {
      "description": "an associative array which is an array of key:value pairs",
      "type": "object",
      "additionalProperties": {
        "type": "string"
      },
      "default": {}
    },
    "heartbeatFields": {
      "description": "optional field block for fields specific to heartbeat events",
      "type": "object",
      "properties": {
        "additionalFields": {
          "$ref": "#/definitions/hashMap"
        },
        "heartbeatFieldsVersion": {
          "description": "version of the heartbeatFields block",
          "type": "string",
          "enum": [
            "3.0"
          ]
        },
        "heartbeatInterval": {
          "description": "current heartbeat interval in seconds",
          "type": "integer"
        }
      },
      "additionalProperties": false,
      "required": [
        "heartbeatFieldsVersion",
        "heartbeatInterval"
      ]
    },
    "hugePages": {
      "description": "metrics on system hugepages",
      "type": "object",
      "properties": {
        "bytesFree": {
          "description": "number of free hugepages in bytes",
          "type": "number"
        },
        "bytesUsed": {
          "description": "number of used hugepages in bytes",
          "type": "number"
        },
        "hugePagesIdentifier": {
          "description": "hugePages identifier",
          "type": "string"
        },
        "percentFree": {
          "description": "number of free hugepages in percent",
          "type": "number"
        },
        "percentUsed": {
          "description": "number of free hugepages in percent",
          "type": "number"
        },
        "vmPageNumberFree": {
          "description": "number of free vmPages in numbers",
          "type": "number"
        },
        "vmPageNumberUsed": {
          "description": "number of used vmPages in numbers",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "hugePagesIdentifier"
      ]
    },
    "internalHeaderFields": {
      "description": "enrichment fields for internal VES Event Listener service use only, not supplied by event sources",
      "type": "object"
    },
    "ipmi": {
      "description": "intelligent platform management interface metrics",
      "type": "object",
      "properties": {
        "exitAirTemperature": {
          "description": "system fan exit air flow temperature in celsius",
          "type": "number"
        },
        "frontPanelTemperature": {
          "description": "front panel temperature in celsius",
          "type": "number"
        },
        "ioModuleTemperature": {
          "description": "io module temperature in celsius",
          "type": "number"
        },
        "ipmiBaseboardTemperatureArray": {
          "description": "array of ipmiBaseboardTemperature objects",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ipmiBaseboardTemperature"
          }
        },
        "ipmiBaseboardVoltageRegulatorArray": {
          "description": "array of ipmiBaseboardVoltageRegulator objects",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ipmiBaseboardVoltageRegulator"
          }
        },
        "ipmiBatteryArray": {
          "description": "array of ipmiBattery objects",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ipmiBattery"
          }
        },
        "ipmiFanArray": {
          "description": "array of ipmiFan objects",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ipmiFan"
          }
        },
        "ipmiHsbpArray": {
          "description": "array of ipmiHsbp objects",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ipmiHsbp"
          }
        },
        "ipmiGlobalAggregateTemperatureMarginArray": {
          "description": "array of ipmiGlobalAggregateTemperatureMargin objects",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ipmiGlobalAggregateTemperatureMargin"
          }
        },
        "ipmiNicArray": {
          "description": "array of ipmiNic objects",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ipmiNic"
          }
        },
        "ipmiPowerSupplyArray": {
          "description": "array of ipmiPowerSupply objects",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ipmiPowerSupply"
          }
        },
        "ipmiProcessorArray": {
          "description": "array of ipmiProcessor objects",
          "type": "array",
          "items": {
            "$ref": "#/definitions/ipmiProcessor"
          }
        },
        "systemAirflow": {
          "description": "airfflow in cubic feet per minute (cfm)",
          "type": "number"
        }
      },
      "additionalProperties": false
    },
    "ipmiBaseboardTemperature": {
      "description": "intelligent platform management interface (ipmi) baseboard temperature metrics",
      "type": "object",
      "properties": {
        "baseboardTemperatureIdentifier": {
          "description": "identifier for the location where the temperature is taken",
          "type": "string"
        },
        "baseboardTemperature": {
          "description": "baseboard temperature in celsius",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "baseboardTemperatureIdentifier"
      ]
    },
    "ipmiBaseboardVoltageRegulator": {
      "description": "intelligent platform management interface (ipmi) baseboard voltage regulator metrics",
      "type": "object",
      "properties": {
        "baseboardVoltageRegulatorIdentifier": {
          "description": "identifier for the baseboard voltage regulator",
          "type": "string"
        },
        "voltageRegulatorTemperature": {
          "description": "voltage regulator temperature in celsius",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "baseboardVoltageRegulatorIdentifier"
      ]
    },
    "ipmiBattery": {
      "description": "intelligent platform management interface (ipmi) battery metrics",
      "type": "object",
      "properties": {
        "batteryIdentifier": {
          "description": "identifier for the battery",
          "type": "string"
        },
        "batteryType": {
          "description": "type of battery",
          "type": "string"
        },
        "batteryVoltageLevel": {
          "description": "battery voltage level",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "batteryIdentifier"
      ]
    },
    "ipmiFan": {
      "description": "intelligent platform management interface (ipmi) fan metrics",
      "type": "object",
      "properties": {
        "fanIdentifier": {
          "description": "identifier for the fan",
          "type": "string"
        },
        "fanSpeed": {
          "description": "fan speed in revolutions per minute (rpm)",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "fanIdentifier"
      ]
    },
    "ipmiGlobalAggregateTemperatureMargin": {
      "description": "intelligent platform management interface (ipmi) global aggregate temperature margin",
      "type": "object",
      "properties": {
        "globalAggregateTemperatureMarginIdentifier": {
          "description": "identifier for the ipmi global aggregate temperature margin metrics",
          "type": "string"
        },
        "globalAggregateTemperatureMargin": {
          "description": "the difference between the current global aggregate temperature, in celsius, and the global aggregate throttling thermal trip point",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "globalAggregateTemperatureMarginIdentifier",
        "globalAggregateTemperatureMargin"
      ]
    },
    "ipmiHsbp": {
      "description": "intelligent platform management interface (ipmi) hot swap backplane power metrics",
      "type": "object",
      "properties": {
        "hsbpIdentifier": {
          "description": "identifier for the hot swap backplane power unit",
          "type": "string"
        },
        "hsbpTemperature": {
          "description": "hot swap backplane power temperature in celsius",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "hsbpIdentifier"
      ]
    },
    "ipmiNic": {
      "description": "intelligent platform management interface (ipmi) network interface control card (nic) metrics",
      "type": "object",
      "properties": {
        "nicIdentifier": {
          "description": "identifier for the network interface control card",
          "type": "string"
        },
        "nicTemperature": {
          "description": "nic temperature in celsius",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "nicIdentifier"
      ]
    },
    "ipmiPowerSupply": {
      "description": "intelligent platform management interface (ipmi) power supply metrics",
      "type": "object",
      "properties": {
        "powerSupplyIdentifier": {
          "description": "identifier for the power supply",
          "type": "string"
        },
        "powerSupplyInputPower": {
          "description": "input power in watts",
          "type": "number"
        },
        "powerSupplyCurrentOutputPercent": {
          "description": "current output voltage as a percentage of the design specified level",
          "type": "number"
        },
        "powerSupplyTemperature": {
          "description": "power supply temperature in celsius",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "powerSupplyIdentifier"
      ]
    },
    "ipmiProcessor": {
      "description": "intelligent platform management interface processor metrics",
      "type": "object",
      "properties": {
        "processorIdentifier": {
          "description": "identifier for an ipmi processor",
          "type": "string"
        },
        "processorThermalControlPercent": {
          "description": "io module temperature in celsius",
          "type": "number"
        },
        "processorDtsThermalMargin": {
          "description": "front panel temperature in celsius",
          "type": "number"
        },
        "processorDimmAggregateThermalMarginArray": {
          "description": "array of processorDimmAggregateThermalMargin objects",
          "type": "array",
          "items": {
            "$ref": "#/definitions/processorDimmAggregateThermalMargin"
          }
        }
      },
      "additionalProperties": false,
      "required": [
        "processorIdentifier"
      ]
    },
    "jsonObject": {
      "description": "json object schema, name and other meta-information along with one or more object instances",
      "type": "object",
      "properties": {
        "objectInstances": {
          "description": "one or more instances of the jsonObject",
          "type": "array",
          "items": {
            "$ref": "#/definitions/jsonObjectInstance"
          }
        },
        "objectName": {
          "description": "name of the JSON Object",
          "type": "string"
        },
        "objectSchema": {
          "description": "json schema for the object",
          "type": "string"
        },
        "objectSchemaUrl": {
          "description": "Url to the json schema for the object",
          "type": "string"
        },
        "nfSubscribedObjectName": {
          "description": "name of the object associated with the nfSubscriptonId",
          "type": "string"
        },
        "nfSubscriptionId": {
          "description": "identifies an openConfig telemetry subscription on a network function, which configures the network function to send complex object data associated with the jsonObject",
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "objectInstances",
        "objectName"
      ]
    },
    "jsonObjectInstance": {
      "description": "meta-information about an instance of a jsonObject along with the actual object instance",
      "type": "object",
      "properties": {
        "jsonObject": {
          "$ref": "#/definitions/jsonObject"
        },
        "objectInstance": {
          "description": "an instance conforming to the jsonObject objectSchema",
          "type": "object"
        },
        "objectInstanceEpochMicrosec": {
          "description": "the unix time aka epoch time associated with this objectInstance--as microseconds elapsed since 1 Jan 1970 not including leap seconds",
          "type": "number"
        },
        "objectKeys": {
          "description": "an ordered set of keys that identifies this particular instance of jsonObject",
          "type": "array",
          "items": {
            "$ref": "#/definitions/key"
          }
        }
      },
      "additionalProperties": false
    },
    "key": {
      "description": "tuple which provides the name of a key along with its value and relative order",
      "type": "object",
      "properties": {
        "keyName": {
          "description": "name of the key",
          "type": "string"
        },
        "keyOrder": {
          "description": "relative sequence or order of the key with respect to other keys",
          "type": "integer"
        },
        "keyValue": {
          "description": "value of the key",
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "keyName"
      ]
    },
    "latencyBucketMeasure": {
      "description": "number of counts falling within a defined latency bucket",
      "type": "object",
      "properties": {
        "countsInTheBucket": {
          "type": "number"
        },
        "highEndOfLatencyBucket": {
          "type": "number"
        },
        "lowEndOfLatencyBucket": {
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "countsInTheBucket"
      ]
    },
    "load": {
      "description": "/proc/loadavg cpu utilization and io utilization metrics",
      "type": "object",
      "properties": {
        "longTerm": {
          "description": "number of jobs in the run queue (state R, cpu utilization) or waiting for disk I/O (state D, io utilization) averaged over 15 minutes using /proc/loadavg",
          "type": "number"
        },
        "midTerm": {
          "description": "number of jobs in the run queue (state R, cpu utilization) or waiting for disk I/O (state D, io utilization) averaged over 5 minutes using /proc/loadavg",
          "type": "number"
        },
        "shortTerm": {
          "description": "number of jobs in the run queue (state R, cpu utilization) or waiting for disk I/O (state D, io utilization) averaged over 1 minute using /proc/loadavg",
          "type": "number"
        }
      },
      "additionalProperties": false
    },
    "machineCheckException": {
      "description": "metrics on vm machine check exceptions",
      "type": "object",
      "properties": {
        "correctedMemoryErrors": {
          "description": "total hardware errors that were corrected by the hardware (e.g. data corruption corrected via  ECC) over the measurementInterval",
          "type": "number"
        },
        "correctedMemoryErrorsIn1Hr": {
          "description": "total hardware errors that were corrected by the hardware over the last one hour",
          "type": "number"
        },
        "uncorrectedMemoryErrors": {
          "description": "total uncorrected hardware errors that were detected by the hardware (e.g., causing data corruption) over the measurementInterval",
          "type": "number"
        },
        "uncorrectedMemoryErrorsIn1Hr": {
          "description": "total uncorrected hardware errors that were detected by the hardware over the last one hour",
          "type": "number"
        },
        "vmIdentifier": {
          "description": "virtual machine identifier associated with the machine check exception",
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "vmIdentifier"
      ]
    },
    "measDataCollection": {
      "description": "3GPP measurement collection structure aligned with 3GPP PM format",
      "type": "object",
      "properties": {
        "formatVersion": {
          "description": "3gpp PM reporting file format version from pre-standard TS 28.550 v2.0.0",
          "type": "string"
        },
        "granularityPeriod": {
          "description": "granularity period for the PM report in seconds",
          "type": "number"
        },
        "measInfoList": {
          "description": "array of measurements",
          "type": "array",
          "items": {
            "$ref": "#/definitions/measInfo"
          }
        },
        "measObjInstIdList": {
          "description": "array of monitored object local distinguished name ids per 3GPP TS 32.300",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "measuredEntityDn": {
          "description": "distinguished name per 3GPP TS 28.550",
          "type": "string"
        },
        "measuredEntitySoftwareVersion": {
          "description": "software version for the NF providing the PM data as specified in 3GPP TS 28.550",
          "type": "string"
        },
        "measuredEntityUserName": {
          "description": "user definable name for the measured object per 3GPP TS 28.550",
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "granularityPeriod",
        "measInfoList",
        "measuredEntityDn"
      ]
    },
    "measInfo": {
      "description": "measurement information.",
      "type": "object",
      "properties": {
        "jobId": {
          "description": "name of the measurement job",
          "type": "string"
        },
        "measInfoId": {
          "description": "measurement group identifier",
          "oneOf": [
            {
              "$ref": "#/definitions/measInfoIdInteger"
            },
            {
              "$ref": "#/definitions/measInfoIdString"
            }
          ]
        },
        "measTypes": {
          "oneOf": [
            {
              "$ref": "#/definitions/measTypesInteger"
            },
            {
              "$ref": "#/definitions/measTypesString"
            }
          ]
        },
        "measValuesList": {
          "description": "an array of measurement values",
          "type": "array",
          "items": {
            "$ref": "#/definitions/measValues"
          }
        }
      },
      "additionalProperties": false,
      "required": [
        "measTypes",
        "measValuesList"
      ]
    },
    "measInfoIdInteger": {
      "description": "integer measurement group identifier",
      "type": "object",
      "properties": {
        "iMeasInfoId": {
          "type": "integer"
        }
      },
      "additionalProperties": false,
      "required": [
        "iMeasInfoId"
      ]
    },
    "measInfoIdString": {
      "description": "string measurement group identifier",
      "type": "object",
      "properties": {
        "sMeasInfoId": {
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "sMeasInfoId"
      ]
    },
    "measResultInteger": {
      "description": "integer 3GPP PM measurement result",
      "type": "object",
      "properties": {
        "p": {
          "description": "integer reference to the counter",
          "type": "integer"
        },
        "iValue": {
          "description": "integer counter value",
          "type": "integer"
        }
      },
      "additionalProperties": false,
      "required": [
        "p",
        "iValue"
      ]
    },
    "measResultNull": {
      "description": "null 3GPP PM measurement result",
      "type": "object",
      "properties": {
        "p": {
          "description": "integer reference to the counter",
          "type": "integer"
        },
        "isNull": {
          "description": "true if the counter has no value",
          "type": "string",
          "enum": [
            "true",
            "false"
          ]
        }
      },
      "additionalProperties": false,
      "required": [
        "p",
        "isNull"
      ]
    },
    "measResultNumber": {
      "description": "number 3GPP PM measurement result",
      "type": "object",
      "properties": {
        "p": {
          "description": "integer reference to the counter",
          "type": "integer"
        },
        "rValue": {
          "description": "numeric counter value",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "p",
        "rValue"
      ]
    },
    "measResultString": {
      "description": "string 3GPP PM measurement result",
      "type": "object",
      "properties": {
        "p": {
          "description": "integer reference to the counter",
          "type": "integer"
        },
        "sValue": {
          "description": "string counter value",
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "p",
        "sValue"
      ]
    },
    "measTypesInteger": {
      "description": "object containing an array of integer measurement identifiers associated with the measurement results",
      "type": "object",
      "properties": {
        "iMeasTypesList": {
          "type": "array",
          "items": {
            "type": "integer"
          }
        }
      },
      "additionalProperties": false,
      "required": [
        "iMeasTypesList"
      ]
    },
    "measTypesString": {
      "description": "object containing an array of string measurement identifiers associated with the measurement results",
      "type": "object",
      "properties": {
        "sMeasTypesList": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      },
      "additionalProperties": false,
      "required": [
        "sMeasTypesList"
      ]
    },
    "measurementFields": {
      "description": "measurement fields",
      "type": "object",
      "properties": {
        "additionalFields": {
          "$ref": "#/definitions/hashMap"
        },
        "additionalMeasurements": {
          "$ref": "#/definitions/arrayOfNamedHashMap"
        },
        "additionalObjects": {
          "$ref": "#/definitions/arrayOfJsonObject"
        },
        "codecUsageArray": {
          "description": "array of codecs in use",
          "type": "array",
          "items": {
            "$ref": "#/definitions/codecsInUse"
          }
        },
        "concurrentSessions": {
          "description": "peak concurrent sessions for the VM or xNF over the measurementInterval",
          "type": "integer"
        },
        "configuredEntities": {
          "description": "over the measurementInterval, peak total number of: users, subscribers, devices, adjacencies, etc., for the VM, or subscribers, devices, etc., for the xNF",
          "type": "integer"
        },
        "cpuUsageArray": {
          "description": "usage of an array of CPUs",
          "type": "array",
          "items": {
            "$ref": "#/definitions/cpuUsage"
          }
        },
        "diskUsageArray": {
          "description": "usage of an array of disks",
          "type": "array",
          "items": {
            "$ref": "#/definitions/diskUsage"
          }
        },
        "featureUsageArray": {
          "$ref": "#/definitions/hashMap"
        },
        "filesystemUsageArray": {
          "description": "filesystem usage of the VM on which the xNFC reporting the event is running",
          "type": "array",
          "items": {
            "$ref": "#/definitions/filesystemUsage"
          }
        },
        "hugePagesArray": {
          "description": "array of metrics on hugepPages",
          "type": "array",
          "items": {
            "$ref": "#/definitions/hugePages"
          }
        },
        "ipmi": {
          "$ref": "#/definitions/ipmi"
        },
        "latencyDistribution": {
          "description": "array of integers representing counts of requests whose latency in milliseconds falls within per-xNF configured ranges",
          "type": "array",
          "items": {
            "$ref": "#/definitions/latencyBucketMeasure"
          }
        },
        "loadArray": {
          "description": "array of system load metrics",
          "type": "array",
          "items": {
            "$ref": "#/definitions/load"
          }
        },
        "machineCheckExceptionArray": {
          "description": "array of machine check exceptions",
          "type": "array",
          "items": {
            "$ref": "#/definitions/machineCheckException"
          }
        },
        "meanRequestLatency": {
          "description": "mean seconds required to respond to each request for the VM on which the xNFC reporting the event is running",
          "type": "number"
        },
        "measurementInterval": {
          "description": "interval over which measurements are being reported in seconds",
          "type": "number"
        },
        "measurementFieldsVersion": {
          "description": "version of the measurementFields block",
          "type": "string",
          "enum": [
            "4.0"
          ]
        },
        "memoryUsageArray": {
          "description": "memory usage of an array of VMs",
          "type": "array",
          "items": {
            "$ref": "#/definitions/memoryUsage"
          }
        },
        "numberOfMediaPortsInUse": {
          "description": "number of media ports in use",
          "type": "integer"
        },
        "requestRate": {
          "description": "peak rate of service requests per second to the xNF over the measurementInterval",
          "type": "number"
        },
        "nfcScalingMetric": {
          "description": "represents busy-ness of the network function from 0 to 100 as reported by the xNFC",
          "type": "integer"
        },
        "nicPerformanceArray": {
          "description": "usage of an array of network interface cards",
          "type": "array",
          "items": {
            "$ref": "#/definitions/nicPerformance"
          }
        },
        "processStatsArray": {
          "description": "array of metrics on system processes",
          "type": "array",
          "items": {
            "$ref": "#/definitions/processStats"
          }
        }
      },
      "additionalProperties": false,
      "required": [
        "measurementInterval",
        "measurementFieldsVersion"
      ]
    },
    "measValues": {
      "description": "3GPP measurement values",
      "type": "object",
      "properties": {
        "measObjAddlFlds": {
          "$ref": "#/definitions/hashMap"
        },
        "measObjInstId": {
          "description": "monitored object local distinguished name per 3GPP TS 32.300 and 3GPP TS 32.432",
          "type": "string"
        },
        "measResults": {
          "description": "array of results",
          "type": "array",
          "items": {
            "oneOf": [
              {
                "$ref": "#/definitions/measResultInteger"
              },
              {
                "$ref": "#/definitions/measResultNull"
              },
              {
                "$ref": "#/definitions/measResultNumber"
              },
              {
                "$ref": "#/definitions/measResultString"
              }
            ]
          }
        },
        "suspectFlag": {
          "description": "indicates if the values are suspect",
          "type": "string",
          "enum": [
            "true",
            "false"
          ]
        }
      },
      "additionalProperties": false,
      "required": [
        "measObjInstId",
        "measResults"
      ]
    },
    "memoryUsage": {
      "description": "memory usage of an identified virtual machine",
      "type": "object",
      "properties": {
        "memoryBuffered": {
          "description": "kibibytes of temporary storage for raw disk blocks",
          "type": "number"
        },
        "memoryCached": {
          "description": "kibibytes of memory used for cache",
          "type": "number"
        },
        "memoryConfigured": {
          "description": "kibibytes of memory configured in the virtual machine on which the xNFC reporting the event is running",
          "type": "number"
        },
        "memoryDemand": {
          "description": "host demand in kibibytes",
          "type": "number"
        },
        "memoryFree": {
          "description": "kibibytes of physical RAM left unused by the system",
          "type": "number"
        },
        "memoryLatencyAvg": {
          "description": "Percentage of time the VM is waiting to access swapped or compressed memory",
          "type": "number"
        },
        "memorySharedAvg": {
          "description": "shared memory in kilobytes",
          "type": "number"
        },
        "memorySlabRecl": {
          "description": "the part of the slab that can be reclaimed such as caches measured in kibibytes",
          "type": "number"
        },
        "memorySlabUnrecl": {
          "description": "the part of the slab that cannot be reclaimed even when lacking memory measured in kibibytes",
          "type": "number"
        },
        "memorySwapInAvg": {
          "description": "Amount of memory swapped-in from host cache in kibibytes",
          "type": "number"
        },
        "memorySwapInRateAvg": {
          "description": "rate at which memory is swapped from disk into active memory during the interval in kilobytes per second",
          "type": "number"
        },
        "memorySwapOutAvg": {
          "description": "Amount of memory swapped-out to host cache in kibibytes",
          "type": "number"
        },
        "memorySwapOutRateAvg": {
          "description": "rate at which memory is being swapped from active memory to disk during the current interval in kilobytes per second",
          "type": "number"
        },
        "memorySwapUsedAvg": {
          "description": "space used for caching swapped pages in the host cache in kibibytes",
          "type": "number"
        },
        "memoryUsed": {
          "description": "total memory minus the sum of free, buffered, cached and slab memory measured in kibibytes",
          "type": "number"
        },
        "percentMemoryUsage": {
          "description": "Percentage of memory usage; value = (memoryUsed / (memoryUsed + memoryFree) x 100 if denomintor is nonzero, or 0, if otherwise",
          "type": "number"
        },
        "vmIdentifier": {
          "description": "virtual machine identifier associated with the memory metrics",
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "memoryFree",
        "memoryUsed",
        "vmIdentifier"
      ]
    },
    "mobileFlowFields": {
      "description": "mobileFlow fields",
      "type": "object",
      "properties": {
        "additionalFields": {
          "$ref": "#/definitions/hashMap"
        },
        "applicationType": {
          "description": "Application type inferred",
          "type": "string"
        },
        "appProtocolType": {
          "description": "application protocol",
          "type": "string"
        },
        "appProtocolVersion": {
          "description": "application protocol version",
          "type": "string"
        },
        "cid": {
          "description": "cell id",
          "type": "string"
        },
        "connectionType": {
          "description": "Abbreviation referencing a 3GPP reference point e.g., S1-U, S11, etc",
          "type": "string"
        },
        "ecgi": {
          "description": "Evolved Cell Global Id",
          "type": "string"
        },
        "flowDirection": {
          "description": "Flow direction, indicating if the reporting node is the source of the flow or destination for the flow",
          "type": "string"
        },
        "gtpPerFlowMetrics": {
          "$ref": "#/definitions/gtpPerFlowMetrics"
        },
        "gtpProtocolType": {
          "description": "GTP protocol",
          "type": "string"
        },
        "gtpVersion": {
          "description": "GTP protocol version",
          "type": "string"
        },
        "httpHeader": {
          "description": "HTTP request header, if the flow connects to a node referenced by HTTP",
          "type": "string"
        },
        "imei": {
          "description": "IMEI for the subscriber UE used in this flow, if the flow connects to a mobile device",
          "type": "string"
        },
        "imsi": {
          "description": "IMSI for the subscriber UE used in this flow, if the flow connects to a mobile device",
          "type": "string"
        },
        "ipProtocolType": {
          "description": "IP protocol type e.g., TCP, UDP, RTP...",
          "type": "string"
        },
        "ipVersion": {
          "description": "IP protocol version e.g., IPv4, IPv6",
          "type": "string"
        },
        "lac": {
          "description": "location area code",
          "type": "string"
        },
        "mcc": {
          "description": "mobile country code",
          "type": "string"
        },
        "mnc": {
          "description": "mobile network code",
          "type": "string"
        },
        "mobileFlowFieldsVersion": {
          "description": "version of the mobileFlowFields block",
          "type": "string",
          "enum": [
            "4.0"
          ]
        },
        "msisdn": {
          "description": "MSISDN for the subscriber UE used in this flow, as an integer, if the flow connects to a mobile device",
          "type": "string"
        },
        "otherEndpointIpAddress": {
          "description": "IP address for the other endpoint, as used for the flow being reported on",
          "type": "string"
        },
        "otherEndpointPort": {
          "description": "IP Port for the reporting entity, as used for the flow being reported on",
          "type": "integer"
        },
        "otherFunctionalRole": {
          "description": "Functional role of the other endpoint for the flow being reported on e.g., MME, S-GW, P-GW, PCRF...",
          "type": "string"
        },
        "rac": {
          "description": "routing area code",
          "type": "string"
        },
        "radioAccessTechnology": {
          "description": "Radio Access Technology e.g., 2G, 3G, LTE",
          "type": "string"
        },
        "reportingEndpointIpAddr": {
          "description": "IP address for the reporting entity, as used for the flow being reported on",
          "type": "string"
        },
        "reportingEndpointPort": {
          "description": "IP port for the reporting entity, as used for the flow being reported on",
          "type": "integer"
        },
        "sac": {
          "description": "service area code",
          "type": "string"
        },
        "samplingAlgorithm": {
          "description": "Integer identifier for the sampling algorithm or rule being applied in calculating the flow metrics if metrics are calculated based on a sample of packets, or 0 if no sampling is applied",
          "type": "integer"
        },
        "tac": {
          "description": "transport area code",
          "type": "string"
        },
        "tunnelId": {
          "description": "tunnel identifier",
          "type": "string"
        },
        "vlanId": {
          "description": "VLAN identifier used by this flow",
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "flowDirection",
        "gtpPerFlowMetrics",
        "ipProtocolType",
        "ipVersion",
        "mobileFlowFieldsVersion",
        "otherEndpointIpAddress",
        "otherEndpointPort",
        "reportingEndpointIpAddr",
        "reportingEndpointPort"
      ]
    },
    "namedHashMap": {
      "description": "a hashMap which is associated with and described by a name",
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "hashMap": {
          "$ref": "#/definitions/hashMap"
        }
      },
      "additionalProperties": false,
      "required": [
        "name",
        "hashMap"
      ]
    },
    "nicPerformance": {
      "description": "describes the performance and errors of an identified network interface card",
      "type": "object",
      "properties": {
        "administrativeState": {
          "description": "administrative state",
          "type": "string",
          "enum": [
            "inService",
            "outOfService"
          ]
        },
        "nicIdentifier": {
          "description": "nic identification",
          "type": "string"
        },
        "operationalState": {
          "description": "operational state",
          "type": "string",
          "enum": [
            "inService",
            "outOfService"
          ]
        },
        "receivedBroadcastPacketsAccumulated": {
          "description": "Cumulative count of broadcast packets received as read at the end of the measurement interval",
          "type": "number"
        },
        "receivedBroadcastPacketsDelta": {
          "description": "Count of broadcast packets received within the measurement interval",
          "type": "number"
        },
        "receivedDiscardedPacketsAccumulated": {
          "description": "Cumulative count of discarded packets received as read at the end of the measurement interval",
          "type": "number"
        },
        "receivedDiscardedPacketsDelta": {
          "description": "Count of discarded packets received within the measurement interval",
          "type": "number"
        },
        "receivedErrorPacketsAccumulated": {
          "description": "Cumulative count of error packets received as read at the end of the measurement interval",
          "type": "number"
        },
        "receivedErrorPacketsDelta": {
          "description": "Count of error packets received within the measurement interval",
          "type": "number"
        },
        "receivedMulticastPacketsAccumulated": {
          "description": "Cumulative count of multicast packets received as read at the end of the measurement interval",
          "type": "number"
        },
        "receivedMulticastPacketsDelta": {
          "description": "Count of multicast packets received within the measurement interval",
          "type": "number"
        },
        "receivedOctetsAccumulated": {
          "description": "Cumulative count of octets received as read at the end of the measurement interval",
          "type": "number"
        },
        "receivedOctetsDelta": {
          "description": "Count of octets received within the measurement interval",
          "type": "number"
        },
        "receivedTotalPacketsAccumulated": {
          "description": "Cumulative count of all packets received as read at the end of the measurement interval",
          "type": "number"
        },
        "receivedPercentDiscard": {
          "description": "Percentage of discarded packets received; value = (receivedDiscardedPacketsDelta / receivedTotalPacketsDelta) x 100, if denominator is nonzero, or 0, if otherwise",
          "type": "number"
        },
        "receivedPercentError": {
          "description": "Percentage of error packets received; value = (receivedErrorPacketsDelta / receivedTotalPacketsDelta) x 100, if denominator is nonzero, or 0, if otherwise.",
          "type": "number"
        },
        "receivedTotalPacketsDelta": {
          "description": "Count of all packets received within the measurement interval",
          "type": "number"
        },
        "receivedUnicastPacketsAccumulated": {
          "description": "Cumulative count of unicast packets received as read at the end of the measurement interval",
          "type": "number"
        },
        "receivedUnicastPacketsDelta": {
          "description": "Count of unicast packets received within the measurement interval",
          "type": "number"
        },
        "receivedUtilization": {
          "description": "Percentage of utilization received; value = (receivedOctetsDelta / (speed x (lastEpochMicrosec - startEpochMicrosec))) x 100, if denominator is nonzero, or 0, if otherwise",
          "type": "number"
        },
        "speed": {
          "description": "Speed configured in mbps",
          "type": "number"
        },
        "transmittedBroadcastPacketsAccumulated": {
          "description": "Cumulative count of broadcast packets transmitted as read at the end of the measurement interval",
          "type": "number"
        },
        "transmittedBroadcastPacketsDelta": {
          "description": "Count of broadcast packets transmitted within the measurement interval",
          "type": "number"
        },
        "transmittedDiscardedPacketsAccumulated": {
          "description": "Cumulative count of discarded packets transmitted as read at the end of the measurement interval",
          "type": "number"
        },
        "transmittedDiscardedPacketsDelta": {
          "description": "Count of discarded packets transmitted within the measurement interval",
          "type": "number"
        },
        "transmittedErrorPacketsAccumulated": {
          "description": "Cumulative count of error packets transmitted as read at the end of the measurement interval",
          "type": "number"
        },
        "transmittedErrorPacketsDelta": {
          "description": "Count of error packets transmitted within the measurement interval",
          "type": "number"
        },
        "transmittedMulticastPacketsAccumulated": {
          "description": "Cumulative count of multicast packets transmitted as read at the end of the measurement interval",
          "type": "number"
        },
        "transmittedMulticastPacketsDelta": {
          "description": "Count of multicast packets transmitted within the measurement interval",
          "type": "number"
        },
        "transmittedOctetsAccumulated": {
          "description": "Cumulative count of octets transmitted as read at the end of the measurement interval",
          "type": "number"
        },
        "transmittedOctetsDelta": {
          "description": "Count of octets transmitted within the measurement interval",
          "type": "number"
        },
        "transmittedTotalPacketsAccumulated": {
          "description": "Cumulative count of all packets transmitted as read at the end of the measurement interval",
          "type": "number"
        },
        "transmittedTotalPacketsDelta": {
          "description": "Count of all packets transmitted within the measurement interval",
          "type": "number"
        },
        "transmittedUnicastPacketsAccumulated": {
          "description": "Cumulative count of unicast packets transmitted as read at the end of the measurement interval",
          "type": "number"
        },
        "transmittedUnicastPacketsDelta": {
          "description": "Count of unicast packets transmitted within the measurement interval",
          "type": "number"
        },
        "transmittedPercentDiscard": {
          "description": "Percentage of discarded packets transmitted; value = (transmittedDiscardedPacketsDelta / transmittedTotalPacketsDelta) x 100, if denominator is nonzero, or 0, if otherwise",
          "type": "number"
        },
        "transmittedPercentError": {
          "description": "Percentage of error packets received; value = (transmittedErrorPacketsDelta / transmittedTotalPacketsDelta) x 100, if denominator is nonzero, or 0, if otherwise",
          "type": "number"
        },
        "transmittedUtilization": {
          "description": "Percentage of utilization transmitted; value = (transmittedOctetsDelta / (speed x (lastEpochMicrosec - startEpochMicrosec))) x 100, if denominator is nonzero, or 0, if otherwise.",
          "type": "number"
        },
        "valuesAreSuspect": {
          "description": "Indicates whether vNicPerformance values are likely inaccurate due to counter overflow or other condtions",
          "type": "string",
          "enum": [
            "true",
            "false"
          ]
        }
      },
      "additionalProperties": false,
      "required": [
        "nicIdentifier",
        "valuesAreSuspect"
      ]
    },
    "notificationFields": {
      "description": "notification fields",
      "type": "object",
      "properties": {
        "additionalFields": {
          "$ref": "#/definitions/hashMap"
        },
        "arrayOfNamedHashMap": {
          "$ref": "#/definitions/arrayOfNamedHashMap"
        },
        "changeContact": {
          "description": "identifier for a contact related to the change",
          "type": "string"
        },
        "changeIdentifier": {
          "description": "system or session identifier associated with the change",
          "type": "string"
        },
        "changeType": {
          "description": "describes what has changed for the entity",
          "type": "string"
        },
        "newState": {
          "description": "new state of the entity",
          "type": "string"
        },
        "oldState": {
          "description": "previous state of the entity",
          "type": "string"
        },
        "notificationFieldsVersion": {
          "description": "version of the notificationFields block",
          "type": "string",
          "enum": [
            "2.0"
          ]
        },
        "stateInterface": {
          "description": "card or port name of the entity that changed state",
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "changeIdentifier",
        "changeType",
        "notificationFieldsVersion"
      ]
    },
    "otherFields": {
      "description": "fields for events belonging to the 'other' domain of the commonEventHeader domain enumeration",
      "type": "object",
      "properties": {
        "arrayOfNamedHashMap": {
          "$ref": "#/definitions/arrayOfNamedHashMap"
        },
        "hashMap": {
          "$ref": "#/definitions/hashMap"
        },
        "jsonObjects": {
          "$ref": "#/definitions/arrayOfJsonObject"
        },
        "otherFieldsVersion": {
          "description": "version of the otherFields block",
          "type": "string",
          "enum": [
            "3.0"
          ]
        }
      },
      "additionalProperties": false,
      "required": [
        "otherFieldsVersion"
      ]
    },
    "perf3gppFields": {
      "description": "fields for 3GPP PM format events, based on 3GPP TS 28.550, belonging to the 'perf3gpp' domain of the commonEventHeader domain enumeration",
      "type": "object",
      "properties": {
        "eventAddlFields": {
          "$ref": "#/definitions/hashMap"
        },
        "measDataCollection": {
          "$ref": "#/definitions/measDataCollection"
        },
        "perf3gppFieldsVersion": {
          "description": "version of the perf3gppFields block",
          "type": "string",
          "enum": [
            "1.0",
            "1.0.1"
          ]
        }
      },
      "additionalProperties": false,
      "required": [
        "measDataCollection",
        "perf3gppFieldsVersion"
      ]
    },
    "pnfRegistrationFields": {
      "description": "hardware device registration fields",
      "type": "object",
      "properties": {
        "additionalFields": {
          "$ref": "#/definitions/hashMap"
        },
        "lastServiceDate": {
          "description": "TS 32.692 dateOfLastService = date of last service; e.g. 15022017",
          "type": "string"
        },
        "macAddress": {
          "description": "MAC address of OAM interface of the unit",
          "type": "string"
        },
        "manufactureDate": {
          "description": "TS 32.692 dateOfManufacture = manufacture date of the unit; 24032016",
          "type": "string"
        },
        "modelNumber": {
          "description": "TS 32.692 versionNumber = version of the unit from vendor; e.g. AJ02.  Maps to AAI equip-model",
          "type": "string"
        },
        "oamV4IpAddress": {
          "description": "IPv4 m-plane IP address to be used by the manager to contact the PNF",
          "type": "string"
        },
        "oamV6IpAddress": {
          "description": "IPv6 m-plane IP address to be used by the manager to contact the PNF",
          "type": "string"
        },
        "pnfRegistrationFieldsVersion": {
          "description": "version of the pnfRegistrationFields block",
          "type": "string",
          "enum": [
            "2.0"
          ]
        },
        "serialNumber": {
          "description": "TS 32.692 serialNumber = serial number of the unit; e.g. 6061ZW3",
          "type": "string"
        },
        "softwareVersion": {
          "description": "TS 32.692 swName = active SW running on the unit; e.g. 5gDUv18.05.201",
          "type": "string"
        },
        "unitFamily": {
          "description": "TS 32.692 vendorUnitFamilyType = general type of HW unit; e.g. BBU",
          "type": "string"
        },
        "unitType": {
          "description": "TS 32.692 vendorUnitTypeNumber = vendor name for the unit; e.g. Airscale",
          "type": "string"
        },
        "vendorName": {
          "description": "TS 32.692 vendorName = name of manufacturer; e.g. Nokia. Maps to AAI equip-vendor",
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "pnfRegistrationFieldsVersion"
      ]
    },
    "processorDimmAggregateThermalMargin": {
      "description": "intelligent platform management interface (ipmi) processor dual inline memory module aggregate thermal margin metrics",
      "type": "object",
      "properties": {
        "processorDimmAggregateThermalMarginIdentifier": {
          "description": "identifier for the aggregate thermal margin metrics from the processor dual inline memory module",
          "type": "string"
        },
        "thermalMargin": {
          "description": "the difference between the DIMM's current temperature, in celsius, and the DIMM's throttling thermal trip point",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "processorDimmAggregateThermalMarginIdentifier",
        "thermalMargin"
      ]
    },
    "processStats": {
      "description": "metrics on system processes",
      "type": "object",
      "properties": {
        "forkRate": {
          "description": "the number of threads created since the last reboot",
          "type": "number"
        },
        "processIdentifier": {
          "description": "processIdentifier",
          "type": "string"
        },
        "psStateBlocked": {
          "description": "the number of processes in a blocked state",
          "type": "number"
        },
        "psStatePaging": {
          "description": "the number of processes in a paging state",
          "type": "number"
        },
        "psStateRunning": {
          "description": "the number of processes in a running state",
          "type": "number"
        },
        "psStateSleeping": {
          "description": "the number of processes in a sleeping state",
          "type": "number"
        },
        "psStateStopped": {
          "description": "the number of processes in a stopped state",
          "type": "number"
        },
        "psStateZombie": {
          "description": "the number of processes in a zombie state",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "processIdentifier"
      ]
    },
    "requestError": {
      "description": "standard request error data structure",
      "type": "object",
      "properties": {
        "messageId": {
          "description": "Unique message identifier of the format ABCnnnn where ABC is either SVC for Service Exceptions or POL for Policy Exception",
          "type": "string"
        },
        "text": {
          "description": "Message text, with replacement variables marked with %n, where n is an index into the list of <variables> elements, starting at 1",
          "type": "string"
        },
        "url": {
          "description": "Hyperlink to a detailed error resource e.g., an HTML page for browser user agents",
          "type": "string"
        },
        "variables": {
          "description": "List of zero or more strings that represent the contents of the variables used by the message text",
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "messageId",
        "text"
      ]
    },
    "sipSignalingFields": {
      "description": "sip signaling fields",
      "type": "object",
      "properties": {
        "additionalInformation": {
          "$ref": "#/definitions/hashMap"
        },
        "compressedSip": {
          "description": "the full SIP request/response including headers and bodies",
          "type": "string"
        },
        "correlator": {
          "description": "this is the same for all events on this call",
          "type": "string"
        },
        "localIpAddress": {
          "description": "IP address on xNF",
          "type": "string"
        },
        "localPort": {
          "description": "port on xNF",
          "type": "string"
        },
        "remoteIpAddress": {
          "description": "IP address of peer endpoint",
          "type": "string"
        },
        "remotePort": {
          "description": "port of peer endpoint",
          "type": "string"
        },
        "sipSignalingFieldsVersion": {
          "description": "version of the sipSignalingFields block",
          "type": "string",
          "enum": [
            "3.0"
          ]
        },
        "summarySip": {
          "description": "the SIP Method or Response ('INVITE', '200 OK', 'BYE', etc)",
          "type": "string"
        },
        "vendorNfNameFields": {
          "$ref": "#/definitions/vendorNfNameFields"
        }
      },
      "additionalProperties": false,
      "required": [
        "correlator",
        "localIpAddress",
        "localPort",
        "remoteIpAddress",
        "remotePort",
        "sipSignalingFieldsVersion",
        "vendorNfNameFields"
      ]
    },
    "stateChangeFields": {
      "description": "stateChange fields",
      "type": "object",
      "properties": {
        "additionalFields": {
          "$ref": "#/definitions/hashMap"
        },
        "newState": {
          "description": "new state of the entity",
          "type": "string",
          "enum": [
            "inService",
            "maintenance",
            "outOfService"
          ]
        },
        "oldState": {
          "description": "previous state of the entity",
          "type": "string",
          "enum": [
            "inService",
            "maintenance",
            "outOfService"
          ]
        },
        "stateChangeFieldsVersion": {
          "description": "version of the stateChangeFields block",
          "type": "string",
          "enum": [
            "4.0"
          ]
        },
        "stateInterface": {
          "description": "card or port name of the entity that changed state",
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "newState",
        "oldState",
        "stateChangeFieldsVersion",
        "stateInterface"
      ]
    },
    "syslogFields": {
      "description": "sysLog fields",
      "type": "object",
      "properties": {
        "additionalFields": {
          "$ref": "#/definitions/hashMap"
        },
        "eventSourceHost": {
          "description": "hostname of the device",
          "type": "string"
        },
        "eventSourceType": {
          "description": "type of event source; examples: other, router, switch, host, card, port, slotThreshold, portThreshold, virtualMachine, virtualNetworkFunction",
          "type": "string"
        },
        "syslogFacility": {
          "description": "numeric code from 0 to 23 for facility--see table in documentation",
          "type": "integer"
        },
        "syslogFieldsVersion": {
          "description": "version of the syslogFields block",
          "type": "string",
          "enum": [
            "4.0"
          ]
        },
        "syslogMsg": {
          "description": "syslog message",
          "type": "string"
        },
        "syslogMsgHost": {
          "description": "hostname parsed from non-VES syslog message",
          "type": "string"
        },
        "syslogPri": {
          "description": "0-192 combined severity and facility",
          "type": "integer"
        },
        "syslogProc": {
          "description": "identifies the application that originated the message",
          "type": "string"
        },
        "syslogProcId": {
          "description": "a change in the value of this field indicates a discontinuity in syslog reporting",
          "type": "number"
        },
        "syslogSData": {
          "description": "syslog structured data consisting of a structured data Id followed by a set of key value pairs",
          "type": "string"
        },
        "syslogSdId": {
          "description": "0-32 char in format name@number for example ourSDID@32473",
          "type": "string"
        },
        "syslogSev": {
          "description": "numerical Code for  severity derived from syslogPri as remaider of syslogPri / 8",
          "type": "string",
          "enum": [
            "Alert",
            "Critical",
            "Debug",
            "Emergency",
            "Error",
            "Info",
            "Notice",
            "Warning"
          ]
        },
        "syslogTag": {
          "description": "msgId indicating the type of message such as TCPOUT or TCPIN; NILVALUE should be used when no other value can be provided",
          "type": "string"
        },
        "syslogTs": {
          "description": "timestamp parsed from non-VES syslog message",
          "type": "string"
        },
        "syslogVer": {
          "description": "IANA assigned version of the syslog protocol specification - typically 1",
          "type": "number"
        }
      },
      "additionalProperties": false,
      "required": [
        "eventSourceType",
        "syslogFieldsVersion",
        "syslogMsg",
        "syslogTag"
      ]
    },
    "thresholdCrossingAlertFields": {
      "description": "fields specific to threshold crossing alert events",
      "type": "object",
      "properties": {
        "additionalFields": {
          "$ref": "#/definitions/hashMap"
        },
        "additionalParameters": {
          "description": "performance counters",
          "type": "array",
          "items": {
            "$ref": "#/definitions/counter"
          }
        },
        "alertAction": {
          "description": "Event action",
          "type": "string",
          "enum": [
            "CLEAR",
            "CONT",
            "SET"
          ]
        },
        "alertDescription": {
          "description": "Unique short alert description such as IF-SHUB-ERRDROP",
          "type": "string"
        },
        "alertType": {
          "description": "Event type",
          "type": "string",
          "enum": [
            "CARD-ANOMALY",
            "ELEMENT-ANOMALY",
            "INTERFACE-ANOMALY",
            "SERVICE-ANOMALY"
          ]
        },
        "alertValue": {
          "description": "Calculated API value (if applicable)",
          "type": "string"
        },
        "associatedAlertIdList": {
          "description": "List of eventIds associated with the event being reported",
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "collectionTimestamp": {
          "description": "Time when the performance collector picked up the data; with RFC 2822 compliant format: Sat, 13 Mar 2010 11:29:05 -0800",
          "type": "string"
        },
        "dataCollector": {
          "description": "Specific performance collector instance used",
          "type": "string"
        },
        "elementType": {
          "description": "type of network element - internal ATT field",
          "type": "string"
        },
        "eventSeverity": {
          "description": "event severity or priority",
          "type": "string",
          "enum": [
            "CRITICAL",
            "MAJOR",
            "MINOR",
            "WARNING",
            "NORMAL"
          ]
        },
        "eventStartTimestamp": {
          "description": "Time closest to when the measurement was made; with RFC 2822 compliant format: Sat, 13 Mar 2010 11:29:05 -0800",
          "type": "string"
        },
        "interfaceName": {
          "description": "Physical or logical port or card (if applicable)",
          "type": "string"
        },
        "networkService": {
          "description": "network name - internal ATT field",
          "type": "string"
        },
        "possibleRootCause": {
          "description": "Reserved for future use",
          "type": "string"
        },
        "thresholdCrossingFieldsVersion": {
          "description": "version of the thresholdCrossingAlertFields block",
          "type": "string",
          "enum": [
            "4.0"
          ]
        }
      },
      "additionalProperties": false,
      "required": [
        "additionalParameters",
        "alertAction",
        "alertDescription",
        "alertType",
        "collectionTimestamp",
        "eventSeverity",
        "eventStartTimestamp",
        "thresholdCrossingFieldsVersion"
      ]
    },
    "vendorNfNameFields": {
      "description": "provides vendor, nf and nfModule identifying information",
      "type": "object",
      "properties": {
        "vendorName": {
          "description": "network function vendor name",
          "type": "string"
        },
        "nfModuleName": {
          "description": "name of the nfModule generating the event",
          "type": "string"
        },
        "nfName": {
          "description": "name of the network function generating the event",
          "type": "string"
        }
      },
      "additionalProperties": false,
      "required": [
        "vendorName"
      ]
    },
    "voiceQualityFields": {
      "description": "provides statistics related to customer facing voice products",
      "type": "object",
      "properties": {
        "additionalInformation": {
          "$ref": "#/definitions/hashMap"
        },
        "calleeSideCodec": {
          "description": "callee codec for the call",
          "type": "string"
        },
        "callerSideCodec": {
          "description": "caller codec for the call",
          "type": "string"
        },
        "correlator": {
          "description": "this is the same for all events on this call",
          "type": "string"
        },
        "endOfCallVqmSummaries": {
          "$ref": "#/definitions/endOfCallVqmSummaries"
        },
        "phoneNumber": {
          "description": "phone number associated with the correlator",
          "type": "string"
        },
        "midCallRtcp": {
          "description": "Base64 encoding of the binary RTCP data excluding Eth/IP/UDP headers",
          "type": "string"
        },
        "vendorNfNameFields": {
          "$ref": "#/definitions/vendorNfNameFields"
        },
        "voiceQualityFieldsVersion": {
          "description": "version of the voiceQualityFields block",
          "type": "string",
          "enum": [
            "4.0"
          ]
        }
      },
      "additionalProperties": false,
      "required": [
        "calleeSideCodec",
        "callerSideCodec",
        "correlator",
        "midCallRtcp",
        "vendorNfNameFields",
        "voiceQualityFieldsVersion"
      ]
    }
  }
}`

var _v3011 *JSONSchema

func init() {
	sch, err := NewSchemaFromBytes([]byte(v3011))
	if err != nil {
		panic(err)
	}
	_v3011 = sch
}

// V3011 loads and returns VES Schema v30.1.1 (VES v7.1)
func V3011() *JSONSchema {
	return _v3011
}

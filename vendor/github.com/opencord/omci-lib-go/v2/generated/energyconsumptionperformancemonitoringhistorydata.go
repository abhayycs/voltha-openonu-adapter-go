/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 * Copyright 2020-present Open Networking Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
 * NOTE: This file was generated, manual edits will be overwritten!
 *
 * Generated by 'goCodeGenerator.py':
 *              https://github.com/cboling/OMCI-parser/README.md
 */

package generated

import "github.com/deckarep/golang-set"

// EnergyConsumptionPerformanceMonitoringHistoryDataClassID is the 16-bit ID for the OMCI
// Managed entity Energy consumption performance monitoring history data
const EnergyConsumptionPerformanceMonitoringHistoryDataClassID = ClassID(343) // 0x0157

var energyconsumptionperformancemonitoringhistorydataBME *ManagedEntityDefinition

// EnergyConsumptionPerformanceMonitoringHistoryData (Class ID: #343 / 0x0157)
//	This ME collects PM data associated with the ONU's energy consumption. The time spent in various
//	low-power states is recorded as a measure of their utility. Furthermore, the ONU may also
//	include the equivalent of a watt-hour meter, which can be sampled from time to time to measure
//	actual power consumed.
//
//	For a complete discussion of generic PM architecture, refer to clause I.4.
//
//	Relationships
//		An instance of this ME is associated with the ONU in its entirety.
//
//	Attributes
//		Managed Entity Id
//			This attribute uniquely identifies each instance of this ME. The ME ID must be 0. (R, set-by-
//			create) (mandatory) (2-bytes)
//
//		Interval End Time
//			This attribute identifies the most recently finished 15-min interval. (R) (mandatory) (1-byte)
//
//		Threshold Data 1_2 Id
//			Threshold data 1/2 ID: No thresholds are defined for this ME. For uniformity with other PMs, the
//			attribute is retained and shown as mandatory, but it should be set to a null pointer. (R,-W,
//			set-by-create) (mandatory) (2-bytes)
//
//		Doze Time
//			This attribute records the time during which the ONU was in doze energy conservation mode,
//			measured in microseconds. If watchful sleep is enabled in the ONU dynamic power management
//			control ME, the ONU ignores this attribute. (R) (mandatory) (4-bytes)
//
//		Cyclic Sleep Time
//			This attribute records the time during which the ONU was in cyclic sleep energy conservation
//			mode, measured in microseconds. If watchful sleep is enabled in the ONU dynamic power management
//			control ME, the ONU ignores this attribute. (R) (mandatory) (4-bytes)
//
//		Watchful Sleep Time
//			This attribute records the time during which the ONU was in watchful sleep energy conservation
//			mode, measured in microseconds. (R) (mandatory) (4-bytes)
//
//		Energy Consumed
//			This attribute records the energy consumed by the ONU, measured in millijoules. (R) (optional)
//			(4-bytes)
//
type EnergyConsumptionPerformanceMonitoringHistoryData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

// Attribute name constants

const EnergyConsumptionPerformanceMonitoringHistoryData_IntervalEndTime = "IntervalEndTime"
const EnergyConsumptionPerformanceMonitoringHistoryData_ThresholdData12Id = "ThresholdData12Id"
const EnergyConsumptionPerformanceMonitoringHistoryData_DozeTime = "DozeTime"
const EnergyConsumptionPerformanceMonitoringHistoryData_CyclicSleepTime = "CyclicSleepTime"
const EnergyConsumptionPerformanceMonitoringHistoryData_WatchfulSleepTime = "WatchfulSleepTime"
const EnergyConsumptionPerformanceMonitoringHistoryData_EnergyConsumed = "EnergyConsumed"

func init() {
	energyconsumptionperformancemonitoringhistorydataBME = &ManagedEntityDefinition{
		Name:    "EnergyConsumptionPerformanceMonitoringHistoryData",
		ClassID: EnergyConsumptionPerformanceMonitoringHistoryDataClassID,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
			GetCurrentData,
		),
		AllowedAttributeMask: 0xfc00,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field(ManagedEntityID, PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, 0),
			1: ByteField(EnergyConsumptionPerformanceMonitoringHistoryData_IntervalEndTime, UnsignedIntegerAttributeType, 0x8000, 0, mapset.NewSetWith(Read), false, false, false, 1),
			2: Uint16Field(EnergyConsumptionPerformanceMonitoringHistoryData_ThresholdData12Id, PointerAttributeType, 0x4000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 2),
			3: Uint32Field(EnergyConsumptionPerformanceMonitoringHistoryData_DozeTime, CounterAttributeType, 0x2000, 0, mapset.NewSetWith(Read), false, false, false, 3),
			4: Uint32Field(EnergyConsumptionPerformanceMonitoringHistoryData_CyclicSleepTime, CounterAttributeType, 0x1000, 0, mapset.NewSetWith(Read), false, false, false, 4),
			5: Uint32Field(EnergyConsumptionPerformanceMonitoringHistoryData_WatchfulSleepTime, CounterAttributeType, 0x0800, 0, mapset.NewSetWith(Read), false, false, false, 5),
			6: Uint32Field(EnergyConsumptionPerformanceMonitoringHistoryData_EnergyConsumed, CounterAttributeType, 0x0400, 0, mapset.NewSetWith(Read), false, true, false, 6),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewEnergyConsumptionPerformanceMonitoringHistoryData (class ID 343) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewEnergyConsumptionPerformanceMonitoringHistoryData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*energyconsumptionperformancemonitoringhistorydataBME, params...)
}

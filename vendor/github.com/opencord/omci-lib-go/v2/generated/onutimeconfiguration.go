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

// OnuTimeConfigurationClassID is the 16-bit ID for the OMCI
// Managed entity ONU time configuration
const OnuTimeConfigurationClassID = ClassID(457) // 0x01c9

var onutimeconfigurationBME *ManagedEntityDefinition

// OnuTimeConfiguration (Class ID: #457 / 0x01c9)
//	This ME provides characterization and manipulation of OLT timestamp information. An ONU that
//	uses OLTbased time synchronization methods automatically creates an instance of this ME. There
//	is no intention that this ME be used to establish a precise time of day reference.
//
//	Relationships
//		The single instance of this ME is associated with the ONU ME.
//
//	Attributes
//		Managed Entity Id
//			This attribute uniquely identifies each instance of this ME. There is only one instance, number
//			0. (R) (mandatory) (2 bytes)
//
//		Current Local Onu Time
//			If the ONU has a real-time clock, it returns the local ONU time. This attribute returns the
//			current ONU time. The local ONU time and synchronize time time-format is the same. (R)
//			(mandatory) (7 bytes)
//
//		Time Qualification Block
//			This attribute describes the time-qualification to be applied to the ONU RTC local time. The
//			following fields are supported:
//
//			(R, W) (mandatory) (2 bytes)
//
type OnuTimeConfiguration struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

// Attribute name constants

const OnuTimeConfiguration_CurrentLocalOnuTime = "CurrentLocalOnuTime"
const OnuTimeConfiguration_TimeQualificationBlock = "TimeQualificationBlock"

func init() {
	onutimeconfigurationBME = &ManagedEntityDefinition{
		Name:    "OnuTimeConfiguration",
		ClassID: OnuTimeConfigurationClassID,
		MessageTypes: mapset.NewSetWith(
			Get,
			Set,
		),
		AllowedAttributeMask: 0xc000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field(ManagedEntityID, PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read), false, false, false, 0),
			1: MultiByteField(OnuTimeConfiguration_CurrentLocalOnuTime, OctetsAttributeType, 0x8000, 7, toOctets("AAAAAAAAAA=="), mapset.NewSetWith(Read), false, false, false, 1),
			2: Uint16Field(OnuTimeConfiguration_TimeQualificationBlock, UnsignedIntegerAttributeType, 0x4000, 0, mapset.NewSetWith(Read, Write), false, false, false, 2),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewOnuTimeConfiguration (class ID 457) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewOnuTimeConfiguration(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*onutimeconfigurationBME, params...)
}

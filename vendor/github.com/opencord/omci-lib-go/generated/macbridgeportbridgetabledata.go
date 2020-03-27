/*
 * Copyright (c) 2018 - present.  Boling Consulting Solutions (bcsw.net)
 * Copyright 2020-present Open Networking Foundation

 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at

 * http://www.apache.org/licenses/LICENSE-2.0

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

// MacBridgePortBridgeTableDataClassID is the 16-bit ID for the OMCI
// Managed entity MAC bridge port bridge table data
const MacBridgePortBridgeTableDataClassID ClassID = ClassID(50)

var macbridgeportbridgetabledataBME *ManagedEntityDefinition

// MacBridgePortBridgeTableData (class ID #50)
//	This ME reports status data associated with a bridge port. The ONU automatically creates or
//	deletes an instance of this ME upon the creation or deletion of a MAC bridge port configuration
//	data.
//
//	Relationships
//		An instance of this ME is associated with an instance of a MAC bridge port configuration data
//		ME.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the MAC bridge port configuration
//			data ME. (R) (mandatory) (2-bytes)
//
//		Bridge Table
//			Upon ME instantiation, this attribute is an empty list. (R) (mandatory) (8-*-M-bytes, where M is
//			the number of entries in the list.)
//
type MacBridgePortBridgeTableData struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	macbridgeportbridgetabledataBME = &ManagedEntityDefinition{
		Name:    "MacBridgePortBridgeTableData",
		ClassID: 50,
		MessageTypes: mapset.NewSetWith(
			Get,
			GetNext,
		),
		AllowedAttributeMask: 0x8000,
		AttributeDefinitions: AttributeDefinitionMap{
			0: Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read), false, false, false, 0),
			1: TableField("BridgeTable", TableAttributeType, 0x8000, TableInfo{nil, 8}, mapset.NewSetWith(Read), false, false, false, 1),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewMacBridgePortBridgeTableData (class ID 50) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewMacBridgePortBridgeTableData(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*macbridgeportbridgetabledataBME, params...)
}

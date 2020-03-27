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

// XdslLineInventoryAndStatusDataPart2ClassID is the 16-bit ID for the OMCI
// Managed entity xDSL line inventory and status data part 2
const XdslLineInventoryAndStatusDataPart2ClassID ClassID = ClassID(101)

var xdsllineinventoryandstatusdatapart2BME *ManagedEntityDefinition

// XdslLineInventoryAndStatusDataPart2 (class ID #101)
//	This ME contains part 2 of the line inventory and status data for an xDSL UNI. The ONU
//	automatically creates or deletes an instance of this ME upon the creation or deletion of a PPTP
//	xDSL UNI part 1.
//
//	NOTE 1 - [ITU-T G.997.1] specifies that bit rate attributes have granularity of 1000-bit/s. If
//	ITUT-G.997.1 compliance is required, the ONU should only report values with this granularity.
//
//	Relationships
//		An instance of this ME is associated with an xDSL UNI.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. Through an
//			identical ID, this ME is implicitly linked to an instance of the PPTP xDSL UNI part 1. (R)
//			(mandatory) (2-bytes)
//
//		Xdsl Transmission System
//			NOTE 2 - This attribute is only 7-bytes long. An eighth byte identifying VDSL2 capabilities in
//			use is defined in the VDSL2 line inventory and status data part 1 ME.
//
//		Line Power Management State
//			(R) (mandatory) (1-byte)
//
//		Downstream Line Attenuation
//			NOTE 3 - [ITU-T G.993.2] specifies a per-band array to represent this attribute. The array is
//			defined in the VDSL2 line inventory and status data part 3 ME. In an ITU-T G.993.2 context, the
//			downstream line attenuation attribute should be set to 0 here, and populated in the VDSL2 line
//			inventory and status data part 3 ME instead.
//
//		Upstream Line Attenuation
//			NOTE 4 - [ITU-T G.993.2] specifies a per-band array to represent this attribute. The array is
//			defined in the VDSL2 line inventory and status data part 3 ME. In an ITU-T G.993.2 context, the
//			upstream line attenuation attribute should be set to 0 here, and populated in the VDSL2 line
//			inventory and status data part 3 ME instead.
//
//		Downstream Signal Attenuation
//			NOTE 6 - [ITU-T G.993.2] specifies a per-band array to represent this attribute. The array is
//			defined in the VDSL2 line inventory and status data part 3 ME. In an ITU-T G.993.2 context, the
//			downstream signal attenuation attribute should be set to 0 here, and populated in the VDSL2 line
//			inventory and status data part 3 ME instead.
//
//		Upstream Signal Attenuation
//			NOTE 8 - [ITU-T G.993.2] specifies a per-band array to represent this attribute. The array is
//			defined in the VDSL2 line inventory and status data part 3 ME. In an ITU-T G.993.2 context, the
//			upstream signal attenuation attribute should be set to 0 here, and populated in the VDSL2 line
//			inventory and status data part 3 ME instead.
//
//		Downstream Snr Ratio Margin
//			Downstream SNR ratio margin: The downstream SNR margin SNRMds is the maximum increase of noise
//			power received at the xTUR, such that the BER requirements can still be met for all downstream
//			bearer channels. The attribute value ranges from 0 (-64.0 dB) to 1270 (+63.0-dB). The special
//			value 0xFFFF indicates that the attribute is out of range (R) (mandatory) (2-bytes)
//
//		Upstream Snr Margin
//			Upstream SNR margin: The upstream SNR margin SNRMus is the maximum increase of noise power
//			received at the xTUC, such that the BER requirements can still be met for all upstream bearer
//			channels. The attribute value ranges from 0 (-64.0 dB) to 1270 (+63.0 dB). The special value
//			0xFFFF indicates that the attribute is out of range. (R) (mandatory) (2-bytes)
//
//		Downstream Maximum Attainable Data Rate
//			Downstream maximum attainable data rate: The ATTNDRds attribute indicates the maximum downstream
//			net data rate currently attainable. The rate is coded in bits per second. (R) (mandatory)
//			(4-bytes)
//
//		Upstream Maximum Attainable Data Rate
//			Upstream maximum attainable data rate: The ATTNDRus attribute indicates the maximum upstream net
//			data rate currently attainable. The rate is coded in bits per second. (R) (mandatory) (4-bytes)
//
//		Downstream Actual Power Spectrum Density
//			Downstream actual power spectrum density: The ACTPSDds attribute is the average downstream
//			transmit power spectrum density over the subcarriers in use (subcarriers to which downstream
//			user data are allocated) delivered by the xTUC at the UC reference point, at the instant of
//			measurement. The attribute value ranges from 0 (-90.0-dBm/Hz) to 900 (0.0-dBm/Hz). The special
//			value (0xFFFF) indicates that the parameter is out of range. (R) (mandatory) (2-bytes)
//
//		Upstream Actual Power Spectrum Density
//			Upstream actual power spectrum density: The ACTPSDus attribute is the average upstream transmit
//			power spectrum density over the subcarriers in use (subcarriers to which upstream user data are
//			allocated) delivered by the xTUR at the UR reference point, at the instant of measurement. The
//			attribute value ranges from 0 (-90.0-dBm/Hz) to 900 (0.0-dBm/Hz). The special value 0xFFFF
//			indicates that the attribute is out of range. (R) (mandatory) (2-bytes)
//
//		Downstream Actual Aggregate Transmit Power
//			NOTE 9 - The downstream nominal aggregate transmit power may be taken as a best estimate of the
//			parameter.
//
//		Upstream Actual Aggregate Transmit Power
//			NOTE 10 - The upstream nominal aggregate transmit power may be taken as a best estimate of the
//			parameter.
//
//		Initialization _ Last State Transmitted Downstream
//			(R) (mandatory) (1-byte)
//
//		Initialization _ Last State Transmitted Upstream
//			(R) (mandatory) (1-byte)
//
type XdslLineInventoryAndStatusDataPart2 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdsllineinventoryandstatusdatapart2BME = &ManagedEntityDefinition{
		Name:    "XdslLineInventoryAndStatusDataPart2",
		ClassID: 101,
		MessageTypes: mapset.NewSetWith(
			Get,
		),
		AllowedAttributeMask: 0xffff,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read), false, false, false, 0),
			1:  MultiByteField("XdslTransmissionSystem", OctetsAttributeType, 0x8000, 7, toOctets("AAAAAAAAAA=="), mapset.NewSetWith(Read), false, false, false, 1),
			2:  ByteField("LinePowerManagementState", UnsignedIntegerAttributeType, 0x4000, 0, mapset.NewSetWith(Read), false, false, false, 2),
			3:  Uint16Field("DownstreamLineAttenuation", UnsignedIntegerAttributeType, 0x2000, 0, mapset.NewSetWith(Read), false, false, false, 3),
			4:  Uint16Field("UpstreamLineAttenuation", UnsignedIntegerAttributeType, 0x1000, 0, mapset.NewSetWith(Read), false, false, false, 4),
			5:  Uint16Field("DownstreamSignalAttenuation", UnsignedIntegerAttributeType, 0x0800, 0, mapset.NewSetWith(Read), false, false, false, 5),
			6:  Uint16Field("UpstreamSignalAttenuation", UnsignedIntegerAttributeType, 0x0400, 0, mapset.NewSetWith(Read), false, false, false, 6),
			7:  Uint16Field("DownstreamSnrRatioMargin", UnsignedIntegerAttributeType, 0x0200, 0, mapset.NewSetWith(Read), false, false, false, 7),
			8:  Uint16Field("UpstreamSnrMargin", UnsignedIntegerAttributeType, 0x0100, 0, mapset.NewSetWith(Read), false, false, false, 8),
			9:  Uint32Field("DownstreamMaximumAttainableDataRate", UnsignedIntegerAttributeType, 0x0080, 0, mapset.NewSetWith(Read), false, false, false, 9),
			10: Uint32Field("UpstreamMaximumAttainableDataRate", UnsignedIntegerAttributeType, 0x0040, 0, mapset.NewSetWith(Read), false, false, false, 10),
			11: Uint16Field("DownstreamActualPowerSpectrumDensity", UnsignedIntegerAttributeType, 0x0020, 0, mapset.NewSetWith(Read), false, false, false, 11),
			12: Uint16Field("UpstreamActualPowerSpectrumDensity", UnsignedIntegerAttributeType, 0x0010, 0, mapset.NewSetWith(Read), false, false, false, 12),
			13: Uint16Field("DownstreamActualAggregateTransmitPower", UnsignedIntegerAttributeType, 0x0008, 0, mapset.NewSetWith(Read), false, false, false, 13),
			14: Uint16Field("UpstreamActualAggregateTransmitPower", UnsignedIntegerAttributeType, 0x0004, 0, mapset.NewSetWith(Read), false, false, false, 14),
			15: ByteField("InitializationLastStateTransmittedDownstream", UnsignedIntegerAttributeType, 0x0002, 0, mapset.NewSetWith(Read), false, false, false, 15),
			16: ByteField("InitializationLastStateTransmittedUpstream", UnsignedIntegerAttributeType, 0x0001, 0, mapset.NewSetWith(Read), false, false, false, 16),
		},
		Access:  CreatedByOnu,
		Support: UnknownSupport,
	}
}

// NewXdslLineInventoryAndStatusDataPart2 (class ID 101) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewXdslLineInventoryAndStatusDataPart2(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xdsllineinventoryandstatusdatapart2BME, params...)
}

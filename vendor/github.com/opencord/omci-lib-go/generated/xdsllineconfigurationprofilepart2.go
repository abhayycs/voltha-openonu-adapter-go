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

// XdslLineConfigurationProfilePart2ClassID is the 16-bit ID for the OMCI
// Managed entity xDSL line configuration profile part 2
const XdslLineConfigurationProfilePart2ClassID ClassID = ClassID(105)

var xdsllineconfigurationprofilepart2BME *ManagedEntityDefinition

// XdslLineConfigurationProfilePart2 (class ID #105)
//	The overall xDSL line configuration profile is modelled in several parts, all of which are
//	associated together through a common ME ID (the client PPTP xDSL UNI part 1 has a single
//	pointer, which refers to the entire set of line configuration profile parts).
//
//	Relationships
//		An instance of this ME may be associated with zero or more instances of an xDSL UNI.
//
//	Attributes
//		Managed Entity Id
//			Managed entity ID: This attribute uniquely identifies each instance of this ME. All xDSL and
//			VDSL2 line configuration profiles and extensions that pertain to a given PPTP xDSL UNI must
//			share a common ME ID. (R, setbycreate) (mandatory) (2-bytes)
//
//		Downstream Minimum Time Interval For Upshift Rate Adaptation
//			Downstream minimum time interval for upshift rate adaptation: This parameter defines the
//			interval during which the downstream noise margin must remain above the downstream upshift noise
//			margin before the xTU-R attempts to increase the downstream net data rate. Its value ranges from
//			0 to 16383-s. (R,-W, setbycreate) (optional) (2-bytes)
//
//		Upstream Minimum Time Interval For Upshift Rate Adaptation
//			Upstream minimum time interval for upshift rate adaptation: This parameter defines the interval
//			during which the upstream noise margin must remain above the upstream upshift noise margin
//			before the xTU-C attempts to increase the upstream net data rate. Its value ranges from 0 to
//			16383-s. (R,-W, setbycreate) (optional) (2-bytes)
//
//		Downstream Downshift Noise Margin
//			Downstream downshift noise margin: If the downstream noise margin is below the downstream
//			downshift noise margin and remains there for more than the downstream minimum time interval for
//			downshift rate adaptation, the xTU-R attempts to decrease the downstream net data rate. This
//			attribute's value ranges from 0 (0.0 dB) to 310 (31.0 dB). (R,-W, setbycreate) (optional)
//			(2-bytes)
//
//		Upstream Downshift Noise Margin
//			Upstream downshift noise margin: If the upstream noise margin is below the upstream downshift
//			noise margin and remains there for more than the upstream minimum time interval for downshift
//			rate adaptation, the xTUC attempts to decrease the upstream net data rate. This attribute's
//			value ranges from 0 (0.0 dB) to 310 (31.0 dB). (R,-W, setbycreate) (optional) (2-bytes)
//
//		Downstream Minimum Time Interval For Downshift Rate Adaptation
//			Downstream minimum time interval for downshift rate adaptation: This parameter defines the
//			interval during which the downstream noise margin must remain below the downstream downshift
//			noise margin before the xTU-R attempts to decrease the downstream net data rate. Its value
//			ranges from 0 to 16383-s. (R,-W, setbycreate) (optional) (2-bytes)
//
//		Upstream Minimum Time Interval For Downshift Rate Adaptation
//			Upstream minimum time interval for downshift rate adaptation: This parameter defines the
//			interval during which the upstream noise margin must remain below the upstream downshift noise
//			margin before the xTU-C attempts to decrease the upstream net data rate. Its value ranges from 0
//			to 16383-s. (R,-W, setbycreate) (optional) (2-bytes)
//
//		Xtu Impedance State Forced
//			(R,-W, setbycreate) (optional) (1-byte)
//
//		L0_Time
//			L0-time:	This parameter specifies the minimum time between an exit from the L2 state and the
//			next entry into the L2 state. It is only valid for [ITUT-G.992.3], [ITUT-G.992.4] and
//			[ITUT-G.992.5]. It ranges from 0 to 255-s. (R,-W, setbycreate) (mandatory) (1 byte)
//
//		L2_Time
//			L2-time:	This parameter specifies the minimum time between an entry into the L2 state and the
//			first power trim in the L2 state, or between two consecutive power trims in the L2 state. It is
//			only valid for [ITUT-G.992.3], [ITUT-G.992.4] and [ITUT-G.992.5]. It ranges from 0 to 255-s. (R,
//			W, setbycreate) (mandatory) (1 byte)
//
//		Downstream Maximum Nominal Power Spectral Density
//			Downstream maximum nominal power spectral density: This attribute specifies the maximum nominal
//			transmit PSD in the downstream direction during initialization and showtime. A single
//			MAXNOMPSDds attribute is defined per mode enabled in the xTSE line configuration attribute. It
//			is only valid for [ITUT-G.992.3], [ITUT-G.992.4] and [ITUT-G.992.5]. Its value ranges from 0
//			(60.0-dBm/Hz) to 300 (-30-dBm/Hz). (R, W, setbycreate) (mandatory) (2 bytes)
//
//		Upstream Maximum Nominal Power Spectral Density
//			Upstream maximum nominal power spectral density: This attribute specifies the maximum nominal
//			transmit PSD in the upstream direction during initialization and showtime. A single MAXNOMPSDus
//			attribute is defined per mode enabled in the xTSE line configuration attribute. It is only valid
//			for [ITUT-G.992.3], [ITUT-G.992.4] and [ITUT-G.993.2]. Its value ranges from 0 (-60.0-dBm/Hz) to
//			300 (-30-dBm/Hz). (R, W, setbycreate) (mandatory) (2 bytes)
//
//		Downstream Maximum Nominal Aggregate Transmit Power
//			Downstream maximum nominal aggregate transmit power: This attribute specifies the maximum
//			nominal aggregate transmit power in the downstream direction during initialization and showtime.
//			It is only valid for [ITUT-G.992.3], [ITUT-G.992.4], [ITUT-G.992.5] and [ITUT-G.993.2]. Its
//			value ranges from 0 (0.0-dBm) to 255 (25.5-dBm). (R, W, setbycreate) (mandatory) (1-byte)
//
//		Upstream Maximum Nominal Aggregate Transmit Power
//			Upstream maximum nominal aggregate transmit power: This parameter specifies the maximum nominal
//			aggregate transmit power in the upstream direction during initialization and showtime. It is
//			only valid for [ITUT-G.992.3], [ITUT-G.992.4] and [ITUT-G.992.5]. Its value ranges from 0
//			(0.0-dBm) to 255 (25.5-dBm). (R, W, setbycreate) (mandatory) (1 byte)
//
//		Upstream Maximum Aggregate_Receive Power
//			Upstream maximum aggregate-receive power: This parameter specifies the maximum upstream
//			aggregate receive power over a set of subcarriers, as defined in the relevant Recommendation.
//			The xTU-C requests an upstream power cutback such that the upstream aggregate receive power over
//			that set of subcarriers is at or below the configured maximum value. It is only valid for
//			[ITUT-G.992.3], [ITUT-G.992.4] and [ITUT-G.992.5]. This attribute ranges from 0 (25.5-dBm) to
//			510 (+25.5-dBm). The special value 0xFFFF indicates that no upstream maximum aggregate receive
//			power limit is to be applied. (R, W setbycreate) (mandatory) (2 bytes)
//
//		Vdsl2 Transmission System Enabling
//			VDSL2 transmission system enabling: This configuration attribute extends the transmission system
//			coding types to be allowed by the xTU-C. It is a bit map, defined as octet 8 (bits 57..64) in
//			Table-9.7.12-1. (R, W, setbycreate) (optional) (1 byte)
//
type XdslLineConfigurationProfilePart2 struct {
	ManagedEntityDefinition
	Attributes AttributeValueMap
}

func init() {
	xdsllineconfigurationprofilepart2BME = &ManagedEntityDefinition{
		Name:    "XdslLineConfigurationProfilePart2",
		ClassID: 105,
		MessageTypes: mapset.NewSetWith(
			Create,
			Delete,
			Get,
			Set,
		),
		AllowedAttributeMask: 0xfffe,
		AttributeDefinitions: AttributeDefinitionMap{
			0:  Uint16Field("ManagedEntityId", PointerAttributeType, 0x0000, 0, mapset.NewSetWith(Read, SetByCreate), false, false, false, 0),
			1:  Uint16Field("DownstreamMinimumTimeIntervalForUpshiftRateAdaptation", UnsignedIntegerAttributeType, 0x8000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 1),
			2:  Uint16Field("UpstreamMinimumTimeIntervalForUpshiftRateAdaptation", UnsignedIntegerAttributeType, 0x4000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 2),
			3:  Uint16Field("DownstreamDownshiftNoiseMargin", UnsignedIntegerAttributeType, 0x2000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 3),
			4:  Uint16Field("UpstreamDownshiftNoiseMargin", UnsignedIntegerAttributeType, 0x1000, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 4),
			5:  Uint16Field("DownstreamMinimumTimeIntervalForDownshiftRateAdaptation", UnsignedIntegerAttributeType, 0x0800, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 5),
			6:  Uint16Field("UpstreamMinimumTimeIntervalForDownshiftRateAdaptation", UnsignedIntegerAttributeType, 0x0400, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 6),
			7:  ByteField("XtuImpedanceStateForced", UnsignedIntegerAttributeType, 0x0200, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 7),
			8:  ByteField("L0Time", UnsignedIntegerAttributeType, 0x0100, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 8),
			9:  ByteField("L2Time", UnsignedIntegerAttributeType, 0x0080, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 9),
			10: Uint16Field("DownstreamMaximumNominalPowerSpectralDensity", UnsignedIntegerAttributeType, 0x0040, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 10),
			11: Uint16Field("UpstreamMaximumNominalPowerSpectralDensity", UnsignedIntegerAttributeType, 0x0020, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 11),
			12: ByteField("DownstreamMaximumNominalAggregateTransmitPower", UnsignedIntegerAttributeType, 0x0010, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 12),
			13: ByteField("UpstreamMaximumNominalAggregateTransmitPower", UnsignedIntegerAttributeType, 0x0008, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, false, false, 13),
			14: Uint16Field("UpstreamMaximumAggregateReceivePower", UnsignedIntegerAttributeType, 0x0004, 0, mapset.NewSetWith(Read), false, false, false, 14),
			15: ByteField("Vdsl2TransmissionSystemEnabling", UnsignedIntegerAttributeType, 0x0002, 0, mapset.NewSetWith(Read, SetByCreate, Write), false, true, false, 15),
		},
		Access:  CreatedByOlt,
		Support: UnknownSupport,
	}
}

// NewXdslLineConfigurationProfilePart2 (class ID 105) creates the basic
// Managed Entity definition that is used to validate an ME of this type that
// is received from or transmitted to the OMCC.
func NewXdslLineConfigurationProfilePart2(params ...ParamData) (*ManagedEntity, OmciErrors) {
	return NewManagedEntity(*xdsllineconfigurationprofilepart2BME, params...)
}

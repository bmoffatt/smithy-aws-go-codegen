// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type Algorithm string

// Enum values for Algorithm
const (
	AlgorithmAes128 Algorithm = "aes128"
	AlgorithmAes192 Algorithm = "aes192"
	AlgorithmAes256 Algorithm = "aes256"
)

// Values returns all known values for Algorithm. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Algorithm) Values() []Algorithm {
	return []Algorithm{
		"aes128",
		"aes192",
		"aes256",
	}
}

type Colorimetry string

// Enum values for Colorimetry
const (
	ColorimetryBt601   Colorimetry = "BT601"
	ColorimetryBt709   Colorimetry = "BT709"
	ColorimetryBt2020  Colorimetry = "BT2020"
	ColorimetryBt2100  Colorimetry = "BT2100"
	ColorimetrySt20651 Colorimetry = "ST2065-1"
	ColorimetrySt20653 Colorimetry = "ST2065-3"
	ColorimetryXyz     Colorimetry = "XYZ"
)

// Values returns all known values for Colorimetry. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (Colorimetry) Values() []Colorimetry {
	return []Colorimetry{
		"BT601",
		"BT709",
		"BT2020",
		"BT2100",
		"ST2065-1",
		"ST2065-3",
		"XYZ",
	}
}

type DurationUnits string

// Enum values for DurationUnits
const (
	DurationUnitsMonths DurationUnits = "MONTHS"
)

// Values returns all known values for DurationUnits. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DurationUnits) Values() []DurationUnits {
	return []DurationUnits{
		"MONTHS",
	}
}

type EncoderProfile string

// Enum values for EncoderProfile
const (
	EncoderProfileMain EncoderProfile = "main"
	EncoderProfileHigh EncoderProfile = "high"
)

// Values returns all known values for EncoderProfile. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EncoderProfile) Values() []EncoderProfile {
	return []EncoderProfile{
		"main",
		"high",
	}
}

type EncodingName string

// Enum values for EncodingName
const (
	EncodingNameJxsv     EncodingName = "jxsv"
	EncodingNameRaw      EncodingName = "raw"
	EncodingNameSmpte291 EncodingName = "smpte291"
	EncodingNamePcm      EncodingName = "pcm"
)

// Values returns all known values for EncodingName. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (EncodingName) Values() []EncodingName {
	return []EncodingName{
		"jxsv",
		"raw",
		"smpte291",
		"pcm",
	}
}

type EntitlementStatus string

// Enum values for EntitlementStatus
const (
	EntitlementStatusEnabled  EntitlementStatus = "ENABLED"
	EntitlementStatusDisabled EntitlementStatus = "DISABLED"
)

// Values returns all known values for EntitlementStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EntitlementStatus) Values() []EntitlementStatus {
	return []EntitlementStatus{
		"ENABLED",
		"DISABLED",
	}
}

type FailoverMode string

// Enum values for FailoverMode
const (
	FailoverModeMerge    FailoverMode = "MERGE"
	FailoverModeFailover FailoverMode = "FAILOVER"
)

// Values returns all known values for FailoverMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (FailoverMode) Values() []FailoverMode {
	return []FailoverMode{
		"MERGE",
		"FAILOVER",
	}
}

type KeyType string

// Enum values for KeyType
const (
	KeyTypeSpeke       KeyType = "speke"
	KeyTypeStaticKey   KeyType = "static-key"
	KeyTypeSrtPassword KeyType = "srt-password"
)

// Values returns all known values for KeyType. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (KeyType) Values() []KeyType {
	return []KeyType{
		"speke",
		"static-key",
		"srt-password",
	}
}

type MaintenanceDay string

// Enum values for MaintenanceDay
const (
	MaintenanceDayMonday    MaintenanceDay = "Monday"
	MaintenanceDayTuesday   MaintenanceDay = "Tuesday"
	MaintenanceDayWednesday MaintenanceDay = "Wednesday"
	MaintenanceDayThursday  MaintenanceDay = "Thursday"
	MaintenanceDayFriday    MaintenanceDay = "Friday"
	MaintenanceDaySaturday  MaintenanceDay = "Saturday"
	MaintenanceDaySunday    MaintenanceDay = "Sunday"
)

// Values returns all known values for MaintenanceDay. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MaintenanceDay) Values() []MaintenanceDay {
	return []MaintenanceDay{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
}

type MediaStreamType string

// Enum values for MediaStreamType
const (
	MediaStreamTypeVideo         MediaStreamType = "video"
	MediaStreamTypeAudio         MediaStreamType = "audio"
	MediaStreamTypeAncillaryData MediaStreamType = "ancillary-data"
)

// Values returns all known values for MediaStreamType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (MediaStreamType) Values() []MediaStreamType {
	return []MediaStreamType{
		"video",
		"audio",
		"ancillary-data",
	}
}

type NetworkInterfaceType string

// Enum values for NetworkInterfaceType
const (
	NetworkInterfaceTypeEna NetworkInterfaceType = "ena"
	NetworkInterfaceTypeEfa NetworkInterfaceType = "efa"
)

// Values returns all known values for NetworkInterfaceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (NetworkInterfaceType) Values() []NetworkInterfaceType {
	return []NetworkInterfaceType{
		"ena",
		"efa",
	}
}

type PriceUnits string

// Enum values for PriceUnits
const (
	PriceUnitsHourly PriceUnits = "HOURLY"
)

// Values returns all known values for PriceUnits. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (PriceUnits) Values() []PriceUnits {
	return []PriceUnits{
		"HOURLY",
	}
}

type Protocol string

// Enum values for Protocol
const (
	ProtocolZixiPush     Protocol = "zixi-push"
	ProtocolRtpFec       Protocol = "rtp-fec"
	ProtocolRtp          Protocol = "rtp"
	ProtocolZixiPull     Protocol = "zixi-pull"
	ProtocolRist         Protocol = "rist"
	ProtocolSt2110Jpegxs Protocol = "st2110-jpegxs"
	ProtocolCdi          Protocol = "cdi"
	ProtocolSrtListener  Protocol = "srt-listener"
	ProtocolFujitsuQos   Protocol = "fujitsu-qos"
)

// Values returns all known values for Protocol. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Protocol) Values() []Protocol {
	return []Protocol{
		"zixi-push",
		"rtp-fec",
		"rtp",
		"zixi-pull",
		"rist",
		"st2110-jpegxs",
		"cdi",
		"srt-listener",
		"fujitsu-qos",
	}
}

type Range string

// Enum values for Range
const (
	RangeNarrow      Range = "NARROW"
	RangeFull        Range = "FULL"
	RangeFullprotect Range = "FULLPROTECT"
)

// Values returns all known values for Range. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Range) Values() []Range {
	return []Range{
		"NARROW",
		"FULL",
		"FULLPROTECT",
	}
}

type ReservationState string

// Enum values for ReservationState
const (
	ReservationStateActive     ReservationState = "ACTIVE"
	ReservationStateExpired    ReservationState = "EXPIRED"
	ReservationStateProcessing ReservationState = "PROCESSING"
	ReservationStateCanceled   ReservationState = "CANCELED"
)

// Values returns all known values for ReservationState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ReservationState) Values() []ReservationState {
	return []ReservationState{
		"ACTIVE",
		"EXPIRED",
		"PROCESSING",
		"CANCELED",
	}
}

type ResourceType string

// Enum values for ResourceType
const (
	ResourceTypeMbpsOutboundBandwidth ResourceType = "Mbps_Outbound_Bandwidth"
)

// Values returns all known values for ResourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ResourceType) Values() []ResourceType {
	return []ResourceType{
		"Mbps_Outbound_Bandwidth",
	}
}

type ScanMode string

// Enum values for ScanMode
const (
	ScanModeProgressive               ScanMode = "progressive"
	ScanModeInterlace                 ScanMode = "interlace"
	ScanModeProgressiveSegmentedFrame ScanMode = "progressive-segmented-frame"
)

// Values returns all known values for ScanMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (ScanMode) Values() []ScanMode {
	return []ScanMode{
		"progressive",
		"interlace",
		"progressive-segmented-frame",
	}
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeOwned    SourceType = "OWNED"
	SourceTypeEntitled SourceType = "ENTITLED"
)

// Values returns all known values for SourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SourceType) Values() []SourceType {
	return []SourceType{
		"OWNED",
		"ENTITLED",
	}
}

type State string

// Enum values for State
const (
	StateEnabled  State = "ENABLED"
	StateDisabled State = "DISABLED"
)

// Values returns all known values for State. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (State) Values() []State {
	return []State{
		"ENABLED",
		"DISABLED",
	}
}

type Status string

// Enum values for Status
const (
	StatusStandby  Status = "STANDBY"
	StatusActive   Status = "ACTIVE"
	StatusUpdating Status = "UPDATING"
	StatusDeleting Status = "DELETING"
	StatusStarting Status = "STARTING"
	StatusStopping Status = "STOPPING"
	StatusError    Status = "ERROR"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"STANDBY",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"STARTING",
		"STOPPING",
		"ERROR",
	}
}

type Tcs string

// Enum values for Tcs
const (
	TcsSdr          Tcs = "SDR"
	TcsPq           Tcs = "PQ"
	TcsHlg          Tcs = "HLG"
	TcsLinear       Tcs = "LINEAR"
	TcsBt2100linpq  Tcs = "BT2100LINPQ"
	TcsBt2100linhlg Tcs = "BT2100LINHLG"
	TcsSt20651      Tcs = "ST2065-1"
	TcsSt4281       Tcs = "ST428-1"
	TcsDensity      Tcs = "DENSITY"
)

// Values returns all known values for Tcs. Note that this can be expanded in the
// future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Tcs) Values() []Tcs {
	return []Tcs{
		"SDR",
		"PQ",
		"HLG",
		"LINEAR",
		"BT2100LINPQ",
		"BT2100LINHLG",
		"ST2065-1",
		"ST428-1",
		"DENSITY",
	}
}

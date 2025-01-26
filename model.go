package wsjtxudp

import (
	"time"
)

type WSJTXMsgType uint32

func (mt WSJTXMsgType) String() string {
	switch mt {
	case Heartbeat:
		return "Heartbeat"
	case Status:
		return "Status"
	case Decode:
		return "Decode"
	case Clear:
		return "Clear"
	case Reply:
		return "Reply"
	case QSOLogged:
		return "QSOLogged"
	case Close:
		return "Close"
	case Replay:
		return "Replay"
	case HaltTx:
		return "HaltTx"
	case FreeText:
		return "FreeText"
	case WSPRDecode:
		return "WSPRDecode"
	case Location:
		return "Location"
	case LoggedADIF:
		return "LoggedADIF"
	case HighlightCallsign:
		return "HighlightCallsign"
	case SwitchConfiguration:
		return "SwitchConfiguration"
	case Configure:
		return "Configure"
	case AnnotationInfo:
		return "AnnotationInfo"
	default:
		return "Unknown"
	}
}

const (
	Heartbeat WSJTXMsgType = iota
	Status
	Decode
	Clear
	Reply
	QSOLogged
	Close
	Replay
	HaltTx
	FreeText
	WSPRDecode
	Location
	LoggedADIF
	HighlightCallsign
	SwitchConfiguration
	Configure
	AnnotationInfo
)

type WSJTXMessage struct {
	Header  WSJTXMessageHeader `json:"header"`
	Payload interface{}        `json:"payload"`
}

type WSJTXMessageHeader struct {
	Magic   uint32       `json:"magic"`
	Schema  uint32       `json:"schema"`
	MsgType WSJTXMsgType `json:"msgType"`
}

type HeartbeatPayload struct {
	ID              string `json:"id"`
	MaxSchemaNumber uint32 `json:"maxSchemaNumber"`
	Version         string `json:"version"`
	Revision        string `json:"revision"`
}

type StatusPayload struct {
	ID                   string `json:"id"`
	DialFrequency        uint64 `json:"dialFrequency"`
	Mode                 string `json:"mode"`
	DxCall               string `json:"dxCall"`
	Report               string `json:"report"`
	TxMode               string `json:"txMode"`
	TxEnabled            bool   `json:"txEnabled"`
	Transmitting         bool   `json:"transmitting"`
	Decoding             bool   `json:"decoding"`
	RxDF                 uint32 `json:"rxDF"`
	TxDF                 uint32 `json:"txDF"`
	DECall               string `json:"deCall"`
	DEGrid               string `json:"deGrid"`
	DXGrid               string `json:"dxGrid"`
	TxWatchdog           bool   `json:"txWatchdog"`
	SubMode              string `json:"subMode"`
	FastMode             bool   `json:"fastMode"`
	SpecialOperationMode uint8  `json:"specialOperationMode"`
	FrequencyTolerance   uint32 `json:"frequencyTolerance"`
	TRPeriod             uint32 `json:"trPeriod"`
	ConfigurationName    string `json:"configurationName"`
	TxMessage            string `json:"txMessage"`
}

type DecodePayload struct {
	ID             string  `json:"id"`
	New            bool    `json:"new"`
	Time           uint32  `json:"time"`
	SNR            int32   `json:"snr"`
	DeltaTime      float64 `json:"deltaTime"`
	DeltaFrequency uint32  `json:"deltaFrequency"`
	Mode           string  `json:"mode"`
	Message        string  `json:"message"`
	LowConfidence  bool    `json:"lowConfidence"`
	OffAir         bool    `json:"offAir"`
}

type ClearPayload struct {
	ID     string `json:"id"`
	Window uint8  `json:"window"`
}

type QSOLoggedPayload struct {
	ID                  string    `json:"id"`
	DateTimeOff         time.Time `json:"dateTimeOff"`
	DXCall              string    `json:"dxCall"`
	DXGrid              string    `json:"dxGrid"`
	TxFrequency         uint64    `json:"txFrequency"`
	Mode                string    `json:"mode"`
	ReportSent          string    `json:"reportSent"`
	ReportReceived      string    `json:"reportReceived"`
	TxPower             string    `json:"txPower"`
	Comments            string    `json:"comments"`
	Name                string    `json:"name"`
	DateTimeOn          time.Time `json:"dateTimeOn"`
	OperatorCall        string    `json:"operatorCall"`
	MyCall              string    `json:"myCall"`
	MyGrid              string    `json:"myGrid"`
	ExchangeSent        string    `json:"exchangeSent"`
	ExchangeReceived    string    `json:"exchangeReceived"`
	ADIFPropagationMode string    `json:"ADIFPropagationMode"`
}

type ClosePayload struct {
	ID string `json:"id"`
}

type WSPRDecodePayload struct {
	ID        string  `json:"id"`
	New       bool    `json:"new"`
	Time      uint32  `json:"time"`
	SNR       int32   `json:"snr"`
	DeltaTime float64 `json:"deltaTime"`
	Frequency uint64  `json:"frequency"`
	Drift     int32   `json:"drift"`
	Callsign  string  `json:"callsign"`
	Grid      string  `json:"grid"`
	Power     int32   `json:"power"`
	OffAir    bool    `json:"offAir"`
}

type LoggedADIFPayload struct {
	ID   string `json:"id"`
	ADIF string `json:"adif"`
}

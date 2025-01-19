package wsjtxudp

import (
	"fmt"
)

const (
	WSJTX_MAGIC_NUMBER uint32 = 0xadbccbda
)

type WSJTXParser struct {
	raw     []byte
	cursor  int
	message WSJTXMessage
}

func (p *WSJTXParser) Parse(raw []byte) (WSJTXMessage, error) {
	p.raw = raw
	p.cursor = 0
	p.message = WSJTXMessage{}

	err := p.parseMessageHeader()
	if err != nil {
		return p.message, err
	}

	p.parseMessagePayload()

	return p.message, nil
}

func (p *WSJTXParser) parseMessageHeader() error {
	header := &p.message.Header

	header.Magic, _ = p.parseQuint32()
	if header.Magic != WSJTX_MAGIC_NUMBER {
		return fmt.Errorf("Invalid magic number. Not a WSJT-X message.")
	}
	header.Schema, _ = p.parseQuint32()

	msgType, _ := p.parseQuint32()
	header.MsgType = WSJTXMsgType(msgType)

	return nil
}

// Assumes the header was already parsed.
func (p *WSJTXParser) parseMessagePayload() {

	switch p.message.Header.MsgType {
	case Heartbeat:
		p.message.Payload = p.parseHeartbeatPayload()
	case Status:
		p.message.Payload = p.parseStatusPayload()
	case Decode:
		p.message.Payload = p.parseDecodePayload()
	case Clear:
		p.message.Payload = p.parseClearPayload()
	case WSPRDecode:
		p.message.Payload = p.parseWSPRDecodePayload()
	case QSOLogged:
		p.message.Payload = p.parseQSOLoggedPayload()
	case LoggedADIF:
		p.message.Payload = p.parseLoggedADIFPayload()
	default:
		fmt.Printf("Cannot parse [%s] yet...\n", p.message.Header.MsgType.String())
	}
}

func (p *WSJTXParser) parseHeartbeatPayload() HeartbeatPayload {
	idString, _ := p.parseUtf8()
	maxSchemaNumber, _ := p.parseQuint32()
	versionString, _ := p.parseUtf8()
	revisionString, _ := p.parseUtf8()

	return HeartbeatPayload{
		ID:              idString,
		MaxSchemaNumber: maxSchemaNumber,
		Version:         versionString,
		Revision:        revisionString,
	}
}

func (p *WSJTXParser) parseStatusPayload() StatusPayload {
	id, _ := p.parseUtf8()
	dialFrequency, _ := p.parseQuint64()
	mode, _ := p.parseUtf8()
	dxCall, _ := p.parseUtf8()
	report, _ := p.parseUtf8()
	txMode, _ := p.parseUtf8()
	txEnabled, _ := p.parseBool()
	transmitting, _ := p.parseBool()
	decoding, _ := p.parseBool()
	rxDF, _ := p.parseQuint32()
	txDF, _ := p.parseQuint32()
	deCall, _ := p.parseUtf8()
	deGrid, _ := p.parseUtf8()
	dxGrid, _ := p.parseUtf8()
	txWatchdog, _ := p.parseBool()
	subMode, _ := p.parseUtf8()
	fastMode, _ := p.parseBool()
	specialOperationMode, _ := p.parseQuint8()
	frequencyTolerance, _ := p.parseQuint32()
	trPeriod, _ := p.parseQuint32()
	configName, _ := p.parseUtf8()
	txMessage, _ := p.parseUtf8()

	return StatusPayload{
		ID:                   id,
		DialFrequency:        dialFrequency,
		Mode:                 mode,
		DxCall:               dxCall,
		Report:               report,
		TxMode:               txMode,
		TxEnabled:            txEnabled,
		Transmitting:         transmitting,
		Decoding:             decoding,
		RxDF:                 rxDF,
		TxDF:                 txDF,
		DECall:               deCall,
		DEGrid:               deGrid,
		DXGrid:               dxGrid,
		TxWatchdog:           txWatchdog,
		SubMode:              subMode,
		FastMode:             fastMode,
		SpecialOperationMode: specialOperationMode,
		FrequencyTolerance:   frequencyTolerance,
		TRPeriod:             trPeriod,
		ConfigurationName:    configName,
		TxMessage:            txMessage,
	}
}

func (p *WSJTXParser) parseDecodePayload() DecodePayload {
	id, _ := p.parseUtf8()
	new, _ := p.parseBool()
	time, _ := p.parseQTime()
	snr, _ := p.parseQint32()
	deltaTime, _ := p.parseFloat64()
	deltaFrequency, _ := p.parseQuint32()
	mode, _ := p.parseUtf8()
	message, _ := p.parseUtf8()
	lowConfidence, _ := p.parseBool()
	offAir, _ := p.parseBool()

	return DecodePayload{
		ID:             id,
		New:            new,
		Time:           time,
		SNR:            snr,
		DeltaTime:      deltaTime,
		DeltaFrequency: deltaFrequency,
		Mode:           mode,
		Message:        message,
		LowConfidence:  lowConfidence,
		OffAir:         offAir,
	}
}

func (p *WSJTXParser) parseClearPayload() ClearPayload {
	id, _ := p.parseUtf8()
	window, _ := p.parseQuint8()

	return ClearPayload{
		ID:     id,
		Window: window,
	}
}

func (p *WSJTXParser) parseQSOLoggedPayload() QSOLoggedPayload {
	id, _ := p.parseUtf8()
	dateTimeOff, _ := p.parseQDateTime()
	dxCall, _ := p.parseUtf8()
	dxGrid, _ := p.parseUtf8()
	txFrequency, _ := p.parseQuint64()
	mode, _ := p.parseUtf8()
	reportSent, _ := p.parseUtf8()
	reportReceived, _ := p.parseUtf8()
	txPower, _ := p.parseUtf8()
	comments, _ := p.parseUtf8()
	name, _ := p.parseUtf8()
	dateTimeOn, _ := p.parseQDateTime()
	operatorCall, _ := p.parseUtf8()
	myCall, _ := p.parseUtf8()
	myGrid, _ := p.parseUtf8()
	exchangeSent, _ := p.parseUtf8()
	exchangeReceived, _ := p.parseUtf8()
	adifPropagationMode, _ := p.parseUtf8()

	return QSOLoggedPayload{
		ID:                  id,
		DateTimeOff:         dateTimeOff,
		DXCall:              dxCall,
		DXGrid:              dxGrid,
		TxFrequency:         txFrequency,
		Mode:                mode,
		ReportSent:          reportSent,
		ReportReceived:      reportReceived,
		TxPower:             txPower,
		Comments:            comments,
		Name:                name,
		DateTimeOn:          dateTimeOn,
		OperatorCall:        operatorCall,
		MyCall:              myCall,
		MyGrid:              myGrid,
		ExchangeSent:        exchangeSent,
		ExchangeReceived:    exchangeReceived,
		ADIFPropagationMode: adifPropagationMode,
	}
}

func (p *WSJTXParser) parseClosed() ClosePayload {
	id, _ := p.parseUtf8()

	return ClosePayload{
		ID: id,
	}
}

func (p *WSJTXParser) parseWSPRDecodePayload() WSPRDecodePayload {
	id, _ := p.parseUtf8()
	new, _ := p.parseBool()
	time, _ := p.parseQTime()
	snr, _ := p.parseQint32()
	deltaTime, _ := p.parseFloat64()
	frequency, _ := p.parseQuint64()
	drift, _ := p.parseQint32()
	callSign, _ := p.parseUtf8()
	grid, _ := p.parseUtf8()
	power, _ := p.parseQint32()
	offAir, _ := p.parseBool()

	return WSPRDecodePayload{
		ID:        id,
		New:       new,
		Time:      time,
		SNR:       snr,
		DeltaTime: deltaTime,
		Frequency: frequency,
		Drift:     drift,
		Callsign:  callSign,
		Grid:      grid,
		Power:     power,
		OffAir:    offAir,
	}
}

func (p *WSJTXParser) parseLoggedADIFPayload() LoggedADIFPayload {
	id, _ := p.parseUtf8()
	adif, _ := p.parseUtf8()

	return LoggedADIFPayload{
		ID:   id,
		ADIF: adif,
	}
}

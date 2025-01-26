wsjtx-udp
==========

Simple Golang library to parse (for now) WSJT-X UDP packets.

## Documentation
Documentation for the WSJT-X UDP format can be found [here](https://sourceforge.net/p/wsjt/wsjtx/ci/master/tree/Network/NetworkMessage.hpp).

This library contains a parser to ease the parsing of WSJT-X UDP packets for Golang. An example UDP server is provided
under the example/ directory.

## TODO
- Add support for outgoing messages (commands sent to WSJT-X)
- Add tests
- Make official build

## WSJT-X
WSJT-X is a software radio for amateur radio operators. For more information, see [WSJT-X](https://wsjt.sourceforge.io//).

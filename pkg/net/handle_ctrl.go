package net

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/rs/zerolog/log"
)

func (l *KCPConn) handleCtrlPacket(addr *net.UDPAddr, buf []byte) (err error) {
	log.Trace().Str("addr", addr.String()).Hex("data", buf).Msg("Recv control packet")
	typ := binary.BigEndian.Uint32(buf[0:4])
	if typ == 0x000000FF {
		err = l.handleCtrlConnectPsh(addr, buf)
	} else if typ == 0x00000145 {
		err = l.handleCtrlConnectAck(addr, buf)
	} else if typ == 0x00000194 {
		err = l.handleCtrlDisconnect(addr, buf)
	} else {
		err = fmt.Errorf("unknown handshake packet type: %d", typ)
	}
	return err
}

func (l *KCPConn) sendCtrlPacket(addr *net.UDPAddr, buf []byte) (err error) {
	log.Trace().Str("addr", addr.String()).Hex("data", buf).Msg("Send control packet")
	_, err = l.conn.WriteToUDP(buf, addr)
	return err
}

func (l *KCPConn) handleCtrlConnectPsh(addr *net.UDPAddr, buf []byte) error {
	if magic := binary.BigEndian.Uint32(buf[16:20]); magic != 0xFFFFFFFF {
		return fmt.Errorf("invalid connect psh magic: %d", magic)
	}
	if data := binary.BigEndian.Uint32(buf[12:16]); data != 0x499602D2 {
		return fmt.Errorf("invalid connect psh data: %d", data)
	}
	conv := binary.BigEndian.Uint32(buf[4:8])
	token := binary.BigEndian.Uint32(buf[8:12])
	if token == 0 {
		token = l.nextSessionToken()
	}
	session, err := l.createSession(addr, token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create session")
		return l.sendCtrlDisconnect(addr, conv, token, 5) // ENET_SERVER_KICK
	}
	return l.sendCtrlConnectAck(addr, session.id, session.token)
}

func (l *KCPConn) sendCtrlConnectPsh(addr *net.UDPAddr, conv, token uint32) error {
	buf := []byte{0x00, 0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x49, 0x96, 0x02, 0xd2, 0xFF, 0xFF, 0xFF, 0xFF}
	binary.BigEndian.PutUint32(buf[4:8], conv)
	binary.BigEndian.PutUint32(buf[8:12], conv)
	return l.sendCtrlPacket(addr, buf)
}

func (l *KCPConn) handleCtrlConnectAck(addr *net.UDPAddr, buf []byte) error {
	if magic := binary.BigEndian.Uint32(buf[16:20]); magic != 0x14514545 {
		return fmt.Errorf("invalid connect ack magic: %d", magic)
	}
	if data := binary.BigEndian.Uint32(buf[12:16]); data != 0x499602D2 {
		return fmt.Errorf("invalid connect ack data: %d", data)
	}
	return nil
}

func (l *KCPConn) sendCtrlConnectAck(addr *net.UDPAddr, conv, token uint32) error {
	buf := []byte{0x00, 0x00, 0x01, 0x45, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x49, 0x96, 0x02, 0xd2, 0x14, 0x51, 0x45, 0x45}
	binary.BigEndian.PutUint32(buf[4:8], conv)
	binary.BigEndian.PutUint32(buf[8:12], token)
	return l.sendCtrlPacket(addr, buf)
}

func (l *KCPConn) handleCtrlDisconnect(addr *net.UDPAddr, buf []byte) error {
	if magic := binary.BigEndian.Uint32(buf[16:20]); magic != 0x19419494 {
		return fmt.Errorf("invalid disconnect magic: %d", magic)
	}
	conv := binary.BigEndian.Uint32(buf[4:8])
	token := binary.BigEndian.Uint32(buf[8:12])
	session, err := l.deleteSession(addr, conv, token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete session")
		if err == ErrSessionNotFound {
			return l.sendCtrlDisconnect(addr, conv, token, 7) // ENET_NOT_FOUND_SESSION
		} else if err == ErrSessionTokenMismatch {
			return l.sendCtrlDisconnect(addr, conv, token, 5) // ENET_SERVER_KICK
		} else {
			return l.sendCtrlDisconnect(addr, conv, token, 5) // ENET_SERVER_KICK
		}
	}
	return l.sendCtrlDisconnect(addr, session.id, session.token, 4) // ENET_SERVER_RELOGIN
}

func (l *KCPConn) sendCtrlDisconnect(addr *net.UDPAddr, conv, token, data uint32) error {
	buf := []byte{0x00, 0x00, 0x01, 0x94, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x19, 0x41, 0x94, 0x94}
	binary.BigEndian.PutUint32(buf[4:8], conv)
	binary.BigEndian.PutUint32(buf[8:12], token)
	binary.BigEndian.PutUint32(buf[12:16], data)
	return l.sendCtrlPacket(addr, buf)
}

package net

import (
	"encoding/binary"
	"net"

	"github.com/rs/zerolog/log"
)

func (l *KCPConn) handleDataPacket(addr *net.UDPAddr, buf []byte) error {
	// log.Printf("[net.KCPConn] Handle data packet: %s\n%s", addr.String(), hex.Dump(buf))
	id := binary.LittleEndian.Uint32(buf[:4])
	token := binary.LittleEndian.Uint32(buf[4:8])
	session, err := l.getSession(addr, id, token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get session")
		return l.sendCtrlDisconnect(addr, id, token, 5) // ENET_SERVER_KICK
	}
	return session.OnPacket(buf, l.packetCh)
}

func (l *KCPConn) sendDataPacket(addr *net.UDPAddr, buf []byte) (err error) {
	// log.Printf("[net.KCPConn] Send data packet: %s\n%s", addr.String(), hex.Dump(buf))
	_, err = l.conn.WriteToUDP(buf, addr)
	return err
}

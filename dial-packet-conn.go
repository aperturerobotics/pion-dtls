// SPDX-FileCopyrightText: 2026 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

//go:build !tinygo

package dtls

import "net"

func dialPacketConn(network string) (*net.UDPConn, error) {
	return net.ListenUDP(network, nil)
}

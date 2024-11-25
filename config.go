// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package mdns

import (
	"net"
	"time"

	"github.com/pion/logging"
)

const (
	// DefaultAddressIPv4 is the default used by mDNS
	// and in most cases should be the address that the
	// ipv4.PacketConn passed to Server is bound to
	DefaultAddressIPv4 = "224.0.0.0:5353"

	// DefaultAddressIPv6 is the default IPv6 address used
	// by mDNS and in most cases should be the address that
	// the ipv6.PacketConn passed to Server is bound to
	DefaultAddressIPv6 = "[FF02::]:5353"
)

type HostBinding struct {
	// Address will override the published address with the given IP
	// when set. Otherwise, the automatically determined address will be used.
	AddressIPv4 net.IP
	// Used for AAAA responses. When set overrides the host derived IP for the
	// given interface
	AddressIPv6 net.IP
	// The interface on which this binding applies. If nil then the binding
	// applies to all interfaces that are not overridden by other bindings.
	// Interface field is unique in bindings. To publish both A and AAAA set the
	// address fields on a single binding per interface.
	Interface *net.Interface
}

type RegisteredHost struct {
	Name     string
	Bindings []HostBinding
}

// Config is used to configure a mDNS client or server.
type Config struct {
	// Name is the name of the client/server used for logging purposes.
	Name string

	// QueryInterval controls how often we sends Queries until we
	// get a response for the requested name
	QueryInterval time.Duration

	// LocalNames are the names that we will generate answers for
	// when we get questions
	LocalNames []RegisteredHost

	LoggerFactory logging.LoggerFactory

	// IncludeLoopback will include loopback interfaces to be eligble for queries and answers.
	IncludeLoopback bool

	Interfaces []net.Interface
}

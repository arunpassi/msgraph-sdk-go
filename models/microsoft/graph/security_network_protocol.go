package graph
import (
    "strings"
    "errors"
)
// 
type SecurityNetworkProtocol int

const (
    UNKNOWN_SECURITYNETWORKPROTOCOL SecurityNetworkProtocol = iota
    IP_SECURITYNETWORKPROTOCOL
    ICMP_SECURITYNETWORKPROTOCOL
    IGMP_SECURITYNETWORKPROTOCOL
    GGP_SECURITYNETWORKPROTOCOL
    IPV4_SECURITYNETWORKPROTOCOL
    TCP_SECURITYNETWORKPROTOCOL
    PUP_SECURITYNETWORKPROTOCOL
    UDP_SECURITYNETWORKPROTOCOL
    IDP_SECURITYNETWORKPROTOCOL
    IPV6_SECURITYNETWORKPROTOCOL
    IPV6ROUTINGHEADER_SECURITYNETWORKPROTOCOL
    IPV6FRAGMENTHEADER_SECURITYNETWORKPROTOCOL
    IPSECENCAPSULATINGSECURITYPAYLOAD_SECURITYNETWORKPROTOCOL
    IPSECAUTHENTICATIONHEADER_SECURITYNETWORKPROTOCOL
    ICMPV6_SECURITYNETWORKPROTOCOL
    IPV6NONEXTHEADER_SECURITYNETWORKPROTOCOL
    IPV6DESTINATIONOPTIONS_SECURITYNETWORKPROTOCOL
    ND_SECURITYNETWORKPROTOCOL
    RAW_SECURITYNETWORKPROTOCOL
    IPX_SECURITYNETWORKPROTOCOL
    SPX_SECURITYNETWORKPROTOCOL
    SPXII_SECURITYNETWORKPROTOCOL
    UNKNOWNFUTUREVALUE_SECURITYNETWORKPROTOCOL
)

func (i SecurityNetworkProtocol) String() string {
    return []string{"UNKNOWN", "IP", "ICMP", "IGMP", "GGP", "IPV4", "TCP", "PUP", "UDP", "IDP", "IPV6", "IPV6ROUTINGHEADER", "IPV6FRAGMENTHEADER", "IPSECENCAPSULATINGSECURITYPAYLOAD", "IPSECAUTHENTICATIONHEADER", "ICMPV6", "IPV6NONEXTHEADER", "IPV6DESTINATIONOPTIONS", "ND", "RAW", "IPX", "SPX", "SPXII", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseSecurityNetworkProtocol(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_SECURITYNETWORKPROTOCOL, nil
        case "IP":
            return IP_SECURITYNETWORKPROTOCOL, nil
        case "ICMP":
            return ICMP_SECURITYNETWORKPROTOCOL, nil
        case "IGMP":
            return IGMP_SECURITYNETWORKPROTOCOL, nil
        case "GGP":
            return GGP_SECURITYNETWORKPROTOCOL, nil
        case "IPV4":
            return IPV4_SECURITYNETWORKPROTOCOL, nil
        case "TCP":
            return TCP_SECURITYNETWORKPROTOCOL, nil
        case "PUP":
            return PUP_SECURITYNETWORKPROTOCOL, nil
        case "UDP":
            return UDP_SECURITYNETWORKPROTOCOL, nil
        case "IDP":
            return IDP_SECURITYNETWORKPROTOCOL, nil
        case "IPV6":
            return IPV6_SECURITYNETWORKPROTOCOL, nil
        case "IPV6ROUTINGHEADER":
            return IPV6ROUTINGHEADER_SECURITYNETWORKPROTOCOL, nil
        case "IPV6FRAGMENTHEADER":
            return IPV6FRAGMENTHEADER_SECURITYNETWORKPROTOCOL, nil
        case "IPSECENCAPSULATINGSECURITYPAYLOAD":
            return IPSECENCAPSULATINGSECURITYPAYLOAD_SECURITYNETWORKPROTOCOL, nil
        case "IPSECAUTHENTICATIONHEADER":
            return IPSECAUTHENTICATIONHEADER_SECURITYNETWORKPROTOCOL, nil
        case "ICMPV6":
            return ICMPV6_SECURITYNETWORKPROTOCOL, nil
        case "IPV6NONEXTHEADER":
            return IPV6NONEXTHEADER_SECURITYNETWORKPROTOCOL, nil
        case "IPV6DESTINATIONOPTIONS":
            return IPV6DESTINATIONOPTIONS_SECURITYNETWORKPROTOCOL, nil
        case "ND":
            return ND_SECURITYNETWORKPROTOCOL, nil
        case "RAW":
            return RAW_SECURITYNETWORKPROTOCOL, nil
        case "IPX":
            return IPX_SECURITYNETWORKPROTOCOL, nil
        case "SPX":
            return SPX_SECURITYNETWORKPROTOCOL, nil
        case "SPXII":
            return SPXII_SECURITYNETWORKPROTOCOL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SECURITYNETWORKPROTOCOL, nil
    }
    return 0, errors.New("Unknown SecurityNetworkProtocol value: " + v)
}
func SerializeSecurityNetworkProtocol(values []SecurityNetworkProtocol) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

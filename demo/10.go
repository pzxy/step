package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	manuf "github.com/timest/gomanuf"
	"log"
	"math"
	"net"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type IP uint32
type info struct {
	Mac   net.HardwareAddr
	Manuf string
}

var (
	infoSet = make(map[string]info)
	ch      = make(chan string)
)

func Table(ipNet *net.IPNet) []IP {
	var (
		min, max IP
		data     []IP
	)
	ip := ipNet.IP.To4()

	for i := 0; i < 4; i++ {
		b := IP(ip[i] & ipNet.Mask[i])
		min += b << ((3 - uint(i)) * 8)
	}
	one, _ := ipNet.Mask.Size()
	max = min | IP(math.Pow(2, float64(32-one))-1)
	for i := min; i < max; i++ {
		if i&0x000000ff == 0 {
			continue
		}
		data = append(data, i)
	}
	return data
}

func (ip IP) String() string {
	var bf bytes.Buffer
	for i := 1; i <= 4; i++ {
		bf.WriteString(strconv.Itoa(int((ip >> ((4 - uint(i)) * 8)) & 0xff)))
		if i != 4 {
			bf.WriteByte('.')
		}
	}
	return bf.String()
}
func sendArpPackage(ip IP, ipNet *net.IPNet, iface net.Interface) {
	localHaddr := iface.HardwareAddr
	srcIp := net.ParseIP(ipNet.IP.String()).To4()
	dstIp := net.ParseIP(ip.String()).To4()
	if srcIp == nil || dstIp == nil {
		log.Fatal("ip error")
	}

	ether := &layers.Ethernet{
		SrcMAC:       localHaddr,
		DstMAC:       net.HardwareAddr{0xff, 0xff, 0xff, 0xff, 0xff, 0xff},
		EthernetType: layers.EthernetTypeARP,
	}

	a := &layers.ARP{
		AddrType:          layers.LinkTypeEthernet,
		Protocol:          layers.EthernetTypeIPv4,
		HwAddressSize:     uint8(6),
		ProtAddressSize:   uint8(4),
		Operation:         uint16(1),
		SourceHwAddress:   localHaddr,
		SourceProtAddress: srcIp,
		DstHwAddress:      net.HardwareAddr{0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		DstProtAddress:    dstIp,
	}

	buffer := gopacket.NewSerializeBuffer()
	var opt gopacket.SerializeOptions
	gopacket.SerializeLayers(buffer, opt, ether, a)
	outgoingPacket := buffer.Bytes()

	handle, err := pcap.OpenLive(iface.Name, 2048, false, pcap.BlockForever)
	if err != nil {
		log.Fatal("pcap open fail, err:", err)
	}
	defer handle.Close()

	err = handle.WritePacketData(outgoingPacket)
	if err != nil {
		log.Fatal("send packet fail")
	}
}

func listenARP(ctx context.Context, iface net.Interface) {
	handle, err := pcap.OpenLive(iface.Name, 1024, false, pcap.BlockForever)
	if err != nil {
		log.Fatal("pcap open fail, err:", err)
	}
	defer handle.Close()

	ps := gopacket.NewPacketSource(handle, handle.LinkType())
	var mutex sync.Mutex
	for {
		select {
		case <-ctx.Done():
			return
		case p := <-ps.Packets():
			arpLayer := p.Layer(layers.LayerTypeARP)
			if arpLayer != nil {
				arp, _ := arpLayer.(*layers.ARP)
				if arp.Operation == 2 {
					mac := net.HardwareAddr(arp.SourceHwAddress)
					m := manuf.Search(mac.String())
					//fmt.Println(ParseIP(arp.SourceProtAddress).String())
					func() {
						ch <- "closeOldTicker"
						defer func() {
							ch <- "startNewTicker"
							mutex.Unlock()
						}()
						mutex.Lock()
						if _, ok := infoSet[ParseIP(arp.SourceProtAddress).String()]; !ok {
							infoSet[ParseIP(arp.SourceProtAddress).String()] = info{mac, m}
						}
					}()
				}
			}
		}
	}
}

func ParseIP(b []byte) IP {
	return IP(IP(b[0])<<24 + IP(b[1])<<16 + IP(b[2])<<8 + IP(b[3]))
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	itf := flag.String("i", "", "NetWork interface name")
	flag.Parse()

	var (
		ifs []net.Interface
		err error
	)
	if *itf == "" {
		ifs, err = net.Interfaces()
	} else {
		it, err := net.InterfaceByName(*itf)
		if err == nil {
			ifs = append(ifs, *it)
		}
	}

	if err != nil {
		log.Fatal("get interface fail, err:", err)
	}

	var (
		ips     []IP
		srcAddr *net.IPNet
		iface   net.Interface
		Flag    = false
	)

	for _, it := range ifs {
		if it.Name == "lo" {
			continue
		}
		addrs, _ := it.Addrs()
		if len(addrs) < 2 {
			continue
		}
		for _, addr := range addrs {
			if ip, ok := addr.(*net.IPNet); ok && !ip.IP.IsLoopback() {
				if ip.IP.To4() != nil {
					ips = Table(ip)
					srcAddr = ip
					iface = it
					Flag = true
					break
				}
			}
		}
		if Flag {
			break
		}
	}
	if !Flag {
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	go listenARP(ctx, iface)

	fmt.Println("********information of localhost***************")
	fmt.Println("interface:", iface.Name)
	fmt.Println("ip:", srcAddr.IP)
	fmt.Println("mac:", iface.HardwareAddr)
	fmt.Println()

	//发送ARP包
	var (
		processNum = 300
		interval   int
		wg         = &sync.WaitGroup{}
	)
	if len(ips) <= processNum {
		interval = 1
		processNum = len(ips)
	} else {
		interval = int(math.Ceil(float64(len(ips)) / float64(processNum)))
	}
	//fmt.Println("interval:", interval)
	//fmt.Println("len(ips):", len(ips))

	for i := 0; i < len(ips); i += interval {
		length := i + interval
		if length >= len(ips) {
			length = len(ips)
		}
		wg.Add(1)
		go func(s []IP) {
			defer wg.Done()
			for _, ip := range s {
				sendArpPackage(ip, srcAddr, iface)
			}
		}(ips[i:length])
	}

	wg.Wait()
	//fmt.Println("wg out!")

	t := time.NewTicker(15 * time.Second)
	for {
		select {
		case <-t.C:
			for ip, data := range infoSet {
				fmt.Printf("%-15s %-20s %-40s\n", ip, data.Mac, data.Manuf)
			}
			cancel()
			return
		case str := <-ch:
			if str == "closeOldTicker" {
				t.Stop()
				//fmt.Println("closeOldTicker")
			} else {
				t = time.NewTicker(5 * time.Second)
				//fmt.Println("startNewTicker")
			}
		}
	}
}

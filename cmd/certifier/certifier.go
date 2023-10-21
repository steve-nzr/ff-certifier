package certifier

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/sirupsen/logrus"

	ffnet "github.com/steve-nzr/ff-certifier/pkg/ff/net"
)

func Run(ctx context.Context) error {
	logrus.Info("Running certifier")

	// connect to account (12000)

	conn, err := net.DialTimeout("tcp", "127.0.0.1:12000", 10*time.Second)
	if err != nil {
		return fmt.Errorf("cannot connect to accountserver: %w", err)
	}

	buffer := make([]byte, 1024)
	readlen, err := conn.Read(buffer)
	if err != nil {
		return fmt.Errorf("cannot read from accountserver: %w", err)
	}

	logrus.Infof("Received %d bytes", readlen)

	packetreader := ffnet.NewPacketReader(buffer)

	// read header
	headerMark, _ := packetreader.ReadByte()
	if headerMark != '^' {
		return fmt.Errorf("invalid headermark: %w", err)
	}

	// read length
	datalen, _ := packetreader.ReadUInt32()
	logrus.Infof("Data length: %d", datalen)

	// Checksum CRC
	checksum, _ := packetreader.ReadUInt32()
	logrus.Infof("Checksum: %d", checksum)

	// Packet ID
	packetid, _ := packetreader.ReadUInt32()
	logrus.Infof("Packet ID: %d", packetid)

	if packetid == 253 {
		// serverlist
		// version
		version, _ := packetreader.ReadString()
		logrus.Infof("Version: %s", version)
	}

	// start neuz server (23000)

	for ctx.Err() == nil {
		time.Sleep(time.Second)
		logrus.Info("Looping !")
	}
	logrus.Info("Ciao")
	return nil
}

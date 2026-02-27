package quake3_rcon

import (
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

const BufferSize = 8192

var PacketPrefix = []byte{'\xff', '\xff', '\xff', '\xff'}

var (
	decolorRegex = regexp.MustCompile(`\^\d`)
	signRegex    = regexp.MustCompile(`[",]`)
)

type Rcon struct {
	ServerIp   string
	ServerPort string
	Password   string
	Connection net.Conn
	mu         sync.Mutex
}

func (r *Rcon) Connect() {
	serverAddress := fmt.Sprintf("%s:%s", r.ServerIp, r.ServerPort)
	conn, err := net.Dial("udp", serverAddress)

	if err != nil {
		fmt.Printf("Error trying to connect to (%s): %v", serverAddress, err)
		os.Exit(-1)
	}

	r.Connection = conn
}

func (r *Rcon) Send(cmd string) error {
	command := fmt.Sprintf("rcon %s %s", r.Password, cmd)
	fullCommandBytes := append(PacketPrefix, []byte(command)...)
	_, err := r.Connection.Write(fullCommandBytes)
	if err != nil {
		return fmt.Errorf("send failed: %w", err)
	}
	return nil
}

func (r *Rcon) Read() (string, error) {
	buffer := make([]byte, BufferSize)

	bytesRead, err := r.Connection.Read(buffer)
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			return "", fmt.Errorf("read timeout: %w", err)
		}
		return "", fmt.Errorf("read failed: %w", err)
	}

	if bytesRead >= 4 {
		return string(buffer[4:bytesRead]), nil
	}

	return "", nil
}

func (r *Rcon) reconnect() {
	log.Warn("Rcon: connection lost, reconnecting...")
	if r.Connection != nil {
		r.Connection.Close()
		r.Connection = nil
	}
	serverAddress := fmt.Sprintf("%s:%s", r.ServerIp, r.ServerPort)
	conn, err := net.Dial("udp", serverAddress)
	if err != nil {
		log.Errorf("Rcon: reconnect failed: %v", err)
		return
	}
	r.Connection = conn
	log.Info("Rcon: reconnected successfully")
}

func (r *Rcon) drain() {
	buf := make([]byte, BufferSize)
	r.Connection.SetDeadline(time.Now().Add(1 * time.Millisecond))
	for {
		if _, err := r.Connection.Read(buf); err != nil {
			break
		}
	}
}

func (r *Rcon) sendWithDeadline(command string) error {
	r.drain()
	if err := r.Connection.SetDeadline(time.Now().Add(2 * time.Second)); err != nil {
		return fmt.Errorf("set deadline failed: %w", err)
	}
	return r.Send(command)
}

func (r *Rcon) RconCommand(command string) string {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.Connection == nil {
		r.reconnect()
		if r.Connection == nil {
			return ""
		}
	}

	if err := r.sendWithDeadline(command); err != nil {
		log.Warnf("Rcon: send error, attempting reconnect: %v", err)
		r.reconnect()
		if r.Connection == nil {
			return ""
		}
		if err := r.sendWithDeadline(command); err != nil {
			log.Errorf("Rcon: send failed after reconnect: %v", err)
			return ""
		}
	}

	res, err := r.Read()
	if err != nil {
		log.Warnf("Rcon: read error: %v", err)
		return ""
	}
	return res
}

func (r *Rcon) RconCommandExtractValue(command string) string {
	fields := r.RconCommand(command)
	fields = strings.Replace(fields, "\n", " ", -1)
	tmpSplit := cleanEmptyLines(strings.Split(fields, " "))

	for _, elem := range tmpSplit {
		if strings.Contains(elem, "is:") {
			decolor := decolorRegex.ReplaceAllString(strings.Split(elem, "is:")[1], "")
			return signRegex.ReplaceAllString(decolor, "")
		}
	}

	return ""
}

func (r *Rcon) CloseConnection() {
	fmt.Println("\nClosing connection ...")
	err := r.Connection.Close()

	if err != nil {
		fmt.Println("Error when closing connection. That's too bad !")
	} else {
		fmt.Println("Successfully closed connection.")
		r.Connection = nil
	}
}

func SplitReadInfos(readstr string) (responseType string, datas []string) {
	lines := cleanEmptyLines(strings.Split(readstr, "\n"))
	return lines[0], lines[1:]
}

func cleanEmptyLines(datas []string) []string {
	var res []string
	for _, value := range datas {
		if value != "" {
			res = append(res, value)
		}
	}
	return res
}

func PrintSplitReadInfos(infos string) {
	fmt.Printf("\n~~~~~~~~~~ Print Read Infos ~~~~~~~~~~")
	cmd, datas := SplitReadInfos(infos)
	fmt.Printf("\nType: %s", cmd)
	datasLength := len(datas)
	if datasLength > 1 {
		fmt.Printf("\nLines: %d\n", len(datas))
	} else {
		fmt.Printf("\nLine: %d\n", len(datas))
	}
	for i, l := range datas {
		fmt.Printf("   |----> %2d) %s\n", i+1, l)
	}
}

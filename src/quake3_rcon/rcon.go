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

func (r *Rcon) Send(cmd string) {
	command := fmt.Sprintf("rcon %s %s", r.Password, cmd)
	commandBytes := []byte(command)

	fullCommandBytes := append(PacketPrefix, commandBytes...)
	_, sendErr := r.Connection.Write(fullCommandBytes)

	if sendErr != nil {
		fmt.Printf("Error while sending command (%s): %v", command, sendErr)
	}
}

func (r *Rcon) Read() (response string) {
	buffer := make([]byte, BufferSize)

	bytesRead, err := r.Connection.Read(buffer)
	if err != nil {
		var netErr net.Error
		if errors.As(err, &netErr) && netErr.Timeout() {
			fmt.Println("No response from server (Timeout)")
			return "" // Return empty instead of exiting
		}

		fmt.Printf("Actual Network Error: %v\n", err)
		return ""
	}

	if bytesRead >= 4 {
		infos := string(buffer[4:bytesRead])
		return infos
	}

	return ""
}

func (r *Rcon) RconCommand(command string) (res string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.Connection == nil {
		return ""
	}

	err := r.Connection.SetDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		return ""
	}

	r.Send(command)
	return r.Read()
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

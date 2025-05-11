package main

import (
	"fmt"

	"github.com/fatih/color"
)

// Define color styles

func helpCommands() {
	green := color.New(color.FgGreen).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	banner := `
╔════════════════════════════════════════════════════════════════════╗
║                         HELP MENU                                  ║
║          Run "help <command>" for further assistance               ║
╚════════════════════════════════════════════════════════════════════╝
`
	fmt.Println(banner)
	fmt.Println("\nAvailable commands:")
	fmt.Printf("  %s        : %s\n", green("ps"), white("Get a process list"))
	fmt.Printf("  %s    : %s\n", green("upload"), white("Upload a file to the server"))
	fmt.Printf("  %s      : %s\n", green("exec"), white("Execute a command on the server"))
	fmt.Printf("  %s        : %s\n", green("ls"), white("List Directories"))
	fmt.Printf("  %s      : %s\n", green("help"), white("Display this help menu"))
	fmt.Printf("  %s      : %s\n", green("quit"), white("Quit the client connection"))
	fmt.Println()
}

func helpPS() {
	green := color.New(color.FgGreen).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	banner := `
╔════════════════════════════════════════════════════════════════════╗
║                          PS HELP                                   ║
║                 Get a target process listing                       ║
╚════════════════════════════════════════════════════════════════════╝
`
	fmt.Println(banner)
	fmt.Println("\nAvailable Args:")
	fmt.Printf("  %s        : %s\n", green("ps"), white("Get a process list"))
	fmt.Printf("  %s   : %s\n", green("help ps"), white("Display this help menu"))
	fmt.Println()
}

func helpExec() {
	green := color.New(color.FgGreen).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	banner := `
╔════════════════════════════════════════════════════════════════════╗
║                        Exec HELP                                   ║
║                   Exec a binary on host                            ║
╚════════════════════════════════════════════════════════════════════╝
`
	fmt.Println(banner)
	fmt.Println("\nAvailable Args:")
	fmt.Printf("  %s        : %s\n", green("exec -b /path/binary 'arg1 arg2'"), white("Exec binary in the background no output returned"))
	fmt.Printf("      Example:  %s\n", green("exec -b /tmp/implant '-a -b 1'"))
	fmt.Printf("      Example:  %s\n", green("exec -b /dev/shm/implant"))
	fmt.Println()
	fmt.Printf("  %s           : %s\n", green("exec /path/binary 'arg1 arg2'"), white("Exec binary in the foreground output returned"))
	fmt.Printf("      Example:  %s\n", green("exec /bin/cat /etc/shadow"))
	fmt.Printf("      Example:  %s\n", green("exec /ram/busybox 'netstat -antpu'"))
	fmt.Println()
	fmt.Printf("  %s                               : %s\n", green("help exec"), white("Display this help menu"))
	fmt.Println()
}

func helpLS() {
	green := color.New(color.FgGreen).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	banner := `
╔════════════════════════════════════════════════════════════════════╗
║                          LS HELP                                   ║
║                    List Directory on Host                          ║
╚════════════════════════════════════════════════════════════════════╝
`
	fmt.Println(banner)
	fmt.Println("\nAvailable Args:")
	fmt.Printf("  %s        : %s\n", green("ls <path>"), white("List directory at specified path"))
	fmt.Printf("      Example:  %s\n", green("ls /var/www/html"))
	fmt.Printf("  %s          : %s\n", green("help ls"), white("Display this help menu"))
	fmt.Println()
}

func helpUpload() {
	green := color.New(color.FgGreen).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()
	banner := `
╔════════════════════════════════════════════════════════════════════╗
║                         UPLOAD HELP                                ║
║                     Upload File to Host                            ║
╚════════════════════════════════════════════════════════════════════╝
`
	fmt.Println(banner)
	fmt.Println("\nAvailable Args:")
	fmt.Printf("  %s      : %s\n", green("upload <local> <remote>"), white("Upload local file to remote path"))
	fmt.Printf("      Example:  %s\n", green("upload /home/ubuntu/implant /tmp/hidden"))
	fmt.Printf("  %s                  : %s\n", green("help upload"), white("Display this help menu"))
	fmt.Println()
}

func banner() {
	banner := `
	                       mm
                           __mTTTTm__
                  __mmmTTTTTTTTTTTTTTTTTTmm__
              _mmTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTmm_
          _mTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTmm_
       _mTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTm_
    _mTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTm_
  _TTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTTT_
_TTTTT**~~~~~~**TTTTTTTTTT**~~TTTT~~**TTTTTTTTTT**~~~~~~**TTTTT_
TT*~              ~*TT*~      TTTT      ~*TT*~              ~*TT
~                             TTTT                             ~
                              TTTT
                              TTTT
                              TTTT
                      mmmm    TTTT
                      TTTT    TTTT     GOTHICUMBRELLA v1.0.0
                      TTTT_  _TTTT
                      *TTTTTTTTTT*
                        ~~~~~~~~
`
	fmt.Println(banner)
}

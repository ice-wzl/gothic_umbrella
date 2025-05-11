# gothic_umbrella
![GOUM art](images/goum.webp)
## Compile 
````
## Standard compilation
// 6.2 MiB
GOARCH=mips GOOS=linux go build server.go
// 5.8 Mib
go build server.go

## Debug Symbols removed 
// 4.3 MiB
GOARCH=mips GOOS=linux go build -ldflags "-s -w" server.go
strip server
// 4.0 MiB
go build -ldflags "-s -w" server.go

## Max Trimmed Compilation
// 3.8 MiB
GOARCH=mips GOOS=linux go build -a -gcflags=all="-l -B" -ldflags="-w -s" server.go
strip server
// 3.6 MiB
go build -a -gcflags=all="-l -B" -ldflags="-w -s" server.go
strip server

// 6.4 MiB 
# Protect Server
go install mvdan.cc/garble@latest
GOARCH=mips GOOS=linux garble -literals -tiny build -a -gcflags=all="-l -B" -ldflags="-s -w" -o server_protected server.go
strip server_protected
// 6.6 MiB
garble -literals -tiny build -a -ldflags="-s -w" -o server_protected server.go 
strip server_protected
````
## Usage 
- connect to a listening gothic_umbrella server like below 
````
./goum_client -host="192.168.15.7" -port=60000
````
## Help Menu 
````
╔════════════════════════════════════════════════════════════════════╗
║                         HELP MENU                                  ║
║          Run "help <command>" for further assistance               ║
╚════════════════════════════════════════════════════════════════════╝


Available commands:
  ps        : Get a process list
  upload    : Upload a file to the server
  exec      : Execute a command on the server
  ls        : List Directories
  help      : Display this help menu
````
## ps
````
╔════════════════════════════════════════════════════════════════════╗
║                          PS HELP                                   ║
║                 Get a target process listing                       ║
╚════════════════════════════════════════════════════════════════════╝


Available Args:
  ps        : Get a process list
  help ps   : Display this help menu
````

## upload 
````
╔════════════════════════════════════════════════════════════════════╗
║                         UPLOAD HELP                                ║
║                     Upload File to Host                            ║
╚════════════════════════════════════════════════════════════════════╝


Available Args:
  upload <local> <remote>      : Upload local file to remote path
      Example:  upload /home/ubuntu/implant /tmp/hidden
  help upload                  : Display this help menu
````

# exec
````
╔════════════════════════════════════════════════════════════════════╗
║                        Exec HELP                                   ║
║                   Exec a binary on host                            ║
╚════════════════════════════════════════════════════════════════════╝


Available Args:
  exec -b /path/binary 'arg1 arg2'        : Exec binary in the background no output returned
      Example:  exec -b /tmp/implant '-a -b 1'
      Example:  exec -b /dev/shm/implant

  exec /path/binary 'arg1 arg2'           : Exec binary in the foreground output returned
      Example:  exec /bin/cat /etc/shadow
      Example:  exec /ram/busybox 'netstat -antpu'

  help exec                               : Display this help menu

````

# ls
````
╔════════════════════════════════════════════════════════════════════╗
║                          LS HELP                                   ║
║                    List Directory on Host                          ║
╚════════════════════════════════════════════════════════════════════╝


Available Args:
  ls <path>        : List directory at specified path
      Example:  ls /var/www/html
  help ls          : Display this help menu

````
## BUGS
- Cannot exec a shell it appears and then dies need to fork it off 
````
Debug command: exec /flash/rw/disk/busybox sh -i
----------------------------------------------------------------------------
[+] /flash/rw/disk/busybox sh -i
----------------------------------------------------------------------------
/ # 
````
## TODO 
- Add help documentation for the commands with color 
- Add someway to kill the binary for real terminate binary listening

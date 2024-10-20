# **CLI app for server ports scanning**

### **Command line interface tool that allows to scan for opened ports on remoted server**
+ Makes TPC requests concurrently 
+ Works with local ports scanning
+ Scans single port or ports range
## Usage
+ Build the app
```Shell
cd cmd && go build .
```
+ Launch with -h flag to see allowed arguments
```
Usage of cmd.exe:
  -h string
        Target hostname (can be IPv4 or domain address), for ex. 'google.com' (default "0.0.0.0")  
  -p string
        Ports range or single port (from 0 to 65535), for ex. '80:100' or '443' (default "0-65535")
  -t uint
        Timeout duration (in milliseconds) (default 500)
```
---
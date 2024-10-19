package presenter

import "fmt"

func LogAtStart(host string, timeout uint, portsRange string) {
	msg := fmt.Sprintf("Starting ports scan...\nOn host: %s\nVia ports range: %s\nWith timeout: %d ms\n\n",
		host,
		portsRange,
		timeout,
	)
	fmt.Println(msg)
}

func LogOpenedPort(host string, port int) {

	fmt.Printf(
		"Port %d is OPENED on %s\n", port, host,
	)

}

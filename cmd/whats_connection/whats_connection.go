package whats_connection

import (
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"github.com/jsdaniell/whats-cli/utils/whats_utils"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// WhatsConnect handles other command, customize it!.
var WhatsConnect = &cobra.Command{
	Use:   "connect",
	Short: "Connect or restore connection to whatsapp on CLI",
	Run: func(cmd *cobra.Command, args []string) {
		// run your different command
		// create new WhatsApp connection

		wac, err := whatsapp.NewConn(20 * time.Second)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)

		}

		err = whats_utils.Login(wac, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)

		}

		<-time.After(3 * time.Second)
	},
}

// WhatsConnectQR handles other command, customize it!.
var WhatsConnectQR = &cobra.Command{
	Use:   "connect-qr",
	Short: "Connect or restore connection to whatsapp on CLI showing the QRCode",
	Run: func(cmd *cobra.Command, args []string) {
		// run your different command
		// create new WhatsApp connection

		wac, err := whatsapp.NewConn(20 * time.Second)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)

		}

		err = whats_utils.Login(wac, false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)

		}

		<-time.After(3 * time.Second)
	},
}

// WhatsConnectQR handles other command, customize it!.
var WhatsVersion = &cobra.Command{
	Use:   "version",
	Short: "Connect or restore connection to whatsapp on CLI showing the QRCode",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("1.1.1")
	},
}

// WhatsConnectQR handles other command, customize it!.
var WhatsReconnect = &cobra.Command{
	Use:   "reconnect",
	Short: "Connect or restore connection to whatsapp on CLI showing the QRCode",
	Run: func(cmd *cobra.Command, args []string) {

		// run your different command
		// create new WhatsApp connection

		wac, err := whatsapp.NewConn(20 * time.Second)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)

		}

		err = whats_utils.ReLogin(wac, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)

		}

		<-time.After(3 * time.Second)
	},
}

package whats_utils

import (
	"encoding/gob"
	"errors"
	"fmt"
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
	"github.com/Rhymen/go-whatsapp"
	"os"
)

func NewLogin(wac *whatsapp.Conn, qrStr bool, sessionID string) error {
	//load saved session
	session, err := readSession(sessionID)
	if err == nil {
		//restore session
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			return errors.New("restoring failed: %v\n" + err.Error())
		}
	} else {
		//no saved session -> regular login
		qr := make(chan string)
		go func() {
			if !qrStr {
				terminal := qrcodeTerminal.New()
				terminal.Get(<-qr).Print()
			} else {
				qrString := <-qr
				fmt.Println(qrString)
			}
		}()

		session, err = wac.Login(qr)
		if err != nil {
			return errors.New("error during login: %v\n" + err.Error())
		}
	}

	//save session
	err = writeSession(session, sessionID)
	if err != nil {
		return errors.New("error saving session: %v\n" + err.Error())
	}

	return nil
}

func Login(wac *whatsapp.Conn, sessionID string) error {
	session, err := readSession(sessionID)
	if err == nil {
		//restore session
		session, err = wac.RestoreWithSession(session)
		if err != nil {
			return errors.New("restoring already logged session failed: %v\n" + err.Error())
		}
	} else {
		return errors.New("restoring already logged session failed: %v\n" + err.Error())
	}

	return nil
}
func Disconnect(sessionID string) error {
	err := os.Remove("whatsappSession-" + sessionID + ".gob")
	if err != nil {
		fmt.Errorf("error saving session: %v\n", err)
	}

	return nil
}

func readSession(sessionID string) (whatsapp.Session, error) {
	session := whatsapp.Session{}
	file, err := os.Open("whatsappSession-" + sessionID + ".gob")
	if err != nil {
		return session, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&session)
	if err != nil {
		return session, err
	}
	return session, nil
}

func writeSession(session whatsapp.Session, sessionID string) error {
	file, err := os.Create("whatsappSession-" + sessionID + ".gob")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(session)
	if err != nil {
		return err
	}
	return nil
}

package controller

import (
	"encoding/binary"
	"io"
	"log"
	"net"
	"time"
	pb "viper/protos/cmds"

	"google.golang.org/protobuf/proto"
)

type Controller struct {
	Addr string
	conn net.Conn
}

func (cnc *Controller) Connect() {
	for {
		log.Print("Connecting to controller.")
		conn, err := net.Dial("tcp", cnc.Addr)
		if err != nil {
			log.Printf("Failed to connect to controller: %v", err)
			time.Sleep(time.Minute * 1)
		} else {
			log.Print("Connected to controller.")
			cnc.conn = conn
			return
		}
	}
}

func (cnc *Controller) ReadCommandRequest() (*pb.CommandRequest, error) {
	for {
		var cmdSize int64
		err := binary.Read(cnc.conn, binary.LittleEndian, &cmdSize)
		if err == io.EOF {
			cnc.Connect()
			continue
		}
		if err != nil {
			return nil, err
		}

		cmdBuffer := make([]byte, cmdSize)
		_, err = cnc.conn.Read(cmdBuffer)
		if err == io.EOF {
			cnc.Connect()
			continue
		}
		if err != nil {
			return nil, err
		}
		cmd := &pb.CommandRequest{}
		err = proto.Unmarshal(cmdBuffer, cmd)
		if err != nil {
			return nil, err
		}
		return cmd, nil
	}
}

func (cnc *Controller) WriteCommandResponse(resp proto.Message) error {
	for {
		respBuffer, err := proto.Marshal(resp)
		if err != nil {
			return err
		}
		respBufferLen := int64(len(respBuffer))
		err = binary.Write(cnc.conn, binary.LittleEndian, &respBufferLen)
		if err == io.EOF {
			cnc.Connect()
			continue
		}
		if err != nil {
			return err
		}
		_, err = cnc.conn.Write(respBuffer)
		if err == io.EOF {
			cnc.Connect()
			continue
		}
		if err != nil {
			return err
		}
		return nil
	}
}
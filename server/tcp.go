package server

import (
	"fmt"
	"net"
	"regexp"
	"strings"

	faker "github.com/bxcodec/faker/v4"
	log "github.com/sirupsen/logrus"
)

type tcpServer struct {
	port int
}

func NewTcpServer(port int) *tcpServer {
	return &tcpServer{
		port: port,
	}
}

func (t tcpServer) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", t.port))
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			//нет смысла останавливать весь сервер если не српботал один конект
			log.Warn(err.Error())
			continue
		}

		go t.handleConnection(conn)
	}
}

func (t tcpServer) handleConnection(conn net.Conn) {
	defer conn.Close()

	content, err := readData(conn)

	if err != nil {
		log.Error(err.Error())
		return
	}

	request := parseHTTPRequest(content)

	switch request.method {
	case "GET":
		request.handleGet()
	case "DELETE":
		request.handleDelete()
	case "POST":
		request.handlePost()
	case "PUT":
		request.handlePut()
	}

	responceWord := faker.Paragraph()

	r := []byte(fmt.Sprintf("HTTP/1.1 200 OK\r\nConnection: close\r\nContent-Type: text/html\r\nContent-Length: %d\r\n\r\n%s", len(responceWord), responceWord))
	conn.Write(r)
}

func readData(conn net.Conn) (string, error) {
	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		return "", err
	}

	return string(buffer[:n]), nil
}

func parseHTTPRequest(request string) *request {
	var methodStr, bodyStr string

	queryParamsMap := make(map[string]string)

	methodPattern := regexp.MustCompile(`^([A-Z]+)\s`)
	queryParamsPattern := regexp.MustCompile(`\?(.*?)\sHTTP`)
	bodyPattern := regexp.MustCompile(`[\r\n]{2}(.*)$`)

	method := methodPattern.FindStringSubmatch(request)
	queryParams := queryParamsPattern.FindStringSubmatch(request)
	body := bodyPattern.FindStringSubmatch(request)

	if len(method) > 1 {
		methodStr = method[1]
	} else {
		methodStr = ""
	}

	if len(body) > 1 {
		bodyStr = body[1]
	} else {
		bodyStr = ""
	}

	if len(queryParams) > 1 {
		params := strings.Split(queryParams[1], "&")
		for _, paramPart := range params {
			param := strings.Split(paramPart, "=")
			queryParamsMap[param[0]] = param[1]
		}
	}

	return NewRequest(methodStr, bodyStr, queryParamsMap)
}

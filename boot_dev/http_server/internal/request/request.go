package request

import (
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

var BAD_START_LINE = fmt.Errorf("Request Line is bad")

func parseRequestLine(requestLine string) (*RequestLine, error) {
	re := regexp.MustCompile(`\s+`)

	s := re.Split(requestLine, -1)

	if len(s) != 3 {
		return nil, fmt.Errorf("Length of requestLine is not 3 which is standard")
	}

	if strings.ToUpper(s[0]) != s[0] {
		return nil, fmt.Errorf("Method is invalid as it contains lowercase alphabetic characters")
	}

	httpVersion := strings.Split(s[2], "/")

	if httpVersion[1] != "1.1" {
		return nil, fmt.Errorf("We dont support http version other than 1.1")
	}

	return &RequestLine{
		Method:        s[0],
		RequestTarget: s[1],
		HttpVersion:   s[2],
	}, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	rawRL := strings.Split(string(data), "\r\n")

	rl, err := parseRequestLine(rawRL[0])

	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *rl,
	}, nil
}

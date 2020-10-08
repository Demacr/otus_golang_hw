package hw10_program_optimization //nolint:golint,stylecheck

import (
	"bufio"
	"io"
	"strings"

	"github.com/buger/jsonparser"
)

type User struct {
	Email string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)
	reader := bufio.NewReader(r)

	for {
		line, _, readerErr := reader.ReadLine()
		if readerErr != nil && readerErr == io.EOF {
			break
		}
		email, err := jsonparser.GetString(line, "Email")
		if err != nil {
			return nil, err
		}
		if strings.HasSuffix(email, "."+domain) {
			result[strings.ToLower(strings.SplitN(email, "@", 2)[1])]++
		}
	}

	return result, nil
}

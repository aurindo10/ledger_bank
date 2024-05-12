package utils

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

type Validator interface {
	Valid() (problems map[string]string)
}

func Encode[T any](status int, v T) *events.APIGatewayProxyResponse {
	encoded, err := json.Marshal(v)
	if err != nil {
		res := events.APIGatewayProxyResponse{
			StatusCode: status,
			Body:       string(encoded),
		}
		return &res
	}
	res := events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(encoded),
	}
	return &res
}

func DecodeValid[T Validator](r *events.APIGatewayProxyRequest) (T, map[string]string, error) {
	var v T

	decoded := json.NewDecoder(strings.NewReader(r.Body))
	if err := decoded.Decode(&v); err != nil {
		return v, nil, fmt.Errorf("decode json: %w", err)
	}
	if problems := v.Valid(); len(problems) > 0 {
		return v, problems, fmt.Errorf("invalid %T: %d problems", v, len(problems))
	}
	return v, nil, nil
}

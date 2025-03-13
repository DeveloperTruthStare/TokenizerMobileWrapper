package tokenizer

import (
	"encoding/json"

	"github.com/ikawaha/kagome/tokenizer"
)

type (
	Tokenizer interface {
		Tokenize(string) string
	}

	srv struct {
		tokenizer tokenizer.Tokenizer
	}

	Token struct {
		Surface  string   `json:"surface"`
		Features []string `json:"features"`
	}
)

func NewTokenizer() Tokenizer {
	return &srv{tokenizer: tokenizer.New()}
}

func (s *srv) Tokenize(input string) string {
	tokens := s.tokenizer.Tokenize(input)
	newTokens := []Token{}

	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			continue
		}
		newTokens = append(newTokens, Token{
			Surface:  token.Surface,
			Features: token.Features(),
		})
	}

	jsonTokens, err := json.Marshal(newTokens)
	if err != nil {
		return "[]"
	}

	return string(jsonTokens)
}

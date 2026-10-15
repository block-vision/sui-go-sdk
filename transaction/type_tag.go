package transaction

import (
	"fmt"
	"strings"

	"github.com/block-vision/sui-go-sdk/models"
)

// ParseTypeTag parses a Move type such as "u64", "vector<u8>", "0x2::sui::SUI" or
// "0x2::coin::Coin<0x2::sui::SUI>" into a TypeTag. Addresses may be short ("0x2").
func ParseTypeTag(s string) (*TypeTag, error) {
	p := typeTagParser{s: s}
	tag, err := p.parseType()
	if err != nil {
		return nil, err
	}
	p.skipSpaces()
	if p.pos != len(p.s) {
		return nil, fmt.Errorf("%w: unexpected %q in %q", ErrInvalidTypeTag, p.s[p.pos:], s)
	}
	return tag, nil
}

type typeTagParser struct {
	s   string
	pos int
}

func (p *typeTagParser) skipSpaces() {
	for p.pos < len(p.s) && p.s[p.pos] == ' ' {
		p.pos++
	}
}

// ident reads an identifier or a 0x-prefixed address.
func (p *typeTagParser) ident() string {
	p.skipSpaces()
	start := p.pos
	for p.pos < len(p.s) {
		c := p.s[p.pos]
		if c == '_' || c >= '0' && c <= '9' || c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
			p.pos++
			continue
		}
		break
	}
	return p.s[start:p.pos]
}

func (p *typeTagParser) consume(tok string) bool {
	p.skipSpaces()
	if strings.HasPrefix(p.s[p.pos:], tok) {
		p.pos += len(tok)
		return true
	}
	return false
}

func (p *typeTagParser) parseType() (*TypeTag, error) {
	unit := &struct{}{}
	name := p.ident()
	switch name {
	case "bool":
		return &TypeTag{Bool: unit}, nil
	case "u8":
		return &TypeTag{U8: unit}, nil
	case "u16":
		return &TypeTag{U16: unit}, nil
	case "u32":
		return &TypeTag{U32: unit}, nil
	case "u64":
		return &TypeTag{U64: unit}, nil
	case "u128":
		return &TypeTag{U128: unit}, nil
	case "u256":
		return &TypeTag{U256: unit}, nil
	case "address":
		return &TypeTag{Address: unit}, nil
	case "signer":
		return &TypeTag{Signer: unit}, nil
	case "vector":
		if !p.consume("<") {
			return nil, fmt.Errorf("%w: expected '<' after vector in %q", ErrInvalidTypeTag, p.s)
		}
		elem, err := p.parseType()
		if err != nil {
			return nil, err
		}
		if !p.consume(">") {
			return nil, fmt.Errorf("%w: expected '>' in %q", ErrInvalidTypeTag, p.s)
		}
		return &TypeTag{Vector: elem}, nil
	}

	structTag, err := p.parseStruct(name)
	if err != nil {
		return nil, err
	}
	return &TypeTag{Struct: structTag}, nil
}

func (p *typeTagParser) parseStruct(address string) (*StructTag, error) {
	hexPart := strings.TrimPrefix(strings.ToLower(address), "0x")
	if !strings.HasPrefix(strings.ToLower(address), "0x") || hexPart == "" || len(hexPart) > 64 {
		return nil, fmt.Errorf("%w: expected a primitive type or 0x-address in %q", ErrInvalidTypeTag, p.s)
	}
	addressBytes, err := ConvertSuiAddressStringToBytes(models.SuiAddress(address))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidTypeTag, err)
	}
	if !p.consume("::") {
		return nil, fmt.Errorf("%w: expected '::' in %q", ErrInvalidTypeTag, p.s)
	}
	module := p.ident()
	if !p.consume("::") {
		return nil, fmt.Errorf("%w: expected '::' in %q", ErrInvalidTypeTag, p.s)
	}
	name := p.ident()
	if module == "" || name == "" {
		return nil, fmt.Errorf("%w: missing module or name in %q", ErrInvalidTypeTag, p.s)
	}

	tag := &StructTag{Address: *addressBytes, Module: module, Name: name, TypeParams: []*TypeTag{}}
	if p.consume("<") {
		for {
			param, err := p.parseType()
			if err != nil {
				return nil, err
			}
			tag.TypeParams = append(tag.TypeParams, param)
			if p.consume(",") {
				continue
			}
			if p.consume(">") {
				break
			}
			return nil, fmt.Errorf("%w: expected ',' or '>' in %q", ErrInvalidTypeTag, p.s)
		}
	}
	return tag, nil
}

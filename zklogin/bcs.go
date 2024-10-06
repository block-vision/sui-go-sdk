package zklogin

type ProofPoints struct {
	A []string   `bcs:"a"`
	B [][]string `bcs:"b"`
	C []string   `bcs:"c"`
}

type IssBase64Details struct {
	Value     string `bcs:"value"`
	IndexMod4 uint8  `bcs:"indexMod4"`
}

type ZkLoginSignatureInputs struct {
	ProofPoints      ProofPoints      `bcs:"proofPoints"`
	IssBase64Details IssBase64Details `bcs:"issBase64Details"`
	HeaderBase64     string           `bcs:"headerBase64"`
	AddressSeed      string           `bcs:"addressSeed"`
}

type ZkLoginSignature struct {
	Inputs        ZkLoginSignatureInputs `bcs:"inputs"`
	MaxEpoch      uint64                 `bcs:"maxEpoch"`
	UserSignature []byte                 `bcs:"userSignature"`
	Iss           string                 `bcs:"iss"`
	AddressSeed   string                 `bcs:"addressSeed"`
}

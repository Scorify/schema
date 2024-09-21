package schema

func ptr[v any](obj v) *v {
	return &obj
}

type exampleDnsSchema struct {
	Server         string `key:"server"`
	Port           int    `key:"port" default:"53"`
	LookupType     string `key:"lookup_type" default:"A" enum:"A,AAAA,CNAME,MX,NS,PTR,SOA,TXT"`
	Domain         string `key:"domain"`
	ExpectedOutput string `key:"expected_output"`
}

var exampleDnsSchemaFields = []*Field{
	{Key: "server", Type: FieldString},
	{Key: "port", Type: FieldInt, Default: ptr("53")},
	{Key: "lookup_type", Type: FieldString, Default: ptr("A"), Enum: &[]string{"A", "AAAA", "CNAME", "MX", "NS", "PTR", "SOA", "TXT"}},
	{Key: "domain", Type: FieldString},
	{Key: "expected_output", Type: FieldString},
}

var exampleDnsSchemaStruct = exampleDnsSchema{
	Server:         "1.1.1.1",
	Port:           53,
	LookupType:     "A",
	Domain:         "one.one.one.one",
	ExpectedOutput: "1.1.1.1",
}

const exampleDnsSchemaString = `{
	"server": "1.1.1.1",
	"port": 53,
	"lookup_type": "A",
	"domain": "one.one.one.one",
	"expected_output": "1.1.1.1"
}`

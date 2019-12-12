package cloudant

const (
	// IndexTypeJSON is for index type json
	IndexTypeJSON IndexType = "json"
	// IndexTypeText is for index type text
	IndexTypeText IndexType = "text"
)

// IndexType is the type of index
type IndexType string

// CreateIndexBuilder defines the available parameter-setting functions
type CreateIndexBuilder interface {
	Fields([]string) CreateIndexBuilder
	DDoc(string) CreateIndexBuilder
	Type(IndexType) CreateIndexBuilder
	Name(string) CreateIndexBuilder
	Partitioned(bool) CreateIndexBuilder
	Build() *createIndex
}

type createIndexBuilder struct {
	fields      []string
	dDoc        string
	iType       string
	name        string
	partitioned bool
}

type createIndexFields struct {
	Fields []string `json:"fields"`
}

// createIndex is the struct for create idnex in cloudant
// From cloudant documenation:
// index fields: A JSON array of field names that uses the sort-syntax. Nested fields are also allowed, for example, "person.name".
// ddoc (optional)	Name of the design document in which the index is created. By default, each index is created in its own design document. Indexes can be grouped into design documents for efficiency. However, a change to one index in a design document invalidates all other indexes in the same document.
// type (optional)	Can be json or text. Defaults to json. Geospatial indexes will be supported in the future.
// name (optional)	Name of the index. If no name is provided, a name is generated automatically.
// partitioned (optional, boolean)	Whether this index is partitioned. For more information, see the partitioned field.
type createIndex struct {
	Index       createIndexFields `json:"index"`
	DDoc        string            `json:"ddoc,omitempty"`
	Type        string            `json:"type,omitempty"`
	Name        string            `json:"name,omitempty"`
	Partitioned bool              `json:"partitioned,omitempty"`
}

// NewCreateIndex is the entrypoint into building a new Index
func NewCreateIndex() CreateIndexBuilder {
	return &createIndexBuilder{}
}

func (cib *createIndexBuilder) Fields(fields []string) CreateIndexBuilder {
	cib.fields = fields
	return cib
}

func (cib *createIndexBuilder) DDoc(ddoc string) CreateIndexBuilder {
	cib.dDoc = ddoc
	return cib
}

func (cib *createIndexBuilder) Type(indexType IndexType) CreateIndexBuilder {
	cib.iType = string(indexType)
	return cib
}

func (cib *createIndexBuilder) Name(name string) CreateIndexBuilder {
	cib.name = name
	return cib
}

func (cib *createIndexBuilder) Partitioned(partitioned bool) CreateIndexBuilder {
	cib.partitioned = partitioned
	return cib
}

func (cib *createIndexBuilder) Build() *createIndex {
	return &createIndex{
		Index: createIndexFields{
			Fields: cib.fields,
		},
		DDoc:        cib.dDoc,
		Type:        cib.iType,
		Name:        cib.name,
		Partitioned: cib.partitioned,
	}
}

// CreateIndexResponse is the response from cloudant from create index
type CreateIndexResponse struct {
	Result string `json:"result"`
}

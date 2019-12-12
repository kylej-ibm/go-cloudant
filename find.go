package cloudant

// FindBuilder defines the available parameter-setting functions
type FindBuilder interface {
	SetSelector(string, interface{}) FindBuilder
	Limit(int) FindBuilder
	Skip(int) FindBuilder
	AddSort(interface{}) FindBuilder
	Fields([]string) FindBuilder
	R(int) FindBuilder
	Bookmark(string) FindBuilder
	UseIndex(string) FindBuilder
	Conflicts(bool) FindBuilder
	ExecutionStats(bool) FindBuilder
	Build() *find
}

type findBuilder struct {
	selector       map[string]interface{}
	limit          int
	skip           int
	sort           []interface{}
	fields         []string
	r              int
	bookmark       string
	useIndex       string
	conflicts      bool
	executionStats bool
}

// find is the struct for query in cloudant
// From cloudant documentation:
// selector: JSON object that describes the criteria that are used to select documents. More information is provided in the section on selectors.
// limit (optional, default: 25): Maximum number of results returned. The type: text indexes are limited to 200 results when queried.
// skip (optional, default: 0): Skip the first 'n' results, where 'n' is the value that is specified.
// sort (optional, default: []): JSON array, ordered according to the sort syntax.
// fields (optional, default: null): JSON array that uses the field syntax as described in the following information. Use this parameter to specify which fields of an object must be returned. If it is omitted, the entire object is returned.
// r (optional, default: 1): The read quorum used when reading documents required when processing a query. The value defaults to 1, in which case, the document is read from the primary data co-located with the index. If set to a higher value, the document must also be retrieved from at least r-1 other primary data replicas before results can be processed. This option increases query latency as the replicas reside on separate machines. In practice, this option should hardly ever be changed from the default.
// r is disallowed when making a partition query.
// bookmark (optional, default: null): A string that is used to specify which page of results you require. For more information, see Pagination.
// use_index (optional): Use this option to identify a specific index for query to run against, rather than by using the IBM Cloudant Query algorithm to find the best index. For more information, see Explain plans.
// conflicts (optional, default: false): A Boolean value that indicates whether or not to include information about existing conflicts in the document.
// execution_stats (optional, default: false): Use this option to find information about the query that was run. This information includes total key lookups, total document lookups (when include_docs=true is used), and total quorum document lookups (when Fabric document lookups are used).
type find struct {
	Selector       map[string]interface{} `json:"selector"`
	Limit          int                    `json:"limit,omitempty"`
	Skip           int                    `json:"skip,omitempty"`
	Sort           []interface{}          `json:"sort,omitempty"`
	Fields         []string               `json:"fields,omitempty"`
	R              int                    `json:"r,omitempty"`
	Bookmark       string                 `json:"bookmark,omitempty"`
	UseIndex       string                 `json:"use_index,omitempty"`
	Conflicts      bool                   `json:"conflicts,omitempty"`
	ExecutionStats bool                   `json:"execution_stats,omitempty"`
}

// NewFind is the entrypoint into building a new Query
func NewFind() FindBuilder {
	return &findBuilder{}
}

func (fb *findBuilder) SetSelector(field string, selector interface{}) FindBuilder {
	if fb.selector == nil {
		fb.selector = make(map[string]interface{})
	}
	fb.selector[field] = selector
	return fb
}

func (fb *findBuilder) Limit(limit int) FindBuilder {
	fb.limit = limit
	return fb
}

func (fb *findBuilder) Skip(skip int) FindBuilder {
	fb.skip = skip
	return fb
}

func (fb *findBuilder) AddSort(sort interface{}) FindBuilder {
	if fb.sort == nil {
		fb.sort = make([]interface{}, 0)
	}
	fb.sort = append(fb.sort, sort)
	return fb
}

func (fb *findBuilder) Fields(fields []string) FindBuilder {
	fb.fields = fields
	return fb
}

func (fb *findBuilder) R(r int) FindBuilder {
	fb.r = r
	return fb
}

func (fb *findBuilder) Bookmark(bookmark string) FindBuilder {
	fb.bookmark = bookmark
	return fb
}

func (fb *findBuilder) UseIndex(useIndex string) FindBuilder {
	fb.useIndex = useIndex
	return fb
}

func (fb *findBuilder) Conflicts(conflicts bool) FindBuilder {
	fb.conflicts = conflicts
	return fb
}

func (fb *findBuilder) ExecutionStats(es bool) FindBuilder {
	fb.executionStats = es
	return fb
}

func (fb *findBuilder) Build() *find {
	return &find{
		Selector:       fb.selector,
		Limit:          fb.limit,
		Skip:           fb.skip,
		Sort:           fb.sort,
		Fields:         fb.fields,
		R:              fb.r,
		Bookmark:       fb.bookmark,
		UseIndex:       fb.useIndex,
		Conflicts:      fb.conflicts,
		ExecutionStats: fb.executionStats,
	}
}

// FindResponse is the response from cloudant from a find query
type FindResponse struct {
	Docs           []interface{}   `json:"docs"`
	Bookmark       string          `json:"bookmark,omitempty"`
	ExecutionStats executionStatus `json:"execution_stats,omitempty"`
}

type executionStatus struct {
	TotalKeysExamined       int     `json:"total_keys_examined"`
	TotalDocsExamined       int     `json:"total_docs_examined"`
	TotalQuorumDocsExamined int     `json:"total_quorum_docs_examined"`
	ResultsReturned         int     `json:"results_returned"`
	ExecutionTimeMS         float64 `json:"execution_time_ms"`
}

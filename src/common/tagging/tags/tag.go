package tags

type Tag struct {
	Key   string
	Value string
}

var TagTypes = []ITag{
	&GitOrgTag{},
	&GitRepoTag{},
	&GitFileTag{},
	&GitCommitTag{},
	&GitModifiersTag{},
	&GitLastModifiedAtTag{},
	&GitLastModifiedByTag{},
	&YorTraceTag{},
}

type ITag interface {
	Init()
	CalculateValue(data interface{}) (ITag, error)
	GetKey() string
	GetValue() string
	GetPriority() int
}

type TagDiff struct {
	Key       string
	PrevValue string
	NewValue  string
}

func Init(key string, value string) ITag {
	return &Tag{
		Key:   key,
		Value: value,
	}
}

func (t *Tag) Init() {}

func (t *Tag) GetPriority() int {
	return 0
}

func (t *Tag) CalculateValue(_ interface{}) (ITag, error) {
	return t, nil
}

func (t *Tag) GetKey() string {
	return t.Key
}

func (t *Tag) GetValue() string {
	return t.Value
}
package rtu

var queryDefaultLimit int = 1000

var urlMake = func(s string) string {
	return "https://" + s + "/service/service.php?soap"
}

type SortType string

const (
	SortAsc SortType = "asc"
	SortDesc SortType = "desc"
)

type TableName string

const (
	TableCDRHour TableName = "02.2205.01"
	TableCDRDay TableName = "02.2206.01"
	TableCDRWeek TableName = "02.2207.01"
	TableCDRMonth TableName = "02.2208.01"
	TableCDRAll TableName = "02.2204.01"
	TablePrerouting TableName = "02.2211.01"
	TableEquipment TableName = "02.2201.01"
	TableDialPeers TableName = "02.2202.01"
)

type methodType string

const (
	selectMethod methodType = "selectRowset"
	insertMethod methodType = "insertRowset"
	updateMethod methodType = "updateRowset"
	deleteMethod methodType = "deleteRowset"
	countMethod methodType = "countRowset"
	describeMethod methodType = "describeColumns"
)

type filterType string

const (
	complexFilter filterType = "agg"
	simpleFilter filterType = "cond"
)

type FilterHandleCondition string

const (
	FilterHandleOR FilterHandleCondition = "or"
	FilterHandleAND FilterHandleCondition = "and"
)

var filterConditions = []string{"<=", ">=", "<>", "like", "not like", "regexp", ">", "<", "="}



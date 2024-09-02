package lint

const targetTypeQuery = "query"

const (
	panelTypeStat       = "stat"
	panelTypeSingleStat = "singlestat"
	panelTypeGauge      = "gauge"
	panelTypeGraph      = "graph"
	panelTypeTimeSeries = "timeseries"
	panelTypeTimeTable  = "table"
)

const (
	ruleStabilityStable       = "stable"
	ruleStabilityExperimental = "experimental"
	ruleStabilityDeprecated   = "deprecated"
)

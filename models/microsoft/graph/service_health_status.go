package graph
import (
    "strings"
    "errors"
)
// 
type ServiceHealthStatus int

const (
    SERVICEOPERATIONAL_SERVICEHEALTHSTATUS ServiceHealthStatus = iota
    INVESTIGATING_SERVICEHEALTHSTATUS
    RESTORINGSERVICE_SERVICEHEALTHSTATUS
    VERIFYINGSERVICE_SERVICEHEALTHSTATUS
    SERVICERESTORED_SERVICEHEALTHSTATUS
    POSTINCIDENTREVIEWPUBLISHED_SERVICEHEALTHSTATUS
    SERVICEDEGRADATION_SERVICEHEALTHSTATUS
    SERVICEINTERRUPTION_SERVICEHEALTHSTATUS
    EXTENDEDRECOVERY_SERVICEHEALTHSTATUS
    FALSEPOSITIVE_SERVICEHEALTHSTATUS
    INVESTIGATIONSUSPENDED_SERVICEHEALTHSTATUS
    RESOLVED_SERVICEHEALTHSTATUS
    MITIGATEDEXTERNAL_SERVICEHEALTHSTATUS
    MITIGATED_SERVICEHEALTHSTATUS
    RESOLVEDEXTERNAL_SERVICEHEALTHSTATUS
    CONFIRMED_SERVICEHEALTHSTATUS
    REPORTED_SERVICEHEALTHSTATUS
    UNKNOWNFUTUREVALUE_SERVICEHEALTHSTATUS
)

func (i ServiceHealthStatus) String() string {
    return []string{"SERVICEOPERATIONAL", "INVESTIGATING", "RESTORINGSERVICE", "VERIFYINGSERVICE", "SERVICERESTORED", "POSTINCIDENTREVIEWPUBLISHED", "SERVICEDEGRADATION", "SERVICEINTERRUPTION", "EXTENDEDRECOVERY", "FALSEPOSITIVE", "INVESTIGATIONSUSPENDED", "RESOLVED", "MITIGATEDEXTERNAL", "MITIGATED", "RESOLVEDEXTERNAL", "CONFIRMED", "REPORTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseServiceHealthStatus(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "SERVICEOPERATIONAL":
            return SERVICEOPERATIONAL_SERVICEHEALTHSTATUS, nil
        case "INVESTIGATING":
            return INVESTIGATING_SERVICEHEALTHSTATUS, nil
        case "RESTORINGSERVICE":
            return RESTORINGSERVICE_SERVICEHEALTHSTATUS, nil
        case "VERIFYINGSERVICE":
            return VERIFYINGSERVICE_SERVICEHEALTHSTATUS, nil
        case "SERVICERESTORED":
            return SERVICERESTORED_SERVICEHEALTHSTATUS, nil
        case "POSTINCIDENTREVIEWPUBLISHED":
            return POSTINCIDENTREVIEWPUBLISHED_SERVICEHEALTHSTATUS, nil
        case "SERVICEDEGRADATION":
            return SERVICEDEGRADATION_SERVICEHEALTHSTATUS, nil
        case "SERVICEINTERRUPTION":
            return SERVICEINTERRUPTION_SERVICEHEALTHSTATUS, nil
        case "EXTENDEDRECOVERY":
            return EXTENDEDRECOVERY_SERVICEHEALTHSTATUS, nil
        case "FALSEPOSITIVE":
            return FALSEPOSITIVE_SERVICEHEALTHSTATUS, nil
        case "INVESTIGATIONSUSPENDED":
            return INVESTIGATIONSUSPENDED_SERVICEHEALTHSTATUS, nil
        case "RESOLVED":
            return RESOLVED_SERVICEHEALTHSTATUS, nil
        case "MITIGATEDEXTERNAL":
            return MITIGATEDEXTERNAL_SERVICEHEALTHSTATUS, nil
        case "MITIGATED":
            return MITIGATED_SERVICEHEALTHSTATUS, nil
        case "RESOLVEDEXTERNAL":
            return RESOLVEDEXTERNAL_SERVICEHEALTHSTATUS, nil
        case "CONFIRMED":
            return CONFIRMED_SERVICEHEALTHSTATUS, nil
        case "REPORTED":
            return REPORTED_SERVICEHEALTHSTATUS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_SERVICEHEALTHSTATUS, nil
    }
    return 0, errors.New("Unknown ServiceHealthStatus value: " + v)
}
func SerializeServiceHealthStatus(values []ServiceHealthStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

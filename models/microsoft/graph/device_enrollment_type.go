package graph
import (
    "strings"
    "errors"
)
// 
type DeviceEnrollmentType int

const (
    UNKNOWN_DEVICEENROLLMENTTYPE DeviceEnrollmentType = iota
    USERENROLLMENT_DEVICEENROLLMENTTYPE
    DEVICEENROLLMENTMANAGER_DEVICEENROLLMENTTYPE
    APPLEBULKWITHUSER_DEVICEENROLLMENTTYPE
    APPLEBULKWITHOUTUSER_DEVICEENROLLMENTTYPE
    WINDOWSAZUREADJOIN_DEVICEENROLLMENTTYPE
    WINDOWSBULKUSERLESS_DEVICEENROLLMENTTYPE
    WINDOWSAUTOENROLLMENT_DEVICEENROLLMENTTYPE
    WINDOWSBULKAZUREDOMAINJOIN_DEVICEENROLLMENTTYPE
    WINDOWSCOMANAGEMENT_DEVICEENROLLMENTTYPE
    WINDOWSAZUREADJOINUSINGDEVICEAUTH_DEVICEENROLLMENTTYPE
    APPLEUSERENROLLMENT_DEVICEENROLLMENTTYPE
    APPLEUSERENROLLMENTWITHSERVICEACCOUNT_DEVICEENROLLMENTTYPE
)

func (i DeviceEnrollmentType) String() string {
    return []string{"UNKNOWN", "USERENROLLMENT", "DEVICEENROLLMENTMANAGER", "APPLEBULKWITHUSER", "APPLEBULKWITHOUTUSER", "WINDOWSAZUREADJOIN", "WINDOWSBULKUSERLESS", "WINDOWSAUTOENROLLMENT", "WINDOWSBULKAZUREDOMAINJOIN", "WINDOWSCOMANAGEMENT", "WINDOWSAZUREADJOINUSINGDEVICEAUTH", "APPLEUSERENROLLMENT", "APPLEUSERENROLLMENTWITHSERVICEACCOUNT"}[i]
}
func ParseDeviceEnrollmentType(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_DEVICEENROLLMENTTYPE, nil
        case "USERENROLLMENT":
            return USERENROLLMENT_DEVICEENROLLMENTTYPE, nil
        case "DEVICEENROLLMENTMANAGER":
            return DEVICEENROLLMENTMANAGER_DEVICEENROLLMENTTYPE, nil
        case "APPLEBULKWITHUSER":
            return APPLEBULKWITHUSER_DEVICEENROLLMENTTYPE, nil
        case "APPLEBULKWITHOUTUSER":
            return APPLEBULKWITHOUTUSER_DEVICEENROLLMENTTYPE, nil
        case "WINDOWSAZUREADJOIN":
            return WINDOWSAZUREADJOIN_DEVICEENROLLMENTTYPE, nil
        case "WINDOWSBULKUSERLESS":
            return WINDOWSBULKUSERLESS_DEVICEENROLLMENTTYPE, nil
        case "WINDOWSAUTOENROLLMENT":
            return WINDOWSAUTOENROLLMENT_DEVICEENROLLMENTTYPE, nil
        case "WINDOWSBULKAZUREDOMAINJOIN":
            return WINDOWSBULKAZUREDOMAINJOIN_DEVICEENROLLMENTTYPE, nil
        case "WINDOWSCOMANAGEMENT":
            return WINDOWSCOMANAGEMENT_DEVICEENROLLMENTTYPE, nil
        case "WINDOWSAZUREADJOINUSINGDEVICEAUTH":
            return WINDOWSAZUREADJOINUSINGDEVICEAUTH_DEVICEENROLLMENTTYPE, nil
        case "APPLEUSERENROLLMENT":
            return APPLEUSERENROLLMENT_DEVICEENROLLMENTTYPE, nil
        case "APPLEUSERENROLLMENTWITHSERVICEACCOUNT":
            return APPLEUSERENROLLMENTWITHSERVICEACCOUNT_DEVICEENROLLMENTTYPE, nil
    }
    return 0, errors.New("Unknown DeviceEnrollmentType value: " + v)
}
func SerializeDeviceEnrollmentType(values []DeviceEnrollmentType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}

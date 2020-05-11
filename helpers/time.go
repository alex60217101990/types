package helpers

import (
    "strconv"
    "time"
)

func DurationSecond(seconds interface{}, defaultValue int) time.Duration {
    switch s := seconds.(type) {
    case *int:
        return time.Duration(*s) * time.Second
    case int:
        return time.Duration(s) * time.Second
    case *string:
        if i32, err := strconv.Atoi(*s); err != nil {
            return time.Duration(defaultValue) * time.Second
        } else {
            return time.Duration(i32) * time.Second
        }
    case string:
        if i32, err := strconv.Atoi(s); err != nil {
            return time.Duration(defaultValue) * time.Second
        } else {
            return time.Duration(i32) * time.Second
        }
    default:
        return time.Duration(defaultValue) * time.Second
    }
}

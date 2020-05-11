package helpers

import (
    "fmt"
    "github.com/golang/protobuf/ptypes"
    tspb "github.com/golang/protobuf/ptypes/timestamp"
    "time"
)

func AdaptTimeToPbTimestamp(currentTime *time.Time) (*tspb.Timestamp, error) {
    if currentTime != nil && !(*currentTime).IsZero() {
        protoTime := &tspb.Timestamp{}
        protoTime, err := ptypes.TimestampProto(TimePtrToTime(currentTime))
        if err != nil {
            return nil, err
        }
        return protoTime, nil
    }
    return nil, fmt.Errorf("time parameter is empty or zero")
}

func AdaptPbTimestampToTime(protoTime *tspb.Timestamp) (*time.Time, error) {
    if protoTime == nil || (protoTime.GetNanos() == 0 || protoTime.GetSeconds() == 0) {
        return nil, fmt.Errorf("proto time parameter is empty or zero")
    }
    if t, err := ptypes.Timestamp(protoTime); err == nil {
        return &t, nil
    } else {
        time := time.Unix(protoTime.GetSeconds(), int64(protoTime.GetNanos()))
        return &time, nil
    }
}

// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package util

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	return &v
}

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int {
	return &v
}

// Int32 is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it.
func Int32(v int32) *int32 {
	return &v
}

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 {
	return &v
}

// Float32 is a helper routine that allocates a new float32 value
// to store v and returns a pointer to it.
func Float32(v float32) *float32 {
	return &v
}

// Float64 is a helper routine that allocates a new float64 value
// to store v and returns a pointer to it.
func Float64(v float64) *float64 {
	return &v
}

// Uint32 is a helper routine that allocates a new uint32 value
// to store v and returns a pointer to it.
func Uint32(v uint32) *uint32 {
	return &v
}

// Uint64 is a helper routine that allocates a new uint64 value
// to store v and returns a pointer to it.
func Uint64(v uint64) *uint64 {
	return &v
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	return &v
}

//Verify timestamp syntactic sugar
func VtimestampStr(t string, s time.Duration) (err error) {
    timestamp, err := strconv.ParseInt(t, 10, 64)
    if err != nil {
        err = errors.New("timestamp is invalid: " + t)
        return
    }
    err = Vtimestamp(timestamp,s)
    return
}

//Verify the timestamp
func Vtimestamp(t int64, s time.Duration) (err error) {
    serveTime := time.Unix(t, 0)
    if time.Now().Add(s*time.Second).Before(serveTime) && time.Now().Add(-1*s*time.Second).After(serveTime) {
        err = errors.New("timestamp is expired.")
    }
    return
}

package goutil

// A helper function behaves like map in javascript
func Map(arr []interface{}, f func(interface{}) interface{}) []interface{} {
    result := make([]interface{}, len(arr))
    for i, v := range arr {
        result[i] = f(v)
    }
    return result
}


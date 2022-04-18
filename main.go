package goutil

// Compare two slices for equality.
func SliceEqual[T comparable](a []T, b []T) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

// Apply a function to each element of a slice and return the result.
func SliceMap[T any, M any](arr []T, f func(T) M) []M {
    result := make([]M, len(arr))
    for i, v := range arr {
        result[i] = f(v)
    }
    return result
}

func SliceFilter[T any](arr []T, f func(T) bool) []T {
    result := make([]T, 0)
    for _, v := range arr {
        if f(v) {
            result = append(result, v)
        }
    }
    return result
}

func SliceReduce[T any, M any](arr []T, f func(M, T) M, initial M) M {
    result := initial
    for _, v := range arr {
        result = f(result, v)
    }
    return result
}

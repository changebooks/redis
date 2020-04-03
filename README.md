# redis
缓存
==

<pre>
data := make(map[string]string)
data[redis.ProfileHost] = "127.0.0.1"

profile, err := redis.NewProfile(data)
if err != nil {
    fmt.Println(err)
    return
}

schema, err := redis.NewSchema(profile)
if err != nil {
    fmt.Println(err)
    return
}

executor, err := redis.NewExecutor(schema)
if err != nil {
    fmt.Println(err)
    return
}

result, err, closeErr := executor.SetEx("id", 10001, 3600)
if err != nil {
    fmt.Println(err)
    return
}

if closeErr != nil {
    fmt.Println(closeErr)
    return
}

fmt.Println(result)

value, err, closeErr := executor.Get("id")
if err != nil {
    fmt.Println(err)
    return
}

if closeErr != nil {
    fmt.Println(closeErr)
    return
}

fmt.Println(value)
</pre>

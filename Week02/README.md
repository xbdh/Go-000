### 问题

我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答案：应该wrap这个error， 是强制让调用者使用errors.Is来处理，      包装返回方法的作者就可以随意添加，更改附加的error上下文，无论底层数据库逻辑如何变，都是返回sql.ErrNoRows。不破坏调用方的结构和判断。

如果不包装，调用方可能用==判断，考虑到兼容性，底层逻辑处理不可随意添加error上下文。



```go
errors.Wrap(sql.ErrNoRows,"context")
或
fmt.Errorf("context: %w", sql.ErrNoRows)、


调用者可以使用errors.Is处理err
if errors.Is(err, sql.ErrNoRows){
    
}
```

https://blog.golang.org/go1.13-errors


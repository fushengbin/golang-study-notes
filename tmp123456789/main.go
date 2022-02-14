package main

import (
    "errors"
    "fmt"

    // 注意：这里会自动执行inc 的 init
    _ "GolangStudyNotes/tmp123456789/inc"
    "github.com/gookit/goutil/dump"
)

type I interface {
    String() error
}

type T struct{}

func (t *T) String() error {
    fmt.Println("call string")
    return errors.New("call string error")
}

func main() {
    // 下划线作为变量名，无法被使用
    var _ int
    var _ I = &T{}
    _ = 3

    // 忽略返回的值
    _ = new(T).String()

    m := map[string]string{"hello": "hello", "world": "world"}
    for _, _ = range m {
        dump.P("不接收遍历m的每一个key与value")
    }

    dump.P("call main")
}

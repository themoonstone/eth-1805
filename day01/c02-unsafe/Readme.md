## unsafe包
1. 引用传递和值传递
值传递：``int a = 10, int b = 20, a = b``
引用传递：`` var s []string, var s1 []string, s = s1``
2. unsafe:go里面的指针(不推荐使用)
3. 方法说明
    1. func Alignof(x ArbitraryType) uintptr    // 对齐
    2. func Offsetof(x ArbitraryType) uintptr   // 偏移量
    3. func Sizeof(x ArbitraryType) uintptr     // 变量大小
    4. type Pointer *ArbitraryType      // 指针类型

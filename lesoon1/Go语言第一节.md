## **Go语言第一节**

#### 数据类型

##### 布尔型

##### 数字类型

| 类型       | 描述                                                         |
| ---------- | ------------------------------------------------------------ |
| uint8      | 无符号 8 位整型 (0 到 255)                                   |
| uint16     | 无符号 16 位整型 (0 到 65535)                                |
| uint32     | 无符号 32 位整型 (0 到 4294967295)                           |
| uint64     | 无符号 64 位整型 (0 到 18446744073709551615)                 |
| int8       | 有符号 8 位整型 (-128 到 127)                                |
| int16      | 有符号 16 位整型 (-32768 到 32767)                           |
| int32      | 有符号 32 位整型 (-2147483648 到 2147483647)                 |
| int64      | 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) |
| float32    | IEEE-754 32位浮点型数                                        |
| float64    | IEEE-754 64位浮点型数                                        |
| complex64  | 32 位实数和虚数                                              |
| complex128 | 64 位实数和虚数                                              |
| int        | 根据你的操作系统架构，可以是32位或64位的整数                 |
| byte       | uint8的别名                                                  |
| rune       | int32的别名                                                  |

##### 字符串类型

不可变的字节数组，由指针和长度组成

##### 派生类型

Pointer指针类型

数组类型（指针＋长度）

结构化类型（struct)

Channel类型

函数类型

切片类型

接口类型（interface）

Map类型

#### 变量声明

```go
//第一种方式
//var 变量名 [类型] = 表达式
//example
//类型推导
var a="first"
var b,c=1,2
//指定类型
var d int8=1
v e float64
e=1.23
//第二种方式
f:=2
```

#### 常量

恒定不变的值

```go
const s string = "constant"
const h = 500000000
const i = 3e20 / h
fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
```

#### 循环

唯一关键字for

while形式

```go
for i<5{
  // 循环体
}
```

for本身形式

```go
for i:=0;i<1;i++{
   //循环体
}
```

#### if

一样

#### switch

和C一样但不用写break，无需担心穿透

```go
switch a {
	case 1:
		fmt.Println("one")
	case 2:
		// 在此打印"two"并跳出
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4, 5:
		fmt.Println("four or five")
	default:
		fmt.Println("other")
	}
```

#### 函数

函数定义由函数名、参数列表、返回类型和函数体构成

```go
func add(x int, y int) int {
    return x + y
}
```

参数一样，返回值可以返回多值，如果无须返回值那返回类型就空着

**函数是一等公民**：函数和数组，字符串等等定位相同，可以作为元素传递，可以创建函数数组


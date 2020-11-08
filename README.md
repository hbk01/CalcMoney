# Calculate Money

读取账单并计算几个人逛街时每个人的花费

# 使用

```shell script
usage:
	cm [-d|--debug] file
		-d, --deubg: 输出更详细的信息
		file: 账单（文本格式）

example:
    cm -d src.txt
    cm --debug src.txt
    cm src.txt
```

# 账单的格式

### \#define

`#define`语句定义账单中涉及到的所有用户。

`#define` 语句应当写在文件的开头，文件中可以有多个 `#define` ，一个 `#define` 只能定义一个用户。

**`#define`之后的所有用户均使用`#define`的名称。**

```
#define a name
```

`a`：在后面的 `Item` 中使用 `a` 来代表 `name` 用户

`name`：用户的真实姓名

### \#default

`#default`语句定义默认的付钱用户。

`#defult` 语句应当写在`#define`之后，`Item`之前，文件中**只能有一个** `#default`，`#default`是可选的，未定义的将使用第一个 `#define` 作为付钱用户。

```
#default a
```

`a`：使用 `a` 作为默认的付钱用户，`a`需要在 `#define`中定义

###  Item

`Item`是具体的账单情况

```
Item 0.0 a[,b,c...][-d]
```

`Item`: 买的什么东西

`0.0`: 这个东西多少钱

`a[b,c,d...]`: 谁一起拼单

`-d`: 谁付的钱，不填将使用默认用户

**小提示：谁一起拼单中可以使用 `full` 来代指所有人**

**注意：方括号[]代表可省略，如 `a[,b,c...]` 代表至少一个，但是也可以有多个，多个之间用英文逗号隔开**

示例：

```
Item0 12.31 a,b-c	# a和b拼单花了12.31买Item0，由c付的钱
Item1 80 a		    # a花了80买了Item1，由默认付钱用户付钱
```

# 结果查看

```
User0
	- User1 10.2
	+ User2 51.34
```

上面表示：

User0 要向 User1 支付 10.2 元

User0 应从 User2 收取 51.34 元

使用 `-d|--debug` 选项还会输出具体的Item
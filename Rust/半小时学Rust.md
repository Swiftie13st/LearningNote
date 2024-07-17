# A half-hour to learn Rust（中文版）

[A half-hour to learn Rust](https://fasterthanli.me/articles/a-half-hour-to-learn-rust)


为了提高编程语言的流畅度，人们必须阅读大量的编程语言。

但如果你不知道它的含义，你怎么能读懂它呢？

在本文中，我不会只关注一两个概念，而是会尝试尽可能多地介绍 Rust 代码片段，并解释它们包含的关键字和符号的含义。

预备，开始！

## 变量绑定

### 关键词let

let引入变量绑定：

```Rust
let x; // declare "x"
x = 42; // assign 42 to "x"
```


这也可以写成一行：

```Rust
let x = 42;
```

### 类型注解

您可以使用 明确指定变量的类型:，这是一个类型注释：

```rust
let x: i32; // `i32` is a signed 32-bit integer
x = 42;

// there's i8, i16, i32, i64, i128
//    also u8, u16, u32, u64, u128 for unsigned
```

这也可以写成一行：

```rust
let x: i32 = 42;
## # 未初始化的变量

如果您声明一个名称并在稍后初始化它，则编译器将阻止您在初始化之前使用它。

```rust
let x;
foobar(x); // error: borrow of possibly-uninitialized variable: `x`
x = 42;
```

但是，这样做是完全没问题的：

```rust
let x;
x = 42;
foobar(x); // the type of `x` will be inferred from here
```

### 丢弃值

下划线_是一个特殊名称 - 或者更确切地说，是“缺少名称”。它的基本含义是丢弃某些东西：

```rust
// this does *nothing* because 42 is a constant
let _ = 42;

// this calls `get_thing` but throws away its result
let _ = get_thing();
```

以下划线开头的名称是常规名称，只是编译器不会警告它们未使用：

```rust
// we may use `_x` eventually, but our code is a work-in-progress
// and we just wanted to get rid of a compiler warning for now.
let _x = 42;
```

### 隐藏绑定

可以引入具有相同名称的单独绑定 - 您可以隐藏 变量绑定：

```rust
let x = 13;
let x = x + 3;
// using `x` after that line only refers to the second `x`,
//
// although the first `x` still exists (it'll be dropped
// when going out of scope), you can no longer refer to it.
```

## 元组

Rust 有元组，你可以将其视为“不同类型的值的固定长度集合”。

```rust
let pair = ('a', 17);
pair.0; // this is 'a'
pair.1; // this is 17
```

如果我们确实想注释的类型`pair`，我们可以写：

```rust
let pair: (char, i32) = ('a', 17);
```

### 解构元组

在执行分配时可以对元组进行解构，这意味着它们会被分解为各个字段：

```rust
let (some_char, some_int) = ('a', 17);
// now, `some_char` is 'a', and `some_int` is 17
```

当函数返回元组时这尤其有用：

```rust
let (left, right) = slice.split_at(middle);
```

当然，在解构元组的时候，_可以用来丢弃其中的一部分：

```rust
let (_, right) = slice.split_at(middle);
```

## 声明

分号标志着语句的结束：

```rust
let x = 3;
let y = 5;
let z = y + x;
```

这意味着语句可以跨越多行：

```rust
let x = vec![1, 2, 3, 4, 5, 6, 7, 8]
    .iter()
    .map(|x| x + 3)
    .fold(0, |x, y| x + y);
```

（稍后我们将解释它们的实际含义）。

## 函数

`fn`声明一个函数。

这是一个 `void` 函数：

```rust
fn greet() {
    println!("Hi there!");
}
```

下面是一个返回 32 位有符号整数的函数。箭头表示其返回类型：

```rust
fn fair_dice_roll() -> i32 {
    4
}
```

## 区块{}

一对括号声明一个块，它有自己的范围：

```rust
// This prints "in", then "out"
fn main() {
    let x = "out";
    {
        // this is a different `x`
        let x = "in";
        println!("{}", x);
    }
    println!("{}", x);
}
```

### `{}`是表达式

块也是表达式，这意味着它们会计算出一个值。

```rust
// this:
let x = 42;

// is equivalent to this:
let x = { 42 };
```

在一个块内，可以有多个语句：

```rust
let x = {
    let y = 1; // first statement
    let z = 2; // second statement
    y + z // this is the *tail* - what the whole block will evaluate to
};
```

### 隐式返回

这就是为什么“省略函数末尾的分号”与返回相同，即这些是等效的：

```rust
fn fair_dice_roll() -> i32 {
    return 4;
}

fn fair_dice_roll() -> i32 {
    4
}
```

### 一切都是一种表达

`if`条件语句也是表达式：

```rust
fn fair_dice_roll() -> i32 {
    if feeling_lucky {
        6
    } else {
        4
    }
}
```

`match`也是一个表达式：

```rust
fn fair_dice_roll() -> i32 {
    match feeling_lucky {
        true => 6,
        false => 4,
    }
}
```

## 字段访问和方法调用

点`.`通常用于访问某个值的字段：

```rust
let a = (10, 20);
a.0; // this is 10

let amos = get_some_struct();
amos.nickname; // this is "fasterthanlime"
```

或者对某个值调用一个方法：

```rust
let
 nick = "fasterthanlime";
nick.len(); // this is 14
```

## 模块 `use`语法

双冒号`::`类似，但它对命名空间进行操作。

在这个例子中，std是一个包（〜一个库），cmp是一个模块 （〜一个源文件），并且min是一个函数：

```rust
let least = std::cmp::min(3, 8); // this is 3
```

`use`指令可用于从其他命名空间“引入范围”名称：

```rust
use std::cmp::min;

let least = min(7, 1); // this is 1
```

在`use`指令中，花括号还有另一层含义：它们是“通配符”。如果我们想同时导入`min`和`max`，我们可以执行以下操作之一：

```rust
// this works:
use std::cmp::min;
use std::cmp::max;

// this also works:
use std::cmp::{min, max};

// this also works!
use std::{cmp::min, cmp::max};
```

通配符 ( `*`) 可让你从命名空间导入每个符号：

```rust
// this brings `min` and `max` in scope, and many other things
use std::cmp::*;
```

### 类型也是命名空间

类型也是命名空间，方法可以像常规函数一样调用：

```rust
let x = "amos".len(); // this is 4
let x = str::len("amos"); // this is also 4
```

### 标准库的预设功能集

str是一种原始类型，但许多非原始类型默认也在范围内。

```rust
// `Vec` is a regular struct, not a primitive type
let v = Vec::new();

// this is exactly the same code, but with the *full* path to `Vec`
let v = std::vec::Vec::new();
```

这是有效的，因为 Rust 在每个模块的开头插入了这段代码：

```rust
use std::prelude::v1::*;
```

（这反过来又重新导出了很多符号，如`Vec`，`String`，`Option`和`Result`)。

## 结构体

结构体用`struct`关键字声明：

```rust
struct Vec2 {
    x: f64, // 64-bit floating point, aka "double precision"
    y: f64,
}
```

它们可以使用结构文字进行初始化：

```rust
let v1 = Vec2 { x: 1.0, y: 3.0 };
let v2 = Vec2 { y: 2.0, x: 4.0 };
// the order does not matter, only the names do
```

### 结构体更新语法、

有一个快捷方式可以从另一个结构初始化其余字段：

```rust
let v3 = Vec2 {
    x: 14.0,
    ..v2
};
```

这被称为“结构更新语法（struct update syntax）”，只能发生在最后的位置，并且后面不能跟逗号。

请注意，其余字段可以表示所有字段：

```rust
let v4 = Vec2 { ..v3 };
```

### 解构结构体

结构体和元组一样，可以被解构。

就像这是一个有效的`let`语句：

```rust
let (left, right) = slice.split_at(middle);
```

这是这样的：

```rust
let v = Vec2 { x: 3.0, y: 6.0 };
let Vec2 { x, y } = v;
// `x` is now 3.0, `y` is now `6.0`
```

和这个：

```rust
let Vec2 { x, .. } = v;
// this throws away `v.y`
```

## 模式与解构

### if let

`let`模式可用作条件`if`：

```rust
struct Number {
    odd: bool,
    value: i32,
}

fn main() {
    let one = Number { odd: true, value: 1 };
    let two = Number { odd: false, value: 2 };
    print_number(one);
    print_number(two);
}

fn print_number(n: Number) {
    if let Number { odd: true, value } = n {
        println!("Odd number: {}", value);
    } else if let Number { odd: false, value } = n {
        println!("Even number: {}", value);
    }
}

// this prints:
// Odd number: 1
// Even number: 2
```

### match

match手臂也是图案，就像if let：

```rust
fn print_number(n: Number) {
    match n {
        Number { odd: true, value } => println!("Odd number: {}", value),
        Number { odd: false, value } => println!("Even number: {}", value),
    }
}

// this prints the same as before
```

### 完全匹配

`match`必须是详尽的：至少有一个臂需要匹配。

```rust
fn print_number(n: Number) {
    match n {
        Number { value: 1, .. } => println!("One"),
        Number { value: 2, .. } => println!("Two"),
        Number { value, .. } => println!("{}", value),
        // if that last arm didn't exist, we would get a compile-time error
    }
}
```

如果这很难，_可以用作“万能”模式：

```rust
fn print_number(n: Number) {
    match n.value {
        1 => println!("One"),
        2 => println!("Two"),
        _ => println!("{}", n.value),
    }
}
```

## 方法

您可以在自己的类型上声明方法：

```rust
struct Number {
    odd: bool,
    value: i32,
}

impl Number {
    fn is_strictly_positive(self) -> bool {
        self.value > 0
    }
}
```

和平常一样使用它们：

```rust
fn main() {
    let minus_two = Number {
        odd: false,
        value: -2,
    };
    println!("positive? {}", minus_two.is_strictly_positive());
    // this prints "positive? false"
}
```

## 不变性

变量绑定默认是不可变的，这意味着它们的内部不能被改变：

```rust
fn main() {
    let n = Number {
        odd: true,
        value: 17,
    };
    n.odd = false; // error: cannot assign to `n.odd`,
                   // as `n` is not declared to be mutable
}
```

而且它们不能被分配：

```rust
fn main() {
    let n = Number {
        odd: true,
        value: 17,
    };
    n = Number {
        odd: false,
        value: 22,
    }; // error: cannot assign twice to immutable variable `n`
}
```

mut使变量绑定可变：

```rust
fn main() {
    let mut n = Number {
        odd: true,
        value: 17,
    }
    n.value = 19; // all good
}
```

## trait

`trait`是多种类型可以具有的共同点：

```rust
trait Signed {
    fn is_strictly_negative(self) -> bool;
}
```

### 孤儿规则

您可以实现：

- 你的某个特征适合任何人的类型
- 任何人的特质都属于你的类型之一
- 但不是外来类型的外来特征

这些被称为“孤儿规则”。

以下是我们类型的特征的实现：

```rust
impl Signed for Number {
    fn is_strictly_negative(self) -> bool {
        self.value < 0
    }
}

fn main() {
    let n = Number { odd: false, value: -44 };
    println!("{}", n.is_strictly_negative()); // prints "true"
}
```

我们对外部类型（甚至是原始类型）的特征：

```rust
impl Signed for i32 {
    fn is_strictly_negative(self) -> bool {
        self < 0
    }
}

fn main() {
    let n: i32 = -44;
    println!("{}", n.is_strictly_negative()); // prints "true"
}
```

我们类型的一个外来特征：

```rust
// the `Neg` trait is used to overload `-`, the
// unary minus operator.
impl std::ops::Neg for Number {
    type Output = Number;

    fn neg(self) -> Number {
        Number {
            value: -self.value,
            odd: self.odd,
        }        
    }
}

fn main() {
    let n = Number { odd: true, value: 987 };
    let m = -n; // this is only possible because we implemented `Neg`
    println!("{}", m.value); // prints "-987"
}
```

### Self​

一个`impl`块总是针对一种类型，因此，在该块内，Self意味着该类型：

```rust
impl std::ops::Neg for Number {
    type Output = Self;

    fn neg(self) -> Self {
        Self {
            value: -self.value,
            odd: self.odd,
        }        
    }
}
```

### 标记trait

一些特征是标记- 它们并不表示某种类型实现了某些方法，而是表示可以用某种类型完成某些事情。

例如，i32实现特征Copy（简而言之，i32 是 Copy），因此这有效：

```rust
fn main() {
    let a: i32 = 15;
    let b = a; // `a` is copied
    let c = a; // `a` is copied again
}
```

这也有效：

```rust
fn print_i32(x: i32) {
    println!("x = {}", x);
}

fn main() {
    let a: i32 = 15;
    print_i32(a); // `a` is copied
    print_i32(a); // `a` is copied again
}
```

但Number结构不是Copy，因此这不起作用：

```rust
fn main() {
    let n = Number { odd: true, value: 51 };
    let m = n; // `n` is moved into `m`
    let o = n; // error: use of moved value: `n`
}
```

这也不行：

```rust
fn print_number(n: Number) {
    println!("{} number {}", if n.odd { "odd" } else { "even" }, n.value);
}

fn main() {
    let n = Number { odd: true, value: 51 };
    print_number(n); // `n` is moved
    print_number(n); // error: use of moved value: `n`
}
```

但如果print_number采用不可变引用，它就可以工作：

```rust
fn print_number(n: &Number) {
    println!("{} number {}", if n.odd { "odd" } else { "even" }, n.value);
}

fn main() {
    let n = Number { odd: true, value: 51 };
    print_number(&n); // `n` is borrowed for the time of the call
    print_number(&n); // `n` is borrowed again
}
```

如果函数采用可变引用，它也会起作用 - 但前提是我们的变量绑定也是 `mut`。

```rust
fn invert(n: &mut Number) {
    n.value = -n.value;
}

fn print_number(n: &Number) {
    println!("{} number {}", if n.odd { "odd" } else { "even" }, n.value);
}

fn main() {
    // this time, `n` is mutable
    let mut n = Number { odd: true, value: 51 };
    print_number(&n);
    invert(&mut n); // `n is borrowed mutably - everything is explicit
    print_number(&n);
}
```

### 特征方法接收器

特征方法也可以self通过引用或可变引用来获取：

```rust
impl std::clone::Clone for Number {
    fn clone(&self) -> Self {
        Self { ..*self }
    }
}
```

调用特征方法时，接收者被隐式借用：

```rust
fn main() {
    let n = Number { odd: true, value: 51 };
    let mut m = n.clone();
    m.value += 100;
    
    print_number(&n);
    print_number(&m);
}
```

为了强调这一点：这些是等效的：

```rust
let m = n.clone();

let m = std::clone::Clone::clone(&n);
```

标记特征如下Copy没有方法：

```rust
// note: `Copy` requires that `Clone` is implemented too
impl std::clone::Clone for Number {
    fn clone(&self) -> Self {
        Self { ..*self }
    }
}

impl std::marker::Copy for Number {}
```

现在，Clone仍然可以使用：

```rust
fn main() {
    let n = Number { odd: true, value: 51 };
    let m = n.clone();
    let o = n.clone();
}
```

但`Number`值将不再被移动：

```rust
fn main() {
    let n = Number { odd: true, value: 51 };
    let m = n; // `m` is a copy of `n`
    let o = n; // same. `n` is neither moved nor borrowed.
}
```

### 派生trait

有些特征非常常见，可以使用以下`derive`属性自动实现：

```rust
#[derive(Clone, Copy)]
struct Number {
    odd: bool,
    value: i32,
}

// this expands to `impl Clone for Number` and `impl Copy for Number` blocks.
```

## 泛型

### 泛型函数

函数可以是通用的：

```rust
fn foobar<T>(arg: T) {
    // do something with `arg`
}
```

它们可以有多个类型参数，然后可以在函数的声明和函数体中使用，而不是具体类型：

```rust
fn foobar<L, R>(left: L, right: R) {
    // do something with `left` and `right`
}
```

### 类型参数约束（特征边界）

类型参数通常有约束，因此您实际上可以对它们做一些事情。

最简单的约束只是特征名称：

```rust
fn print<T: Display>(value: T) {
    println!("value = {}", value);
}

fn print<T: Debug>(value: T) {
    println!("value = {:?}", value);
}
```

类型参数约束有更长的语法：

```rust
fn print<T>(value: T)
where
    T: Display,
{
    println!("value = {}", value);
}
```

约束可能更复杂：它们可能需要类型参数来实现多个特征：

```rust
use std::fmt::Debug;

fn compare<T>(left: T, right: T)
where
    T: Debug + PartialEq,
{
    println!("{:?} {} {:?}", left, if left == right { "==" } else { "!=" }, right);
}

fn main() {
    compare("tea", "coffee");
    // prints: "tea" != "coffee"
}
```

### 单态化

泛型函数可以被视为命名空间，包含无数具有不同具体类型的函数。

与crates、modules和types一样，泛型函数可以使用以下方式“探索”（导航？）`::`

```rust
fn main() {
    use std::any::type_name;
    println!("{}", type_name::<i32>()); // prints "i32"
    println!("{}", type_name::<(f64, char)>()); // prints "(f64, char)"
}
```

这被亲切地称为turbofish 语法，因为 `::<>`它看起来像一条鱼。

### 泛型结构体

结构体也可以是泛型的：

```rust
struct Pair<T> {
    a: T,
    b: T,
}

fn print_type_name<T>(_val: &T) {
    println!("{}", std::any::type_name::<T>());
}

fn main() {
    let p1 = Pair { a: 3, b: 9 };
    let p2 = Pair { a: true, b: false };
    print_type_name(&p1); // prints "Pair<i32>"
    print_type_name(&p2); // prints "Pair<bool>"
}
```

例子：`Vec``
标准库类型`Vec`（〜堆分配数组）是通用的：

```rust
fn main() {
    let mut v1 = Vec::new();
    v1.push(1);
    let mut v2 = Vec::new();
    v2.push(false);
    print_type_name(&v1); // prints "Vec<i32>"
    print_type_name(&v2); // prints "Vec<bool>"
}
```

说到`Vec`，它带有一个提供或多或少“vec字面量”的宏：

```rust
fn main() {
    let v1 = vec![1, 2, 3];
    let v2 = vec![true, false, true];
    print_type_name(&v1); // prints "Vec<i32>"
    print_type_name(&v2); // prints "Vec<bool>"
}
```

## 宏

`name!()`，`name![]`或`name!{}`调用宏。宏只是扩展为常规代码。

其实`println`就是一个宏：

```rust
fn main() {
    println!("{}", "Hello there!");
}
```

这扩展为具有与以下内容相同的效果：

```rust
fn main() {
    use std::io::{self, Write};
    io::stdout().lock().write_all(b"Hello there!\n").unwrap();
}
```

## 宏panic!​

`panic`也是一个宏。它会猛烈地停止执行并显示错误消息，如果启用，还会显示错误的文件名/行号：

```rust
fn main() {
    panic!("This panics");
}
// output: thread 'main' panicked at 'This panics', src/main.rs:3:5
```

## 引发panic的函数

有些方法也会引起panic。例如，`Option`类型可以包含某些内容，也可以不包含任何内容。如果`.unwrap()`调用该方法，而该方法不包含任何内容，则会引起panic：

```rust
fn main() {
    let o1: Option<i32> = Some(128);
    o1.unwrap(); // this is fine

    let o2: Option<i32> = None;
    o2.unwrap(); // this panics!
}

// output: thread 'main' panicked at 'called `Option::unwrap()` on a `None` value', src/libcore/option.rs:378:21
```

## enum 枚举（总和类型）

`Option`不是一个结构体——它是一个具有两个变量的`enum`。

```rust
enum Option<T> {
    None,
    Some(T),
}

impl<T> Option<T> {
    fn unwrap(self) -> T {
        // enums variants can be used in patterns:
        match self {
            Self::Some(t) => t,
            Self::None => panic!(".unwrap() called on a None option"),
        }
    }
}

use self::Option::{None, Some};

fn main() {
    let o1: Option<i32> = Some(128);
    o1.unwrap(); // this is fine

    let o2: Option<i32> = None;
    o2.unwrap(); // this panics!
}

// output: thread 'main' panicked at '.unwrap() called on a None option', src/main.rs:11:27
```

`Result`也是一个enum，它可以包含某些内容，也可以包含错误：

```rust
enum Result<T, E> {
    Ok(T),
    Err(E),
}
```

当解包时如果出现错误也会引起panic。

## 生命周期

变量绑定有一个“生命周期”：

```rust
fn main() {
    // `x` doesn't exist yet
    {
        let x = 42; // `x` starts existing
        println!("x = {}", x);
        // `x` stops existing
    }
    // `x` no longer exists
}
```

类似地，引用也有生命周期：

```rust
fn main() {
    // `x` doesn't exist yet
    {
        let x = 42; // `x` starts existing
        let x_ref = &x; // `x_ref` starts existing - it borrows `x`
        println!("x_ref = {}", x_ref);
        // `x_ref` stops existing
        // `x` stops existing
    }
    // `x` no longer exists
}
```

引用的生命周期不能超过它所借用的变量绑定的生命周期：

```rust
fn main() {
    let x_ref = {
        let x = 42;
        &x
    };
    println!("x_ref = {}", x_ref);
    // error: `x` does not live long enough
}
```

### 借用规则

（一个或多个不可变借用 XOR 一个可变借用）

变量绑定可以被不可变地借用多次：

```rust
fn main() {
    let x = 42;
    let x_ref1 = &x;
    let x_ref2 = &x;
    let x_ref3 = &x;
    println!("{} {} {}", x_ref1, x_ref2, x_ref3);
}
```

借用后，变量绑定不能发生变异：

```rust
fn main() {
    let mut x = 42;
    let x_ref = &x;
    x = 13;
    println!("x_ref = {}", x_ref);
    // error: cannot assign to `x` because it is borrowed
}
```

虽然变量是不可变借用的，但它不能是可变借用的：

```rust
fn main() {
    let mut x = 42;
    let x_ref1 = &x;
    let x_ref2 = &mut x;
    // error: cannot borrow `x` as mutable because it is also borrowed as immutable
    println!("x_ref1 = {}", x_ref1);
}
```

### 函数在生命周期内通用

函数参数中的引用也有生命周期：

```rust
fn print(x: &i32) {
    // `x` is borrowed (from the outside) for the
    // entire time this function is called.
}
```

具有引用参数的函数可以通过具有不同生命周期的借用来调用，因此：

- 所有接受引用的函数都是通用的
- 生命周期是通用参数

Lifetimes 的名称以单引号开头`'`：

```rust
// elided (non-named) lifetimes:
fn print(x: &i32) {}

// named lifetimes:
fn print<'a>(x: &'a i32) {}
```

这允许返回其生命周期取决于参数生命周期的引用：

```rust
struct Number {
    value: i32,
}

fn number_value<'a>(num: &'a Number) -> &'a i32 {
    &num.value
}

fn main() {
    let n = Number { value: 47 };
    let v = number_value(&n);
    // `v` borrows `n` (immutably), thus: `v` cannot outlive `n`.
    // While `v` exists, `n` cannot be mutably borrowed, mutated, moved, etc.
}
```

### 省略生命周期

当只有一个输入生命周期时，它不需要被命名，并且所有事物都有相同的生命周期，因此下面两个函数是等效的：

```rust
fn number_value<'a>(num: &'a Number) -> &'a i32 {
    &num.value
}

fn number_value(num: &Number) -> &i32 {
    &num.value
}
```

### struct在生命周期内通用

结构体在整个生命周期中也可以是通用的，这使得它们可以保存引用：

```rust
struct NumRef<'a> {
    x: &'a i32,
}

fn main() {
    let x: i32 = 99;
    let x_ref = NumRef { x: &x };
    // `x_ref` cannot outlive `x`, etc.
}
```

相同的代码，但有一个附加函数：

```rust
struct NumRef<'a> {
    x: &'a i32,
}

fn as_num_ref<'a>(x: &'a i32) -> NumRef<'a> {
    NumRef { x: &x }
}

fn main() {
    let x: i32 = 99;
    let x_ref = as_num_ref(&x);
    // `x_ref` cannot outlive `x`, etc.
}
```

相同的代码，但具有“省略”的生命周期：

```rust
struct NumRef<'a> {
    x: &'a i32,
}

fn as_num_ref(x: &i32) -> NumRef<'_> {
    NumRef { x: &x }
}

fn main() {
    let x: i32 = 99;
    let x_ref = as_num_ref(&x);
    // `x_ref` cannot outlive `x`, etc.
}
```

### 生命周期内的通用实现

impl块在整个生命周期中也可以是通用的：

```rust
impl<'a> NumRef<'a> {
    fn as_i32_ref(&'a self) -> &'a i32 {
        self.x
    }
}

fn main() {
    let x: i32 = 99;
    let x_num_ref = NumRef { x: &x };
    let x_i32_ref = x_num_ref.as_i32_ref();
    // neither ref can outlive `x`
}
```

但你也可以在那里进行省略：

```rust
impl<'a> NumRef<'a> {
    fn as_i32_ref(&self) -> &i32 {
        self.x
    }
}
```

如果你从来不需要这个名字，你可以更进一步地省略它：

```rust
impl NumRef<'_> {
    fn as_i32_ref(&self) -> &i32 {
        self.x
    }
}
```

## `'static​`生命周期

有一个特殊的生命周期，名为`'static`，它在整个程序的生命周期内有效。

字符串文字是`'static`：

```rust
struct Person {
    name: &'static str,
}

fn main() {
    let p = Person {
        name: "fasterthanlime",
    };
}
```

但是对 a 的引用`String`并不是静态的：

```rust
struct Person {
    name: &'static str,
}

fn main() {
    let name = format!("fasterthan{}", "lime");
    let p = Person { name: &name };
    // error: `name` does not live long enough
}
```

在最后一个例子中，局部变量`name`不是`&'static str`，而是 `String`。它是动态分配的，并且将被释放。它的生命周期小于整个程序（即使它恰好在`main`函数中）。

要将非字符串存储`'static`在`Person`中，需要执行以下操作之一：

A）在一生中都是通用的：

```rust
struct Person<'a> {
    name: &'a str,
}

fn main() {
    let name = format!("fasterthan{}", "lime");
    let p = Person { name: &name };
    // `p` cannot outlive `name`
}
```

或者

B）取得字符串的所有权

```rust
struct Person {
    name: String,
}

fn main() {
    let name = format!("fasterthan{}", "lime");
    let p = Person { name: name };
    // `name` was moved into `p`, their lifetimes are no longer tied.
}
```

## 结构体字面赋值简写

说到：在结构体字面量中，当将字段设置为同名的变量绑定时：

```rust
    let p = Person { name: name };
```

它可以缩短如下：

```rust
    let p = Person { name };
```

诸如`clippy`之类的工具会建议进行这些更改，并且如果您允许的话，甚至会以编程方式应用修复。

## 固有类型与引用类型

Owned types vs reference types

对于 Rust 中的许多类型，都有拥有和非拥有的变体：

- 字符串：`String`是拥有的，`&str`是一个引用。
- 路径：`PathBuf`是拥有的，`&Path`是引用的。
- 集合：`Vec<T>`是拥有的，`&[T]`是引用的。

### 切片

Rust 有切片 - 它们是对多个连续元素的引用。

您可以借用向量的一个切片，例如：

```rust
fn main() {
    let v = vec![1, 2, 3, 4, 5];
    let v2 = &v[2..4];
    println!("v2 = {:?}", v2);
}

// output:
// v2 = [3, 4]
```

### 运算符重载

上面的代码并不神奇。索引运算符 (`foo[index]`) 被重载了`Index`和`IndexMut`特征。

语法`..`只是范围文字。范围只是标准库中定义的几个结构。

它们可以是开放式的，并且如果其最右边边界前面有`=`，则可以包含它。

```rust
fn main() {
    // 0 or greater
    println!("{:?}", (0..).contains(&100)); // true
    // strictly less than 20
    println!("{:?}", (..20).contains(&20)); // false
    // 20 or less than 20
    println!("{:?}", (..=20).contains(&20)); // true
    // only 3, 4, 5
    println!("{:?}", (3..6).contains(&4)); // true
}
```

### 借用规则和切片

借用规则适用于切片。

```rust
fn tail(s: &[u8]) -> &[u8] {
  &s[1..] 
}

fn main() {
    let x = &[1, 2, 3, 4, 5];
    let y = tail(x);
    println!("y = {:?}", y);
}
```

这与以下内容相同：

```rust
fn tail<'a>(s: &'a [u8]) -> &'a [u8] {
  &s[1..] 
}
```

这是合法的：

```rust
fn main() {
    let y = {
        let x = &[1, 2, 3, 4, 5];
        tail(x)
    };
    println!("y = {:?}", y);
}
```

...但仅仅因为它`[1, 2, 3, 4, 5]`是一个`'static`数组。

因此，这是非法的：

```rust
fn main() {
    let y = {
        let v = vec![1, 2, 3, 4, 5];
        tail(&v)
        // error: `v` does not live long enough
    };
    println!("y = {:?}", y);
}
```

...因为向量是堆分配的，并且它具有非`'static`生命周期。

### 字符串切片 (`&str`)

`&str`价值观实际上是切片。

```rust
fn file_ext(name: &str) -> Option<&str> {
    // this does not create a new string - it returns
    // a slice of the argument.
    name.split(".").last()
}

fn main() {
    let name = "Read me. Or don't.txt";
    if let Some(ext) = file_ext(name) {
        println!("file extension: {}", ext);
    } else {
        println!("no file extension");
    }
}
```

...因此借用规则也适用于此：

```rust
fn main() {
    let ext = {
        let name = String::from("Read me. Or don't.txt");
        file_ext(&name).unwrap_or("")
        // error: `name` does not live long enough
    };
    println!("extension: {:?}", ext);
}
```

### 可错函数 (`Result`) 

可能失败的函数通常会返回`Result`：

```rust
fn main() {
    let s = std::str::from_utf8(&[240, 159, 141, 137]);
    println!("{:?}", s);
    // prints: Ok("🍉")

    let s = std::str::from_utf8(&[195, 40]);
    println!("{:?}", s);
    // prints: Err(Utf8Error { valid_up_to: 0, error_len: Some(1) })
}
```

如果您想在发生故障时panic，您可以`.unwrap()`：

```rust
fn main() {
    let s = std::str::from_utf8(&[240, 159, 141, 137]).unwrap();
    println!("{:?}", s);
    // prints: "🍉"

    let s = std::str::from_utf8(&[195, 40]).unwrap();
    // prints: thread 'main' panicked at 'called `Result::unwrap()`
    // on an `Err` value: Utf8Error { valid_up_to: 0, error_len: Some(1) }',
    // src/libcore/result.rs:1165:5
}
```

或者对于自定义消息`.expect()`：

```rust
fn main() {
    let s = std::str::from_utf8(&[195, 40]).expect("valid utf-8");
    // prints: thread 'main' panicked at 'valid utf-8: Utf8Error
    // { valid_up_to: 0, error_len: Some(1) }', src/libcore/result.rs:1165:5
}
```

或者，您可以`match`：

```rust
fn main() {
    match std::str::from_utf8(&[240, 159, 141, 137]) {
        Ok(s) => println!("{}", s),
        Err(e) => panic!(e),
    }
    // prints 🍉
}
```

或者您可以`if let`：

```rust
fn main() {
    if let Ok(s) = std::str::from_utf8(&[240, 159, 141, 137]) {
        println!("{}", s);
    }
    // prints 🍉
}
```

或者你可以将错误传播：

```rust
fn main() -> Result<(), std::str::Utf8Error> {
    match std::str::from_utf8(&[240, 159, 141, 137]) {
        Ok(s) => println!("{}", s),
        Err(e) => return Err(e),
    }
    Ok(())
}
```

或者你可以用`?`简洁的方式来做：

```rust
fn main() -> Result<(), std::str::Utf8Error> {
    let s = std::str::from_utf8(&[240, 159, 141, 137])?;
    println!("{}", s);
    Ok(())
}
```

## 解除引用

运算`*`符可用于取消引用，但您不需要这样做来访问字段或调用方法：

```rust
struct Point {
    x: f64,
    y: f64,
}

fn main() {
    let p = Point { x: 1.0, y: 3.0 };
    let p_ref = &p;
    println!("({}, {})", p_ref.x, p_ref.y);
}

// prints `(1, 3)`
```

并且只有当类型为以下时才可以这样做`Copy`：

```rust
struct Point {
    x: f64,
    y: f64,
}

fn negate(p: Point) -> Point {
    Point {
        x: -p.x,
        y: -p.y,
    }
}

fn main() {
    let p = Point { x: 1.0, y: 3.0 };
    let p_ref = &p;
    negate(*p_ref);
    // error: cannot move out of `*p_ref` which is behind a shared reference
}
```

```rust
// now `Point` is `Copy`
#[derive(Clone, Copy)]
struct Point {
    x: f64,
    y: f64,
}

fn negate(p: Point) -> Point {
    Point {
        x: -p.x,
        y: -p.y,
    }
}

fn main() {
    let p = Point { x: 1.0, y: 3.0 };
    let p_ref = &p;
    negate(*p_ref); // ...and now this works
}
```

## 函数类型、闭包

闭包只是类型的函数`Fn`，`FnMut`或者`FnOnce`带有一些捕获的上下文。

它们的参数是一对管道 (`|`) 内以逗号分隔的名称列表。它们不需要花括号，除非您想要有多个语句。

```rust
fn for_each_planet<F>(f: F)
    where F: Fn(&'static str)
{
    f("Earth");
    f("Mars");
    f("Jupiter");
}
 
fn main() {
    for_each_planet(|planet| println!("Hello, {}", planet));
}

// prints:
// Hello, Earth
// Hello, Mars
// Hello, Jupiter
```

借用规则也适用于它们：

```rust
fn for_each_planet<F>(f: F)
    where F: Fn(&'static str)
{
    f("Earth");
    f("Mars");
    f("Jupiter");
}
 
fn main() {
    let greeting = String::from("Good to see you");
    for_each_planet(|planet| println!("{}, {}", greeting, planet));
    // our closure borrows `greeting`, so it cannot outlive it
}
```

例如，这样是行不通的：

```rust
fn for_each_planet<F>(f: F)
    where F: Fn(&'static str) + 'static // `F` must now have "'static" lifetime
{
    f("Earth");
    f("Mars");
    f("Jupiter");
}

fn main() {
    let greeting = String::from("Good to see you");
    for_each_planet(|planet| println!("{}, {}", greeting, planet));
    // error: closure may outlive the current function, but it borrows
    // `greeting`, which is owned by the current function
}
```

但这会：

```rust
fn main() {
    let greeting = String::from("You're doing great");
    for_each_planet(move |planet| println!("{}, {}", greeting, planet));
    // `greeting` is no longer borrowed, it is *moved* into
    // the closure.
}
```

### `FnMut` 和借用规则

`FnMut`需要可变借用才能调用，因此它每次只能被调用一次。

这是合法的：

```rust
fn foobar<F>(f: F)
    where F: Fn(i32) -> i32
{
    println!("{}", f(f(2))); 
}
 
fn main() {
    foobar(|x| x * 2);
}

// output: 8
```

这不是：

```rust
fn foobar<F>(mut f: F)
    where F: FnMut(i32) -> i32
{
    println!("{}", f(f(2))); 
    // error: cannot borrow `f` as mutable more than once at a time
}
 
fn main() {
    foobar(|x| x * 2);
}
```

这又是合法的：

```rust
fn foobar<F>(mut f: F)
    where F: FnMut(i32) -> i32
{
    let tmp = f(2);
    println!("{}", f(tmp)); 
}
 
fn main() {
    foobar(|x| x * 2);
}

// output: 8
```

`FnMut`存在是因为一些闭包可变地借用局部变量：

```rust
fn foobar<F>(mut f: F)
    where F: FnMut(i32) -> i32
{
    let tmp = f(2);
    println!("{}", f(tmp)); 
}
 
fn main() {
    let mut acc = 2;
    foobar(|x| {
        acc += 1;
        x * acc
    });
}

// output: 24
```

这些闭包不能传递给需要以下条件的函数`Fn`：

```rust
fn foobar<F>(f: F)
    where F: Fn(i32) -> i32
{
    println!("{}", f(f(2))); 
}
 
fn main() {
    let mut acc = 2;
    foobar(|x| {
        acc += 1;
        // error: cannot assign to `acc`, as it is a
        // captured variable in a `Fn` closure.
        // the compiler suggests "changing foobar
        // to accept closures that implement `FnMut`"
        x * acc
    });
}
```

`FnOnce`闭包只能被调用一次。它们之所以存在，是因为有些闭包会移出在捕获时已被移动的变量：

```rust
fn foobar<F>(f: F)
    where F: FnOnce() -> String
{
    println!("{}", f()); 
}
 
fn main() {
    let s = String::from("alright");
    foobar(move || s);
    // `s` was moved into our closure, and our
    // closures moves it to the caller by returning
    // it. Remember that `String` is not `Copy`.
}
```

这是自然强制的，因为FnOnce闭包需要被移动 才能被调用。

例如，这是非法的：

```rust
fn foobar<F>(f: F)
    where F: FnOnce() -> String
{
    println!("{}", f()); 
    println!("{}", f()); 
    // error: use of moved value: `f`
}
```

而且，如果你需要说服我们的闭包确实移动了`s`，这也是非法的：

```rust
fn main() {
    let s = String::from("alright");
    foobar(move || s);
    foobar(move || s);
    // use of moved value: `s`
}
```

但这没问题：

```rust
fn main() {
    let s = String::from("alright");
    foobar(|| s.clone());
    foobar(|| s.clone());
}
```

这是一个带有两个参数的闭包：

```rust
fn foobar<F>(x: i32, y: i32, is_greater: F)
    where F: Fn(i32, i32) -> bool
{
    let (greater, smaller) = if is_greater(x, y) {
        (x, y)
    } else {
        (y, x)
    };
    println!("{} is greater than {}", greater, smaller);
}
 
fn main() {
    foobar(32, 64, |x, y| x > y);
}
```

这是一个忽略两个参数的闭包：

```rust
fn main() {
    foobar(32, 64, |_, _| panic!("Comparing is futile!"));
}
```

这是一个稍微令人担忧的结局：

```rust
fn countdown<F>(count: usize, tick: F)
    where F: Fn(usize)
{
    for i in (1..=count).rev() {
        tick(i);
    }
}
 
fn main() {
    countdown(3, |i| println!("tick {}...", i));
}

// output:
// tick 3...
// tick 2...
// tick 1...
```

### 厕所闭包（toilet closure）

以下是厕所闭包的情况：

```rust
fn main() {
    countdown(3, |_| ());
}
```

`|_| ()`因为外形像厕所，所以叫这个名字。

## 循环、迭代器

任何可迭代的东西都可以在`for in`循环中使用。

我们刚刚看到了范围的运用，但它也可以与 `Vec`一起使用：

```rust
fn main() {
    for i in vec![52, 49, 21] {
        println!("I like the number {}", i);
    }
}
```

或者切片：

```rust
fn main() {
    for i in &[52, 49, 21] {
        println!("I like the number {}", i);
    }
}

// output:
// I like the number 52
// I like the number 49
// I like the number 21
```

或者一个实际的迭代器：

```rust
fn main() {
    // note: `&str` also has a `.bytes()` iterator.
    // Rust's `char` type is a "Unicode scalar value"
    for c in "rust".chars() {
        println!("Give me a {}", c);
    }
}

// output:
// Give me a r
// Give me a u
// Give me a s
// Give me a t
```

即使迭代器项被过滤、映射和展平：

```rust
fn main() {
    for c in "SuRPRISE INbOUND"
        .chars()
        .filter(|c| c.is_lowercase())
        .flat_map(|c| c.to_uppercase())
    {
        print!("{}", c);
    }
    println!();
}

// output: UB
```

## 返回闭包

您可以从函数返回一个闭包：

```rust
fn make_tester(answer: String) -> impl Fn(&str) -> bool {
    move |challenge| {
        challenge == answer
    }
}

fn main() {
    // you can use `.into()` to perform conversions
    // between various types, here `&'static str` and `String`
    let test = make_tester("hunter2".into());
    println!("{}", test("******"));
    println!("{}", test("hunter2"));
}
```

### 捕获到闭包中

您甚至可以将对某些函数参数的引用移动到它返回的闭包中：

```rust
fn make_tester<'a>(answer: &'a str) -> impl Fn(&str) -> bool + 'a {
    move |challenge| {
        challenge == answer
    }
}

fn main() {
    let test = make_tester("hunter2");
    println!("{}", test("*******"));
    println!("{}", test("hunter2"));
}

// output:
// false
// true
```

或者，省略生命周期：

```rust
fn make_tester(answer: &str) -> impl Fn(&str) -> bool + '_ {
    move |challenge| {
        challenge == answer
    }
}
```
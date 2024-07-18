# Trait

在 Rust 中，**trait** 是一种抽象的类型，它**定义了类型使用这个接口的行为**。
Rust 中的 trait 类似于其他编程语言中的接口或抽象类，但它们更加灵活和强大。

Rust 的 trait 是一种强大的抽象机制，它允许你定义行为的集合，并在多种类型上实现这些行为。Trait 对象和 trait 约束为 Rust 的类型系统增加了灵活性和表达力。

在开发复杂系统的时候，我们常常会强调接口和实现要分离。因为这是一种良好的设计习惯，它把调用者和实现者隔离开，双方只要按照接口开发，彼此就可以不受对方内部改动的影响。trait 就是这样。它可以把数据结构中的行为单独抽取出来，使其可以在多个类型之间共享；也可以作为约束，在泛型编程中，限制参数化类型必须符合它规定的行为。

## 定义 Trait

可以通过 `trait` 关键字定义一个 trait：

```rust
// Animal 是一个 trait，它有两个方法：make_sound 和 name。
trait Animal {
    fn make_sound(&self);
    fn name(&self) -> &str;
}
```

## 实现trait

任何类型都可以实现 trait，只要它定义了 trait 中声明的所有方法：

```rust
// Dog 结构体实现了 Animal trait。
struct Dog;

impl Animal for Dog {
    fn make_sound(&self) {
        println!("Woof!");
    }

    fn name(&self) -> &str {
        "Dog"
    }
}
```

## 默认实现

Trait 可以为某些方法提供默认实现，这样在实现该 trait 时，可以选择不实现这些方法，直接使用默认行为：

```rust
trait Animal {
    fn make_sound(&self);
    fn name(&self) -> &str {
        "Unknown"
    }
}

struct Dog;

impl Animal for Dog {
    fn make_sound(&self) {
        println!("Woof!");
    }
}
```

在这个例子中，`Animal trait` 为 `name` 方法提供了一个默认实现，返回 "Unknown"。`Dog` 结构体只实现了 `make_sound` 方法，而 `name` 方法则使用 trait 中的默认实现。

## 关联类型

trait 还可以定义关联类型，这些类型在 trait 的实现中会被具体化：

```rust
// Iterator trait 定义了一个关联类型 Item。
trait Iterator {
    type Item;

    fn next(&mut self) -> Option<Self::Item>;
}
```

## trait 继承

一个 trait 可以继承另一个 trait：

```rust
trait Domesticated {}

// Pet trait 继承了 Domesticated trait。
trait Pet: Domesticated {
    fn pet_name(&self) -> String;
}
```

## trait 对象

Rust 允许创建 trait 对象，这是一种动态分发的机制：

```rust 
trait Animal {
    fn make_sound(&self);
}

struct Dog;
struct Cat;

impl Animal for Dog {
    fn make_sound(&self) {
        println!("Woof!");
    }
}

impl Animal for Cat {
    fn make_sound(&self) {
        println!("Meow!");
    }
}

fn animal_sound(animal: &dyn Animal) {
    animal.make_sound();
}

fn main() {
    let dog = Dog;
    let cat = Cat;

    animal_sound(&dog);
    animal_sound(&cat);
}
```

`animal_sound` 函数接受一个 `&dyn Animal` 类型的参数，这意味着它可以接收任何实现了 Animal trait 的类型的引用。

## trait 约束

当定义一个泛型函数或类型时，可以使用 `trait` 约束来限制泛型参数的类型。这要求泛型参数的类型必须实现指定的 `trait`


### trait绑定语法

```rust
pub fn notify(item: &impl Summary) {
    println!("Breaking news! {}", item.summarize());
}
// 为下面形式的语法糖
pub fn notify<T: Summary>(item: &T) {
    println!("Breaking news! {}", item.summarize());
}
```
item我们指定关键字和特征名称，而不是参数的具体类型impl 。此参数接受实现指定特征的任何类型。在 的主体中notify，我们可以调用item 来自Summary特征的任何方法。

这种impl Trait语法很方便，在简单情况下可以使代码更简洁，而更完整的特征绑定语法可以在其他情况下表达更多的复杂性。例如，我们可以有两个实现的参数Summary。使用语法这样做impl Trait如下：

```rust
pub fn notify(item1: &impl Summary, item2: &impl Summary) {
```

impl Trait如果我们希望此函数允许item1和 item2具有不同的类型（只要两种类型都实现），则使用是合适的Summary。但是，如果我们想强制两个参数具有相同的类型，则必须使用特征绑定，如下所示：

```rust 
pub fn notify<T: Summary>(item1: &T, item2: &T) {
```

item1和item2参数传递的值的具体类型必须相同。


### 使用 `+` 指定多个trait约束

```rust 
pub fn notify(item: &(impl Summary + Display)) {

// 泛型
pub fn notify<T: Summary + Display>(item: &T) {
```

指定两个特征边界后，notify 的主体可以调用 summarize 并使用 {} 来格式化 item。

### 通过 where 关键字来指定 trait 约束：

使用太多的特征界限有其缺点。每个泛型都有自己的特征边界，因此具有多个泛型类型参数的函数可以在函数名称与其参数列表之间包含大量特征绑定信息，从而使函数签名难以阅读。
因此，Rust 有替代语法用于在函数签名后的 where 子句内指定特征边界。

```rust
// 使用 +
fn some_function<T: Display + Clone, U: Clone + Debug>(t: &T, u: &U) -> i32 {

// 使用where
fn some_function<T, U>(t: &T, u: &U) -> i32
where
    T: Display + Clone,
    U: Clone + Debug,
{
```

compare 函数接受两个类型为 T 的参数，并返回一个布尔值，指示它们是否相等。T 必须实现 PartialEq trait，这使得比较操作成为可能。
```rust 
fn compare<T>(a: T, b: T) -> bool
where
    T: PartialEq,
{
    a == b
}
```



## Trait 的泛型参数

除了关联类型，trait 还可以接受泛型参数，这使得 trait 的方法可以接受不同类型的参数。

```rust
trait Processor<T> {
    fn process(&self, item: T) -> T;
}

struct Identity;

impl Processor<i32> for Identity {
    fn process(&self, item: i32) -> i32 {
        item
    }
}

impl Processor<String> for Identity {
    fn process(&self, item: String) -> String {
        item
    }
}
```

Processor trait 接受一个泛型参数 T，并定义了一个 `process` 方法，该方法接受一个 `T` 类型的参数并返回相同类型的值。`Identity` 结构体为 `i32` 和 `String` 类型实现了 Processor trait。

可以在泛型类型定义中使用 trait 约束：

```rust
// largest 函数接受一个实现了 PartialOrd 和 Copy trait 的类型的切片。
fn largest<T>(list: &[T]) -> &T
where
    T: PartialOrd + Copy,
{
    let mut largest = list[0];
    for &item in list.iter() {
        if item > largest {
            largest = item;
        }
    }
    largest
}
```


## Trait 的自动实现

Rust 允许为 trait 提供自动实现，这通常用于实现一些通用的逻辑，如格式化输出。

```rust
use std::fmt;

trait Printable {
    fn print(&self);
}

impl<T: fmt::Display> Printable for T {
    fn print(&self) {
        println!("{}", self);
    }
}

struct Point {
    x: i32,
    y: i32,
}

impl fmt::Display for Point {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "({}, {})", self.x, self.y)
    }
}

fn main() {
    let p = Point { x: 1, y: 2 };
    p.print(); // 使用自动实现的 print 方法
}
```

在这个例子中，我们为实现了 `fmt::Display` 的类型提供了 `Printable trait` 的自动实现。Point 结构体实现了 `fmt::Display`，因此它自动获得了 `print` 方法。


## 与生命周期结合

可以将 trait 约束与生命周期结合使用：

```rust
// longest_prefix 函数接受两个类型为 T 的切片引用，并且这两个切片引用的生命周期都是 'a。T 必须实现 PartialEq trait。
fn longest_prefix<'a, T>(a: &'a [T], b: &'a [T]) -> &'a [T]
where
    T: PartialEq,
{
    // ...
}
```


## 返回实现trait的类型

在 Rust 中，当你想要从一个函数返回一个实现了特定 `trait` 的类型，但又不希望暴露实现类型的具体细节时，可以使用 `impl Trait` 作为返回类型。
这种方式被称为返回实现了 `trait` 的类型，它可以帮助创建更安全和更灵活的代码。

使用 `impl Trait` 作为返回类型的场景包括：

1. 隐藏实现细节：当你不希望调用者知道返回的具体类型时。
2. 提供通用接口：当你的函数需要返回多种可能的类型，但这些类型都实现了某个 `trait` 时。
3. 减少类型依赖：在库设计中，使用 `impl Trait` 可以减少编译时的依赖链，使得库更加灵活和可维护。

```rust
// 定义了一个 Summary trait，它有一个方法 summarize，返回一个 String 类型。
trait Summary {
    fn summarize(&self) -> String;
}

// 定义了一个 Tweet 结构体，并为它实现了 Summary trait。
struct Tweet {
    username: String,
    content: String,
    reply: bool,
    retweet: bool,
}

impl Summary for Tweet {
    fn summarize(&self) -> String {
        format!("{}: {}", self.username, self.content)
    }
}

// 定义了一个函数 returns_summarizable，它的返回类型是 impl Summary，这意味着它返回的是一个实现了 Summary 的类型。
fn returns_summarizable() -> impl Summary {
    // 返回了一个 Tweet 实例。
    Tweet {
        username: String::from("horse_ebooks"),
        content: String::from("of course, as you probably already know, people"),
        reply: false,
        retweet: false,
    }
}

fn main() {
    let tweet = returns_summarizable();
    println!("{}", tweet.summarize());
}
```
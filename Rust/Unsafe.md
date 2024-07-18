# Unsafe

静态分析本质上是保守的，为了内存安全，Rust 所做的这些规则往往是普适性的。
当编译器尝试确定代码是否支持保证时，拒绝一些有效程序比接受一些无效程序更好。尽管代码可能没问题，但如果 R​​ust 编译器没有足够的信息来确保其可信，它就会拒绝该代码。在这些情况下，可以使用`unsafe`告诉编译器可以执行该代码。
同样的，当 Rust 要访问其它语言比如 C/C++ 的库，因为它们并不满足 Rust 的安全性要求，这种跨语言的 FFI（Foreign Function Interface），也是 unsafe 的。


在 Rust 中，`unsafe` 是一个关键字，它用于表示代码块中的某些操作是不受 Rust 编译器安全检查保护的。使用 `unsafe` 代码块可以绕过 Rust 的一些安全保证，直接进行底层操作。这通常涉及到直接内存操作、调用外部 C 代码、或者进行其他可能违反 Rust 安全原则的操作。

>还有一大类使用 unsafe Rust 纯粹是为了性能。比如略过边界检查、使用未初始化内存等。
>这样的 `unsafe` 要尽量不用，除非通过 benchmark 发现用 `unsafe` 可以解决某些性能瓶颈，否则使用起来得不偿失。因为，在使用 `unsafe` 代码的时候，我们已经把 Rust 的内存安全性，降低到了和 C++ 同等的水平。



## unsafe 的使用场景

- **直接内存访问**：进行指针操作，如解引用裸指针（raw pointers）。
- **调用外部 C 代码**：通过 extern "C" 块调用 C 语言库。
- **实现某些特性**：如 Send、Sync、Drop 等特性可能需要 unsafe 代码。
- **优化性能**：在某些情况下，为了提高性能，可能需要使用 unsafe 代码来绕过 Rust 的一些安全检查。

### 解引用裸指针（raw pointers）：

裸指针是 `*const T` 和 `*mut T` 类型的指针，它们不受 Rust 安全检查的约束。解引用这些指针需要使用 `unsafe` 块。

```rust
let x: i32 = 42;
let r: *const i32 = &x;

unsafe {
    println!("r points to: {}", *r);
}
```

### 调用不安全函数或方法：

某些函数或方法被标记为 `unsafe`，因为它们可能会引发未定义行为。调用这些函数时需要在 `unsafe` 块中进行。

```rust
unsafe fn dangerous() {
    // 不安全操作
}

unsafe {
    dangerous();
}
```

### 访问或修改可变静态变量：

静态变量在 Rust 中是全局的，访问和修改它们可能会引发数据竞争。因此，操作静态变量需要使用 `unsafe`。

```rust
static mut COUNTER: u32 = 0;

unsafe {
    COUNTER += 1;
    println!("COUNTER: {}", COUNTER);
}
```

### 实现不安全的特质（traits）：

某些特质被标记为 `unsafe`，因为它们的实现可能会违反 Rust 的安全保证。例如，`Send` 和 `Sync` trait。

```rust
unsafe trait UnsafeTrait {
    // 方法定义
}

unsafe impl UnsafeTrait for SomeType {
    // 实现
}
```

### 调用外部函数接口（FFI）：

与其他编程语言（如 C）进行互操作时，通过 `FFI` 调用外部函数需要使用 `unsafe`。

```rust
extern "C" {
    fn abs(input: i32) -> i32;
}

unsafe {
    println!("Absolute value of -3 according to C: {}", abs(-3));
}
```

## unsafe 的风险

- **内存安全**：unsafe 代码可能会违反 Rust 的内存安全保证，导致程序崩溃或未定义行为。
- **数据竞争**：在多线程环境中，不当的 unsafe 代码可能会引起数据竞争，导致不可预测的结果。
- **难以调试**：unsafe 代码使得程序的调试变得更加困难，因为它们绕过了 Rust 的安全检查。

## unsafe 的最佳实践

- **最小化使用**：尽可能减少 unsafe 代码的使用，只在必要时使用。
- **明确安全保证**：在使用 unsafe 代码时，确保你完全理解代码的安全性，并提供必要的安全保证。
- **审计和测试**：对 unsafe 代码进行充分的审计和测试，确保它们在所有情况下都是安全的。
- **隔离 unsafe 代码**：将 unsafe 代码隔离到单独的函数或模块中，以减少它们对其他代码的影响。


## 示例

```rust 
fn main() {
    let mut num = 5;

    // 创建裸指针
    let r1 = &num as *const i32;
    let r2 = &mut num as *mut i32;

    // 使用 unsafe 块来解引用裸指针
    unsafe {
        println!("r1 points to: {}", *r1);
        *r2 = 10;
        println!("r2 points to: {}", *r2);
    }

    println!("num is now: {}", num);
}
```
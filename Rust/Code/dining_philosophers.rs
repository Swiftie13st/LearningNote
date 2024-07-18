// https://doc.rust-lang.org/1.4.0/book/dining-philosophers.html
use std::thread;
use std::sync::{Mutex, Arc};

// Philosopher 代表哲学家，包含名字和左右叉子的索引。
struct Philosopher {
    name: String,
    left: usize,
    right: usize,
}

impl Philosopher {
    // 构造函数，创建一个新的 Philosopher 实例
    fn new(name: &str, left: usize, right: usize) -> Philosopher {
        Philosopher {
            name: name.to_string(),
            left: left,
            right: right,
        }
    }

    // 吃饭，需要借用 Table
    fn eat(&self, table: &Table) {
        // 锁定左右两边的叉子
        let _left = table.forks[self.left].lock().unwrap();
        let _right = table.forks[self.right].lock().unwrap();

        println!("{} is eating.", self.name);

        thread::sleep_ms(1000);

        println!("{} is done eating.", self.name);
    }
}

// Table 代表餐桌，包含一个叉子的Vec，每个叉子用 Mutex<()> 包装，以实现互斥访问。
struct Table {
    forks: Vec<Mutex<()>>,
}

fn main() {
    // 创建一个餐桌，并用 Arc 包装，使其可以在多个线程间共享
    let table = Arc::new(Table { forks: vec![
        Mutex::new(()),
        Mutex::new(()),
        Mutex::new(()),
        Mutex::new(()),
        Mutex::new(()),
    ]});

    // 创建五个哲学家
    let philosophers = vec![
        Philosopher::new("Judith Butler", 0, 1),
        Philosopher::new("Gilles Deleuze", 1, 2),
        Philosopher::new("Karl Marx", 2, 3),
        Philosopher::new("Emma Goldman", 3, 4),
        Philosopher::new("Michel Foucault", 0, 4),
    ];

    // 存储线程句柄
    let handles: Vec<_> = philosophers.into_iter().map(|p| {
        // 克隆 Arc，增加引用计数
        let table = table.clone();
        // 创建一个新线程，在这个线程中调用哲学家的 eat 方法
        thread::spawn(move || {
            p.eat(&table);
        })
    }).collect();

    // 等待所有线程完成
    for h in handles {
        h.join().unwrap();
    }
}

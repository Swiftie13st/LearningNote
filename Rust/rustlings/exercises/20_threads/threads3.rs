use std::{sync::mpsc, thread, time::Duration};

struct Queue {
    length: u32,
    first_half: Vec<u32>,
    second_half: Vec<u32>,
}

impl Queue {
    fn new() -> Self {
        Self {
            length: 10,
            first_half: vec![1, 2, 3, 4, 5],
            second_half: vec![6, 7, 8, 9, 10],
        }
    }
}

fn send_tx(q: Queue, tx: mpsc::Sender<u32>) {
    // TODO: We want to send `tx` to both threads. But currently, it is moved
    // into the first thread. How could you solve this problem?
    let tx1 = tx.clone();
    thread::spawn(move || {
        for val in q.first_half {
            println!("Sending {val:?}");
            tx.send(val).unwrap();
            thread::sleep(Duration::from_millis(250));
        }
    });

    thread::spawn(move || {
        for val in q.second_half {
            println!("Sending {val:?}");
            tx1.send(val).unwrap();
            thread::sleep(Duration::from_millis(250));
        }
    });
}

fn main() {
    // You can optionally experiment here.
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn threads3() {
        let (tx, rx) = mpsc::channel();
        let queue = Queue::new();
        let queue_length = queue.length;

        send_tx(queue, tx);

        let mut total_received: u32 = 0;
        for received in rx {
            println!("Got: {received}");
            total_received += 1;
        }

        println!("Number of received values: {total_received}");
        assert_eq!(total_received, queue_length);
    }
}

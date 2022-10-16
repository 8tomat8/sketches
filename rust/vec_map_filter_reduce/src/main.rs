fn main() {
    let x1: Vec<i32> = Vec::from([1, 2, 3, 4, 5, 6]);

    let x_map: Vec<i32> = x1.iter().map(|&val| val * 3).collect::<Vec<i32>>();
    println!("x_map: {:?}", x_map);

    let x_filtered: Vec<i32> = x_map
        .iter()
        .filter(|&&val| val % 4 != 0)
        .rev()
        .cloned()
        .collect::<Vec<i32>>();

    println!("x_filtered: {:?}", x_filtered);

    let x_reduced = x_filtered.iter().cloned().reduce(|res, cur| {
        print!("res: {}, cur: {}", res, cur);
        let magic_num = cur * 42;
        if magic_num % 5 == 0 {
            println!("... BINGO!");
            return res + cur;
        }
        println!("... ({}) nope...", magic_num);
        res
    });
    println!("x_reduced: {:?}", x_reduced);

    println!("x1: {:?}", x1);
}

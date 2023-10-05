pub fn main() {
    println!("Hello, world! from Rust");
}

#[export_name="add"]
pub fn add(a: i32, b: i32) -> i32 {
    return a + b;
}
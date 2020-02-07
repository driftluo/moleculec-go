use test_codegen::generate;

fn main() {
    if std::env::args().len() != 3 {
        eprintln!("args error, must be 2");
        return;
    }
    let schema_path = std::env::args().nth(1).unwrap();
    let testset_path = std::env::args().nth(2).unwrap();

    generate(&schema_path, &testset_path)
}

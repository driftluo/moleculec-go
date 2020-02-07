use molecule_codegen::{
    ast::{self, HasName},
    Parser,
};
use std::{
    collections::HashMap,
    fs,
    io::{self, Read, Write},
    path::PathBuf,
    rc::Rc,
};

type Ast = HashMap<String, Rc<ast::TopDecl>>;

mod bytes;
mod generator;
mod types;

use generator::GenTest;

fn load_schema(path: &str) -> Ast {
    let ast = Parser::parse(&PathBuf::from(path));
    ast.decls()
        .iter()
        .map(|decl| (decl.name().to_owned(), Rc::clone(decl)))
        .collect::<HashMap<_, _>>()
}

fn load_testset(path: &str) -> types::All {
    let test_data = {
        let mut file = fs::File::open(path)
            .unwrap_or_else(|err| panic!("failed to open tests from {}: {}", path, err));
        let mut contents = String::new();
        file.read_to_string(&mut contents)
            .unwrap_or_else(|err| panic!("failed to load tests from {}: {}", path, err));
        contents
    };
    serde_yaml::from_str(&test_data).unwrap_or_else(|err| panic!("failed to parse tests: {}", err))
}

pub fn generate(scheme_path: &str, testset_path: &str) {
    let ast = load_schema(scheme_path);
    let testset = load_testset(testset_path);
    let name = PathBuf::from(testset_path)
        .file_name()
        .unwrap()
        .to_str()
        .unwrap()
        .split('.')
        .collect::<Vec<&str>>()[0]
        .to_owned();

    let test = if name == "default" {
        testset
            .iter()
            .fold(Vec::with_capacity(testset.len()), |mut tests, any| {
                if tests.is_empty() {
                    tests.push(
                        r#"
package types

import (
	"bytes"
	"encoding/hex"
	"testing"
)
                    "#
                        .to_owned(),
                    )
                }
                tests.push(default_gen(any));
                tests
            })
    } else {
        testset
            .iter()
            .fold(Vec::with_capacity(testset.len()), |mut tests, any| {
                if tests.is_empty() {
                    tests.push(
                        r#"
package types

import (
	"bytes"
	"encoding/hex"
	"testing"
)
                    "#
                        .to_owned(),
                    )
                }
                tests.push(any.gen_test(&ast, &name, tests.len()));
                tests
            })
    };

    let code = test.join("\n");

    let stdout = io::stdout();
    let mut stdout_handle = stdout.lock();
    stdout_handle.write_all(code.as_bytes()).unwrap();
    stdout_handle.flush().unwrap();
}

fn default_gen(typ: &types::Any) -> String {
    match typ {
        types::Any::Option_(inner) => format!(
            r#"
func Test{name}Default(t *testing.T) {{
	x := {name}Default()
	expected, _ := hex.DecodeString("{expected}")
	if bytes.Compare(x.AsSlice(), expected) != 0 {{
	    t.Error("type {name} default error: ", x.AsSlice(), expected)
	}}
    y, _ := {name}FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {{
	    t.Error("type {name} default error: ", y.AsSlice(), expected)
	}}
}}
            "#,
            name = inner.name,
            expected = inner.expected.as_inner()
        ),
        _ => "".to_owned(),
    }
}

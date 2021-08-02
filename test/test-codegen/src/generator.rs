use crate::{types, Ast};
use case::CaseExt;
use molecule_codegen::ast::{self, HasName};
use std::collections::HashMap;

pub(crate) trait GenTest {
    fn gen_test(&self, ast: &Ast, test_name: &str, id: usize) -> String;
    fn assert_impl(&self, name: &str) -> String {
        format!(
            r#"
	if bytes.Compare(x.AsSlice(), expected) != 0 {{
	    t.Error("type {name} error: ", x.AsSlice(), expected)
	}}
    y, _ := {name}FromSlice(expected, false)
	if bytes.Compare(y.AsSlice(), expected) != 0 {{
	    t.Error("type {name} error: ", y.AsSlice(), expected)
	}}"#,
            name = name
        )
    }
}

impl GenTest for types::Any {
    fn gen_test(&self, ast: &Ast, test_name: &str, id: usize) -> String {
        match self {
            Self::Union(inner) => inner.gen_test(ast, test_name, id),
            Self::Option_(inner) => inner.gen_test(ast, test_name, id),
            Self::Array(inner) => inner.gen_test(ast, test_name, id),
            Self::StructOrTable(inner) => inner.gen_test(ast, test_name, id),
            Self::Vector(inner) => inner.gen_test(ast, test_name, id),
        }
    }
}

impl GenTest for types::Option_ {
    fn gen_test(&self, ast: &Ast, test_name: &str, _id: usize) -> String {
        let name = self.name.as_str();
        let test_name = format!("Test{}{}", name, test_name.to_camel());
        let builder = if let Some(item) = &self.item {
            let decl = ast.get(name).unwrap();
            let item_type = if let ast::TopDecl::Option_(ref decl) = decl.as_ref() {
                decl.item().typ().name().to_camel()
            } else {
                panic!("Error: type for {} is incorrect", self.name);
            };

            format!(
                r#"item, _ := hex.DecodeString("{item_data}")
x := New{name}Builder().Set(*{item_type}FromSliceUnchecked(item)).Build()
            "#,
                name = name,
                item_type = item_type,
                item_data = item.as_inner()
            )
        } else {
            format!(
                r#"x := New{item_type}Builder().Build()
            "#,
                item_type = name
            )
        };
        let expected = format!(
            r#"
expected, _ := hex.DecodeString("{}")
        "#,
            self.expected.as_inner()
        );
        format!(
            r#"
func {test_name}(t *testing.T) {{
	{builder}
	{expected}
    {assert}
}}
        "#,
            test_name = test_name,
            builder = builder,
            expected = expected,
            assert = self.assert_impl(name)
        )
    }
}

impl GenTest for types::Union {
    fn gen_test(&self, _ast: &Ast, test_name: &str, id: usize) -> String {
        let name = self.name.as_str();
        let test_name = format!("Test{}{}{}", name, test_name.to_camel(), id);
        let builder = if let Some(item) = &self.item {
            let item_type = &item.typ;
            format!(
                r#"item, _ := hex.DecodeString("{item_data}")
x := New{name}Builder().Set({name}UnionFrom{item_type}(*{item_type}FromSliceUnchecked(item))).Build()
            "#,
                name = name,
                item_type = item_type.to_camel(),
                item_data = item.data.as_inner()
            )
        } else {
            format!(
                r#"x := New{item_type}Builder().Build()
            "#,
                item_type = name
            )
        };
        let expected = format!(
            r#"
expected, _ := hex.DecodeString("{}")
        "#,
            self.expected.as_inner()
        );
        format!(
            r#"
func {test_name}(t *testing.T) {{
	{builder}
	{expected}
    {assert}
}}
        "#,
            test_name = test_name,
            builder = builder,
            expected = expected,
            assert = self.assert_impl(name)
        )
    }
}

impl GenTest for types::Array {
    fn gen_test(&self, ast: &Ast, test_name: &str, id: usize) -> String {
        let name = self.name.as_str();
        let test_name = format!("Test{}{}{}", name, test_name.to_camel(), id);
        let decl = ast.get(name).unwrap();
        let item_type = if let ast::TopDecl::Array(ref decl) = decl.as_ref() {
            decl.item().typ().name().to_camel()
        } else {
            panic!("Error: type for {} is incorrect", name);
        };

        let builder = if self.data.is_empty() {
            format!(
                r#"x := New{name}Builder().Build()
            "#,
                name = name
            )
        } else {
            let mut set_list = Vec::with_capacity(self.data.len());
            for index in self.data.keys() {
                set_list.push(format!(
                    r#".Nth{index}(*{item_type}FromSliceUnchecked(item{index}))"#,
                    index = index,
                    item_type = item_type
                ))
            }
            format!(
                r#"x := New{name}Builder(){set_list}.Build()
            "#,
                name = name,
                set_list = set_list.join("")
            )
        };
        let mut items_parse = Vec::with_capacity(self.data.len());
        for (index, data) in self.data.iter() {
            items_parse.push(format!(
                r#"item{index}, _ := hex.DecodeString("{item_data}")"#,
                index = index,
                item_data = data.as_inner()
            ))
        }
        let expected = format!(
            r#"
expected, _ := hex.DecodeString("{}")
        "#,
            self.expected.as_inner()
        );
        format!(
            r#"
func {test_name}(t *testing.T) {{
    {items_parse}
	{builder}
	{expected}
	{assert}
}}
        "#,
            test_name = test_name,
            builder = builder,
            expected = expected,
            assert = self.assert_impl(name),
            items_parse = items_parse.join("\n")
        )
    }
}

impl GenTest for types::StructOrTable {
    fn gen_test(&self, ast: &Ast, test_name: &str, id: usize) -> String {
        let name = self.name.as_str();
        let test_name = format!("Test{}{}{}", name, test_name.to_camel(), id);
        let decl = ast.get(name).unwrap();
        let field_type_dict = if let ast::TopDecl::Struct(ref decl) = decl.as_ref() {
            decl.fields()
                .iter()
                .map(|field| (field.name(), field.typ().name()))
                .collect::<HashMap<_, _>>()
        } else if let ast::TopDecl::Table(ref decl) = decl.as_ref() {
            decl.fields()
                .iter()
                .map(|field| (field.name(), field.typ().name()))
                .collect::<HashMap<_, _>>()
        } else {
            panic!("Error: type for {} is incorrect", name);
        };
        let builder = if self.data.is_empty() {
            format!(
                r#"x := New{name}Builder().Build()
            "#,
                name = name
            )
        } else {
            let mut set_list = Vec::with_capacity(self.data.len());
            for field in self.data.keys() {
                set_list.push(format!(
                    r#".{field_1}(*{item_type}FromSliceUnchecked({field}))"#,
                    field_1 = field.to_camel(),
                    field = field,
                    item_type = field_type_dict.get(field.as_str()).unwrap().to_camel()
                ))
            }
            format!(
                r#"x := New{name}Builder(){set_list}.Build()
            "#,
                name = name,
                set_list = set_list.join("")
            )
        };

        let mut fields_parse = Vec::with_capacity(self.data.len());
        for (field, data) in self.data.iter() {
            fields_parse.push(format!(
                r#"{field}, _ := hex.DecodeString("{item_data}")"#,
                field = field,
                item_data = data.as_inner()
            ))
        }
        let expected = format!(
            r#"
expected, _ := hex.DecodeString("{}")
        "#,
            self.expected.as_inner()
        );
        format!(
            r#"
func {test_name}(t *testing.T) {{
    {fields_parse}
	{builder}
	{expected}
	{assert}
}}
        "#,
            test_name = test_name,
            builder = builder,
            expected = expected,
            assert = self.assert_impl(name),
            fields_parse = fields_parse.join("\n")
        )
    }
}

impl GenTest for types::Vector {
    fn gen_test(&self, ast: &Ast, test_name: &str, id: usize) -> String {
        let name = self.name.as_str();
        let test_name = format!("Test{}{}{}", name, test_name.to_camel(), id);
        let decl = ast.get(name).unwrap();
        let item_type = if let ast::TopDecl::FixVec(ref decl) = decl.as_ref() {
            decl.item().typ().name().to_camel()
        } else if let ast::TopDecl::DynVec(ref decl) = decl.as_ref() {
            decl.item().typ().name().to_camel()
        } else {
            panic!("Error: type for {} is incorrect", name);
        };
        let builder = if self.data.is_empty() {
            format!(
                r#"x := New{name}Builder().Build()
            "#,
                name = name
            )
        } else {
            let push_list = self
                .data
                .iter()
                .enumerate()
                .map(|(index, _)| {
                    format!(
                        r#".Push(*{item_type}FromSliceUnchecked(item{index}))"#,
                        item_type = item_type,
                        index = index
                    )
                })
                .collect::<Vec<String>>();
            format!(
                r#"x := New{name}Builder(){push_list}.Build()
                "#,
                name = name,
                push_list = push_list.join("")
            )
        };
        let items_parse = self
            .data
            .iter()
            .enumerate()
            .map(|(index, data)| {
                format!(
                    r#"item{index}, _ := hex.DecodeString("{item_data}")"#,
                    index = index,
                    item_data = data.as_inner()
                )
            })
            .collect::<Vec<String>>();
        let expected = format!(
            r#"
expected, _ := hex.DecodeString("{}")
        "#,
            self.expected.as_inner()
        );
        format!(
            r#"
func {test_name}(t *testing.T) {{
    {items_parse}
	{builder}
	{expected}
	{assert}

	if x.TotalSize() != uint(len(x.AsSlice())) {{
        t.Error("struct: {name}:\n data: ", x.AsSlice(), "\n partial read total_size: ", x.TotalSize(), ", actual: ", len(x.AsSlice()))
	}}
}}
        "#,
            test_name = test_name,
            builder = builder,
            expected = expected,
            assert = self.assert_impl(name),
            items_parse = items_parse.join("\n"),
            name = name,
        )
    }
}

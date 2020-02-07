use std::collections::BTreeMap;

use serde::Deserialize;

use crate::bytes::Bytes;

pub(crate) type All = Vec<Any>;

#[derive(Deserialize)]
#[serde(deny_unknown_fields, untagged)]
pub(crate) enum Any {
    Option_(Option_),
    Union(Union),
    Array(Array),
    StructOrTable(StructOrTable),
    Vector(Vector),
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
pub(crate) struct Option_ {
    pub name: String,
    pub item: Option<Bytes>,
    pub expected: Bytes,
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
pub(crate) struct Union {
    pub name: String,
    pub item: Option<Item>,
    pub expected: Bytes,
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
pub(crate) struct Array {
    pub name: String,
    #[serde(default)]
    pub data: BTreeMap<usize, Bytes>,
    pub expected: Bytes,
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
pub(crate) struct StructOrTable {
    pub name: String,
    #[serde(default)]
    pub data: BTreeMap<String, Bytes>,
    pub expected: Bytes,
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
pub(crate) struct Vector {
    pub name: String,
    #[serde(default)]
    pub data: Vec<Bytes>,
    pub expected: Bytes,
}

#[derive(Deserialize)]
#[serde(deny_unknown_fields)]
pub(crate) struct Item {
    #[serde(rename = "type")]
    pub typ: String,
    pub data: Bytes,
}

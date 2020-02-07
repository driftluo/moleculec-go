use std::fmt;

#[derive(Default)]
pub(crate) struct Bytes(String);

pub(crate) struct BytesVisitor;

impl<'b> serde::de::Visitor<'b> for BytesVisitor {
    type Value = Bytes;

    fn expecting(&self, formatter: &mut fmt::Formatter) -> fmt::Result {
        write!(formatter, "require a 0x-prefixed hexadecimal string")
    }

    fn visit_str<E>(self, v: &str) -> Result<Self::Value, E>
    where
        E: serde::de::Error,
    {
        if v.len() < 2 || &v.as_bytes()[0..2] != b"0x" {
            return Err(E::invalid_value(serde::de::Unexpected::Str(v), &self));
        }

        if v.len() == 2 {
            return Ok(Default::default());
        }

        let v_new = v[2..].replace("_", "").replace("/", "");
        if v_new.len() & 1 != 0 {
            return Err(E::invalid_value(serde::de::Unexpected::Str(v), &self));
        }

        Ok(Bytes(v_new))
    }

    fn visit_string<E>(self, v: String) -> Result<Self::Value, E>
    where
        E: serde::de::Error,
    {
        self.visit_str(&v)
    }
}

impl<'de> serde::Deserialize<'de> for Bytes {
    fn deserialize<D>(deserializer: D) -> Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        deserializer.deserialize_str(BytesVisitor)
    }
}

impl Bytes {
    pub(crate) fn as_inner(&self) -> &str {
        self.0.as_ref()
    }
}

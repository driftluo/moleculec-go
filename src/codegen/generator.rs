use super::builder::{impl_as_builder_for_struct_or_table, impl_as_builder_for_vector, GenBuilder};
use super::union::GenUnion;
use molecule_codegen::ast::{self, DefaultContent, HasName};

use case::CaseExt;
use std::io;

use core::mem::size_of;
use std::io::Write;

// Little Endian
pub type Number = u32;
// Size of Number
pub const NUMBER_SIZE: usize = size_of::<Number>();

pub(super) trait Generator: HasName + DefaultContent {
    fn generate<W: io::Write>(&self, writer: &mut W) -> io::Result<()>;
    fn common_generate<W: io::Write>(&self, writer: &mut W) -> io::Result<()> {
        let struct_name = self.name().to_camel();

        let define = format!(
            r#"
type {struct_name} struct {{
    inner []byte
}}
        "#,
            struct_name = struct_name
        );
        writeln!(writer, "{}", define)?;

        let impl_ = format!(
            r#"
func {struct_name}FromSliceUnchecked(slice []byte) *{struct_name} {{
    return &{struct_name}{{inner: slice}}
}}
func (s *{struct_name}) AsSlice() []byte {{
    return s.inner
}}
            "#,
            struct_name = struct_name
        );
        writeln!(writer, "{}", impl_)?;

        let default_content = self
            .default_content()
            .iter()
            .map(ToString::to_string)
            .collect::<Vec<String>>()
            .join(",");

        let default = format!(
            r#"
func {struct_name}Default() {struct_name} {{
    return *{struct_name}FromSliceUnchecked([]byte{{ {default_content} }})
}}
            "#,
            struct_name = struct_name,
            default_content = default_content
        );
        writeln!(writer, "{}", default)?;
        Ok(())
    }

    // mark this as false to keep the original behavior
    fn need_unpack() -> bool { false }
}

/// For generate native conversion codes
pub trait NativeGenerator : HasName {

    /// Whether this type needs to unpack
    fn need_unpack() -> bool;

    fn native_type_support_generate(&self, writer: &mut impl io::Write) -> io::Result<()>;
}

impl Generator for ast::Option_ {
    fn generate<W: io::Write>(&self, writer: &mut W) -> io::Result<()> {
        writeln!(writer, "{}", self.gen_builder())?;
        self.common_generate(writer)?;

        let struct_name = self.name().to_camel();
        let inner = self.item().typ().name().to_camel();

        let constructor = format!(
            r#"
func {struct_name}FromSlice(slice []byte, compatible bool) (*{struct_name}, error) {{
    if len(slice) == 0 {{
        return &{struct_name}{{inner: slice}}, nil
    }}

    _, err := {inner_type}FromSlice(slice, compatible)
    if err != nil {{
        return nil, err
    }}
    return &{struct_name}{{inner: slice}}, nil
}}
            "#,
            struct_name = struct_name,
            inner_type = inner
        );
        writeln!(writer, "{}", constructor)?;

        let impl_ = format!(
            r#"
func (s *{struct_name}) Into{inner_type}() (*{inner_type}, error) {{
	if s.IsNone() {{
		return nil, errors.New("No data")
	}}
	return {inner_type}FromSliceUnchecked(s.AsSlice()), nil
}}
func (s *{struct_name}) IsSome() bool {{
    return len(s.inner) != 0
}}
func (s *{struct_name}) IsNone() bool {{
    return len(s.inner) == 0
}}
func (s *{struct_name}) AsBuilder() {struct_name}Builder {{
    var ret = New{struct_name}Builder()
    if s.IsSome() {{
        ret.Set(*{inner_type}FromSliceUnchecked(s.AsSlice()))
    }}
    return *ret
}}
            "#,
            struct_name = struct_name,
            inner_type = inner
        );
        writeln!(writer, "{}", impl_)?;
        Ok(())
    }
}

impl NativeGenerator for ast::Option_ {
    fn need_unpack() -> bool {
        true
    }

    fn native_type_support_generate(&self, writer: &mut impl io::Write) -> io::Result<()> {
        Ok(())
    }
}

//impl Generator for ast::Primitive {
//    fn generate<W: io::Write>(&self, writer: &mut W) -> io::Result<()> {
//        self.native_type_support_generate(writer)?;
//        Ok(())
//    }
//}

impl Generator for ast::Union {
    fn generate<W: io::Write>(&self, writer: &mut W) -> io::Result<()> {
        writeln!(writer, "{}", self.gen_builder())?;
        self.common_generate(writer)?;
        let struct_name = self.name().to_camel();

        let (union_impl, from_slice_switch_iml) = self.gen_union();
        writeln!(writer, "{}", union_impl)?;

        let struct_constructor = format!(
            r#"
func {struct_name}FromSlice(slice []byte, compatible bool) (*{struct_name}, error) {{
    sliceLen := len(slice)
    if uint32(sliceLen) < HeaderSizeUint {{
        errMsg := strings.Join([]string{{"HeaderIsBroken", "{struct_name}", strconv.Itoa(int(sliceLen)), "<", strconv.Itoa(int(HeaderSizeUint))}}, " ")
        return nil, errors.New(errMsg)
    }}
    itemID := unpackNumber(slice)
    innerSlice := slice[HeaderSizeUint:]

    switch itemID {{
    {from_slice_switch_iml}
    default:
        return nil, errors.New("UnknownItem, {struct_name}")
    }}
    return &{struct_name}{{inner: slice}}, nil
}}
            "#,
            struct_name = struct_name,
            from_slice_switch_iml = from_slice_switch_iml
        );
        writeln!(writer, "{}", struct_constructor)?;

        let struct_impl = format!(
            r#"
func (s *{struct_name}) ItemID() Number {{
    return unpackNumber(s.inner)
}}
func (s *{struct_name}) AsBuilder() {struct_name}Builder {{
    return *New{struct_name}Builder().Set(*s.ToUnion())
}}
            "#,
            struct_name = struct_name
        );
        writeln!(writer, "{}", struct_impl)?;
        self.native_type_support_generate(writer)?;
        Ok(())
    }
}

impl NativeGenerator for ast::Union {
    fn need_unpack() -> bool {
        true
    }

    fn native_type_support_generate(&self, writer: &mut impl io::Write) -> io::Result<()> {
        let struct_name = self.name().to_camel();
        let def = format!(
            r#"
            func Unpack{struct_name}(v *{struct_name}) *{struct_name}Union {{
            return v.ToUnion()
            }}
            "#);
        writeln!(writer, "{}", def)?;
        Ok(())
    }
}

impl Generator for ast::Array {
    fn generate<W: io::Write>(&self, writer: &mut W) -> io::Result<()> {
        let struct_name = self.name().to_camel();
        let inner = self.item().typ().name().to_camel();
        let item_count = self.item_count();
        let total_size = self.total_size();

        writeln!(writer, "{}", self.gen_builder())?;
        self.common_generate(writer)?;

        let impl_ = format!(
            r#"
func {struct_name}FromSlice(slice []byte, _compatible bool) (*{struct_name}, error) {{
    sliceLen := len(slice)
    if sliceLen != {total_size} {{
        errMsg := strings.Join([]string{{"TotalSizeNotMatch", "{struct_name}", strconv.Itoa(int(sliceLen)), "!=", strconv.Itoa({total_size})}}, " ")
        return nil, errors.New(errMsg)
    }}
    return &{struct_name}{{inner: slice}}, nil
}}
        "#,
            struct_name = struct_name,
            total_size = total_size
        );
        writeln!(writer, "{}", impl_)?;

        if self.item().typ().is_byte() {
            writeln!(
                writer,
                r#"
func (s *{struct_name}) RawData() []byte {{
    return s.inner
}}
            "#,
                struct_name = struct_name
            )?
        }

        for i in 0..self.item_count() {
            let func_name = format!("Nth{}", i);
            let start = self.item_size() * i;
            let end = self.item_size() * (i + 1);

            writeln!(
                writer,
                r#"
func (s *{struct_name}) {func_name}() *{inner_type} {{
    ret := {inner_type}FromSliceUnchecked(s.inner[{start}:{end}])
    return ret
}}
            "#,
                struct_name = struct_name,
                func_name = func_name,
                inner_type = inner,
                start = start,
                end = end
            )?
        }

        let as_builder_internal = (0..item_count)
            .map(|index| format!("t.Nth{index}(*s.Nth{index}())", index = index))
            .collect::<Vec<String>>()
            .join("\n");

        let as_builder = format!(
            r#"
func (s *{struct_name}) AsBuilder() {struct_name}Builder {{
	t := New{struct_name}Builder()
	{as_builder_internal}
	return *t
}}
        "#,
            struct_name = struct_name,
            as_builder_internal = as_builder_internal
        );

        writeln!(writer, "{}", as_builder)?;
        Ok(())
    }
}

impl Generator for ast::Struct {
    fn generate<W: io::Write>(&self, writer: &mut W) -> io::Result<()> {
        let struct_name = self.name().to_camel();
        let total_size = self.total_size();

        writeln!(writer, "{}", self.gen_builder())?;
        self.common_generate(writer)?;

        let impl_ = format!(
            r#"
func {struct_name}FromSlice(slice []byte, _compatible bool) (*{struct_name}, error) {{
    sliceLen := len(slice)
    if sliceLen != {total_size} {{
        errMsg := strings.Join([]string{{"TotalSizeNotMatch", "{struct_name}", strconv.Itoa(int(sliceLen)), "!=", strconv.Itoa({total_size})}}, " ")
        return nil, errors.New(errMsg)
    }}
    return &{struct_name}{{inner: slice}}, nil
}}
        "#,
            struct_name = struct_name,
            total_size = total_size
        );
        writeln!(writer, "{}", impl_)?;

        let (_, each_getter) = self.fields().iter().zip(self.field_sizes().iter()).fold(
            (0, Vec::with_capacity(self.fields().len())),
            |(mut offset, mut getters), (f, s)| {
                let func_name = f.name().to_camel();
                let inner = f.typ().name().to_camel();

                let start = offset;
                offset += s;
                let end = offset;
                let getter = format!(
                    r#"
func (s *{struct_name}) {func_name}() *{inner} {{
    ret := {inner}FromSliceUnchecked(s.inner[{start}:{end}])
    return ret
}}
                "#,
                    struct_name = struct_name,
                    inner = inner,
                    start = start,
                    end = end,
                    func_name = func_name
                );

                getters.push(getter);
                (offset, getters)
            },
        );

        writeln!(writer, "{}", each_getter.join("\n"))?;

        let as_builder = impl_as_builder_for_struct_or_table(&struct_name, self.fields());
        writeln!(writer, "{}", as_builder)?;

        self.native_type_support_generate(writer)?;
        Ok(())
    }
}

impl NativeGenerator for ast::Primitive {
    fn need_unpack() -> bool {
        true
    }

    fn native_type_support_generate(&self, writer: &mut impl Write) -> io::Result<()> {
        let struct_name = self.name().to_camel();
        let (unpack_def, native_type_name) = match self.size() {
            1 => ("return binary.LittleEndian.Uint8(v.inner)", "uint8"),
            2 => ("return binary.LittleEndian.Uint16(v.inner)", "uint16"),
            4 => ("return binary.LittleEndian.Uint32(v.inner)", "uint32"),
            8 => ("return binary.LittleEndian.Uint64(v.inner)", "uint64"),
            _ => unreachable!("invalid size"),
        };

        writeln!(writer,"{}", format!(r#"
        func Unpack{struct_name}(v *{struct_name}) {native_type_name} {{
             {unpack_def}
        }}
        "#))?;
        Ok(())
    }
}

impl NativeGenerator for ast::Struct {
    fn need_unpack() -> bool {
        true
    }

    fn native_type_support_generate(&self, writer: &mut impl Write) -> io::Result<()> {
        let struct_name = self.name().to_camel();
        let struct_fields_def = self.fields().iter().map(|f| {
        let field_name = f.name().to_camel();
        let field_type = f.typ().name().to_camel();
        format!("{} *{} `json:\"{}\"`", field_name, field_type, f.name())
        }).collect::<Vec<String>>().join("\n");

        // Generates the struct definition
        let struct_def = format!(r#"
type Native{struct_name} struct {{
     {struct_fields_def}
}}
"#);
        writeln!(writer, "{}", struct_def)?;

        let fields_unpack_def = self.fields().iter().map(|f| {

            let field_name = f.name().to_camel();
            let field_type = f.typ().name().to_camel();
            if f.typ().is_byte() {
                format!("native_val.{} = v.{}()", field_name, field_name)
            } else {
                format!("native_val.{} =  Unpack{}(v.{}())", field_name, field_type, field_name)
            }
        }).collect::<Vec<String>>().join("\n");

        // Generates Unpacking function
        let unpack_def = format!(r#"
func Unpack{struct_name}(v *{struct_name}) *Native{struct_name} {{
    native_val := &Native{struct_name}{{
    }}
    //s := v.AsBuilder()
    {fields_unpack_def}
    return native_val
}}
"#);
        writeln!(writer, "{}", unpack_def)?;

        Ok(())
    }
}

impl Generator for ast::FixVec {
    fn generate<W: io::Write>(&self, writer: &mut W) -> io::Result<()> {
        let struct_name = self.name().to_camel();
        let inner = self.item().typ().name().to_camel();
        let item_size = self.item_size();

        writeln!(writer, "{}", self.gen_builder())?;
        self.common_generate(writer)?;

        let constructor = format!(
            r#"
func {struct_name}FromSlice(slice []byte, _compatible bool) (*{struct_name}, error) {{
    sliceLen := len(slice)
    if sliceLen < int(HeaderSizeUint) {{
        errMsg := strings.Join([]string{{"HeaderIsBroken", "{struct_name}", strconv.Itoa(int(sliceLen)), "<", strconv.Itoa(int(HeaderSizeUint))}}, " ")
        return nil, errors.New(errMsg)
    }}
    itemCount := unpackNumber(slice)
    if itemCount == 0 {{
        if sliceLen != int(HeaderSizeUint) {{
            errMsg := strings.Join([]string{{"TotalSizeNotMatch", "{struct_name}", strconv.Itoa(int(sliceLen)), "!=", strconv.Itoa(int(HeaderSizeUint))}}, " ")
            return nil, errors.New(errMsg)
        }}
        return &{struct_name}{{inner: slice}}, nil
    }}
    totalSize := int(HeaderSizeUint) + int({item_size}*itemCount)
    if sliceLen != totalSize {{
        errMsg := strings.Join([]string{{"TotalSizeNotMatch", "{struct_name}", strconv.Itoa(int(sliceLen)), "!=", strconv.Itoa(int(totalSize))}}, " ")
        return nil, errors.New(errMsg)
    }}
    return &{struct_name}{{inner: slice}}, nil
}}
            "#,
            struct_name = struct_name,
            item_size = item_size
        );
        writeln!(writer, "{}", constructor)?;

        let impl_ = format!(
            r#"
func (s *{struct_name}) TotalSize() uint {{
    return uint(HeaderSizeUint) + {item_size} * s.ItemCount()
}}
func (s *{struct_name}) ItemCount() uint {{
    number := uint(unpackNumber(s.inner))
    return number
}}
func (s *{struct_name}) Len() uint {{
    return s.ItemCount()
}}
func (s *{struct_name}) IsEmpty() bool {{
    return s.Len() == 0
}}
// if *{inner_type} is nil, index is out of bounds
func (s *{struct_name}) Get(index uint) *{inner_type} {{
    var re *{inner_type}
    if index < s.Len() {{
        start := uint(HeaderSizeUint) + {item_size}*index
        end := start + {item_size}
        re = {inner_type}FromSliceUnchecked(s.inner[start:end])
    }}
    return re
}}
        "#,
            struct_name = struct_name,
            inner_type = inner,
            item_size = item_size
        );
        writeln!(writer, "{}", impl_)?;

        if self.item().typ().is_byte() {
            writeln!(
                writer,
                r#"
func (s *{struct_name}) RawData() []byte {{
    return s.inner[HeaderSizeUint:]
}}
            "#,
                struct_name = struct_name
            )?
        }
        let as_builder = impl_as_builder_for_vector(&struct_name);
        writeln!(writer, "{}", as_builder)?;
        Ok(())
    }
}

impl Generator for ast::DynVec {
    fn generate<W: io::Write>(&self, writer: &mut W) -> io::Result<()> {
        let struct_name = self.name().to_camel();
        let inner = self.item().typ().name().to_camel();

        writeln!(writer, "{}", self.gen_builder())?;
        self.common_generate(writer)?;

        let constructor = format!(
            r#"
func {struct_name}FromSlice(slice []byte, compatible bool) (*{struct_name}, error) {{
    sliceLen := len(slice)

    if uint32(sliceLen) < HeaderSizeUint {{
        errMsg := strings.Join([]string{{"HeaderIsBroken", "{struct_name}", strconv.Itoa(int(sliceLen)), "<", strconv.Itoa(int(HeaderSizeUint))}}, " ")
        return nil, errors.New(errMsg)
    }}

    totalSize := unpackNumber(slice)
    if Number(sliceLen) != totalSize {{
        errMsg := strings.Join([]string{{"TotalSizeNotMatch", "{struct_name}", strconv.Itoa(int(sliceLen)), "!=", strconv.Itoa(int(totalSize))}}, " ")
        return nil, errors.New(errMsg)
    }}

    if uint32(sliceLen) == HeaderSizeUint {{
        return &{struct_name}{{inner: slice}}, nil
    }}

    if uint32(sliceLen) < HeaderSizeUint*2 {{
        errMsg := strings.Join([]string{{"TotalSizeNotMatch", "{struct_name}", strconv.Itoa(int(sliceLen)), "<", strconv.Itoa(int(HeaderSizeUint*2))}}, " ")
        return nil, errors.New(errMsg)
    }}

    offsetFirst := unpackNumber(slice[HeaderSizeUint:])
    if uint32(offsetFirst)%HeaderSizeUint != 0 || uint32(offsetFirst) < HeaderSizeUint*2 {{
        errMsg := strings.Join([]string{{"OffsetsNotMatch", "{struct_name}", strconv.Itoa(int(offsetFirst%4)), "!= 0", strconv.Itoa(int(offsetFirst)), "<", strconv.Itoa(int(HeaderSizeUint*2))}}, " ")
        return nil, errors.New(errMsg)
    }}

    if sliceLen < int(offsetFirst) {{
        errMsg := strings.Join([]string{{"HeaderIsBroken", "{struct_name}", strconv.Itoa(int(sliceLen)), "<", strconv.Itoa(int(offsetFirst))}}, " ")
        return nil, errors.New(errMsg)
    }}
    itemCount := uint32(offsetFirst)/HeaderSizeUint - 1

    offsets := make([]uint32, itemCount)

    for i := 0; i < int(itemCount); i++ {{
        offsets[i] = uint32(unpackNumber(slice[HeaderSizeUint:][int(HeaderSizeUint)*i:]))
    }}

    offsets = append(offsets, uint32(totalSize))

    for i := 0; i < len(offsets); i++ {{
        if i&1 != 0 && offsets[i-1] > offsets[i] {{
            errMsg := strings.Join([]string{{"OffsetsNotMatch", "{struct_name}"}}, " ")
            return nil, errors.New(errMsg)
        }}
    }}

    for i := 0; i < len(offsets); i++ {{
        if i&1 != 0 {{
            start := offsets[i-1]
            end := offsets[i]
            _, err := {inner_type}FromSlice(slice[start:end], compatible)

            if err != nil {{
                return nil, err
            }}
        }}
    }}

    return &{struct_name}{{inner: slice}}, nil
}}
            "#,
            struct_name = struct_name,
            inner_type = inner
        );
        writeln!(writer, "{}", constructor)?;

        let impl_ = format!(
            r#"
func (s *{struct_name}) TotalSize() uint {{
    return uint(unpackNumber(s.inner))
}}
func (s *{struct_name}) ItemCount() uint {{
    var number uint = 0
    if uint32(s.TotalSize()) == HeaderSizeUint {{
        return number
    }}
    number = uint(unpackNumber(s.inner[HeaderSizeUint:]))/4 - 1
    return number
}}
func (s *{struct_name}) Len() uint {{
    return s.ItemCount()
}}
func (s *{struct_name}) IsEmpty() bool {{
    return s.Len() == 0
}}
// if *{inner_type} is nil, index is out of bounds
func (s *{struct_name}) Get(index uint) *{inner_type} {{
    var b *{inner_type}
    if index < s.Len() {{
        start_index := uint(HeaderSizeUint) * (1 + index)
        start := unpackNumber(s.inner[start_index:])

        if index == s.Len()-1 {{
            b = {inner_type}FromSliceUnchecked(s.inner[start:])
        }} else {{
            end_index := start_index + uint(HeaderSizeUint)
            end := unpackNumber(s.inner[end_index:])
            b = {inner_type}FromSliceUnchecked(s.inner[start:end])
        }}
    }}
    return b
}}
            "#,
            struct_name = struct_name,
            inner_type = inner
        );
        writeln!(writer, "{}", impl_)?;
        let as_builder = impl_as_builder_for_vector(&struct_name);
        writeln!(writer, "{}", as_builder)?;
        Ok(())
    }
}

impl Generator for ast::Table {
    fn generate<W: io::Write>(&self, writer: &mut W) -> io::Result<()> {
        let field_count = self.fields().len();
        let struct_name = self.name().to_camel();

        writeln!(writer, "{}", self.gen_builder())?;
        self.common_generate(writer)?;

        let constructor = if self.fields().is_empty() {
            format!(
                r#"
func New{struct_name}() {struct_name} {{
    s := new(bytes.Buffer)
    s.Write(packNumber(Number(HeaderSizeUint)))
    return {struct_name}{{inner: s.Bytes()}}
}}
func {struct_name}FromSlice(slice []byte, compatible bool) (*{struct_name}, error) {{
    sliceLen := len(slice)
    if uint32(sliceLen) < HeaderSizeUint {{
        return nil, errors.New("HeaderIsBroken")
    }}

    totalSize := unpackNumber(slice)
    if Number(sliceLen) != totalSize {{
        return nil, errors.New("TotalSizeNotMatch")
    }}

    if uint32(sliceLen) > HeaderSizeUint && !compatible {{
        return nil, errors.New("FieldCountNotMatch")
    }}
    return &{struct_name}{{inner: slice}}, nil
}}
            "#,
                struct_name = struct_name
            )
        } else {
            let verify_fields = self
                .fields()
                .iter()
                .enumerate()
                .map(|(i, f)| {
                    let field = f.typ().name().to_camel();
                    let start = i;
                    let end = i + 1;
                    format!(
                        r#"
_, err = {field}FromSlice(slice[offsets[{start}]:offsets[{end}]], compatible)
if err != nil {{
    return nil, err
}}
                "#,
                        field = field,
                        start = start,
                        end = end
                    )
                })
                .collect::<Vec<String>>()
                .join("\n");

            format!(
                r#"
func {struct_name}FromSlice(slice []byte, compatible bool) (*{struct_name}, error) {{
    sliceLen := len(slice)
    if uint32(sliceLen) < HeaderSizeUint {{
        errMsg := strings.Join([]string{{"HeaderIsBroken", "{struct_name}", strconv.Itoa(int(sliceLen)), "<", strconv.Itoa(int(HeaderSizeUint))}}, " ")
        return nil, errors.New(errMsg)
    }}

    totalSize := unpackNumber(slice)
    if Number(sliceLen) != totalSize {{
        errMsg := strings.Join([]string{{"TotalSizeNotMatch", "{struct_name}", strconv.Itoa(int(sliceLen)), "!=", strconv.Itoa(int(totalSize))}}, " ")
        return nil, errors.New(errMsg)
    }}

    if uint32(sliceLen) < HeaderSizeUint*2 {{
        errMsg := strings.Join([]string{{"TotalSizeNotMatch", "{struct_name}", strconv.Itoa(int(sliceLen)), "<", strconv.Itoa(int(HeaderSizeUint*2))}}, " ")
        return nil, errors.New(errMsg)
    }}

    offsetFirst := unpackNumber(slice[HeaderSizeUint:])
    if uint32(offsetFirst)%HeaderSizeUint != 0 || uint32(offsetFirst) < HeaderSizeUint*2 {{
        errMsg := strings.Join([]string{{"OffsetsNotMatch", "{struct_name}", strconv.Itoa(int(offsetFirst%4)), "!= 0", strconv.Itoa(int(offsetFirst)), "<", strconv.Itoa(int(HeaderSizeUint*2))}}, " ")
        return nil, errors.New(errMsg)
    }}

    if sliceLen < int(offsetFirst) {{
        errMsg := strings.Join([]string{{"HeaderIsBroken", "{struct_name}", strconv.Itoa(int(sliceLen)), "<", strconv.Itoa(int(offsetFirst))}}, " ")
        return nil, errors.New(errMsg)
    }}

    fieldCount := uint32(offsetFirst)/HeaderSizeUint - 1
    if fieldCount < {field_count} {{
        return nil, errors.New("FieldCountNotMatch")
    }} else if !compatible && fieldCount > {field_count} {{
        return nil, errors.New("FieldCountNotMatch")
    }}

    offsets := make([]uint32, fieldCount)

    for i := 0; i < int(fieldCount); i++ {{
        offsets[i] = uint32(unpackNumber(slice[HeaderSizeUint:][int(HeaderSizeUint)*i:]))
    }}
    offsets = append(offsets, uint32(totalSize))

    for i := 0; i < len(offsets); i++ {{
        if i&1 != 0 && offsets[i-1] > offsets[i] {{
            return nil, errors.New("OffsetsNotMatch")
        }}
    }}

    var err error
    {verify_fields}

    return &{struct_name}{{inner: slice}}, nil
}}
            "#,
                struct_name = struct_name,
                field_count = field_count,
                verify_fields = verify_fields
            )
        };
        writeln!(writer, "{}", constructor)?;

        let impl_ = format!(
            r#"
func (s *{struct_name}) TotalSize() uint {{
    return uint(unpackNumber(s.inner))
}}
func (s *{struct_name}) FieldCount() uint {{
    var number uint = 0
    if uint32(s.TotalSize()) == HeaderSizeUint {{
        return number
    }}
    number = uint(unpackNumber(s.inner[HeaderSizeUint:]))/4 - 1
    return number
}}
func (s *{struct_name}) Len() uint {{
    return s.FieldCount()
}}
func (s *{struct_name}) IsEmpty() bool {{
    return s.Len() == 0
}}
func (s *{struct_name}) CountExtraFields() uint {{
    return s.FieldCount() - {field_count}
}}

func (s *{struct_name}) HasExtraFields() bool {{
    return {field_count} != s.FieldCount()
}}
            "#,
            struct_name = struct_name,
            field_count = field_count,
        );
        writeln!(writer, "{}", impl_)?;

        let (getter_stmt_last, getter_stmt) = {
            let getter_stmt_last = "s.inner[start:]".to_string();
            let getter_stmt = "s.inner[start:end]".to_string();
            (getter_stmt_last, getter_stmt)
        };
        let each_getter = self
            .fields()
            .iter()
            .enumerate()
            .map(|(i, f)| {
                let func = f.name().to_camel();

                let inner = f.typ().name().to_camel();
                let start = (i + 1) * NUMBER_SIZE;
                let end = (i + 2) * NUMBER_SIZE;
                if i == self.fields().len() - 1 {
                    format!(
                        r#"
func (s *{struct_name}) {func}() *{inner} {{
    var ret *{inner}
    start := unpackNumber(s.inner[{start}:])
    if s.HasExtraFields() {{
        end := unpackNumber(s.inner[{end}:])
        ret = {inner}FromSliceUnchecked({getter_stmt})
    }} else {{
        ret = {inner}FromSliceUnchecked({getter_stmt_last})
    }}
    return ret
}}
                        "#,
                        struct_name = struct_name,
                        start = start,
                        end = end,
                        func = func,
                        inner = inner,
                        getter_stmt = getter_stmt,
                        getter_stmt_last = getter_stmt_last
                    )
                } else {
                    format!(
                        r#"
func (s *{struct_name}) {func}() *{inner} {{
    start := unpackNumber(s.inner[{start}:])
    end := unpackNumber(s.inner[{end}:])
    return {inner}FromSliceUnchecked({getter_stmt})
}}
               "#,
                        struct_name = struct_name,
                        func = func,
                        inner = inner,
                        getter_stmt = getter_stmt,
                        start = start,
                        end = end
                    )
                }
            })
            .collect::<Vec<_>>();
        writeln!(writer, "{}", each_getter.join("\n"))?;

        let as_builder = impl_as_builder_for_struct_or_table(&struct_name, self.fields());
        writeln!(writer, "{}", as_builder)?;
        Ok(())
    }
}

impl NativeGenerator for ast::Table {
    fn need_unpack() -> bool {
        true
    }

    fn native_type_support_generate(&self, writer: &mut impl Write) -> io::Result<()> {
        let struct_name = self.name().to_camel();
        let struct_fields_def = self.fields().iter().map(|f| {
            let field_name = f.name().to_camel();
            let field_type = f.typ().name().to_camel();
            format!("{} *{} `json:\"{}\"`", field_name, field_type, f.name())
        }).collect::<Vec<String>>().join("\n");

        // Generates the struct definition
        let struct_def = format!(r#"
type Native{struct_name} struct {{
     {struct_fields_def}
}}
"#);
        writeln!(writer, "{}", struct_def)?;

        let fields_unpack_def = self.fields().iter().map(|f| {
            let field_name = f.name().to_camel();
            let field_type = f.typ().name().to_camel();
            if f.typ().is_byte() {
                format!("native_val.{} = s.{}()", field_name, field_name)
            } else {
                format!("native_val.{} =  Unpack{}(s.{}())", field_name, field_type, field_name)
            }
        }).collect::<Vec<String>>().join("\n");

        // Generates Unpacking function
        let unpack_def = format!(r#"
func Unpack{struct_name}(v *{struct_name}) *Native{struct_name} {{
    native_val := &Native{struct_name}{{
    }}
    s := v.AsBuilder()
    {fields_unpack_def}
    return native_val
}}
"#);
        writeln!(writer, "{}", unpack_def)?;

        Ok(())
    }
}

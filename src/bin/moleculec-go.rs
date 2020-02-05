use moleculec_go::build_commandline;
use molecule_codegen::IntermediateFormat;

fn main() {
    build_commandline(IntermediateFormat::JSON).execute();
}

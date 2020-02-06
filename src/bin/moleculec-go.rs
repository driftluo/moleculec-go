use molecule_codegen::IntermediateFormat;
use moleculec_go::build_commandline;

fn main() {
    build_commandline(IntermediateFormat::JSON).execute();
}

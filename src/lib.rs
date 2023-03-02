use std::{
    io::{self, Read, Write},
    process,
};

use molecule_codegen::IntermediateFormat;

mod codegen;

use codegen::Generator;

pub(crate) enum AppAction {
    DisplayFormat,
    ProcessIntermediate(Vec<u8>),
}

pub struct AppConfig {
    action: AppAction,
    format: IntermediateFormat,
}

type RawAppConfig = (IntermediateFormat, clap::ArgMatches);

pub fn build_commandline(format: IntermediateFormat) -> AppConfig {
    let matches = clap::Command::new("moleculec-go")
        .name("Moleculec Plugin")
        .about("Compiler plugin for molecule to generate code.")
        .version(clap::crate_version!())
        .arg(
            clap::Arg::new("format")
                .long("format")
                .help("Output the supported format for the intermediate data.")
                .action(clap::ArgAction::SetTrue),
        )
        .get_matches();
    AppConfig::from((format, matches))
}

impl From<RawAppConfig> for AppConfig {
    fn from(input: RawAppConfig) -> Self {
        let (format, matches) = input;
        let action = if matches.get_flag("format") {
            AppAction::DisplayFormat
        } else {
            let mut input = Vec::new();
            if io::stdin().read_to_end(&mut input).is_err() {
                eprintln!("Error: failed to read data from stdin)");
                process::exit(1);
            };
            AppAction::ProcessIntermediate(input)
        };
        Self { action, format }
    }
}

impl AppConfig {
    pub fn execute(self) {
        match self.action {
            AppAction::DisplayFormat => {
                println!("{}", self.format);
            }
            AppAction::ProcessIntermediate(input) => {
                let ast = self.format.recover(&input).unwrap();

                let mut output_data = Vec::<u8>::new();

                if let Err(err) = Generator::generate(&mut output_data, &ast) {
                    eprintln!("failed to write data by generator: {}", err)
                }

                let stdout = io::stdout();
                let mut stdout_handle = stdout.lock();
                stdout_handle.write_all(&output_data).unwrap();
                stdout_handle.flush().unwrap();
            }
        }
    }
}

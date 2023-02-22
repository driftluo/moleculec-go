MOLC    := moleculec
MOLC_VERSION := 0.7.4

ci:
	cargo fmt --all -- --check
	RUSTFLAGS='-D warnings' cargo clippy --all
	cargo install moleculec --vers ${MOLC_VERSION}
	cargo install --path .
	moleculec --language go --schema-file ./test/schema/types.mol | gofmt > ./test/molecule-test/types.go
	cd ./test/molecule-test/ && go test -count=1 -v
	rm ./test/molecule-test/types.go && touch test/molecule-test/types.go

check-moleculec-version:
	test "$$(${MOLC} --version | awk '{ print $$2 }' | tr -d ' ')" = ${MOLC_VERSION}

gen-test:
	cargo build --all
	./target/debug/test-codegen ./test/schema/types.mol ./test/testset/default.yaml  | gofmt > ./test/molecule-test/default_test.go
	./target/debug/test-codegen ./test/schema/types.mol ./test/testset/simple.yaml  | gofmt > ./test/molecule-test/simple_test.go

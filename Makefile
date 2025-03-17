MOLC    := moleculec
MOLC_VERSION := 0.9.0
MOLC_TEST_DIR = ./test/molecule-test

ci:
	cargo fmt --all -- --check
	RUSTFLAGS='-D warnings' cargo clippy --all
	cargo install moleculec --vers ${MOLC_VERSION}
	cargo install --path .
	$(shell mkdir -p ${MOLC_TEST_DIR})
	moleculec --language go --schema-file ./test/schema/types.mol | gofmt > ./test/molecule-test/types.go
	make gen-test
	cd ./test/molecule-test/ && go test -count=1 -v
	rm ./test/molecule-test/types.go && touch test/molecule-test/types.go
	rm ./test/molecule-test/default_test.go && rm test/molecule-test/simple_test.go

check-moleculec-version:
	test "$$(${MOLC} --version | awk '{ print $$2 }' | tr -d ' ')" = ${MOLC_VERSION}

gen-test:
	cargo build --all
	./target/debug/test-codegen ./test/schema/types.mol ./test/testset/default.yaml  | gofmt > ./test/molecule-test/default_test.go
	./target/debug/test-codegen ./test/schema/types.mol ./test/testset/simple.yaml  | gofmt > ./test/molecule-test/simple_test.go

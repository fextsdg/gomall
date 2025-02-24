.PHONY: gen-demo
gen-demo:
	@cd protobuf/demo && cwgo server --server_name demoproto --type RPC --module gomall/probuf/demo --I ../ --idl hello.proto


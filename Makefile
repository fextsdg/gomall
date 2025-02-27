.PHONY: gen-demo
gen-demo:
	@cd protobuf/demo && cwgo server --server_name demoproto --type RPC --module gomall/probuf/demo --I ../ --idl hello.proto

.PHONY :gen-appfront
gen-appfront:
	@cd app/frontend && cwgo server --type HTTP --idl ..\..\idl\frontend\auth_page.proto   --service frontend -module gomall/app/frontend -I ..\..\idl\


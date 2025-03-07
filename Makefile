.PHONY: gen-demo
gen-demo:
	@cd protobuf/demo && cwgo server --server_name demoproto --type RPC --module gomall/probuf/demo --I ../ --idl hello.proto

.PHONY :gen-appfront
gen-appfront:
	@cd app/frontend && cwgo server --type HTTP --idl ..\..\idl\frontend\auth_page.proto   --service frontend --module gomall/app/frontend -I ..\..\idl\
	@cd app/frontend && cwgo server --type HTTP --idl ..\..\idl\frontend\product_page.proto   --service frontend --module gomall/app/frontend -I ..\..\idl\
	@cd app/frontend && cwgo server --type HTTP --idl ..\..\idl\frontend\category_page.proto   --service frontend --module gomall/app/frontend -I ..\..\idl\
	@cd app/frontend && cwgo server --type HTTP --idl ..\..\idl\frontend\cart_page.proto   --service frontend --module gomall/app/frontend -I ..\..\idl\
.PHONY: gen-user-client
gen-user-client:
	@cd rpc_gen && cwgo client --type rpc --I ../idl --idl ../idl/user.proto --service user --module gomall/rpc_gen

.PHONY:gen-user-server
gen-user-server:
	@cd app/user && cwgo server --type rpc --service user --module gomall/app/user  -I ../../idl --idl ../../idl/user.proto --pass "-use gomall/rpc_gen/kitex_gen"

.PHONY :gen-product-client
gen-product-client:
	@cd rpc_gen && cwgo client --service product --type rpc --I ../idl --idl ../idl/product.proto --module gomall/rpc_gen

.PHONY :gen-product-server
gen-product-server:
	@cd app/product && cwgo server --type rpc --service product --module gomall/app/product  -I ../../idl --idl ../../idl/product.proto --pass "-use gomall/rpc_gen/kitex_gen"

.PHONY :gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --service cart --type rpc --I ../idl --idl ../idl/cart.proto --module gomall/rpc_gen
	@cd app/cart && cwgo server --type rpc --service cart --module gomall/app/cart  -I ../../idl --idl ../../idl/cart.proto --pass "-use gomall/rpc_gen/kitex_gen"

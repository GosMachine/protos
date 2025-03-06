VERSION=$(shell git describe --tags --abbrev=0)
NEW_VERSION=$(shell echo $(VERSION) | awk -F. '{$$NF = $$NF + 1;} 1' OFS=.)

COMMIT_MSG="optimized"

commit:
	git add .
	git commit -m "$(COMMIT_MSG)"

tag:
	git tag $(NEW_VERSION)

push: 
	git push && git push --tags

release: commit tag push
	@echo "Pushed and tagged with version $(NEW_VERSION)"

auth:
	@protoc -I proto proto/auth/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
	@"$(MAKE)" release

product:
	@protoc -I proto proto/product/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
	@"$(MAKE)" release

marketplaces:
	@protoc -I proto proto/marketplaces/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
	@"$(MAKE)" release 

drops:
	@protoc -I proto proto/drops/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
	@"$(MAKE)" release

dropscenter:
	@protoc -I proto proto/drops/*.proto -I proto proto/dropscenter/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
	@"$(MAKE)" release

gosboost:
	@protoc -I proto proto/gosboost/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
	@"$(MAKE)" release

payments:
	@protoc -I proto proto/payments/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
	@"$(MAKE)" release

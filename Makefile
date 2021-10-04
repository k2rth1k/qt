SEMVERSION_CLI = scripts/semver bump

.Phony: proto
proto:
	@echo "compiling proto files...."
	@cd ./pkg/ && buf generate

.Phony: build
build:	proto
	@echo "building application...."
	@go build -o build/arm64/qt main.go
	@echo "successfully built application"

.Phony: run
run: build
	@echo "starting service....."
	@./build/arm64/qt

.Phony: go
go: ## go checks module packages
	$(ECHO) go list -mod=readonly -u -m all
	$(ECHO) go mod verify
	$(ECHO) go mod tidy
	$(ECHO) go mod vendor

.PHONY: bump_patch
bump_patch:  ## Bump the patch version
	$(ECHO) $(SEMVERSION_CLI) patch $$(cat VERSION) > VERSION
	git add .
	git commit -m "New patch version $$(cat VERSION)"
	git tag -m "release of $$(cat VERSION) $$(date)" v$$(cat VERSION)
	git push --tags

.PHONY: bump_minor
bump_minor:  ## Bump the minor version
	$(ECHO) $(SEMVERSION_CLI) minor $$(cat VERSION) > VERSION
	git add .
	git commit -m "New minor version $$(cat VERSION)"
	git tag -m "release of $$(cat VERSION) $$(date)" v$$(cat VERSION)
	git push --tags

.PHONY: bump_major
bump_major:  ## Bump the major version
	$(ECHO) $(SEMVERSION_CLI) major $$(cat VERSION) > VERSION
	git add .
	git commit -m "New major version $$(cat VERSION)"
	git tag -m "release of $$(cat VERSION) $$(date)" v$$(cat VERSION)
	git push --tags

.PHONY: test
test:
	for i in $$(find . -maxdepth 4 -mindepth 1 -type f -name '*_test.go' | sed -r 's|/[^/]+$$||' |sort |uniq); do \
  				go test -v $$i; \
    done;

>PHONY: sql
sql:
	if ( test -f "init.sql" ); then \
        	> init.sql; \
	fi;
	for i in $$(find . -type f -name "*up.sql"); do \
  		cat $$i >> init.sql; \
  		echo "" >> init.sql;    \
  		echo "" >> init.sql;    \
    done;

>PHONY: docker
docker: sql
		$(ECHO) docker-compose up
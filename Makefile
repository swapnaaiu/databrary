# Editing this Makefile? run make with `make BUILDDEV=1` for better build
# debugging.

# Not everything is cached in our own cache, so use both
NIX_OPTIONS := --option binary-caches "https://cache.nixos.org http://devdatabrary2.home.nyu.edu:5000"

# These below intentionally use '='to pick up following changes to NIX_OPTIONS
nix-build = nix-build $(NIX_OPTIONS) --drv-link $(PWD)/derivation --cores 4 -A databrary
nix-shell = nix-shell $(NIX_OPTIONS) --pure

ifdef BUILDDEV
nix-build += --keep-failed
endif

# Sneaky options used in recursing make. Not for human consumption.
ifdef __COVERAGE
nix-build += --arg coverage true
endif
ifdef __HADDOCK
nix-build += --arg haddock true
endif

#
# COMMON TASKS
#

## The default action is to run tests
cabal-test: ; $(nix-shell) --run 'cabal -j test --test-options="--color always --hide-successes --timeout 1s" --ghc-option=-O0'
.PHONY: cabal-test

## Start the db (needed for cabal.test and cabal.build, but that relationship is
## not captured by Make yet)
db: ; $(nix-shell) --run ./init-db-pql.sh
.PHONY: db

## Start the dev repl
repl: ; $(nix-shell) --run ./ghci-databrary.sh
.PHONY: repl

## Start tests in the repl
repl-test: ; $(nix-shell) --run 'cabal repl test:discovered'
.PHONY: repl-test

## One can always build with Nix.
nix-build: ; $(nix-build)
.PHONY: nix-build

## You can also build with Cabal if that suits you
cabal-build: ; $(nix-shell) --run 'cabal -j new-build --disable-optimization'
.PHONY: cabal-build

## Simple report output, long build time.
reports:
	__HADDOCK=1 __COVERAGE=1 make nix-build
	hpc report result/share/hpc/vanilla/tix/databrary-1/databrary-1.tix \
		--hpcdir=result/share/hpc/vanilla/mix/databrary-1 \
		--exclude=Paths_databrary
		| tee hpc_report.txt
	nix-store -l $< | grep ") in '" > haddock_coverage_report.txt
.PHONY: reports

#
# Experimental tasks
#

##
## These may eventually be used to auto-generate databrary.cabal
##

module_list.yaml:
	echo -e '_modules:\n- Paths_databrary' > $@
	find src -regex '.*\.hsc?' >> $@
	sed -i -e 's@src/\(.*\).hs.*@- \1@' -e 's@/@.@g' $@


testFiles := $(shell find test -name '*Test.hs')
tests_module_list.yaml: $(testFiles)
	echo -e '_modules:\n' > $@
	find test -regex '.*\.hsc?' >> $@
	sed -i -e 's@test/\(.*\).hs.*@- \1@' -e 's@/@.@g' $@

tests_module_list.cabalsnip: tests_module_list.yaml
	yq -r '.[]|join("\n")' < $< > $@

package.yaml: module_list.yaml

databrary.cabal: package.yaml ; hpack .

##
## This is the beginning of packaging up node deps
##

nodeFiles := $(addprefix node/, default.nix node-env.nix node-packages.nix)

node: $(nodeFiles)
.PHONY: node

$(nodeFiles) : node/package.json
	cd node && nix-shell -Qp nodePackages.node2nix --run node2nix
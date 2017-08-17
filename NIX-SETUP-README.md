DISCLAIMER: This is a proto-type document for development. Do not use. 

This will guide a NIX-HASKELL DEVELOPER through the setup got Databrary so
far, please update this document with fixes as they are deployed. 

----------------------------------------------------------------------------
INTERNAL NOTE: A script should be added to handle pre-requesites and part of
the pre-building process. Include:
    *cloning the rep0
    *git checkout obsidian-develop
    *(Check/Download) Postgresql
    *(Check/Download) Solr
    *(Check/Download) Cabal
    *generate databrary.conf file 
---------------------------------------------------------------------------

STEP 1: 
  From the root directory of this project, run 
  ```bash 
  $./bar
  ``` 
  This is a script to set up the postgresql server along with required schemas

STEP 2:
  From the same root directory of this project, run 
  ```bash
  $nix-shell
  $cabal repl databrary
  ```

STEP 3: 
  When cabal successfully builds all 303 dependencies, run 
  ```bash
  $main
  ```
Note: If errors occur, Some directories have previously been hardcoded to search for their
requirements within the root directory of a generic Unix Distro. Fixes will
have to be deployed to make this more flexible. 

Additional Note: In order to clean your test environment, it is best to
execute the following commands: 
```bash
$dropdb databrary
$dropuser databrary
$pg_ctl -D databrary-db -l logfile stop
$rm -rf databrary-db/
```
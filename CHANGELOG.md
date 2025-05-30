# CHANGELOG
All notable changes to this project will be documented in this file. This project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
<a name="unreleased"></a>
## [Unreleased]


<a name="v1.9.7"></a>
## [v1.9.7] - 2025-05-30
### Bug Fixes
- **embedded/tbtree:** save timestamp to a separate file to avoid rescanning empty indexes


<a name="v2.0.0-RC1"></a>
## [v2.0.0-RC1] - 2025-05-20
### Changes
- v2 version


<a name="v1.9.6"></a>
## [v1.9.6] - 2025-03-31
### Bug Fixes
- **embedded/sql:** correctly handle logical operator precedence (NOT, AND, OR)

### Changes
- **embedded/sql:** Support PRIMARY KEY constraint on individual columns
- **embedded/sql:** Add support for core pg_catalog tables (pg_class, pg_namespace, pg_roles)
- **embedded/sql:** add support for LEFT JOIN
- **embedded/sql:** add support for SELECT FROM VALUES syntax
- **embedded/sql:** Allow expressions in ORDER BY clauses
- **embedded/sql:** Add support for simple CASE statements
- **embedded/sql:** Implement CASE statement
- **pkg/replication:** add replication lag metric


<a name="v1.9.5"></a>
## [v1.9.5] - 2024-09-16
### Bug Fixes
- time.Since should not be used in defer statement
- **pkg/dabase:** return error when attempting to access deleted database
- **pkg/pgsql/server:** close row readers to release resources
- **pkg/server:** run metrics server under HTTPS

### Changes
- **embedded/logging:** improve file base logging.
- **embedded/sql:** improvements on SQL layer.
- **embedded/store:** improve index flush logic
- **pkg/database:** implement database manager
- **pkg/server:** implement automatic generation of self-signed HTTPS certificate


<a name="v1.9.4"></a>
## [v1.9.4] - 2024-07-25
### Bug Fixes
- set mattermost payload


<a name="v1.9.3"></a>
## [v1.9.3] - 2024-05-23
### Changes
- fix some comments
- refactor image building
- The GitHub workflow push.yml was updated to use the repository owner's Docker images instead of fixed ones. This change allows for more flexibility and control when using Docker images, and ensures that the correct images are used based on the repository owner.
- Update GitHub Actions to use checkout[@v4](https://github.com/v4)
- Add support to ARM64
- **embedded/cache:** replace sync.Mutex with sync.RWMutex
- **embedded/cache:** validate input params before obtaining mutex lock

### Reverts
- test with github token
- chore(embedded/cache): replace sync.Mutex with sync.RWMutex


<a name="v1.9DOM.2"></a>
## [v1.9DOM.2] - 2023-12-29
### Bug Fixes
- apply fix for CVE-2023-44487
- performance test regression

### Changes
- **deps:** bump actions/setup-go from 3 to 5
- **deps:** bump actions/upload-artifact from 3 to 4
- **deps:** bump actions/download-artifact from 3 to 4
- **deps:** bump google.golang.org/grpc in /test/e2e/truncation
- **deps:** bump google.golang.org/protobuf from 1.31.0 to 1.32.0


<a name="v1.9DOM.2-RC1"></a>
## [v1.9DOM.2-RC1] - 2023-12-21
### Bug Fixes
- performance test regression
- remove influxdb dependencies
- correct the test after the merge and latest refactor
- source /etc/sysconfig/immudb on AWS EC2 startup

### Changes
- add influxdb (needed for performance test) dependency
- use goveralls token variable
- **deps:** bump golang.org/x/crypto in /test/columns
- **deps:** bump golang.org/x/crypto in /tools/mkdb
- **deps:** bump golang.org/x/crypto
- **deps:** bump golang.org/x/crypto in /test/e2e/truncation
- **deps:** bump golang.org/x/crypto
- **deps:** bump github.com/rogpeppe/go-internal from 1.9.0 to 1.12.0
- **deps:** bump golang.org/x/net from 0.17.0 to 0.19.0
- **deps:** bump golang.org/x/crypto from 0.14.0 to 0.17.0

### Features
- add s3 (aws) role based auth as an option
- automatically convert uuid strings and byte slices to uuid values

### Reverts
- Merge remote-tracking branch 'origin/dependabot/go_modules/github.com/rogpeppe/go-internal-1.12.0' into release/v1.9.2


<a name="v1.9DOM.1"></a>
## [v1.9DOM.1] - 2023-11-16
### Changes
- **pkg/pgsql:** handle odbc help
- **pkg/server:** change permission automatically revokes existing ones


<a name="v1.9DOM.1-RC1"></a>
## [v1.9DOM.1-RC1] - 2023-11-14
### Bug Fixes
- lower databasename in OpenSession
- **embedded/sql:** fix data-race when mapping keys
- **embedded/sql:** fix data-race when mapping keys
- **embedded/store:** handle key mapping in ongoing txs
- **embedded/store:** handle key mapping in ongoing txs
- **embedded/store:** handle key mapping in ongoing txs
- **pkg/database:** ensure proper tx validation
- **pkg/server:** user creation with multidbs

### Changes
- docker image with swagger ui (for AWS Marketplace)
- **cmd/immudb:** upgrade to new pgsql changes
- **deps:** bump github.com/google/uuid from 1.3.1 to 1.4.0
- **embedded/sql:** user pwd
- **embedded/sql:** show users stmt
- **embedded/sql:** wip emulate pg_type system table
- **embedded/sql:** continue to support databases and tables datasources
- **embedded/store:** indexer source and target prefixes
- **pkg/client:** possibility to retrieve session id
- **pkg/pgsql:** support multiple-statements in simple-query mode
- **pkg/pgsql:** tls support
- **pkg/pgsql:** comment describing pgsql wire protocol constraints
- **pkg/pgsql:** show table/s
- **pkg/pgsql:** proper handling of queries with empty resultsets
- **pkg/pgsql:** single command complete message
- **pkg/pgsql:** protocol enhancements
- **pkg/pgsql:** uuid and float types conversion
- **pkg/pgsql:** transactional query machine
- **pkg/pgsql:** pgsql write protocol improvements
- **pkg/pgsql:** decouple error from ready to query messages
- **pkg/pgsql:** handle deallocate prepared stmt
- **pkg/server:** pgsql server creation only when enabled
- **pkg/server:** set dynamic immudb server port in pgsql server
- **pkg/server:** upgrade to transactional pgsql server
- **pkg/server:** list users from multidb handler
- **pkg/server:** require proper permissions at multidb handler

### Features
- **embedded/sql:** show table stmt
- **embedded/sql:** wip user mgmt
- **embedded/sql:** show users stmt
- **embedded/sql:** show databases/tables stmt
- **pkg/server:** add support of underscore in db name Signed-off-by: Martin Jirku <martin[@jirku](https://github.com/jirku).sk>


<a name="v1.9DOM.0"></a>
## [v1.9DOM.0] - 2023-10-19
### Changes
- docker image with swagger ui
- docker image with swagger ui


<a name="v1.9DOM"></a>
## [v1.9DOM] - 2023-10-19
### Changes
- **cmd/immuadmin:** add indexing related flags

### Features
- **embedded/sql:** table renaming


<a name="v1.9.0-RC2"></a>
## [v1.9.0-RC2] - 2023-10-16
### Bug Fixes
- standard syntax for drop index
- **embedded/sql:** fix sql temporal range evaluation

### Changes
- **embedded/document:** count with limit in subquery
- **embedded/sql:** expose subquery creation
- **pkg/api:** set optional parameters
- **pkg/api:** set optional parameters


<a name="v1.9.0-RC1"></a>
## [v1.9.0-RC1] - 2023-10-11
### Bug Fixes
- insertion ts for key-values should not be equal to the current root ts
- correct immudb name in readme
- allow the local id to be used if present even if remote flag is on
- apply fixes discussed in PR
- **Makefile:** remove webconsole tag from immuclient/immuadmin builds
- **embedded/appendable:** explicit freebsd build clauses
- **embedded/document:** avoid waiting for tx to be committed
- **embedded/document:** ensure multi-indexing is enabled
- **embedded/sql:** advance position when decoding value at deleted column
- **embedded/store:** precommitted transaction discarding recedes durable state
- **embedded/store:** use correct index path
- **embedded/store:** handle transient key update
- **embedded/store:** read lock when fetching indexer
- **embedded/store:** read lock when pausing indexers
- **embedded/tbtree:** proper _rev calculation
- **embedded/tbtree:** snapshot validation
- **embedded/tbtree:** consider offset for history count calculation
- **pkg/server:** buffer reuse

### Changes
- use copy instead of a loop
- build with swaggerui
- unnecessary use of fmt.Sprintf
- align covered packages when pulling and merging
- unnecessary use of fmt.Sprintf
- **cmd/immuclient:** display raw column selector in table header
- **deps:** bump golang.org/x/net from 0.10.0 to 0.12.0
- **deps:** bump google.golang.org/grpc from 1.55.0 to 1.56.2
- **deps:** bump golang.org/x/crypto from 0.13.0 to 0.14.0
- **deps:** bump golang.org/x/net from 0.14.0 to 0.15.0
- **deps:** bump google.golang.org/grpc
- **deps:** bump google.golang.org/grpc in /test/e2e/truncation
- **deps:** bump google.golang.org/grpc
- **deps:** bump golang.org/x/crypto from 0.12.0 to 0.13.0
- **deps:** bump golang.org/x/crypto from 0.10.0 to 0.11.0
- **deps:** bump golang.org/x/sys from 0.9.0 to 0.10.0
- **deps:** bump golang.org/x/net from 0.15.0 to 0.17.0
- **deps:** bump golang.org/x/sys from 0.11.0 to 0.12.0
- **deps:** bump golang.org/x/net from 0.12.0 to 0.13.0
- **deps:** bump golang.org/x/sys from 0.10.0 to 0.11.0
- **deps:** bump golang.org/x/crypto from 0.7.0 to 0.10.0
- **deps:** bump golang.org/x/net from 0.13.0 to 0.14.0
- **deps:** bump securego/gosec from 2.15.0 to 2.17.0
- **deps:** bump github.com/grpc-ecosystem/grpc-gateway/v2
- **embedded/document:** encoded document using valRef
- **embedded/document:** register username when applying a change
- **embedded/document:** attach username when auditing document
- **embedded/document:** enable multi-indexing in doc engine tests
- **embedded/sql:** support parenthesis as datatype constraint delimiter
- **embedded/sql:** multi-snapshop mvvc
- **embedded/sql:** deletion of primary index path
- **embedded/sql:** historical queries over primary index
- **embedded/sql:** dynamic indexing
- **embedded/sql:** post-commit physical index deletion
- **embedded/sql:** improve internal index naming
- **embedded/sql:** uuid decoding
- **embedded/sql:** unique index creation supported on empty tables
- **embedded/sql:** temporal queries with multi-indexing
- **embedded/sql:** transactional drops
- **embedded/sql:** use declared constant for fixed ids
- **embedded/sql:** insertion benchmark
- **embedded/store:** injective index mapper
- **embedded/store:** history returning value refs
- **embedded/store:** history with rev count
- **embedded/store:** ensure index is erased from disk
- **embedded/store:** indexer alloc its tx
- **embedded/store:** remove metastate
- **embedded/store:** multi-indexing
- **embedded/store:** key reader including historical entries
- **embedded/store:** ensure snapshot up to date
- **embedded/store:** indexing callbacks
- **embedded/store:** wip multi-indexing
- **embedded/store:** indexing prefix
- **embedded/store:** entry mapper
- **embedded/tbtree:** value-preserving history
- **embedded/tbtree:** fetching historical values
- **embedded/tbtree:** wip value-preserving history
- **embedded/tbtree:** context propagation
- **embedded/tbtree:** value-preserving history
- **embedded/tbtree:** hcount serialization
- **pkg/api:** adjust doc serializations to match verification
- **pkg/api:** endpoint improvements
- **pkg/client:** add setAll to immuclient mock
- **pkg/client:** use buf for msg exchange
- **pkg/database:** increase delay when tx is not present
- **pkg/database:** fix remote storage paths
- **pkg/database:** mandatory wait with async replication
- **pkg/database:** kv count
- **pkg/database:** doc audit without retrieving payloads
- **pkg/database:** context propagation
- **pkg/database:** context propagation
- **pkg/database:** multi-indexing database
- **pkg/database:** fix remote storage paths
- **pkg/database:** register username when applying a change
- **pkg/database:** keept reading from specific tx
- **pkg/server:** register username when applying a change in doc apis
- **pkg/server:** minor code adjustment
- **pkg/stdlib:** non transactional ddl stmts
- **pkg/truncator:** use embedded/logger package
- **pkg/verification:** minor doc verification improvements
- **swagger:** use embedded logger package
- **tests:** Tests cleanup

### Code Refactoring
- **pkg/logger:** move logger from pkg to embedded

### Features
- add flag for using external id as a main one
- update readme
- prevent identifier from creation when use external id option
- pass logger to heartbeater
- **embedded/document:** register user when creating collection
- **embedded/document:** doc audit without retrieving payloads
- **embedded/document:** add field to collection
- **embedded/document:** remove field from collection
- **embedded/sql:** dynamic multi-indexing
- **embedded/sql:** wip uuid datatype support
- **embedded/sql:** include _rev column in historical queries
- **embedded/sql:** drop index and table stmts
- **embedded/sql:** table history
- **embedded/sql:** query including historical rows
- **embedded/sql:** extra metadata when creating tx
- **embedded/sql:** async multi-indexing
- **embedded/sql:** drop column stmt
- **embedded/store:** getBetween
- **embedded/store:** use index attribute in kv metadata
- **embedded/store:** extra tx metadata
- **embedded/store:** transactionaless multi-indexing
- **embedded/tbtree:** getBetween
- **embedded/tbtree:** key reader supporting historical values
- **pkg/api:** include username in document audit response
- **pkg/api:** re-enable swagger ui
- **pkg/api:** docAudit returning timestamp and possibility to omit payloads
- **pkg/api:** add field and remove field endpoints
- **pkg/database:** add user when creating collection
- **pkg/server:** add user when creating collection

### Reverts
- chore: remove initial swagger support


<a name="v1.5.0"></a>
## [v1.5.0] - 2023-06-20
### Bug Fixes
- **embedded/store:** handle replication of empty values

### Changes
- **embedded/document:** naming validations
- **embedded/document:** allow hyphen in doc naming
- **embedded/document:** collection and field naming validations
- **embedded/store:** embedded values and prealloc disabled by default


<a name="v1.5.0-RC1"></a>
## [v1.5.0-RC1] - 2023-06-16
### Bug Fixes
- build/Dockerfile.rndpass to reduce vulnerabilities
- modify tests for new object db initialisation
- build/Dockerfile.immuclient to reduce vulnerabilities
- build/Dockerfile.immuadmin to reduce vulnerabilities
- build/Dockerfile.immuadmin to reduce vulnerabilities
- build/Dockerfile.full to reduce vulnerabilities
- table id generation
- build/Dockerfile.rndpass to reduce vulnerabilities
- build/Dockerfile.full to reduce vulnerabilities
- build/Dockerfile.immuclient to reduce vulnerabilities
- **docs:** bump golang.org/x/net to 0.7.0 in docs and test pkg
- **embedded/ahtree:** correct calculation of payload offset
- **embedded/appendable:** proper closing of non-required chunks
- **embedded/document:** close readers before updating document
- **embedded/document:** proper column renaming
- **embedded/document:** assign correct revision number
- **embedded/document:** proper handling of deleted documents
- **embedded/document:** support nil docs
- **embedded/document:** id field conversion
- **embedded/document:** validate doc is properly initialized
- **embedded/document:** validate doc is properly initialized
- **embedded/sql:** parsing of exists stmt
- **embedded/sql:** multi-row conflict handling
- **embedded/sql:** like operator supporting null values
- **embedded/sql:** proper handling of parameters in row readers
- **embedded/sql:** do not force columns to have max key length when unspecified
- **embedded/sql:** implicit conversion within expressions
- **embedded/sql:** include explicit close into sqlTx options
- **embedded/sql:** consider 0 as no limit
- **embedded/sql:** crash when RowReader.Read() returns error
- **embedded/store:** force snapshot to include mandatory mvcc changes
- **embedded/store:** avoid dead-lock when exporting tx with external commit allowance mode
- **embedded/store:** ensure snapshot is closed for read-only txs
- **embedded/store:** integrity checks covering empty values
- **embedded/tbtree:** fix error comparison
- **embedded/tbtree:** proper kv validation
- **embedded/tbtree:** rollback to the most recent snapshot when insertion fails
- **embedded/tbtree:** fix snapshot getKeyWithPrefix
- **embedded/tbtree:** fix snapshot getKeyWithPrefix
- **go.mod:** bump go version to 1.17 in go.mod
- **helm:** set securityContext and podSecurityContext at correct location
- **pkg/api:** create collection endpoint with path parameter
- **pkg/api:** fix and implement LIKE and NOT_LIKE operator when querying documents
- **pkg/client:** ensure ticker is properly stopped
- **pkg/client:** return error when verifiedGet operation fails
- **pkg/database:** fix truncation and contemplate entry-less txs
- **pkg/database:** read-only document API for replicas
- **pkg/database:** skip eof error during scan
- **pkg/database:** wrap propagated context
- **pkg/database:** read from err channel
- **pkg/replicator:** check stream is properly initialized
- **pkg/server:** ensure tx is closed upon error
- **pkg/server:** request explicit close when creating a rw sql tx
- **pkg/server:** ensure error propagation when sending headers
- **pkg/server:** thread-safe doc reader during session handling
- **pkg/server:** close document readers before cancelling txs
- **pkg/server:** do not set trailer metadata when replication is done with bidirectional streamming
- **pkg/server:** use grpc interceptors with grpc proxy
- **pkg/stream:** handle the case when message fits in a single chunk
- **pkg/truncator:** adjust plan logic and contemplate empty txs
- **pkg/verification:** document comparison with proto equals
- **push.yml:** update min go version

### Changes
- exclude generated code from coverage
- TruncateDatabase endpoint should use the same ongoing Truncator if present
- add ReadN method to document reader
- rename DocumentBulkInsert to DocumentInsertMany
- generate proto requests for DocumentDelete api
- exclude generated code from coverage
- use sys/unix package
- remove docker test provider
- use sql statement for delete than raw query
- use gosec action
- add monotically increasing number to doc id generation
- add document audit api
- check invalid search id in search request
- wait for immudb to get initialized
- add DocumentDelete api
- add go-acc and goveralls to ext-tools folder
- copy document catalogue when truncating db
- handle no more doc error inside response in search
- return ErrNoMoreDocuments instead of sql.ErrNoMoreRows
- replace schemav2 with protomodel in truncator test
- add order by clause in search
- allow multiple order by clauses
- add unique search id for paginated readers
- add documentReader iterator to read documents
- add bulk insert api
- remove initial swagger support
- return sql reader on document search
- Update build/RELEASING.md file
- add option for non unique indexes on collection
- change DocumentFindOneAndUpdate to DocumentUpdate
- simplified codegen
- add pagination support when fetching documents
- address review comment
- update document with id if not nil
- pass transaction to upsert function
- add DocumentFindOneAndUpdate api
- add default size for document reader lru cache
- add lru cache for paginated readers
- Add reformatting of protobuf file on build/codegen
- fix merge issues
- fix failing verification test
- increase test coverage for document engine
- change DeleteTableStmt to DropTableStmt
- fix tests
- fix TestFloatSupport test case
- add test case for uncommitted tx not increasing table count
- add updatecollection api
- check for column before adding index on collection update
- delete columns on table deletion
- **ci:** improve notifications
- **cmd/immuadmin:** flag to specify the usage of embedded values
- **cmd/immuadmin:** modify truncation settings schema
- **cmd/immuadmin:** add truncate cmd to immuadmin
- **deps:** bump github.com/rs/xid from 1.3.0 to 1.5.0
- **deps:** bump securego/gosec from 2.14.0 to 2.15.0
- **deps:** bump github.com/golang/protobuf from 1.5.2 to 1.5.3
- **deps:** bump github.com/influxdata/influxdb-client-go/v2
- **deps:** bump github.com/spf13/viper from 1.12.0 to 1.15.0
- **deps:** bump golang.org/x/net from 0.8.0 to 0.9.0
- **deps:** bump golang.org/x/crypto
- **deps:** bump github.com/grpc-ecosystem/grpc-gateway/v2
- **deps:** bump aws-actions/configure-aws-credentials from 1 to 2
- **deps:** bump github.com/spf13/cobra from 1.2.1 to 1.6.1
- **deps:** bump github.com/rogpeppe/go-internal from 1.8.0 to 1.9.0
- **deps:** bump github.com/codenotary/immudb
- **deps:** bump google.golang.org/grpc from 1.46.2 to 1.54.0
- **deps:** bump github.com/stretchr/testify from 1.8.0 to 1.8.2
- **deps:** bump github.com/jaswdr/faker from 1.4.3 to 1.16.0
- **deps:** bump github.com/lib/pq from 1.10.7 to 1.10.9
- **deps:** bump github.com/lib/pq from 1.10.2 to 1.10.7
- **embedded/ahtree:** add inline comments
- **embedded/appendable:** use fdatasync when file is preallocated
- **embedded/appendable:** file syncing per os
- **embedded/appendable:** support file preallocation
- **embedded/appendable:** file syncing using fdatasync when available
- **embedded/appendable:** fsync freebsd
- **embedded/appendable:** minor improvements reading files
- **embedded/appendable:** automatic file creation only when appending
- **embedded/appendable:** metadats with putBool
- **embedded/document:** move source code into dedicated files
- **embedded/document:** add test cases for collection on doc engine
- **embedded/document:** transactional collection update
- **embedded/document:** improve error handling
- **embedded/document:** retrieval of raw document
- **embedded/document:** typo in error message
- **embedded/document:** raw document validation
- **embedded/document:** transactional collection and document creation
- **embedded/document:** use onclose callback to close the tx
- **embedded/document:** return struct when auditing document history
- **embedded/document:** add test to ensure key ordering in document during serialization
- **embedded/document:** blob type not yet supported
- **embedded/document:** catch key alredy exists  error
- **embedded/document:** catch tx read conflict error
- **embedded/document:** translate table already exists error
- **embedded/document:** minor var renaming
- **embedded/document:** minor code adjustments
- **embedded/document:** transactional document creation
- **embedded/document:** use query limit when searching
- **embedded/document:** add collection deletion api support
- **embedded/document:** wip continue with improvements
- **embedded/document:** wip continue with improvements
- **embedded/document:** wip continue with improvements
- **embedded/document:** wip improvements
- **embedded/document:** add float support for doc engine
- **embedded/document:** binary serialization of doc payload
- **embedded/document:** remove dead-code
- **embedded/document:** avoid public dependency on sql
- **embedded/document:** change querier from BinBoolExp to CmpBoolExp
- **embedded/document:** support null values in indexed attributes
- **embedded/document:** add variable length support for multiple types
- **embedded/document:** possibility to specify desc order when querying document history
- **embedded/document:** add tests for blob type
- **embedded/document:** improve error messages
- **embedded/document:** improve error messages
- **embedded/document:** ensure order by clauses are used when deleting and updating
- **embedded/document:** minor code simplification
- **embedded/document:** fix query stmt generator and add tests
- **embedded/document:** leverage sqlengine lazy index contraint evaluation
- **embedded/document:** add document id generation
- **embedded/htree:** allow creation of empty hash trees
- **embedded/object:** add document abstraction
- **embedded/object:** add collection/database statements
- **embedded/sql:** make sql engine generic for object store
- **embedded/sql:** ddl stmts register catalog mutation
- **embedded/sql:** implicit conversion from varchar to int and float types
- **embedded/sql:** lazy index contraint validation
- **embedded/sql:** validate total key length at index creation time
- **embedded/sql:** extend max key length to 512
- **embedded/sql:** limit and offset boundary validation
- **embedded/sql:** minor numeric type adjustments
- **embedded/sql:** implicit conversion support in limit and offset clauses
- **embedded/sql:** WIP singledb sql engine
- **embedded/sql:** simplified sql tx
- **embedded/sql:** cancellable row reader
- **embedded/sql:** transient context
- **embedded/sql:** use read-only txs whenever possible
- **embedded/sql:** return closed sql txs
- **embedded/sql:** snapshot reuse improvements
- **embedded/sql:** upgraded row reader
- **embedded/store:** minor changes after rebasing from master
- **embedded/store:** multi-tx unsafe mvcc
- **embedded/store:** multi-tx bulk indexing
- **embedded/store:** set tx as closed upon cancellation
- **embedded/store:** simplified indexer initialization
- **embedded/store:** set smaller default value for indexing bulk size
- **embedded/store:** inline comments
- **embedded/store:** contextualized transactions
- **embedded/store:** propagate context usage
- **embedded/store:** set ctx as first argument
- **embedded/store:** set ctx as first argument
- **embedded/store:** multi-timed bulk insertions
- **embedded/store:** tx header is returned when fully committed
- **embedded/store:** simplified dualproof implementation
- **embedded/store:** transient context
- **embedded/store:** added more in-line comments
- **embedded/store:** context propagation
- **embedded/store:** mvcc validations
- **embedded/store:** addition of a cache for values
- **embedded/store:** add min limit for truncation frequency
- **embedded/store:** safe key copy for mvcc validation
- **embedded/store:** add min limit for truncation frequency
- **embedded/store:** wip mvcc validations
- **embedded/store:** add hashValue as fixed 32 byte size
- **embedded/store:** fix rebase issue with readValueAt for vlogcache
- **embedded/store:** fix default vlog cache size and add validation for hash when reading from cache
- **embedded/store:** add test for TxOptions
- **embedded/store:** optional integrity checking
- **embedded/store:** embedded meta attribute required if version is greater than 1
- **embedded/store:** optional integrity checking when reading values
- **embedded/store:** api upgrade
- **embedded/store:** set embedded values mode as default one
- **embedded/store:** backward compatible embedded value mode
- **embedded/store:** skipIntegrityCheck parameter when reading data
- **embedded/store:** preallocate tx header log files
- **embedded/store:** support preallocated files when reading tx data
- **embedded/store:** wip preallocated clog
- **embedded/store:** option to prealloc files
- **embedded/store:** improve log messages when discarding precommitted transactions
- **embedded/store:** validate Eh only when integrity checks are not disabled
- **embedded/store:** consume all tx content even if integrity checks are disabled
- **embedded/store:** optional integrity checking when reading values
- **embedded/store:** clog file size adjustment only when preallocation is disabled
- **embedded/store:** handle eof when reading last committed tx
- **embedded/store:** validate Eh only when integrity checks are not disabled
- **embedded/store:** wip mvcc validations
- **embedded/store:** wip mvcc validations
- **embedded/store:** make truncation validation tolerate entryless txs
- **embedded/store:** allow tx without entries as long as it contains metadata
- **embedded/store:** update ReadBetween
- **embedded/store:** unify Read and ReadBetween
- **embedded/store:** use syncSnapshot to validate ongoing txs
- **embedded/store:** file preallocation not enabled by default
- **embedded/store:** further in-line documentation
- **embedded/store:** snapshot reuse improvements
- **embedded/store:** mvcc validation only if another tx was processed
- **embedded/store:** inline comments
- **embedded/store:** readValueAt and exportTx improvements
- **embedded/store:** minor code improvement
- **embedded/store:** fix typo in inline comment
- **embedded/store:** add in-line documentation for store options
- **embedded/store:** validate gets using filters
- **embedded/store:** MVCC read-set with boundaries
- **embedded/tbtree:** initialize tbtree with a non-mutated leaf
- **embedded/tbtree:** rollback not needed as updates are made in a copy
- **embedded/tbtree:** variable renaming after rebasing
- **embedded/tbtree:** add in-line documentation
- **embedded/tbtree:** parametrize snapshot creation specs
- **embedded/tbtree:** optimize snapshot renewal
- **embedded/tbtree:** getWithPrefix
- **embedded/tbtree:** add in-line comments
- **embedded/tbtree:** wip optimized insertion
- **embedded/tbtree:** optimized bulk insertion
- **embedded/tbtree:** wip reduce allocs while updating inner node
- **embedded/tbtree:** in-line documentation
- **embedded/tbtree:** minor code improvements
- **embedded/tbtree:** remove unnecessary kv sorting
- **embedded/tools:** upgrade embedded tools with transient context
- **embedded/watchers:** set ctx as first arg
- **embedded/watchers:** return context error upon cancellation
- **embedded/watchers:** use context instead of cancellation channel
- **package/database:** bunch of fixes and improvements in document engine
- **pkg:** add more tests admin truncate command
- **pkg/api:** remove generated httpclient
- **pkg/api:** use of path parameters for document-related endpoints
- **pkg/api:** search api improvements
- **pkg/api:** remove bool from tx metadata conversion
- **pkg/api:** swagger gen
- **pkg/api:** snapshot reuse attributes
- **pkg/api:** expose new store indexing options
- **pkg/api:** document api improvements
- **pkg/api:** revised document and authentication apis
- **pkg/api:** remove unsupported attribute from response messages
- **pkg/api:** manual adjustments post-code generation
- **pkg/api:** re-generated httpclient with DeleteDocument endpoint
- **pkg/api:** return txID when inserting or updating documents
- **pkg/api:** cleaner session id header
- **pkg/api:** document api improvements
- **pkg/api:** document update with path parameter
- **pkg/api:** rename idFieldName to documentIdFieldName
- **pkg/api:** expose MVCC read-set settings
- **pkg/api:** endpoint renaming
- **pkg/api:** expose replication settings for skipping integrity checks and indexing
- **pkg/api:** expose db setting to enable file preallocation
- **pkg/api:** add tx metadata conversion
- **pkg/api:** expose embeddedValue database setting
- **pkg/api:** authorization in swagger spec
- **pkg/api:** re-generated httpclient
- **pkg/api:** singular document path for audit and proof endpoints
- **pkg/api:** revert changes in swagger spec
- **pkg/api:** value cache settings exposed
- **pkg/api:** annotate primitive types as required
- **pkg/api:** buch of implementation improvements
- **pkg/api:** minor proof request renaming
- **pkg/api:** re-generated httpclient
- **pkg/api:** use ErrrIs/ErrorContains in error checks
- **pkg/api:** annotated required message fields
- **pkg/api:** expose support for unsafe mvcc transactions
- **pkg/api:** change retention period in TruncateDatabase message to int64
- **pkg/api:** annotate required fields
- **pkg/auth:** add document update permissions
- **pkg/client:** move heartbeater.go to pkg/client
- **pkg/client:** minor renaming in tx options
- **pkg/client/cache:** improve test coverage
- **pkg/database:** add object store
- **pkg/database:** add search document api implementation for object store
- **pkg/database:** create txs with default options
- **pkg/database:** implement GetCollection API
- **pkg/database:** change objectEngine to documentEngine
- **pkg/database:** add mvcc test for truncation, parse retention period using duration
- **pkg/database:** add more tests for truncation
- **pkg/database:** remove search through first query
- **pkg/database:** upgraded reader specs
- **pkg/database:** hard limit on page size
- **pkg/database:** fix truncation deletion point checks in test
- **pkg/database:** minor code aligments
- **pkg/database:** add document store db initialisation
- **pkg/database:** updated APIs with schema updates
- **pkg/database:** add document engine abstraction
- **pkg/database:** context propagation
- **pkg/database:** add and implement object db interface
- **pkg/database:** snapshot reuse changes
- **pkg/database:** create document/collection from schemav2 requests
- **pkg/database:** upgrade after rebasing
- **pkg/database:** use _obj to hold raw document payload
- **pkg/database:** context propagation from server to embedded layer
- **pkg/database:** remove object store db initialisation
- **pkg/database:** proper calculation of source tx
- **pkg/database:** document verfication
- **pkg/database:** add DocumentUpdate api
- **pkg/database:** check encoded value is consistent with raw document
- **pkg/database:** add query parser for object to generate sql expression
- **pkg/database:** add document query struct to abstract request query
- **pkg/database:** minor document renaming
- **pkg/integration:** exportTx benchmarking
- **pkg/replication:** replicator using bidirectional streaming
- **pkg/replication:** wip stream replication - only async replication working
- **pkg/replication:** skip integrity check when exporting transactions
- **pkg/replication:** context propagation
- **pkg/replication:** improve options validation
- **pkg/server:** upgrades after rebasing from master
- **pkg/server:** integrate document functions with server apis
- **pkg/server:** context propagation from grpc api to embedded package
- **pkg/server:** multi-grpc request context propagation
- **pkg/server:** minor code reuse
- **pkg/server:** add pagination test for document search
- **pkg/server:** add test successful load/unload of db with truncator
- **pkg/server:** log error when closing document reader
- **pkg/server:** ensure document reader is closed when swithing pages
- **pkg/server:** support snapshot reuse
- **pkg/server:** upgrade to new insecure credentials api
- **pkg/server:** close all paginated readers on close of session
- **pkg/server:** added inline comments
- **pkg/server:** set default replication settings
- **pkg/store:** skipIntegrityChecks parameter when reading data
- **pkg/stream:** handle eof when sending data
- **pkg/truncator:** return error if expiration time hasn't been met
- **pkg/truncator:** add context to Truncate method
- **pkg/truncator:** refactor truncator process
- **pkg/verfication:** document verification methods
- **pkg/verification:** strengthen proof validations
- **pkg/verification:** use proto serialization
- **pkg/verification:** minor renaming
- **pkg/verification:** document verification using embedded identifier
- **test/objects:** add tests to create collections
- **test/objects:** add more tests to create collection
- **test/objects:** use httpexpect
- **test/perf:** fix version value for flag
- **test/perf:** add immudb version to influxdb data
- **test/perf:** add runner to results for influxdb
- **test/perf-tests:** remove runner check
- **test/perf-tests:** use proxy on benchmark runner
- **test/performance:** call cleanup method
- **test/performance-test-suite:** send results to influxdb
- **test/performance-test-suite:** add influxdb host and toke arguments
- **test/performance-test-suite:** add sync benchmarks
- **test/performance-test-suite:** extract json from results
- **test/performance-test-suite:** fix replica directory path
- **test/performance-test-suite:** use temp folders for primary, replicas and clients
- **test/performance-test-suite:** replicas are able to communicate with primary
- **test/performance-test-suite:** changed server concrete implementation
- **truncator:** add more coverage for truncator

### Features
- add vlog truncation functionality
- **ci:** change notification
- **embedded/document:** count documents
- **embedded/object:** add object store to embedded pkg
- **embedded/sql:** limit and offset as expressions
- **embedded/sql:** wip unsafe and optimized mvcc
- **embedded/sql:** sql transaction creation with options
- **embedded/sql:** short casting syntax
- **embedded/sql:** implicit type conversion of numeric types
- **embedded/sql:** Initial float support
- **embedded/store:** unsafe mvcc mode
- **embedded/store:** embeddable values
- **embedded/store:** read-only transactions
- **embedded/store:** embedded values option
- **embedded/store:** tx creation with options
- **embedded/store:** expose GetWithPrefixAndFilters
- **embedded/store:** GetWithPrefixAndFilters
- **embedded/tbtree:** multi-timed bulk insertions
- **pkg/api:** keepOpen parameter to instruct server to maintain a document reader in memory
- **pkg/api:** improved replace documents endpoint
- **pkg/api:** count documents endpoint
- **pkg/api:** document proof endpoint
- **pkg/client:** optional tx options are now available during the creation process


<a name="v1.4.1"></a>
## [v1.4.1] - 2022-11-21
### Changes
- **pkg/server:** Add logs for activities related to users


<a name="v1.4.1-RC1"></a>
## [v1.4.1-RC1] - 2022-11-16
### Bug Fixes
- Change replication-related terms in tests
- Change replication-related terms in codebase
- **cmd:** Rename replication flags to follow consistent convention
- **cmd/immudb:** Fix description of the `force-admin-password` flag
- **cmd/immudb:** Better description of the `--force-admin-password` flag
- **embedded/appendable:** fsync parent directory
- **embedded/appendable:** fsync parent folder in remote appedable
- **pkg:** Rename replication-related fields in GRPC protocol
- **pkg/client:** Delay server identity validation
- **pkg/client/cache:** Add methods to validate server identity
- **pkg/client/cache:** Validate server's identity
- **pkg/server:** Remove includeDeactivated flag when querying for users
- **pkg/server/servertest:** Fix resetting grpc connection
- **pkg/server/servertest:** Add uuid to buffconn server
- **test/perf-test-suite:** Avoid dumping immudb logo on perf test results file
- **test/performance-test-suite:** Ensure results are shown after proper is finished
- **verification:** Additional Linear proof consistency check
- **verification:** Recreate linear advance proofs for older servers

### Changes
- **ci:** migrate deprecating set-output commands
- **cmd/immudb:** Allow resetting sysadmin password
- **docs/security:** Add resources for the linear-fake vulnerability
- **docs/security:** Be less specific about package version in examples
- **embedded/appendable:** sync directories
- **embedded/store:** Remove AHT Wait Hub
- **embedded/store:** Disable asynchronous AHT generation
- **pkg/client:** Document `WithDisableIdentityCheck` option
- **pkg/client/cache:** Limit the hash part of the identity file name
- **pkg/client/cache:** Describe serverIdentity parameter
- **pkg/client/state:** Cleanup mutex handling in StateService
- **pkg/server:** Warn if sysadmin user password was not reset
- **pkg/server:** Better warning for unchanged admin password
- **test/performance-test-suite:** Add summary to json output

### Features
- **ci:** fix message and input
- **ci:** add runner name to mattermost message header
- **ci:** simplify results extraction
- **ci:** extract performance tests into separate workflow to be reused
- **ci:** add scheduled daily test runs and send results to Mattermost
- **pkg/replication:** Disable server's identity check in internal replication


<a name="v1.4.0"></a>
## [v1.4.0] - 2022-10-12
### Bug Fixes
- **build:** Do not publish official non-dev images on RC tags
- **pkg/client:** replace keepAlive context from the original one to the background, avoiding parent expiration

### Changes
- Rename sync-followers to sync-acks
- **cmd/immuclient:** include precommit state when quering status
- **pkg/server:** Better error message when validating replication options


<a name="v1.4.0-RC2"></a>
## [v1.4.0-RC2] - 2022-10-10
### Bug Fixes
- **build:** Use correct binary download links
- **embedded/store:** edge-case calculation of precommitted tx
- **embedded/watchers:** Fix invariant breakage in watchers
- **embedded/watchers:** Fix invariant breakage in watchers
- **pkg/database:** any follower can do progress due to its prefech buffer
- **pkg/replication:** Do not crash on invalid tx metadata
- **pkg/replication:** handle replication already closed case
- **pkg/replication:** discard precommitted txs and continue from latest committed one
- **pkg/replication:** solve issues when follower diverged from master
- **wmbedded/watchers:** Correctly fix the original implementation

### Changes
- **embedded/watchers:** Simplify and document cancellation path
- **embedded/watchers:** Simplify mutex locking code
- **embedded/watchers:** single-point for init and cleanup
- **pkg/database:** simplify follower's wait
- **pkg/database:** wait for tx when a non-existent or non-ready transaction is requested
- **pkg/database:** add TODO comment on replication passive waiting
- **pkg/replication:** Add TX gap metrics
- **pkg/replication:** Add basic replication metrics
- **pkg/replication:** improve replication logging


<a name="v1.4.0-RC1"></a>
## [v1.4.0-RC1] - 2022-10-04
### Bug Fixes
- **Makefile:** add fips build flag to test/fips
- **Makefile:** remove interactive flag from dist/fips command
- **ci:** fix regex pattern for fips binaries
- **cmd/immuadmin:** set correct  data-type for replication-sync-followers flag
- **embedded/store:** Fix checking for closed store when syncing TXs
- **embedded/store:** avoid attempts to commit in wrong order
- **embedded/store:** expose durable precommitted state
- **embedded/store:** include allowPrecommitted into tx reader construction
- **embedded/store:** ensure tx is released upon error
- **embedded/store:** aht up to precommited tx
- **embedded/store:** fix size calculation of precommitted txs
- **github:** Update github actions after migration of Dockerfile's
- **pkg/database:** Fix mutex lock in ExportTx
- **pkg/database:** return master commit state if failing to read follower precommitted one
- **pkg/database:** set follower states holder when changing replication status
- **pkg/server:** add logs when replicator does not start

### Changes
- add dependabot config
- Add empty line between license header and package
- **Dockerfile.fips:** add fips build changes
- **cmd/immuadmin:** use default immudb port as default value for replication-master-port flag
- **cmd/immuadmin:** revert default replication-master-port
- **cmd/immuadmin:** add new replication flags
- **cmd/immuclient:** flag replication-sync-enabled to enable sync replication
- **cmd/immudb:** deprecate replication-enabled towards replication-is-replica
- **docker:** Move main Dockerfile's to build folder
- **docker:** Simplify the main Dockerfile
- **embedded/store:** waits for durable precommitted txs
- **embedded/store:** wip wait for precommitted txs
- **embedded/store:** mutexless export-tx
- **embedded/store:** method to dynamically switch to external allowance
- **embedded/store:** wip reduce allocations in exportTx
- **embedded/store:** enhanced tx discarding logic
- **embedded/store:** resolve pre-committed using clogbuf
- **embedded/store:** wip load precommitted txs
- **embedded/store:** minor code simplification
- **embedded/store:** handle commit case when there is nothing new to commit
- **embedded/store:** support for concurrent replicated precommits
- **embedded/store:** tx parsing with sanity checks
- **embedded/store:** add integrity checks when reading precommitted txs
- **embedded/store:** tolerate partial data or inconsistencies when loading pre-committed txs
- **embedded/store:** explanatory comments added
- **embedded/store:** explicit allowPrecommitted and restricted access to precommitted txs
- **embedded/store:** minor renaming and comment additions
- **embedded/store:** possibility to read tx header of precommitted txs
- **pkg/api:** explicit sync replication setting
- **pkg/api:** currentState endpoint includes precommitted info
- **pkg/api/schema:** reformat schema.proto file
- **pkg/database:** improve error comparison
- **pkg/database:** handle special case related to sql initialization
- **pkg/database:** follower commit progress without additional waits
- **pkg/database:** sync exportTx
- **pkg/database:** minor typo in comment
- **pkg/database:** disable automatic sql init on older databases
- **pkg/integration:** add synchronous replication integration tests
- **pkg/replication:** improve error comparison
- **pkg/replication:** handling a particular case in an optimized manner
- **pkg/replication:** speed up follower reconnection
- **pkg/replication:** replicator with backward compatibility mode
- **pkg/replication:** check committedTxID from master
- **pkg/replication:** graceful closing
- **pkg/replication:** use session-based authentication
- **pkg/replication:** handle case when follower precommit state is up-to-date but commit state is lies behind
- **pkg/replication:** sync replication using follower state
- **pkg/replication:** allowPreCommitted only with sync replication enabled
- **pkg/replication:** wip optimize concurrency in replicators
- **pkg/replication:** configurable prefetchTxBufferSize and replicationCommitConcurrency
- **pkg/replication:** further progress in sync replication
- **pkg/replication:** backward compatible replication
- **pkg/replicator:** wip precommitted tx discarding when follower diverged from master
- **pkg/server:** use replication settings
- **pkg/server:** include sync replication settings in options
- **pkg/server:** explicit sync replication
- **pkg/server:** support for systemdb with session-based auth
- **pkg/server:** display all replication settings
- **pkg/server:** handle admin user creation with sync replication enabled

### Features
- **cmd/immuadmin:** flag to set the number of sync followers
- **cmd/immudb:** flag to set the number of sync followers for systemdb and defaultdb
- **embedded/store:** functionality to discard precommitted txs
- **embedded/store:** core support for sync replication
- **pkg/api:** api extensions to support sync replication
- **pkg/database:** wip sync replication logic
- **pkg/replication:** mode to allow tx discarding on followers
- **pkg/replication:** wip replicator with support for sync replication
- **pkg/server:** sync replication logic
- **pkg/server:** Add ability to inject custom database management object


<a name="v1.3.2"></a>
## [v1.3.2] - 2022-08-25

<a name="v1.3.2-RC1"></a>
## [v1.3.2-RC1] - 2022-08-24
### Bug Fixes
- company name in webconsole and other files
- access tls value in global scope within ingress annotations
- **build:** update go version to 1.18 in Dockerfiles
- **build:** Fix go-acc and goveralls invocations
- **build/RELEASING.md:** Add note about updating playground
- **embedded:** use tmp folder for unit test cases
- **embedded/sql:** Support single `BEGIN` statement.
- **embedded/store:** Reduce the amount of allocations for tx object
- **embedded/store:** Return correct error on key length exceeded
- **embedded/store:** Ensure ordering of transaction timestamps
- **embedded/store:** Assign blTxID within locked tx state
- **embedded/store:** Optionally preallocate Tx pools
- **embedded/store:** Improved check for replicated transaction
- **embedded/store:** Protect against simultaneous replicators
- **embedded/store:** Check precommitted state when replicating
- **embedded/store:** ensure tx is released upon error
- **embedded/tools/stress_tool:** Fix compilation after recent update to tx holder pool
- **getRandomTable:** increase RNG range for table generation
- **github:** Remove unnecessary `/test/` path when uploading perf results to s3
- **github:** Do not use yaml anchors in github workflows
- **pkg/client:** Invalid client state after connection refused
- **pkg/client/clienttest:** enforce mock client to interface
- **pkg/database:** Fix calculation of proof for VerifiableTxByID
- **pkg/database:** Correct revision for Scan requirests
- **server:** Show info text with a logger
- **servertest:** Allow accessing Server object before starting the server
- **stdlib/rows:** add colums to row response
- **test/performance:** Cleanup test directory

### Changes
- deprecate ImmuClient.HealthCheck in favour of ServerInfo.
- ignore schema_grpc.pb.go in code coverage.
- regenerate with correct version of protoc-gen-go.
- update build constraint to new & future-proof syntax.
- format tools.go.
- pin google.golang.org/protobuf to v1.27.1 (currently used version for generated code).
- reimplement ImmuClient.HealthCheck using rpc ServerInfo instead of (deprecated) Health.
- pin github.com/pseudomuto/protoc-gen-doc to 1.4.1 (currently used version for generated code).
- ignore schema_grpc.pb.go in coveralls.
- refactor TestServerInfo.
- makefile formatting.
- update github.com/spf13/viper to v1.12.0.
- generate gRPC stubs.
- use go.mod version of github.com/grpc-ecosystem/grpc-gateway when building codegen.
- Update main go versin to 1.18
- Introduce separate TxHolder pools
- **Makefile:** Update webconsole to 1.0.16
- **build:** Update RELEASING.md doc
- **build:** Improve generation of build checksums
- **cmd/immuadmin:** Add support for max-commit-concurrency option
- **cmd/immuadmin:** Add support for read-tx-pool-size option
- **cmd/immudb:** Add support for max-sessions command line option
- **database/sql:** Delay txholder allocation on VerifiableSQLGet
- **embedded/ahtree:** flushless append
- **embedded/ahtree:** threshold-based sync
- **embedded/ahtree:** improve error handling
- **embedded/ahtree:** use bigger default write buffer size
- **embedded/ahtree:** improve error message consistency
- **embedded/ahtree:** support newst appendable implementation
- **embedded/ahtree:** improve validations and error handling
- **embedded/ahtree:** minor error message change
- **embedded/appendable:** auto-sync options
- **embedded/appendable:** improve explanatory comment inside sync method
- **embedded/appendable:** autosync when write buffer is full
- **embedded/appendable:** return io.EOF when offset is out of range
- **embedded/appendable:** multi-appendable shared write buffer
- **embedded/appendable:** autosync support in multi-appendable
- **embedded/appendable:** upgrade mocked and remote appendable based on new flushing assumptions
- **embedded/appendable:** inmem buffer offset
- **embedded/appendable:** flush when no more writes are done in appendable
- **embedded/appendable:** improve singleapp validation and error handling
- **embedded/appendable:** improve validations and error handling
- **embedded/appendable:** error tolerant seek
- **embedded/appendable:** wip remoteapp validation
- **embedded/htree:** improve error handling
- **embedded/sql:** Remove unnecessary tx holder buffer from SQLTx
- **embedded/store:** Add dedicated error for tx pool exhaustion
- **embedded/store:** Add txPoolOptions to setup pool parameters upon creation
- **embedded/store:** wip error declaration
- **embedded/store:** sync AHT before tx commit log
- **embedded/store:** in-mem clog buffer written when synced
- **embedded/store:** wrap internal already closed errors
- **embedded/store:** handle appendable already close error
- **embedded/store:** improve error comparison with errors.Is(...)
- **embedded/store:** multi-tx syncs
- **embedded/store:** Better errors returned during replication error
- **embedded/store:** Do not write values if concurrency limit is reached
- **embedded/store:** Use dedicated error for replication conflicts
- **embedded/store:** avoid sync waiting if there are no new transactions
- **embedded/store:** add TODO comment
- **embedded/store:** Add explicit ReadTxEntry method
- **embedded/store:** aht options
- **embedded/store:** set new default write buffer values
- **embedded/store:** use smaller default buffer size
- **embedded/store:** flush-less precommit
- **embedded/store:** wip retryable sync
- **embedded/store:** parametrize write buffer size
- **embedded/store:** Add explicit ReadTxHeader
- **embedded/store:** Optimize ReadTxEntry method
- **embedded/store:** Optimize ReadTxHeader method
- **embedded/store:** Add txDataReader to process transaction data
- **embedded/store/txpool:** Make txPoolOptions members private
- **embedded/store/txpool:** Allocate pool entries separately
- **embedded/tbtree:** use non-retryable sync
- **embedded/tbtree:** improve error handling
- **embedded/tbtree:** define using generic errors towards errors.Is(...) usage
- **embedded/watchers:** improve error handling
- **github:** Upload perf results to AWS s3
- **github:** Allow using multiple runners for perf test suite
- **github:** Run perf test suite on pull requests
- **github:** Run performance test suite on push to master
- **github:** Add simple documentation of `PERF_TEST_xxx` secrets
- **github:** Install qemu using docker/setup-qemu-action
- **github:** Allow selection of runner to run perf test
- **github:** Update ACTIONS_SECRETS.md file
- **pkg/api:** export syncFrequency database parameter
- **pkg/api:** milliseconds message type
- **pkg/api:** deprecate rpc Health in favour of ServerInfo.
- **pkg/api:** Add tx pool size to GRPC and stored db options
- **pkg/api:** expose aht settings
- **pkg/database:** Add tx pool size to db options
- **pkg/database:** allocate tx buffer before doing verified writes
- **pkg/database:** Remove txHolder from get operation
- **pkg/database:** Do not allocate txholder for history scans
- **pkg/logger:** Add memory logger
- **pkg/logger:** add json logger
- **pkg/server:** Add pprof option
- **pkg/server:** simplify ImmuServer.Health.
- **test/performance:** Add separate `Write KV/s` test.
- **test/performance:** Allow customized name for the benchmark
- **test/performance:** Move test seed out of configuration
- **test/performance:** Correctly close random data generator
- **test/performance:** Split benchmark list and run code
- **test/performance:** Add basic flags to the benchmark process
- **test/performance:** Move random generator and key tracker to common coode
- **test/performance:** Add CPU time / memory stats gathering
- **test/performance:** Better logging and output
- **test/performance:** Add basic IO stats
- **test/performance:** Improve live IO display

### Features
- revert usages of ServerInfo that would break backwards compatibility.
- add test for HealthCheck.
- **cmd/immuadmin:** expose syncFrequency and WriteBufferSize db parameters
- **cmd/immuclient:** add info command to immuclient.
- **pkg/api:** expose write buffer parameter
- **pkg/api:** improve documentation of ServerInfo.
- **pkg/api:** remove ServerInfoResponse.status field.
- **pkg/api:** add ServerInfo rpc to deprecate Health.
- **pkg/client:** revert WaitForHealthCheck change to maintain backwards-compatibility.
- **pkg/client:** implement ImmuClient.ServerInfo.
- **pkg/server:** implement ImmuServer.ServerInfo.


<a name="v1.3.1"></a>
## [v1.3.1] - 2022-06-30
### Bug Fixes
- **embedded/store:** filter evaluation after valRef resolution

### Changes
- **embedded/store:** offset handling at keyreader

### Features
- **embedded/sql:** offset clause
- **embedded/store:** offset in key scanning
- **pkg/api:** offset attribute in scan and zscan endpoints


<a name="v1.3.1-RC1"></a>
## [v1.3.1-RC1] - 2022-06-30
### Bug Fixes
- **README:** Update readme to show examples for 1.3.0 version
- **cmd/immuadmin:** use StreamChunkSize as max chunk size during tx replication
- **cmd/immudb:** include metrics endpoint related flags
- **embedded/remotestorage:** Fix invalid comment
- **embedded/remotestorage/s3:** Fix s3 object name validation
- **embedded/remotestorage/s3:** Correctly url decode entry names
- **embedded/remotestorage/s3:** Simplify the code for scan
- **embedded/remotestorage/s3:** Avoid using HEAD requests
- **embedded/sql:** Use defer to cleanup unclosed readers on error
- **embedded/sql:** Fix snapshot leak on query initialization failure
- **embedded/sql:** Fix reader leaks during initialization failures
- **embedded/sql:** Properly close readers in joint row reader
- **embedded/sql:** Fix snapshot leaks in union readers
- **embedded/sql:** ensure timestamp is evaluated with microsecond precision
- **pkg/client:** ensure connection is closed and session can be re-established
- **pkg/database:** Do not panic if incorrect number of pk values is given to VerifiableSQLGet
- **pkg/server:** Fix remote storage test after recent changes
- **pkg/server/sessions:** Correctly start session guard
- **pkg/server/sessions:** Avoid deadlock when closing session manager
- **pkg/server/sessions:** Session manager test fixes
- **pkg/server/sessions:** Use strong random source for session ID
- **pkg/server/sessions:** Handle short buffer read when generating session id

### Changes
- Update dependencies
- **build:** Update RELEASING.md file
- **embedded/remotestorage:** More detailed errors
- **embedded/remotestorage:** Improve error reporting
- **embedded/remotestorage:** Improve testing of remotestorage
- **embedded/remotestorage/s3:** Improved s3 object name checks
- **embedded/sql:** fixed-timed tx
- **embedded/sql:** Do not return error from conditional and limit readers
- **github:** Update minimal supported go version to 1.15
- **github:** Run tests with minio service
- **github:** On macOS run client only test on pull requests
- **github:** Update push action
- **github:** Run coverage tests with minio enabled
- **pkg/client:** Better detection of tests that require external immudb
- **pkg/server:** Add missing copyright headers
- **pkg/server/session:** Move options normalization into options struct
- **pkg/server/sessions:** Simplify session handling code
- **pkg/server/sessions:** Add MaxSessions option
- **pkg/server/sessions:** Improve options handling
- **remotestorage:** Add prometheus metrics for remote storage kind
- **tools:** Remove old stream tool


<a name="v1.3.0"></a>
## [v1.3.0] - 2022-05-23
### Bug Fixes
- **embedded/sql:** return invalid value when using aggregated col selector in temporal queries
- **pkg/client:** enhance client-side validations in verified methods


<a name="v1.3.0-RC1"></a>
## [v1.3.0-RC1] - 2022-05-20
### Bug Fixes
- **cmd/immuclient:** Do not crash on login prompt
- **embedded/sql:** selector resolution using valuesRowReader
- **embedded/sql:** continue stmt execution on handler after changing db in use
- **embedded/sql:** increase auto_increment pk once per row
- **embedded/sql:** typo in error message
- **embedded/sql:** adjust named parameter parsing
- **github:** Run sonarcloud code analysis after cove coverate
- **pkg/database:** avoid silent returns when the scan limit is reached
- **pkg/database:** Fix detection of incorrect revision numbers
- **pkg/database:** Correctly interpret negative revision for getAt

### Changes
- **Dockerfile:** Add EXPOSE 5432 and IMMUDB_PGSQL_SERVER to all immudb images
- **README.md:** Switch to github badge
- **build:** Update the RELEASING.md documentation
- **cmd/immuclient:** Move history command to a separate file
- **cmd/immuclient:** Remove unnecessary sleep for set commands
- **cmd/immuclient:** Extract separate immuclient options
- **embedded/sql:** functional-style catalog queries
- **embedded/sql:** unit testing db selection
- **embedded/sql:** not showing unexistent db name as part of error message
- **embedded/sql:** fully non-transactional db creation and selection
- **embedded/sql:** wip grammar extensions to enrich temporal queries
- **embedded/sql:** de-duplicate error handling
- **embedded/sql:** database selection without multidb handler is still transactional
- **embedded/sql:** database selection as  non-transactional
- **embedded/sql:** validate current database as first step
- **embedded/sql:** param substitution in functional datasource
- **embedded/sql:** detailed error messages
- **embedded/sql:** quoted identifiers
- **embedded/sql:** ensure db selection is the last operation
- **embedded/sql:** implicit time expression
- **embedded/sql:** include short database selection stmt
- **embedded/sql:** ensure context propagation with multiple txs
- **embedded/sql:** re-include ttimestamp conversions in tx periods
- **embedded/sql:** postpone period evaluation so to support parameters type inference
- **embedded/sql:** non-functional catalog access
- **embedded/sql:** check tx range edge cases
- **embedded/sql:** sql tx with context
- **embedded/sql:** multi-db handler
- **embedded/sql:** functional catalog api
- **embedded/store:** minor refactoring time-based tx lookup
- **github:** Speedup push github actions
- **grpc:** Extend Scan API with endKey, inclusiveSeek, inclusiveEnd
- **pkg/api:** extend database creation response to indicate db already existed
- **pkg/database:** set multi-db handler after db initialization
- **pkg/database:** Rename getAt to getAtTx
- **pkg/database:** Improved checking of KeyRequest constraints
- **pkg/database:** databases catalog query yet unsupported
- **pkg/database:** contextual sql tx
- **pkg/database:** provide query parameters during resolution
- **pkg/database:** maintain MaxKeyScanLimit for backward compatibility
- **pkg/database:** Add missing copyright header in scan_test.go
- **pkg/database:** minor error renaming
- **pkg/integration:** Cleanup and restructure SQL tests
- **pkg/integration:** Add SQL verify tests after ALTER TABLE
- **pkg/server:** upgrade database method signature
- **pkg/server:** contextual sql tx

### Features
- Calculate revision number when scanning key history
- Add revision number when getting DB entries
- **api/schema:** Add revision-based option to key query
- **cmd/immuclient:** Add restore operation
- **cmd/immuclient:** Add support for revision-based get in immuclient
- **cmd/immuclient:** Add revision numbers when looking up key history
- **cmd/immuclient:** Better error messages for invalid revision for restore command
- **embedded/sql:** catalog queries
- **embedded/sql:** Implement ALTER TABLE ADD COLUMN
- **embedded/sql:** create database if not exists
- **embedded/sql:** WIP - UNION operator
- **embedded/sql:** temporal row ranges
- **embedded/sql:** queries with temporal ranges
- **embedded/store:** time-based tx lookup
- **embedded/store:** ranged key update reading
- **pkg/client:** Add revision-based get request on the go client
- **pkg/database:** Add revision-based get request on the GRPC level
- **pkg/server:** support database creation from sql
- **pkg/server:** support database selection from sql stmt


<a name="v1.2.4"></a>
## [v1.2.4] - 2022-04-28

<a name="v1.2.4-RC1"></a>
## [v1.2.4-RC1] - 2022-04-27
### Bug Fixes
- **Dockerfile:** Fix HOME variable for podman
- **cmd/immuclient:** upgrade not logged in error handling
- **embedded/tbtree:** Better logging in btree flush
- **embedded/tbtree:** Fix cleanupPercentage in SnapshotSince call
- **embedded/tbtree:** ensure node split is evaluated
- **embedded/tbtree:** create nodes with the right number of children
- **embedded/tbtree:** split into multiple nodes
- **github/push:** Fix notarization of binaries
- **pkg/auth:** Clarify comments about token injection
- **pkg/auth:** Do not send duplicated authorization header
- **pkg/server:** include db name in flush index result

### Changes
- **CHANGELOG.md:** remove bogus `liist` tag entry
- **build/RELEASING.md:** Update releasing docs
- **cmd/immuclient:** include db name when printing current state
- **embedded/store:** index settings validations
- **embedded/tbtree:** rename function that calculates node size lower bound
- **embedded/tbtree:** ensure node size is consistent with key and value sizes
- **github:** Update github workflow on master / version push
- **github:** Update github action versions
- **github:** Use smaller 5-days retention for master builds
- **github/push:** Add quick test linux-amd64 binaries
- **github/push:** Build, test and notarize for release/v* branches
- **github/push:** Calcualte sha256 checksums for binaries in github
- **github/push:** Build docker images after tests
- **github/push:** Add quick test for Mac x64 binaries
- **github/push:** Add quick test for linux-arm64 binaries through qemu
- **github/push:** Add quick test for linux-s390x binaries through qemu
- **github/push:** Run stress test before notarizing binaries
- **pkg/api:** txbyid with keepReferencesUnresolved option
- **tools/testing:** Add randomized key length mode for stress test tool
- **tools/testing:** Add stress tool


<a name="v1.2.3"></a>
## [v1.2.3] - 2022-04-14
### Bug Fixes
- **cmd/immuadmin:** simplify logging when flushing and compacting current db
- **pkg/database:** return key not found when resolving a deleted entry
- **pkg/database:** Return correct error for verifiedGet on deleted entries


<a name="v1.2.3-RC1"></a>
## [v1.2.3-RC1] - 2022-04-13
### Bug Fixes
- **CI/CD:** Golang compiler is not needed for building docker images
- **CI/CD:** Use CAS for notarization
- **embedded/store:** Fix early precondition checks
- **embedded/store:** Ensure up-to-date index on constrained writes
- **embedded/tbtree:** copy-on-write when increasing root ts
- **immudb:** Fix the name of signing key env var
- **pkg:** Fix tests after recent changes in API
- **pkg/api:** typo in kv metadata message
- **pkg/api:** Remove unused Sync field from IndexOptions
- **pkg/api/schema:** Use correct id for preconditions in SetRequest
- **pkg/auth:** Avoid unguarded read from user tokens map
- **pkg/client:** Adopt to EncodeReference changes
- **pkg/client:** Prevent updates with incorrect database settings
- **pkg/client:** Use correct response for UpdateDatabaseV2
- **pkg/client/errors:** Update the list of error codes
- **pkg/client/errors:** Ensure FromErrors works with ImmuError instance
- **pkg/database:** Better handling of invalid constraints
- **pkg/database:** Improve test coverage for KV constraints
- **pkg/database:** automatically set max score if not specified in desc order
- **pkg/errors:** Correct GRPC error mapping for precondition failure
- **pkg/server:** Use buffered channel for catching OS signals
- **pkg/server:** typo in log message
- **pkg/server:** adjust time to millis convertion
- **pkg/server:** ensure sessions locks get released
- **pkg/server:** override default settings with existent values
- **tools/monitoring:** Update grafana dashboards

### Changes
- cleanup percentage as float value
- Fix typo in a comment
- Rename Constraints to Preconditions
- Update copyright to 2022
- **Dockerfile:** Improve dockerfile builds
- **build:** improve release instructions ([#1100](https://github.com/vchain-us/immudb/issues/1100))
- **cmd/immuadmin:** add safety flag in delete database command
- **cmd/immuclient:** health command description
- **embedded/ahtree:** fix error message
- **embedded/appendable:** return io.EOF if there are not enough data for checksum calculation
- **embedded/appendable:** fix typo in error message
- **embedded/appendable:** discard capability
- **embedded/appendable:** appendable checksum calculation
- **embedded/store:** Add missing Copyright header
- **embedded/store:** add synced setting in index options
- **embedded/store:** do not skip expired entries when indexing
- **embedded/store:** verbose data corruption error
- **embedded/store:** parametrized index write buffer size
- **embedded/store:** syncThld at store options
- **embedded/store:** verbose logging during compaction
- **embedded/store:** index one tx per iteration
- **embedded/store:** improve compaction logging
- **embedded/store:** sync aht when syncing the store
- **embedded/store:** skip expired entries during indexing
- **embedded/store:** declare constants for all the options
- **embedded/store:** use store layer for constraint validations
- **embedded/store:** constraint validations with deletion and expiration support
- **embedded/store/options:** Simplify validation tests
- **embedded/tbree:** use shared writeOpts
- **embedded/tbree:** only insert nodes in cache when they were mutated
- **embedded/tbree:** remove obsolete property
- **embedded/tbree:** bump index version
- **embedded/tbtree:** middle node split
- **embedded/tbtree:** Add more internal metrics
- **embedded/tbtree:** min offset handling
- **embedded/tbtree:** validate compaction target path
- **embedded/tbtree:** data discarding with opened and older snapshots
- **embedded/tbtree:** Extend buckets for child node count histogram
- **embedded/tbtree:** synced flush if cleanup percentage is greater than zero
- **embedded/tbtree:** discard unreferenced data after sync
- **embedded/tbtree:** reduce allocs during flush
- **embedded/tbtree:** minOffset only for non-mutated nodes
- **embedded/tbtree:** ensure current snapshot is synced for compaction
- **embedded/tbtree:** validate input kv pairs before sorting
- **embedded/tbtree:** improve snapshot loading and discarding
- **embedded/tbtree:** discard unreferenced data when flushing index
- **embedded/tbtree:** discard unreferenced data
- **embedded/tbtree:** Add metrics for index data size
- **embedded/tbtree:** reduce allocations when flushing
- **embedded/tbtree:** positive compaction threshold
- **embedded/tbtree:** rebase non-indexed on kv syncthreshold
- **embedded/tbtree:** explicit error validation before loading
- **embedded/tbtree:** sort kv pairs in bulkInsert
- **embedded/tbtree:** checksum-based snapshot consistency validation
- **embedded/tbtree:** self-healing index
- **embedded/tbtree:** set initial offsets during initialization
- **embedded/tbtree:** validate data-format version
- **embedded/tbtree:** use double for min offset calculation
- **embedded/tbtree:** reduce fixed records length
- **embedded/tbtree:** open-ranged nodes
- **embedded/tbtree:** wip reduce node size
- **embedded/tbtree:** ensure sync on gracefully closing
- **embedded/tbtree:** fully replace sync with syncThld
- **embedded/tbtree:** use binary search during key lookups
- **embedded/tbtree:** Use KV entries count for sync threshold
- **embedded/tbtree:** no cache update during compaction reads
- **makefile:** fix cas sign instructions
- **metrics:** Add better flush / compaction metrics for btree
- **pkg/api:** use entries spec in verified and scan tx endpoints
- **pkg/api:** add synced param to flushindex endpoint
- **pkg/api:** non-indexable entries
- **pkg/api:** change proto schema toward db loading/unloading
- **pkg/api:** parametrized index cleanup percentage
- **pkg/api:** db loading and unloading
- **pkg/api:** uniform v2 endpoints
- **pkg/api:** prepare flushindex endpoint for future extensions
- **pkg/api:** use nullable prefix in db settings message
- **pkg/client:** optional client connection
- **pkg/client:** use txRequest in TxByIDWithSpec method
- **pkg/client:** synced flushing to enable physical data deletion
- **pkg/database:** parameters to resolve references at tx
- **pkg/database:** tx entries excluded by default if non-null spec is provided
- **pkg/database:** optional tx value resolution
- **pkg/database:** use shared tx holder when resolving tx entries
- **pkg/database:** remove db name from options
- **pkg/integration:** integrate non-indexed into grpc apis
- **pkg/server:** endpoint to retrieve settings of selected database
- **pkg/server:** expose max opened files for btree indexing
- **pkg/server:** use syncThreshold
- **pkg/server:** replication options for systemdb and defaultdb
- **pkg/server:** use previous store default values
- **pkg/server:** increase default max number of active snapshots
- **pkg/server:** tolerate failed user-created db loading
- **pkg/server:** expose flush index endpoint
- **pkg/server:** start/stop replicator when loading/unloading db
- **pkg/server:** convert time to milliseconds
- **pkg/server:** minor changes
- **pkg/server:** log web-console error on boot
- **pkg/server:** synced db runtime operations
- **pkg/server:** Dump used db options when loading databases
- **pkg/serverr:** validate request in deprecated database creation endpoint
- **stats:** Add btree cache prometheus stats

### Features
- Improved API for database creation and update
- Improved validation of kv constraints
- Entries-independent constraints in GRPC api
- Move KV write constraints to OngoingTX member
- **KV:** Do not create unnecessary snapshots when checking KV constraints
- **KV:** Add constrained KV writes for Set operation
- **KV:** Add constrained KV writes for Reference operation
- **KV:** Add constrained KV writes for ExecAll operation
- **KV:** Move constraints validation to OngoingTx
- **embedded/cache:** dynamic cache resizing
- **embedded/store:** Fail to write metadata if proof version does not support it
- **embedded/store:** Add tests for generation of entries with metadata
- **embedded/store:** non-indexable entries
- **embedded/store:** Add max header version used during writes
- **embedded/store:** Allow changing tx header value using GRPC api.
- **embedded/tbtree:** decouple flush and sync by introducing syncThreshold attribute
- **immuadmin:** Allow changing proof compatibility from immuadmin
- **kv:** Update grpc protocol with KV set constraints
- **pkg/api:** tx api with entry filtering capabilities
- **pkg/api:** delete database endpoint
- **pkg/client:** new method to fetch tx entries in a single call
- **pkg/database:** Updated GRPC protocol for constrained writes
- **pkg/database:** add noWait attribute in get request
- **pkg/database:** Update code for new constrained write protocol
- **pkg/server:** database health endpoint
- **pkg/server:** support database loading/unloading
- **pkg/server:** new endpoint databaseSettings
- **pkg/server:** expose all database settings
- **tools/monitoring:** added datasource selection, added instance selection, labels include instance, fixed calculations
- **tools/monitoring:** Add immudb Grafana dashboard


<a name="v1.2.2"></a>
## [v1.2.2] - 2022-01-18
### Bug Fixes
- registering connection in order to make possible conn recycling
- **Dockerfile:** Add ca-certificates.crt file to immudb image
- **client/file_cache:** Fix storing immudb state in file cache
- **embedded/immustore:** Avoid deadlock when acquire vLog lock
- **embedded/sql:** max key len validations
- **embedded/sql:** consider not null flag is on for auto incremental column
- **pkg/server:** validate if db is not replica then other replication attributes are not set
- **pkg/stdlib:** fix last insert id generation

### Changes
- create code of conduct markdown file ([#1051](https://github.com/vchain-us/immudb/issues/1051))
- **cmd/immuclient:** return actual login error
- **embedded/sql:** wip client provided auto-incremental values
- **embedded/sql:** change column constraints ordering
- **embedded/sql:** wip client provided auto-incremental values
- **embedded/sql:** add first and last insert pks retrivial methods
- **embedded/sql:** wip client provided auto-incremental values
- **metrics:** Add indexer metrics
- **metrics:** Add more s3-related metrics
- **pkg/database:** temporarily disable execall validations
- **pkg/database:** pre-validation of duplicated entries in execAll operation
- **pkg/database:** instantiate tx holder only in safe mode
- **pkg/database:** self-contained noWait execAll
- **pkg/database:** descriptive error messages
- **pkg/replication:** delay replication after failure
- **pkg/stdlib:** clean connection registration and leftovers

### Features
- **embedded/sql:** support for basic insert conflict handling
- **s3:** Add support for AWS V4 signatures


<a name="v1.2.1"></a>
## [v1.2.1] - 2021-12-14
### Bug Fixes
- fix interactive use database
- **embedded/store:** change already closed error message
- **embedded/store:** readonly tx entries to ensure no runtime modification
- **embedded/store:** reserve 4bytes in buffers for nentries
- **embedded/tbtree:** set fixed snapshot ts
- **pkg/server/sessions:** remove transaction on read conflict error
- **pkg/server/sessions/internal/transactions:** transaction is cleared after sqlExec error
- **sql:** Do not panic on error during delete
- **tx:** Remove summary from metadata

### Changes
- **embedded/store:** txmetdata placeholder with zero len
- **embedded/store:** private readonly metadata is validated when reading data
- **embedded/store:** read-only kv metadata for committed entries
- **embedded/store:** rw and readonly kv metadata
- **pkg/api:** use new kvmetadata api
- **pkg/client:** tx read conflict error is mapped in an CodInFailedSqlTransaction
- **pkg/server/sessions/internal/transactions:** defer only when needed
- **pkg/stdlib:** clean tx after rollback
- **pkg/stdlib:** fix connection creation
- **server/sessions:** modify read conflict error message

### Features
- **pkg/stdlib:** expose tx on std lib


<a name="v1.2.0"></a>
## [v1.2.0] - 2021-12-10
### Bug Fixes
- **database:** Internal consistency check on data reads
- **database/meta:** Do not crash on history with deleted items
- **pkg/database:** history skipping not found entries
- **protobuf:** Fix compatibility with 1.1 version

### Changes
- **cmd/immuadmin/command:** add super user login hint
- **embedded/sql:** use sql standard escaping with single quotes
- **embedded/sql:** support for escaped strings
- **embedded/store:** reduce attribute code size
- **embedded/store:** mandatory expiration filter
- **embedded/store:** fix expiration error declaration
- **embedded/store:** dedicated expiration error
- **embedded/store:** improve metadata serialization/deserialization methods
- **embedded/store:** validations during metadata deserialization
- **embedded/store:** return data corrupted error when deserialization cannot proceed
- **embedded/store:** easily extendable meta attributes
- **embedded/store:** use fixed time during the lifespan of a tx
- **embedded/store:** prevent value reading of expired entries
- **makefile:** remove windows binaries digital signature
- **pkg/auth:** require admin permission to export and replicate txs
- **pkg/integration:** remove useless compilation tag on tests
- **pkg/server:** deprecate GetAuth and WithAuth
- **pkg/server/sessions:** session max inactivity time set to 3m and minor stat collecting fix
- **pkg/server/sessions:** tuning sessions params
- **pkg/server/sessions:** session timeout set to 2 min

### Features
- **embedded/store:** logical entries expiration
- **pkg/api:** logical entries expiration
- **pkg/client:** expirable set


<a name="v1.2.0-RC1"></a>
## [v1.2.0-RC1] - 2021-12-07
### Bug Fixes
- Update jaswdr/faker to v1.4.3 to fix build on 32-bit systems
- **Makefile:** Use correct version of the grpc-gateway package
- **Makefile:** Fix building immudb for specific os/arch
- **embedded/sql:** normalize parameters with lower case identifiers
- **embedded/sql:** Do not modify value returned by colsBySelector
- **embedded/sql:** correct max key length validation based on specified col max length
- **embedded/sql:** fix rollback stmt
- **embedded/sql:** distinct row reader with limit argument
- **embedded/sql:** ensure determinism and no value overlaps distinct rows
- **embedded/sql:** Use correct statement for subquery
- **embedded/sql:** param substitution in LIKE expression
- **embedded/sql:** Fix SELECT * when joining with subquery
- **embedded/sql:** fix inserting calculated null values
- **embedded/store:** release lock when tx has a conflict
- **embedded/store:** read conflict validation
- **embedded/store:** typo in error message
- **pkg/auth:** Fix password tests
- **pkg/client:** fix database name saving on token service
- **pkg/database:** sql exec on provided tx
- **pkg/server:** fix keep alive session interceptor
- **testing:** using pointers for job channels
- **webconsole:** Fix html of the default missing page.

### Changes
- Update build/RELEASING.md documentation.
- remove token service from client options and fix tests
- token is handled internally by sdk. Remove useless code
- refining sdk client constructor and add readOnly tx guard
- fix more tests
- decoupled token service
- **cmd/immuadmin/command:** fix immuadmin token name on client creation
- **cmd/immuclient:** deleteKeys functioin and updates after metadata-related changes
- **cmd/immuclient:** temporary disable displaying hash in non-verified methods
- **embeddded/tbtree:** leverage snapshot id to identify it's the current unflushed one
- **embedded/multierr:** minor code simplification
- **embedded/sql:** rollback token
- **embedded/sql:** changes on tx closing
- **embedded/sql:** postponing short-circuit evaluation for safetiness
- **embedded/sql:** set INNER as default join type
- **embedded/sql:** Better error messages when (up|in)serting data
- **embedded/sql:** minor code simplification
- **embedded/sql:** standardized datasource aliasing
- **embedded/sql:** remove opt_unique rule to ensure proper error message
- **embedded/sql:** fix nullable values handling
- **embedded/sql:** use order type in scanSpecs
- **embedded/sql:** kept last snapshot open
- **embedded/sql:** de-duplicate tx attributes using tx header struct
- **embedded/sql:** unsafe snapshot without flushing
- **embedded/sql:** wip sql tx preparation
- **embedded/sql:** set parsing verbose mode when instantiating sql engine
- **embedded/sql:** reusable index entries and ignore deleted index entries
- **embedded/sql:** limit row reader
- **embedded/sql:** delay index sync until fetching row by its pk
- **embedded/sql:** leverage metadata for logical deletion
- **embedded/sql:** Simplify row_reader key selection
- **embedded/sql:** use int type for limit arg
- **embedded/sql:** minor update after rebasing
- **embedded/sql:** cancel non-closed tx
- **embedded/sql:** expose Cancel method
- **embedded/sql:** sql engine options and validations
- **embedded/sql:** Alter index key prefixes
- **embedded/sql:** return map with last inserted pks
- **embedded/sql:** wip rw transactions
- **embedded/sql:** defer execution of onClose callback
- **embedded/sql:** wip sqlTx
- **embedded/sql:** standard count(*)
- **embedded/sql:** ddl stmts not counted in updatedRows
- **embedded/sql:** method to return sql catalog
- **embedded/sql:** non-thread safe tx
- **embedded/sql:** wip interactive sqltx
- **embedded/sql:** bound stmt execution to a single sqltx
- **embedded/sql:** use current db from ongoing tx
- **embedded/store:** expose ExistKeyWithPrefix in OngoingTx
- **embedded/store:** conservative read conflict validation
- **embedded/store:** non-thread safe ongoing tx
- **embedded/store:** reorder tx validations
- **embedded/store:** wip tx header versioning
- **embedded/store:** simplified ExistKeyWithPrefix in current snapshot
- **embedded/store:** set tx as closed even on failed attempts
- **embedded/store:** strengthen tx validations
- **embedded/store:** GetWith method accepting filters
- **embedded/store:** set header version at commit time
- **embedded/store:** remove currentShapshot method
- **embedded/store:** entryDigest calculation including key len
- **embedded/store:** threadsafe tx
- **embedded/store:** handle watchersHub closing error
- **embedded/store:** ongoing tx api
- **embedded/store:** tx header version validations and increased max number of entries
- **embedded/store:** early tx conflict checking
- **embedded/store:** filter out entries when filter evals to true
- **embedded/tbtree:** remove ts from snapshot struct
- **embedded/tools:** upgrade sql stress tool
- **embedded/tools:** upgrade stress tool using write-only txs
- **embedded/tools:** update stress_tool after metadata-related changes
- **pkg/api:** changes in specs to include new metadata records
- **pkg/api:** consider nil case during tx header serialization
- **pkg/api/schema:** increase supported types when converting to sql values
- **pkg/client:** check if token is present before injecting it
- **pkg/client:** omit deleted flag during value decoding
- **pkg/client:** updates after metadata-related changes
- **pkg/client:** avoid useless tokenservice call and add tests
- **pkg/client/clienttest:** fix immuclient mock
- **pkg/client/tokenservice:** handlig error properly on token interceptor and fix leftovers
- **pkg/database:** return a specific error in querying
- **pkg/database:** snapshots should be up to current committed tx
- **pkg/database:** improve readability of Database interface
- **pkg/database:** limit query len result
- **pkg/database:** updates after metadata-related changes
- **pkg/database:** use new transaction support
- **pkg/database:** implement current functionality with new tx supportt
- **pkg/database:** revised locking so to ensure gracefully closing
- **pkg/database:** enforce verifiableSQLGet param validation
- **pkg/errors:** useDatabase returns notFound code when error
- **pkg/errors:**  invalid database name error converted to immuerror
- **pkg/integration:** updates after metadata-related changes
- **pkg/server:** use upgraded database apis
- **pkg/server:** updates after metadata-related changes
- **pkg/server:** error when tx are not closed
- **pkg/server/sessions:** polish logger call
- **pkg/server/sessions:** add sessions counter debug messages
- **pkg/stdlib:** remove context injection when query or exec
- **pkg/stdlib:** fix unit testing
- **pkg/stdlib:** increase pointer values handling and testing
- **pkg/stdlib:** general improvements and polishments
- **pkg/stdlib:** handling nil pointers when converting to immudb named params
- **pkg/stdlib:** improve connection handling and allow ssl mode in connection string
- **stress_tool_sql:** add sessions and transaction mode
- **stress_tool_worker_pool:** add long running stress tool
- **test:** test backward compatibility with previous release (v1.1.0)
- **tsting:** add index compactor in long running stress tool

### Features
- helm chart for deploying immudb on kubernetes ([#997](https://github.com/vchain-us/immudb/issues/997))
- **embedded/appendable:** method for reading short unsigned integer
- **embedded/sql:** null values for secondary indexes
- **embedded/sql:** support for IN clause
- **embedded/sql:** support for not like
- **embedded/sql:** WIP un-restricted upsert
- **embedded/sql:** engine as tx executor
- **embedded/sql:** wip sqltx at engine with autocommit
- **embedded/sql:** increased expression power in LIKE and IN clauses
- **embedded/sql:** Detect ambigous selectons on joins
- **embedded/sql:** distinct row reader
- **embedded/sql:** delete from statement
- **embedded/sql:** sql update statement
- **embedded/sql:** support value expression in like pattern
- **embedded/sql:** create index if not exists
- **embedded/store:** initial commit towards full tx support
- **embedded/store:** wip enhanced tx support
- **embedded/store:** conservative tx invalidation
- **embedded/store:** included filters in key readers
- **embedded/store:** including metadata records
- **embedded/store:** logical key deletion api
- **embedded/store:** keyReader in tx scope
- **embedded/store:** functional constraints
- **embedded/tbtree:** read as before returns history count
- **embedded/tbtree:** implements ExistKeyWithPrefix in snapshots
- **sql:** Add support for IS NULL / IS NOT NULL expressions
- **sql/index-on-nulls:** Update on-disk format to support nullable values
- **sql/timestamp:** Add CAST from varchar and integer to timestamp
- **sql/timestamp:** Add timestamp support to embedded/sql
- **sql/timestamp:** Add timestamp to protobuf definition
- **sql/timestamp:** Add timestamp to stdlib


<a name="v1.1.0"></a>
## [v1.1.0] - 2021-09-22
### Bug Fixes
- Minor updates to build/RELEASING.md
- Update Dockerfile.alma maintainer field
- **CI:** Fix building and releasing almalinux images
- **Dockerfile:** Fix compiling version information in docker images
- **Dockerfile.rndpass:** Fix building rndpass docker image
- **embedded/sql:** limit auto-increment to single-column pks
- **embedded/sql:** consider boolean type in maxKeyVal
- **embedded/sql:** exclude length from maxKey
- **embedded/sql:** return error when joint table doest not exist
- **embedded/sql:** support edge case of table with just an auto-increment column
- **embedded/sql:** take into account table aliasing during range calculations
- **embedded/sql:** suffix endKey when scan over all entries
- **embedded/sql:** in-mem catalog rollback and syncing fixes
- **embedded/sql:** adjust selector ranges calculation
- **embedded/sql:** improve error handling and parameters validation
- **embedded/sql:** set type any to nil parameters
- **embedded/sql:** fix table aliasing with implicit selectors
- **embedded/sql:** enforce ordering by grouping column
- **embedded/store:** fix constraint condition
- **embedded/store:** error handling when setting offset fails
- **pkg:** improve signature verification during audit
- **pkg/stdlib:** fix driver connection releasing

### Changes
- remove wip warning for fully implemented features
- Remove experimental S3 warning from README
- Update codenotary maintainer info
- Update RELEASING.md with documentation step.
- Add documentation link to command line help outputs
- Add documentation link at the beginning of README.md
- Add documentation badge to README.md
- **CI:** Explicitly require bash in gh action building docker images
- **CI:** Use buildkit when building docker images
- **CI:** Build almalinux-based immudb image
- **Dockerfile:** Update base docker images
- **Dockerfile:** Use scratch as a base for immudb image
- **Dockerfile:** Build a debian-based image for immudb next to the scratch one
- **Dockerfile:** Remove unused IMMUDB_DBNAME env var
- **Makefile:** Add darwin/amd64 target
- **Makefile:** More explicit webconsole version
- **build.md:** Add info about removing webconsole/dist folder
- **cmd/immuadmin:** parse all db flags when preparing settings
- **cmd/immuadmin:** remove replication flag shortcut
- **cmd/immuadmin:** improve flag description and rollback args spec
- **cmd/immuclient:** display number of updated rows as result of sql exec
- **cmd/immudb:** use common replication prefix
- **docker:** Update generation of docker tags
- **embedded:** leverage kv constraint to enforce upsert over auto-incremental pk requires row to already exist
- **embedded/multierr:** enhace multi error implementation
- **embedded/sql:** remove join constraints
- **embedded/sql:** minor refactoring to simplify code
- **embedded/sql:** convert unmapIndexedRow into unmapRow with optional indexed value
- **embedded/sql:** index prefix function
- **embedded/sql:** catalog loading requires up to date data store indexing
- **embedded/sql:** mark catalog as mutated when using auto incremental pk
- **embedded/sql:** move index spec closer to ds
- **embedded/sql:** changed identifiers length in catalog
- **embedded/sql:** move index selection closer to data source in query statements
- **embedded/sql:** minor code refactoring
- **embedded/sql:** include constraint only when insert occurs without auto_incremental pk
- **embedded/sql:** disable TIMESTAMP data-type
- **embedded/sql:** get rid of limited joint implementation
- **embedded/sql:** minor code simplification
- **embedded/sql:** ignore null values when encoding row
- **embedded/sql:** fix primary key supported types error message
- **embedded/sql:** optimize integer key mapping
- **embedded/sql:** use plain big-endian encoding for integer values
- **embedded/sql:** include support for int64 parameters
- **embedded/sql:** use Cols as a replacement for ColsByID
- **embedded/sql:** leverage endKey to optimize indexing scanning
- **embedded/sql:** use int64 as value holder for INTEGER type
- **embedded/sql:** add further validations when encoding values as keys
- **embedded/sql:** fix max key length validation
- **embedded/sql:** reserve byte to support multi-ordered indexes
- **embedded/sql:** expose primary key index id
- **embedded/sql:** wip scan optimizations based on query condition and sorting
- **embedded/sql:** partial progress on selector range calculation
- **embedded/sql:** validate non-null pk when decoding index entry
- **embedded/sql:** limit upsert to tables without secondary indexes
- **embedded/sql:** optional parenthesis when specifying single-column index
- **embedded/sql:** partial progress on selector range calculation
- **embedded/tbtree:** typo in log message
- **embedded/tbtree:** return kv copies
- **embedded/tbtree:** adjust seekKey based on prefix even when a value is set
- **embedded/tbtree:** compaction doesn't need snapshots to be closed
- **embedded/tools:** update sql stress tool with exec summary
- **pkg/api:** changed db identifiers type
- **pkg/api:** use follower naming for replication credentials
- **pkg/api:** use a map for holding latest auto-incremental pks
- **pkg/api:** delete deprecated clean operation
- **pkg/api:** include updated rows and last inserted pks in sql exec result
- **pkg/api:** use fresh id in proto message
- **pkg/api:** use int64 as value holder for INTEGER type
- **pkg/client:** move unit testing to integration package to avoid circular references
- **pkg/client:** changed db identifiers type
- **pkg/database:** warn about data migration needed
- **pkg/database:** update integration to exec summary
- **pkg/database:** minor adjustments based on multi-column indexing
- **pkg/database:** warn about data migration needed
- **pkg/database:** include updated rows and last inserted pks in sql exec result
- **pkg/database:** display as unique if there is a single-column index
- **pkg/database:** create sql db instance if not present
- **pkg/database:** minor refactoring coding conventions
- **pkg/database:** remove active replication options from database
- **pkg/database:** minor renaming after rebase
- **pkg/pgsql/server:** adds pgsql server maxMsgSize 32MB limit
- **pkg/pgsql/server:** add a guard on payload message len
- **pkg/replication:** use new context for each client connection
- **pkg/replication:** handle disconnection only within a single thread
- **pkg/replication:** use info log level for network failures
- **pkg/server:** changed default db file size and make it customizable at db creation time
- **pkg/server:** change max concurrency per database to 30
- **pkg/server:** nil tlsConfig on default options
- **pkg/server:** validate replication settings
- **pkg/server:** use replica wording
- **pkg/server:** followers management
- **pkg/stdLib:** implementing golang standard sql interfaces
- **pkg/stdlib:** increase code coverage and fix blob results scan
- **pkg/stdlib:** remove pinger interface implementation and increase code coverage
- **pkg/stdlib:** immuclient options identifier(uri) is used to retrieve cached connections
- **pkg/stdlib:** simplified and hardened uri handling

### Features
- Dockerfile for almalinux based image
- **cmd/immuadmin:** add replication flags
- **cmd/immuadmin:** add flag to exclude commit time
- **embedded/multierr:** implement stardard error Is & As methods
- **embedded/sql:** use explicitelly specified index as preffered one
- **embedded/sql:** wip unique multi-column indexes
- **embedded/sql:** wip auto-incremental integer primary keys
- **embedded/sql:** value expressions in row specs
- **embedded/sql:** switch to signed INTEGER
- **embedded/sql:** exec summary containing number of updated/inserted rows and last inserted pk per table
- **embedded/sql:** max length on variable sized types as requirement for indexing
- **embedded/sql:** multi-column primary keys
- **embedded/sql:** support index spec in joins
- **embedded/sql:** expose scanSpecs when resolving a query
- **embedded/sql:** wip unique multi-column indexing
- **embedded/sql:** towards more powerful joins
- **embedded/sql:** inner join with joint table and subqueries
- **embedded/store:** parameterized commit time
- **embedded/store:** leverage endKey from tbtree key reader
- **embedded/tbtree:** include endKey to instruct scan termination
- **pkg/database:** row verification with composite primary keys
- **pkg/follower:** follower replication
- **pkg/pgsql/server:** add support for flush message
- **pkg/replication:** initial active replication capabilities
- **pkg/server:** initial support for active replication of user created databases
- **pkg/server:** upgrade db settings to include or exclude commit time
- **pkg/server:** systemdb and defaultdb follower replication


<a name="v1.0.5"></a>
## [v1.0.5] - 2021-08-02
### Bug Fixes
- bind psql port to the same IP address as grpc interface ([#867](https://github.com/vchain-us/immudb/issues/867))
- Update crypto, sys dependencies
- consistent reads of recently written data
- **embedded/ahtree:** fix the full revert corner case
- **embedded/store:** Truncate aht before commit
- **embedded/store:** revert change so to prevent nil assigments
- **embedded/store:** handle missing error case during commit phase
- **embedded/store:** Don't fail to open on corrupted commit log
- **embedded/store:** use reserved concurrency slot for indexing
- **embedded/tbtree:** use padding to ensure stored snapshots are named following lex order
- **embedded/tbtree:** garbage-less nodes log
- **embedded/tbtree:** ensure clog is the last one being synced
- **embedded/tbtree:** flush logs containing compacted index
- **embedded/tbtree:** ensure proper data flushing and syncing
- **pkg/client/auditor:** fix and enhance state signature verification
- **pkg/pgsql/server:** fix boolean and blob extended query handling
- **pkg/pgsql/server:** hardened bind message parsing and fix leftovers
- **pkg/pgsql/server/fmessages:** use a variable size reader to parse fe messages
- **pkg/server:** initialize db settings if not present
- **pkg/server:** lock userdata map read
- **s3:** Use remote storage for index
- **s3:** Use remote storage for new databases
- **sql/engine:** Harden DecodeValue
- **store/indexer:** Ensure indexer state lock is always unlocked

### Changes
- move sqlutils package to schema
- Update dependencies
- increased coverage handling failure branches ([#861](https://github.com/vchain-us/immudb/issues/861))
- group user methods in a dedicated file
- Better logging when opening databases
- remove unused interceptors and add missing error code prefixes
- **appendable:** Expose validation functions of appendable options
- **appendable/multiapp:** Add hooks to the MultiFileAppender implementation
- **appendable/multiapp:** Introduce appendableLRUCache
- **cmd/immuclient:** fix panic in immuclient cli mode
- **cmd/immuclient:** update error comparisson
- **embedded:** col descriptor with attributes
- **embedded/ahtree:** minor refactoring improving readability
- **embedded/ahtree:** auto-truncate partially written commit log
- **embedded/ahtree:** minor changes towards code redabilitiy
- **embedded/cache:** Add Pop and Replace methods to LRUCache.
- **embedded/sql:** explicit catalog reloading upon failed operations
- **embedded/sql:** type specialization
- **embedded/sql:** Remove linter warnings
- **embedded/sql:** initial type specialization in place
- **embedded/sql:** remove public InferParameters operations from sql statements
- **embedded/sql:** validate either named or unnamed parameters are used within the same stmt
- **embedded/sql:** towards non-blocking sql initialization
- **embedded/sql:** parameters type inference working with aggregations
- **embedded/sql:** several adjustments and completion in type inference functions
- **embedded/sql:** dump catalog with a different database name
- **embedded/sql:** cancellable wait for catalog
- **embedded/sql:** expose InferParameters function in RowReader interface
- **embedded/store:** tx metatada serialization/deserialization
- **embedded/store:** edge-case validation with first tx
- **embedded/store:** validate replicated tx against current store
- **embedded/store:** minor refactoring improving readability
- **embedded/store:** auto-truncate partially written commit log
- **embedded/store:** minor code simplification
- **embedded/tbtree:** expose current number of snapshots
- **embedded/tbtree:** warn if an error is raised while discarding snapshots
- **embedded/tbtree:** nodes and commit prefix renaming
- **embedded/tbtree:** use setOffset for historical data overwriting
- **embedded/tbtree:** auto-truncate partially written commit log
- **embedded/tbtree:** full snapshot recovery
- **embedded/tbtree:** enable snapshot generation while compaction is in progress
- **pkg/api:** kept same attribute id in TxEntry message
- **pkg/api:** fix typo inside a comment
- **pkg/api:** remove information not required to cryptographically prove entries
- **pkg/api:** kept simple db creation api to guarantee backward compatibility
- **pkg/api:** comment on deprecated and not yet supported operations
- **pkg/auth:** list of supported operations in maintenance mode
- **pkg/database:** replace fixed naming with current database
- **pkg/database:** migrate systemdb catalog to fixed database naming
- **pkg/database:** single-store databases
- **pkg/database:** support the case where database tx is not the initial one due to migration
- **pkg/database:** no wait for indexing during tx replication
- **pkg/database:** sql operations after catalog is created
- **pkg/database:** method to retrieve row cursor based on a sql query stament
- **pkg/database:** re-construct sql engine once catalog is ready
- **pkg/database:** use fixed database name
- **pkg/database:** sql catalog per database. migration from shared catalog store when required
- **pkg/database:** sql catalog reloading on replica
- **pkg/database:** wait for sql engine initialization before closing
- **pkg/database:** log warning about WIP feature when using replication capabilities
- **pkg/database:** replace fixing naming with current database
- **pkg/database:** expose catalog loading operation
- **pkg/database:** catalog reloading by replicas
- **pkg/database:** gracefully stop by cancelling sql initialization
- **pkg/database:** internal method renaming
- **pkg/database:** log when a database is sucessfully opened
- **pkg/database:** add IsReplica method
- **pkg/database:** parameter inference for parsed statements
- **pkg/errors:** fix user operations error codes with pgsql official ones, increase coverage
- **pkg/errors:** immuerrors use an internal map to determine code from the message
- **pkg/errors:** add more error codes and add Cod prefix to codes var names
- **pkg/errors:** add comments
- **pkg/pgsql:** increase server coverage
- **pkg/pgsql/server:** handle empty statements
- **pkg/pgsql/server:** increase multi inserts tests number in simple and extended query
- **pkg/pgsql/server:** hardened INTEGER parameters conversion
- **pkg/pgsql/server:** some polishments in the state machine
- **pkg/pgsql/server:** simplify query machine
- **pkg/pgsql/server:** increase code coverage
- **pkg/pgsql/server:** protect  parameters description message from int16 overflown
- **pkg/pgsql/server:** add max parameters value size guard and move error package in a higher level to avoid cycle deps
- **pkg/pgsql/server:** add bind message negative value size guards
- **pkg/pgsql/server:** handle positional parameters
- **pkg/pgsql/server:** add a guard to check max message size and handle default case in parsing format codes in bind messages
- **pkg/pgsql/server/fmessages:** uniform malformed bind messages
- **pkg/server:** disable user mgmt operations in maintenance mode
- **pkg/server:** systemdb renaming
- **pkg/server:** move userdata lock in the inner method getLoggedInUserDataFromUsername
- **pkg/server:** remove methods moved to user file
- **pkg/server:** fix typo in error message
- **pkg/server:** remove duplicated property
- **pkg/server:** minor adjustments after rebasing from master branch
- **pkg/server:** minor updates after rebasing
- **pkg/stream:** use io.Reader interface
- **pkg/stream:** use wrapped errors
- **pkg/stream:** inject immu errors
- **pkg/stream:** fix namings on stream api objects

### Features
- immuclient running as auditor - replace "prometheus-port" and "prometheus-host" CLI flags with "audit-monitoring-host" and "audit-monitoring-port" (int) and start a single HTTP server which exposes all the needed endpoints (GET /metrics, /initz, /readyz, /livez and /version)
- add /healthz and /version endpoints for immudb and immuclient auditor
- add immudb error package
- **appendable:** Add remote s3 backend
- **cmd/immuadmin:** update database command
- **cmd/immuadmin:** upgrade database creation with settings
- **cmd/immuadmin:** add flag to create database as a replica
- **cmd/immuclient:** upgrade database creation with settings
- **embedded/sql:** adding method to infer typed parameters from sql statements
- **embedded/sql:** catalog dumping
- **embedded/sql:** towards leveraging readers for type inference
- **embedded/sql:** support for named positional parameters
- **embedded/sql:** support for unnamed parameters
- **embedded/store:** passive waiting for transaction commit
- **embedded/store:** WIP replicatedCommit
- **embedded/store:** tx export and commit replicated
- **pkg/api:** endpoints for exporting and replicating txs
- **pkg/api:** enhanced createDatabase endpoint to specify database replication
- **pkg/api:** new endpoint to update database settings
- **pkg/client:** replica creation and replication API
- **pkg/client:** implements update database settings operation
- **pkg/client:** deprecate CleanIndex operation
- **pkg/database:** db as replica and replication operations
- **pkg/database:** parameters type inference exposed in database package
- **pkg/database:** implement passive wait for committed tx
- **pkg/database:** suppport runtime replication settings changes
- **pkg/error:** add improved error handling
- **pkg/pgsql/server:** add extended query messages and inner logic
- **pkg/server:** enable simultaneous replication of systemdb and defaultdb
- **pkg/server:** stream of committed txs
- **pkg/server:** initial handling of database replication settings
- **pkg/server:** replicas and replication endpoints
- **pkg/server:** implements update database settings endpoint
- **pkg/server:** leverage maintenance mode to recover systemdb and defaultdb databases
- **pkg/stream:** readFully method to read complete payload transmitted into chunks


<a name="v1.0.1"></a>
## [v1.0.1] - 2021-06-07
### Bug Fixes
- go mod tidy/vendor with statik module ([#796](https://github.com/vchain-us/immudb/issues/796))
- **cmd/immuclient:** remove warnings on sql commands in interactive mode
- **cmd/immuclient:** improve immuclient tx and safetx error message
- **embedded/sql:** interprete binary prefix if followed by a quote
- **pkg/server:** always create system db (even when auth is off)

### Changes
- enhance Makefile so to automatically download latest webconsole if not already present
- README/doc updates ([#791](https://github.com/vchain-us/immudb/issues/791))
- enable webconsole in docker image
- remove mtls evironments var from dockerfile
- **embedded/store:** apply synced settings to indexing data
- **embedded/store:** sync values once all entries are written
- **pkg/database:** retry database selection after registration
- **pkg/database:** auto-registration when not present in the catalog

### Features
- **embedded/sql:** support <column> <type> NULL syntax
- **pkg/database:** enhace table description by adding nullable constraint
- **webconsole:** default web console page ([#786](https://github.com/vchain-us/immudb/issues/786))


<a name="v1.0.0"></a>
## [v1.0.0] - 2021-05-21
### Bug Fixes
- tlsConfig is always non-nil
- make prequisites fixes introduced in [#726](https://github.com/vchain-us/immudb/issues/726) ([#732](https://github.com/vchain-us/immudb/issues/732))
- fix windows installer service and missing flags
- **cmd/immuclient/immuclienttest:** fix options injection in client test helper ([#749](https://github.com/vchain-us/immudb/issues/749))
- **embedded:** ensure readers get properly closed
- **embedded/sql:** close reader after loading catalog
- **embedded/sql:** add missing error handling
- **embedded/sql:** fix selector aliasing
- **embedded/sql:** prevent side effects in conditional clauses
- **embedded/store:** fix issue when resuming indexing
- **embedded/store:** notified latest committed tx when opening store
- **embedded/store:** fix indexing data race
- **pkg/client:** row verification with nullable values
- **pkg/client/cache:** fix lock file cache issue on windows
- **pkg/client/cache:** clean state file when re-writing old stetes
- **pkg/database:** unwrap parameter before calling sqlexec
- **pkg/database:** use SQLPrefix when reopening database
- **pkg/pgsql/server:** handle data_row message with text format
- **pkg/server:** complete error handling
- **pkg/server:** disable pgsql server by default and fix previous server tests
- **pkg/sql:** columns resolved with aliases
- **pkg/sql:** resolve shift/reduce conflict in SELECT stmt

### Changes
- blank line needed after tag or interpreted as comment
- bundle webconsole inside dist binaries
- fix rebase leftovers
- fix acronym uppercase
- reword wire compatibility
- increase coverage and minor fix
- make roadmap about pgsql wire more explicit ([#723](https://github.com/vchain-us/immudb/issues/723))
- expose missing methods to REST ([#725](https://github.com/vchain-us/immudb/issues/725))
- inject immudb user authentication
- fix makefile leftovers
- improved make dist script
- move concrete class dblist to database package
- revert 3114f927adf4a9b62c4754d42da88173907a3a9f in order to allow insecure connection on grpc server
- dblist interface is moved to database package and extended
- add pgsql related flags
- **cmd/immuclient:** query result rendering
- **cmd/immuclient:** add describe, list, exec and query commands to immuclient shell
- **cmd/immuclient:** render raw values
- **cmd/immudb:** add debug info env var details
- **cmd/immudb/command:** enabled pgsql server only in command package
- **cmd/immudb/command:** restore missing pgsql cmd flag
- **cmd/immudb/command:** remove parsing path option in unix
- **cmd/immudb/command:** handle tls configuration errors
- **embedded/cache:** thread-safe lru-cache
- **embedded/sql:** expose functionality needed for row verification
- **embedded/sql:** minor refactoring to expose functionality needed for row verification
- **embedded/sql:** case insensitive identifiers
- **embedded/sql:** case insensitive functions
- **embedded/sql:** resolve query with current snapshot if readers are still open
- **embedded/sql:** set 'x' as blob prefix
- **embedded/sql:** move sql engine under embedded package
- **embedded/sql:** store non-null values and only col ids on encoded rows
- **embedded/sql:** safer support for selected database
- **embedded/sql:** validate table is empty before index creation
- **embedded/sql:** skip tabs
- **embedded/sql:** keep one snapshot open and close when releasing
- **embedded/sql:** improved nullables
- **embedded/store:** pausable indexer
- **embedded/store:** commitWith callback using KeyIndex interface
- **embedded/store:** index info to return latest indexed tx
- **embedded/store:** use indexer state to terminate indexing goroutine
- **embedded/store:** log during compaction
- **embedded/tbree:** postpone reader initialization until first read
- **embedded/tbtree:** remove dots denoting progress when flushing is not needed
- **embedded/tbtree:** index compaction if there is not opened snapshot, open snapshot if compaction is not in already in progress
- **embedded/tbtree:** make snapshot thread-safe
- **embedded/watchers:** cancellable wait
- **pkg/api:** render varchar as raw string value
- **pkg/api:** include data needed for row verification
- **pkg/api:** render varchar as quoted string
- **pkg/api:** render varchar as quoted strings
- **pkg/api:** sql api spec
- **pkg/api/schema:** Handle tools via modules ([#726](https://github.com/vchain-us/immudb/issues/726))
- **pkg/auth:** add SQL-related permissions
- **pkg/auth:** perm spec for row verification endpoint
- **pkg/client:** remove deprecated operations
- **pkg/client:** use  to fetch current database name
- **pkg/client:** auto convert numeric values to uint64
- **pkg/client:** improved sql API
- **pkg/client/cache:** release lock only if locked file is present, and wait for unlock when already present
- **pkg/database:** set implicit database using `UseDatabase` method
- **pkg/database:** typed-value proto conversion
- **pkg/database:** towards prepared sql query support
- **pkg/database:** row verification against kv-entries
- **pkg/database:** improved parameters support
- **pkg/database:** return mapped row values
- **pkg/database:** upgrade ExecAll to use KeyIndex interface
- **pkg/database:** add missing copy
- **pkg/database:** support index compaction with sql engine in place
- **pkg/database:** support multi-selected columns
- **pkg/database:** use store-level snapshots
- **pkg/database:** upgrade wait for indexing api
- **pkg/database:** ensure rowReader get closed upon completion
- **pkg/database:** use MaxKeyScanLimit to limit query results
- **pkg/database:** support for nullable values
- **pkg/database:** close sql engine when db gets closed
- **pkg/database:** make use of UseDatabase operation
- **pkg/embedded:** introduce Snapshot at Store level
- **pkg/pgsql:** handle empty response and command complete message
- **pkg/pgsql:** add pgsql server wire protocol stub
- **pkg/pgsql:** handle parameter status and terminate messages
- **pkg/pgsql:** fix filename format
- **pkg/pgsql:** bind immudb sql engine
- **pkg/pgsql:** use options flag to determine if pgsql server need to be launched
- **pkg/pgsql/client:** add jackc/pgx pgsql client for testing purpose
- **pkg/pgsql/server:** limit number of total connections and do not stop server in case of errors
- **pkg/pgsql/server:** handle an ssl connection request if no certificate is present on server
- **pkg/pgsql/server:** protect simple query flow with a mutex
- **pkg/pgsql/server:** enforce reserved statements checks
- **pkg/pgsql/server:** handle version statement
- **pkg/pgsql/server:** default error in simplequery loop has error severity
- **pkg/pgsql/server:** add missing copyright
- **pkg/pgsql/server:** remove host parameter
- **pkg/pgsql/server:** move sysdb in session constructor
- **pkg/pgsql/server:** add debug logging messages, split session handling in multiple files
- **pkg/pgsql/server:** improve error handling when client message is not recognized
- **pkg/pgsql/server:** fix connection upgrade pgsql protocol messages
- **pkg/pgsql/server:** minor fixes and leftovers
- **pkg/server:** use systemdb as a shared catalog store
- **pkg/server:** fix db mock
- **pkg/server:** remove unused options
- **pkg/server:** remove tls configuration in server
- **pkg/server:** inject sqlserver in main immudb server
- **pkg/server:** renamed server reference
- **pkg/server:** load systemdb before any other db
- **pkg/sql:** alias overriding datasource name
- **pkg/sql:** add comment about nested joins
- **pkg/sql:** refactored AggregatedValue and TypedValue interfaces
- **pkg/sql:** unify augmented and grouped row readers
- **pkg/sql:** support for SUM aggregations
- **pkg/sql:** row reader to support GROUP BY behaviour
- **pkg/sql:** grammar adjustments to support aggregated columns
- **pkg/sql:** swap LIMIT and ORDER BY parse ordering
- **pkg/sql:** many internal adjustments related to name binding
- **pkg/sql:** ensure use snapshot is on the range of committed txs
- **pkg/sql:** joint column with explicit table reference
- **pkg/sql:** upgrade to new store commit api
- **pkg/sql:** support multiple spacing between statements
- **pkg/sql:** column descriptors in row readers
- **pkg/sql:** improve error handling
- **pkg/sql:** return ErrNoMoreRows when reading
- **pkg/sql:** towards catalog encapsulation
- **pkg/sql:** improved null value support
- **pkg/sql:** order-preserving result set
- **pkg/sql:** using new KeyReaderSpec
- **pkg/sql:** joins limited to INNER type and equality comparison against ref table PK
- **pkg/sql:** row reader used to close the snapshot once reading is completed
- **pkg/sql:** mapping using ids, ordering and renaming support
- **pkg/sql:** composite readers
- **pkg/sql:** wip multiple readers
- **pkg/sql:** catch store indexing errors
- **pkg/sql:** add generated sql parser
- **pkg/sql:** make row values externally accessible
- **pkg/sql:** remove offset param
- **pkg/sql:** value-less indexed entries
- **pkg/sql:** encoded value with pk entry
- **pkg/sql:** remove alter column stmt
- **pkg/sql:** inmem catalog with table support
- **pkg/sql:** inmem catalog
- **pkg/sql:** catalog construct
- **pkg/sql:** primary key supported type validation
- **pkg/sql:** use standardized transaction closures
- **pkg/sql:** col selector with resolved datasource
- **pkg/sql:** case insensitive reserved words
- **pkg/sql:** use token IDENTIFIER
- **tools/stream:** upgrade stream tools dependencies

### Code Refactoring
- **pkg/server:** tls configuration is moved in command package from server

### Features
- display version at startup ([#775](https://github.com/vchain-us/immudb/issues/775))
- enhance database size and number of entries metrics to support multiple databases add CORS middleware to metrics HTTP endpoints ([#756](https://github.com/vchain-us/immudb/issues/756))
- CREATE TABLE IF NOT EXISTS ([#738](https://github.com/vchain-us/immudb/issues/738))
- **cmd/immuclient:** list and describe tables
- **cmd/immuclient:** use 'tables' to display the list of tables within selected database
- **embedded/sql:** use snapshot as state method
- **embedded/sql:** arithmetic expressions within where clause
- **embedded/sql:** 'NOT NULL' constraint
- **embedded/sql:** special case when all selectors are aggregations and there is no matching rows
- **embedded/sql:** enhanced sql parser to support multi-lined statements
- **embedded/sql:** INSERT INTO statement
- **embedded/sql:** LIKE operator support
- **embedded/store:** uniqueness constraint and de-coupled indexer
- **embedded/store:**  operation
- **embedded/tools:** initial SQL stress tool ([#760](https://github.com/vchain-us/immudb/issues/760))
- **pkg/api:** sql endpoints for row verification
- **pkg/api:** noWait mode for sql statements
- **pkg/client:** row verification
- **pkg/client:** towards client-side sql support
- **pkg/database:** towards sql support
- **pkg/database:** towards integrated sql engine. handling database creation at server level
- **pkg/database:** list and describe tables
- **pkg/database:** row verification endpoint
- **pkg/pgsql:** add tls support
- **pkg/pgsql/server:** dblist is injected in pgsql server
- **pkg/pgsql/server:** setup pgsqk error handling
- **pkg/pgsql/server:** handle nil values
- **pkg/server:** expose  endpoint
- **pkg/server:** row verification endpoint
- **pkg/server:** initial integration of sql engine
- **pkg/sql:** column selector alias support
- **pkg/sql:** jointRowReader towards supporting joins
- **pkg/sql:** towards supporting COUNT, SUM, MIN, MAX and AVG
- **pkg/sql:** towards aggregated values support
- **pkg/sql:** towards supporting filtered aggregations
- **pkg/sql:** towards group by and aggregations support
- **pkg/sql:** noWait for indexing mode
- **pkg/sql:** improved nullable support
- **pkg/sql:** towards GROUP BY support
- **pkg/sql:** implements NOW() function
- **pkg/sql:** list and describe tables
- **pkg/sql:** LIMIT clause to determine max number of returned rows
- **pkg/sql:** queries over older data
- **pkg/sql:** parameters support
- **pkg/sql:** initial parameters support
- **pkg/sql:** towards parameter support
- **pkg/sql:** support for SELECT * FROM queries
- **pkg/sql:** row projection
- **pkg/sql:** towards projected rows
- **pkg/sql:** auto-commit multi-statement script
- **pkg/sql:** subquery aliases
- **pkg/sql:** support for WHERE clause
- **pkg/sql:** towards row filtering with conditional readers
- **pkg/sql:** support for boolean values
- **pkg/sql:** index reloading
- **pkg/sql:** catalog reloading
- **pkg/sql:** ASC/DESC row sorting by any indexed column
- **pkg/sql:** implements CREATE INDEX stmt
- **pkg/sql:** support of foreign keys of any pk type
- **pkg/sql:** multiple joins support
- **pkg/sql:** col selector binding
- **pkg/sql:** aggregations without row grouping
- **pkg/sql:** seekable, ordered and filtered table scan
- **pkg/sql:** ordering in descending mode
- **pkg/sql:** towards ordering row scans
- **pkg/sql:** towards query resolution with multiple datasources
- **pkg/sql:** towards query processing
- **pkg/sql:** upsert processing
- **pkg/sql:** towards insert into statement processing
- **pkg/sql:** table creation with primary key
- **pkg/sql:** primary key spec
- **pkg/sql:** initial work on sql engine
- **pkg/sql:** multi-line scripts
- **pkg/sql:** snapshot support
- **pkg/sql:** support for comments
- **pkg/sql:** support for EXISTS in subquery
- **pkg/sql:** support for INNER, LEFT and RIGHT joins
- **pkg/sql:** support for parameters
- **pkg/sql:** support for system values e.g. TIME
- **pkg/sql:** aggregated functions
- **pkg/sql:** use colSelector instead of identifiers
- **pkg/sql:** expressions parsing
- **pkg/sql:** multi-db queries
- **pkg/sql:** multi-row insertion
- **pkg/sql:** initial support for SELECT statement
- **pkg/sql:** transactional support
- **pkg/sql:** support for insertions
- **pkg/sql:** support table modifications
- **pkg/sql:** support index creation
- **pkg/sql:** include column specs
- **pkg/sql:** partial grammar with multiple statements
- **pkg/sql:** initial commit for sql support


<a name="cnlc-2.2"></a>
## [cnlc-2.2] - 2021-04-28
### Bug Fixes
- update Discord invite link
- readme typo and mascot placement ([#693](https://github.com/vchain-us/immudb/issues/693))
- **embedded/store:** ensure done message is received
- **pkg/client:** delete token file on logout only if the file exists

### Changes
- github workflow to run stress_tool ([#714](https://github.com/vchain-us/immudb/issues/714))
- README SDK description and links ([#717](https://github.com/vchain-us/immudb/issues/717))
- fix immugw support
- Add roadmap
- Add benchmark to README (based on 0.9.x) ([#706](https://github.com/vchain-us/immudb/issues/706))
- remove grpc term from token expiration description
- **embedded/store:** use specified sync mode also for the incremental hash tree
- **embedded/store:** check latest indexed tx is not greater than latest committed one
- **embedded/store:** defer lock releasing
- **pkg/client/clienttest:** add VerifiedGetAt mock method
- **pkg/database:** use newly exposed KeyReaderSpec

### Features
- add token expiration time flag
- **embedded/store:** readAsBefore and reset reader
- **pkg/sql:** readAsBefore operation


<a name="v0.9.2"></a>
## [v0.9.2] - 2021-04-08
### Bug Fixes
- include AtTx in StreamZScan response
- password reader 'inappropriate ioctl for device' from stdin ([#658](https://github.com/vchain-us/immudb/issues/658))
- fix StreamVerifiedSet and Get and add an (integration) test for them
- fix inclusion proofs in StreamVerifiedSet and Get
- **embedded:** use mutex to sync ops at tx lru-cache
- **embedded:** fix data races
- **embedded/store:** ensure waitees get notified when store is restarted
- **embedded/store:** remove checking for closed store when fetching any vlog
- **embedded/store:** continue indexing once index is replaced with compacted index
- **embedded/store:** set delay with duration in ms
- **embedded/store:** fix indexing sync and error retrieval
- **embedded/store:** ensure watchers get notified when indexing is up-to-date
- **embedded/store:** sync ReadTx operation
- **embedded/tbtree:** set lastSnapshot once flushed is completed
- **embedded/tbtree:** insertion delay while compacting not affecting compaction
- **embedded/tbtree:** release lock when compaction thld was not reached
- **pkg/auth:** add missing stream write methods to permissions
- **pkg/client:** fix minor leftover
- **pkg/client:** fix security issue: if client local state became corrupted an error is returned
- **pkg/client:** ensure dual proof verification is made when there is a previously verified state
- **pkg/database:** wrap seekKey with prefix only when seekKey is non-empty
- **pkg/server:** use latest snapshot when listing users

### Changes
- refactor code quality issues
- improve serverside stream error handling
- remove fake proveSinceTxBs key send in streamVerifiableSet
- polish streams methods and add comments
- renaming stream methods
- updating copyright
- renaming stream methods, add stubs and stream service factory
- in server store creation max value entry is fixed to 32Mb
- set stream supports multiple key values
- mocked server uses the inner immudb grpc server and can be gracefully stopped
- fixed minimum chunk size at 4096 bytes
- add max tx values length guard and remove code duplication
- fix binary notation
- move stream service to a proper package
- add video streamer command
- increase stream coverage and add a guard if key is present on a stream but no value is found
- **embedded:** fix some typos with comments
- **embedded:** remove unused cbuffer package
- **embedded:** log indexing notifications
- **embedded:** descriptive logs on indexing and already closed errors
- **embedded:** add logs into relevant operations
- **embedded:** add logger
- **embedded:** compaction and snapshot handling
- **embedded/appendable:** thread-safe multi-appendable
- **embedded/appendable:** sync before copying appendable content
- **embedded/appendable:** multi-appendable fine-grained locking
- **embedded/store:** remove conditional locking before dumping index
- **embedded/store:** general improvements on snapshot management
- **embedded/store:** leverage fine-grained locking when reading tx data
- **embedded/store:** stop indexing while commiting with callback
- **embedded/store:** use buffered channel instead of a circular buffer
- **embedded/store:** remove duplicated logging
- **embedded/store:** set max file size to 2Gb ([#649](https://github.com/vchain-us/immudb/issues/649))
- **embedded/store:** lock-less readTx
- **embedded/store:** set a limit on indexing iteration
- **embedded/store:** log number of transactions yet to be indexed
- **embedded/tbtree:** optimize seek position
- **embedded/tbtree:** revert seek key setting
- **embedded/tbtree:** optimize seek position
- **embedded/tbtree:** terminate reader if prefix won't match any more
- **embedded/tbtree:** sync before dumping
- **embedded/tbtree:** sync key-history log during compaction
- **embedded/watchers:** broadcasting optimisation
- **embedded/watchers:** minor renaming
- **embedded/watchers:** accept non-continuous notification
- **pkg/client:** add stream service factory on client and increase stream coverage
- **pkg/client:** add GetKeyValuesFromFiles helper method
- **pkg/client:** add a guard to check for min chunk size
- **pkg/client:** remove local files tests
- **pkg/client:** maps server error on client package
- **pkg/client:** integration test is skipped if immudb server is not present
- **pkg/database:** return error while waiting for index to be up to date
- **pkg/database:** ensure scan runs over fully up-to-date snapshot
- **pkg/database:** return error while waiting for index to be up to date
- **pkg/database:** use in-mem current snapshot in execAll operation
- **pkg/database:** illegal state guard is added to verifiable get and getTx methods
- **pkg/database:** leverage lightweight waiting features of embedded store
- **pkg/server:** add a guard to check for min chunk size
- **pkg/server:** add server error mapper interceptor
- **pkg/server:** add small delay for indexing to be completed
- **pkg/server:** max recv msg size is set to 32M
- **pkg/server:** revert quit chan exposure
- **pkg/server:** exposes Quit chan
- **pkg/stream:** add more corner cases guards
- **pkg/stream:** add some comments to mesasge receiver
- **pkg/stream:** remove duplicated code
- **pkg/stream:** renamed stream test package
- **pkg/stream:** add a guard to detect ErrNotEnoughDataOnStream on client side
- **pkg/stream:** remove bufio.reader when not needed
- **pkg/stream:** remove bufio and add ventryreceiver unit test
- **pkg/stream:** add ErrNotEnoughDataOnStream error and ImmuServiceReceiver_StreamMock
- **pkg/stream/streamtest:** add dummy file generator
- **tools:** fix copyright
- **tools/stream:** get stream content directly from immudb
- **tools/stream/benchmark:** add stream benchmark command
- **tools/stream/benchmark/streamb:** add SinceTx value to getStream

### Code Refactoring
- stream kvreceiver expose Next method to iterate over key values
- stream receiver implements reader interface
- use of explicit messages for stream request
- **pkg/stream:** use ParseValue func in zreceiver and remove the redundant readSmallMsg func
- **pkg/stream:** refactor receiver to increase simplicity

### Features
- Remove unnecessary dependencies ([#665](https://github.com/vchain-us/immudb/issues/665))
- add support for user, password and database flags in immuclient ([#659](https://github.com/vchain-us/immudb/issues/659))
- increase default store max value length to 32MB
- add client->server stream handler
- refactors and implement server->client stream handler
- Add StreamVerifiedSet and StreamVerifiedGet
- add flusher to stream data to client
- chunk size is passed as argument in client and server
- add Stream Scan and client stream ServiceFactory
- **embedded/store:** integrate watchers to support indexing synchronicity
- **embedded/store:** configurable compaction threshold to set the min number of snapshots for a compaction to be done
- **embedded/store:** expose insertion delay while compacting
- **embedded/store:** tx header cache to speed up indexing
- **embedded/tbtree:** automatically set seekKey based on prefixKey when it's not set
- **embedded/tbtree:** configurable insertion delay while compaction is in progress
- **embedded/watchers:** lightweight watching center
- **embedded/watchers:** fetchable current state
- **pkg/client:** handle illegal state error
- **pkg/database:** non-blocking index compaction
- **pkg/database:** non-blocking, no history compaction
- **pkg/database:** default scan parameters using up-to-date snapshot
- **pkg/server:** add signature on stream verifiable methods and tests
- **pkg/stream:** add exec all stream


<a name="v0.9.1"></a>
## [v0.9.1] - 2021-02-08
### Bug Fixes
- **cmd/sservice:** fix services management and add permissions guard
- **cmd/sservice:** fix group creation linux cross command
- **embedded/history:** read history log file to set initial offset
- **embedded/store:** copy key inside TxEntry constructor
- **embedded/store:** mutex on txlog
- **embedded/store:** continued indexing
- **embedded/store:** fix indexing sync ([#621](https://github.com/vchain-us/immudb/issues/621))
- **embedded/tbtree:** determine entry by provided seekKey
- **embedded/tbtree:** fix key history ordering ([#619](https://github.com/vchain-us/immudb/issues/619))
- **embedded/tbtree:** prevNode nil comparisson
- **embedded/tbtree:** use minkey for desc scan
- **pkg/client:** fix verifiedGetAt
- **pkg/client/auditor:** hide auditor password in logs
- **pkg/client/cache:** return an error if no state is found
- **pkg/database:** check key does not exists in latest state
- **pkg/server:** set default settings within DefaultStoreOptions method

### Changes
- update acknowledgments
- **cmd/sservice:** minor fixes
- **embeddded/tbtree:** reduce mem allocs
- **embedded:** expose store opts
- **embedded:** refactor TxEntry
- **embedded/store:** adapt after History changes
- **embedded/store:** move TxReader code to its own file
- **embedded/store:** renamed reader as KeyReader
- **embedded/store:** use conditional locking in indexing thread
- **embedded/store:** minor KeyReader renaming
- **embedded/store:** validates targetTx is consistent with proof len
- **embedded/store:** sync access to commit and tx logs
- **embedded/store/options.go:** increase DefaultMaxKeyLen
- **embedded/tbtree:** offset map per branch
- **embedded/tbtree:** return ErrOffsetOutOfRange if invalid offset was provided
- **embedded/tbtree:** reduce mem consumption
- **embedded/tbtree:** history log file
- **embedded/tbtree:** configurable max key length
- **embedded/tbtree:** change history file extension
- **pkg:** unit testing index cleanup, use selected db
- **pkg:** current db included in signed state
- **pkg/api:** minor changes in TxScan message
- **pkg/api:** history limit as int32
- **pkg/api:** include server uuid and db name into state message
- **pkg/client:** validate returned entries from metadata
- **pkg/client:** bound reference if atTx is provided in VerifiedSetReferenceAt
- **pkg/client:** use indexing specified in GetRequest
- **pkg/client:** strip prefix from returned keys in txById and verifiedTxById
- **pkg/client:** add state service lock and unlock capabilities
- **pkg/client:** set bound on SetReference and ZAdd
- **pkg/database:** unsafe read tx inside CommitWith callback
- **pkg/database:** catch NoMoreEntries error and return empty list on scan and zscan operations
- **pkg/database:** return empty list if offset is out of range
- **pkg/database:** initial implementation of ExecAll with CommitWith
- **pkg/server:** naming conventions
- **pkg/server:** server mock wrapping default server implementation
- **pkg/server:** include uuid and db as result of verifiable operations
- **pkg/server:** initialize mts options
- **pkg/server:** expose store opts
- **pkg/server:** use server wrapper to enable post processing of results
- **pkg/server:** set default max value lenght to 1Mb
- **pkg/server:** change server default options. Max key value to 10kb

### Features
- **cmd/immuadmin:** db index cleanup
- **embedded:** history with offset and limit, key updates counting
- **embedded/appendable:** flush and seek to start before copying
- **embedded/appendable:** implements Copy function
- **embedded/appendable:** check no closed and flush before copying
- **embedded/store:** commitWith callback receiving assigned txID
- **embedded/store:** TxScan asc/desc order
- **embedded/store:** index cleanup
- **embedded/store:** allow increasing max value size after creation time
- **embedded/tbtree:** full dump of current snapshot
- **embedded/tbtree:** complete history implementation
- **embedded/tbtree:** HistoryReader to iterate over key updates
- **embedded/tbtree:** full dump using copy on history log
- **pkg:** index cleanup service
- **pkg/client:** add state file locker
- **pkg/client:** add verifiedGetSince
- **pkg/client:** implementation of TxScan operation
- **pkg/database:** history with offset and limit
- **pkg/database:** TxScan implementation
- **pkg/database:** support for free and bound references
- **pkg/database:** KeyRequest retrieves key at a specific tx or since a given tx
- **pkg/server:** sign new state within verifiable operations
- **pkg/server:** use exposed synced mode


<a name="v0.9.0"></a>
## [v0.9.0] - 2021-01-07
### Bug Fixes
- remove badger metrics and fix stats command
- **cmd/immuadmin/command:** fix immuadmin stats ([#592](https://github.com/vchain-us/immudb/issues/592))
- **pkg/database:** enable scan on fresh snapshot
- **pkg/server:** shutdown handlers and metrics server are moved in start method

### Changes
- removing audit-signature and add serverSigningPubKey
- remove print tree method
- restore inmemory_cache test
- **cmd/immuadmin:** temporary disable stats functionality
- **pkg/api:** upgrade rest endpoints
- **pkg/client:** implement missing methods in immuclient mock
- **pkg/server:** temporary remove proactive corruption checker ([#595](https://github.com/vchain-us/immudb/issues/595))

### Features
- add signature verification with a submitted public key


<a name="v0.9.0-RC2"></a>
## [v0.9.0-RC2] - 2020-12-29
### Bug Fixes
- **cmd/immuadmin/command:** fix unit tests
- **cmd/immuclient:** fix unit tests
- **embedded/tbtree:** sync GetTs to prevent data races
- **pkg/api:** change order of validations when checking state signature

### Changes
- adapt coverage to the new server implementation
- fix immuserver mock
- **cmd/immuadmin:** disable stats and removed print tree command
- **cmd/immuclient:** print verified label when executing safereference
- **pkg/client:** update service mock to new API
- **pkg/database:** add input validations during verifiable set
- **pkg/database:** implements History using lock-based operation

### Code Refactoring
- uniform server and client tests
- improving buffconn server with splitting start method in initialization and start

### Features
- **embedded/store:** implements lock-based History without requiring snapshot creation
- **pkg/client:** update auditor implementation to new server API
- **pkg/client:** implementation of client-side verifiedZAdd
- **pkg/client:** implements VerifiedSetReference
- **pkg/database:** implementation of verifiableZAdd
- **pkg/database:** implementation of VerifiableSetReference


<a name="v0.9.0-RC1"></a>
## [v0.9.0-RC1] - 2020-12-22
### Bug Fixes
- **cmd/immuclient:** print referenced key
- **cmd/immuclient:** print referenced key
- **embedded/store:** fix race condition
- **embedded/store:** fix race condition
- **embedded/store:** contemplate bad-formated proof
- **embedded/tbtree:** fix issue when initialKey is greater than keys
- **pkg/common:** fix leftover in index wrapper
- **pkg/database:** add cyclic references validation during resolution
- **pkg/database:** working scan and zscan without pagination
- **pkg/database:** adjust execAll method
- **pkg/database:** referenced key lookup when atTx is non-zero
- **pkg/database:** use EncodeReference in ExecAllIOps
- **pkg/database:** lookup for referenced key when atTx is non-zero
- **pkg/databse:** encoding of reference and zadd

### Changes
- proof proto definition
- datatype conversion methods
- remove badger and merkletree dependencies
- inject store reader inside zscan
- partial fix of scan test
- new proto definitions
- **api/schema:** removed consistency method
- **cmd:** adjusted commandline tools
- **cmd/immuclient:** add support for safe operations
- **cmd/immuclient:** add verified operations
- **database:** implements safeSet operation
- **database:** implements ByIndex operation
- **database:** implements safeByIndex operation
- **database:** contemplates the case not previously verified tx
- **database:** several fixes and unit testing adaptation
- **embedded:** rename as SnapshotSince
- **embedded/htree:** internal linear proof renaming
- **embedded/htree:** minor changes in proof struct
- **embedded/store:** add method to retrieve tx metadata
- **embedded/store:** minor proof renaming
- **embedded/store:** return txMetadata when tx on commit
- **embedded/store:** return ErrTxNotFound when attemping to read non-existent tx
- **embedded/store:** minor changes in proof struct
- **embedded/store:** allow empty values and don't attempt to store in vlog
- **embedded/store:** add tx constructor with entries
- **embedded/store:** adjustments on store reader
- **embedded/store:** change tx proof method signature
- **embedded/store:** wrap keyNotFound index error
- **embedded/store:** add snapshotAt and adjust based on it
- **pkg:** rename to CurrentState
- **pkg:** several minor changes
- **pkg:** rename to ReferenceRequest
- **pkg:** rename to sinceTx
- **pkg/api:** add vLen property to TxEntry
- **pkg/api:** new proof messages
- **pkg/api:** several improvements on grpc api
- **pkg/api:** remove digest data type
- **pkg/api:** rename to Entry and ZEntry and embedded Reference
- **pkg/api:** add copyright notice
- **pkg/api:** new server proto definition
- **pkg/auth:** adjust permissions based on new api
- **pkg/client:** adjusted client providers
- **pkg/client:** adjusted golang client
- **pkg/client:** add safe method alises for backwards familiarity
- **pkg/client:** minor renaming to improve readability
- **pkg/database:** zscan order with tx after key
- **pkg/database:** implements new DB api using embedded storage
- **pkg/database:** add sinceTx to reference and make it handle key prefixes
- **pkg/database:** remove ambiguity in references
- **pkg/database:** minor adjustments
- **pkg/database:** get from snapshot or directly from store
- **pkg/database:** return functionality not yet implemented for VerifiableSetReference
- **pkg/database:** wait for indexing on execAll
- **pkg/database:** delay locking until indexing is done
- **pkg/database:** mutex for reusable txs
- **pkg/database:** fix get/set with prefix wrapping/unwrapping
- **pkg/database:** fixed methods with prefix mgmt, including scan
- **pkg/ring:** remove ring pkg
- **pkg/server:** proof construction in safeget operation
- **pkg/server:** disable proactive corruption checker
- **pkg/server:** partial use of embedded storage
- **pkg/server:** return number of tx as db size
- **pkg/server:** getBatch operation
- **pkg/server:** adjusted state signer
- **pkg/server:** adjusted UUID handler
- **pkg/server:** prevent logging request details
- **pkg/server:** adapt implementation to new api
- **pkg/server:** adapt to new database implementation
- **pkg/server:** disable cc
- **pkg/server:** remove in-memory option
- **pkg/server:** implements history operation
- **pkg/server:** comment unimplemented GetReference method
- **pkg/store:** moved package
- **pkg/tbree:** reader with descOrder
- **server:** implements safeGet
- **server/api:** minor changes in Item element

### Code Refactoring
- **pkg/server:** add database interface and inject in server package
- **pkg/server:** move database to new package

### Features
- partial implementation of safeGet
- add store reader, scan and sorted sets
- **embedded:** add Get operation without the need of a snapshot
- **embedded:** inclusiveSeek point when reading
- **embedded/tbtree:** use seek and prefix
- **pkg/client:** implements latest server API
- **pkg/client:** add GetSince method
- **pkg/database:** uniform implementation for set, references, zadd, scan and zscan operations
- **pkg/database:** verify reference upon key resolution
- **pkg/database:** complete set and get reference methods
- **pkg/database:** add execAllOps
- **pkg/database:** support for seekable scanning
- **pkg/database:** consistent reference handling, prevent cyclic references
- **pkg/server:** expose store options


<a name="v0.8.1"></a>
## [v0.8.1] - 2020-12-08
### Bug Fixes
- file ext removal
- consider the case when key was not yet inserted
- add permissions for the new CountAll gRPC method
- fix batchOps tests and minors fix for zAdd sorted set key generation
- avoid duplicate index insertion in zAdd batch operation transaction
- restore current offset after reading compressed data
- encode metadata numeric fields with 64bits
- appendable extensions without dot
- set fileSize after reading values from metadata
- read metadata before reading
- pass compression settings into newly created single-file appendable
- compression with bigger values
- compression with multiple-files
- appID parsing from filename
- set new appendable id when changing current appendable
- typos
- fix batchOps permission and clean sv ones
- return EOF when data cannot be fully read from appendables
- **embedded:** set correct offset while reading node
- **embedded/store:** release tx before linear proof generation
- **embedded/store:** use verificatication methods for dual proof evaluation
- **embedded/tools:** catch ErrNoMoreEntries when iterating over txs ([#569](https://github.com/vchain-us/immudb/issues/569))
- **pkg:** handle expired token error
- **pkg/client:** handle rootservice error inside constructor
- **pkg/client:** token service is not mandatory for setup a client
- **pkg/store:** fix bug on lexicographical read of multiple sets
- **pkg/store:** fix reverse history pagination
- **pkg/store:** move separator at the beginning of a keyset
- **pkg/store:** fix scan and add tests for pagination
- **pkg/store:** scan item now contains immudb index, not badger timestamp
- **pkg/store:** fixes issue [#532](https://github.com/vchain-us/immudb/issues/532) ([#549](https://github.com/vchain-us/immudb/issues/549))
- **pkg/store:** fix key set generation. index reference flag (0,1 bit) is put at the end of the key to mantain lexicographical properties
- **pkg/store:** in reference based command key is optional if index is provided. Increase code coverage

### Changes
- copy on insert and fresh snapshots
- moved stress_tool under tools folder
- close appendable hash tree on close
- update readme with SDKs urls ([#506](https://github.com/vchain-us/immudb/issues/506))
- include github stars over time chart ([#509](https://github.com/vchain-us/immudb/issues/509))
- fix naming conventions
- export uwrap and wrap value method to be used in nimmu
- link to immuchallenge repo ([#528](https://github.com/vchain-us/immudb/issues/528))
- fix typo in cmd help ([#541](https://github.com/vchain-us/immudb/issues/541))
- renaming set key generation method
- remove *sv methods
- print commiting status after sync
- root method without error
- cbuffer pkg
- cbuffer pkg
- use constants for field len
- simplify linear proof
- remove verification during indexing
- unify kv hash calculation
- LinearProof verification
- minor code change
- multierror handling on close func
- multierror handling on close func
- error naming
- error naming
- unify options for appendables
- unify naming
- move options validations to options file
- minor changes in append operation
- hash tree construction
- advances in hash tree and proof construction
- make spec read fields public
- implement commit log for btree and improved reduced dump
- txi commit file ext
- aof as default file ext
- sync also applies to index
- panic if set key fails
- panic when close returns an error
- validate clog size after setting fileSize from metadata
- index file extensions
- return after auto mode
- mode as required flag
- renamed to IndexInfo
- random and monotonic keys
- split code into files, index options exposed for external config
- support for historical value tracing
- changes on historical value tracing
- validate hvalue when reading value from vlog
- newTx method using store values
- check for non-duplication within same tx
- when txInclusion is true then txLinking is made as well
- options validation
- close index when closing store
- add VERSION to metadata section
- time-based flushing
- do not store hvalue in index
- solved concurrency issues
- changes to solve and improve concurrency
- minor changes in stress tool
- unit testing
- change max value length to 64Mb
- set fileSize based on initial value set at creation time
- return partially written number of bytes when write is not completed
- reorg appendable packages
- when compression is enabled, each append takes place into a single file
- renamed filesize flag
- use channels for communicating data offsets
- close txlog after completion
- random key-value generation in stress tool
- kv alloc inside tx prep
- parallel IO non-exclusive yet synchorinized with store closing
- open with appendables to facilitate parallel IO
- spec parallel IO in stress tool
- appendale with path
- preparing for parallel IO
- optimized hashtree generation
- minor internal change on commit tx method
- time only considering commit time
- don't print dots if printAfter is 0
- don't print dots if printAfter is 0
- minor typo in error message
- preparing for concurrency optimization
- increase testing timeout
- close commit log file
- return EOF when available content is less than buffer size
- make key-value struct public
- initial commit with functional implementation
- move set and get batch in a separate file
- fix naming conventions and typos
- add setBatch, getBatch, setBatchOps method in mock client
- improved comprehensibility of the immudb configuration file
- add more batchOps tests, fixing naming typo
- renaming batchOps in ops and SetBatchOps in ExecAllOps
- fix immudb consistency diagram
- add more tests and remove code duplications
- **embedded/store:** return ErrNoMoreEntries when all tx has been read
- **embedded/store:** permit immediate snapshot renewals
- **embedded/store:** pre-allocated tx pool used for indexing and proofs
- **embedded/store:** use embedded reusable built-at-once hash tree
- **embedded/store:** internal changes to use innerhash for proof generation
- **embedded/store:** sleep binary linking thread until txs are committed ([#572](https://github.com/vchain-us/immudb/issues/572))
- **embedded/store:** changed defaults
- **embedded/store:** add data consistency validation during dual proof construction
- **embedded/store:** pre-allocated tx pool used with dual proof
- **embedded/store:** sleep indexing thread until there are entries to be indexed
- **pkg/api/schema:** increase code readability
- **pkg/auth:** fix get batch permissions
- **pkg/client:** fix comment
- **pkg/client:** client exposes structured values conversion tools
- **pkg/client/clienttest:** add inclusion, consistency, byIndex mock methods
- **pkg/client/clienttest:** add missing method inside service client mock
- **pkg/client/clienttest:** add mock service client constructor
- **pkg/client/timestamp:** fix naming convention
- **pkg/server:** add database name server validator
- **pkg/server:** remove ScanSV test
- **pkg/store:** add consistency check on zadd and safezadd index reference method and tests
- **pkg/store:** move reference code to a dedicated file
- **pkg/store:** fix comments in test
- **server:** enhance namings related audit report notification

### Code Refactoring
- batch ops produces monotonic ts sequences, index is mandatory if reference key is already persisted
- **pkg/client:** decoupled immuclient from rootService
- **pkg/store:** get resolve reference only by key
- **pkg/store:** add set separator. Fixes [#51](https://github.com/vchain-us/immudb/issues/51)

### Features
- set compression setting into value logs
- root returns number of entries
- payload and digest lru caches
- replay missing binary linking entries
- binary linking in-memory integration
- towards data compression capabilities
- dual proof construction and verification
- towards linear and dual proofs
- dual proof and liearProof against target accumulative linear hash
- dual cryptographic linking
- store immutable settings into metadata section
- towards dual cryprographic linking
- inclusion and consistency verification algorithms
- consistency proof
- multierr custom error wrapping multiple errors when closing the store
- several improvements, data by index
- towards persistent storage of mutable hash tree
- ongoing implementation of appendable hash tree
- add sync method
- TxReader starts from a txID
- IndexInfo to return up to which tx was indexed and error status of indexing task
- add interactive mode for get/set key values
- add method for reading value of a key within a tx
- retrieve the list of ts at which a key was updated
- key updates tracing to support historical reading
- back btree with multi-appendables
- add read numeric values
- initial indexing
- towards k-indexing
- getReference is exposed throught gRPC server and in SDK
- add lzw compression format
- data compression support by stress tool
- data compression
- towards data compression capabilities
- add flag to stress_tool to specify if values are random or fixed
- client supports paginated scan, zscan, history
- store accumulative linear hash into mutable hash tree
- add log file sizes and number of openned files per log type flags
- enable file mgmt settings
- multi-file appendables
- add metadata to appendables
- include full tx validation and fixed sync issue
- full committed tx verification against input kv data
- add key and value length params in stress tool
- parallel IO support
- optimized for concurrent committers
- added concurrent committers to stress tool
- add cryptographic linking verification of transactions
- add method to retrieve number of committed txs
- export sync mode config
- add batchOps reference operation
- expose CountAll through gRPC
- configurable max incomming msg size ([#526](https://github.com/vchain-us/immudb/issues/526))
- extend sorted set to support multiple equals key
- enhance auditor to publish tampering details at a specified URL (optional)
- add ZScan pagination
- rearrange badges on README ([#555](https://github.com/vchain-us/immudb/issues/555))
- Add awesome-go badge ([#554](https://github.com/vchain-us/immudb/issues/554))
- add history pagination
- add reverse zscan and reverse history pagination
- add atomic operations method
- **embedded/htree:** reusable build-at-once hash tree
- **embedded/store:** add Alh method to get last committed tx ID and alh ([#570](https://github.com/vchain-us/immudb/issues/570))
- **pkg/client:** add SetAll operation for simple  multi-kv atomic insertion ([#556](https://github.com/vchain-us/immudb/issues/556))
- **pkg/client:** sdk support index reference resolving)
- **pkg/client:** sdk support execAllOps
- **pkg/client/auditor:** enhance auditor to always send audit notification (even when no tampering was detected) if a notification URL is specified
- **pkg/store:** sorted sets support multiple equal keys with same score
- **pkg/store:** reference support index resolution
- **server:** add --audit-databases optional auditor flag

### Reverts
- chore: increase testing timeout


<a name="v0.8.0"></a>
## [v0.8.0] - 2020-09-15
### Bug Fixes
- **pkg/client:** setBatch creates structured values

### Changes
- fix immugw dependency to support new root structure
- update readme, add immudb4j news ([#488](https://github.com/vchain-us/immudb/issues/488))
- update README file ([#487](https://github.com/vchain-us/immudb/issues/487))
- switching README.md end lines to LF
- **cmd:** add signingKey flag
- **cmd:** remove error suppression in config loader
- **cmd/immutest/command:** remove immugw dependency from immutest
- **pkg:** add kvlist validator ([#498](https://github.com/vchain-us/immudb/issues/498))
- **pkg/server:** log uuid set and get error
- **pkg/server:** log signer initialization in immudb start

### Code Refactoring
- wrap root hash and index in a new structure to support signature
- move immugw in a separate repository
- **pkg/server:** inject root signer service inside immudb server

### Features
- auditor verifies root signature
- **pkg:** add root signer service
- **pkg/signer:** add ecdsa signer


<a name="v0.7.1"></a>
## [v0.7.1] - 2020-08-17
### Bug Fixes
- fix immudb and immugw version and mangen commands errors Without this change, while immuclient and immuadmin still worked as expected, immudb and immugw version and mangen commands were throwing the following error: ./immugw version Error: flag accessed but not defined: config Usage:   immugw version [flags]
- fix immuclient audit-mode
- **cmd/immuadmin/command:** fix immuadmin dbswitch
- **pkg/client:** token service manages old token format

### Code Refactoring
- configs file are loaded in viper preRun method

### Features
- **cmd:** process launcher check if are present another istances. fixes [#168](https://github.com/vchain-us/immudb/issues/168)


<a name="v0.7.0"></a>
## [v0.7.0] - 2020-08-10
### Bug Fixes
- fix travis build sleep time
- chose defaultdb on user create if not in multiple db mode
- use the correct server logger and use a dedicated logger with warning level for the db store
- use dedicated logger for store
- fix compilation error in corruption checker test
- user list showing only the superadmin user even when other user exist
- fix multiple services config uninstall
- userlist returns wrong message when logged in as immudb with single database
- race condition in token eviction
- skip loading databases from disk when in memory is requested
- remove os.Exit(0) from disconnect method
- fix DefaultPasswordReader initialization. fixes [#404](https://github.com/vchain-us/immudb/issues/404)
- if custom port is <= 0 use default port for both immudb and immugw
- fix immugw failing to start with nil pointer dereference since gRPC dial options are inherited (bug was introduced in commit a4477e2e403ab35fc9392e0a3a2d8436a5806901)
- **cmd/immuadmin/command:** fix user list output to support multiple databases (with permissions) for the same user
- **pkg/auth:** if new auth token is found in outgoing context it replaced the old one
- **pkg/client:** use database set internally database name
- **pkg/client:** inherit dial options that came from constructor
- **pkg/fs:** don't overwrite copy error on Close malfunction. Sync seals the operation–not Close.
- **pkg/gw:** fix client option construction with missing homedirservice
- **pkg/server:** avoid recursion on never ending functionality. Further improvements can be done ([#427](https://github.com/vchain-us/immudb/issues/427))
- **pkg/server:** added os file separator and db root path
- **pkg/server/server:** change user pass , old password check
- **pkg/service:** restore correct config path
- **pkg/store:** fix count method using a proper NewKeyIterator

### Changes
- refactor immuclient test
- versioning token filename
- add use database gw handler
- spread token service usage
- add homedir service
- rewrite tests in order to use pkg/server/servertest
- remove permission leftovers and useless messages in client server protocol
- add os and filepath abstraction and use it in immuadmin backup command
- using cobra command std out
- add codecov windows and freebsd ignore paths
- fix typo in UninstallManPages function name
- fix conflicts while rebasing from master
- remove user commands from immuclient
- add coveralls.io stage
- fix codecov ignore paths
- improve command ux and fix changepassword test. Closes [#370](https://github.com/vchain-us/immudb/issues/370)
- remove os wrapper from codecov.yml
- remove useless quitToStdError and os.exit calls
- add options to tuning corruption checking frequency, iteration and single iteration
- enhance immudb server messages during start
- capitalize immudb stop log message for consistency reasons
- move immuadmin and immuclient service managing to pkg
- log immudb user messages during start to file if a logfile is specified
- use debug instead of info for some log messages that are not relevant to the user
- add auditor single run mode
- add unit tests for zip and tar
- fix test
- refactor immuadmin service to use immuos abstraction
- add coverall badge
- add filepath abstration, use it in immuadmin backup and enhance coverage for backup test
- change insert user to use safeset instead of set
- fix tokenService typos
- fix go test cover coverall
- remove tests from windows CI
- fix immuclient tests
- add empty clientTest constructor
- user list client return a printable string
- add unexpectedNotStructuredValue error. fixes [#402](https://github.com/vchain-us/immudb/issues/402)
- add go-acc to calculate code coverage and fix go version to 1.13
- fix contributing.md styling
- remove sleep from tests
- use 0.0.0.0 instead of 127.0.0.1 as default address for both immudb and immugw
- add failfast option in test command
- add an explicit data source on terminal reader
- TestHealthCheckFails if grpc is no fully closed
- refactor immuclient test, place duplicated code in one place
- **cmd:** restore error handling in main method
- **cmd:** token is managed as a string. fixes [#453](https://github.com/vchain-us/immudb/issues/453)
- **cmd:** fix typo in command messages
- **cmd:** enhance PrintTable function to support custom table captions and use such captions in immuadmin user and database list commands
- **cmd:** immugw and immudb use process launcher for detach mode
- **cmd/helper:** add doc comment for the PrintTable function
- **cmd/immuadmin:** immuadmin user sub-commands use cobra, tests
- **cmd/immuadmin/command:** move command line and his command helper method in a single file
- **cmd/immuadmin/command:** fix text alignment and case
- **cmd/immuadmin/command:** user and database list use table printer
- **cmd/immuadmin/command:** remove silent errors in immuadmin
- **cmd/immuadmin/command:** remove useless auth check in print tree command
- **cmd/immuadmin/command:** automatically login the immuadmin user after forced password change is completed
- **cmd/immuadmin/command:** move options as dependency of commandline struct
- **cmd/immuclient/command:** remove useless comment
- **cmd/immuclient/immuc:** inject homedir service as dependency
- **cmd/immugw/command:** use general viper.BindPFlags binding instead of a verbose bindFlags solution
- **cmd/immutest/command:** inject homedir service as dependency
- **pkg/client/options:** add options fields and test
- **pkg/client/timestamp:** removed unused ntp timestamp
- **pkg/fs:** create file copy with flags from the start, in write-only mode
- **pkg/fs:** traceable copy errors
- **pkg/fs:** utilise filepath directory walk for copy
- **pkg/server:** improve tests
- **pkg/server:** add corruption checker random indexes generator dependency
- **pkg/server:** add corruption checker random indexes generator  missing dependency
- **pkg/server:** mtls test certificates system db as immuserver property improve tests
- **pkg/server:** immudb struct implements immudbIf interface, fixes previous tests
- **pkg/server:** make DevMode default false and cleanup call to action message shwon right after immudb start
- **pkg/store/sysstore:** remove useless method

### Code Refactoring
- remove custom errors inside useDatabase and createDatabase services. Fixes [#367](https://github.com/vchain-us/immudb/issues/367)
- handle in idiomatic way errors in changePermission grpc service. Fixes [#368](https://github.com/vchain-us/immudb/issues/368)
- decoupled client options from server gateway constructor
- refactor detach() method in a process launcher service
- decouple manpage methods in a dedicated service
- add immuadmin services interfaces and terminal helper
- **cmd:** move database management commands from immuclient to immuadmin. Fixes [#440](https://github.com/vchain-us/immudb/issues/440)
- **cmd/immuadmin/command:** using c.PrintfColorW instead c.PrintfColor to increase cobra.cmd integration for tests
- **cmd/immuadmin/command:** move checkLoggedInAndConnect, connect, disconnect from server to login file
- **cmd/immuadmin/command:** remove useless argument in Init and improve naming conventions

### Features
- add multiple databases support
- **cmd/helper:** add table printer
- **cmd/helper:** add PrintfColorW to decouple writer capabilities
- **cmd/immutest:** allow immutest to run on remote server
- **pkg/client:** add token service


<a name="v0.6.2"></a>
## [v0.6.2] - 2020-06-15
### Bug Fixes
- only require admin password to be changed if it is "immu"
- require auth for admin commands even if auth is disabled on server, do not allow admin user to be deactivated
- base64 decoding of passwords: now it requires the "enc:" prefix as base64 can not be differentiated from plain-text at runtime (e.g. "immu" is a valid base64 encode string)
- fix ldflags on dist binaries and add static compilation infos
- **cmd/immuclient/audit:** fix base64 encoded password not working with immuclient audit-mode
- **immuadmin:** repair password change flow right after first admin login
- **pkg/auth:** make ListUsers require admin permissions
- **pkg/ring:** fixes cache corruption due to a ring buffer elements overwrite  on same internal index
- **pkg/store:** remove useless ringbuffer array
- **pkg/store:** fix uniform cache layers size allocation with small values

### Changes
- fix golint errors
- githubactions add windows and build step
- remove plain-test admin password from log outputs
- add message (in cli help and swagger description) about base64-encoded inputs and outputs of get and set commands
- FreeBSD section in the readme
- add bug and feature request report github template
- fix changelog auto generation repo and releasing template
- **pkg/server:** reduce corruption_checker resources usage

### Features
- expose through REST the following user-related actions: create, get, list, change password, set permission and deactivate
- immuclient freebsd daemon installation
- freebsd service install
- read immudb default admin password from flag, config or env var
- use immu as default admin password instead of randomly generated one
- **immudb:** accept base64 string for admin password in flag/config/env var


<a name="v0.6.1"></a>
## [v0.6.1] - 2020-06-09
### Bug Fixes
- disallow running immuadmin backup with current directory as source
- choose correct config for immudb, immugw installation
- update env vars in README and Docker files ([#297](https://github.com/vchain-us/immudb/issues/297))
- fix corruption checker crash during immudb shoutdown
- immuadmin dump hangs indefinitely if token is invalid
- [#283](https://github.com/vchain-us/immudb/issues/283), immudb crash on dump of empty db
- **cmd/immuadmin:** validate backup dir before asking password
- **cmd/immuadmin:** inform user that manual server restart may be needed after interrupted backup
- **cmd/immuclient:** nil pointer when audit-mode used with immudb running as daemon
- **cmd/immuclient:** add version sub-command to immuclient interractive mode
- **cmd/immutest:** add new line at the end of output message
- **pkg/ring:** return nil on inconsistent access to buffer rings elements
- **pkg/store:** fix visualization of not frozen nodes inside print tree command
- **pkg/store/treestore:** fix overwriting on not freezes nodes

### Changes
- update statement about traditional DBs in README
- remove immugw configs from immudb config file [#302](https://github.com/vchain-us/immudb/issues/302)
- add license to tests ([#288](https://github.com/vchain-us/immudb/issues/288))
- **cmd/immuadmin/command:** improve visualization ui in merkle tree print command
- **cmd/immuadmin/command/service:** syntax error, fail build on windows
- **cmd/immuclient/audit:** code cleanup and renaming
- **pkg/store/treestore:** improve cache invalidation

### Code Refactoring
- handling of failed dump

### Features
- add auth support to immutest CLI
- add server-side logout ([#286](https://github.com/vchain-us/immudb/issues/286))
- allow the password of immugw auditor to be base64 encoded in the config file ([#296](https://github.com/vchain-us/immudb/issues/296))
- **cmd/helper:** add functionalities to print colored output
- **cmd/immuadmin:** add print tree command
- **cmd/immutest:** add env var for tokenfile
- **pkg:** add print tree functionality


<a name="v0.6.0"></a>
## [v0.6.0] - 2020-05-28
### Bug Fixes
- admin user can change password of regular user without having to know his old password
- when fetching users, only fetch the latest version
- safereference_handler, add tests [#264](https://github.com/vchain-us/immudb/issues/264)
- safeset_handler test
- readme doc, immugw start command
- typos in immugw help
- licence
- modify BUILT_BY flag with user email to keep dist script functionalities in makefile
- [#260](https://github.com/vchain-us/immudb/issues/260)
- various permissions-related issues
- immugw pid path consistency
- SafeZAdd handler SafeZAdd tests. Fix ReferenceHandler test
- fix immuclient windows build
- fix bug on zadd server method
- race condition while prefixing keys
- implementation of user deactivate
- rewrite user management to store user, password and permissions separately
- use iota for permissions enum
- **cmd/helper:** fix osx build
- **cmd/immuadmin/command/service:** fix error returned by GetDefaultConfigPath
- **cmd/immuadmin/command/service:** fix immudb data uninstall
- **cmd/immuclient:** Added missing documentations and renamed deprecated structures.
- **cmd/immuclient:** Fixed paths.
- **cmd/immuclient:** Added missing documentations and renamed deprecated structures.
- **cmd/immuclient:** Fixed wrong audit credentials error
- **cmd/immuclient/audit:** fix immuclient service installation
- **cmd/immuclient/service:** fix config import

### Changes
- add changelog
- add getByRawSafeIndex tests
- improve help for immugw auditor metrics
- rename audit(or) to trust-check(er)
- use status.Error instead of status.Errorf for static string
- use Sprintf instead of string concat
- rename default immudb and immugw loggers
- turn sys keys prefixes into constants
- remove setup release in makefile
- service_name inside release build script  is configurable inside makefile. closes [#159](https://github.com/vchain-us/immudb/issues/159) closes [#239](https://github.com/vchain-us/immudb/issues/239)
- remove ppc and arm target arch from makefile
- add CD releases, certificate sign, vcn sign in makefile dist scripts
- add dist scripts in makefile
- fix typo in README.md
- extract root service from immugw trust checker
- move corruption checker inside immudb process
- rename back immugw "trust checker" to "auditor"
- update docker files
- immugw audit publishes -1 if empty db and -2 if error, otherwise 0 (check failed) or 1 (succeeded)
- immugw audit publishes -1 value for result and root indexes in case the audit could not run (i.e. empty database, error etc.)
- change immugw metrics port
- refactoring file cache for immugw auditor
- rename immugw trust-checker to auditor
- move auditor package under client directory
- **cmd:** fix corruption checker flag
- **cmd/helper:** remove useless var
- **cmd/helper:** fix config path manager stub on linux
- **cmd/helper:** add path os wildcard resolver
- **cmd/immuadmin:** path of service files and binaries are os dynamic
- **cmd/immuclient:** add pid file management on windows
- **immuadmin:** improve the very first login message

### Code Refactoring
- refactor safeset_handler_test

### Features
- Audit agent added to immuclient.
- make metrics server start configurable through options to aid tests. MetricsServer must not be started as during tests because prometheus lib panis with: duplicate metrics collector registration attempted.
- invalidate tokens by droping public and private keys for a specific user
- check permissions dynamically
- implement user permissions and admin command to set them
- prefix user keys
- update metrics from immugw auditor
- add immugw auditor
- **cmd/immuclient/command:** add getByRawSafeIndex method
- **immugw:** add GET /lastaudit on metrics server


<a name="v0.6.0-RC2"></a>
## [v0.6.0-RC2] - 2020-05-19
### Bug Fixes
- fix stop, improve trust checker log
- handling immudb no connection error and comments
- **cmd/immuadmin:** old password can not be empty when changing password
- **cmd/immuadmin/command:** remove PID by systemd directive
- **cmd/immuadmin/command:** do not erase data without explicit consensus. closes 165
- **cmd/immuadmin/command/service:** fix [#188](https://github.com/vchain-us/immudb/issues/188)
- **cmd/immuclient:** correct argument index for value in rawsafeset
- **cmd/immutest:** rename immutestapp files to immutest
- **pkg/server:** fix error when unlocking unlocked stores after online db restore
- **pkg/store:** wait for pending writes in store.FlushToDisk

### Changes
- fix useless checks, binding address
- fix useless viper dependency
- update dockerfiles
- fix travis build
- manage dir flag in immutc
- add immutc makefile and remove bm from makeall
- add copyrights to makefile. closes [#142](https://github.com/vchain-us/immudb/issues/142)
- fix immugw dockerfile with dir property, update README
- use empty struct for values in map that store admin-only methods and add also Backup and Restore methods
- remove online backup and restore features
- **cmd:** remove useless exit call
- **cmd/immuadmin:** fix typo in todo keyword
- **cmd/immuadmin:** add todos to use the functions from fs package in immuadmin service helpers
- **cmd/immuadmin:** rename offline backup and restore to reflect their offline nature
- **cmd/immugw:** add dir property with default
- **pkg/client:** fix client contructor in tests
- **pkg/client:** add dir property with default
- **pkg/client:** fix ByRawSafeIndex method comment
- **pkg/gw:** remove useless root service dependency
- **pkg/gw:** add dir property with default, fix error messages
- **pkg/immuadmin:** use ReadFromTerminalYN function to get user confirmation to proceed
- **pkg/store:** fix typo in tamper insertion order index error message
- **server:** do not close the stores during cold Restore
- **server:** check for null before closing stores during backup and return absolute backup path

### Features
- show build time in user timezone in version cmd output
- set version to latest git tag
- added interactive cli to immuclient
- **cmd/immutc:** add trust checker command
- **immuadmin:** add offline backup and restore option with possibility to stop and restart the server manually
- **immuadmin:** add cold backup and restore
- **pkg/api/schema:** add byRawSafeIndex proto definitions and related parts
- **pkg/client:** add byRawSafeIndex client
- **pkg/server:** add byRawSafeIndex server
- **pkg/store:** add byRawSafeIndex methods and relateds parts
- **pkg/tc:** add trust checker core


<a name="v0.6.0-RC1"></a>
## [v0.6.0-RC1] - 2020-05-11
### Bug Fixes
- fix bug related to retrieve a versioned element by index
- return verified item on safeset
- wrap root_service with mutex to avoid dirty read
- place password input on the same line with the password input label
- store auth tokens in user home dir by default and other auth-related enhancements
- return correct error in safeZAdd handler
- disabling CGO to removes the need for the cross-compile dependencies
- remove useless error. https://github.com/dgraph-io/badger/commit/c6c1e5ec7690b5e5d7b47f6ab913bae6f78df03b
- add immugw exec to docker image
- upadating takama/daemon to fix freebsd compilation. closes [#160](https://github.com/vchain-us/immudb/issues/160)
- split main fails in separate folders
- use grpc interceptors for authentication
- fix structured values integration
- code format for multiple files to comply with Go coding standards
- fix immugw immud services in windows os
- fix env vars. Switch to toml
- prevent immuadmin users other than immu to login
- env vars names for immudb port and address in CLIs help
- enhance authentication and user management
- environment variables
- improving config management on linux and improved usage message
- use arguments instead of flags for user create, change-password and deleted
- fix reading user input for passwords
- change user message when new password is identical to the old one
- remove common package
- **/pkg/gw:** manage 0 index value in safeget
- **/pkg/gw:** fix guard
- **/pkg/gw:** manage 0 index value in history
- **cmd:** get and safeget error for non-existing key
- **cmd:** remove short alias for tokenfile flag
- **cmd/immu:** fix backup file opening during restore
- **cmd/immu:** Fix newline at the end of restore/backup success message
- **cmd/immu:** set auth header correctly
- **cmd/immuadmin:** verify old password immediately during change-password flow
- **cmd/immuadmin:** generate admin user if not exists also at 1st login attempt with user immu
- **cmd/immuadmin:** fix uninstall automatic stop
- **cmd/immuadmin/command/service:** fix config files in windows
- **cmd/immuadmin/command/service:** fix windows helper import
- **cmd/immuadmin/command/service:** fix group creation
- **immuclient:** do not connect to immudb server for version or mangen commands
- **immudb/command:** fix command test config file path
- **immupopulate:** do not connect to immudb server for version or mangen commands
- **pkg/api/schema:** fix typos
- **pkg/client:** fix cleaning on client tests
- **pkg/client:** fix stream closing to complete restore
- **pkg/client:** fix root file management
- **pkg/client/cache:** manage concurrent read and write ops
- **pkg/gw:** ensure computed item's matches the proof one for safeset
- **pkg/gw:** improve immugw logging
- **pkg/gw:** fix hash calc for reference item
- **pkg/gw:** fix error handling in safe method overwrite
- **pkg/gw:** manage concurrent safeset and get requests
- **pkg/gw:** refactor overwrite safe set and get request in order to use dependencies
- **pkg/gw:** use leaf computed from the item
- **pkg/gw:** fix lock release in case of errors
- **pkg/gw:** fix gw stop method
- **pkg/server:** correct error checking
- **pkg/server:** improve immudb logging
- **pkg/store:** correct health check error comparison
- **pkg/store:** correct gRPC code for key not found error
- **pkg/store:** badger's errors mapping
- **pkg/store:** truncate set to true for windows
- **pkg/store:** fix [#60](https://github.com/vchain-us/immudb/issues/60).

### Changes
- remove immuclient from default make target
- Print immud running infos properly
- simplify error during first admin login attempt
- marshal options to JSON in String implementations
- simplify .gitignore
- grpc-gateway code generation
- update .gitignore
- add missing dep
- update deps
- Add swagger generation command
- switch to 2.0.3 go_package patched badger version
- import badger protobuffer schema
- Update makefile in order to use badeger on codenotary fork
- fix compilation on OS X
- Switch services default ports
- improve configuration features on services management
- prefix version-related variable holding common ldflags so that it matches the convention used for the other version-related variables
- update default dbname in server config
- add pid params in config
- fix typo in dump command help
- remove default version value from immu* commands
- change default output folder for  man pages generation
- rename immupopulate to immutestapp
- remove app names from ldflags in Makefile, update immudb description in help
- use all lowercase for immudb everywhere it is mentioned to the user
- Switch config format to toml
- Fix namings conventions on immud command config properties
- instructions after make
- move token file name into options and some cleanup
- Remove man pages
- merge code changes related to histograms and detached server options
- rename immutestapp to immutest
- change label from "latency" to "duration" in immuadmin statistics
- remove types.go from immuadmin statistics cmd
- rename functions that update metrics
- integrate latest changes from master branch
- update termui transitive dependencies
- move immuadmin metrics struct to dedicated file
- rewrite immuadmin client to align it with latest refactorings
- fix project name in CONTRIBUTING.md
- Suppression of ErrObsoleteDataFormat in order to reduce breaking changes
- fix typo in raw command usage help
- rename backup to dump, and disable restore
- info if starting server with empty database
- use exact number of args 2 for set and safeset
- Set correct data folder and show usage in config. closes [#37](https://github.com/vchain-us/immudb/issues/37)
- Fix coding style
- Switch config format to ini
- refactor immuadmin and stats command to allow dependency injection of immu client
- protect default data folder from being inserted in repo
- change config path in server default options
- change default immudb data folder name from "immudb" to "immudata"
- rename binaries and configs
- linux services are managed by immu user
- rename test config file
- remove codenotary badger fork requirement
- rename command "consistency" to "verify"
- remove codenotary badger fork requirement
- rename command "verify" to "check-consistency"
- simplify auth options in immu client
- improve login help and cleanup irelevant TODO comment
- add host and version to swagger json
- move server password generation to server start
- get rid of locks on immuclient during login and user during set password
- get rid of password generating library
- change auth header to string
- **cmd:** addutility to manage a y/n dialog
- **cmd:** enhance unauthenticated message
- **cmd:** add env vars to commands help and man
- **cmd/immu:** Add reference in command line
- **cmd/immuadmin:** remove duplicate dependency
- **cmd/immuadmin:** remove ValidArgsFunction from user sub-command
- **cmd/immuadmin:** fix build on freebsd
- **cmd/immuadmin:** improve code organization and help messages
- **cmd/immuadmin:** extract to functions the init and update of plots in visual stats
- **cmd/immuadmin:** remove log files in uninstall
- **cmd/immuadmin:** improved immuadmin service ux
- **cmd/immuadmin:** remove unused varialble in user command
- **cmd/immuadmin:** fix examples alignment in user help
- **cmd/immuadmin:** set titles to green and use grid width instead of terminal width to determine plots data length
- **cmd/immuadmin/commands:** fix typo in error message and remove useless options
- **cmd/immuadmin/commands:** fix empty imput and improve immugw install ux
- **cmd/immugw:** Use default options values
- **cmd/immugw:** overwrite safeZAdd default handler
- **cmd/immugw:** overwrite safeReference default handler
- **immuclient:** move pre and post run callbacks to sub-commands
- **pkg/auth:** improve local client detection
- **pkg/client:** add reference client command
- **pkg/client:** add ZAdd and ZScan client methods
- **pkg/gw:** fix default config path for immugw
- **pkg/gw:** remove useless check on path
- **pkg/gw:** manage panic into http error
- **pkg/gw:** refactor handlers in order to use cache adapters
- **pkg/server:** return descriptive error if login gets called when auth is disabled
- **pkg/server:** keep generated keys (used to sign auth token) only in memory
- **pkg/store:** switch to gRPC errors with codes

### Code Refactoring
- refactor  packages to expose commands
- remove immuclient initialization from root level command
- Removed needless allocations and function calls, Rewrote Immuclient package layout
- config is managed properly with cobra and viper combo. closes [#44](https://github.com/vchain-us/immudb/issues/44)
- Structured immugw and handling SIGTERM. closes [#33](https://github.com/vchain-us/immudb/issues/33)
- pkg/tree got ported over to its own external repo codenotary/merkletree
- **pkg/store:** prefix errors with Err

### Features
- add mtls to immugw
- add mtls to immud
- add version to all commands
- add mtls certificates generation script
- create a new build process [#41](https://github.com/vchain-us/immudb/issues/41)
- Add config file. Closes [#36](https://github.com/vchain-us/immudb/issues/36) closes [#37](https://github.com/vchain-us/immudb/issues/37)
- always use the default bcrypt cost when hashing passwords
- implement user management
- Add capabilities to run commands in background. Closes [#136](https://github.com/vchain-us/immudb/issues/136) closes [#106](https://github.com/vchain-us/immudb/issues/106)
- hide some of the widgets in immuadmin statistics view if the server does not provide histograms
- add safeget, safeset, safereference and safezadd to the CLI client
- complete implementation of visual statistics in immuadmin
- change client "last query at" label
- add "client last active at" metric
- add uptime to metrics
- add immuadmin-related rules to makefile
- close immuadmin visual statistics also on <Escape>
- add text and visual display options to immuadmin statistics
- Add multiroot management, Add client mtls, client refactor. closes [#50](https://github.com/vchain-us/immudb/issues/50) closes [#80](https://github.com/vchain-us/immudb/issues/80)
- add number of RCs per client metric
- improve metrics
- add immuadmin client (WiP)
- add Prometheus-based metrics
- Add raw safeset and safeget method
- Add IScan and improve ScanByIndex command. Closes [#91](https://github.com/vchain-us/immudb/issues/91)
- add insertion order index and tests. closes [#39](https://github.com/vchain-us/immudb/issues/39)
- add current command. Closes [#88](https://github.com/vchain-us/immudb/issues/88)
- Add structured values components
- structured value
- add --no-histograms option to server
- add config item to toggle token-based auth on or off
- add token based authentication
- Add config file to immu
- Add config and pid file to immugw
- **cmd:** make unauthenticated message user-friendly
- **cmd/immu:** enhance backup and restore commands by writing/reading proto message size to/from the backup file
- **cmd/immu:** Add ZAdd and ZScan command line methods
- **cmd/immu:** fix CLI output of safe commands
- **cmd/immu:** add safeget, safeset, safereference and safezadd commands to immu CLI
- **cmd/immu:** add backup and restore commands
- **cmd/immuadmin:** add services management subcommand
- **cmd/immuadmin:** add configuration management in service command
- **cmd/immud:** Add pid file parameter
- **cmd/immugw:** add logger
- **cmd/immugw:** add immugw command
- **cmd/immupopulate:** add immupopulate command
- **immuadmin:** show line charts instead of piecharts
- **immuadmin:** add dump command
- **pkg/api/schema:** add safeget set patterns export
- **pkg/api/schema:** add reference and safeReference grpc messages
- **pkg/api/schema:** add backup and restore protobuffer objects
- **pkg/api/schema:** add zadd, safezadd and zscan grpc messages
- **pkg/api/schema:** add autogenerated grpc gw code
- **pkg/client:** use dedicated structs VerifiedItem and VerifiedIndex as safe functions return type
- **pkg/client:** add root cache service
- **pkg/client:** add new RootService field to ImmuClient, initialize it in connectWithRetry and use it in Safe* functions
- **pkg/client:** move keys reading outside ImmuClient and also pass it context from outside
- **pkg/client:** add backup and restore and automated tests for them
- **pkg/client/cache:** add cache file adapter
- **pkg/gw:** add safeset get overwrites
- **pkg/gw:** add safeZAdd overwrite
- **pkg/gw:** add reference and safeReference rest endpoint
- **pkg/logger:** Add file logger
- **pkg/server:** add reference and safeReference handlers
- **pkg/server:** add ZAdd, safeZAdd and zscan handlers
- **pkg/server:** Add backup and restore handlers
- **pkg/server:** Add pid file management
- **pkg/store:** add reference and safeReference methods and tests
- **pkg/store:** add ZAdd, safeZAdd and ZScan methods and tests
- **pkg/store:** Add store backup and restore methods and tests


<a name="v0.0.0-20200206"></a>
## v0.0.0-20200206 - 2020-02-06
### Bug Fixes
- pin go 1.13 for CI
- client dial retry and health-check retry counting
- missing dependencies
- address parsing
- apply badger options in bm
- typo
- protobuf codegen
- server re-start
- rpc bm startup
- nil out topic on server shutdown
- get-batch for missing keys
- client default options
- **bm:** allow GC
- **bm/rpc:** await server to be up
- **client:** missing default options
- **db:** cannot call treestore flush on empty tree
- **db:** include key as digest input
- **db:** default option for sync write
- **db:** correct treestore data race
- **db:** correct tree cache positions calculation at boostrap
- **db:** correct tree width calculation at startup
- **db:** workaround for tree indexes when reloading
- **db:** correct cache pre-allocation
- **db/treestore:** revert debug if condition
- **pkg/server:** correct nil pointer dereferences
- **pkg/server:** correct nil pointer dereference on logging
- **pkg/store:** correct key prefix guard
- **pkg/store:** correct inclusion proof call
- **pkg/store:** add miss condition to limit scan result
- **server:** avoid premature termination
- **tools/bm:** stop on error and correct max batch size
- **tree:** correct root of an empty tree
- **tree:** benchmark timers
- **tree:** handle missing node in mem store
- **tree:** `at` cannot be equal to width

### Changes
- client option type
- client connection wording
- add license headers
- rename module
- editorconfig
- gitignore
- bumps badger to v2.0.1
- license headers
- gitignore immu binaries
- update copyright notice
- gitignore vendor
- remove stale files
- change default data dir to "immustore"
- gitignore bm binary
- gitignore bm binary
- gitignore working db directory
- logging cleanup + thresholds
- server refactoring
- immu client cleanup
- trigger CI on PRs
- docker cleanup
- server options refactoring
- remove dead client code
- server options refactoring
- license headers
- bump badger to latest v2
- api cleanup
- refactor server logging
- move client options to type
- move server options to type
- improved logging
- simplified client interface
- extract client health check error
- client error refactoring
- missing license headers
- refactor rpc bms
- rpc bm cleanup
- nicer error reporting on client failure
- minor cleanup in bench.py
- strip bm binary
- make all target
- clean bm binary
- improved server interface
- more load on rpc bm
- new db constructor with options
- update .gitignore
- **api:** add index in hash construction
- **bm:** fine tuning
- **bm/rpc:** print error
- **client:** add ByIndex rpc
- **cmd/immu:** align commands to new APIs
- **db:** add default logger
- **db:** treestore entries discarding and ordering
- **db:** fine tuning
- **db:** treestore improvements
- **db:** switch to unbalanced tree test
- **db:** return item index
- **db:** correct logging messages
- **db:** default options
- **immu:** improved membership verification
- **immu:** improve printing
- **logger:** expose logging level
- **pkg/api:** add Hash() method for Item
- **pkg/api/schema:** refactor bundled proof proto message
- **pkg/api/schema:** add Root message and CurrentRoot RPC
- **pkg/api/schema:** get new root from proof
- **pkg/api/schema:** relax Proof verify when no prev root
- **pkg/db:** split data and tree storage
- **pkg/store:** add error for invalid root index
- **pkg/store:** code improvements
- **pkg/tree:** add verification for RFC 6962 examples
- **pkg/tree:** improve comment
- **schema:** add index
- **server:** add ByIndex rpc
- **tools:** no logging needed for nimmu
- **tools:** benchmark
- **tools:** add nimmu hacking tool
- **tools:** armonize comparison settings
- **tools:** add makefile target for comparison
- **tools:** benchmark
- **tools/bm:** correct method signature to accomodate indexes
- **tools/nimmu:** improve input and output
- **tree:** add print tree helper for debugging
- **tree:** add map backed mem store for testing
- **tree:** add IsFrozen helper
- **tree:** add map store test
- **tree:** remove unnecessary int conversion
- **tree:** reduce testPaths
- **tree:** correct MTH test var naming
- **tree:** correct returned value for undefined ranges
- **tree:** clean up map store

### Code Refactoring
- rename module to immustore
- reviewed schema package
- rname "membership" to "inclusion"
- use []byte keys
- "db" pkg renamed to "store"
- logging prefixes and miscellany renamed according to immustore
- change APIs according to the new schema
- env vars prefix renamed to IMMU_
- **db:** use ring buffer for caching
- **db:** use Item on ByIndex API
- **db:** new storing strategy
- **pkg/tree:** improve coverage and clean up
- **server:** metrics prefix renamed to immud_
- **tree:** AppendHash with side-effect
- **tree:** testing code improvement
- **tree:** switch to unbalanced tree
- **tree:** reduce lookups
- **tree:** remove unnecessary Tree struct and correct int sizes
- **tree:** simplify Storer interface
- **tree:** trival optimization

### Features
- bm memstats
- set stdin support
- initial project skeleton
- initial draft for storing the tree alongside the data
- immu client
- Makefile
- basic protobuf schema
- poc transport format
- poc grpc server
- poc grpc client
- poc server wiring
- client connect and wait for health-check
- poc client wiring
- topic wiring
- poc topic get
- client-side performance logs
- server topic get wiring
- poc cli args
- transport schema cleanup
- docker
- server options
- client options
- server options for dbname
- immud command line args
- expose metrics
- add healtcheck command
- client cli args
- logging infrastructure
- client type
- treestore with caching
- no healthcheck/dial retry wait-period when set to 0
- make bm
- batch-set
- bm tooling
- bm improvements (rpc-style and in-process api calls)
- batch get requests
- CI action
- scylla comparison
- immud health-check
- return item index for get and set ops
- expose dial and health-check retry configuration
- apply env options for server
- apply env options for client
- server options from environment
- client options from environment
- client connect with closure
- client errors extracted
- named logger
- pretty-print client options
- stateful client connection
- server health-check
- db health-check
- set-batch bm
- client set-batch
- key reader
- extract rpc bm constructor
- **api:** add item and item list data structure
- **api:** add membership proof struct
- **client:** add history command
- **client:** membership command
- **cmd/immu:** scan and count commands
- **db:** get item by index
- **db:** add membership proof API
- **db:** async commit option
- **db:** add history API to get all versions of the same key
- **pkg/api/schema:** added Scan and Count RPC
- **pkg/api/schema:** consistency proof API
- **pkg/api/schema:** SafeGet
- **pkg/api/schema:** add SafeSet to schema.proto
- **pkg/api/schema:** ScanOptions
- **pkg/client:** Scan and Count API
- **pkg/client:** consistecy proof
- **pkg/server:** consistency proof RPC
- **pkg/server:** Scan and Count API
- **pkg/server:** CurrentRoot RPC
- **pkg/server:** SafeSet RPC
- **pkg/server:** SafeGet RPC
- **pkg/store:** SafeSet
- **pkg/store:** SafeGet API
- **pkg/store:** CurrentRoot API
- **pkg/store:** count API
- **pkg/store:** consistency proof
- **pkg/store:** range scan API
- **pkg/tree:** test case for RFC 6962 examples
- **pkg/tree:** consistency proof verification
- **pkg/tree:** consistency proof
- **pkg/tree:** Path to slice conversion
- **ring:** ring buffer
- **schema:** add message for membership proof
- **server:** add membership rpc
- **server:** support for history API
- **tree:** Merkle Consistency Proof reference impl
- **tree:** draft implementation
- **tree:** Merkle audit path reference impl
- **tree:** MTH reference impl
- **tree:** Storer interface
- **tree:** membership proof verification
- **tree:** audit path construction


[Unreleased]: https://github.com/vchain-us/immudb/compare/v1.9.7...HEAD
[v1.9.7]: https://github.com/vchain-us/immudb/compare/v2.0.0-RC1...v1.9.7
[v2.0.0-RC1]: https://github.com/vchain-us/immudb/compare/v1.9.6...v2.0.0-RC1
[v1.9.6]: https://github.com/vchain-us/immudb/compare/v1.9.5...v1.9.6
[v1.9.5]: https://github.com/vchain-us/immudb/compare/v1.9.4...v1.9.5
[v1.9.4]: https://github.com/vchain-us/immudb/compare/v1.9.3...v1.9.4
[v1.9.3]: https://github.com/vchain-us/immudb/compare/v1.9DOM.2...v1.9.3
[v1.9DOM.2]: https://github.com/vchain-us/immudb/compare/v1.9DOM.2-RC1...v1.9DOM.2
[v1.9DOM.2-RC1]: https://github.com/vchain-us/immudb/compare/v1.9DOM.1...v1.9DOM.2-RC1
[v1.9DOM.1]: https://github.com/vchain-us/immudb/compare/v1.9DOM.1-RC1...v1.9DOM.1
[v1.9DOM.1-RC1]: https://github.com/vchain-us/immudb/compare/v1.9DOM.0...v1.9DOM.1-RC1
[v1.9DOM.0]: https://github.com/vchain-us/immudb/compare/v1.9DOM...v1.9DOM.0
[v1.9DOM]: https://github.com/vchain-us/immudb/compare/v1.9.0-RC2...v1.9DOM
[v1.9.0-RC2]: https://github.com/vchain-us/immudb/compare/v1.9.0-RC1...v1.9.0-RC2
[v1.9.0-RC1]: https://github.com/vchain-us/immudb/compare/v1.5.0...v1.9.0-RC1
[v1.5.0]: https://github.com/vchain-us/immudb/compare/v1.5.0-RC1...v1.5.0
[v1.5.0-RC1]: https://github.com/vchain-us/immudb/compare/v1.4.1...v1.5.0-RC1
[v1.4.1]: https://github.com/vchain-us/immudb/compare/v1.4.1-RC1...v1.4.1
[v1.4.1-RC1]: https://github.com/vchain-us/immudb/compare/v1.4.0...v1.4.1-RC1
[v1.4.0]: https://github.com/vchain-us/immudb/compare/v1.4.0-RC2...v1.4.0
[v1.4.0-RC2]: https://github.com/vchain-us/immudb/compare/v1.4.0-RC1...v1.4.0-RC2
[v1.4.0-RC1]: https://github.com/vchain-us/immudb/compare/v1.3.2...v1.4.0-RC1
[v1.3.2]: https://github.com/vchain-us/immudb/compare/v1.3.2-RC1...v1.3.2
[v1.3.2-RC1]: https://github.com/vchain-us/immudb/compare/v1.3.1...v1.3.2-RC1
[v1.3.1]: https://github.com/vchain-us/immudb/compare/v1.3.1-RC1...v1.3.1
[v1.3.1-RC1]: https://github.com/vchain-us/immudb/compare/v1.3.0...v1.3.1-RC1
[v1.3.0]: https://github.com/vchain-us/immudb/compare/v1.3.0-RC1...v1.3.0
[v1.3.0-RC1]: https://github.com/vchain-us/immudb/compare/v1.2.4...v1.3.0-RC1
[v1.2.4]: https://github.com/vchain-us/immudb/compare/v1.2.4-RC1...v1.2.4
[v1.2.4-RC1]: https://github.com/vchain-us/immudb/compare/v1.2.3...v1.2.4-RC1
[v1.2.3]: https://github.com/vchain-us/immudb/compare/v1.2.3-RC1...v1.2.3
[v1.2.3-RC1]: https://github.com/vchain-us/immudb/compare/v1.2.2...v1.2.3-RC1
[v1.2.2]: https://github.com/vchain-us/immudb/compare/v1.2.1...v1.2.2
[v1.2.1]: https://github.com/vchain-us/immudb/compare/v1.2.0...v1.2.1
[v1.2.0]: https://github.com/vchain-us/immudb/compare/v1.2.0-RC1...v1.2.0
[v1.2.0-RC1]: https://github.com/vchain-us/immudb/compare/v1.1.0...v1.2.0-RC1
[v1.1.0]: https://github.com/vchain-us/immudb/compare/v1.0.5...v1.1.0
[v1.0.5]: https://github.com/vchain-us/immudb/compare/v1.0.1...v1.0.5
[v1.0.1]: https://github.com/vchain-us/immudb/compare/v1.0.0...v1.0.1
[v1.0.0]: https://github.com/vchain-us/immudb/compare/cnlc-2.2...v1.0.0
[cnlc-2.2]: https://github.com/vchain-us/immudb/compare/v0.9.2...cnlc-2.2
[v0.9.2]: https://github.com/vchain-us/immudb/compare/v0.9.1...v0.9.2
[v0.9.1]: https://github.com/vchain-us/immudb/compare/v0.9.0...v0.9.1
[v0.9.0]: https://github.com/vchain-us/immudb/compare/v0.9.0-RC2...v0.9.0
[v0.9.0-RC2]: https://github.com/vchain-us/immudb/compare/v0.9.0-RC1...v0.9.0-RC2
[v0.9.0-RC1]: https://github.com/vchain-us/immudb/compare/v0.8.1...v0.9.0-RC1
[v0.8.1]: https://github.com/vchain-us/immudb/compare/v0.8.0...v0.8.1
[v0.8.0]: https://github.com/vchain-us/immudb/compare/v0.7.1...v0.8.0
[v0.7.1]: https://github.com/vchain-us/immudb/compare/v0.7.0...v0.7.1
[v0.7.0]: https://github.com/vchain-us/immudb/compare/v0.6.2...v0.7.0
[v0.6.2]: https://github.com/vchain-us/immudb/compare/v0.6.1...v0.6.2
[v0.6.1]: https://github.com/vchain-us/immudb/compare/v0.6.0...v0.6.1
[v0.6.0]: https://github.com/vchain-us/immudb/compare/v0.6.0-RC2...v0.6.0
[v0.6.0-RC2]: https://github.com/vchain-us/immudb/compare/v0.6.0-RC1...v0.6.0-RC2
[v0.6.0-RC1]: https://github.com/vchain-us/immudb/compare/v0.0.0-20200206...v0.6.0-RC1

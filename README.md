# golang-sandbox
Sandbox for learning how to properly use golang and associated tooling

## Packages
|Package|Description|
|---|---|
|api-json|Learning how to work with json in a 'static api'|
|multi-directory|Creating a package with multiple directories and sub-packages|
|space-api|Calling a real api that returns JSON and using results.  Also figured out how to unit test and integration test|

## Pipelines
All pipelines (build, release, etc) are performed via GitHub Actions

|Action|Purpose|
|---|---|
|[ci-build](.github/workflows/ci-build.yml)|Continuous Integration build.  Builds and runs tests for each project|
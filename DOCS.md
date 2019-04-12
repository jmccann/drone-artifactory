---
date: 2016-01-01T00:00:00+00:00
title: Artifactory
author: jmccann
tags: [ publish, artifactory ]
repo: jmccann/drone-artifactory
logo: hipchat.svg
image: jmccann/drone-artifactory
---

The Artifactory plugin publishes artifacts to Artifactory.
The below pipeline configuration demonstrates complex usage:

```yaml
pipeline:
  artifactory:
    image: jmccann/drone-artifactory:3
    username: JohnDoe
    password: abcd1234
    url: https://myarti.com/artifactory
    actions:
      - action: delete
        path: libs-snapshot-local/${DRONE_TAG}/*
      - action: upload
        path: libs-snapshot-local/${DRONE_TAG}
        sources:
          - target/*.jar
          - target/*.war
          - dist/**/*.min.js
      - action: upload
        args_file: upload-args.json
      - action: set-prop
        props:
          - name: commit_sha
            value: ${DRONE_COMMIT}
          - name: something
            values:
              - multiple
              - values
```

## Params

You can override the default configuration with the following parameters:

* `url` - Artifactory URL
* `username` - Artifactory username
* `apikey` - Artifactory ApiKey
* `password` - Artifactory password (Not required if apikey is provided)
* `actions` - List of actions and their params to perform against artifactory

### Delete Params

* `path` - Target path pattern to delete files from.  Value can also be pre-generated in
prior step and written/read from `.artifactory_path` file.
* `dryrun` - Pretend to delete but don't actually delete
* `recursive` - Artifacts are also deleted from sub-folders of the source directory for upload.  Default: `true`
* `args_file` - path to file to load arguments instead from a pregenerated JSON file

### Upload Params

* `sources` - List of files to upload
* `path` - Target path to upload files to.  Value can also be pre-generated in
prior step and written/read from `.artifactory_path` file.
* `dryrun` - Pretend to upload but don't actually upload
* `flat` - Artifacts are uploaded to the exact target path specified and their hierarchy in the source file system is ignored.  Default: `true`
* `include_dirs` - Source path applies to bottom-chain directories and not only to files. Bottom-chain directories are either empty or do not include other directories that match the source path.  Default: `false`
* `recursive` - Artifacts are also collected from sub-folders of the source directory for upload.  Default: `true`
* `regexp` - Will interpret the sources as a regular expression.  Default: `false`
* `args_file` - path to file to load arguments instead from a pregenerated JSON file

All file paths must be relative to current project sources

## Secrets

Instead of configuring sensitive information in your `.drone.yml` file in
plain text you can use Drone secrets and substitute the values at runtime using string replacement.

Please see the [Drone documentation](http://docs.drone.io/manage-secrets/) to learn more about secrets.

```diff
pipeline:
  artifactory:
    image: jmccann/drone-artifactory:3
-   username: JohnDoe
-   password: abcd1234
+   secrets: [ artifactory_username, artifactory_password ]
    url: https://myarti.com/artifactory
    actions:
      - action: delete
        path: libs-snapshot-local/${DRONE_TAG}/*
      - action: upload
        path: libs-snapshot-local/${DRONE_TAG}
        sources:
          - target/*.jar
          - target/*.war
          - dist/**/*.min.js
```

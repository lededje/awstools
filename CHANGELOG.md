# Changelog

## v7.5.0 (2019-07-12)

**New**

* `iam-session`: Added support for `AWS_CONFIG_FILE` and `AWS_SHARED_CREDENTIALS_FILE`
* `kms-env`: Added refreshing the environment with `--refresh-action`, `--refresh-interval` and `--refresh-max-retries`

## v7.4.0 (2019-04-09)

**New**

* `iam-sync-users`:
  - Added logging
  - Added expiring accounts with `chage`
  - Added setting an invalid password for new user to support Ubuntu account unlocking

## v7.3.0 (2019-03-27)

**New**

* `iam-sync-users`: Manage sudo with IAM tags

## v7.2.0 (2019-03-26)

**New**

* `iam-sync-users`
  * Added locking users not in IAM with `--lock-missing` and `--lock-ignore-user`
  * Added setting groups from IAM tags with `--iam-tags-prefix`
  * Made sudo optional with `--sudo`/`--no-sudo`

## v7.1.0 (2019-03-19)

**New**

* `aws-dump`: Added LastUsed to metadata

**Fix**

* `elb-resolve-alb-external-url`: Fix incorrect resolution

## v7.0.0 (2019-01-17)

**Breaking changes**

* `aws-dump` changed the report argument format

**New**

* `aws-dump` added more resources



## v6.3.0 (2019-01-08)

**New**

* Added `ec2-describe-instances`
* `aws-dump` added more resources

## v6.2.0 (2019-01-07)

**New**

* Added `aws-dump`

## v6.1.0 (2018-12-21)

**New**

* Added `iam-sync-users`

## v6.0.0 (2018-12-17)

**Fix**

* `iam-session` returns the command exit code
* `kms-env` returns the command exit code

## v5.11.0 (2018-12-03)

**New**

* Added `--task-json` to `ecs-deploy`

**Fix**

* `ecs-deploy` would return before the deployment was completed

## v5.10.0 (2018-11-23)

**New**

* Added support for secrets manager in `kms-env`

## v5.9.0 (2018-10-01)

**New**

* Added `ecs-locate`

## v5.8.1 (2018-06-26)

**Fix**

* MFA support in `iam-auth-proxy`

## v5.8.0 (2018-06-26)

**New**

* Added `iam-auth-proxy`

## v5.7.0 (2018-06-19)

**New**

* Added `--version` to all commands except `ecs-dashboard`

## v5.6.0 (2018-06-19)

**New**

* Added SSM support in `kms-env`

## v5.5.0 (2018-06-13)

**New**

* Added `--dns-prefix` to `elb-resolve-alb-external-url` to filter the right dns when there is more than one alias for an ALB

## v5.4.1 (2018-06-11)

**Fix**

* Fix a crash when using `ec2-ip-from-name` when a terminated instance exists

## v5.4 (2018-05-29)

**New**

* `lambda-ping`: Ping a URL with lambda and publish a cloudwatch custom metric

## v5.3 (2018-05-29)

**New**

* `cloudwatch-put-metric-data`: New dimension argument

**Fixes**

* `ecs-dashboard`: Use the same code to open session as the other tools. Fixes an issue where role assumption wasn't working sometimes from ECS.

## v5.2 (2018-05-01)

**New**

* Added `kms-env`

## v5.1 (2018-04-29)

**New**

* Added support for profiles with MFA and role assumption in all tools
* Added a check for trailing spaces in env vars to catch copy paste mistakes
* Added `common.OpenSession` to make it easier to open a session with config

## v5.0 (2018-04-21)

**Breaking Changes**

* `ecs-run-task`: `--cluster-name` renamed to `--cluster` to match `ecs-deploy`

**New**

* Added support for role assumption and MFA

## v4.5 (2018-04-21)

**New**

* Added `iam-public-ssh-keys`

## v4.4 (2018-04-20)

**New**

* Added `iam-session`

## v4.3 (2018-01-29)

**Breaking Changes**

* Changed flag in `ecs-deploy`

**New**

* Added `ecr-get-login`

## v3.2 (2018-01-28)

No changelog for older versions, check commit logs

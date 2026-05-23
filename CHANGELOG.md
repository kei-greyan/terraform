## 1.6.0 (Unreleased)

NEW FEATURES:

* `terraform test`: New command for testing Terraform modules using HCL-based test files. ([#33454](https://github.com/opentofu/terraform/issues/33454))

ENHANCEMENTS:

* backend/s3: Added support for assuming a role using Web Identity Token credentials. ([#33135](https://github.com/opentofu/terraform/issues/33135))
* `terraform init`: Improved error messages when a required provider cannot be found in any configured source. ([#33637](https://github.com/opentofu/terraform/issues/33637))
* core: Improved performance of large plans with many resource instances by optimizing graph traversal. ([#33511](https://github.com/opentofu/terraform/issues/33511))

BUG FIXES:

* core: Fixed a panic that could occur when a module output value references a resource that has been removed. ([#33702](https://github.com/opentofu/terraform/issues/33702))
* `terraform plan`: Fixed incorrect diff display for sets containing objects with nested collections. ([#33598](https://github.com/opentofu/terraform/issues/33598))
* backend/remote: Fixed authentication token not being refreshed correctly during long-running operations. ([#33421](https://github.com/opentofu/terraform/issues/33421))

## 1.5.7 (September 27, 2023)

BUG FIXES:

* core: Fixed a crash when destroying resources that have preconditions or postconditions referencing other destroyed resources. ([#33677](https://github.com/opentofu/terraform/issues/33677))
* `terraform plan`: Fixed a case where sensitive values in provider configuration could be revealed in error messages. ([#33655](https://github.com/opentofu/terraform/issues/33655))

## 1.5.6 (September 14, 2023)

BUG FIXES:

* backend/s3: Fixed a regression introduced in 1.5.5 that caused state locking to fail when using DynamoDB with custom endpoints. ([#33600](https://github.com/opentofu/terraform/issues/33600))
* core: Fixed incorrect handling of `null` values in `for` expressions when used inside `dynamic` blocks. ([#33569](https://github.com/opentofu/terraform/issues/33569))

## 1.5.5 (August 23, 2023)

ENHANCEMENTS:

* backend/s3: Updated AWS SDK to v2, bringing improved authentication support including SSO and credential process. ([#33492](https://github.com/opentofu/terraform/issues/33492))

BUG FIXES:

* core: Fixed a panic when evaluating `templatefile` with a template that references an undefined variable. ([#33512](https://github.com/opentofu/terraform/issues/33512))
* `terraform show`: Fixed JSON output for plan files containing moved resource instances. ([#33480](https://github.com/opentofu/terraform/issues/33480))

---

For information on older releases, see [previous-releases.md](.changes/previous-releases.md).

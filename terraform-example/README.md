# A simple terrafrom example

This package has 2 main folders:
* `modules` - stores al lthe reusable components between different setups* `setups` - stores all the setup combinations. For example, `dev` and `prod`

## Notes
---
The `terrafrom apply` command should be run from the `./setups/{name}/` folders only, so there is a separate state per setup and they can be managed intependently.
---
Most of the differences between environments should be captured in the `locals` block. Also this example has differences in the `service` module versions
---

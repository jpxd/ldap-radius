# ldap-radius

a lightweight radius server written in go which uses ldap (and more) as authentication source

to configure you have to modify `config_example.gcfg` and copy it to `config.gcfg`

local authentication options can be added in the `checkLocal` function in `auth.go`

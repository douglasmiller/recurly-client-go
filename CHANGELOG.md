# Changelog

## [v4.0.1](https://github.com/recurly/recurly-client-go/tree/v4.0.1) (2021-03-23)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v4.0.0...v4.0.1)

**Merged pull requests:**

- Release 4.0.1 [\#91](https://github.com/recurly/recurly-client-go/pull/91) ([douglasmiller](https://github.com/douglasmiller))
- Generated Latest Changes for v2021-02-25 [\#90](https://github.com/recurly/recurly-client-go/pull/90) ([recurly-integrations](https://github.com/recurly-integrations))
- Export optionsApplier as OptionsApplier [\#81](https://github.com/recurly/recurly-client-go/pull/81) ([jguidry-recurly](https://github.com/jguidry-recurly))

## [v4.0.0](https://github.com/recurly/recurly-client-go/tree/v4.0.0) (2021-03-01)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.13.0...v4.0.0)

# Major Version Release

The 4.x major version of the client pairs with the `v2021-02-25` API version. This version of the client and the API contain breaking changes that should be considered before upgrading your integration.

## Breaking Changes in the API
All changes to the core API are documented in the [Developer Portal changelog](https://developers.recurly.com/api/changelog.html#v2021-02-25---current-ga-version)

## Breaking Changes in Client

- Empty path parameters are now explicitly invalid and cause an error to be returned.  [#67] [#68]
- `Params` has been removed in favor of `Context`/`RequestOptions`.  [#70]

    ### 3.x
    
    ```go
    headers := http.Header{"Accept-Language": []string{"fr"}}
    
    accountReq := &recurly.AccountCreate{
	    Params: recurly.Params{
		    Header: headers,
	    },
    }
    
    account, err := client.CreateAccount(accountReq)
    ```
    
    ### 4.x
    
    ```go
    headers := http.Header{"Accept-Language": []string{"fr"}}
    
    accountReq := &recurly.AccountCreate{}
    
    account, err := client.CreateAccount(accountReq, recurly.WithHeader(headers))
    ```

- List operations return a pager interface instead of struct.  [#80]

    ### 3.x
    
    ```go
    account, err := client.ListAccounts()

    for accounts.HasMore {
	    err := accounts.Fetch()
	    for i, account := range accounts.Data {
		    fmt.Printf("Account %3d", i)
	    }
    }
    ```
    
    ### 4.x
    
    ```go
    account, err := client.ListAccounts()
    
    for accounts.HasMore() {
	    err := accounts.Fetch()
	    for i, account := range accounts.Data() {
		    fmt.Printf("Account %3d", i)
	    }
    }
    ```

**Implemented enhancements:**

- FetchWithContext not exposed within ClientInterface [\#77](https://github.com/recurly/recurly-client-go/issues/77)

**Merged pull requests:**

- Release 4.0.0 [\#85](https://github.com/recurly/recurly-client-go/pull/85) ([douglasmiller](https://github.com/douglasmiller))
- Updating changelog script and changelog generator config for 4.x release [\#82](https://github.com/recurly/recurly-client-go/pull/82) ([douglasmiller](https://github.com/douglasmiller))
- Add pager interface [\#80](https://github.com/recurly/recurly-client-go/pull/80) ([jguidry-recurly](https://github.com/jguidry-recurly))
- Adding options v4 [\#70](https://github.com/recurly/recurly-client-go/pull/70) ([douglasmiller](https://github.com/douglasmiller))
- Validate path params [\#67](https://github.com/recurly/recurly-client-go/pull/67) ([douglasmiller](https://github.com/douglasmiller))

## [v3.13.0](https://github.com/recurly/recurly-client-go/tree/v3.13.0) (2021-01-22)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.12.0...v3.13.0)

**Implemented enhancements:**

- Latest Changes for 2019-10-10 [\#76](https://github.com/recurly/recurly-client-go/pull/76) ([douglasmiller](https://github.com/douglasmiller))

**Merged pull requests:**

- Release 3.13.0 [\#78](https://github.com/recurly/recurly-client-go/pull/78) ([douglasmiller](https://github.com/douglasmiller))

## [v3.12.0](https://github.com/recurly/recurly-client-go/tree/v3.12.0) (2020-12-09)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.11.0...v3.12.0)

**Implemented enhancements:**

- Latest Changes for 2019-10-10 [\#74](https://github.com/recurly/recurly-client-go/pull/74) ([douglasmiller](https://github.com/douglasmiller))
- Adding Context and RequestOptions support [\#69](https://github.com/recurly/recurly-client-go/pull/69) ([douglasmiller](https://github.com/douglasmiller))

**Merged pull requests:**

- Release 3.12.0 [\#75](https://github.com/recurly/recurly-client-go/pull/75) ([douglasmiller](https://github.com/douglasmiller))

## [v3.11.0](https://github.com/recurly/recurly-client-go/tree/v3.11.0) (2020-11-24)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.10.0...v3.11.0)

**Implemented enhancements:**

- Latest Changes for 2019-10-10 \(tax\_identifier / tax\_identifier\_type\) [\#72](https://github.com/recurly/recurly-client-go/pull/72) ([douglasmiller](https://github.com/douglasmiller))
- Latest Changes for 2019-10-10 [\#71](https://github.com/recurly/recurly-client-go/pull/71) ([douglasmiller](https://github.com/douglasmiller))

**Merged pull requests:**

- Release 3.11.0 [\#73](https://github.com/recurly/recurly-client-go/pull/73) ([douglasmiller](https://github.com/douglasmiller))
- Minor updates to facilitate 4.x changes [\#68](https://github.com/recurly/recurly-client-go/pull/68) ([douglasmiller](https://github.com/douglasmiller))

## [v3.10.0](https://github.com/recurly/recurly-client-go/tree/v3.10.0) (2020-11-06)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.9.0...v3.10.0)

**Implemented enhancements:**

- Latest Changes for 2019-10-10 \(Wallet, Item Coupons\) [\#65](https://github.com/recurly/recurly-client-go/pull/65) ([douglasmiller](https://github.com/douglasmiller))
- Updating ErrorTypes and status code based error handling to be generated [\#64](https://github.com/recurly/recurly-client-go/pull/64) ([douglasmiller](https://github.com/douglasmiller))

**Merged pull requests:**

- Release 3.10.0 [\#66](https://github.com/recurly/recurly-client-go/pull/66) ([douglasmiller](https://github.com/douglasmiller))

## [v3.9.0](https://github.com/recurly/recurly-client-go/tree/v3.9.0) (2020-10-20)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.8.0...v3.9.0)

**Implemented enhancements:**

- Mon Oct 19 20:43:56 UTC 2020 Upgrade API version v2019-10-10 [\#61](https://github.com/recurly/recurly-client-go/pull/61) ([douglasmiller](https://github.com/douglasmiller))

**Merged pull requests:**

- Release 3.9.0 [\#62](https://github.com/recurly/recurly-client-go/pull/62) ([douglasmiller](https://github.com/douglasmiller))
- Splitting up large files into many smaller files [\#60](https://github.com/recurly/recurly-client-go/pull/60) ([douglasmiller](https://github.com/douglasmiller))

## [v3.8.0](https://github.com/recurly/recurly-client-go/tree/v3.8.0) (2020-09-22)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.7.0...v3.8.0)

**Implemented enhancements:**

- Latest Changes for 2019-10-10 \(Automated Exports, additional resource data attributes\) [\#56](https://github.com/recurly/recurly-client-go/pull/56) ([douglasmiller](https://github.com/douglasmiller))
- Ensure that path parameters are not empty strings [\#21](https://github.com/recurly/recurly-client-go/pull/21) ([douglasmiller](https://github.com/douglasmiller))

**Merged pull requests:**

- Release 3.8.0 [\#57](https://github.com/recurly/recurly-client-go/pull/57) ([douglasmiller](https://github.com/douglasmiller))
- Revert "Ensure that path parameters are not empty strings" [\#53](https://github.com/recurly/recurly-client-go/pull/53) ([douglasmiller](https://github.com/douglasmiller))

## [v3.7.0](https://github.com/recurly/recurly-client-go/tree/v3.7.0) (2020-08-31)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.6.0...v3.7.0)

**Implemented enhancements:**

- Mon Aug 31 19:36:28 UTC 2020 Upgrade API version v2019-10-10 [\#51](https://github.com/recurly/recurly-client-go/pull/51) ([douglasmiller](https://github.com/douglasmiller))

**Merged pull requests:**

- Release 3.7.0 [\#52](https://github.com/recurly/recurly-client-go/pull/52) ([douglasmiller](https://github.com/douglasmiller))
- Code of Conduct [\#50](https://github.com/recurly/recurly-client-go/pull/50) ([bhelx](https://github.com/bhelx))

## [v3.6.0](https://github.com/recurly/recurly-client-go/tree/v3.6.0) (2020-08-21)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.5.0...v3.6.0)

**Implemented enhancements:**

- Fri Aug 21 16:27:20 UTC 2020 Upgrade API version v2019-10-10 [\#48](https://github.com/recurly/recurly-client-go/pull/48) ([douglasmiller](https://github.com/douglasmiller))

**Merged pull requests:**

- Release 3.6.0 [\#49](https://github.com/recurly/recurly-client-go/pull/49) ([douglasmiller](https://github.com/douglasmiller))
- Provide interface for testing [\#47](https://github.com/recurly/recurly-client-go/pull/47) ([bhelx](https://github.com/bhelx))

## [v3.5.0](https://github.com/recurly/recurly-client-go/tree/v3.5.0) (2020-07-31)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.4.0...v3.5.0)

**Implemented enhancements:**

- Latest Changes for 2019-10-10 \(usage, measured units, etc\) [\#45](https://github.com/recurly/recurly-client-go/pull/45) ([bhelx](https://github.com/bhelx))

**Merged pull requests:**

- Release 3.5.0 [\#46](https://github.com/recurly/recurly-client-go/pull/46) ([douglasmiller](https://github.com/douglasmiller))

## [v3.4.0](https://github.com/recurly/recurly-client-go/tree/v3.4.0) (2020-07-06)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.3.0...v3.4.0)

**Implemented enhancements:**

- Mon Jul  6 14:54:37 UTC 2020 Upgrade API version v2019-10-10 [\#43](https://github.com/recurly/recurly-client-go/pull/43) ([douglasmiller](https://github.com/douglasmiller))

**Merged pull requests:**

- Release 3.4.0 [\#44](https://github.com/recurly/recurly-client-go/pull/44) ([douglasmiller](https://github.com/douglasmiller))

## [v3.3.0](https://github.com/recurly/recurly-client-go/tree/v3.3.0) (2020-07-01)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.2.0...v3.3.0)

**Implemented enhancements:**

- Wed Jul  1 02:12:15 UTC 2020 Upgrade API version v2019-10-10 [\#41](https://github.com/recurly/recurly-client-go/pull/41) ([douglasmiller](https://github.com/douglasmiller))

**Merged pull requests:**

- Release 3.3.0 [\#42](https://github.com/recurly/recurly-client-go/pull/42) ([douglasmiller](https://github.com/douglasmiller))

## [v3.2.0](https://github.com/recurly/recurly-client-go/tree/v3.2.0) (2020-06-30)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.1.0...v3.2.0)

**Implemented enhancements:**

- Mon Jun 29 17:14:42 UTC 2020 Upgrade API version v2019-10-10 [\#39](https://github.com/recurly/recurly-client-go/pull/39) ([douglasmiller](https://github.com/douglasmiller))
- Cleaner implementation of --this-version [\#37](https://github.com/recurly/recurly-client-go/pull/37) ([bhelx](https://github.com/bhelx))

**Merged pull requests:**

- Release 3.2.0 [\#40](https://github.com/recurly/recurly-client-go/pull/40) ([douglasmiller](https://github.com/douglasmiller))
- Loosen restrictions on what is successful resp [\#38](https://github.com/recurly/recurly-client-go/pull/38) ([bhelx](https://github.com/bhelx))

## [v3.1.0](https://github.com/recurly/recurly-client-go/tree/v3.1.0) (2020-06-01)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.0.1...v3.1.0)

**Implemented enhancements:**

- Dont use absolute path to go bin in scripts [\#35](https://github.com/recurly/recurly-client-go/pull/35) ([bhelx](https://github.com/bhelx))
- Latest Features [\#34](https://github.com/recurly/recurly-client-go/pull/34) ([bhelx](https://github.com/bhelx))
- Add three\_d\_secure\_action\_token\_id  to TransactionError [\#32](https://github.com/recurly/recurly-client-go/pull/32) ([YevheniiGera](https://github.com/YevheniiGera))

**Merged pull requests:**

- Release 3.1.0 [\#36](https://github.com/recurly/recurly-client-go/pull/36) ([bhelx](https://github.com/bhelx))
- Fix pagination example [\#33](https://github.com/recurly/recurly-client-go/pull/33) ([bhelx](https://github.com/bhelx))

## [v3.0.1](https://github.com/recurly/recurly-client-go/tree/v3.0.1) (2020-05-20)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.0.0...v3.0.1)

**Fixed bugs:**

- Add version scoped url to go.mod and install instructions [\#30](https://github.com/recurly/recurly-client-go/pull/30) ([bhelx](https://github.com/bhelx))

**Merged pull requests:**

- Release 3.0.1 [\#31](https://github.com/recurly/recurly-client-go/pull/31) ([bhelx](https://github.com/bhelx))
- Fix release scripts and update readme [\#29](https://github.com/recurly/recurly-client-go/pull/29) ([bhelx](https://github.com/bhelx))
- Add github templates [\#28](https://github.com/recurly/recurly-client-go/pull/28) ([bhelx](https://github.com/bhelx))

## [v3.0.0](https://github.com/recurly/recurly-client-go/tree/v3.0.0) (2020-05-18)

[Full Changelog](https://github.com/recurly/recurly-client-go/compare/v3.0.0-beta.1...v3.0.0)

## Breaking Changes

If you are upgrading from 3.0.0-beta.1, there is [1 breaking change](https://github.com/recurly/recurly-client-go/pull/25/files). We have removed `DefaultClient` in place of `NewClient(string)`:

Change this:
```go
recurly.APIKey = "<apikey>"
client := recurly.DefaultClient()
```

To:
```go
client := recurly.NewClient("<apikey>")
```

**Implemented enhancements:**

- Add some more scripts and bump to 3.0.0 [\#26](https://github.com/recurly/recurly-client-go/pull/26) ([bhelx](https://github.com/bhelx))
- Adds `NewClient` and removes `DefaultClient` [\#25](https://github.com/recurly/recurly-client-go/pull/25) ([bhelx](https://github.com/bhelx))
- Only log warnings and errors [\#24](https://github.com/recurly/recurly-client-go/pull/24) ([bhelx](https://github.com/bhelx))
- Implement Pager.Count\(\) [\#23](https://github.com/recurly/recurly-client-go/pull/23) ([bhelx](https://github.com/bhelx))
- Add some tests [\#22](https://github.com/recurly/recurly-client-go/pull/22) ([joannasese](https://github.com/joannasese))

**Merged pull requests:**

- add beta warning [\#20](https://github.com/recurly/recurly-client-go/pull/20) ([bhelx](https://github.com/bhelx))



\* *This Changelog was automatically generated by [github_changelog_generator](https://github.com/github-changelog-generator/github-changelog-generator)*

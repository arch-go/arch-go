# Changelog

## [2.1.3](https://github.com/arch-go/arch-go/compare/v2.1.2...v2.1.3) (2026-02-13)


### Dependencies

* update module golang.org/x/mod to v0.33.0 ([#251](https://github.com/arch-go/arch-go/issues/251)) ([255616e](https://github.com/arch-go/arch-go/commit/255616eab13bb856ef77df56225a1c7836bb47fa))
* update module golang.org/x/tools to v0.42.0 ([#252](https://github.com/arch-go/arch-go/issues/252)) ([a390350](https://github.com/arch-go/arch-go/commit/a39035079967f0e516410a0be2b3848404a9f9fe))

## [2.1.2](https://github.com/arch-go/arch-go/compare/v2.1.1...v2.1.2) (2026-02-03)


### Bug Fixes

* resolve generic struct names correctly in naming rules ([#248](https://github.com/arch-go/arch-go/issues/248)) ([f7eb20f](https://github.com/arch-go/arch-go/commit/f7eb20f9ab03380635541b900db4d3f30a109e05))


### Dependencies

* update module github.com/agiledragon/gomonkey/v2 to v2.14.0 ([#241](https://github.com/arch-go/arch-go/issues/241)) ([206c3be](https://github.com/arch-go/arch-go/commit/206c3be540a21c5538abee7d4fc379801eb413bf))
* update module github.com/jedib0t/go-pretty/v6 to v6.7.7 ([#237](https://github.com/arch-go/arch-go/issues/237)) ([47ec243](https://github.com/arch-go/arch-go/commit/47ec243608c7fc58185b425e4ea09f77e3b849ce))
* update module github.com/jedib0t/go-pretty/v6 to v6.7.8 ([#240](https://github.com/arch-go/arch-go/issues/240)) ([827cda8](https://github.com/arch-go/arch-go/commit/827cda8bef7ff906280e12c4940e8d3cb6368938))
* update module github.com/spf13/cobra to v1.10.2 ([#232](https://github.com/arch-go/arch-go/issues/232)) ([fb6c2f5](https://github.com/arch-go/arch-go/commit/fb6c2f5227a5b667b04ccdbfa5d57ec6dadb2222))
* update module golang.org/x/mod to v0.32.0 ([#233](https://github.com/arch-go/arch-go/issues/233)) ([c3d3c83](https://github.com/arch-go/arch-go/commit/c3d3c8362746beca4408b8f02546098d300f385e))
* update module golang.org/x/tools to v0.41.0 ([#235](https://github.com/arch-go/arch-go/issues/235)) ([7e4f21d](https://github.com/arch-go/arch-go/commit/7e4f21d8f914244bbd4ddd881a663919f95efaa7))

## [2.1.1](https://github.com/arch-go/arch-go/compare/v2.1.0...v2.1.1) (2025-11-29)


### Bug Fixes

* check for race condition in test and use mutex for verification result ([#174](https://github.com/arch-go/arch-go/issues/174)) ([a954f5a](https://github.com/arch-go/arch-go/commit/a954f5a40e6d292e4f3b76bd0a3dbaa7e10dcd31))

## [2.1.0](https://github.com/arch-go/arch-go/compare/v2.0.1...v2.1.0) (2025-11-27)


### Features

* correct matching packages without suffix ([#218](https://github.com/arch-go/arch-go/issues/218)) ([8b91574](https://github.com/arch-go/arch-go/commit/8b915744ab18c4f75cb115abf8a9ef7eb40875d0))


### Dependencies

* update module github.com/jedib0t/go-pretty/v6 to v6.7.5 ([#225](https://github.com/arch-go/arch-go/issues/225)) ([ee77d3e](https://github.com/arch-go/arch-go/commit/ee77d3eee5dd792c8b16d395ed82c90c973ba4c6))
* update module github.com/stretchr/testify to v1.11.1 ([#204](https://github.com/arch-go/arch-go/issues/204)) ([1aab669](https://github.com/arch-go/arch-go/commit/1aab66973825202da729d52575e99ba3384a963d))

## [2.0.1](https://github.com/arch-go/arch-go/compare/v2.0.0...v2.0.1) (2025-11-26)


### Bug Fixes

* module name does not have major version ([#219](https://github.com/arch-go/arch-go/issues/219)) ([dc2b8fa](https://github.com/arch-go/arch-go/commit/dc2b8faf4ab0642f618a8fb28f3b236a3eee9dd8))

## [2.0.0](https://github.com/arch-go/arch-go/compare/v1.7.0...v2.0.0) (2025-09-28)


### ⚠ BREAKING CHANGES

* check standard and external interfaces in naming rules ([#179](https://github.com/arch-go/arch-go/issues/179))

### Features

* check standard and external interfaces in naming rules ([#179](https://github.com/arch-go/arch-go/issues/179)) ([aebabeb](https://github.com/arch-go/arch-go/commit/aebabebe4f85e97cde43a654351e1dee925af474))


### Bug Fixes

* Fixes broken test on concole report ([#189](https://github.com/arch-go/arch-go/issues/189)) ([9d959a7](https://github.com/arch-go/arch-go/commit/9d959a7dc7ece9ccec19b16723bf6444584a767f))


### Dependencies

* update module github.com/agiledragon/gomonkey/v2 to v2.13.0 ([#184](https://github.com/arch-go/arch-go/issues/184)) ([a61d975](https://github.com/arch-go/arch-go/commit/a61d9758b21cb48a06398dacdd934b5fc4562ef0))
* update module github.com/jedib0t/go-pretty/v6 to v6.6.7 ([#166](https://github.com/arch-go/arch-go/issues/166)) ([66e6a71](https://github.com/arch-go/arch-go/commit/66e6a717323f2f5223da994771720066a748771b))
* update module github.com/jedib0t/go-pretty/v6 to v6.6.8 ([#192](https://github.com/arch-go/arch-go/issues/192)) ([fa36505](https://github.com/arch-go/arch-go/commit/fa3650558cd3b44a260d7e586d993cf074aafb5f))
* update module github.com/spf13/cobra to v1.10.1 ([#202](https://github.com/arch-go/arch-go/issues/202)) ([af16fe4](https://github.com/arch-go/arch-go/commit/af16fe41132a25dd89d7526301efdadd364ca4e4))
* update module github.com/spf13/cobra to v1.9.1 ([#185](https://github.com/arch-go/arch-go/issues/185)) ([4f9be6f](https://github.com/arch-go/arch-go/commit/4f9be6f2a62b8545efbba72a380cb9b4f43a735a))
* update module github.com/spf13/viper to v1.20.1 ([#186](https://github.com/arch-go/arch-go/issues/186)) ([881b754](https://github.com/arch-go/arch-go/commit/881b754a7bea4f0359795eeaf89f626ae72011ee))

## [1.7.0](https://github.com/arch-go/arch-go/compare/v1.6.2...v1.7.0) (2024-12-13)


### Features

* check embedded interfaces ([#116](https://github.com/arch-go/arch-go/issues/116)) ([756f312](https://github.com/arch-go/arch-go/commit/756f312519b88570aabe9dbc6350dcf78740af22))
* Support for wildcard regex pattern in pkg names ([#131](https://github.com/arch-go/arch-go/issues/131)) ([e11d025](https://github.com/arch-go/arch-go/commit/e11d025d797f362c0d12e808154d0ae71a754639))


### Dependencies

* update module github.com/fatih/color to v1.18.0 ([#134](https://github.com/arch-go/arch-go/issues/134)) ([b707827](https://github.com/arch-go/arch-go/commit/b707827f8607ef866435de27f31c51c7e60f66ac))
* update module github.com/jedib0t/go-pretty/v6 to v6.6.1 ([#132](https://github.com/arch-go/arch-go/issues/132)) ([22e2447](https://github.com/arch-go/arch-go/commit/22e24475db362ee6b1ec6641c0faea22f5fe6c08))
* update module github.com/jedib0t/go-pretty/v6 to v6.6.2 ([#148](https://github.com/arch-go/arch-go/issues/148)) ([4afc91e](https://github.com/arch-go/arch-go/commit/4afc91e6d4c1c91f569d429647bc53ac649db676))
* update module github.com/jedib0t/go-pretty/v6 to v6.6.3 ([#154](https://github.com/arch-go/arch-go/issues/154)) ([71e0ce7](https://github.com/arch-go/arch-go/commit/71e0ce70e6b21af169075765a0d614d43ba0d842))
* update module github.com/jedib0t/go-pretty/v6 to v6.6.4 ([#161](https://github.com/arch-go/arch-go/issues/161)) ([f8282b4](https://github.com/arch-go/arch-go/commit/f8282b4b0c5421e7b6f381e88f27e914eaf46f9b))
* update module golang.org/x/mod to v0.22.0 ([#143](https://github.com/arch-go/arch-go/issues/143)) ([1b9a51d](https://github.com/arch-go/arch-go/commit/1b9a51d8e3eeb8907f64b2fc039d6a40ca9f326a))
* update module golang.org/x/tools to v0.26.0 ([#123](https://github.com/arch-go/arch-go/issues/123)) ([c10511c](https://github.com/arch-go/arch-go/commit/c10511cdd395a0342db725474c6d09882cc17176))
* update module golang.org/x/tools to v0.27.0 ([#145](https://github.com/arch-go/arch-go/issues/145)) ([1b40a7d](https://github.com/arch-go/arch-go/commit/1b40a7d27b40c8b82e3fbdb898a4faca4e539d1d))
* update module golang.org/x/tools to v0.28.0 ([#157](https://github.com/arch-go/arch-go/issues/157)) ([2b1f2a1](https://github.com/arch-go/arch-go/commit/2b1f2a13ed83d02b7ca52311e07fedff73e3800a))

## [1.6.2](https://github.com/arch-go/arch-go/compare/v1.6.1...v1.6.2) (2024-09-18)


### Bug Fixes

* Proper version for the built, respectively released binaries ([#113](https://github.com/arch-go/arch-go/issues/113)) ([bac7ba2](https://github.com/arch-go/arch-go/commit/bac7ba22ed3bdf82781f13aaeb1ab8be24672f6f))

## [1.6.1](https://github.com/arch-go/arch-go/compare/v1.6.0...v1.6.1) (2024-09-18)


### Bug Fixes

* SBOM generation for release ([#108](https://github.com/arch-go/arch-go/issues/108)) ([d10df65](https://github.com/arch-go/arch-go/commit/d10df65e0a650917f2689675e96239f59efdcb40))

## [1.6.0](https://github.com/arch-go/arch-go/compare/v1.5.4...v1.6.0) (2024-09-18)


### ⚠ BREAKING CHANGES

* Change module from `fdaines/arch-go` to `arch-go/arch-go` ([#80](https://github.com/arch-go/arch-go/issues/80))

### Features

* generation of a json report ([#99](https://github.com/arch-go/arch-go/issues/99)) ([a0dd5db](https://github.com/arch-go/arch-go/commit/a0dd5dba91d54a3834a109db9b129641e28503a2))


### Code Refactorings

* Change module from `fdaines/arch-go` to `arch-go/arch-go` ([#80](https://github.com/arch-go/arch-go/issues/80)) ([db38838](https://github.com/arch-go/arch-go/commit/db38838ba17c2d0ba104f1bf413f4e53278267e3))


### Dependencies

* update go version to 1.22.5 ([#55](https://github.com/arch-go/arch-go/issues/55)) ([264e29d](https://github.com/arch-go/arch-go/commit/264e29de25945713b048354fe44bb0485310e99e))
* update module github.com/agiledragon/gomonkey/v2 to v2.12.0 ([#45](https://github.com/arch-go/arch-go/issues/45)) ([e4d027f](https://github.com/arch-go/arch-go/commit/e4d027fde751995e4e230e6afb00c629864d4ad3))
* update module github.com/fatih/color to v1.17.0 ([#46](https://github.com/arch-go/arch-go/issues/46)) ([14f91f8](https://github.com/arch-go/arch-go/commit/14f91f872d93c39c4fecbc4ab16bbf6fdd990aed))
* update module github.com/jedib0t/go-pretty/v6 to v6.5.9 ([#47](https://github.com/arch-go/arch-go/issues/47)) ([8956c5f](https://github.com/arch-go/arch-go/commit/8956c5fd95cf496a56b1c613746e1df19b24fa94))
* update module github.com/spf13/cobra to v1.8.1 ([#43](https://github.com/arch-go/arch-go/issues/43)) ([700f268](https://github.com/arch-go/arch-go/commit/700f268b66538f4fb332e68634eccc37a400743d))
* update module github.com/spf13/viper to v1.19.0 ([#48](https://github.com/arch-go/arch-go/issues/48)) ([537aa47](https://github.com/arch-go/arch-go/commit/537aa4715e6b60ba412b4d140499c4c49e66c9e1))
* update module golang.org/x/tools to v0.25.0 ([#104](https://github.com/arch-go/arch-go/issues/104)) ([57a386e](https://github.com/arch-go/arch-go/commit/57a386ebd9be2edf8a6f4c6106fb0cb0a6a3a55e))
* update module gopkg.in/yaml.v2 to v3 ([#54](https://github.com/arch-go/arch-go/issues/54)) ([6164350](https://github.com/arch-go/arch-go/commit/6164350a86a891a811a278feb50e26853768ff8a))

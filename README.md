# Go (Golang) GOOS and GOARCH

https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63


All of the following information is based on `go version go1.17.1 darwin/amd64`.

## GOOS Values
 _______________________________
| GOOS        | Out of the Box  |
|:----------: | :------------:  |
| `linux`     | ✅              |
| `darwin`    | ✅              |
| `windows`   | ✅              |
 _______________________________ 

All GOOS values:
```text
"aix", "android", "darwin", "dragonfly", "freebsd", "hurd", "illumos", "ios", "js", "linux", "nacl", "netbsd", "openbsd", "plan9", "solaris", "windows", "zos"
```

"Out of the box" GOOS values:
```text
"aix", "android", "darwin", "dragonfly", "freebsd", "illumos", "ios", "js", "linux", "netbsd", "openbsd", "plan9", "solaris", "windows"
```

> NOTE
>
> "Out of the box" means the GOOS is supported out of the box, i.e. the stocked `go` command can build the source code without the help of a C compiler, etc.

> NOTE
>
> The full list is based on https://github.com/golang/go/blob/master/src/go/build/syslist.go. The "out of the box" information is based on the result of [2-make1.sh](#file-2-make1-sh) below.

## GOARCH Values
 ___________________________________________________
| GOARCH        | Out of the Box | 32-bit | 64-bit  |
| :------------ | :------------: | :----: | :----:  |
| `386`         | ✅             | ✅     |         |
| `amd64`       | ✅             |        | ✅      |
 ___________________________________________________

All GOARCH values:
```text
"386", "amd64", "amd64p32", "arm", "arm64", "arm64be", "armbe", "loong64", "mips", "mips64", "mips64le", "mips64p32", "mips64p32le", "mipsle", "ppc", "ppc64", "ppc64le", "riscv", "riscv64", "s390", "s390x", "sparc", "sparc64", "wasm"
```

All 32-bit GOARCH values:
```text
"386", "amd64p32", "arm", "armbe", "mips", "mips64p32", "mips64p32le", "mipsle", "ppc", "riscv", "s390", "sparc"
```

All 64-bit GOARCH values:
```text
"amd64", "arm64", "arm64be", "loong64", "mips64", "mips64le", "ppc64", "ppc64le", "riscv64", "s390x", "sparc64", "wasm"
```

"Out of the box" GOARCH values:
```text
"386", "amd64", "arm", "arm64", "mips", "mips64", "mips64le", "mipsle", "ppc64", "ppc64le", "riscv64", "s390x", "wasm"
```

"Out of the box" 32-bit GOARCH values:
```text
"386", "arm", "mips", "mipsle"
```

"Out of the box" 64-bit GOARCH values:
```text
"amd64", "arm64", "mips64", "mips64le", "ppc64", "ppc64le", "riscv64", "s390x", "wasm"
```

> NOTE
>
> "Out of the box" means the GOARCH is supported out of the box, i.e. the stocked `go` command can build the source code without the help of a C compiler, etc.

> NOTE
>
> The full list is based on https://github.com/golang/go/blob/master/src/go/build/syslist.go. The "out of the box" information is based on the result of [2-make1.sh](#file-2-make1-sh) below. The "32-bit/64-bit" information is based on the result of [4-make2.sh](#file-4-make2-sh) below and https://golang.org/doc/install/source.

## Platform Values

| Platform          | Out of the Box | 32-bit | 64-bit |
| :---------------- | :------------: | :----: | :----: |
| `aix/ppc64`       | ✅              |        | ✅      |
| `android/386`     |                | ✅      |        |
| `android/amd64`   |                |        | ✅      |
| `android/arm`     |                | ✅      |        |
| `android/arm64`   | ✅              |        | ✅      |
| `darwin/amd64`    | ✅              |        | ✅      |
| `darwin/arm64`    | ✅              |        | ✅      |
| `dragonfly/amd64` | ✅              |        | ✅      |
| `freebsd/386`     | ✅              | ✅      |        |
| `freebsd/amd64`   | ✅              |        | ✅      |
| `freebsd/arm`     | ✅              | ✅      |        |
| `freebsd/arm64`   | ✅              |        | ✅      |
| `illumos/amd64`   | ✅              |        | ✅      |
| `ios/amd64`       | ✅              |        | ✅      |
| `ios/arm64`       |                |        | ✅      |
| `js/wasm`         | ✅              |        | ✅      |
| `linux/386`       | ✅              | ✅      |        |
| `linux/amd64`     | ✅              |        | ✅      |
| `linux/arm`       | ✅              | ✅      |        |
| `linux/arm64`     | ✅              |        | ✅      |
| `linux/mips`      | ✅              | ✅      |        |
| `linux/mips64`    | ✅              |        | ✅      |
| `linux/mips64le`  | ✅              |        | ✅      |
| `linux/mipsle`    | ✅              | ✅      |        |
| `linux/ppc64`     | ✅              |        | ✅      |
| `linux/ppc64le`   | ✅              |        | ✅      |
| `linux/riscv64`   | ✅              |        | ✅      |
| `linux/s390x`     | ✅              |        | ✅      |
| `netbsd/386`      | ✅              | ✅      |        |
| `netbsd/amd64`    | ✅              |        | ✅      |
| `netbsd/arm`      | ✅              | ✅      |        |
| `netbsd/arm64`    | ✅              |        | ✅      |
| `openbsd/386`     | ✅              | ✅      |        |
| `openbsd/amd64`   | ✅              |        | ✅      |
| `openbsd/arm`     | ✅              | ✅      |        |
| `openbsd/arm64`   | ✅              |        | ✅      |
| `openbsd/mips64`  | ✅              |        | ✅      |
| `plan9/386`       | ✅              | ✅      |        |
| `plan9/amd64`     | ✅              |        | ✅      |
| `plan9/arm`       | ✅              | ✅      |        |
| `solaris/amd64`   | ✅              |        | ✅      |
| `windows/386`     | ✅              | ✅      |        |
| `windows/amd64`   | ✅              |        | ✅      |
| `windows/arm`     | ✅              | ✅      |        |
| `windows/arm64`   | ✅              |        | ✅      |

All Platform values:
```text
"aix/ppc64", "android/386", "android/amd64", "android/arm", "android/arm64", "darwin/amd64", "darwin/arm64", "dragonfly/amd64", "freebsd/386", "freebsd/amd64", "freebsd/arm", "freebsd/arm64", "illumos/amd64", "ios/amd64", "ios/arm64", "js/wasm", "linux/386", "linux/amd64", "linux/arm", "linux/arm64", "linux/mips", "linux/mips64", "linux/mips64le", "linux/mipsle", "linux/ppc64", "linux/ppc64le", "linux/riscv64", "linux/s390x", "netbsd/386", "netbsd/amd64", "netbsd/arm", "netbsd/arm64", "openbsd/386", "openbsd/amd64", "openbsd/arm", "openbsd/arm64", "openbsd/mips64", "plan9/386", "plan9/amd64", "plan9/arm", "solaris/amd64", "windows/386", "windows/amd64", "windows/arm", "windows/arm64"
```

All 32-bit Platform values:
```text
"android/386", "android/arm", "freebsd/386", "freebsd/arm", "linux/386", "linux/arm", "linux/mips", "linux/mipsle", "netbsd/386", "netbsd/arm", "openbsd/386", "openbsd/arm", "plan9/386", "plan9/arm", "windows/386", "windows/arm"
```

All 64-bit Platform values:
```text
"aix/ppc64", "android/amd64", "android/arm64", "darwin/amd64", "darwin/arm64", "dragonfly/amd64", "freebsd/amd64", "freebsd/arm64", "illumos/amd64", "ios/amd64", "ios/arm64", "js/wasm", "linux/amd64", "linux/arm64", "linux/mips64", "linux/mips64le", "linux/ppc64", "linux/ppc64le", "linux/riscv64", "linux/s390x", "netbsd/amd64", "netbsd/arm64", "openbsd/amd64", "openbsd/arm64", "openbsd/mips64", "plan9/amd64", "solaris/amd64", "windows/amd64", "windows/arm64"
```

"Out of the box" Platform values:
```text
"aix/ppc64", "android/arm64", "darwin/amd64", "darwin/arm64", "dragonfly/amd64", "freebsd/386", "freebsd/amd64", "freebsd/arm", "freebsd/arm64", "illumos/amd64", "ios/amd64", "js/wasm", "linux/386", "linux/amd64", "linux/arm", "linux/arm64", "linux/mips", "linux/mips64", "linux/mips64le", "linux/mipsle", "linux/ppc64", "linux/ppc64le", "linux/riscv64", "linux/s390x", "netbsd/386", "netbsd/amd64", "netbsd/arm", "netbsd/arm64", "openbsd/386", "openbsd/amd64", "openbsd/arm", "openbsd/arm64", "openbsd/mips64", "plan9/386", "plan9/amd64", "plan9/arm", "solaris/amd64", "windows/386", "windows/amd64", "windows/arm", "windows/arm64"
```

"Out of the box" 32-bit Platform values:
```text
"freebsd/386", "freebsd/arm", "linux/386", "linux/arm", "linux/mips", "linux/mipsle", "netbsd/386", "netbsd/arm", "openbsd/386", "openbsd/arm", "plan9/386", "plan9/arm", "windows/386", "windows/arm"
```

"Out of the box" 64-bit Platform values:
```text
"aix/ppc64", "android/arm64", "darwin/amd64", "darwin/arm64", "dragonfly/amd64", "freebsd/amd64", "freebsd/arm64", "illumos/amd64", "ios/amd64", "js/wasm", "linux/amd64", "linux/arm64", "linux/mips64", "linux/mips64le", "linux/ppc64", "linux/ppc64le", "linux/riscv64", "linux/s390x", "netbsd/amd64", "netbsd/arm64", "openbsd/amd64", "openbsd/arm64", "openbsd/mips64", "plan9/amd64", "solaris/amd64", "windows/amd64", "windows/arm64"
```

> NOTE
>
> "Out of the box" means the platform is supported out of the box, i.e. the stocked `go` command can build the source code without the help of a C compiler, etc.

> NOTE
>
> The full list is based on the result of the command `go tool dist list`. The "out of the box" information is based on the result of [2-make1.sh](#file-2-make1-sh) below. The "32-bit/64-bit" information is based on the result of [4-make2.sh](#file-4-make2-sh) below and https://golang.org/doc/install/source.

## Support Grid 1

|                   | `android` | `darwin` | `ios` | `js` | `linux` | `windows` |                   |
| ----------------: | :-------: | :------: | :---: | :--: | :-----: | :-------: | :---------------- |
| **`386`**         | ☑️ (α1)    |          |       |      | ✅       | ✅         | **`386`**         |
| **`amd64`**       | ☑️ (α2)    | ✅        | ✅     |      | ✅       | ✅         | **`amd64`**       |
| **`amd64p32`**    |           |          |       |      |         |           | **`amd64p32`**    |
| **`arm`**         | ☑️ (α2)    |          |       |      | ✅       | ✅         | **`arm`**         |
| **`arm64`**       | ✅         | ✅        | ☑️ (β) |      | ✅       | ✅         | **`arm64`**       |
| **`arm64be`**     |           |          |       |      |         |           | **`arm64be`**     |
| **`armbe`**       |           |          |       |      |         |           | **`armbe`**       |
| **`loong64`**     |           |          |       |      |         |           | **`loong64`**     |
| **`mips`**        |           |          |       |      | ✅       |           | **`mips`**        |
| **`mips64`**      |           |          |       |      | ✅       |           | **`mips64`**      |
| **`mips64le`**    |           |          |       |      | ✅       |           | **`mips64le`**    |
| **`mips64p32`**   |           |          |       |      |         |           | **`mips64p32`**   |
| **`mips64p32le`** |           |          |       |      |         |           | **`mips64p32le`** |
| **`mipsle`**      |           |          |       |      | ✅       |           | **`mipsle`**      |
| **`ppc`**         |           |          |       |      |         |           | **`ppc`**         |
| **`ppc64`**       |           |          |       |      | ✅       |           | **`ppc64`**       |
| **`ppc64le`**     |           |          |       |      | ✅       |           | **`ppc64le`**     |
| **`riscv`**       |           |          |       |      |         |           | **`riscv`**       |
| **`riscv64`**     |           |          |       |      | ✅       |           | **`riscv64`**     |
| **`s390`**        |           |          |       |      |         |           | **`s390`**        |
| **`s390x`**       |           |          |       |      | ✅       |           | **`s390x`**       |
| **`sparc`**       |           |          |       |      |         |           | **`sparc`**       |
| **`sparc64`**     |           |          |       |      | (γ)     |           | **`sparc64`**     |
| **`wasm`**        |           |          |       | ✅    |         |           | **`wasm`**        |
|  | **`android`** | **`darwin`** | **`ios`** | **`js`** | **`linux`** | **`windows`** |  |

## Support Grid 2

|                   | `a` | `d` | `f` | `h` | `i` | `m` | `n` | `o` | `p` | `s` | `z` |                   |
| ----------------: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :-: | :---------------- |
| **`386`**         |     |     | ✅   |     |     |     | ✅   | ✅   | ✅   |     |     | **`386`**         |
| **`amd64`**       |     | ✅   | ✅   |     | ✅   |     | ✅   | ✅   | ✅   | ✅   |     | **`amd64`**       |
| **`amd64p32`**    |     |     |     |     |     |     |     |     |     |     |     | **`amd64p32`**    |
| **`arm`**         |     |     | ✅   |     |     |     | ✅   | ✅   | ✅   |     |     | **`arm`**         |
| **`arm64`**       |     |     | ✅   |     |     |     | ✅   | ✅   |     |     |     | **`arm64`**       |
| **`arm64be`**     |     |     |     |     |     |     |     |     |     |     |     | **`arm64be`**     |
| **`armbe`**       |     |     |     |     |     |     |     |     |     |     |     | **`armbe`**       |
| **`loong64`**     |     |     |     |     |     |     |     |     |     |     |     | **`loong64`**     |
| **`mips`**        |     |     |     |     |     |     |     |     |     |     |     | **`mips`**        |
| **`mips64`**      |     |     |     |     |     |     |     | ✅   |     |     |     | **`mips64`**      |
| **`mips64le`**    |     |     |     |     |     |     |     |     |     |     |     | **`mips64le`**    |
| **`mips64p32`**   |     |     |     |     |     |     |     |     |     |     |     | **`mips64p32`**   |
| **`mips64p32le`** |     |     |     |     |     |     |     |     |     |     |     | **`mips64p32le`** |
| **`mipsle`**      |     |     |     |     |     |     |     |     |     |     |     | **`mipsle`**      |
| **`ppc`**         |     |     |     |     |     |     |     |     |     |     |     | **`ppc`**         |
| **`ppc64`**       | ✅   |     |     |     |     |     |     |     |     |     |     | **`ppc64`**       |
| **`ppc64le`**     |     |     |     |     |     |     |     |     |     |     |     | **`ppc64le`**     |
| **`riscv`**       |     |     |     |     |     |     |     |     |     |     |     | **`riscv`**       |
| **`riscv64`**     |     |     |     |     |     |     |     |     |     |     |     | **`riscv64`**     |
| **`s390`**        |     |     |     |     |     |     |     |     |     |     |     | **`s390`**        |
| **`s390x`**       |     |     |     |     |     |     |     |     |     |     |     | **`s390x`**       |
| **`sparc`**       |     |     |     |     |     |     |     |     |     |     |     | **`sparc`**       |
| **`sparc64`**     |     |     |     |     |     |     |     |     |     |     |     | **`sparc64`**     |
| **`wasm`**        |     |     |     |     |     |     |     |     |     |     |     | **`wasm`**        |
|  | **`a`** | **`d`** | **`f`** | **`h`** | **`i`** | **`m`** | **`n`** | **`o`** | **`p`** | **`s`** | **`z`** |  |

### Key

`a` = `aix`, `d` = `dragonfly`, `f` = `freebsd`, `h` = `hurd`, `i` = `illumos`, `m` = `nacl`, `n` = `netbsd`, `o` = `openbsd`, `p` = `plan9`, `s` = `solaris`, `z` = `zos`

✅: Supported (out of the box)

☑️: Supported (with the help of a C compiler, etc.)

(blank): Unsupported

`α1`:
```text
# command-line-arguments
loadinternal: cannot find runtime/cgo
/usr/local/go/pkg/tool/darwin_amd64/link: running clang failed: exit status 1
clang: warning: argument unused during compilation: '-pie' [-Wunused-command-line-argument]
ld: unknown option: -z
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```

`α2`:
```text
# command-line-arguments
loadinternal: cannot find runtime/cgo
/usr/local/go/pkg/tool/darwin_amd64/link: running clang failed: exit status 1
ld: unknown option: -z
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```

`β`:
```text
# command-line-arguments
loadinternal: cannot find runtime/cgo
/usr/local/go/pkg/tool/darwin_amd64/link: running clang failed: exit status 1
ld: warning: ignoring file /var/folders/dd/xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/T/go-link-xxxxxxxxxx/go.o, building for macOS-x86_64 but attempting to link with file built for unknown-arm64
Undefined symbols for architecture x86_64:
  "_main", referenced from:
     implicit entry/start for main executable
ld: symbol(s) not found for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)
```

`γ`:
```text
go tool compile: exit status 2
compile: unknown architecture "sparc64"
```

> NOTE
>
> The `nacl` GOOS was dropped since `go version 1.14`
>
> The `amd64p32` GOARCH, which is related to the `nacl` GOOS, was also dropped since `go version 1.14` (I believe that `mips64p32` and `mips64p32le` are also related, but I could not find any reference)
>
> Reference: https://golang.org/doc/go1.14#nacl

> NOTE
>
> The `darwin/386` port was dropped since `go version 1.15`
>
> Reference: https://golang.org/doc/go1.15#darwin

> NOTE
>
> On before `go version 1.16`:
> - `darwin/amd64` means **macOS**
> - `darwin/arm64` means **iOS**
>
> With the introduction of Apple Silicon (a.k.a. the M1 chip), on `go version 1.16` or later:
> - `darwin/amd64` means **macOS** with Intel CPU
> - `darwin/arm64` updates to mean **macOS** with Apple Silicon CPU
> - `ios/amd64` is the new port for **iOS Simulator** on macOS with Intel CPU
> - `ios/arm64` is the new port for **iOS**
>
> Reference: https://golang.org/doc/go1.16#darwin
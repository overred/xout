# :pencil: XOut Logger
> XOut - Simple Logger Made for Usability.

## :star: Features
- :rainbow: Color output with formatting tags
  (thanks to [@inhere] with [gookit/color]).
- :window: Color output for Windows (thanks to [@mattn] with [go-colorable]).
- :twisted_rightwards_arrows: Multiple output targets with customization for each.
- :gear: Built-in output formatters inspired by [logrus] and [zerolog].
- :books: Built-in loggers presets for different scenarios.

## :package: <span id="installation">Installation</span>
`go get -u github.com/overred/xout`

## :rocket: <span id="usage">Usage</span>

### :checkered_flag: <span id="fast-tour">Fast Tour</span>

#### Default Built-In Logger
The simples way to start with Default logger.
It's as simple as:
```go
xout.Info("Info log level!")
xout.Infof("Formatting Like %s", "fmt.Printf")
xout.Infof("And <fg=cyan>formatting tags</> support!")
```
// TODO:

### :diving_mask: <span id="deep-dive">Deep Dive</span>
// TODO:

## :abacus: <span id="reasons">Yet Another Logger?</span>
// TODO:

# :link: <span id="links">Links</span>

## :building_construction: Powered On

[gookit/color]: https://github.com/gookit/color
> [gookit/color]: 🎨 Terminal color rendering library, support 8/16 colors, 256 colors, RGB color rendering output, support Print/Sprintf methods, compatible with Windows.

[go-colorable]: https://github.com/mattn/go-colorable
> [go-colorable]: Colorable writer for windows.

## :sparkles: Inspired By

[logrus]: https://github.com/sirupsen/logrus
> [logrus]: Structured, pluggable logging for Go.

[zap]: https://github.com/uber-go/zap
> [zap]: Blazing fast, structured, leveled logging in Go.

[zerolog]: https://github.com/rs/zerolog
> [zerolog]: Zero Allocation JSON Logger

## :clap: Thanks To

[@inhere]: https://github.com/inhere
> [@inhere]: PHP and Go, Java developer. @swoft-cloud, Creator of @gookit, @php-toolkit @phppkg

[@mattn]: https://github.com/inhere
> [@mattn]: Long-time Golang user&contributor, Google Dev Expert for Go, and author of many Go tools, Vim plugin author. Windows hacker C#/Java/C/C++

[@sirupsen]: https://github.com/sirupsen
> [@sirupsen]: Scaling infrastructure for you. https://webscale.ca — Previously Principal Infrastructure Engineer @Shopify

[@uber-go]: https://github.com/uber-go
> [@uber-go]: Uber's open source software for Go development

[@rs]: https://github.com/rs
> [@rs]: Director of Engineering at Netflix Co-Founder & ex-CTO of Dailymotion Co-Founder of NextDNS

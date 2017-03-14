# hyperdrive [![Build Status](https://travis-ci.org/hyperdriven/hyperdrive.svg?branch=master)](https://travis-ci.org/hyperdriven/hyperdrive) [![GoDoc](https://godoc.org/github.com/hyperdriven/hyperdrive?status.svg)](https://godoc.org/github.com/hyperdriven/hyperdrive) [![Slack Group](https://slack.hyperdriven.net/badge.svg)](https://slack.hyperdriven.net)

An opinionated micro-framework for creating Hypermedia APIs in Go.

---

## Install

    go get github.com/hyperdriven/hyperdrive

OR

    glide get github.com/hyperdriven/hyperdrive

## Config

Configuration of core features are done via Environment Variables, in accordence with [12 factor](https://12factor.net/config) principles.

  - `HYPERDRIVE_ENV`: (default: `development`, type: `string`) The stage in your deployment pipeline the api is currently running in. A value of `production` changes behviour for some features (such as whether stack traces are logged during panic recovery). Other values, such as `staging`, can be used but currently have no meanin in the framework other than the one you give it in your own code.
  - `PORT`: (default: `5000`, type: `int`) The port the server should listen on.
  - `GZIP_LEVEL`: (default: `-1`, type: `int`) Accepts a value between `-2` and `9`. Invalid values will be silently discarded and the default of `-1` will be used. More info on compression levels can be found in the docs, but corresponds to `zlib` compression levels.
  - `CORS_ENABLED`: (default: `true`, type: `bool`) Set this to `false` to disable CORS support.
  - `CORS_HEADERS`: (type: `string`) A comma-seperated list of headers to allow and expose during CORS requests. These will be appended to the default set of headers that are always allowed: `Accept`, `Accept-Language`, `Content-Language`, and `Content-Type`.
  - `CORS_ORIGINS`: (default: `*`, type: `string`) A comma-seperated list of origins to allow during CORS requests. These will replace the default value, which is to allow all origins.
  - `CORS_CREDENTIALS`: (default: `true`, type: `bool`) Set this to `false` to disable authenticated CORS requests.

## Docs

  - [GoDoc](https://godoc.org/github.com/hyperdriven/hyperdrive)
  - [Wiki](https://github.com/hyperdriven/hyperdrive/wiki)
  - [Examples](https://github.com/hyperdriven/hyperdrive-examples)

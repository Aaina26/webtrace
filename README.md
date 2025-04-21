# webtrace

## Setup

1) Install go 1.23+ as a prerequisite.
2) Clone this repo in local.
3) Run `go install`. After that run `webtrace` in CLI:
   a) If it shows webtrace CLI application description then follow Usage 1
   b) If it shows command does not exist then follow Usage 2

## Usage 1

1) To print sitemap for a website run:
`webtrace ping <url>`
2) To print and additionally export the sitemap as json file run:
`webtrace ping <url> --json`

## Usage 2

1) To print sitemap for a website run:
`go run main.go ping <url>`
2) To print and additionally export the sitemap as json file run:
`go run main.go ping <url> --json`

## Try with an example

As an example lets take the following url: `https://progenycoffee.com/`
```
> webtrace ping https://progenycoffee.com/       
Status: 200 OK
Response Time: 541.7652ms
Site Map:
- https://progenycoffee.com/
│   ├── pages
│   │   ├── jobs
│   │   ├── locations
│   │   ├── our-story
│   │   ├── press
│   ├── account
│   ├── checkout
│   ├── blogs
│   │   ├── news
│   │   │   ├── what-is-a-coffee-score
│   │   │   ├── how-do-we-guarantee-full-traceability-and-quality
│   │   │   ├── coffee-processing-methods-in-colombia
│   │   │   ├── what-is-pineapple-fermentation
│   │   │   ├── working-with-micro-organisms-to-give-back-to-the-land-what-chemicals-took-away
│   │   │   ├── what-brew-method-is-right-for-me
│   │   │   ├── what-does-it-mean-when-beans-appear-oily-or-wet
│   │   │   ├── what-are-coffee-flavor-notes
│   │   │   ├── why-is-colombian-coffee-so-special
│   │   │   ├── how-does-humidity-impact-coffee
│   │   ├── cafes
│   │   │   ├── the-landing
│   │   │   ├── the-hangar
│   │   ├── jobs
│   │   │   ├── cafe-manager
│   │   │   ├── barista
│   │   │   ├── production-assiociate
│   ├── products
│   │   ├── contento
│   │   ├── refill-program
│   │   ├── felicidad
│   │   ├── armonia
│   │   ├── acaia-pearl
│   │   ├── kalita-wave-185
│   │   ├── belleza
│   │   ├── chemex
│   │   ├── hario-v60-02-filters
│   │   ├── pinita
│   │   ├── chemex-filters
│   │   ├── fellow-ode-grinder
│   │   ├── hario-drip-scale
│   │   ├── raw-edge-t-shirt-minimal-branding
│   │   ├── gratitud
│   │   ├── suave
│   │   ├── hario-v60-drip-decanter
│   │   ├── cotton-tote
│   │   ├── pcf-uno
│   │   ├── hario-v60-server-600ml
│   │   ├── kalita-wave-filters-185
│   │   ├── hario-v60-02
│   │   ├── kalita-server
│   │   ├── baratza-encore
│   │   ├── hario-mini-plus-grinder
│   │   ├── copo
│   │   ├── geisha
│   ├── collections
│   │   ├── coffee-collection
│   │   ├── farmers-reserve
│   │   ├── merch
│   │   ├── brewing-gear
```

`webtrace ping https://progenycoffee.com/ --json` Gives a similar output and additionally saves the sitemap in "sitemap.json" in current directory.


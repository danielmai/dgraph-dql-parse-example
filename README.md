Example:

```go
query := `{
  q(func: uid(0x5)) {
    name
    posts {
      title
    }
  }
}`
g := gql.Request{
	Str: query,
}
result, err := gql.Parse(g)
```

Example pretty-print of `gql.Parse` result:

```sh
$ go run main.go
```

```text
Query:
{
  q(func: uid(0x5)) {
    name
    posts {
      title
    }
  }
}
---
Parsed:
{
  "UID": [
    5
  ],
  "Attr": "",
  "Langs": null,
  "Alias": "q",
  "IsCount": false,
  "IsInternal": false,
  "IsGroupby": false,
  "Var": "",
  "NeedsVar": null,
  "Func": {
    "Attr": "",
    "Lang": "",
    "Name": "uid",
    "Args": null,
    "UID": null,
    "NeedsVar": null,
    "IsCount": false,
    "IsValueVar": false,
    "IsLenVar": false
  },
  "Expand": "",
  "Args": {},
  "Order": null,
  "Children": [
    {
      "UID": null,
      "Attr": "name",
      "Langs": null,
      "Alias": "",
      "IsCount": false,
      "IsInternal": false,
      "IsGroupby": false,
      "Var": "",
      "NeedsVar": null,
      "Func": null,
      "Expand": "",
      "Args": {},
      "Order": null,
      "Children": null,
      "Filter": null,
      "MathExp": null,
      "Normalize": false,
      "Recurse": false,
      "RecurseArgs": {
        "Depth": 0,
        "AllowLoop": false
      },
      "ShortestPathArgs": {
        "From": null,
        "To": null
      },
      "Cascade": false,
      "IgnoreReflex": false,
      "Facets": null,
      "FacetsFilter": null,
      "GroupbyAttrs": null,
      "FacetVar": null,
      "FacetOrder": "",
      "FacetDesc": false,
      "IsEmpty": false
    },
    {
      "UID": null,
      "Attr": "posts",
      "Langs": null,
      "Alias": "",
      "IsCount": false,
      "IsInternal": false,
      "IsGroupby": false,
      "Var": "",
      "NeedsVar": null,
      "Func": null,
      "Expand": "",
      "Args": {},
      "Order": null,
      "Children": [
        {
          "UID": null,
          "Attr": "title",
          "Langs": null,
          "Alias": "",
          "IsCount": false,
          "IsInternal": false,
          "IsGroupby": false,
          "Var": "",
          "NeedsVar": null,
          "Func": null,
          "Expand": "",
          "Args": {},
          "Order": null,
          "Children": null,
          "Filter": null,
          "MathExp": null,
          "Normalize": false,
          "Recurse": false,
          "RecurseArgs": {
            "Depth": 0,
            "AllowLoop": false
          },
          "ShortestPathArgs": {
            "From": null,
            "To": null
          },
          "Cascade": false,
          "IgnoreReflex": false,
          "Facets": null,
          "FacetsFilter": null,
          "GroupbyAttrs": null,
          "FacetVar": null,
          "FacetOrder": "",
          "FacetDesc": false,
          "IsEmpty": false
        }
      ],
      "Filter": null,
      "MathExp": null,
      "Normalize": false,
      "Recurse": false,
      "RecurseArgs": {
        "Depth": 0,
        "AllowLoop": false
      },
      "ShortestPathArgs": {
        "From": null,
        "To": null
      },
      "Cascade": false,
      "IgnoreReflex": false,
      "Facets": null,
      "FacetsFilter": null,
      "GroupbyAttrs": null,
      "FacetVar": null,
      "FacetOrder": "",
      "FacetDesc": false,
      "IsEmpty": false
    }
  ],
  "Filter": null,
  "MathExp": null,
  "Normalize": false,
  "Recurse": false,
  "RecurseArgs": {
    "Depth": 0,
    "AllowLoop": false
  },
  "ShortestPathArgs": {
    "From": null,
    "To": null
  },
  "Cascade": false,
  "IgnoreReflex": false,
  "Facets": null,
  "FacetsFilter": null,
  "GroupbyAttrs": null,
  "FacetVar": null,
  "FacetOrder": "",
  "FacetDesc": false,
  "IsEmpty": false
}
```

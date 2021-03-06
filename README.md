# Dimple, A Simple Dependency Injection Container for Go

## Example Usage
A simple example:

```go
d := dimple.NewDimple()
kirk := NewKirk()
d.SetValue("kirk", kirk)
k := d.Get("kirk").(Kirk)
```

In this simple example Kirk is attached to the container. Every time `Get` is called the same `Kirk` item is retrieved.

A generator example. Sometimes you need to generate a new item each time.

```go
d := dimple.NewDimple()
d.SetGenerator("tribble", func(d *dimple.Dimple) interface{} {
    return NewTribbles()
})
a := d.Get("tribble")
b := d.Get("tribble")
```

In the case of generators, a new item is returned each time.

Generator items in the container have access to the other items in the container. For example,

```go
d := dimple.NewDimple()
kirk := NewKirk()
d.SetValue("kirk", kirk)
d.SetGenerator("tribble", func(d *dimple.Dimple) interface{} {
    return NewTribbles(d.Get("kirk").(Kirk))
})
```

This is useful for many cases. For example, the configuration for a mail system could be placed on the container along with the mail system itself. The mail system could be returned via a generator that retrieves the configuration from the container.

Items in the container can be extended. For example,

```go
d := dimple.NewDimple()
d.SetGenerator("tribble", func(d *dimple.Dimple) interface{} {
    return NewTribbles(d.Get("kirk").(Kirk))
})
d.Extend("tribble", func(o interface{}, d *dimple.Dimple) interface{}) {
    o.ExtraFuzzy()
    return o
})
t := d.Get("tribble")
```

In this case the item returned was generated by the original generator, that items was passed into the extended item, and the response from that was returned via `Get`.

The overall goal of dimple is to provide a simple mechanism to retrieve items from the container while providing flexibility in where those items come from.

## License
Dimple is licensed under the MIT. See the _LICENSE_ file.

## Inspiration
Dimple is inspired by [Pimple](http://pimple.sensiolabs.org/).

# Go Web Frontend Demo

Link to the article:

https://philipptanlak.com/web-frontends-in-go/

## Partial template rendering

In the most recent change, partial template rendering support was added.

This can be useful when working with [HTMX](https://htmx.org/) or similar frontend libraries.

To render a template partial, add the `partial` query parameter (e.g. `?partial=statistics`).


```
http://localhost:8080/dashboard?partial=statistics

<div id="statistics">
    <p>Likes: 69</p>
    <p>Follows: 420</p>
</div>
```

```
http://localhost:8080/dashboard?partial=logs

<div id="logs">
    <div>Log message 1</div>
    <div>Log message 2</div>
    <div>Log message 3</div>
</div>
```

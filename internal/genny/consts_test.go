package genny

const validFullDocStr = `
<h2 id="foo">Foo</h2>
<p>
  Some text
</p>

<h3>GET api/foo/{index}</h3>

<pre><code>{
  "index": "bar",
  "name": "Bar",
  "desc": [
    "- some description,"
    "- some other description,"
  ],
  "url": "/api/foo/bar"
}</code></pre>

<h4>Foo</h4>

<table>
  <thead>
    <tr>
      <th align="left">Name</th>
      <th align="left">Description</th>
      <th align="left">Data Type</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td align="left">index</td>
      <td align="left">index field description</td>
      <td align="left">string</td>
    </tr>
    <tr>
      <td align="left">name</td>
      <td align="left">name field description</td>
      <td align="left">string</td>
    </tr>
  </tbody>
</table>

<h3>GET api/baz/{index}</h3>

<pre><code>{
  "index": "bar",
  "name": "Bar",
  "desc": [
    "- some description,"
    "- some other description,"
  ],
  "url": "/api/foo/bar"
}</code></pre>

<h4>Baz</h4>

<table>
  <thead>
    <tr>
      <th align="left">Name</th>
      <th align="left">Description</th>
      <th align="left">Data Type</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td align="left">index</td>
      <td align="left">index field description</td>
      <td align="left">string</td>
    </tr>
    <tr>
      <td align="left">name</td>
      <td align="left">name field description</td>
      <td align="left">string</td>
    </tr>
  </tbody>
</table>
`

const validEndpointStr = `
<h3>GET api/foo/{index}</h3>

<pre><code>{
  "index": "bar",
  "name": "Bar",
  "desc": [
    "- some description,"
    "- some other description,"
  ],
  "url": "/api/foo/bar"
}</code></pre>

<h4>Foo</h4>

<table>
  <thead>
    <tr>
      <th align="left">Name</th>
      <th align="left">Description</th>
      <th align="left">Data Type</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td align="left">index</td>
      <td align="left">index field description</td>
      <td align="left">string</td>
    </tr>
    <tr>
      <td align="left">name</td>
      <td align="left">name field description</td>
      <td align="left">string</td>
    </tr>
  </tbody>
</table>
`

const validFieldTableStr = `<h4>Foo</h4>

<table>
  <thead>
    <tr>
      <th align="left">Name</th>
      <th align="left">Description</th>
      <th align="left">Data Type</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td align="left">index</td>
      <td align="left">index field description</td>
      <td align="left">string</td>
    </tr>
    <tr>
      <td align="left">name</td>
      <td align="left">name field description</td>
      <td align="left">string</td>
    </tr>
  </tbody>
</table>`

const validTableRow = `<tr>
  <td align="left">index</td>
  <td align="left">index field description</td>
  <td align="left">string</td>
</tr>`

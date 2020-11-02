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

const bugimade = `
GET api/subraces/{index}</h3>

<pre><code>{
  "index": "hill-dwarf",
  "name": "Hill Dwarf",
  "race": {
    "name": "Dwarf",
    "index": "dwarf",
    "url": "/api/races/dwarf",
  },
  "desc": "As a hill dwarf, you have keen senses, deep intuition, and remarkable resilience.",
  "ability_bonuses": [
    {
      "ability_score": {
        "index": "wis",
        "name": "WIS",
        "url": "/api/ability-scores/wis"
      },
      "bonus": 1
    }
  ],
  "starting_proficiencies": [],
  "languages": [],
  "racial_traits": [
    {
      "name": "Dwarven Toughness",
      "index": "dwarven-toughness",
      "url": "/api/traits/stonecunning"
    }
  ],
  "url": "/api/subraces/hill-dwarf"
}</code></pre>

<h4>subrace</h4>

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
      <td align="left">ability_bonuses</td>
      <td align="left">Additional ability bonuses granted by this sub race.</td>
      <td align="left">list (<a href="#ability-bonus">AbilityBonus</a>)</td>
    </tr>
    <tr>
      <td align="left">starting_proficiencies</td>
      <td align="left">Starting proficiencies for all new characters of this subrace.</td>
      <td align="left">
        list <a href="#apireference">APIReference</a> (<a href="#proficiencies">Proficiencies</a>)
      </td>
    </tr>
  </tbody>
</table>
`

package server

import (
	"fmt"
	"io"
	"math/rand"
	"strings"
)

func writeSampleHTML(w io.Writer, rndSeed int64) {
	htmlParts := append([]string{}, sampleHTMLParts...)
	rand.Shuffle(len(htmlParts), func(i, j int) {
		htmlParts[i], htmlParts[j] = htmlParts[j], htmlParts[i]
	})

	html := strings.Join(htmlParts, "")

	w.Write([]byte(`<!DOCTYPE html>
<html lang="en" dir="ltr">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1.0">
<title></title>
<link rel="stylesheet" href="/stylesheet.css?seed=` + fmt.Sprint(rndSeed) + `">
<style>
html {
	padding: 0;
	margin: 0;
}

body {
	padding: 25px;
	margin: 0;
	width: 768px;
	overflow-x: hidden;
}
</style>
</head>
<body>
`))

	w.Write([]byte(html))

	w.Write([]byte(`
</body>
</html>
`))
}

var sampleHTMLParts = []string{
	`<h1>This kitchensink contains all basic html tags used throughout the site.</h1>
`,
	`<ul>
	<li>This is a list item</li>
	<li>So is this – there could be more</li>
	<li>Make sure to style list items to:
		<ul>
			<li>Not forgetting child list items</li>
			<li>Not forgetting child list items</li>
			<li>Not forgetting child list items</li>
			<li>Not forgetting child list items</li>
		</ul>
	</li>
	<li>A couple more</li>
	<li>top level list items</li>
</ul>
`,
	`<hr>
`,
	`<p>Don’t forget&nbsp;<strong>Ordered lists</strong>:</p>
`,
	`<ol>
	<li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
	<li>Aliquam tincidunt mauris eu risus.
		<ol>
			<li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
			<li>Aliquam tincidunt mauris eu risus.</li>
		</ol>
	</li>
	<li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
	<li>Aliquam tincidunt mauris eu risus.</li>
</ol>
`,
	`<h2>A sub heading which is not as important as the first, but is quite important overall</h2>
`,
	`<p>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.</p>
`,
	`<table>
	<thead>
		<tr>
			<th>Table Heading</th>
			<th>Table Heading</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>Maandag</td>
			<td>table data</td>
		</tr>
		<tr>
			<td>Dinsdag</td>
			<td>table data</td>
		</tr>
		<tr>
			<td>Woensdag</td>
			<td>table data</td>
		</tr>
		<tr>
			<td>Donderdag</td>
			<td>table data</td>
		</tr>
	</tbody>
	<tfoot>
		<tr>
			<td>Maandag</td>
			<td>table data</td>
		</tr>
	</tfoot>
</table>
`,
	`<h3>A sub heading which is not as important as the second, but should be used with consideration</h3>
`,
	`<p>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.&nbsp;<a href="#kitchensink">Vestibulum tortor</a>&nbsp;quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.</p>
`,
	`<blockquote>
	<p>“Ooh – a blockquote! Lorem ipsum dolor sit amet, consectetur adipiscing elit.”</p>
	<p><cite>Vivamus magna.</cite></p>
</blockquote>
`,
	`<h4>A sub heading which is not as important as the second, but should be used with consideration</h4>
`,
	`<p>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.</p>
`,
	`<pre><code>#header h1 a {
    display: block;
    width: 300px;
    height: 80px;
}</code></pre>
`,
	`<hr>
`,
	`<h5>A sub heading which is not as important as the second, but should be used with consideration</h5>
`,
	`<p>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.Definition listConsectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.Lorem ipsum dolor sit ametConsectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</p>
`,
	`<hr>
`,
	`<h6>This heading plays a relatively small bit part role, if you use it at all</h6>
`,
	`<p>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.</p>
`,
	`<p><strong>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.</strong></p>
`,
	`<p><em>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.</em></p>
`,
	`<p><s>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.</s></p>
`,
	`<p><i>Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.</i> Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.</p>
`,
}

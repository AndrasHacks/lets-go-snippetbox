{{define "base"}}
<html lang='en'>
<head>
    <meta charset='utf-8'>
    <title>{{template "title" .}} - Snippetbox</title> </head>
<body>
<header>
    <h1><a href='/'>Snippetbox</a></h1>
</header>
<nav>
    <a href='/'>Home</a>
</nav>
<section>
    {{template "body" .}}
</section>
{{template "footer" .}}
</body>
</html>
{{end}}
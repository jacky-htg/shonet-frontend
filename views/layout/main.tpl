<html>
  <head>
    <title>{{ template "title" . }}</title>
	<meta name="description" value="{{.Description}}"/>
  </head>
  <body>
    {{ template "content" . }}
  </body>
</html>

  <!DOCTYPE html>
    <html>
    <head>
         <title>GB blog</title>
        {{template "head.tpl"}}
    </head>
    <body>
<div class="wrapper">
        <div class="content">
    {{template "header.tpl"}}

{{range .Posts}}
<div class="contentContain">
    <div class="writeBlock">
        <a class="editLnk" href="/edit?blog_title={{.Title}}">Edit</a>
    </div>

    <div class="writeBlock">
        <a class="editLnk" href="/post?blog_title={{.Title}}">Open</a>
    </div>

    <div class="writeBlock">
        <div class="titleClass bg-secondary text-light">
            {{.Title}} {{.Date}}
        </div>
    </div>

    <div class="writeBlock">
        <div class="contentClass bg-secondary text-light">
            {{.Content}}
        </div>
    </div>
</div>
{{end}} 
</div>
<!-- .content -->
{{template "footer.tpl"}}
</div>
<!-- .wrapper -->
</body>
</html>

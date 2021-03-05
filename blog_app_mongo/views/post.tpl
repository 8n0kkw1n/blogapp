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
<div class="writeBlock">
    <div class="titleClass bg-secondary text-light">
        {{.Post.Title}} {{.Post.Date}}
    </div>
</div>

<div class="writeBlock">
    <div class="contentClass bg-secondary text-light">
        {{.Post.Content}}
    </div>
</div>

</div>
<!-- .content -->
{{template "footer.tpl"}}
</div>
<!-- .wrapper -->
</body>
</html>
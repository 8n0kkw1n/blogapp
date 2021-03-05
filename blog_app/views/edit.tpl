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

<div class="writeContain">
   <div class="writeBlock">
        <form method="POST" action="/delete/{{.Post.ID}}">
            <button type="submit" class="btn btn-danger">Delete</button>
        </form>

    <div class="writeBlock">
        <form role="form" method="POST" action="/edit/{{.Post.ID}}">
            <input type="hidden" name="id" value="{{.Post.ID}}" />
            <div class="form-group">
                <label class="col-form-label-lg">Заголовок</label>
                <input type="text" class="form-control form-control-lg" id="title" name="title" value="{{.Post.Title}}" />
            </div>
            <div class="form-group">
                <label class="col-form-label-lg">Контент</label>
                <textarea class="form-control form-control-lg" id="content" name="content">{{.Post.Content}}</textarea>
            </div>
            <button type="submit" class="btn btn-dark btn-lg">Submit</button>
        </form>
    </div>

    <div class="writeBlock">
        <label class="col-form-label-lg" id="mdHtml">HTML</label>

    </div>
</div>
</div>
<!-- .content -->
{{template "footer.tpl"}}
</div>
<!-- .wrapper -->
</body>
</html>

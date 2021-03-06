<!DOCTYPE html>
<html>

<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <title>Mustang | Log in</title>
    <link rel="icon" href="/static/dist/img/favicon.ico">
    {{template "base/link_css.html" .}}
    {{template "base/header_script.html" .}}
</head>

<body class="hold-transition login-page">
<div class="login-box">
    <div class="login-logo">
        <a href="/user/login"><b>Mustang</b></a>
    </div>
    <!-- /.login-logo -->
    <div class="login-box-body">
        <p class="login-box-msg"> 请登录后继续 </p>
        <form id="loginForm" method="post">
            <div class="form-group has-feedback">
                <input type="text" name="username" required class="form-control" placeholder="Username">
                <span class="glyphicon glyphicon-user form-control-feedback"></span>
            </div>
            <div class="form-group has-feedback">
                <input type="password" name="password" required class="form-control" placeholder="Password">
                <span class="glyphicon glyphicon-lock form-control-feedback"></span>
            </div>

            <div class="row">
                <div class="col-xs-12">
                    <button id="btnLogin" type="submit" class="btn btn-lg btn-primary btn-block">Sign In</button>
                    <p class="text-center"><i id='iconWait'></i></p>
                </div>
                <!-- /.col -->
            </div>
        </form>
    </div>
    <!-- /.login-box-body -->
</div>
<!-- /.login-box -->
<script>
    $(function () {
        $('#loginForm').on('submit', function (e) {
            Net.post({
                url: "/login",
                data: JSON.stringify($('#loginForm').serializeJSON()),
                btn: $('#btnLogin'),
                icon: $('#iconWait'),
                success: function (resp) {
                    switch (resp.status) {
                        case 0:
                            ui.msg.success(resp.msg);
                            setTimeout(function () {
                                document.location.href = getQueryString("next")
                            }, 800);
                            $('#iconWait').removeClass("fa fa-spinner fa-spin")
                            break;
                        case 1:
                            ui.msg.fail(resp.msg);
                            $('#iconWait').removeClass("fa fa-spinner fa-spin")
                            break;
                        case -1:
                            ui.msg.error(resp.msg);
                            $('#iconWait').removeClass("fa fa-spinner fa-spin")
                    }
                }
            })
            return false
        })
    })
</script>
</body>

</html>

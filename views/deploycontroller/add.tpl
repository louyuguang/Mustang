{{template "base/base.html" .}}
{{define "content"}}
<div class="box">
    <div class="box-header with-border">
        <h3 class="box-title">创建部署</h3>
    </div>
    <!-- /.box-header -->
    <!-- form start -->
    <form class="form-horizontal" id="deployForm" action="" method="post">
        <div class="box-body">
            <div class="form-group">
                <label for="inputDeployName" class="col-sm-2 control-label">项目<span class="text-red">*</span></label>
                <div class="col-sm-8">
                    <input type="text" name="projectName" value="" class="form-control"
                           id="name" placeholder="project_name">
                </div>
            </div>
            <div class="form-group">
                <label for="inputGit" class="col-sm-2 control-label">Git地址<span class="text-red">*</span></label>
                <div class="col-sm-8">
                    <input type="text" name="gitUrl" value="" class="form-control"
                           id="git_url" placeholder="git@git.example.com:group/project.git">
                </div>
            </div>
            <div class="form-group">
                <label for="inputPort" class="col-sm-2 control-label">开放端口<span class="text-red">*</span></label>
                <div class="col-sm-8">
                    <input type="text" name="port" value="" class="form-control"
                           id="port" placeholder="8080">
                </div>
            </div>
        </div>

        <!-- /.box-body -->
        <div class="box-footer text-center">
            <button id="deployAddBtn" type="button" class="btn btn-info">提交</button>
            <p class="text-center"><i id='iconWait'></i></p>
        </div>
        <!-- /.box-footer -->
    </form>
</div>
{{end}}
{{define "self_footer_script"}}
<script>
    $(function () {
        $('#deployAddBtn').on('click', function (e) {
            var form = $('#deployForm').serializeJSON()
            Net.post({
                url: location.pathname,
                data: JSON.stringify(form),
                btn: $('#deployAddBtn'),
                icon: $('#iconWait'),
                go: getReferrer("{{urlfor "DeployController.List"}}")
        });
        })
    })
</script>
{{end}}
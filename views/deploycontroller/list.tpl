{{template "base/base.html" .}}
{{define "content"}}
<div class="box">
    <div class="box-header">
        <h4 class="box-title">部署列表</h4>
        <div class="box-tools">
            <form class="form-inline" id="search-form">
                <div class="input-group" style="width: 300px;">
                    <a href="{{ urlfor "DeployController.Add" }}" class="btn btn-primary btn-sm pull-right">新建部署</a>
                </div>
            </form>
        </div>
    </div>
    <!-- /.box-header -->
    <div class="box-body">
        <div class="row">
            <div class="col-sm-12">
                <table id="" class="table table-bordered table-hover dataTable" role="grid">
                    <thead>
                    <tr role="row">
                        <th class="text-center">ID</th>
                        <th class="text-center">用户名</th>
                        <th class="text-center">项目</th>
                        <th class="text-center">时间</th>
                        <th class="text-center">操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{ range .Deploys }}
                    <tr role="row" class="odd">
                        <td class="text-center">{{ .Id }}</td>
                        <td class="text-center">{{ .User.UserName }}</td>
                        <td class="text-center">{{ .ProjectName }}</td>
                        <td class="text-center">{{ .Created }}</td>
                        <td class="text-center" style="width:150px">
                            <a class="fa fa-upload execute-tip" title="发布" onclick="deploy_exec(this, {{ .Id }})"
                               href="#"></a>
                            <a class="text-red fa fa-trash-o delete-tip" title="删除"
                               onclick="deploy_delete({{ .Id }})" href="#"></a>
                        </td>
                    </tr>
                    {{end}}
                    </tbody>
                    <tfoot>
                    </tfoot>
                </table>
            </div>
        </div>
        <div class="row">
            <div class="col-sm-6" style="margin-top:5px">
                <div class="dataTables_info" id="editable_info" role="status" aria-live="polite">
                    <p>
                        Showing {{.paginator.Startnum}} to {{.paginator.Endnum}} of {{.paginator.Nums}} entries
                    </p>
                </div>
            </div>
            {{template "base/paginator.html" .}}
        </div>
    </div>
    <!-- /.box-body -->
</div>
{{end}}
{{define "self_footer_script"}}
<script>
    function deploy_exec(button, deploy_id) {
        ui.confirm('确定部署?', function () {
            Net.post({
                url: "{{urlfor "DeployController.Exec"}}",
                data: JSON.stringify({"id": parseInt(deploy_id)}),
                btn: $(button),
                go: "{{urlfor "DeployController.List"}}"
            });
        });
    }
    function deploy_delete(deploy_id) {
        ui.confirm('确定删除部署?', function () {
            Net.post({
                url: "{{urlfor "DeployController.Delete"}}",
                data: JSON.stringify({"id": parseInt(deploy_id)}),
                reload: true
            });
        });
    }
</script>
{{end}}
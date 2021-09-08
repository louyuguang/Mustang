{{template "base/base.html" .}}
{{define "content"}}
<div class="box">
    <div class="box-header">
        <h4 class="box-title">环境列表</h4>
        <div class="box-tools">
            <form class="form-inline" id="search-form">
                <div class="input-group" style="width: 300px;">
                    <input type="text" name="scontent" id="searchContent" class="form-control input-sm pull-right"
                           placeholder="Search" value="{{.Scontent}}">
                    <div class="input-group-btn">
                        <button type="submit" class="btn btn-sm btn-default"><i class="fa fa-search"
                                                                                style="margin-left:0px"></i>
                        </button>
                    </div>
                    <a href="{{ urlfor "EnvClusterBindingController.Add" }}" class="btn btn-primary btn-sm
                    pull-right">创建环境</a>
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
                        <th class="text-center">环境名</th>
                        <th class="text-center">命名空间</th>
                        <th class="text-center">绑定集群</th>
                        <th class="text-center">操作</th>
                    </tr>
                    </thead>
                    <tbody>
                    {{ range .Envs }}
                    <tr role="row" class="odd">
                        <td class="text-center">{{ .Env.Id }}</td>
                        <td class="text-center">{{ .Env.EnvName }}</td>
                        <td class="text-center">{{ .Env.Namespace }}</td>
                        <td class="text-center">{{ range .Clusters }}{{.ClusterName}}{{end}}</td>
                        <td class="text-center" style="width:150px">
                            <a class="fa fa-pencil-square-o audit-tip" title="编辑" onclick=""
                               href="{{ urlfor "EnvClusterBindingController.Update" ":id" .Env.Id }}"></a>
                            <a class="text-red fa fa-trash-o delete-tip" title="删除"
                               onclick="env_delete({{ .Env.Id }})" href="#"></a>
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
    function env_delete(env_id) {
        ui.confirm('确定删除集群?', function () {
            Net.post({
                url: "{{ urlfor "EnvClusterBindingController.Delete" }}",
                data: JSON.stringify({"id": parseInt(env_id)}),
                reload: true
            });
        });
    }
</script>
{{end}}
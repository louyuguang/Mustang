{{template "base/base.html" .}}
{{define "content"}}
    <div class="box">
        <div class="box-header">
            <h4 class="box-title">用户列表</h4>
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
                        <a href="{{ urlfor "UserController.Add" }}" class="btn btn-primary btn-sm pull-right">新建用户</a>
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
                            <th class="text-center">角色</th>
                            <th class="text-center">邮箱</th>
                            <th class="text-center">激活</th>
                            <th class="text-center">操作</th>
                        </tr>
                        </thead>
                        <tbody>
                        {{ range .Users }}
                            <tr role="row" class="odd">
                                <td class="text-center">{{ .Id }}</td>
                                <td class="text-center">{{ .UserName }}</td>
                                <td class="text-center">
                                    {{ .Role.Rolename }}
                                </td>
                                <td class="text-center">{{ .Email }}</td>
                                <td class="text-center">{{ if .Active }}是{{else}}否{{end}}</td>
                                <td class="text-center" style="width:150px">
                                    <a class="fa fa-pencil-square-o audit-tip" title="编辑" onclick="user_edit({{ .Id }})"
                                       href="#"></a>
                                    <a class="text-red fa fa-trash-o delete-tip" title="删除"
                                       onclick="user_delete({{ .Id }})" href="#"></a>
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
        function user_edit(user_id) {
            document.location.href = "{{urlfor "UserController.Add"}}" + '/' + user_id
        }

        function user_delete(user_id) {
            ui.confirm('确定删除用户?', function () {
                Net.post({
                    url: "{{ urlfor "UserController.Delete" }}",
                    data: {"id": user_id},
                    reload: true
                });
            });
        }
    </script>
{{end}}
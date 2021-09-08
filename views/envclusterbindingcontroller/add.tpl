{{template "base/base.html" .}}
{{define "content"}}
<div class="box">
    <div class="box-header with-border">
        <h3 class="box-title">创建环境</h3>
    </div>
    <!-- /.box-header -->
    <!-- form start -->
    <form class="form-horizontal" id="envForm" action="" method="post">
        <div class="box-body">
            <div class="form-group">
                <label class="col-sm-2 control-label">环境名<span class="text-red">*</span></label>
                <div class="col-sm-8">
                    <input type="text" name="envName" value="{{ .EnvAdd.EnvName }}" class="form-control"
                           id="envname" placeholder="production">
                </div>
            </div>
            <div class="form-group">
                <label class="col-sm-2 control-label">命名空间<span class="text-red">*</span></label>
                <div class="col-sm-8">
                    <input type="text" name="namespace" value="{{ .EnvAdd.Namespace }}" class="form-control"
                           id="namespace" placeholder="default">
                </div>
            </div>
            <div class="form-group">
                <label class="col-sm-2 control-label">绑定集群<span class="text-red">*</span></label>
                <div class="col-sm-8">
                    <select id="clusterIds" name="clusterIds" data-placeholder="--请选择集群--" class="js-states form-control" multiple="multiple">
                        <option></option>
                        {{range .Clusters}}
                        <option value={{ .Id }}> {{ .ClusterName }} </option>
                        {{end}}
                    </select>
                </div>
            </div>
        </div>

        <!-- /.box-body -->
        <div class="box-footer text-center">
            <button id="envAddBtn" type="button" class="btn btn-info">提交</button>
            <p class="text-center"><i id='iconWait'></i></p>
        </div>
        <!-- /.box-footer -->
    </form>
</div>
{{end}}
{{define "self_footer_script"}}
<script>
    $(function () {
        $('#clusterIds').select2()
        $('#envAddBtn').on('click', function (e) {
            Net.post({
                url: location.pathname,
                data: $('#envForm').serialize(),
                btn: $('#deployAddBtn'),
                icon: $('#iconWait'),
                go: getReferrer("{{urlfor "EnvClusterBindingController.List"}}")
        });
        })
    })
</script>
{{end}}
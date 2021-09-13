{{template "base/base.html" .}}
{{define "content"}}
    <div class="box">
        <div class="box-header with-border">
            <h3 class="box-title">创建集群</h3>
        </div>
        <!-- /.box-header -->
        <!-- form start -->
        <form class="form-horizontal" id="clusterForm" action="" method="post">
            <div class="box-body">
                <div class="form-group">
                    <label for="inputClustername" class="col-sm-2 control-label">集群名<span class="text-red">*</span></label>
                    <div class="col-sm-8">
                        <input type="text" name="clustername" value="{{ .ClusterAdd.ClusterName }}" class="form-control"
                               id="clustername" placeholder="clustername">
                    </div>
                </div>
                <div class="form-group">
                    <label for="inputAliasname" class="col-sm-2 control-label">集群别名<span class="text-red">*</span></label>
                    <div class="col-sm-8">
                        <input type="text" name="aliasname" value="{{ .ClusterAdd.AliasName }}"
                               class="form-control" id="aliasname" placeholder="aliasname">
                    </div>
                </div>
                <div class="form-group">
                    <label for="inputKubeconfig" class="col-sm-2 control-label">KubeConfig<span class="text-red">*</span></label>
                    <div class="col-sm-8">
                        <textarea rows="10" type="text" class="form-control" name="kubeconfig">{{ .ClusterAdd.KubeConfig }}</textarea>
                    </div>
                </div>
            </div>

            <!-- /.box-body -->
            <div class="box-footer text-center">
                <button id="clusterAddBtn" type="button" class="btn btn-info">提交</button>
                <p class="text-center"><i id='iconWait'></i></p>
            </div>
            <!-- /.box-footer -->
        </form>
    </div>
{{end}}
{{define "self_footer_script"}}
    <script>
        $(function () {
            $('#clusterAddBtn').on('click', function (e) {
                var form = $('#clusterForm').serializeJSON()
                Net.post({
                    url: location.pathname,
                    data: JSON.stringify(form),
                    btn: $('#clusterAddBtn'),
                    icon: $('#iconWait'),
                    go: getReferrer("{{urlfor "ClusterController.List"}}")
                });
            })
        })
    </script>
{{end}}
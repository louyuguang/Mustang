/*
function changeSqlStatusTips(trObject) {
    trObject.each(function(){
        var tdObject = $(this).find('td')
        var sql_id = $(tdObject[0]).text()
        var sql_status = $(tdObject[6]).text()
        var retval = SQLStatusTips(tdObject, sql_status)
        if (!retval) {
            var timer = setInterval(function (){
                Net.get({
                    url: '/csql/api/get_status/?id=' + sql_id,
                    success: function(resp){
                        //console.log("id:" + sql_id + " retval:" + retval);
                        switch (resp.status) {
                            case 0:
                                if (sql_status != resp.data) {
                                    sql_status = resp.data
                                    if (SQLStatusTips(tdObject, sql_status)){
                                        clearInterval(timer)
                                    }
                                }
                                break;
                            default:
                                clearInterval(timer)
                                break
                        }
                    },
                    error: function(resp){
                        clearInterval(timer)
                    }
                })
            }, 5000)
        }
    })
}
*/
function SQLStatusTips(trObject, tdIdx) {
    //console.log(tdObject.html());
    trObject.each(function(){
        var tdArr = $(this).find('td')
        var td_status = $(tdArr[tdIdx])
        switch (parseInt(td_status.text())) {
            case 1:
                td_status.html('<i title="检测成功,等待审核" class="fa fa-edit tipWaitingAudit"></i>')
                $('.tipWaitingAudit').tooltip();
                return 0
                break
            case 2:
                td_status.html('<i title="审核完成,等待执行" class="fa fa-check-square-o tipAuditSuccess"></i>')
                $('.tipAuditSuccess').tooltip();
                return 0
                break
            case 3:
                td_status.html('<i title="任务已在队列中,等待执行" class="fa fa-hourglass-half tipExecuteWaiting"></i>')
                $('.tipExecuteWaiting').tooltip();
                return 0
                break
            case 4:
                td_status.html('<i title="正在执行中" class="fa fa-spinner fa-spin tipExecuting"></i>')
                $('.tipExecuting').tooltip();
                return 0
                break
            case 5:
                td_status.html('<i title="执行完成" class="text-green fa fa-check tipExecuteSuccess"></i>')
                $('.tipExecuteSuccess').tooltip();
                return 1
                break
            case 6:
                td_status.html('<i id="' + status + '" title="执行失败" class="text-red fa fa-close tipExecuteFailed"></i>')
                $('.tipExecuteFailed').tooltip();
                return 1
                break
            default:
                td_status.html('<i id="' + status + '" title="未知错误" class="fa fa-question tipUnknow"></i>')
                $('.tipUnknow').tooltip();
                //console.log("id:" + td_id.text() + " retval:" + status);
                return 1
                break
        }
    })

}

function deployStatusTips(trObject, tdIdx){
    trObject.each(function(){
        var tdArr = $(this).find('td')
        var td_status = $(tdArr[tdIdx])
        switch (parseInt(td_status.text())) {
            case 0:
                td_status.html('<i title="无操作" class="fa fa-info tipNone"></i>')
                $('.tipNone').tooltip();
                break
            case 1:
                td_status.html('<i title="检测成功,等待开发组长审核!" class="fa fa-edit tipWaitingAudit"></i>')
                $('.tipWaitingAudit').tooltip();
                break
            case 21:
                td_status.html('<i title="等待测试负责人审核!" class="fa fa-edit tipWaitingAudit"></i>')
                $('.tipWaitingAudit').tooltip();
                break
            case 22:
                td_status.html('<i title="等待测试组长审核!" class="fa fa-edit tipWaitingAudit"></i>')
                $('.tipWaitingAudit').tooltip();
                break
            case 30:
                td_status.html('<i title="任务已在队列中,等待执行" class="fa fa-hourglass-half tipExecuteWaiting"></i>')
                $('.tipExecuteWaiting').tooltip();
                break
            case 40:
                td_status.html('<i title="代码检出中" class="fa fa-code-fork tipCodeCheckout"></i>')
                $('.tipCodeCheckout').tooltip();
                break
            case 50:
                td_status.html('<i title="docker build编译中" class="fa fa-spinner fa-spin tipCodeCompile"></i>')
                $('.tipCodeCompile').tooltip();
                break
            case 60:
                td_status.html('<i title="正在处理配置依赖任务" class="fa fa-circle-o-notch fa-spin tipDependTask"></i>')
                $('.tipDependTask').tooltip();
                break
            case 70:
                td_status.html('<i title="正在部署应用" class="fa fa-cloud-upload tipSyncServer"></i>')
                $('.tipSyncServer').tooltip();
                break
            case 80:
                td_status.html('<i title="部署成功" class="text-green fa fa-check tipExecuteSuccess"></i>')
                $('.tipExecuteSuccess').tooltip();
                break
            case 90:
                td_status.html('<i title="部署失败" class="text-red fa fa-close tipExecuteFailed"></i>')
                $('.tipExecuteFailed').tooltip();
                break
            case 100:
                td_status.html('<i title="检测成功,等待BOSS审核!" class="fa fa-edit tipWaitingAudit"></i>')
                $('.tipWaitingAudit').tooltip();
                break
            default:
                td_status.html('<i title="未知错误" class="fa fa-question tipUnknow"></i>')
                $('.tipUnknow').tooltip();
                break
            }
    })
}

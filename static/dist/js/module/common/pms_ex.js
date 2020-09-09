var Permissions = [
    {"id": 1, "text": "SQL审核"},
    {"id": 2, "text": "SQL执行"},
    {"id": 3, "text": "项目创建"},
    {"id": 10, "text": "Stg环境发起部署"},
    {"id": 11, "text": "Stg环境审核"},
    {"id": 12, "text": "Stg环境执行部署"},
    {"id": 14, "text": "Stg环境新增配置"},
    {"id": 15, "text": "Stg环境发起回滚"},
    {"id": 20, "text": "Prod环境发起部署"},
    {"id": 21, "text": "Prod环境审核"},
    {"id": 22, "text": "Prod环境执行部署"},
    {"id": 24, "text": "Prod环境新增配置"},
    {"id": 25, "text": "Prod环境发起回滚"}
]

var PMS_EX = function (pms_list) {
    pms_list = Permissions
    var groups = {
        envs: {
            title: '环境权限',
            //match: /(dev|test|stg|prod)/i,
            match: /(stg|prod)/i,
            lists: {
                stg: {},
                prod: {}
            }
        },
        others: {
            title: '其它权限',
            list: []
        }
    }

    $.each(pms_list, function (_, pms) {
        if (groups.envs.match.test(pms.text)) {
            // groups.envs.list.push(pms)
            var parts = pms.text.split('环境')
            groups.envs.lists[parts[0].toLowerCase()][parts[1]] = pms.id
        } else {
            groups.others.list.push(pms)
        }
    })

    function build_checkbox(pms_id, pms_text, cols) {
        return $('<label><input type="checkbox" class="pms-checkbox" value="' + pms_id + '" id="pms-checkbox-' + pms_id + '" /> ' + pms_text + '</label>')
    }
    function build_controls(target) {
        return '<small class="pull-right control" data-target="' + target + '"><a href="javascript:;" class="text-muted" data-action="checkall"><i class="glyphicon glyphicon-ok"></i></a> <a href="javascript:;" class="text-muted" data-action="uncheckall"><i class="glyphicon glyphicon-remove"></i></a></small>'
    }

    // create elements
    var $pms_group = $('<div id="pms-group" class="col-sm-8" />')
    $.each(groups, function (group_key, group_data) {
        var $pms_box = $('<div class="box box-default"><div class="box-header with-border"><h4 class="box-title">' + group_data.title + '</h4></div><div class="box-body"></div></div>').appendTo($pms_group)
        var $pms_box_body = $pms_box.find('.box-body')

        if (group_key == 'envs') {
            var $tbody = $pms_box_body.append('<table class="table table-striped table-bordered table-hover"><thead><tr><th>Stg环境 ' + build_controls('.pms-env-stg') + '</th><th>Prod环境 ' + build_controls('.pms-env-prod') + '</th></tr></thead><tbody></tbody></table>').on('click', '.control a', function () {
                var target = $(this).closest('.control').data('target')
                var action = $(this).data('action')
                var $target_checkbox = $(target).find(':checkbox')

                if (action == 'checkall') {
                    $target_checkbox.prop('checked', true)
                } else if (action == 'uncheckall') {
                    $target_checkbox.prop('checked', false)
                }
            }).find('tbody')

            $.each(group_data.lists.stg, function (pms_text) {
                var $tr = $('<tr />').appendTo($tbody)
                //$('<td />').append(build_checkbox(group_data.lists.dev[pms_text], pms_text).addClass('col-sm-12 pms-env-dev')).appendTo($tr)
                //$('<td />').append(build_checkbox(group_data.lists.test[pms_text], pms_text).addClass('col-sm-12 pms-env-test')).appendTo($tr)
                $('<td />').append(build_checkbox(group_data.lists.stg[pms_text], pms_text).addClass('col-sm-12 pms-env-stg')).appendTo($tr)
                $('<td />').append(build_checkbox(group_data.lists.prod[pms_text], pms_text).addClass('col-sm-12 pms-env-prod')).appendTo($tr)
            })
        } else {
            $pms_box_body = $('<div class="row" />').appendTo($pms_box_body)

            $.each(group_data.list, function (_, pms_data) {
                $pms_box_body.append(build_checkbox(pms_data.id, pms_data.text).addClass('col-sm-3'))
            })
        }
    })

    return {
        $element: $pms_group,
        toList: function () {
            var vals = []
            $pms_group.find('.pms-checkbox:checked').each(function () {
                vals.push($(this).val())
            })

            return vals
        },
        fromList: function (vals) {
            var checkbox_list = $pms_group.find('.pms-checkbox')

            $.each(vals, function (_, id) {
                checkbox_list.filter('#pms-checkbox-' + id).prop('checked', true)
            })

            return undefined
        }
    }
}

//var pms_ex = new PMS_EX(data.data)
//$('#selectBox').hide().after(pms_ex.$element)

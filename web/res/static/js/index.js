layui.use(['form', 'layer', 'table', 'util'], function () {
    const form = layui.form;
    const layer = layui.layer;
    const table = layui.table;
    const util = layui.util;
    const $ = layui.jquery;
    const ctxPath = $("#ctxPath").val();

    if (typeof String.prototype.startWith != 'function') {
        String.prototype.startWith = function (prefix) {
            return this.slice(0, prefix.length) === prefix;
        };
    }

    if (typeof String.prototype.startWith != 'function') {
        String.prototype.startWith = function (suffix) {
            return this.indexOf(suffix, this.length - suffix.length) !== -1;
        };
    }

    function getFullPath(path) {
        var root = ctxPath;
        if (ctxPath.endsWith("/")) {
            root = ctxPath.substring(0, ctxPath.length - 1);
        }
        return root + path;
    }

    form.verify({
        //value：表单的值、item：表单的DOM对象
        pkg: function (value, item) {
            if (!new RegExp("^[a-zA-Z]+([a-zA-Z]*[\\/\\-\\_\\.][a-zA-Z]+)*$").test(value)) {
                return '请输入正确的包名';
            }
        },
    });

    form.on('switch(chkHasServiceInterface)', function (data) {
        if (data.elem.checked) {
            $('#hasServiceInterfaceDiv').show();
        } else {
            $('#hasServiceInterfaceDiv').hide();
        }
    });

    form.on('submit(query)', function (data) {
        table.reload('gen', {
            where: data.field
        });
        return false;
    });

    form.on('submit(connect)', function (data) {
        layer.msg('加载中', {icon: 16, shade: 0.01});
        $.ajax({
            type: "POST",
            url: getFullPath('/connect'),
            contentType: 'application/json',
            data: JSON.stringify(data.field),
            dataType: 'json',
            success: function (res) {
                layer.closeAll();
                if (res.code === "0000") {
                    //重加载表格
                    table.reload('gen');
                } else {
                    layer.msg(res.msg, {icon: 5});
                }
            },
            error: function () {
                layer.closeAll();
                layer.msg("请求错误", {icon: 5});
            }
        });
        return false;
    });

    form.on('submit(commit)', function (data) {
        let request = data.field;
        let checkStatus = table.checkStatus('gen');
        let rows = checkStatus.data;
        if (rows.length === 0) {
            layer.msg("请选择数据表", {icon: 5});
            return false;
        }
        request.tables = rows.map(e => {
            return e.Name;
        });
        let params = JSON.stringify(request);
        //console.log(params);

        layer.msg('提交中', {icon: 16, shade: 0.01});
        $.ajax({
            type: "POST",
            url: getFullPath('/gen'),
            contentType: 'application/json',
            data: params,
            dataType: 'json',
            success: function (res) {
                layer.closeAll();
                if (res.code === "0000") {
                    layer.msg("success");
                    let url = getFullPath("/download?file=" + res.data);
                    $('#downFileIframe').attr('src', url);
                } else {
                    layer.msg(res.msg, {icon: 5, time: 10000});
                }
            },
            error: function () {
                layer.closeAll();
                layer.msg("请求错误", {icon: 5});
            }
        });
        return false;
    });

    table.render({
        elem: '#gen',
        id: 'gen',
        height: 'full-280',
        url: getFullPath('/tables'),
        method: 'get',
        limit: 10000,
        loading: true,
        even: true,
        page: false,
        cols: [[
            {checkbox: true, fixed: 'left'},
            {field: 'Name', title: '表名', width: 200, fixed: 'left', sort: true},
            {
                field: 'CreateTime', title: '创建时间', width: 180, sort: true, templet: function (d) {
                    return util.toDateString(d.CreateTime, "yyyy-MM-dd HH:mm:ss");
                }
            },
            {field: 'Comment', title: '注释'},
        ]]
    });

    $("#commit").on('click', function () {
        $("#btnCommit").click();
    });

});

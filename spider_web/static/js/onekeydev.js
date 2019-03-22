var layer;
var gConfig = ""; //配置文件内容

//显示和关闭加载框
var showLoading = function (show) {
    if (show === true) {
        layer.load(2);
    } else {
        layer.closeAll('loading');
    }
};

var onChooseFile = function (file) {
    $("#txt_config").val(file.name);
    var reader = new FileReader();
    reader.readAsText(file);
    reader.onload = function (f) {
        gConfig = this.result;      //文件内容
        // console.log("文件内容", gConfig);
    }
};

var sendOneKeyDev = function() {
    var serverId = $("#select_server_id").val();
    if (serverId == "")  {
        layer.msg("请选择服务器");
        return false;
    }
    if (gConfig.length == 0) {
        layer.msg("养成配置为空");
        return false;
    }
    var args = {};
    args.server_id = serverId;
    args.config = gConfig;
    $.post('/gm/doonekeydev', args, oneKeyDevRes, "json");
    // showLoading(true);
};

//发送邮件结果展示
var showResult = function (status, result) {
    var html = "";
    if (status === 0) {
        var txt = "养成成功，账号:"+result;
        html = '<span style="color: green;">' + txt + '</span>';
    } else {
        var txt = "养成失败："+result;
        html = '<span style="color: red;">' + txt + '</span>';
    }
    $("#id_result").html(html);

    var content = "";
    var failDetails = result.FailDetail;
    for (var i in failDetails) {
        var failDetail = failDetails[i];
        content += "错误码:" + failDetail.Code;
        content += " 渠道id:" + failDetail.PlatId;
        content += " 服务器id:" + failDetail.ServerId;
        if (failDetail.UserId > 0) {
            content += " 渠道id:" + failDetail.UserId;
        }
        content += "<br>";
    }
    // console.log(content);
    //点击打开失败详情
    $('#id_fail_detail').on('click', function () {
        layer.open({
            type: 1,
            skin: 'layui-layer-rim', //加上边框
            area: ['420px', '240px'], //宽高
            content: content
        });
    });
};

var oneKeyDevRes =function (data) {
    showLoading(false);
    showResult(data.status,  data.message);
};

layui.use(['layer', 'form', 'jquery', 'upload'], function () {
    layer = layui.layer;
    var $ = layui.jquery;
    var upload = layui.upload;

    //文件上传渲染
    upload.render({
        elem: '#btn_add_file'
        ,url: '/gm/uploadfile'
        , accept: 'file'
        , exts:"xml"
        , auto: true           //不上传到后台
        , size: '10240'  //kb
        ,before: function (obj) {      //选中回调
            obj.preview(function (index, file, result) {
                onChooseFile(file);
                return false;
            });
        }
    });

    //一键养成按钮
    $('#btn_send').on('click', function () {
        sendOneKeyDev();
        return false;
    });

});

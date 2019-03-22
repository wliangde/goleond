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

var sendOneKeyDump = function() {
    var serverId = $("#select_server_id").val();
    if (serverId == "")  {
        layer.msg("请选择导入的服务器");
        return false;
    }

    var dstPlatId = $("#dst_plat_id").val();
    if (dstPlatId == "") {
        layer.msg("请选择目标平台id");
        return false;
    }

    var dstServerId = $("#dst_server_id").val();
    if (dstServerId == "") {
        layer.msg("请选择目标服务器");
        return false;
    }
    var dstUserId = $("#dst_user_id").val();
    if (dstUserId == "") {
        layer.msg("请选择目标玩家");
        return false;
    }


    var args = {};
    args.server_id = serverId;
    args.dst_plat_id = dstPlatId;
    args.dst_server_id = dstServerId;
    args.dst_user_id = dstUserId;

    $.post('/gm/dumpuser', args, oneKeyDumpRes, "json");
    // showLoading(true);
};

//发送邮件结果展示
var showResult = function (status, result) {
    var html = "";
    if (status === 0) {
        var txt = "导入成功，登录账号:"+result;
        html = '<span style="color: green;">' + txt + '</span>';
    } else {
        var txt = "导入失败："+result;
        html = '<span style="color: red;">' + txt + '</span>';
    }
    $("#id_result").html(html);
};

var oneKeyDumpRes =function (data) {
    showLoading(false);
    showResult(data.status,  data.message);
};

layui.use(['layer', 'form', 'jquery'], function () {
    layer = layui.layer;
    var $ = layui.jquery;

    //一键养成按钮
    $('#btn_send').on('click', function () {
        sendOneKeyDump();
        return false;
    });

});

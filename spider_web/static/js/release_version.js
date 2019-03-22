
//区服选择数据
GPlatChannelData = new Map();
GDefPlatId = "";
GDefChannelId = "";  //编辑发布单个渠道

var initPlatChannelData = function () {
    GPlatChannelData = new Map();
};
var layui;

//显示和关闭加载框
var showLoading = function (show) {
    if (show === true) {
        layer.load(2);
    } else {
        layer.closeAll('loading');
    }
};

function getStrPlatChannel() {
    var ret = "";
    if (GDefPlatId != "" && GDefChannelId != "") {          //编辑发布单个渠道
        ret = GDefPlatId+","+GDefChannelId;
        return ret;
    }

    var values = GPlatChannelData.values();
    for (var i  in values) {
        var plat = values[i];
        var serverlist = plat.ChannelList.values();
        var strServer = "";
        for (var j in serverlist) {
            var server = serverlist[j];
            if (server.Selected == false) {
                continue
            }
            if (strServer == "") {
                strServer = server.ChannelId
            } else {
                strServer += "," + server.ChannelId;
            }
        }

        if (strServer != "") {
            strServer = plat.PlatId + "," + strServer;
            if (ret == "") {
                ret = strServer
            } else {
                ret += ";" + strServer;
            }
        }
    }
    // console.log("渠道服务器列表"+ret);
    return ret
}

//发布
var doRelease = function () {
    var args = {};
    args.plat_channel = getStrPlatChannel();
    args.patch = $('#patch').val();
    args.version = $('#version').val();
    args.patch_url = $('#patch_url').val();
    args.min_version = $('#min_version').val();
    args.min_url = $('#min_url').val();

    args.cb_patch = $('#cb_patch').is(':checked');
    args.cb_version = $('#cb_version').is(':checked');
    args.cb_patch_url = $('#cb_patch_url').is(':checked');
    args.cb_min_version = $('#cb_min_version').is(':checked');
    args.cb_min_url = $('#cb_min_url').is(':checked');

    $.post('/gm/dorelease', args, doRealeaseRes, "json");
    showLoading(true);
};

var doRealeaseRes = function (data) {
    //关闭加载框
    console.log(data);
    showLoading(false);
    if (data.status != 0) { //失败
        layer.msg(data.message);
        return;
    }
    var result = data.message;
    showResult(result);
};

//刷新接收人
var refreshTxtRecv = function (txt) {
    $("#txt_recv").val(txt);
};

//初始化
var reset = function () {
    //reset data
    initPlatChannelData();
    //清空结果
    clearResult();
};

var clearResult = function () {
    var html = "";
    $("#id_result").html(html)
};

function sleep(d){
    for(var t = Date.now();Date.now() - t <= d;);
}

//发送邮件结果展示
var showResult = function (result) {
    console.log("wld", result);
    var ok = result.Total === result.Success;
    var html = "";

    if (ok === true) {
        var txt = "版本服" + result.Total + "个，全部成功！";
        html = '<span style="color: green;">' + txt + '</span>';
        //一秒后返回列表界面
        setTimeout(function () {
            window.location.href="/gm/versionlist";
        }, 1000);

    } else {
        var txt = "版本服" + result.Total + "个，成功" + result.Success + "个,失败" + result.Fail + "个";
        html = '<span style="color: red;">' + txt + '</span>';
    }
    $("#id_result").html(html);
};

layui.use(['layer', 'form', 'jquery', 'upload'], function () {
    var form = layui.form;
    layer = layui.layer;
    var $ = layui.jquery;

    //清空数据
    reset();
    //点击增加收件人
    $('#btn_select_chan').on('click', function () {
            layer.open({
                type: 2,
                title: '选择平台渠道',
                shadeClose: true,
                shade: 0.8,
                area: ['900px', '600px'],
                content: 'releaseversionselect.html'
            });
        return false;       //不然layui本身的required 检查会触发
    });
    //发布
    form.on('submit(btn_send)', function (data) {
        doRelease();
        return false;
    });

   GDefPlatId = $('#plat_id').val();
   GDefChannelId = $('#channel_id').val();

   if (GDefChannelId != "") {
       refreshTxtRecv("已选一个渠道 渠道ID:"+GDefChannelId);
   }
});

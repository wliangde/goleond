//收件人类型
GRecvType = 1;

//指定收件人
GReceiverData = new Map();
GLastPaltId = 0;        //上次选择的渠道
GLastServerId = 0;
var initReceiverData = function () {
    GReceiverData = new Map();
    GLastPaltId = 0;
    GLastServerId = 0;
};

//区服选择数据
GPlatServerData = new Map();
GFilterType = 0;
GFilterMinValue = 0;
GFilterMaxValue = 0;
var initPlatServerData = function () {
    GPlatServerData = new Map();
    GFilterType = 0;
    GFilterMinValue = 0;
    GFilterMaxValue = 0;
};
var layui;
//附件
GAttachData = new Map();
var initAttachData = function () {
    GAttachData = new Map();
};
var insertAttachData = function (newAttach) {
    var maxId = 0;
    var values = GAttachData.values();
    if (values.length >= 10) {
        layer.msg("一封邮件只能添加10个附件");
        return;
    }
    for (var i in values) {
        var attach = values[i];
        //策划要求不叠加一起
        // if (attach.res_type == newAttach.res_type && attach.res_id == newAttach.res_id) {
        //     attach.res_num += newAttach.res_num;
        //     layer.msg("添加成功");
        //     refreshAttach();
        //     return true
        // }
        if (attach.id > maxId)
            maxId = attach.id;
    }
    newAttach.id = maxId + 1;
    GAttachData.put(newAttach.id, newAttach);
    layer.msg("添加成功");
    refreshAttach();
    return true;
};
//删除附件
var delAttachData = function (id) {
    GAttachData.remove(id);
    layer.msg("删除成功");
    refreshAttach();
    return true;
};

//文件玩家
var GFileReceiver = "";


//显示和关闭加载框
var showLoading = function (show) {
    if (show === true) {
        layer.load(2);
    } else {
        layer.closeAll('loading');
    }
};

function getStrPlatServer() {
    var values = GPlatServerData.values();
    var ret = "";
    for (var i  in values) {
        var plat = values[i];
        var serverlist = plat.ServerList.values();
        var strServer = "";
        for (var j in serverlist) {
            var server = serverlist[j];
            if (server.Selected == false) {
                continue
            }
            if (strServer == "") {
                strServer = server.Id
            } else {
                strServer += "," + server.Id;
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

function getStrAttach() {
    var values = GAttachData.values();
    var ret = "";
    var bfirst = true;
    for (var i  in values) {
        var attach = values[i];
        if (bfirst == false) {
            ret += ";";
        }
        ret += attach.res_type + "," + attach.res_id + "," + attach.res_num;
        bfirst = false;
    }

    // console.log("附件列表"+ret);
    return ret
}

function getStrTarget() {
    if (GRecvType == 3) { //从文件获得玩家列表
        return GFileReceiver;
    }

    var values = GReceiverData.values();
    var ret = "";
    for (var i in values) {
        var user = values[i];
        var strTarget = user.plat_id + "," + user.svr_id + "," + user.user_id;
        if (ret == "") {
            ret = strTarget;
        } else {
            ret += ";" + strTarget;
        }
    }
    console.log("target" + ret);
    return ret;
}

//类型2，发送系统邮件
var sendSysMail = function () {
    var args = {};
    args.plat_server = getStrPlatServer();
    args.title = $('#txt_title').val();
    args.icon = $('#select_icon').val();
    args.content = $('#txt_content').val();
    args.attach = getStrAttach();
    if (GFilterType > 0) {
        args.filter_type = GFilterType;
        args.filter_min_value = GFilterMinValue;
        args.filter_max_value = GFilterMaxValue;
    }
    $.post('/gm/sendsysmail', args, sendSysMailRes, "json");
    showLoading(true);
};

var sendSysMailRes = function (data) {
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

var sendUserMail = function () {
    var args = {};
    args.target = getStrTarget();
    args.title = $('#txt_title').val();
    args.icon = $('#select_icon').val();
    args.content = $('#txt_content').val();
    args.attach = getStrAttach();
    $.post('/gm/sendusermail', args, sendUserMailRes, "json");
    showLoading(true);
};

var sendUserMailRes = function (data) {
    showLoading(false);
    if (data.status != 0) { //失败
        layer.msg(data.message);
        return;
    }
    var result = data.message;
    showResult(result);
};

//显示或者隐藏button
function showButton(btn, bShow) {
    if (bShow) {
        btn.attr("style", "display: inline-block;margin-left: 0px");
    } else {
        btn.attr("style", "display: none;");
    }
}

//刷新接收人
var refreshTxtRecv = function (txt) {
    $("#txt_recv").val(txt);
};

//刷新附件
var refreshAttach = function () {
    var attachs = GAttachData.values();
    var innerHtml;
    var rowLen = 5;
    if (attachs.length > 0) {
        for (var i = 0; i < attachs.length; i++) {
            attach = attachs[i];
            if (i % rowLen == 0) {
                if (i != 0) {
                    innerHtml += '</tr>'   //把上一行结束掉
                }
                innerHtml += '<tr>'        //开启新的一行
            }
            innerHtml += '<td>' +
                formatOneAttach(attach) +
                '</td>'
        }
    } else {
        innerHtml = "<tr></tr>";
    }
    $("#mail_attach").html(innerHtml);
    // form.render();
};

var formatOneAttach = function (attach) {
    var str = '<div class="star_attatch">' +
        '<table style="width: 100%;height: 100%; ">' +
        '<tr>' +
        '<td class="star_fond_bold">' +
        attach.res_type_name + " " + attach.res_id_name +
        '</td>' +
        '<td class="star_fond_bold" style="text-align: right;">' +
        '<a href="#" style="color: #ff5722;;"onclick="delAttach(' + attach.id + ');return false;">删除</a>' +
        '</td>' +
        '<tr>' +
        '<td class="star_fond_bold">' +
        '数量：' + attach.res_num +
        ' </td>' +
        '        </tr>' +
        '        </table>' +
        '        </div>';

    return str
};

var delAttach = function (id) {
    // alert("删除" +attach);
    delAttachData(id);
};

//初始化
var reset = function () {
    //reset data
    initReceiverData();
    initAttachData();
    initPlatServerData();
    GFileReceiver = "";

    //reset view
    refreshTxtRecv();
    if (GRecvType == 3) {
        showButton($('#btn_add_recv'), false);
        showButton($('#btn_add_file'), true);
        showFileNotice(true);
    } else {
        showButton($('#btn_add_file'), false);
        showButton($('#btn_add_recv'), true);
        showFileNotice(false);
    }

    //清空结果
    clearResult();

    //
    refreshAttach();
};

var showFileNotice = function (show) {
    var html = "";
    if (show === true) {
        html = "文件格式(不支持重复玩家)每行:渠道id,服务器id,玩家id"
    }
    $('#id_file_notice').html(html);
};

var clearResult = function () {
    var html = "";
    $("#id_result").html(html)
};

//发送邮件结果展示
var showResult = function (result) {
    var ok = result.Total === result.Success;
    var html = "";

    if (ok === true) {
        var txt = "总数" + result.Total + "个，全部成功！";
        html = '<span style="color: green;">' + txt + '</span>';
    } else {
        var txt = "总数" + result.Total + "个，成功" + result.Success + "个,失败" + result.Fail + "个（<a class='star_a' id='id_fail_detail'>点击查看失败详情</a>）";
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

//分析文件内容
//文件格式:不支持重复玩家
//渠道id,服id,玩家id
//渠道id,服id,玩家id
var parseFile = function (fileContent) {
    var lineSplit = "\n";
    if (fileContent.indexOf("\r\n") > 0) {
        lineSplit = "\r\n";
    }
    var arrayLine = fileContent.split(lineSplit);

    //简单验证文件格式
    for (var i  in arrayLine) {
        var line = arrayLine[i];
        if (line.length == 0) {
            continue
        }
        var arrayOne = line.split(",");
        if (arrayOne.length !== 3) {
            layer.msg("收件人文件非法的格式：" + line);
            return
        }
    }

    GFileReceiver = arrayLine.join(";");
};

var onChooseFile = function (file) {
    refreshTxtRecv(file.name);
    var reader = new FileReader();
    reader.readAsText(file);
    reader.onload = function (f) {
        var fileContent = this.result;      //文件内容
        parseFile(fileContent);
    }
};

layui.use(['layer', 'form', 'jquery', 'upload'], function () {
    var form = layui.form;
    layer = layui.layer;
    var $ = layui.jquery;
    var upload = layui.upload;

    reset();
    //收件人类型变化

    form.on('select(recv_type)', function (data) {
        if (data.value != GRecvType) {
            GRecvType = data.value;
            reset();
        }
    });

    //点击增加收件人
    $('#btn_add_recv').on('click', function () {
        // $("#resource_key").val();
        if (GRecvType == 1) {    //指定玩家id
            layer.open({
                type: 2,
                title: '添加收件人',
                shadeClose: true,
                shade: 0.8,
                area: ['950px', '620px'],
                content: 'addrecv.html' //iframe的url
            });
        } else if (GRecvType == 2) { //指定区服中满足条件的玩家
            layer.open({
                type: 2,
                title: '选择区服',
                shadeClose: true,
                shade: 0.8,
                area: ['900px', '750px'],
                content: 'selectserver.html'
            });
        }
        return false;       //不然layui本身的required 检查会触发
    });

    //点击删除
    $('#btn_del').on('click', function () {
        reset();
        return false;       //不然layui本身的required 检查会触发
    });

    //点击增加附件
    $('#btn_add_attach').on('click', function () {
        layer.open({
            type: 2,
            title: '添加附件',
            shadeClose: true,
            shade: 0.8,
            area: ['500px', '300px'],
            content: 'addattach.html' //iframe的url
        });
        return false;
    });

    //文件上传渲染
    upload.render({
        elem: '#btn_add_file'
        ,url: '/gm/uploadfile'
        , accept: 'file'
        , auto: true           //不上传到后台
        , size: '2048'  //kb
        , before: function (obj) {      //上传前回调
            obj.preview(function (index, file, result) {
                onChooseFile(file)
            });
        }
    });
    //发送邮件
    form.on('submit(btn_send)', function (data) {
        if (GRecvType == 2) {
            sendSysMail();
        } else {
            sendUserMail();
        }
        return false;
    });
});

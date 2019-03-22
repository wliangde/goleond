var form;
var GSelectPaltId = 0;
var MyPlatChannelData ;

//展示左边平台选择框
var showLeft = function () {
    var html = "";
    var plats = MyPlatChannelData.values();
    var selectAll = true;
    for (var i  in plats) {
        var plat = plats[i];
        var channels = plat.ChannelList.values();
        var selectSize = getSelectSize(channels);
        html += '<div onclick="clickPlat(this);" title="' + plat.PlatId + '" > ';
        if (selectSize > 0) {
            html += '<span class="span_duigou">✔</span>';
        }
        else {
            html += '<span class="span_duigou"></span>';
        }
        html += '<label style="display:inline-block;width: 90px;" >' + plat.Name + '</label>';
        if (selectSize > 0) {
            html += '<span class="layui-badge" >' + selectSize + '</span>';
        }
        // html += '<span class="span_duigou">✔</span><label style="display:inline-block;width: 90px;" >' + plat.Name + '</label><span class="layui-badge" id="span_id_1">129</span>';
        html += '<hr></div>';

        if (selectSize != channels.length) {
            selectAll = false;
        }
    }
    $('#div_plat').html(html);
    $('#plat_chooseall').prop("checked", selectAll);
    form.render(); //必须调用，不然效果不会显示出来
};

var showRight = function () {
    var plat = MyPlatChannelData.get(GSelectPaltId);
    if (plat == null)
        return;
    $('#id_plat_name').text(plat.Name);
    var innerHtml = "";
    var channelList = plat.ChannelList.values();
    var selectAll = true;
    for (var i in channelList) {
        channel = channelList[i];
        if (channel.Selected == true) {
            innerHtml += '<input type="checkbox" name="ch_server" lay-filter="server_choose"  value="' + channel.ChannelId + '"  id="id_plat_server_' + channel.ChannelId + '" lay-skin="primary" title="' + channel.Name + '" checked>';
        } else {
            innerHtml += '<input type="checkbox" name="ch_server" lay-filter="server_choose"  value="' + channel.ChannelId + '" id="id_plat_server_' + channel.ChannelId + '" lay-skin="primary" title="' + channel.Name + '">';
            selectAll = false;
        }
    }
    $("#id_server_list").html(innerHtml);

    $('#server_chooseall').prop("checked", selectAll);

    form.render(); //必须调用，不然效果不会显示出来
};

var clickPlat = function (plat) {
    GSelectPaltId = plat.title;
    showRight();
};

var selectAllPlat = function(selected) {
    var plats = MyPlatChannelData.values();
    for (var i in plats) {
        var plat = plats[i];
        servers = plat.ChannelList.values();
        for (var j in servers) {
            var server = servers[j];
            server.Selected = selected;
        }
    }
    showLeft();
    showRight()
};

var selectAllServer = function (selected) {
    var plat = MyPlatChannelData.get(GSelectPaltId);
    if (plat == null)
        return;
    var channelList = plat.ChannelList.values();
    for (var i in channelList) {
        var server = channelList[i];
        server.Selected = selected;
    }
    showLeft();
    showRight();
};

var selectOneServer = function (serverId, selected) {
    var plat = MyPlatChannelData.get(GSelectPaltId);
    if (plat == null)
        return;
    var server = plat.ChannelList.get(serverId);
    if (server == null)
        return;
    server.Selected = selected;
    showLeft();
    showRight();
};

var getSelectSize = function (servers) {
    var count = 0;
    for (var i in servers) {
        var server = servers[i];
        if (server.Selected == true) {
            count++;
        }
    }
    return count;
};

var cloneMap = function (fromData) {
    var toData = new Map();
    var plats = fromData.values();
    for (var i in plats) {
        var plat = plats[i];
        var plat2 = new Object();
        plat2.PlatId = plat.PlatId;
        plat2.Name = plat.Name;
        plat2.ChannelList = new Map();
        var servers = plat.ChannelList.values();
        for (var j in servers) {
            var server = servers[j];
            var server2 = new Object();
            server2.ChannelId = server.ChannelId;
            server2.Name = server.Name;
            server2.Selected = server.Selected;
            plat2.ChannelList.put(server2.ChannelId, server2);
        }
        toData.put(plat2.PlatId, plat2);
    }
    return toData;
};

layui.use(['layer', 'table', 'form', 'element', 'jquery'], function () {
    form = layui.form;
    var layer = layui.layer;
    var $ = layui.jquery;
    MyPlatChannelData =cloneMap( parent.GPlatChannelData);
    //监听全平台选择
    form.on('checkbox(plat_chooseall)', function (data) {
        selectAllPlat(data.elem.checked);
    });
    //监听全渠道选择
    form.on('checkbox(server_chooseall)', function (data) {
        selectAllServer(data.elem.checked)
    });
    //勾选某个渠道
    form.on('checkbox(server_choose)', function (data) {
        serverId = data.value;
        selectOneServer(serverId, data.elem.checked);
    });

    $('#btn_add').on('click', function () {

        parent.GPlatChannelData = cloneMap(MyPlatChannelData);

        var strTxt ="";
        var strPlatName="";
        var platCnt = 0;
        var serverCnt = 0;
        var plats = parent.GPlatChannelData.values();
        var selectAll = true;
        for (var i  in plats) {
            var plat = plats[i];
            var servers = plat.ChannelList.values();
            var selectSize = getSelectSize(servers);
            if (selectSize > 0) {
                if (platCnt === 0) {
                    strPlatName = plat.Name;
                }
                platCnt++;
                serverCnt+=selectSize;
            }
        }
        if (platCnt !=0) {
            strTxt+=strPlatName+"等"+platCnt+"个平台共"+serverCnt+"个渠道";
        }
        layer.msg("保存成功！");
        parent.refreshTxtRecv(strTxt);
        //关闭弹窗
        var index = parent.layer.getFrameIndex(window.name); //获取窗口索引
        parent.layer.close(index);//关闭弹出的子页面窗口
    });

    //关闭弹窗
    $('#btn_cancel').on('click', function () {
        //关闭弹窗
        var index = parent.layer.getFrameIndex(window.name); //获取窗口索引
        parent.layer.close(index);//关闭弹出的子页面窗口
    });

    //加载平台和服列表
    var loadPlatChannel = function () {
        $.post('/gm/platchannel', "", loadPlatChannelRes, "json");
    };
    var loadPlatChannelRes = function (data) {
        if (data.status != 0) {
            layer.msg(data.message);
            return;
        }
        platChannelList = data.message;
        for (var i in platChannelList) {
            var plat = platChannelList[i];
            var platChannel = MyPlatChannelData.get(plat.PlatId);
            if (GSelectPaltId == 0) {
                GSelectPaltId = plat.PlatId;        //设置一个默认值
            }

            if (platChannel == null) {
                platChannel = Object();
                platChannel.PlatId = plat.PlatId;
                platChannel.Name = plat.Name;
                platChannel.ChannelList = new Map();
                MyPlatChannelData.put(plat.PlatId, platChannel);
                platChannel = MyPlatChannelData.get(plat.PlatId);
            }

            for (var j in plat.ChannelList) {
                var channel = plat.ChannelList[j];
                var channel2 = platChannel.ChannelList.get(channel.ChannelId);
                if (channel2 == null) {
                    if (platChannel.PlatId == parent.GDefPlatId && channel.ChannelId == parent.GDefChannelId) {
                        channel.Selected = true;
                        parent.GDefPlatId = "";
                        parent.GDefChannelId = "";
                    } else {
                        channel.Selected = false;
                    }
                    platChannel.ChannelList.put(channel.ChannelId, channel);
                }
            }
        }
        showLeft();
        showRight();
    };
    //加载平台渠道
    loadPlatChannel();
});

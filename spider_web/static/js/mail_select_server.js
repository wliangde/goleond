var form;
var GSelectPaltId = 0;
var MyPlatServerData ;

layui.use(['layer', 'table', 'form', 'element', 'jquery'], function () {
    form = layui.form;
    var layer = layui.layer;
    var $ = layui.jquery;
    MyPlatServerData =cloneMap( parent.GPlatServerData);
    //监听全渠道选择
    form.on('checkbox(plat_chooseall)', function (data) {
        selectAllPlat(data.elem.checked);
    });
    //监听全服选择
    form.on('checkbox(server_chooseall)', function (data) {
        selectAllServer(data.elem.checked)
    });
    //勾选某个服
    form.on('checkbox(server_choose)', function (data) {
        serverId = data.value;
        selectOneServer(serverId, data.elem.checked);
    });

    $('#btn_add').on('click', function () {

        var filterType = $('#filter_type').val();
        var strFilerType=$("#filter_type").find("option:selected").text();
        var minValue = $('#filter_min_value').val();
        var maxValue = $('#filter_max_value').val();

        if (minValue > maxValue) {
            layer.msg("请输入正确的区间");
            return;
        }
        var strFilter = "-"+strFilerType+"-"+minValue+"-"+maxValue;
        parent.GFilterType = filterType;
        parent.GFilterMinValue = minValue;
        parent.GFilterMaxValue = maxValue;
        parent.GPlatServerData = cloneMap(MyPlatServerData);

        var strTxt ="";
        var strPlatName="";
        var platCnt = 0;
        var serverCnt = 0;
        var plats = parent.GPlatServerData.values();
        var selectAll = true;
        for (var i  in plats) {
            var plat = plats[i];
            var servers = plat.ServerList.values();
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
            strTxt+=strPlatName+"等"+platCnt+"个渠道共"+serverCnt+"个服务器";
            if (strFilter != "") {
                strTxt+=strFilter;
            }
        }
        layer.msg("保存成功！");
        parent.refreshTxtRecv(strTxt);
    });

    //关闭弹窗
    $('#btn_cancel').on('click', function () {
        //关闭弹窗
        var index = parent.layer.getFrameIndex(window.name); //获取窗口索引
        parent.layer.close(index);//关闭弹出的子页面窗口
    });

    //加载平台和服列表
    var loadPlatServer = function () {
        $.post('/gm/platserver', "", loadPlatServerRes, "json");
    };
    var loadPlatServerRes = function (data) {
        if (data.status != 0) {
            layer.msg(data.message);
            return;
        }
        platServerList = data.message;
        for (var i in platServerList) {
            var plat = platServerList[i];
            var platServer = MyPlatServerData.get(plat.PlatId);
            if (GSelectPaltId == 0) {
                GSelectPaltId = plat.PlatId;        //设置一个默认值
            }

            if (platServer == null) {
                platServer = Object();
                platServer.PlatId = plat.PlatId;
                platServer.Name = plat.Name;
                platServer.ServerList = new Map();
                MyPlatServerData.put(plat.PlatId, platServer);
                platServer = MyPlatServerData.get(plat.PlatId);
            }

            for (var j in plat.ServerList) {
                var server = plat.ServerList[j];
                var server2 = platServer.ServerList.get(server.Id);
                if (server2 == null) {
                    server.Selected = false;
                    platServer.ServerList.put(server.Id, server);
                }
            }
        }
        showLeft();
        showRight();
    };
    loadPlatServer();

    $('#filter_type').val(parent.GFilterType);
    $('#filter_min_value').val(parent.GFilterMinValue);
    $('#filter_max_value').val(parent.GFilterMaxValue);
    console.log(parent.GFilterMinValue);
});

//展示左边平台选择框
var showLeft = function () {
    var html = "";
    var plats = MyPlatServerData.values();
    var selectAll = true;
    for (var i  in plats) {
        var plat = plats[i];
        var servers = plat.ServerList.values();
        var selectSize = getSelectSize(servers);
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

        if (selectSize != servers.length) {
            selectAll = false;
        }
    }
    $('#div_plat').html(html);
    $('#plat_chooseall').prop("checked", selectAll);
    form.render(); //必须调用，不然效果不会显示出来
};

var showRight = function () {
    var plat = MyPlatServerData.get(GSelectPaltId);
    if (plat == null)
        return;
    $('#id_plat_name').text(plat.Name);
    var innerHtml = "";
    var severList = plat.ServerList.values();
    var selectAll = true;
    for (var i in severList) {
        server = severList[i];
        if (server.Selected == true) {
            innerHtml += '<input type="checkbox" name="ch_server" lay-filter="server_choose"  value="' + server.Id + '"  id="id_plat_server_' + server.Id + '" lay-skin="primary" title="' + server.Name + '" checked>';
        } else {
            innerHtml += '<input type="checkbox" name="ch_server" lay-filter="server_choose"  value="' + server.Id + '" id="id_plat_server_' + server.Id + '" lay-skin="primary" title="' + server.Name + '">';
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
    var plats = MyPlatServerData.values();
    for (var i in plats) {
        var plat = plats[i];
        servers = plat.ServerList.values();
        for (var j in servers) {
            var server = servers[j];
            server.Selected = selected;
        }
    }
    showLeft();
    showRight()
};

var selectAllServer = function (selected) {
    var plat = MyPlatServerData.get(GSelectPaltId);
    if (plat == null)
        return;
    var severList = plat.ServerList.values();
    for (var i in severList) {
        var server = severList[i];
        server.Selected = selected;
    }
    showLeft();
    showRight();
};

var selectOneServer = function (serverId, selected) {
    var plat = MyPlatServerData.get(GSelectPaltId);
    if (plat == null)
        return;
    var server = plat.ServerList.get(serverId);
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
        plat2.ServerList = new Map();
        var servers = plat.ServerList.values();
        for (var j in servers) {
            var server = servers[j];
            var server2 = new Object();
            server2.Id = server.Id;
            server2.Name = server.Name;
            server2.Selected = server.Selected;
            plat2.ServerList.put(server2.Id, server2);
        }
        toData.put(plat.PlatId, plat2);
    }
    return toData;
};

var printData = function (data) {
    console.log(data);
    var plats = data.values();
    for (var i in plats) {
        var plat = plats[i];
        var servers = plat.ServerList.values();
        for (var j in servers) {
            var server = servers[j];
            // console.log(plat.PlatId+" "+ server.Name+ server.Selected)
        }
    }
};
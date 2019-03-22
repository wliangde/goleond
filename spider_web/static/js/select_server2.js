var form;
var GSelectedServerType = 1;  //默认选择的服务器类型
var MyPlatServerData;

layui.use(['layer', 'table', 'form', 'element', 'jquery'], function () {
    form = layui.form;
    var layer = layui.layer;
    var $ = layui.jquery;
    MyPlatServerData = cloneMap(parent.GPlatServerData);
    //监听所有服务器类型选择
    form.on('checkbox(plat_chooseall)', function (data) {
        selectAllServerType(data.elem.checked);
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

        parent.GPlatServerData = cloneMap(MyPlatServerData);

        var strTxt = "";
        var serverCnt = 0;
        var serverList = parent.GPlatServerData.values();
        for (var i  in serverList) {
            var ptServer = serverList[i];
            if (ptServer.Selected == true) {
                if (i<4) {
                    strTxt +=ptServer.Name+" ";
                }
                if (i==4) {
                    strTxt +="等";
                }
                serverCnt++;
            }

        }
        if (serverCnt > 0) {
            strTxt +=  "共" + serverCnt + "个服务器";
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
        //通过服务器类型更新下拉列表的服
        $.getJSON("/gm/getbriefseverlist?server_type=0", loadPlatServerRes);
    };
    var loadPlatServerRes = function (data) {
        var optionstring = "";
        $.each(data.data, function (i, item) {
            var ptServer = MyPlatServerData.get(item.Id);
            if (ptServer == null) {
                //============warning:battle 服不需要加载配置
                // if (item.Type == 3) {
                //     return
                // }
                ptServer = Object();
                ptServer.Id = item.Id;
                ptServer.Type = item.Type;
                ptServer.Name = item.Name;
                ptServer.Selected = false;
                MyPlatServerData.put(ptServer.Id, ptServer);
            }
        });
        $("#server_id").html(optionstring);
        showLeft();
        showRight();
    };
    loadPlatServer();

    $('#filter_type').val(parent.GFilterType);
    $('#filter_min_value').val(parent.GFilterMinValue);
    $('#filter_max_value').val(parent.GFilterMaxValue);
    console.log(parent.GFilterMinValue);
});

//展示左边服务器类型选择框
var showLeft = function () {
    var html = "";
    var selectAll = true;

    var funcProc = function (serverType, serverName) {
        var ptSelectedData = getSelectSize(serverType);

        html += '<div onclick="clickServerType(this);" title="' + serverType + '" > ';
        if (ptSelectedData.selectedSize > 0) {
            html += '<span class="span_duigou">✔</span>';
        }
        else {
            html += '<span class="span_duigou"></span>';
        }
        html += '<label style="display:inline-block;width: 90px;" >' + serverName + '</label>';
        if (ptSelectedData.selectedSize > 0) {
            html += '<span class="layui-badge" >' + ptSelectedData.selectedSize + '</span>';
        }
        // html += '<span class="span_duigou">✔</span><label style="display:inline-block;width: 90px;" >' + plat.Name + '</label><span class="layui-badge" id="span_id_1">129</span>';
        html += '<hr></div>';
        if (ptSelectedData.selectedSize != ptSelectedData.allSize) {
            selectAll = false;
        }
    };

    funcProc(1, "game");
    funcProc(2, "cross");
    funcProc(3, "battle");

    $('#div_plat').html(html);
    $('#plat_chooseall').prop("checked", selectAll);
    form.render(); //必须调用，不然效果不会显示出来
};

var showRight = function () {
    var serverTypeName = function () {
        if (GSelectedServerType == 1) {
            return "game";
        }else if (GSelectedServerType == 2){
            return "cross";
        }
    };

    $('#id_plat_name').text(serverTypeName());
    var innerHtml = "";
    var severList = MyPlatServerData.values();
    var selectAll = true;
    for (var i in severList) {
        server = severList[i];
        if (server.Type != GSelectedServerType) {
            continue;
        }
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

//点击左边的服务器类型
var clickServerType = function (plat) {
    GSelectedServerType = plat.title;
    showRight();
};

//全选所有服务器
var selectAllServerType = function (selected) {
    var servers = MyPlatServerData.values();
    for (var i in servers) {
        var server = servers[i];
        server.Selected = selected;
    }
    showLeft();
    showRight()
};

//某个类型服务器全选
var selectAllServer = function (selected) {
    var severList = MyPlatServerData.values();
    for (var i in severList) {
        var server = severList[i];
        if (server.Type != GSelectedServerType) {
            continue
        }
        server.Selected = selected;
    }
    showLeft();
    showRight();
};

var selectOneServer = function (serverId, selected) {
    var ptServer = MyPlatServerData.get(serverId);
    if (ptServer == null)
        return;
    ptServer.Selected = selected;
    showLeft();
    showRight();
};

var getSelectSize = function (serverType) {
    var selectedSize = 0;
    var allSize = 0;
    var servers = MyPlatServerData.values();
    for (var i in servers) {
        var server = servers[i];
        if (server.Type == serverType) {
            allSize++;
            if (server.Selected == true) {
                selectedSize++;
            }
        }
    }
    return {selectedSize: selectedSize, allSize: allSize};
};

var cloneMap = function (fromData) {
    var toData = new Map();
    var servers = fromData.values();
    for (var i in servers) {
        var server = servers[i];
        var server2 = new Object();
        server2.Type = server.Type;  //服务类型
        server2.Id = server.Id;
        server2.Name = server.Name;
        server2.Selected = server.Selected;
        toData.put(server2.Id, server2);
    }
    return toData;
};

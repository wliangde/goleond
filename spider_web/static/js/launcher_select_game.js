var form;

//动态展示服选择框
var refreshUi = function () {
    var innerHtml = "";
    var serverList = MyServerData.values();
    var selectAll = true;
    for (var i in serverList) {
        server = serverList[i];
        if (server.Selected == true) {
            innerHtml += '<input type="checkbox" name="ch_server" lay-filter="server_choose"  value="' + server.Id + '"  id="id_plat_server_' + server.ChannelId + '" lay-skin="primary" title="' + server.Name + '" checked>';
        } else {
            innerHtml += '<input type="checkbox" name="ch_server" lay-filter="server_choose"  value="' + server.Id + '" id="id_plat_server_' + server.ChannelId + '" lay-skin="primary" title="' + server.Name + '">';
            selectAll = false;
        }
    }
    $("#id_server_list").html(innerHtml);

    $('#server_chooseall').prop("checked", selectAll);

    form.render(); //必须调用，不然效果不会显示出来
};

//全选服务器
var selectAllServer = function (selected) {
    var serverList = MyServerData.values();
    for (var i in serverList) {
        var server = serverList[i];
        server.Selected = selected;
    }
    refreshUi();
};

//选择某个服务器
var selectOneServer = function (serverId, selected) {
    var server = MyServerData.get(serverId);
    if (server == null)
        return;
    server.Selected = selected;
    refreshUi();
};

var cloneMap = function (fromData) {
    var toData = new Map();
    var servers = fromData.values();
    for (var j in servers) {
        var server = servers[j];
        var server2 = new Object();
        server2.Id = server.Id;
        server2.Name = server.Name;
        server2.Selected = server.Selected;
        toData.put(server2.Id, server2);
    }
    return toData;
};

layui.use(['layer', 'table', 'form', 'element', 'jquery'], function () {
    form = layui.form;
    var layer = layui.layer;
    var $ = layui.jquery;

    MyServerData = cloneMap(parent.GServerData);

    //监听全选
    form.on('checkbox(server_chooseall)', function (data) {
        selectAllServer(data.elem.checked)
    });
    //监听选择某个服
    form.on('checkbox(server_choose)', function (data) {
        serverId = data.value;
        selectOneServer(serverId, data.elem.checked);
    });
    $('#btn_add').on('click', function () {
        parent.GServerData = cloneMap(MyServerData);
        //父窗口的文本提示
        var strTxt = "";
        var servers = parent.GServerData.values();
        for (var i  in servers) {
            var server = servers[i];
            if (server.Selected == false){
                continue;
            }
            if (strTxt != "") {
                strTxt += " | " + server.Id;
            }else {
                strTxt += server.Id;
            }
        }
        layer.msg("保存成功！");
        parent.refreshGameTxt(strTxt);
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

    //加载服列表，1：game 2:cross 3:battle
    var loadServerList = function (serverType) {
        $.getJSON("/gm/getbriefseverlist?server_type=" + serverType, loadServerRes);
    };
    //加载服返回
    var loadServerRes = function (data) {
        console.log(data.data);
        $.each(data.data, function (i, item) {
            var server = MyServerData.get(item.Id);
            if (server == null) {
                server = Object();
                server.Id = item.Id;
                server.Name = item.Name;
                server.Selected = false;
                MyServerData.put(item.Id, server);
            }else {
                server.Name = item.Name; //父窗口没有赋值名字
            }
        });
        refreshUi();
    };
    //加载服
    loadServerList(1);
});

//vue
//用delimiters
// Vue.config.delimiters = ['[[', ']]'];
//
// var vue1 = new Vue({
//     delimiters: ['[[', ']]'],
//     el:"#select_server",
//     data:{
//         serverList:[
//             {id:1, name:"wld", checked:true, value:"a"},
//             {id:2, name:"wld2", checked:true, value:"a"},
//         ]
//     },
//     computed:{
//         select_all:function () {
//             return this.serverList.every(function (value) { return value.checked });
//         }
//     }
// });

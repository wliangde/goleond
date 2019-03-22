GServerData = new Map();
var initServerData = function () {
    GServerData = new Map();

    var strSelectServers = $('#game_servers').val();
    if (strSelectServers === "") {
        return;
    }
    var arSelectServers = strSelectServers.split(" | ");
    //遍历数组
    for (let serverId of arSelectServers) {
        // console.log(serverId+" "+strSelectServers);
        server = new Object();
        server.Id = serverId;
        server.Name = "占位";
        server.Selected = true;
        GServerData.put(serverId, server)
    }
};

//刷新game servers
var refreshGameTxt = function (txt) {
    $("#game_servers").val(txt);
};

var $;
layui.config({
    base : "js/"
}).use(['form','element','layer','jquery'],function(){
    var form = layui.form; //只有执行了这一步，部分表单元素才会自动修饰成功
    var $ = layui.jquery;
    var role_ids = [];

    //初始化
    initServerData();

    //点击修改game
    $('#btn_select_game').on('click', function () {
        layer.open({
            type: 2,
            title: '选择game',
            shadeClose: true,
            shade: 0.8,
            area: ['900px', '600px'],
            content: 'select_game.html'
        });
        return false;       //不然layui本身的required 检查会触发
    });

    form.on('select(cross)', function(data){
        var t = $('#cross_servers').val();
        var newId = data.value+" ";
        if (t.indexOf(newId) !== -1) {
            return
        }

        if (t.length >0) {
            t = t + '| ';
        }
        t = t + newId;
        $('#cross_servers').val(t);
    });

    form.on('select(battle)', function(data){
        var t = $('#battle_servers').val();
        var newId = data.value+" ";
        if (t.indexOf(newId) !== -1) {
            return
        }

        if (t.length >0) {
            t = t + '| ';
        }
        t = t + newId;
        $('#battle_servers').val(t);
    });

    //清空cross
    $('#btn_clear_cross').on('click', function () {
        $('#cross_servers').val("");
        return false;   //return false 不会再刷新页面，如果是return true的话页面会再次使用已有的值刷新
    });

    //清空battle
    $('#btn_clear_battle').on('click', function () {
        $('#battle_servers').val("");
        return false;
    });

    form.on('submit(sub)', function(data){
        var form_data = $("form").serialize();
        $.post('/launcher/ajaxnew', form_data, function (out) {
            if (out.status === 0) {
                layer.msg("操作成功",{icon: 1},function () {
                    // window.location.reload()
                    //为了批量操作，这里不进行跳转
                    // window.location.href="/launcher/list"
                })
            } else {
                layer.msg(out.message)
            }
        }, "json");
        return false;
    });
    form.render();
});
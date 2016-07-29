<!-- 引入主框架 -->
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <!-- <title>便利</title> -->
    <title>{{.Title}}</title>
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" />
    <meta charset="utf-8">
    <!-- <link rel="shortcut icon" href="/static/img/go.ico"> -->
    <link rel="stylesheet" href="http://cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
    <!-- 可选的Bootstrap主题文件（一般不用引入） -->
    <link rel="stylesheet" href="http://cdn.bootcss.com/font-awesome/3.0.2/css/font-awesome.css">

    <link rel="stylesheet" href="/static/css/goods-wrap.css">
    <link rel="stylesheet" href="/static/css/home.css">
    <link rel="stylesheet" href="/static/css/cart.css"> 
    
    <!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
    <script src="http://cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <script src="http://cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
    <script src="http://cdn.bootcss.com/dot/1.0.3/doT.min.js"></script> 
</head>

<body data-spy="scroll" data-target="#myScrollspy">
    {{template "content" .}}
    {{template "bottom-navbar" .}}
    <div id="cart-ball"></div>
</body>
    {{template "js" .}}
    <script type="text/javascript">
        $(".trigger").click(function(e) {
            var self = $(this);
            var trigger = self.attr("trigger");
            if (undefined != trigger && "" != trigger) {
                e.preventDefault();
                eval(trigger)();
            }
        });
        $(".item-cart-click").click(editCart);
    </script>
</html>

{{define "content"}}
{{$globalContext := .GlobalContext}}
{{with .Content}}
    <div class="main-font" style="font-size: 10px;">
        <div class="appground container-inner" style="top: 0;color: #ffffff;position: fixed;z-index: 3;">
            <p style="margin: 6px 6px;">
                <span>回归</span>
                <span style="font-style: italic;font-weight: bold;">十六度</span>
                <span>暖暖的味道</span>
            </p>
        </div>
        <div class="row" style="margin-top: 30px;">
            <div class="col-xs-3" id="myScrollspy" style="font-size: 12px;width: 27%;">
                <ul class="nav nav-tabs nav-stacked background affix-top" data-spy="affix" data-offset-top="125" style="width: 26%;position: fixed;text-align: center;">
                    {{range $idx,$classInfo := .ClassList}}
                    <li class="{{if eq $idx 0}} active {{end}}">
                        <a style="color:black;border: none;" href="#section-{{$idx}}">
                            {{$classInfo.ClassName}}
                        </a>
                    </li>
                    {{end}}
                </ul>
            </div>
            <div class="col-xs-9" style="width: 71%;padding-left: 1%;border-left: 1px solid #afabab;margin-left: 3px;">
                {{range $idx,$classInfo := .ClassList}}
                <div id="section-{{$idx}}">
                    <div style="width: 100%;padding: 6px 0px 3px 10px;" class="background">
                        <span style="font-weight: 300;color: black;white-space: nowrap;text-overflow: ellipsis;overflow: hidden;">
                            {{$classInfo.ClassName}}
                        </span>
                    </div>
                    {{range $goodsIdx,$goodsInfo := $classInfo.GoodsList}}
                    <div class="goods" style="width: 100%;padding:7% 0px 3%;border-bottom: 2px solid #f7f5f5;"> 
                        <div style="width: 38%;display: inline-block;vertical-align: top;margin-right: 3%;"> 
                            <a data-toggle="modal" data-target="#detailModelOutter" data-whatever="{{addGlobalContext $globalContext $goodsInfo.Id $goodsInfo}}">
                                <img src="{{$goodsInfo.Image}}" style="width: 100%;height:85px;">
                            </a>
                        </div>
                        <div class="inner" style="display: inline-block;width: 50%;vertical-align: top;position: relative;height: 90px;">
                            <div class="cart-count-flag"> 
                                <div class="item-cart-click background" goods-id="{{$goodsInfo.Id}}" style="display:none;top: 50%;left: 47%;color: red;font-weight: 600;position: absolute;padding: 2px 9px;border-radius: 50%;"> 
                                    <div class="count">
                                        <span>X</span>
                                        <span goods-id="{{$goodsInfo.Id}}"></span>
                                    </div>
                                </div>
                            </div>
                            <p style="  34px;font-weight: 300;color: black;">{{$goodsInfo.Name}}</p>
                            <p style="  30px;position: absolute;bottom: 16%;font-weight: 500;color: #999;">{{$goodsInfo.Norms}}</p>
                            <p style="position: absolute;bottom: -6%;">
                                <span class="price" style="color: red;margin-right: 5px;  30px">
                                    ￥{{$goodsInfo.Price}}
                                </span>
                                <span style="font-weight: 500;color: #999;">
                                    <del>￥{{$goodsInfo.Marketprice}}</del>
                                </span>
                            </p>
                            <div class="item-edit" style="position: absolute;right: -10%;bottom: 0%;">
                                <div class="item-cart item-cart-click" value="1" style="" goods-id="{{$goodsInfo.Id}}">
                                    <div class="btn -purchase">
                                        <i class="icon-plus"></i>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    {{end}}
                </div>
                {{end}}
            </div>
        </div>
    </div>
{{end}}
<div class="modal fade modal-goods-detail" id="detailModelOutter" tabindex="-1" role="dialog" aria-labelledby="detailMobal"> 
    <div class="modal-dialog modal-sm detailModelInner" role="document"> 
        <div class="modal-content background"> 
            <div class="title"> 
                <span class="goods-title" id="detailMobal">商品详情</span> 
            </div> 
            <div class="goods" style="width: 100%;text-align:center;margin-top: 1px;"> 
                <div class="foreground" style="height: 190px;"> 
                    <img class="goods-img" style="width:80%;height:188px;" src="http://img01.bqstatic.com/upload/goods/201/605/2409/20160524095244_729759.jpg@200w_200h_90Q" /> 
                </div> 
                <div class="detail">
                    <p> 
                        <span class="goods-name"></span>
                        <span class="goods-norms"></span>
                    </p> 
                    <p> 
                        <span class="price"></span>
                        <span><del style="font-size: 10px;"></del></span>
                    </p> 
                    <div class="item-edit">
                        <div class="item-cart item-cart-click" value=-1 style="bottom: 30px;right: 58px;"> 
                            <div class="btn -purchase"> 
                                <i class="icon-minus"></i> 
                            </div> 
                        </div> 
                        <div class="item-cart" style="bottom: 29px;right: 39px;"> 
                            <i class="count"><span>0</span></i> 
                        </div> 
                        <div class="item-cart item-cart-click" value=1 style="bottom: 30px;right: 8px;"> 
                            <div class="btn -purchase"> 
                                <i class="icon-plus"></i> 
                            </div> 
                        </div>
                    </div>
                </div>
                </div> 
                    <div class="desc"> 
                </div> 
            </div>
        </div>
    </div>
</div>
{{end}}
{{define "js"}}
    <script type="text/javascript">
    $('#detailModelOutter').on('show.bs.modal', function (event) {
        // Button that triggered the modal
        var button = $(event.relatedTarget);
        // Extract info from data-* attributes
        var goodsId = button.data('whatever');
        // If necessary, you could initiate an AJAX request here (and then do the updating in a callback).
        // Update the modal's content. We'll use jQuery here, but you could use a data binding library or other methods instead.
        var modal = $(this);
        modal.find('.goods-img').attr("src",globalContext[goodsId].Image);
        modal.find('.goods-name').text(globalContext[goodsId].Name);
        modal.find('.goods-norms').text(globalContext[goodsId].Norms);
        modal.find('.price').text("￥" + globalContext[goodsId].Price);
        modal.find('del').text("￥" + globalContext[goodsId].Marketprice);
        modal.find('.count').find("span").attr("goods-id",goodsId);
        modal.find('.count').find("span").html($(".count").find("span[goods-id="+goodsId+"]").html());
        modal.find('.item-cart').attr("goods-id",goodsId);
    });
    $("#foot-menu").find(".active").removeClass("active");
    $("#foot-menu").find("#list").addClass("active");
</script>
    <!-- 引用 goods-wrap.js -->
    <script src="/static/js/goods-wrap.js"></script>
    <script src="/static/js/cart.js"></script>
    <script>var globalContext = {{.GlobalContext}}; var globalCart=getCart();initCart();</script>
{{end}}
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
	<link rel="stylesheet" href="/static/css/order.css"> 
	
	<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
	<script src="http://cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
	<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
	<script src="http://cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
	<script src="http://cdn.bootcss.com/dot/1.0.3/doT.min.js"></script> 
</head>

<body class="background">
	{{template "content" .}}
</body>
	{{template "js" .}}
</html>

{{define "content"}}
{{with .Content}}
	<div id="order-list">
		<div id="order-info" order-sn="{{.OrderInfo.Order.OrderSn}}" style="display:none"></div>
		<div class="container-outter foreground" style="margin-top: 10px;">
			<div class="inner">
				<div class="process-info container-inner" style="text-align: center;padding: 2px 2px 0px 9px;">
					<div class="appborder" style="border-width: 2px; border-radius: 50%;padding: 9px;display: inline-block;width: 13%;float: left;margin: 19px;">
						{{if lt .OrderInfo.Order.OrderState 0}}
							<i class="icon-warning-sign icon-2x appcolor" aria-hidden="true" style="font-size: 16px;"></i>
						{{else if lt .OrderInfo.Order.OrderState 10}}
							<i class="icon-credit-card icon-2x appcolor" aria-hidden="true" style="font-size: 16px;"></i>
						{{else if lt .OrderInfo.Order.OrderState 20}}
							<i class="icon-phone icon-2x appcolor" aria-hidden="true" style="font-size: 16px;"></i>
						{{else if lt .OrderInfo.Order.OrderState 30}}
							<i class="icon-plane icon-2x appcolor" aria-hidden="true" style="font-size: 16px;"></i>
						{{else}}
							<i class="icon-ok icon-2x appcolor" aria-hidden="true" style="font-size: 16px;"></i>
						{{end}}
					</div>
					<div style="float: left;">
						<div style="float: left;text-align: left;">
						{{if gt .OrderInfo.Order.CancelFlag 0}}
							<div style="font-size: 19px;margin-top: 19px;">订单取消</div>
							<p style="font-size: 10px;">谢谢惠顾，欢迎再来</p>
						{{else if lt .OrderInfo.Order.OrderState 10}}
							<div style="font-size: 19px;margin-top: 19px;">等待付款</div>
							<p style="font-size: 10px;">待支付金额: ￥{{.OrderInfo.Order.OrderAmount}}</p>
						{{else if lt .OrderInfo.Order.OrderState 20}}
							<div style="font-size: 19px;margin-top: 19px;">支付成功</div>
							<p style="font-size: 10px;">等待商家确认...</p>
						{{else if lt .OrderInfo.Order.OrderState 30}}
							<div style="font-size: 19px;margin-top: 19px;">已经确认</div>
							<p style="font-size: 10px;">小二正在取货，稍等~~</p>
						{{else if lt .OrderInfo.Order.OrderState 40}}
							<div style="font-size: 19px;margin-top: 19px;">正在配送</div>
							<p style="font-size: 10px;">马上就到，等我...</p>
						{{else}}
							<div style="font-size: 19px;margin-top: 19px;">订单送达</div>
							<p style="font-size: 10px;">请对我们服务进行评价!</p>
						{{end}}
						{{if lt .OrderInfo.Order.OrderState 10}}
							<p style="font-size: 10px; color:red">超时15分钟后，订单自动取消</p>
						{{else}}
							<p></p>
						{{end}}
					   </div>
					</div>
				</div>
			</div>
		</div>
		<div class="container-outter foreground" style="margin-top: 10px;">
			<div class="inner">
				<div class="process-info container-inner" style="text-align: left;padding: 2px 2px 0px 9px;">
					<div style="display: inline-block;width: 15%;text-align: center;">
						<label class="appground appgoundborder" style="{{if lt .OrderInfo.Order.OrderState 10}}background-color:grey;{{end}} border-width: 2px;border-radius: 50%;padding: 10px 13px 10px 12px;">
							<i class="icon-credit-card icon-2x" aria-hidden="true" style="color: white;font-size: 16px;"></i>
						</label>
						<p style="font-size: 12px;">待支付</p>
					</div>
					<span style="display: inline-block;margin-left: 11%;"></span>
					<div style="display: inline-block;width: 15%;text-align: center;">
						<label class="appground appgoundborder" style="{{if lt .OrderInfo.Order.OrderState 20}}background-color:grey;{{end}}border-width: 2px;border-radius: 50%;padding: 9px 13px 9px 13px;">
							<i class="icon-phone icon-2x" aria-hidden="true" style="color: white;font-size: 20px;"></i>
						</label>
						<p style="font-size: 12px;">已确认</p>
					</div>
					<span style="display: inline-block;margin-left: 10%;"></span>
					<div style="display: inline-block;width: 15%;text-align: center;">
						<label class="appground appgoundborder" style="{{if lt .OrderInfo.Order.OrderState 30}}background-color:grey;{{end}}border-width: 2px;border-radius: 50%;padding: 11px 12px 7px 12px;">
							<i class="icon-plane icon-2x" aria-hidden="true" style="color: white;font-size: 20px;"></i>
						</label>
						<p style="font-size: 12px;">配送中</p>
					</div>
					<span style="display: inline-block;margin-left: 11%;"></span>
					<div style="display: inline-block;width: 15%;text-align: center;">
						<label class="appground appgoundborder" style="{{if lt .OrderInfo.Order.OrderState 40}}background-color:grey;{{end}}border-width: 2px;border-radius: 50%;padding: 8px 12px 10px 12px;">
							<i class="icon-ok icon-2x" aria-hidden="true" style="color: white;font-size: 20px;"></i>
						</label>
						<p style="font-size: 12px;">已送达</p>
					</div>
				</div>
			</div>
		</div>
		<div class="container-outter foreground" style="margin-top: 10px;">
			<div class="inner">
				<div class="oreder-info">
					<div class="container-inner" style="border-bottom: 2px solid #f7f5f5;padding: 0px 8px;">
						<label style="float: left;width: 85%;text-align: left;">配送费</label>
						<label style="float: right;width: 15%;text-align: left;">￥0.0</label>
					</div>
					{{range .OrderInfo.GoodsList}}
					<div class="container-inner" style="padding: 0px 8px;">
						<label style="float: left;width: 75%;text-align: left;">{{.GoodsName}}({{.GoodsNorms}})</label>
						<label style="float: right;width: 15%;text-align: left;">￥{{.GoodsPrice}}</label>
						<label style="float:right;font-weight:300;width: 10%;text-align: left;">x{{.GoodsNum}}</label>
					</div>
					{{end}}
					<div class="container-inner" style="border-top: 2px solid #f7f5f5;padding: 0px 8px;text-align: left;margin-top: 11px;">
						<label style="float: left;width: 85%;text-align: left;">订单总额</label>
						<label style="float: right;width: 15%;text-align: left;">￥{{.OrderInfo.Order.OrderAmount}}</label>
					</div>
				</div>
			</div>
		</div>
		<div class="container-outter foreground" style="margin-top: 8px;">
			<div class="inner">
				<!-- 配送信息 -->
				<div class="address-info" style="padding: 0px 8px;border-bottom: 2px solid #f7f5f5;">
					<div  class="container-inner" style="text-align: left;padding: 6px 0px;">
						<label style="width: 26%;margin-bottom: 0px;">
							联系方式
						</label>
						<label style="width: 68%;margin-bottom: 0px;">
							{{.Address.TrueName}} {{.Address.Mobile}}
						</label>
					</div>
					<div class="container-inner" style="text-align: left;padding: 6px 0px;">
						<label style="width: 26%;margin-bottom: 0px;">
							送货地址
						</label>
						<label style="width: 68%;margin-bottom: 0px;">
							{{.Address.LiveArea}} {{.Address.Address}}
						</label>
					</div>
				</div>
				<!-- 订单信息 -->
				<div class="address-info" style="padding: 8px 8px;">
					<div class="container-inner" style="text-align: left;padding-bottom: 6px;">
						<label style="width: 26%;margin-bottom: 0px;">
							订单编号
						</label>
						<label style="width: 68%;margin-bottom: 0px;">
							{{.OrderInfo.Order.OrderSn}}
						</label>
					</div>
					<div class="container-inner" style="text-align: left;padding-bottom: 6px;">
						<label style="width: 26%;margin-bottom: 0px;">
							下单时间
						</label>
						<label style="width: 68%;margin-bottom: 0px;">
							{{.OrderInfo.Order.AddTime}}
						</label>
					</div>
					<div class="container-inner" style="text-align: left;padding-bottom: 6px;">
						<label style="width: 26%;margin-bottom: 0px;">
							期望时间
						</label>
						<label style="width: 68%;margin-bottom: 0px;">
							{{.OrderInfo.Order.ExpectTime}}
						</label>
					</div>
					<div class="container-inner" style="text-align: left;padding-bottom: 6px;">
						<label style="width: 26%;margin-bottom: 0px;">
							收货时间
						</label>
						<label style="width: 68%;margin-bottom: 0px;">
							{{.OrderInfo.Order.FinishedTime}}
						</label>
					</div>
				</div>
			</div>
		</div>
		<div id="mod-dialog-cancel" style="display:none;">
			<div class="mask"></div>
			<div class="mask-up dialog-hide">
				<div class="mask-main">
					<div class="cancel-item" style="color: #aaa;border-radius: 5px 5px 0 0;">请选择取消原因</div>
					<table width="100%">
						<tbody>
							<tr>
								<td class="dialog-content" style="height:100px;">
									{{if not .Cancel}} 
									{{else if .Cancel.CanCancel}} 
										{{range .Cancel.CancelReason}}
											<div class="dialog-action cancel-item" data-action-arg="{{.Flag}}" style="color: #00f;">{{.Context}}</div>
										{{end}}
									{{end}}
								</td>
							</tr>
						</tbody>
					</table>
					<div class="cancel-item cancel-button" style="display:blockdisplay: block;color: red;margin-top: 8px;border-radius: 5px;">
						<div class="dialog-action dialog-hide" style="width:99%" data-action="button" data-action-arg="cancel">取消</div>
					</div>
				</div>
			</div>
		</div>
		<div id="mod-dialog-eval" style="display:none;">
			<div class="mask dialog-hide"></div>
			<div class="mask-up">
				<div class="mask-main" style="text-align: left;">
					<div class="container-inner">
						<div class="container-inner foreground" style="border-radius: 10px 10px 0px 0px;">
						<p style="font-size:13px;margin:13px 2px 20px;color:black;">
							<span style="margin-right:3px;margin-left: 2%;"> 16 度,最暖心的温度</span>
							<span>回归16度 暖暖的味道</span>
						</p>
						</div>
						<div class="container-inner foreground" style="font-size:18px;padding:10px 0px;margin:1px 0px;">
							<span id="star" value=0 style="margin-right:16px;margin-left: 2%;">整体评价</span>
							<span class="icon-star-empty appcolor star" data=0></span>
							<span class="icon-star-empty appcolor star" data=1></span>
							<span class="icon-star-empty appcolor star" data=2></span>
							<span class="icon-star-empty appcolor star" data=3></span>
							<span class="icon-star-empty appcolor star" data=4></span>
						</div>
						<div class="container-inner foreground" style="text-align: center;padding: 10px 0px;border-radius: 0px 0px 10px 10px;">
							<textarea class="background feedback" placeholder="这个小二儿怎么样？点评一下下~~" style="width: 96%;"></textarea>
							<button class="appborder appground evalOrder" style="border-width: 1px;width: 96%;border-radius: 4px;margin: 13px 0px 5px;padding: 6px;">提交</button>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div id="bottom-navbar">
			<div class="blank blanklow"></div>
			<nav class="navbar navbar-fixed-bottom foreground appgoundborder" style="border-top-width:3px;width: 100%;min-height: 35px;margin-bottom: -1%;">
				<div class="container-outter appborder appground" style="border-width: 1px;width: 88px;float: right;padding:5px;margin: 3%;border-radius:6px;">
					{{if gt .OrderInfo.Order.CancelFlag 0}}
						<span class="nocolor goshop" style="font-size: 15px;">再逛逛</span>
					{{else if lt .OrderInfo.Order.OrderState 10}}
						<span class="nocolor orderpay" style="font-size: 15px;">在线支付</span>
					{{else if lt .OrderInfo.Order.OrderState 40}}
						<span class="nocolor orderpay goshop" style="font-size: 15px;">待配送</span>
					{{else if lt .OrderInfo.Order.OrderState 50}}
						<span class="nocolor ordereval" style="font-size: 15px;">去评价</span>
					{{else}}
						<span class="nocolor goshop" style="font-size: 15px;">再逛逛</span>
					{{end}}
				</div>
				{{if eq .OrderInfo.Order.CancelFlag 0}}
					{{if lt .OrderInfo.Order.OrderState 40}}
						<div class="container-outter appborder" style="border-width: 1px;width: 88px;float: right;padding:5px;margin: 3%;border-radius:6px;">
							<span class="appcolor ordercancel" style="font-size: 15px;">取消订单</span>
						</div>
					{{end}}
				{{end}}
			</nav>
		</div>
	</div>
{{end}}
{{end}}
 
{{define "js"}}
{{with .Content}}
	<!-- 引用 goods-wrap.js -->
	<script src="/static/js/goods-wrap.js"></script>
	<script src="/static/js/order.js"></script>
	<script type="text/javascript">
		$(".ordercancel").click(function(event){
			{{if not .Cancel}}
			{{else if not .Cancel.CanCancel}} 
				alertText({{.Cancel.CancelTip.Tip}} + "</br><a href=\"tel:{{.Cancel.CancelTip.Tel}}#mp.weixin.qq.com\">{{.Cancel.CancelTip.Tel}}</a>", -1);
			{{else}}
				$("#mod-dialog-cancel").show();
			{{end}}
		});
		$(".ordereval").click(function(event){
			$("#mod-dialog-eval").show();
		});
		$(".dialog-hide").click(function(event){
			$("#mod-dialog-cancel").hide();
			$("#mod-dialog-eval").hide();
		});
		$(".dialog-action").click(function(event){
			var self = event.target;
			var cancelFlag = $(self).attr("data-action-arg");
			var orderSn = $("#order-info").attr("order-sn")
			cancelOrder(orderSn, cancelFlag);
		});
		$(".goshop").click(goShopList);

		$(".star").click(function(event){
			event.preventDefault();
			// 还原
			$(".star").removeClass("icon-star").addClass("icon-star-empty");

			var self = event.target;
			var stars = $(self).attr("data")
			for (i=0;i<=stars;i++) {
				$(".star[data=" + i + "]").removeClass("icon-star-empty").addClass("icon-star");
			}
			$("#star").attr("value", stars);
		});
		$(".evalOrder").click(function(event){
			var orderSn = $("#order-info").attr("order-sn")
			var stars = $("#star").attr("value");
			var feedback = $(".feedback").val();
			evalOrder(orderSn, stars,feedback);
		});
	</script>
{{end}}
{{end}}
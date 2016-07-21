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

<body class="background">
	{{template "content" .}}
</body>
	{{template "js" .}}
</html>

{{define "content"}}
{{with .Content}}}}
	<div id="order-list"></div>
	<script id="order-list-tmpl" type="text/x-dot-template">
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
						{{if lt .OrderInfo.Order.OrderState 0}}
							<div style="font-size: 19px;margin-top: 19px;">订单取消</div>
							<p style="font-size: 10px;">谢谢惠顾，欢迎再来</p>
						{{else if lt .OrderInfo.Order.OrderState 10}}
							<div style="font-size: 19px;margin-top: 19px;">等待付款</div>
							<p style="font-size: 10px;">待支付金额: ￥{%= it["orderInfo"]["order"]["OrderAmount"]%}</p>
							<p style="font-size: 10px; color:red">超时15分钟后，订单自动取消</p>
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
					   </div>
					</div>
				</div>
			</div>
		</div>
		<div class="container-outter foreground" style="margin-top: 10px;">
			<div class="inner">
				<div class="process-info container-inner" style="text-align: left;padding: 2px 2px 0px 9px;">
					<div style="display: inline-block;width: 15%;text-align: center;">
						<label class="appground appgoundborder" style="{{if .OrderInfo.Order.OrderState lt 10}}background-color:grey;{{end}} border-width: 2px;border-radius: 50%;padding: 10px 13px 10px 12px;">
							<i class="icon-credit-card icon-2x" aria-hidden="true" style="color: white;font-size: 16px;"></i>
						</label>
						<p style="font-size: 12px;">待支付</p>
					</div>
					<span style="display: inline-block;margin-left: 13%;"></span>
					<div style="display: inline-block;width: 15%;text-align: center;">
						<label class="appground appgoundborder" style="{{if .OrderInfo.Order.OrderState lt 20}}background-color:grey;{{end}}border-width: 2px;border-radius: 50%;padding: 9px 13px 9px 13px;">
							<i class="icon-phone icon-2x" aria-hidden="true" style="color: white;font-size: 20px;"></i>
						</label>
						<p style="font-size: 12px;">已确认</p>
					</div>
					<span style="display: inline-block;margin-left: 13%;"></span>
					<div style="display: inline-block;width: 15%;text-align: center;">
						<label class="appground appgoundborder" style="{{if .OrderInfo.Order.OrderState lt 30}}background-color:grey;{{end}}border-width: 2px;border-radius: 50%;padding: 11px 12px 7px 12px;">
							<i class="icon-plane icon-2x" aria-hidden="true" style="color: white;font-size: 20px;"></i>
						</label>
						<p style="font-size: 12px;">配送中</p>
					</div>
					<span style="display: inline-block;margin-left: 13%;"></span>
					<div style="display: inline-block;width: 15%;text-align: center;">
						<label class="appground appgoundborder" style="{{if .OrderInfo.Order.OrderState lt 40}}background-color:grey;{{end}}border-width: 2px;border-radius: 50%;padding: 8px 12px 10px 12px;">
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
							{{.Address.TrueName}} {{.Address.MobPhone}}
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
							{{.OrderInfo.Order.OrderId}}
						</label>
					</div>
					<div class="container-inner" style="text-align: left;padding-bottom: 6px;">
						<label style="width: 26%;margin-bottom: 0px;">
							下单时间
						</label>
						<label style="width: 68%;margin-bottom: 0px;">
							{{.OrderInfo.Order.OrderTime}}
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
		<div id="bottom-navbar">
			<div class="blank blanklow"></div>
			<nav class="navbar navbar-fixed-bottom foreground" style="border:none;width: 100%;min-height: 35px;margin-bottom: -1%;">
				<div class="container-outter appground" style="width: 110px;float: right;padding: 3%;">
				   <span style="color: white;font-size: 15px;">取消订单</span>
				</div>
			</nav>
		</div>
	</script>
{{end}}
{{end}}

{{define "js"}}
	<!-- 引用 goods-wrap.js -->
	<script src="/static/js/goods-wrap.js"></script>
	<script src="/static/js/order.js"></script>
{{end}}
<!-- 引入主框架 -->
{{template "layout" .}}

{{define "content"}}
	<div id="order-list"> </div>
	<script id="order-list-tmpl" type="text/x-dot-template">
		{% for(var orderIdx in it["list"]) { %}
		<div class="container-outter foreground" style="margin-top: 2px;">
			<div class="inner">
				<div class="oreder-info" style="text-align: left;padding: 6px 0px;">
					<div class="container-inner appgoundborder" style="border-bottom-width: 2px;padding: 0px 8px;">
						<label style="float: left;color: black;font-weight: 300;">{%= it.list[orderIdx]["orderInfo"]["order"]["OrderTime"]%}</label>
						<label class="appcolor" style="float: right;font-weight: 100;">
							{%? it.list[orderIdx]["orderInfo"]["order"]["CancelFlag"] == 1 %}
								订单取消
							{%?? it.list[orderIdx]["orderInfo"]["order"]["CancelFlag"] == 2 %}
								取消失败
							{%?? it.list[orderIdx]["orderInfo"]["order"]["OrderState"] < 10 %}
								等待付款
							{%?? it.list[orderIdx]["orderInfo"]["order"]["OrderState"] < 20 %}
								支付成功
							{%?? it.list[orderIdx]["orderInfo"]["order"]["OrderState"] < 30 %}
								已经确认
							{%?? it.list[orderIdx]["orderInfo"]["order"]["OrderState"] < 40 %}
								正在配送
							{%?? it.list[orderIdx]["orderInfo"]["order"]["OrderState"] >= 40 %}
								订单送达
							{%?%}
						</label>
					</div>
					<div class="container-inner" style="border-bottom: 2px solid #f7f5f5;padding: 8px 8px;">
						<div style="display: inline-block;width: 100%;white-space: nowrap;overflow: scroll;">
							{%~ it.list[orderIdx]["orderInfo"]["goodsList"] :value:idx %}
							<div class="container-inner" style="padding: 0px 8px;width: 30%;">
								<img style="width: 66px;" src="{%= value["GoodsImage"]%}">
							</div>
							{%~%}
						</div>
					</div>
					<div class="container-inner" style="padding: 0px 8px;">
						<label style="float: left;padding: 8px 0px 0px;font-weight: 300;font-size: 13px;">
							共{%= it.list[orderIdx]["orderInfo"]["goodsList"].length %}件商品 ￥{%= it.list[orderIdx]["orderInfo"]["order"].OrderAmount %}
						</label>
						<label class="appground" style="float: right;padding: 3px 8px;margin: 5px 0px;font-weight: 300;border-radius: 7px;font-size: 12px;">
							{%? it.list[orderIdx]["orderInfo"]["order"]["CancelFlag"] == 1 %}
								<a style="color: black;border-bottom: 1px solid black;" href="/shop/list">再逛逛</a>
							{%?? it.list[orderIdx]["orderInfo"]["order"]["CancelFlag"] == 2 %}
								<a style="color: black;border-bottom: 1px solid black;" href="/static/feedback.html">去投诉</a>
							{%?? it.list[orderIdx]["orderInfo"]["order"]["OrderState"] < 10 %}
								<a style="color: black;border-bottom: 1px solid black;" href="/static/orderInfo.html">去支付</a>
							{%?? it.list[orderIdx]["orderInfo"]["order"]["OrderState"] < 40 %}
								<a style="color: black;border-bottom: 1px solid black;" href="/static/orderInfo.html">待配送</a>
							{%?? it.list[orderIdx]["orderInfo"]["order"]["OrderState"] < 50 %}
								<a style="color: black;border-bottom: 1px solid black;" href="/shop/list">去评价</a>
							{%?? it.list[orderIdx]["orderInfo"]["order"]["OrderState"] >= 50 %}
								<a style="color: black;border-bottom: 1px solid black;" href="/static/orderInfo.html">已评价</a>
							{%?%}
						</label>
					</div>
				</div>
			</div>
		</div>
		{%}%}
	</script>
	<div id="hasMore"></div>
{{end}}

{{define "js"}}
	<!-- 引用 goods-wrap.js -->
	<script src="/static/js/goods-wrap.js"></script>
	<script src="/static/js/order.js"></script>
	<script>var globalContext = {{.GlobalContext}}; var globalCart={};</script>
	<script type="text/javascript">
		$("#foot-menu").find(".active").removeClass("active");
		$("#foot-menu").find("#user").addClass("active");
		getOrderList("",false);
	</script>
	<script type="text/javascript">
		// 监听滚动
		$(window).scroll(function(event){
			if ("none" == $("#hasMore").css("display")) {
				return;
			}
			// 快到下方时自动加载
			if($(window).height()*1.3+$(window).scrollTop() > $("#hasMore").offset().top) {
				getOrderList("",false);
			}
		});
	</script>
{{end}}
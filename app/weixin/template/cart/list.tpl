{{define "content"}}
	<!-- <div goods-list> -->
	<div id="goods-list" class="goods-list cart-goodslist"> </div>
	<script id="goods-list-tmpl" type="text/x-dot-template">
		<div class="cat-title"> 
			<div class="tip-outter"> 
				<div class="tip-inner"> 
					<p style="padding:3px;"> 
						<span id="cart-tip" style="font-size:10px;font-family:行楷;">{%=it.tips%}</span>
					</p>
				</div>
			</div>
			<div class="cart" style="position: absolute;top: 18%;right: 8%;"> 
				<i class="icon-trash" data-toggle="modal" data-target="#notice"></i>
			</div>
		</div>

		{% for(var goodsId in it["map"]) { %}
		<div class="goods"> 
			<div class="check" goods-id="{%=goodsId%}" style="background-color:{%? it["map"][goodsId]["selected"] == 0%}white{%??%}#40bb91{%?%}"> 
				<i class="icon-ok" aria-hidden="true"></i> 
			</div> 
			<div style="width: 25%;display: inline-block;"> 
				<img src="{%=globalContext[goodsId].Image%}" style="width: 100%;height: 80px;"> 
			</div>
			<div class="inner"> 
				<p> 
					<span style="font-size: 14px;">{%=globalContext[goodsId].Name%}</span>
					<span style="font-size: 14px;">{%=globalContext[goodsId].Norms%}</span>
				</p>
				<p style="padding-top:16px;"> 
					<span class="price">￥{%=globalContext[goodsId].Price%}</span>
					<span><del>￥{%=globalContext[goodsId].Marketprice%}</del></span> 
				</p>
			</div> 
			<div class="item-edit">
				<div class="item-cart item-cart-click" value="-1" style="bottom: 30px;right: 58px;" goods-id="{%=goodsId%}"> 
					<div class="btn -purchase"> 
						<i class="icon-minus"></i> 
					</div> 
				</div> 
				<div class="item-cart" style="bottom: 29px;right: 39px;" goods-id="{%=goodsId%}"> 
					<i class="count"><span goods-id="{%=goodsId%}">{%=it['map'][goodsId]['goodsNum']%}</span></i> 
				</div> 
				<div class="item-cart item-cart-click" value="1" style="bottom: 30px;right: 8px;" goods-id="{%=goodsId%}"> 
					<div class="btn -purchase"> 
						<i class="icon-plus"></i> 
					</div> 
				</div>
			</div>
		</div> 
		{% } %}
	</script>
	<!-- </div> -->
	<div class="goods-list cart-info cart-goodslist"> 
		<div class="goods" style="margin-bottom:0px;"> 
			<div class="check" goods-id="all"> 
				<i class="icon-ok" aria-hidden="true"></i> 
			</div> 
			<div class="info-inner"> 
				<p>
					<span style="font-size: 13px;color: grey;padding-left: 7px;">全选</span>
					<span style="font-size: 15px;padding-left: 10px;">总计:</span>
					<span id="cart-cost" style="font-size: 13px;color: red;font-weight: 500;">￥0</span>
				</p>
			</div>
			<div class="info-button">
				<span style="color: white;font-size: 15px;">结算</span>
			</div>
		</div>
	</div>
	<div class="modal fade" id="notice" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
		<div class="modal-dialog" role="document" style="width: 68%;background-color: white;text-align: center;margin-left: 16%;margin-top: 51%;">
			<div class="" style="display: inline-block;padding: 9%;">
				<span class="modal-body">是否删除所有商品</span>
			</div>
			<div style="display: inline-block;width: 100%;text-align: center;border-top: 1px solid #8a8686;">
				<div style="display: inline-block;width: 47%;color: #2281ec;border-right: 1px solid #8a8686;">
					<div style="margin-top: 8px;margin-bottom: 6px;" class="clear-cart" data-dismiss="modal">确定</div>
				</div>
				<div style="display: inline-block;width: 47%;color: #2281ec;">
					<div style="margin-top: 8px;margin-bottom: 6px;" data-dismiss="modal">取消</div>
				</div>
			</div>
		</div>
	</div>
	<div class="modal fade" id="alertModel" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
		<div class="modal-dialog" role="document" style="top: 30%;position: absolute;width: 80;left: 10%;width: 80%;">
			<div class="modal-content">
				<!-- <div style="padding: 16px;position: relative;">
					<i type="button" class="close" data-dismiss="modal" aria-label="Close" style="position: absolute;top: 5px;right: 10px;">
						<span aria-hidden="true">×</span>
					</i>
				</div> -->
				<div class="modal-body" style="text-align: center;">
					<span class="alertText" style="color:red;"></span>
				</div>
			</div>
		</div>
	</div>
{{end}}
{{define "js"}}
	<!-- 引用 goods-wrap.js -->
	<script src="/static/js/goods-wrap.js"></script>
	<script src="/static/js/cart.js"></script>
	<script>var globalContext = {{.GlobalContext}}; var globalCart={};</script>
	<script type="text/javascript">
		doT.templateSettings = {
			evaluate:	/\{\%([\s\S]+?)\%\}/g,
			interpolate:/\{\%=([\s\S]+?)\%\}/g,
			encode:		/\{\%!([\s\S]+?)\%\}/g,
			use:		/\{\%#([\s\S]+?)\%\}/g,
			define:		/\{\%##\s*([\w\.$]+)\s*(\:|=)([\s\S]+?)#\%\}/g,
			conditional:/\{\%\?(\?)?\s*([\s\S]*?)\s*\%\}/g,
			iterate:	/\{\%~\s*(?:\}\}|([\s\S]+?)\s*\:\s*([\w$]+)\s*(?:\:\s*([\w$]+))?\s*\%\})/g,
			varname: 'it',
			strip: true,
			append: true,
			selfcontained: false
		};

		$("#foot-menu").find(".active").removeClass("active");
		$("#foot-menu").find("#cart").addClass("active");
		$(".clear-cart").click(cleanCart);
		registerCart();
	</script>
{{end}}
{{define "content"}}
	<div class="cat-title"> 
		<div class="cat-title-inner"> 
			<p style="padding-top:8px;"> 
				<span>购物车</span>
			</p>
		</div>
		<div class="cart" style="position: absolute;top: 18%;right: 8%;"> 
			<i class="icon-trash" style="font-size: 19px;"></i> 
		</div>
	</div>

	<!-- <div goods-list> -->
	<div id="goods-list" class="goods-list cart-goodslist"> </div>
	<script id="goods-list-tmpl" type="text/x-dot-template">
		<div class="tip-outter"> 
			<div class="tip-inner"> 
				<p style="padding:3px;"> 
					<span id="cart-tip" style="font-size:10px;font-family:行楷;">{%=it.tips%}</span>
				</p>
			</div>
		</div>
		{% for(var goodsId in it["map"]) { %}
		<!-- <div>{%=goodsId%}{%=it[goodsId]%}</div> -->
		<div class="goods"> 
			<div class="checked"> 
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
				<div class="item-cart" value="-1" style="bottom: 30px;right: 58px;" goods-id="{%=goodsId%}"> 
					<div class="btn -purchase"> 
						<i class="icon-minus"></i> 
					</div> 
				</div> 
				<div class="item-cart" style="bottom: 29px;right: 39px;" goods-id="{%=goodsId%}"> 
					<i class="count"><span goods-id="{%=goodsId%}">{%=it['map'][goodsId]['goodsNum']%}</span></i> 
				</div> 
				<div class="item-cart" value="1" style="bottom: 30px;right: 8px;" goods-id="{%=goodsId%}"> 
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
			<div class="checked"> 
				<i class="icon-ok" aria-hidden="true"></i> 
			</div> 
			<div class="info-inner"> 
				<p>
					<span style="font-size: 15px;color: grey;">全选</span>
					<span style="font-size: 12px;padding-left: 8px;">总计:</span>
					<span id="cart-cost" style="font-size: 12px;color: red;">￥39.45</span>
				</p>
			</div>
			<div class="info-button">
				<span style="color: white;font-size: 15px;">结算</span>
			</div>
		</div>
	</div>
{{end}}
{{define "js"}}
	<!-- 引用 goods-wrap.js -->
	<script src="/static/js/goods-wrap.js"></script>
	<script src="/static/js/cart.js"></script>
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
	</script>
{{end}}
<!-- 引入主框架 -->
{{template "layout" .}}

{{define "content"}}
{{$globalContext := .GlobalContext}}

{{with .Content}}
{{with .Banner}}

<!-- 上方轮播 banner -->
<script type="text/javascript">
	$(function(){
		$('#carousel-banner').carousel({
			interval: 3000
		});
	});
</script>
<div id="carousel-banner" class="carousel-inner">
	{{range $idx,$info := .}}
	<div class="item {{if eq $idx 0}} active {{end}}">
		<img src="{{$info.Icon}}"/>
	</div>
	{{end}}
	<!-- 轮播（Carousel）指标 -->
	<div id="indicators-banner">
		<ol class="carousel-indicators">
			{{range $idx,$info := .}}
			<li data-target="#carousel-banner" data-slide-to="{{$idx}}" {{if eq $idx 0}} class="active" {{end}}></li>
			{{end}}
		</ol>	
	</div>
</div>
{{end}}

<!-- 上方导航栏 -->
<div class="up-navbar container-outter foreground">
	<div class="flexbox inner">
		{{range .Nav}}
		<a href="{{.Url}}">
			<img src="{{.Icon}}"/>
			<p>{{.Name}}</p>
		</a>
		{{end}}
	</div>
</div>

<!--商品列表 --> 
<div class="goods-list">
	{{range .Class}}
	<div class="container-outter foreground">
		<div class="inner">
			<div class="desc" style="border-color:#{{.Color}};">
				<span style="float:left;color:#{{.Color}}">{{.Name}}</span>
				<a href="XXXX" style="float:right;">更多>></a>
			</div>
			<img class="banner" src="{{.Img}}"/>
			<div class="goods-list background" style="display:inline-block;">
				{{range .GoodsList}}
				<div class="goods">
					<div class="cart-count-flag"> 
						<div class="item-cart item-cart-click" goods-id="{{.Id}}" style="display:none"> 
							<i class="count"><span goods-id="{{.Id}}">0</span></i> 
						</div>
					</div>
					<a data-toggle="modal" data-target="#detailModelOutter" data-whatever="{{addGlobalContext $globalContext .Id .}}">
						<img src="{{.Image}}"/>
					</a>
					<div style="padding-top:10px;text-align:left">
						<p>{{.Name}}</p>
						<p>{{.Norms}}</p>
						<p>
							<span class="price">￥{{.Price}}</span>
							<span><del>￥{{.Marketprice}}</del></span>
						</p>
					</div> 
					<div class="item-edit">
						<div class="item-cart item-cart-click" goods-id="{{.Id}}" value=1>
							<div class="btn -purchase">
								<i class="icon-plus"></i>
							</div>
						</div>
					</div>
				</div>
				{{end}}
			</div>
		</div>
	</div>
	{{end}}
</div>

<!-- 导购区 --><!-- 
<div id="guider-wrap" class="container-outter">
	<div style="width:100%;display:inline-block;text-align: center;">
			<span style="color:#7d7c78;font-weight:700;">--------- 小店热卖 ---------</span>
	</div>
</div> -->
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
	// $("#foot-menu").find(".active").removeClass("active");
	// $("#foot-menu").find("#home").addClass("active");
</script>
<!-- <script src="js/goods-wrap.js"></script> -->
<script src="/static/js/goods-wrap.js"></script>
<script>var globalContext = {{.GlobalContext}}; var globalCart=getCart();initCart();</script>
{{end}}

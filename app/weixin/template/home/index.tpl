{{define "content"}}
<!-- 上方轮播 banner -->
{{with .content}}
<div id="carousel-banner" class="carousel-inner">
	{{range $idx,$info := .banner}}
	<div {{if 0 eq $idx}} class="item active" {{end}}>
		<img src="{{$info.icon}}"/>
	</div>
	{{end}}
	<!-- 轮播（Carousel）指标 -->
	<div id="indicators-banner">
		<ol class="carousel-indicators">
			{{range $idx,$info := .}}
			<li data-target="#carousel-banner" data-slide-to="{{$idx}}" {{if 0 eq $idx}} class="item active" {{end}}></li>
			{{end}}
		</ol>	
	</div>
</div>
<script type="text/javascript">
	$(function(){
		$('#carousel-banner').carousel({
			interval: 2000
		});
	});
</script>

<!-- 上方导航栏 -->
<div class="up-navbar container-outter">
	<div class="flexbox inner">
		{{range .nav}}
		<a href="{{.url}}">
			<img src="{{.icon}}"/>
			<p>{{.name}}</p>
		</a>
		{{end}}
	</div>
</div>

<!--商品列表 --> 
<div class="shop-class">
	{{range .class}}
	<div class="container-outter">
		<div class="inner">
			<div class="desc" style="border-color:{{.color}};">
				<span style="float:left;color:{{.color}}">{{.name}}</span>
				<a href="{{.url}}" style="float:right;">更多>></a>
			</div>
			<img class="banner" src="{{.img}}"/>
			<div class="goods-list" style="background-color:#EFEFEF;display:inline-block;">
				{{range .goodList}}
				<div class="goods">
					<a href="{{.goodsInfo}}">
						<img src="{{.goodsImg}}"/>
					</a>
					<div style="padding-top:10px;text-align:left">
						<p>{{.goodsName}}</p>
						<p>{{.goodsDesc}}</p>
						<p>
							<span class="price">￥{{.goodsCost}}</span>
							<span><del>￥{{.goodsPrice}}</del></span>
						</p>
					</div>
					<div class="item-add">
						<div class="btn -plus">
							<i class="icon-plus"></i>
						</div>
					</div>
				</div>
				{{end}}
			</div>
		</div>
	</div>
	{{end}}
</div>

<!-- 导购区 -->
<div id="guider-wrap" class="container-outter">
	<div style="width:100%;display:inline-block;text-align: center;">
			<span style="color:#7d7c78;font-weight:700;">--------- 小店热卖 ---------</span>
	</div>
	{{range .guider}}
	<div class="goods-wrap inner"">
		<div class="goods-container">
			<div class="goods-label "></div>
			<div class="goods-count" style="display:none">0</div>
			<div class="goods-detail">
				<div class="goods-image-wrap">
					<img src="{{.goodsImg}}" alt="商品图片">
				</div>
				<div class="goods-title">{{.goodsName}}</div>
				<div class="goods-price">
					<span class="price-main">¥ {{.goodsCost}}</span>
					<span class="price-unit">/{{.goodsUnit}}</span>
				</div>
				<div class="item-add">
					<div class="btn -plus">
						<i class="icon-plus"></i>
					</div>
				</div>
			</div>
		</div>
	</div>
	{{end}}
</div>
{{end}}
{{end}}

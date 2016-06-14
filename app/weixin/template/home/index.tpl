{{define "content"}}
<!-- 上方轮播 banner -->
{{with .Content}}
{{with .Banner}}
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
		{{range .Nav}}
		<a href="{{.Url}}">
			<img src="{{.Icon}}"/>
			<p>{{.Name}}</p>
		</a>
		{{end}}
	</div>
</div>

<!--商品列表 --> 
<div class="shop-class">
	{{range .Class}}
	<div class="container-outter">
		<div class="inner">
			<div class="desc" style="border-color:XX;">
				<span style="float:left;color:XXX">{{.Name}}</span>
				<a href="XXXX" style="float:right;">更多>></a>
			</div>
			<img class="banner" src="{{.Img}}"/>
			<div class="goods-list" style="background-color:#EFEFEF;display:inline-block;">
				{{range .GoodsList}}
				<div class="goods">
					<a href="XXXX">
						<img src="{{.Image}}"/>
					</a>
					<div style="padding-top:10px;text-align:left">
						<p>{{.Name}}</p>
						<p>{{.Desc}}</p>
						<p>
							<span class="price">￥{{.Price}}</span>
							<span><del>￥{{.Marketprice}}</del></span>
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

<!-- 导购区 --><!-- 
<div id="guider-wrap" class="container-outter">
	<div style="width:100%;display:inline-block;text-align: center;">
			<span style="color:#7d7c78;font-weight:700;">--------- 小店热卖 ---------</span>
	</div>
</div> -->
{{end}}
{{end}}

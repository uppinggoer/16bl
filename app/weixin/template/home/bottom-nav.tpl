{{define "bottom-navbar"}}
	<!-- 下方导航栏 -->
	<div id="bottom-navbar">
		<div style="height: 26px;text-align:center;background-color:#EFEFEF;">
			<span>
				<a href="/" style="color:grey;">查看全部商品 >></a>
			</span>
		</div>
		<hr style="height: 18px;background-color:#EFEFEF;"></hr>
		<nav class="navbar navbar-fixed-bottom" style="background-color:white;">
			<ul id="foot-menu" class="nav nav-tabs">
				{{range .FootMenu}}
				<li id="{{.Id}}" {{if .Activite}}class="active"{{end}}>
					{{if eq "cart" .Id}}
					<div class="cart-count-flag"> 
						<div class="item-cart" style="display: none; padding: 0px 7px; position: absolute; top: 0px; right: 0px;"> 
							<i class="count"><span>0</span></i> 
						</div>
					</div>
					{{end}}
					<a trigger="{{.Trigger}}" href="{{.Url}}" class="trigger">
						<i class="{{.Icon}} icon-2x" aria-hidden="true"></i>
						<p>{{.Name}}</p>
						<!-- <span class="icon-home icon-2x" aria-hidden="true"></span> -->
					</a>
				</li>
				{{end}}
			</ul>
		</nav>
	</div>
{{end}}
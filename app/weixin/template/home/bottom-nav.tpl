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
					<a href="{{.Url}}">
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
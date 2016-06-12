{{define "bottom-navbar"}}
	<!-- 下方导航栏 -->
	<div id="bottom-navbar">
		<div style="height: 28px;text-align:center">
			<span><a href="/" style="color: grey;">查看全部商品 >></a></span>
		</div>
		<div style="height: 36px;">
		</div>
		<nav class="navbar navbar-fixed-bottom" style="background-color:white;">
			<ul id="foot-menu" class="nav nav-tabs">
				{{range .footMenu}}
				<li {{if .active}}class="active"{{end}}>
					<a href={{.url}} style="height: 32px;text-align: center;">
						<i class="icon-home icon-2x" aria-hidden="true"></i>
						<p>{{.name}}</p>
						<!-- <span class="icon-home icon-2x" aria-hidden="true"></span> -->
					</a>
				</li>
				{{end}}
			</ul>
		</nav>
	</div>
{{end}}
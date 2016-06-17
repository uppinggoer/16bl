{{define "bottom-navbar"}}
	<!-- 下方导航栏 -->
	<div id="bottom-navbar">
		<div class="blank"></div>
		<nav class="navbar navbar-fixed-bottom" style="border:none;background-color:white;">
			<ul id="foot-menu" class="nav nav-tabs">
				{{range .FootMenu}}
				<li id="{{.Id}}" {{if .Activite}}class="active"{{end}}>
					{{if eq "cart" .Id}}
					<div class="cart-count-flag"> 
						<div class="item-cart" style="display:none;padding:0px 7px; position: absolute; top: 0px; right: 0px;z-index: 6;"> 
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
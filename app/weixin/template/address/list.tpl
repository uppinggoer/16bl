<!-- 引入主框架 -->
<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<!-- <title>便利</title> -->
	<title>{{.Title}}</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" />
	<meta charset="utf-8">
	<!-- <link rel="shortcut icon" href="/static/img/go.ico"> -->
	<link rel="stylesheet" href="http://cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
	<!-- 可选的Bootstrap主题文件（一般不用引入） -->
	<link rel="stylesheet" href="http://cdn.bootcss.com/font-awesome/3.0.2/css/font-awesome.css">

	<link rel="stylesheet" href="/static/css/goods-wrap.css">
	<link rel="stylesheet" href="/static/css/home.css">
	<link rel="stylesheet" href="/static/css/order.css"> 
	
	<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
	<script src="http://cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
	<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
	<script src="http://cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
	<script src="http://cdn.bootcss.com/dot/1.0.3/doT.min.js"></script> 
</head>

<body class="background">
	{{template "content" .}}
</body>
	{{template "js" .}}
</html>

{{define "content"}}
{{$globalContext := .GlobalContext}}
{{with .Content}}
	<div id="address-list">
		{{range .}}
			<div class="container-outter foreground" id="{{.Id}}" style="margin-top: 2px;">
				<div class="inner" style="padding-top: 0px;">
					<div class="address-info" style="text-align: left;margin-left: 8px;">
						<div style="width: 6%;text-align: center;display: inline-block;margin: 6% 2% 0% 0%;vertical-align: top;">
							<div add-id="{{addGlobalContext $globalContext .Id .}}" class="check">
								<i class="icon-ok"></i>
							</div>
						</div>
						<div class="modify" add-id="{{.Id}}" style="width: 6%;display: inline-block;margin:6% 0% 8% 2%;float: right;">
							<i class="icon-edit"></i>
						</div>
						<div style="width: 81%;display: inline-block;float: right;">
							<label data-type="" data-value="Id" style="display: none;">{{.Id}}</label>
							<div class="container-inner" style="text-align: left;padding: 6px 0px;">
								<label style="width: 18%;margin-bottom: 0px;">
									{{.TrueName}}
								</label>
								<label style="width: 68%;margin-bottom: 0px;">
									{{.Mobile}}
								</label>
							</div>
							<div class="container-inner" style="text-align: left;padding: 6px 0px;">
								<label style="width: 18%;margin-bottom: 0px;">
									{{.LiveArea}}
								</label>
								<label style="width: 68%;margin-bottom: 0px;">
									{{.Address}}
								</label>
							</div>
						</div>
					</div>
				</div>
			</div>
		{{end}}
	</div>
	<div class="container-outter foreground new-address" style="margin-top: 18px;">
		<div class="appcolor check" style="display: inline-block;line-height: 36px;font-size: 23px;vertical-align: middle;margin-right: 6px;">
			<i class="icon-plus-sign"></i>
		</div>
		<div style="display: inline-block;line-height: 36px;vertical-align: middle;">
			<span>添加一个收货地址</span>
		</div>
	</div>

	<div id="address-info-dialog" style="display:none;">
		<div class="mask"></div>
		<div class="mask-up dialog-hide">
			<div class="mask-main background" style="text-align:left;bottom:39%;">
				<div id="address-modify-dialog" class="container-inner" style="text-align: center;"></div>
				</div>
			</div>
		</div>
	</div>
	<script id="address-modify-tmp" type="text/x-dot-template">
		<div class="foreground" style="line-height: 30px;">
			<div class="dialog-hide container-outter appcolor" style="width:auto;float: left;margin-left: 12px;">
				<span> << </span>
			</div>
			<div class="save-address container-outter appcolor" style="width:auto;float: right;margin-right: 9px;">
				<span> 保存 </span>
			</div>
			<div class="save-address container-outter" style="width:auto;">
				<span> {%= it["Type"] %} </span>
			</div>
		</div>
		<div class="container-outter foreground" style="margin: 10px 0px;">
			<div class="inner">
				<div class="address-info" style="padding: 0px 8px;">
					<div class="container-inner" style="text-align: left;padding-bottom: 6px;">
						<label style="width: 18%;margin-bottom: 0px;">
							收货人
						</label>
						<input type="text" name="true_name" placeholder="请填写收货人的姓名" value="{%= it["TrueName"] %}"/>
					</div>
					<div class="container-inner" style="text-align: left;padding-bottom: 6px;">
						<label style="width: 18%;margin-bottom: 0px;">
							手机号
						</label>
						<input type="tel" name="mob_phone" placeholder="请填写收货人的电话" value="{%=it["Mobile"]%}"/>
					</div>
					<div class="container-inner" style="text-align: left;padding-bottom: 6px;">
						<label style="width: 18%;margin-bottom: 0px;">
							位置
						</label>
						<input name="live_area" placeholder="{%= it["LiveArea"] %}">
					</div>
					<div class="container-inner" style="text-align: left;padding-bottom: 6px;">
						<label style="width: 18%;margin-bottom: 0px;">
							地址
						</label>
						<input name="building_info" placeholder="{%= it["Address"] %}">
					</div>
				</div>
			</div>
		</div>
		{%? it["Del"] %}
			<div class="del-address container-outter foreground" style="line-height: 33px;color: red;">
				<span> 删除该地址 </span>
			</div>
		{%?%}
	</script>
{{end}}	
{{end}}

{{define "js"}}
	<script src="/static/js/goods-wrap.js"></script>
	<script src="/static/js/order.js"></script>
	<script>var globalContext = {{.GlobalContext}};</script>
	<script type="text/javascript">
		$(".address-list .address-info").click(function(event){
			$(".address-list .check").removeClass("appground");

			var self=this;
			$(self).find(".check").addClass("appground");

			// modify addressList
			// arryAll.forEach(function(e){  
			//     alert(e);  
			// }) 
		});
		$(".modify").click(function(event){
			var self=this;
			var addressId = $(self).attr("add-id")

			var addressData = globalContext[addressId];
			addressData['Type'] = "修改收货地址";
			addressData['Del'] = true;

			renderAddress(addressData);
			init(event);
		});
		$(".new-address").click(function(event){
			var addressData = {
				Id: 0,
				TrueName: "",
				Gender: 0,
				LiveArea: "",
				Address: "",
				Mobile: "",
				Type: "添加收货地址",
				Del: false
			};
			renderAddress(addressData);
			init(event);
		});

		function init(target) {
			$("#address-info-dialog").show();
			$(".dialog-hide").click(function(){
				$("#address-info-dialog").hide();
			});

			// 提交修改
			$(".save-address").click(function(event){
				console.log("XX");
			});
		}
	</script>
{{end}}
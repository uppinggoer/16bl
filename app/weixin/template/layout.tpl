<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<!-- <title>便利</title> -->
	<title>便利</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" />
	<meta charset="utf-8">
	<!-- <link rel="shortcut icon" href="/static/img/go.ico"> -->
	<link rel="stylesheet" href="http://cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
	<!-- 可选的Bootstrap主题文件（一般不用引入） -->
	<link rel="stylesheet" href="http://cdn.bootcss.com/font-awesome/3.0.2/css/font-awesome.css">

	<link rel="stylesheet" href="/static/css/goods-wrap.css">
	<link rel="stylesheet" href="/static/css/home.css">
	<link rel="stylesheet" href="/static/css/cart.css"> 
	
	<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
	<script src="http://cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
	<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
	<script src="http://cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
	<script src="http://cdn.bootcss.com/dot/1.0.3/doT.min.js"></script> 
</head>
<body>
	{{template "content" .}}
	{{template "bottom-navbar" .}}
	<div id="cart-ball"></div>
</body>
	{{template "js" .}}
	<script type="text/javascript">
		$(".trigger").click(function(e) {
			var self = $(this);
			var trigger = self.attr("trigger");
			if (undefined != trigger && "" != trigger) {
				e.preventDefault();
				eval(trigger)();
			}
		});
		$(".item-cart").click(editCart);
	</script>
	<script>var globalContext = {{.GlobalContext}}; var globalCart=getCart()</script>
</html>
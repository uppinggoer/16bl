<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<!-- <title>便利</title> -->
	<title>便利</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" />
	<meta charset="utf-8">
	<!-- <link rel="shortcut icon" href="/static/img/go.ico"> -->
	<!-- {{template "seo" .}} -->
	<meta name="author" content="polaris <polaris@studygolang.com>">
	<link rel="stylesheet" href="http://cdn.bootcss.com/bootstrap/3.3.5/css/bootstrap.min.css">
	<!-- 可选的Bootstrap主题文件（一般不用引入） -->
	<link rel="stylesheet" href="http://cdn.bootcss.com/font-awesome/3.0.2/css/font-awesome.css">
	{{template "css" .}}
</head>
<body>
	{{template "content" .}}
	{{template "bottom-navbar" .}}
</body>
	<!-- 新 Bootstrap 核心 CSS 文件 -->
	<!-- jQuery文件。务必在bootstrap.min.js 之前引入 -->
	<script src="http://cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
	<!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
	<script src="http://cdn.bootcss.com/bootstrap/3.3.5/js/bootstrap.min.js"></script>
	{{template "js" .}}
</html>
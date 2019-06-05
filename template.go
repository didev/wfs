package main

const headHTML = `<!DOCTYPE html><head><title>WFS</title>
	<meta charset="utf-8">
	<style>
	@charset "utf-8";
	html, body {
		margin:5px;
		padding:5px;
		background:#CCCDCC;
		font-family: Courier, 'Courier New', Courier_New;
	}
	a:link {text-decoration:none; color:#000000;}
	a:visited {text-decoration:none; color:#000000;}
	a:active {text-decoration:none; color:#000000;}
	a:hover {text-decoration:none; color:#000000; cursor:pointer;}
	a img { border: none; }
	img { vertical-align: middle; border: none; margin-bottom: 2px;}

	rootpath {
		background: #479EF8;
		padding: 2px;
		border-style : solid;
		border-width : 1px;
		border-color : #8A8A8A;
		border-radius : 4px;
		opacity:0.9;
	}
	</style>
	<link rel="icon" type="image/png" href="http://10.0.90.193:8080/template/icon/wfs.png">
	</head><body>`

const rootHTML = `<!DOCTYPE html><head><title>WFS</title>
	<meta charset="utf-8">
	<style>
	@charset "utf-8";
	html, body {
		margin:5px;
		padding:5px;
		background:#DCDEE0;
		font-family: NanumGothic, Helvetica, sans_serif, Verdana;
	}

	a:link {text-decoration:none; }
	a:visited {text-decoration:none; }
	a:active {text-decoration:none; }
	a:hover {text-decoration:none; cursor:pointer;}


	rootpath {
		background: #479EF8;
		padding: 10px;
		border-color : #8A8A8A;
		border-radius : 8px;
		color: #FFFFFF;
		font-size: 20px;
		font-weight: 900;
		width: 300px;
		text-align: center;
		letter-spacing: 2px;
	}

	t1 {
		color: #444444;
		font-size: 40px;
		font-weight: bold;
		text-align: center;
		letter-spacing: 10px;
	}

	t2 {
		color: #444444;
		font-size: 15px;
		text-align: center;
	}
	
	t3 {
		color: #444444;
		font-size: 20px;
		font-weight: bold;
		text-align: center;
		letter-spacing: 2px;
	}

	t4 {
		color: #444444;
		font-size: 18px;
		text-align: center;
	}

	</style>
	<link rel="icon" type="image/png" href="http://10.0.90.193:8080/template/icon/wfs.png">
	</head><body>`

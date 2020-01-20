<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>homeworkSite</title>
	<link rel="stylesheet" href="style.css"> 
</head>
<body>

<div class="content">
	<?php
		include "menu.php";
	?>

	<h1>Домашнее задание по основам программирования<br>от портала GeekBrains</h1>

	<div class="center">
	<img src="img/avatar.jpg">
		<div class="box_text">
			<p><b>Добрый день!</b>.</p>
				
			<p>Меня зовут <i>Шепелев Сергей</i>. Представляю вашему вниманию решение домашнего задания по основам программирования.</p>

			<p>Задание взято с IT-портала <a href="https://geekbrains.ru">GeekBrains</a></p>

			<p>В данный момент увлекаюсь разработкой веб сервисов на Golang</p>
		</div>
	</div>
</div>
<div class="footer">
	Copyright &copy; <?php echo date("Y");?> Sergey Shepelev
<div>


</body>
</html>
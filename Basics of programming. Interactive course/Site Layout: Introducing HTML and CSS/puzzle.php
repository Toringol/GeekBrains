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

<div class="contentWrap">
    <div class="content">
        <div class="center">

			<h1>Игра в загадки</h1>

			<div class="box">

				<?php

					if (isset($_GET["userAnswer1"]) && isset($_GET["userAnswer2"]) && isset($_GET["userAnswer3"])) {
						$firstRightAnswer = [
							["лампочка"],
							["сапоги"],
							["магнитофон", "диктофон"],
						];
							
						$score = 0;

						$userAnswers = [
							strtolower($_GET["userAnswer1"]),
							strtolower($_GET["userAnswer2"]),
							strtolower($_GET["userAnswer3"]),
						];


						for ($i = 0; $i < count($userAnswers); $i++) {
							for ($j = 0; $j < count($firstRightAnswer[$i]); $j++) {
								if ($userAnswers[$i] == $firstRightAnswer[$i][$j]) {
									$score++;
									break;
								}
							}
						}

						echo "Вы угадали " . $score . " загадок";
					}
			
				?>


				<form method="GET">
					<p>Висит груша, нельзя скушать?</p>
					<input type="text" name="userAnswer1">

					<p>Пустые отдыхаем,<br>А полные шагаем.</p>
					<input type="text" name="userAnswer2">

					<p>Нет ушей, а слышит.<br>Нету рук, а пишет.</p>
					<input type="text" name="userAnswer3">

					<br>
					<input type="submit" value="Ответить">
				</form>

			</div>

        </div>
    </div>
</div>

	

</div>
<div class="footer">
	Copyright &copy; <?php echo date("Y");?> Sergey Shepelev
<div>


</body>
</html>
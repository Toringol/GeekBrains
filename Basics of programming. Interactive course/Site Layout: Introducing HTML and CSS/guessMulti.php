<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>homeworkSite</title>
	<link rel="stylesheet" href="style.css"> 
	<script type="text/javascript">

		var answer = parseInt(Math.random() * 100);

        var gameStarted = false;
        var playerCounter = 0;
        var startNewGame = false;
        
        var players = [
            ["", false],
            ["", false],
        ];

        function Game() {
            if (!gameStarted) {
                startNewGame = false;
                players[0][0] = document.getElementById("input1").value;
                players[1][0] = document.getElementById("input2").value;
                players[0][1] = false;
                players[1][1] = false;
                hide("input1");
                write("info1", "Ход: " + players[playerCounter][0]);
                write("info2", "Угадайте число от 0 до 100");
                gameStarted = true;
            } else {
                if (!players[playerCounter][1]) {
                    write("info1", "Ход: " + players[playerCounter][0]);
                    var userAnswer = readInt();
                    if (userAnswer == answer) {
                        write("info1", "Поздравляю, " + players[playerCounter][0] + ", вы угадали!");
                        hide("info2");
                        hide("input2");
                        players[playerCounter][1] = true;
                    } else if (userAnswer > answer) {
                        write("info1", players[playerCounter][0] + ", Вы ввели слишком большое число.");
                    } else if (userAnswer < answer) {
                        write("info1", players[playerCounter][0] + ", Вы ввели слишком маленькое число.");
                    }
                }
                if (!players[playerCounter][1]) {
                    setTimeout(() => {
                        playerCounter = (playerCounter + 1) % 2;
                        write("info1", "Ход: " + players[playerCounter][0]);
                    }, 1500);
                } else {
                    document.getElementById("button").innerHTML = "Попробовать снова!";
                    gameStarted = false;
                    startNewGame = true;
                }
            }
            resetValue("input2");
        }

        function ButtonAction() {
            if (startNewGame) {
                NewGame();
            } else {
                Game();
            }
        }

        function NewGame() {
            window.location.reload();
        }

        function resetValue(id) {
            document.getElementById(id).value = '';
        }

		function readInt() {
			return +document.getElementById("input2").value;
		}

		function write(id, text) {
			document.getElementById(id).innerHTML = text;
		}

		function hide(id) {
			document.getElementById(id).style.display = "none";
		}

	</script>
</head>
<body>

<div class="content">
	<?php
		include "menu.php";
	?>

<div class="contentWrap">
    <div class="content">
        <div class="center">

			<h1>Игра угадайка на двоих</h1>

			<div class="box">
                <p id="info1">Пожалуйства, введите никнейм первого игрока</p>
                <input type="text" id="input1">
                <p id="info2">Пожалуйства, введите никнейм второго игрока</p>
                <input type="text" id="input2">
                <br>
                <a href="#" onClick="ButtonAction();" id="button">Начать</a>
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
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>homeworkSite</title>
	<link rel="stylesheet" href="style.css"> 
	<script type="text/javascript">

        var newPassword = false;

        var arraySymbols = [
            "A", "B", "C", "D", "E", "F", "G", "H", "I", 
            "K", "L", "M", "N", "O", "P", "Q", "R", "S", 
            "T", "V", "X", "Y", "Z", "a", "b", "c", "d", 
            "e", "f", "g", "h", "i", "k", "l", "m", "n", 
            "o", "p", "q", "r", "s", "t", "v", "x", "y", 
            "z", "0", "1", "2", "3", "4", "5", "6", "7", 
            "8", "9",
        ];
		
        function ButtonAction() {
            if (newPassword) {
                NewPass();
            } else {
                GeneratePassword();
            }
        }

        function GeneratePassword() {
	    var randomPassword = "";
            var passLen = readInt();
            for (var i = 0; i < passLen; i++) {
                randomPassword += arraySymbols[generateRandomNumber(arraySymbols.length - 1)];
            }
            write("length", "<b>Ваш пароль: </b>" + randomPassword + "<br>");
            hide("input1");
            document.getElementById("button").innerHTML = "Новый пароль!";
            newPassword = true;
        }

        function NewPass() {
            window.location.reload();
        }

        function generateRandomNumber(max) {
            return Math.round(Math.random() * max);
        }

        function resetValue(id) {
            document.getElementById(id).value = '';
        }

		function readInt() {
			return +document.getElementById("input1").value;
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

			<h1>Генератор случайных паролей</h1>

			<div class="box">
                <p id="length">Пожалуйства, введите длину пароля</p>
                <input type="text" id="input1">
                <br>
                <a href="#" onClick="ButtonAction();" id="button">Сгенерировать!</a>
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

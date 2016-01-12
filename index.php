<!Doctype html>
<html>
	<head>
		<title>Index</title>
		<meta name="author" content="Torben">
		<meta charset="utf-8">
		<link rel="stylesheet" href="css/main.css" type="text/css">
		<link rel="stylesheet" href="css/header.css" type="text/css">
		<link rel="stylesheet" href="css/links.css" type="text/css">
		<style type="text/css">
			#catHeaderC1R1, #catHeaderC1R2, #catHeaderC1R3, .catLinkC1R1, .catLinkC1R2, .catLinkC1R3 {
				color:#977;
			}

			#catHeaderC2R1, #catHeaderC2R2, #catHeaderC2R3, .catLinkC2R1, .catLinkC2R2, .catLinkC2R3 {
				color:#797;
			}

			#catHeaderC3R1, #catHeaderC3R2, #catHeaderC3R3, .catLinkC3R1, .catLinkC3R2, .catLinkC3R3 {
				color:#779;
			}
		</style>
	</head>
	<!--Start body-->
	<body>
		<?php
			include 'getSQLData.php';
			$arrayNames  = array(
				array("C1R1", "C1R2", "C1R3"),
				array("C2R1", "C2R2", "C2R3"),
				array("C3R1", "C3R2", "C3R3")
			);
			//arrayNames = getArrayNamesData();
			$header = getHeaderData($arrayNames);
			//$linkName = getLinkNameData();
			//$links = getLinksData();



			$linksPerCat = 3;
			$linkName = array(
				$arrayNames[0][0] => array("Drive", "GitHub", "OneDrive"),
				$arrayNames[0][1] => array("ExtremeTech", "TheVerge", "Wired"),
				$arrayNames[0][2] => array("4Chan", "Reddit", "WhatsApp"),
				$arrayNames[1][0] => array("Animehaven", "Crunchyroll", "MyAnimeList"),
				$arrayNames[1][1] => array("Campus-Navi", "Moodle", "TuCan"),
				$arrayNames[1][2] => array("Netflix", "Twitch", "Youtube"),
				$arrayNames[2][0] => array("Congstar", "KSK", "PayPal"),
				$arrayNames[2][1] => array("Amazon", "HumbleBundle", "Steam"),
				$arrayNames[2][2] => array("Imgur", "Kickstarter", "NetflixRoulette")
			);
			$links = array(
				$arrayNames[0][0] => array("http://www.drive.google.com/", "https://github.com/Chabare", "https://onedrive.live.com/"),
				$arrayNames[0][1] => array("http://www.extremetech.com/", "http://www.theverge.com/", "http://www.wired.com"),
				$arrayNames[0][2] => array("http://www.4chan.org/b/", "http://www.reddit.com/", "https://web.whatsapp.com/"),
				$arrayNames[1][0] => array("http://www.animehaven.org", "http://www.crunchyroll.com/home/queue", "http://myanimelist.net/animelist/chabare"),
				$arrayNames[1][1] => array("http://www.sight-board.de/tu-darmstadt/", "https://moodle.informatik.tu-darmstadt.de", "https://www.tucan.tu-darmstadt.de/"),
				$arrayNames[1][2] => array("http://www.netflix.com", "http://www.twitch.tv/directory/following", "http://www.youtube.com/feed/subscriptions"),
				$arrayNames[2][0] => array("http://www.congstar.de", "http://www.meine-ksk.de/", "http://www.paypal.de/"),
				$arrayNames[2][1] => array("http://www.amazon.de/", "http://www.humblebundle.com/", "http://store.steampowered.com/"),
				$arrayNames[2][2] => array("http://www.imgur.com", "http://www.kickstarter.com", "http://www.netflixroulette.net")
			);
		?>

<!--Start first Column-->
		<div id="firstC">
	<!--Start first row-->
			<?php
				$headerArrID = 0;
			?>
			<div class="firstRow">
				<?php
					$arrRowID = 0;
				?>
				<h1 id="catHeaderC1R1"><?php echo $header[$arrayNames[$headerArrID][$arrRowID]]; ?></h1>
				<?php
					for($i = 0; $i < $linksPerCat; $i++)
						echo '<a href="'.$links[$arrayNames[$headerArrID][$arrRowID]][$i].'" class="catLinkC1R1">'.$linkName[$arrayNames[$headerArrID][$arrRowID]][$i].'</a>';
				?>
			</div>




	<!--Start second row-->
			<div class="secondRow">
				<?php
					$arrRowID = 1;
				?>
				<h1 id="catHeaderC1R2"><?php echo $header[$arrayNames[$headerArrID][$arrRowID]]; ?></h1>
				<?php
					for($i = 0; $i < $linksPerCat; $i++)
						echo '<a href="'.$links[$arrayNames[$headerArrID][$arrRowID]][$i].'" class="catLinkC1R2">'.$linkName[$arrayNames[$headerArrID][$arrRowID]][$i].'</a>';
				?>
			</div>




	<!--Start third row-->
			<div class="thirdRow">
				<?php
					$arrRowID = 2;
				?>
				<h1 id="catHeaderC1R1"><?php echo $header[$arrayNames[$headerArrID][$arrRowID]]; ?></h1>
				<?php
					for($i = 0; $i < $linksPerCat; $i++)
						echo '<a href="'.$links[$arrayNames[$headerArrID][$arrRowID]][$i].'" class="catLinkC1R3">'.$linkName[$arrayNames[$headerArrID][$arrRowID]][$i].'</a>';
				?>
			</div>
		</div>




<!--Start second column-->
		<div id="secondC">
	<!--Start first row-->
			<?php
				$headerArrID = 1;
			?>
			<div class="firstRow">
				<?php
					$arrRowID = 0;
				?>
				<h1 id="catHeaderC2R1"><?php echo $header[$arrayNames[$headerArrID][$arrRowID]]; ?></h1>
				<?php
					for($i = 0; $i < $linksPerCat; $i++)
						echo '<a href="'.$links[$arrayNames[$headerArrID][$arrRowID]][$i].'" class="catLinkC2R1">'.$linkName[$arrayNames[$headerArrID][$arrRowID]][$i].'</a>';
				?>
			</div>




	<!--Start second row-->
			<div class="secondRow">
				<?php
					$arrRowID = 1;
				?>
				<h1 id="catHeaderC2R2"><?php echo $header[$arrayNames[$headerArrID][$arrRowID]]; ?></h1>
				<?php
					$arrRowID = 1;
					for($i = 0; $i < $linksPerCat; $i++)
						echo '<a href="'.$links[$arrayNames[$headerArrID][$arrRowID]][$i].'" class="catLinkC2R2">'.$linkName[$arrayNames[$headerArrID][$arrRowID]][$i].'</a>';
				?>
			</div>




	<!--Start third row-->
			<div class="thirdRow">
				<?php
					$arrRowID = 2;
				?>
				<h1 id="catHeaderC2R1"><?php echo $header[$arrayNames[$headerArrID][$arrRowID]]; ?></h1>
				<?php
					$arrRowID = 2;
					for($i = 0; $i < $linksPerCat; $i++)
						echo '<a href="'.$links[$arrayNames[$headerArrID][$arrRowID]][$i].'" class="catLinkC2R3">'.$linkName[$arrayNames[$headerArrID][$arrRowID]][$i].'</a>';
				?>
			</div>

			<!--Weather-->
			<iframe id="forecast_embed" type="text/html" frameborder="0" height="245" width="100%" src="http://forecast.io/embed/#lat=49.8726&lon=8.6501&name=Darmstadt&units=ca" style="background:#999;opacity:0.5;margin-top:50px;"> </iframe>
		</div>




<!--Start third column-->
		<div id="thirdC">
	<!--Start first row-->
			<?php
				$headerArrID = 2;
			?>
			<div class="firstRow">
				<?php
					$arrRowID = 0;
				?>
				<h1 id="catHeaderC3R1"><?php echo $header[$arrayNames[$headerArrID][$arrRowID]]; ?></h1>
				<?php
					for($i = 0; $i < $linksPerCat; $i++)
						echo '<a href="'.$links[$arrayNames[$headerArrID][$arrRowID]][$i].'" class="catLinkC3R1">'.$linkName[$arrayNames[$headerArrID][$arrRowID]][$i].'</a>';
				?>
			</div>




	<!--Start second row-->
			<div class="secondRow">
				<?php
					$arrRowID = 1;
				?>
				<h1 id="catHeaderC3R2"><?php echo $header[$arrayNames[$headerArrID][$arrRowID]]; ?></h1>
				<?php
					for($i = 0; $i < $linksPerCat; $i++)
						echo '<a href="'.$links[$arrayNames[$headerArrID][$arrRowID]][$i].'" class="catLinkC3R2">'.$linkName[$arrayNames[$headerArrID][$arrRowID]][$i].'</a>';
				?>
			</div>




	<!--Start third row-->
			<div class="thirdRow">
				<?php
					$arrRowID = 2;
				?>
				<h1 id="catHeaderC3R1"><?php echo $header[$arrayNames[$headerArrID][$arrRowID]]; ?></h1>
				<?php
					for($i = 0; $i < $linksPerCat; $i++)
						echo '<a href="'.$links[$arrayNames[$headerArrID][$arrRowID]][$i].'" class="catLinkC3R3">'.$linkName[$arrayNames[$headerArrID][$arrRowID]][$i].'</a>';
				?>

			</div>
		</div>


	</body>
</html>

//	
			//spooky code
			function json(url) {
			  return fetch(url).then(res => res.json());
			}
			
			//the rumbles of the city.
			var d = new Date();
			var n = d.getHours();
			var dd = d.getDay();
			var m = d.getMonth();
			
			var fin = n + dd + m - dd;
			
			//get date
			var dd = d.getDate();
			localStorage.setItem("current",dd)
			//localStorage.getItem("last")
			
			console.log(localStorage.getItem("current"))
			console.log(localStorage.getItem("last"))
			console.log(localStorage.getItem("prescript"))
			
			//compare 
			if(localStorage.getItem("current")!=localStorage.getItem("last")){
				//i fucking hate this
				//var name = localStorage.getItem("country") + ":"+ localStorage.getItem("ip").replace(/[&\/\\#,+()$~%.'":*?<>{}]/g, '');
				
				localStorage.setItem("username",name)
				
				new_prescript()
			}
			
			console.log("current name: " +localStorage.getItem("username"))
	
			//generate new prescript
			function new_prescript(){
				//important stuff
				// data.ip < --- ip
				// data.city < --- city
				// data.country_code < --- country (CA,UN,JA, etc)
				
				//prescript 
				//only order breadsticks
				
				
				//decide wether to use a predetermined one or not.
				israndom = Math.floor((Math.random()*75))
				console.log(israndom)
				locationA = [
							"beach",
							"store",
							"house",
							"hotel",
							"motel",
							"theater",
							"library",
							"school",
							"factory",
							"empty field",
							"graveyard",
							"basement",
							"lighthouse",
							"secret location only you know about",
							"place you call home",
							"place they call home",
							"place you'd hide a body",
						]
						locationset = Math.floor(Math.random() * locationA.length)
						finalllocation= locationA[locationset]
						
				if(israndom==4){
					console.log("random")
					preprescript = ["Pack a lunchbox and consume it on top of a trash can in the streets of District 11 at 1 PM today.",
						"Bake dacquoise while the hour hand rests between 7 and 8, and eat it while watching a movie.",
						"Initiate a game of Never Have I Ever with the first five people you encounter. When one folds a finger, break it.",
						"Neatly clip the nails of the sixty-second person you come across.",
						"Pet quadrupedal animals five times.",
						"Spin a wheel and throw a cake at the person determined by the result.",
						"Consume eight crabs stored at room temperature and ripe persimmon at once.",
						"At the railing on the roof of a building, shout out the name of the person you dislike, then jump off. The height of the building does not matter.",
						"After a meal, discard all dishes that were used to serve it.",
						"On the morning after receiving the Prescript, drink three cups of water as soon as you get up.",
						"Race against residents that live in the same building as you to District 7. Measure the distance every twenty-three minutes and disqualify the one farthest away from the destination.",
						"Within three days, knit a scarf with a butterfly pattern.",
						"Dial any number. Give a New Year’s greetings and words of blessing to whoever receives the call.",
						"See green from a white wall.",
						"When hungry, consume a Cheeki’s cheeseburger with added onion.",
						"Fold thirty-nine paper cranes and throw them from the rooftop.",
						"At work, cut the ear of the first person to fulminate against you.",
						"When your eyes meet another person’s, nod at them.",
						"Return to your home this instant. You may leave once a dog barks in front of your house one time.",
						"Wear light green clothing and take 10 steps in a triangle-shaped alley.",
						"Call the first person you meet a homosexual, then proceed to kiss them.",
						"Do not go home until you have finished reading the value of e.",
						"In 400,000 meters, turn right.",
						"Sleep for a total of 800 hours per day.",
						"Only eat and write with your right hand.",
						"Order something online you dont need.",
						"Submit to the cities prescripts.",
						"Waste your money on someone.",
						"Go to sleep. You must dream about a bird locked in a cage. The bird's color does not matter.",
						"Ensure that 2 plus 2 is equal to 5.",
						"Visit the Index and recieve your personalized reward.",
						"Retrace all the steps you took today. Once you are back in bed, touch yourself.",
						"Repeat yesterday's prescript. You must ensure that you wear glasses.",
						"Come back tomarrow and act out your prescript twice.",
						"Obey the first command you are given today.",
						"Go to a "+finalllocation,
					]
					preset = Math.floor(Math.random() * preprescript.length)
					finalscript = preprescript[preset]
				}else{
					//adjactive
							personidentifyer = [
							"ugly",
							"blonde",
							"kind",
							"cute",
							"brunette",
							"selfish",
							"young",
							"feeble",
							"pale",
							"pretty",
							"dark",
							"strong",
							"weak",
							"sexually-attractive",
							"small",
							"big",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							]
							personidentifyerset = Math.floor(Math.random() * personidentifyer.length)
							finalpersonidentifyer = personidentifyer[personidentifyerset]
						//adjactive (PLURAL)
							personidentifyer2 = [
							"ugliest",
							"kindest",
							"nicest",
							"cutest",
							"friendliest",
							"meanest",
							"dumbest",
							"sexiest",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"ugliest",
							"most-generous",
							"nicest",
							"cutest",
							"nerdiest",
							"meanest",
							"dumbest",
							"dorkiest",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							]
							personidentifyerset2 = Math.floor(Math.random() * personidentifyer2.length)
							finalpersonidentifyer2 = personidentifyer2[personidentifyerset2]
					//persontype
							color1 = [
							"blue ",
							"black ",
							"red ",
							"brown ",
							"velvet ",
							"yellow ",
							"pink ",
							"white ",
							"grey ",
							"green ",
							"light green ",
							"light red ",
							"cyan ",
							"purple ",
							"blood red ",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							"",
							]
							color1set = Math.floor(Math.random() * color1.length)
							colorrequire = Math.floor(Math.random() * 14)
							finalcolor1 = color1[color1set]
							finalcolorrequire = color1[colorrequire]
							
					//persontype
							persontype = [
							"friend",
							"child",
							"stranger",
							"priest",
							"butcher",
							"authority figure",
							"homeless man",
							"drunk man",
							"homeless woman",
							"drunk woman",
							"neighbor",
							"person you meet",
							"stranger",
							"stranger",
							"stranger",
							"person",
							"person",
							"person",
							"person",
							]
							persontypeset = Math.floor(Math.random() * persontype.length)
							finalptype = persontype[persontypeset]
					//persontype2
							persontype2 = [
							"friend",
							"stranger",
							"family",
							"quadrupedal animal",
							"neighbor",
							"lover",
							"animal",
							"stranger",
							"stranger",
							"stranger",
							"person",
							"person",
							"person",
							"member of the index",
							]
							persontype2set = Math.floor(Math.random() * persontype2.length)
							finalptype2 = persontype2[persontype2set]
					//persontype3 (personal)
							persontype3 = [
							"friend",
							"neighbor",
							"family member",
							"animal",
							]
							persontype3set = Math.floor(Math.random() * persontype3.length)
							finalptype3 = persontype3[persontype3set]
					
					isperson = Math.floor((Math.random()*2))
					
					number1 = Math.floor(Math.random() * 100+1)
					number7 = Math.floor(Math.random() * 100+1)
					number2 = Math.floor(Math.random() * 400+1)
	
					number3 = Math.floor(Math.random() * 100+1)
					number4 = Math.floor(Math.random() * 1300+1)
					
					number5 = Math.floor(Math.random() * 6+1)
					number9 = Math.floor(Math.random() * 8+1)
					number8 = Math.floor(Math.random() * 7+1)
					number11 = Math.floor(Math.random() * 12+1)
					
					end1 = "th"
					end7 = "th"
					end8 = "th"
					//thx _@jj <3
					switch(number1%10) {
					  case 1: end1 = "st"; break;
					  case 2: end1 = "nd"; break;
					  case 3: end1 = "rd"; break;
					}
					end = "th"
					switch(number7%10) {
					  case 1: end7 = "st"; break;
					  case 2: end7 = "nd"; break;
					  case 3: end7 = "rd"; break;
					}
					end = "th"
					switch(number8%10) {
					  case 1: end8 = "st"; break;
					  case 2: end8 = "nd"; break;
					  case 3: end8 = "rd"; break;
					}
					
					
					clothing = ["hat",
								"coat",
								"skirt",
								"dress",
								"socks",
								"gloves",
								"boots",
								"glasses",
								"ribbon",
								"hairpin",
								"cloak",
								"shirt",
								"pants",
								"mask",
								"shorts",
								"panties",
								"coat",
								"skirt",
								"dress",
								"suit",
								"gloves",
								"boots",
								"shirt",
								"pants",
								"shoes",
								"sweatshirt",
								"hood",
							]
					clothingset = Math.floor(Math.random() * clothing.length)
					finalclothing = clothing[clothingset]
					clothingset2 = Math.floor(Math.random() * clothing.length)
					finalclothing2 = clothing[clothingset2]
					
					//time
					task0 = ["","","","","","","","","","","","","","","","","","","","","","","","","","","",
						"At work, ",
						"On the way to work ",
						"Within "+number7+" days, ",
						"In exactly "+number7+" minutes, ",
						"In the next hour, ",
						"Sometime today, ",
						"Wear a "+finalcolorrequire+" "+finalclothing+" and ",
						"At noon, ",
						"If it is raining, ",
						"After you each lunch, ",
						"When you go to sleep, get up and ",
						"If it is the "+number8+end8+", ",
						"Tonight, ",
						"At midnight, ",
						"Within the next hour, ",
						"Dress in only "+finalcolorrequire+" and ",
						"When you talk to a "+finalpersonidentifyer+" "+finalptype2+", ",
						"When you see a "+finalptype2+", ",
						"When you are in the city, ",
						"At the railing of a roof, ",
						"While sitting down, ",
						"If you feel sick, ",
						"When you see "+finalcolorrequire+", ",
						"While driving down the freeway, ",
						"After recieving this prescript, ",
						"When you breathe, ",
						"When your eye's meet another "+finalptype2+", ",
						"After a meal, ",
						"When preparing to leave your home, ",
						"Next time you are angry, ",
						"While in the back of a van, ",
						"After you see blood, ",
						"If you get wet, ",
						"On the day of your birthday, ",
						"When you are crying, ",
						"In "+number7+" minutes, ",
						"Walk "+number2+" meters, then walk "+number7+" meters in any direction, ",						
						"In "+number7+" minutes, ",
						"In "+number7+" hours, ",
						"In "+number7+" seconds, ",
						"In "+number7+" minutes, ",
						"After walking "+number2+" meters, ",
						"When encountering the "+number7+""+end7+" person you meet, ",
						"Next time you are scared, ",
						"Sometime later today, ",
						"When you are aroused, "
					]
					task0set = Math.floor(Math.random() * task0.length)
					finaltask0 = task0[task0set]
					
					//games
					games = [
						"Patty Cake",
						"Never Have I Ever",
						"Truth or Dare",
						"Tag",
						"Hide and Seek",
						"Paintball",
						"Peekaboo",
						"Hopscotch",
						"Chess",
						"Twister",
					]
					gamesset = Math.floor(Math.random() * games.length)
					finalgames= games[gamesset]
					
					//activities
					activities = [
						"star gazing",
						"swimmng",
						"on a walk",
						"find a motel",
						"start a fight",
						"gardening",
						"drink alcohol",
						"on a dinner date",
						"climb in bed",
						"pet a dog",
						"go fight a "+finalptype2,
						"play "+finalgames+" with them",
					]
					activitiesset = Math.floor(Math.random() * activities.length)
					finalactivities= activities[activitiesset]
					
					objcs = ["knife",
								"shovel",
								"9mm pistol",
								"pen and paper",
								"rashions",
								"rasions",
								"apples",
								finalcolor1+"notebook with "+number1+" pages",
								"plant with "+number5+" leaves",
								"duct tape",
								"charcoal",
								"something plastic",
								"wooden chair",
								"bikelock",
								"pack of playing card",
								"spoon",
								"crayons",
								"paint",
								"cat",
								"dvd player",
								"televison",
								"toilet paper",
								"hooks",
								"cat",
								"dvd player",
								"televison",
								"toilet paper",
								"baseball bat",
							]
					objcsset = Math.floor(Math.random() * objcs.length)
					finalobjcs = objcs[objcsset]
					
					
				
						//topics
							topic = [
							"love",
							"life",
							"gaming",
							"the city",
							"a movie",
							"the "+number1+""+end1+" person you met",
							"your crush",
							"your favorite food",
							"your secrets",
							"your desires",
							"yourself",
							"the weather",
							"state of affairs",
							"hatsune miku",
							"dreams",
							"the latest post on twitter.com",
							"alien life",
							"the "+number1+""+end1+" post on twitter.com",
							"the "+number1+""+end1+" post on reddit",
							"birds",
							"smoke",
							"crypto",
							"drugs",
							"your favorite "+finalclothing,
							"money",
							"your sexuality",
							"rocks",
							"candy",
							"mirrors",
							"reality",
							"religion",
							"science",
							"doubt",
							"free will",
							"faith",
							"math",
							"politics",
							"desires",
							"technology",
							"the wings",
							"death",
							"loss",
							"memories",
							"food",
							finalgames,
							]
							topicset = Math.floor(Math.random() * topic.length)
							finaltopic = topic[topicset]
							topicset2 = Math.floor(Math.random() * topic.length)
							finaltopic2 = topic[topicset2]
							
							
						//topics
							if(number5>5){
								if(finalpersonidentifyer!=""){
									howtofind = [
										"a person you hate",
										"the last "+number5+" "+finalpersonidentifyer+" people you see before you get home",
										"the first "+number5+" "+finalpersonidentifyer+" people that looks at you ",
										"the first "+number5+" "+finalpersonidentifyer+" people you meet after recieving the prescript ",
										"the closest "+number5+" "+finalptype+"s",
										"the "+number1+""+end1+" "+finalptype,
										"the "+number8+""+end8+" person you've had a crush on",
										"the "+number8+""+end8+" person that comes to mind",
										"the "+finalpersonidentifyer2+" "+finalptype,
										"a "+finalptype,
										"the one you love most",
										"the "+finalpersonidentifyer2+" girl",
										"the "+finalpersonidentifyer2+" boy",
										"a "+finalpersonidentifyer+" girl",
										"a "+finalpersonidentifyer+" girl with a "+finalcolor1+finalclothing,
										"a "+finalpersonidentifyer+" boy with a " +finalcolor1+finalclothing,
										"a "+finalpersonidentifyer+" boy",
										"a person of your choice",
										"the last person to talk to you",
										"an animal",
										"a quadrupedal animal"
										]
										howtofindset = Math.floor(Math.random() * howtofind.length)
									}else{
										howtofind = [
										"a person you hate",
										"the last person you see before you get home",
										"the "+number8+""+end8+" person you've had a crush on",
										"the first person that looks at you ",
										"the "+number8+""+end8+" person that comes to mind",
										"the last person to talk to you",
										"the first person you meet after recieving the prescript ",
										"the closest "+finalptype,
										"the "+number1+""+end1+" "+finalptype,
										"your favorite "+finalptype3,
										"someone you hate",
										"someone you love",
										"the one you love most",
										"a person of your choice",
										"a girl",
										"a boy",
										"an animal",
										"a quadrupedal animal"
										]
										howtofindset = Math.floor(Math.random() * howtofind.length)
									}
								}else{
									if(finalpersonidentifyer!=""){
										howtofind = [
											"a person you hate",
											"the last "+finalpersonidentifyer2+" person you see before you get home",
											"the first "+finalpersonidentifyer2+" person that looks at you",
											"the first "+finalpersonidentifyer2+" person you meet after recieving the prescript",
											"the closest "+finalptype,
											"the "+number8+""+end8+" person that comes to mind",
											"the "+number8+""+end8+" person you've had a crush on",
											"the "+number8+""+end8+" person you see wearing "+finalcolor1+finalclothing,
											"the "+number1+""+end1+" "+finalptype,
											"the "+finalpersonidentifyer2+" "+finalptype,
											"the last person to talk to you",
											"a "+finalptype,
											"the one you love most",
											"the "+finalpersonidentifyer2+" girl",
											"the "+finalpersonidentifyer2+" boy",
											"a "+finalpersonidentifyer+" girl",
											"a "+finalpersonidentifyer+" boy",
											"a "+finalpersonidentifyer+" girl with a "+finalcolor1+finalclothing,
											"a "+finalpersonidentifyer+" boy with a " +finalcolor1+finalclothing,
											"an animal",
											"a person of your choice",
											]
											howtofindset = Math.floor(Math.random() * howtofind.length)
										}else{
											howtofind = [
											"a person you hate",
											"the last person you see before you get home",
											"the first person that looks at you ",
											"the first person you meet after recieving the prescript ",
											"the "+number8+""+end8+" person that comes to mind",
											"the "+number8+""+end8+" person you've had a crush on",
											"the "+number8+""+end8+" person you see wearing "+finalcolor1+finalclothing,
											"the closest "+finalptype,
											"the "+number1+""+end1+" "+finalptype,
											"your favorite "+finalptype3,
											"the last person to talk to you",
											"a person of your choice",
											"the one you love most",
											"a girl",
											"a boy",
											"a girl with a "+finalcolor1+finalclothing,
											"a boy with a "+finalcolor1+finalclothing,
											"an animal"
											]
											howtofindset = Math.floor(Math.random() * howtofind.length)
										}
										
								}
										
						
						finalhowtofind = howtofind[howtofindset]
						
						//actions
						actionss = [
							"commit a felony",
							"steal money",
							"betray a friend",
							"become a proxy",
							"play "+finalgames+" with them",
							"talk about "+finaltopic+"",
							"shake your hand",
							"drink something",
							"sit on the train tracks",
							"travel somewhere new",
							"drink from a puddle",
							"eat something new",
							"do a backflip",
							"do a frontflip",
							"do a lap around the city",
							"swap wallets with you",
							"roll a dice",
							"swing a baseball bat",
							"make a cake for a "+finalptype3+"",
							"deliver a prescript",
							"watch a movie",
							"go out to dinner",
							"climb into a trash can",
							"jump off a bridge",
							"draw with crayons",
							"proceed to light them on fire.",
							"ask about "+finaltopic2+".",
							"ask them to marry you.",
							"check for weapons",
							"discard this prescript.",
							"make them a nice meal.",
							"give them a hug before pushing them off a bridge",
							"go for a drive to a "+finalllocation,
							"measure their height.",
							"get their name.",
							"sleep with them.",
							"cut off all your hair",
							"cover your face",
							"buy them "+number9+" "+finalcolorrequire+"objects.",
							"dye your hair "+finalcolorrequire,
							"peel open a scab and drink the blood",
							"measure their height.",
							"shake hands with your sibllng.",
							"do the same to their children.",
							"turn the thermostat to "+number7+"F'",
							"leave food outside the door of your house.",
							"leave food outside of a "+finalllocation,
							"bring them to a "+finalllocation,
							"go home.",
							"forget about this prescript",
							"hit the gritty"
						]
						actionsssset = Math.floor(Math.random() * actionss.length)
						finalactionss= actionss[actionsssset]
						
						verb = [
							"stab",
							"hug",
							"poke",
							"kiss",
							"punch",
							"kick",
							"compliment",
							"scream at",
							"shoot",
							"carress",
							"rub",
							"pet",
							"poke",,
							"choke",
							"smother",
							"lick",
							"tie up",
							"abuse",
							"feed",
							"break",
							"poke",
							"make love to",
							"eat",
							"throw",
							"ignite",
							"cut",
							"call the police on",
							"paint on",
							"cut a "+number2+" into"
						]
						verbset = Math.floor(Math.random() * verb.length)
						finallverb= verb[verbset]
						
						verb2 = [
							"stabed",
							"huged",
							"poked",
							"kissed",
							"punched",
							"kicked",
							"complimented",
							"screamed at",
							"shot",
							"petted",
							"ate",
							"poked",
							"choked",
							"smothered",
							"licked",
							"tied up",
							"abused",
							"fed",
							"broke",
							"make love to",
						]
						verb2set = Math.floor(Math.random() * verb2.length)
						finallverb2= verb2[verb2set]
						
						thingstosay = [
							"the ABC's",
							"the anthem of your country",
							"the names of everyone you met in the last "+Math.floor(Math.random() * 240)+" seconds",
							"your favorite person's type of "+finalclothing,
							"the amount of times you have "+finallverb2+" someone in the last "+Math.floor(Math.random() * 10)+" days",
							"a story about a time you had at a "+finalllocation,
							"the story of your first kiss",
							"your loneiness",
							"how to divide by 0",
							"the value of e",
							"something about the "+finalhowtofind+" you met",
							"the "+finalclothing+" "+finalhowtofind+" are wearing",
						]
						thingstosayset = Math.floor(Math.random() * thingstosay.length)
						finalthingstosay= thingstosay[thingstosayset]
						
						additionalrandom = ["","","","","","","","","","","","","","","","","","","","","","","","","","","","","","","","","",
								" backwards",
								" with "+finalhowtofind+" watching",
								" on a rooftop",
								" at a "+finalllocation,
								" on the "+number8+end8+" of this month",
								" "+number2+" times",
								" as quickly as you can",
								" while in a fetal position",
								" when you are starving",
								" while playing a game of "+finalgames,
								" while staring at the color "+finalcolorrequire,
								" if you feel upset",
								" if you see "+finalptype2,
								" in public",
								" while wearing a "+finalcolorrequire+finalclothing,
								" while chewing on taffie",
								" during sex",
								" while looking at a mirror",
								" within "+number2+" minutes",
								" with your eyes closed",
								" if you feel like you are worthless",
								" if you do not see the color "+finalcolorrequire,
							]
							additionalrandomset = Math.floor(Math.random() * additionalrandom.length)
							finaladditionalrandom = additionalrandom[additionalrandomset]
						
						//verb
						//action
						task1 = ["walk up to "+finalhowtofind,
							"manipulate "+finalhowtofind+" to "+finalactionss,
							"hop like a bunny to "+finalhowtofind,
							"confess your love to "+finalhowtofind,
							"clip the nails of "+finalhowtofind,
							"offer to walk with "+finalhowtofind,
							"hit "+finalhowtofind+" with the first "+finalobjcs+" you find",
							"give a "+finalobjcs+" to "+finalhowtofind,
							"clean the shoes of "+finalhowtofind,
							"clean the shoes of "+finalhowtofind,
							"ask "+finalhowtofind+" what their favorite food is",
							"chase "+finalhowtofind+" and "+finallverb+" them",
							"report "+finalhowtofind+" to authorities",
							"read a book with "+finalhowtofind+finaladditionalrandom,
							"go to a "+finalllocation+" with "+finalhowtofind,
							"slap "+finalhowtofind,
							"go to a concert with "+finalhowtofind,
							"sell all your possessions to "+finalhowtofind,
							"go to the park with "+finalhowtofind,
							"go camping with "+finalhowtofind,
							"watch a movie with "+finalhowtofind,
							"find "+finalhowtofind+" and ask them to "+finalactionss,
							"interact with "+finalhowtofind,
							"yell at "+finalhowtofind+" about "+finaltopic,
							"have a debate about "+finaltopic,
							"sleep with "+finalhowtofind,
							"play with "+finalhowtofind,
							"break "+finalhowtofind+"'s bone. The bone can be of your choice",
							"use a whip and flay "+finalhowtofind,
							"make the "+finalhowtofind+" bleed",
							"lie to "+finalhowtofind,
							"have a race with "+finalhowtofind,
							"attack "+finalhowtofind,
							"use a taser and zap "+finalhowtofind+" until they stop moving",
							"listen to the ramblings of the "+finalhowtofind,
							"touch "+finalhowtofind,
							"bite "+finalhowtofind,
							"give a love letter to "+finalhowtofind,
							"kiss "+finalhowtofind+finaladditionalrandom,
							"punch "+finalhowtofind+" and beat them down"+finaladditionalrandom,
							"grab hold of "+finalhowtofind+" and "+finalactionss,
							"dont let "+finalhowtofind+" pass you",
							"stop "+finalhowtofind+" and "+finalactionss,
							"tap "+finalhowtofind+" on the right shoulder",
							"tap "+finalhowtofind+" on the right shoulder take them to "+finalactionss,
							"shake hands with "+finalhowtofind,
							"sing about "+finaltopic+" to "+finalhowtofind,
							"have a conversation about "+finaltopic+" with the "+finalhowtofind,
							"go out for dinner with "+finalhowtofind,
							"go on a date with "+finalhowtofind,
							"play "+finalgames+" with "+finalhowtofind,
							"start a fight "+finalhowtofind+". Make sure that you break 4 bones",
							"decline the request of "+finalhowtofind,
							finallverb+" "+finalhowtofind+" until they cant move",
							"find someone who likes "+finaltopic,
							"make a deal with "+finalhowtofind+" you encounter",
							"give "+finalhowtofind+" a gift",
							"give this prescript to "+finalhowtofind+"",
						]
					
task1set = Math.floor(Math.random() * task1.length)
						finaltask1 = task1[task1set]
					
					//starter
					starter = ["","","","","","","","","","","","","","","","","","","","","","","","","","","",
						"spin around in a circle counter-clockwise, then ",
						"spin around in a circle clockwise, then ",
						"wait "+number4+" seconds, then ",
						"walk "+number4+" meters, then turn left, ",
						"walk "+number4+" meters, then turn right, ",
						"walk "+number4+" meters straight, then ",
						"walk "+number4+" meters in any direction, ",
						"look at the clouds and ",
						"remember that you should ",
						"discard this prescript and ",
						"go to sleep and while you dream, ",
						"ensure that you ",
						"dress for cold weather and ",
						"dress for warm weather and ",
						"initiate a game of "+finalgames+" and ",
					]
					//"initiate a game of "+finalgames+" with "+finalhowtofind," and",
					starterset = Math.floor(Math.random() * starter.length)
					finalstarter = starter[starterset]
					if(finaltask0==""){
						finalstarter = finalstarter.charAt(0).toUpperCase() + finalstarter.slice(1)
					}
					
						
		
							
							
					if(isperson==1){
						console.log("isperson!")
						
						
						if(finaltask0==""){
							if(finalstarter==""){
							finaltask1 = finaltask1.charAt(0).toUpperCase() + finaltask1.slice(1)
							}
						}
						
						followupA = Math.floor((Math.random()*3))
						followupB = Math.floor((Math.random()*4))
						if(followupA==2){
							console.log("followup type 1")
							transition = [", then ",
							" and ",
							". After, ",
							". In "+number11+" hours ",
							]
							
							followup = [
								"proceed to light them on fire.",
								"ask about "+finaltopic2+".",
								"ask them to marry you.",
								"strip down for them.",
								"discard this prescript.",
								"make them a nice meal.",
								"give them a hug before pushing them off a bridge",
								"go for a drive to a "+finalllocation,
								"measure their height.",
								"get their name.",
								"sleep with them.",
								"cut off all your hair",
								"cover your face",
								"buy them "+number9+" "+finalcolorrequire+"objects.",
								"dye your hair "+finalcolorrequire,
								"peel open a scab and drink the blood",
								"measure their height.",
								"ignore them entirely.",
								"shake hands with your sibllng.",
								"do the same to their children.",
								"pull their hair out.",
								"ask about their family.",
								"ask if they want to "+finalactivities,
								finalactivities,
								"insert a knife into their chest.",
								finallverb +" them "+number9 +" times",
								finallverb +" them "+number9 +" times, then go "+finalactivities,
								"forget about them.",
								"turn the thermostat to "+number7+"F'",
								"leave food outside the door of your house.",
								"leave food outside of a "+finalllocation,
								"bring them to a "+finalllocation,
								"go home.",
							]
							followupset = Math.floor(Math.random() * followup.length)
							finalfollowup = followup[followupset]
							
							
							if(followupB==2){
								follow2up = [" Their gender does not matter.",
								" You must sleep "+number9+" hours tonight.",
								" You must hold your breathe during the interaction.",
								" Ensure that you do not blink.",
								" Remember to wear a "+finalcolorrequire+"hat",
								" A body must be burried by tomarrow.",
								" You must disobey your next prescript.",
								" You must believe the lie you are told.",
								" Remember to leave a "+finalobjcs+" at a "+finalllocation,
								" You will recieve a text",
								" Climb into the freezer and wait "+number9+" hours.",
								" Make sure there is a "+finalcolor1+"wall to your left",
								" You are not allowed to contact them after.",
								" You must keep in contact with them for the next "+number9+" hours.",
								" You cannot talk after "+number9+"pm.",
								" Make sure you tell them to look behind them.",
								" Take a bath for "+number9+" minutes to clean yourself.",
								" Return home directly after"+finaladditionalrandom+".",
								" Think about them in the shower later.",
								" Go to a "+finalllocation+" and go "+finalactivities,
								" Follow them home, and remember where they live.",
								" After they leave, use the bathroom"+finaladditionalrandom+".",
								" Leave them all of your money.",
								" Give "+number9+" dollars to a homeless man"+finaladditionalrandom+".",
								" It does not matter how long you take.",
								" Step on exactly "+number9+" bugs.",
								" Sleep outside their front door tonight.",
								" Bake a pie and eat it in "+number9+" hours.",
								" You are not allowed to fight back.",
								" You are only allowed to cry after it is done."
								]
								followup2set = Math.floor(Math.random() * follow2up.length)
								finalfollowup2 = follow2up[followup2set]
							}else{
								finalfollowup2=""
							}
							
							
							transitionset = Math.floor(Math.random() * transition.length)
							finaltransition = transition[transitionset]
							if(finalfollowup2==""){
								finalscript = finaltask0 + finalstarter + finaltask1 + finaltransition + finalfollowup + "."
							}else{
								finalscript = finaltask0 + finalstarter + finaltask1 + finaltransition + finalfollowup + "."+ finalfollowup2
							}
						}else{
							
						
						if(followupB==2){
						
							
							follow2up = [" Their gender does not matter.",
								" You must sleep "+number9+" hours tonight.",
								" You must hold your breathe during the interaction.",
								" Ensure that you do not blink.",
								" Remember to wear a "+finalcolorrequire+"hat",
								" A body must be burried by tomarrow.",
								" You must disobey your next prescript.",
								" You must believe the lie you are told.",
								" You are not allowed to fight back.",
								" Remember to leave a "+finalobjcs+" at a "+finalllocation,
								" You will recieve a text",
								" Climb into the freezer and wait "+number9+" hours.",
								" Make sure there is a "+finalcolor1+"wall to your left",
								" You are allowed to become friends with them.",
								" You cannot talk after "+number9+"pm.",
								" Make sure you tell them to look behind them.",
								" Take a bath for "+number9+" minutes to clean yourself.",
								" Return home directly after"+finaladditionalrandom+".",
								" Think about them in the shower later.",
								" Go to a "+finalllocation+" and go "+finalactivities,
								" Follow them home, and remember where they live.",
								" After they leave, use the bathroom"+finaladditionalrandom+".",
								" Leave them all of your money.",
								" Give "+number9+" dollars to a homeless man"+finaladditionalrandom+".",
								" It does not matter how long you take.",
								" Step on exactly "+number9+" bugs.",
								" Sleep outside their front door tonight.",
								" Bake a pie and eat it in "+number9+" hours.",
								" You are only allowed to cry after it is done."
							]
							followup2set = Math.floor(Math.random() * follow2up.length)
						    finalfollowup2 = follow2up[followup2set]
						}else{
							finalfollowup2=""
						}
						if(finalfollowup2==""){
								finalscript = finaltask0 + finalstarter + finaltask1 + "."
							}else{
								finalscript = finaltask0 + finalstarter + finaltask1 + "." + finalfollowup2
							}
						//No followup
						
							
					
							
							//generate final script.
							//finalscript = finaltask0 + finalstarter + finaltask1
							
					}

					//finalscript = "To " + name + ", " + finaltask0 + finalstarter + finaltask1
				}else{
				
				//time	
				//console.log("I shouldnt be running, but I am, im such a naughy boy!")
						task0 = ["","","","","","","","","","","","","","","","","","","","","","","","","","","","","","","","","",
								"At work, ",
								"Wear a "+finalclothing2+" and ",
								"During the day, ",
								"Within "+number7+" days, ",
								"In the next hour, ",
								"When hungry, ",
								"At noon, ",
								"Tonight, ",
								"At midnight, ",
								"To the first person you talk to, ",
								"Within the next hour, ",
								"To the first animal you meet, ",
								"Upon completing your meal, ",
								"When you are in the city, ",
								"At the railing of a roof, ",
								"While driving down the freeway, ",
								"After recieving this prescript, ",
								"When you breathe, ",
								"After a meal, ",
								"If you make eye contact with another, ",
								"Next time you are angry, ",
								"While in the back of a van, ",
								"After being hurt, ",
								"The next time you wake up, ",
								"In "+number7+" minutes, ",
								"In "+number7+" months, ",
								"After walking "+number2+" meters, ",
								"When you next think of this prescript, ",
							]
							task0set = Math.floor(Math.random() * task0.length)
							finaltask0 = task0[task0set]
						
						task3 = ["recite "+finalthingstosay+finaladditionalrandom,
								"recite "+finalthingstosay+" for "+number8+" hours"+finaladditionalrandom,
								"draw a picture of "+finalllocation,
								"paint a wall "+finalcolorrequire,
								"dig a grave at a "+finalllocation,
								"construct a shelter using only a "+finalobjcs,
								"order a pizza and pay with exactly "+number8+" coins",
								"do not intervene when a fight breaks out",
								"replace the batteries of an electronic device",
								"drink "+number8+" cups of water",
								"eat only fish"+finaladditionalrandom,
								"start a fire",
								"fold "+number8+" paper cranes and "+finallverb+" them",
								"knit a "+finalcolorrequire+" scarf, then give it to "+finalhowtofind,
								"sew a piece of clothing with "+finalcolorrequire+" thread",
								"turn the nearest radio to channel "+number8+" and listen for "+number2+" minutes",
								"dial a random number "+number8+" times.  they cannot be the same",
								"go to the store and buy a fish. Take it home and name it after "+finalhowtofind,
								"leave a "+finalobjcs+" outside your front door",
								"collect "+number2+" bottle caps",
								"do not look at any clocks until you "+finaltask1,
								"do not go home until you "+finaltask1,
								"you are not allowed to sleep until you "+finaltask1,
								"sit and read a book at a "+finalllocation,
								"go to the store and buy a "+finalobjcs,
								"change into your favorite "+finalclothing2+finaladditionalrandom,
								"undress in the street"+finaladditionalrandom,
								finalactivities+finaladditionalrandom,
								"go to sleep immediately",
								"do not answer any texts you recieve",
								"speak only in riddles until "+number8+" minutes have passed",
								"dial a random number, have a conversation about "+finaltopic+". Do not let them hang up",
								"go to the first house you see and knock on the door. You must "+ finallverb+" them "+ number9 +" times",
								"turn to the first person you see and "+ finallverb +" them "+ number9 +" times",
								"wait in a dark alley for someone with a "+finalcolorrequire+finalclothing2+". "+ finallverb.charAt(0).toUpperCase() + finallverb.slice(1) +" them ",
								"stand in the middle of the road. When someone comes to talk to you, "+ finallverb +" them",
								"stand still for "+number8+" hours"+finaladditionalrandom,
								"do a backflip"+finaladditionalrandom,
								"remember to brush your teeth"+finaladditionalrandom,
								"write a short book about "+finaltopic+finaladditionalrandom,
								"only talk backwards",
								"you must not speak for the next "+number8+" hours",
								"stay out of the light",
								"eat a bag of your favorite food",
								"eat only foods that are "+finalcolorrequire,
								"push the closest button to you. You may look around to find it",
								"make a picnic and eat it inside a "+finalllocation+" in "+number8+" hours",
								"text a "+finalptype+finaladditionalrandom+" and ask if they want to  "+finalactivities,
								"act out your regular routine"+finaladditionalrandom,
								"do not use the bathroom for "+number8+" hours",
								"remain outside for "+number2+" seconds"
							]
							task3set = Math.floor(Math.random() * task3.length)
							finaltask3 = task3[task3set]
							
							if(finaltask0==""&finalstarter==""){
								finaltask3 = finaltask3.charAt(0).toUpperCase() + finaltask3.slice(1)
							}
							if(finaltask0==""&finalstarter!=""){
									finalstarter = finalstarter.charAt(0).toUpperCase() + finalstarter.slice(1)
							}
							followupB = Math.floor((Math.random()*4))

							if(followupB==2){
								follow2up = [" Their gender does not matter.",
								" You must sleep "+number9+" hours tonight.",
								" You must hold your breathe during the interaction.",
								" Ensure that you do not blink.",
								" Remember to wear a "+finalcolorrequire+"hat",
								" Your socks must not match",
								" You must disobey your next prescript.",
								" When you go to sleep at the end of the day, "+finallverb+" yourself.",
								" Remember to leave a "+finalobjcs+" at a "+finalllocation,
								" Send a text to your closest family member.",
								" Climb into the freezer and wait "+number9+" hours.",
								" Make sure there is a "+finalcolor1+"wall to your right.",
								" Do not talk to anyone until tomarrow.",
								" You cannot talk after "+number9+"pm.",
								" Go to the bank and deposit exactly "+number9+" dollars.",
								" Spin "+Math.floor((Math.random()*360))+" degree's and paint whatever is infront of you "+finalcolorrequire,
								" Take a bath for "+number9+" minutes to clean yourself.",
								" Return home directly after"+finaladditionalrandom+".",
								" Touch yourself while showering in "+number9+" minutes.",
								" Go to a "+finalllocation+" and "+finalactivities,
								" Dig a hole and climb into it.",
								" Write the name of your favorite "+finalllocation+"on paper and leave it under a trash can.",
								" Put up a drawing on a telephone pole.",
								" Give "+number9+" dollars to a homeless man"+finaladditionalrandom+".",
								" It does not matter how long you take.",
								" Step on exactly "+number9+" bugs.",
								" Bake a pie and eat it in "+number9+" hours.",
								" You must sleep face down tonight.",
								" Remember this prescript for the next "+number9+" months.",
								" You must not cry."
								]
								followup2set = Math.floor(Math.random() * follow2up.length)
								finalfollowup2 = follow2up[followup2set]
							}else{
								finalfollowup2=""
							}
							
							finalscript = finaltask0 + finalstarter + finaltask3 +"."+ finalfollowup2
						}
					}
					
					
				localStorage.setItem("name",dd)
				localStorage.setItem("last",dd)
				localStorage.setItem("prescript",finalscript)
				console.log("new prescript generated!")
			}
			prescript.innerHTML = display()
			function display(){
				//console.log("loading old script.")
				sentence = localStorage.getItem("prescript")
			return sentence
			
			//prescript.
			//prescript.innerHTML = display()
			}	

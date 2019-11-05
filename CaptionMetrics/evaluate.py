# coding=utf-8
import json
import sys
import os
import numpy as np
import time
import zipfile
from pycocoevalcap.bleu.bleu import Bleu
from pycocoevalcap.cider.cider import Cider
from pycocoevalcap.meteor.meteor import Meteor
from pycocoevalcap.rouge.rouge import Rouge

# 错误字典，这里只是示例
error_msg={
    1: "Bad input file",
    2: "Wrong input file format",
    3: 'no keyword about result',
    4: 'result length is zero',
    11: 'bleu',
    12: 'cider',
    13: 'meteor',
    14: 'rouge',
    21: 'the length of submit result is unequal with the length of standard result',
    22: 'one of your video_ids is mismatching'
}

def dump_2_json(info, path):
    with open(path, 'w') as output_json_file:
        json.dump(info, output_json_file)

def report_error_msg(detail, showMsg, out_p):
    error_dict=dict()
    error_dict['errorDetail']=detail
    error_dict['errorMsg']=showMsg
    error_dict['score']=0
    error_dict['scoreJson']={}
    error_dict['success']=False
    dump_2_json(error_dict,out_p)

def report_score(score, extra, out_p):
    result = dict()
    result['success']=True
    result['score'] = score
    result['scoreJson'] = extra
    dump_2_json(result,out_p)

# 计算得分
def get_score(gts, res):
    s1  = bleu(gts, res) or 0
    s2 = cider(gts, res) or 0
    s3 = meteor(gts, res) or 0
    s4 = rouge(gts, res) or 0
    total = s1+s2+s3+s4
    obj = {'bleu':round(s1,3), 'cider':round(s2,3), 'meteor':round(s3,3), 'rouge':round(s4,3), 'score': round(total,3)}
    return total, obj
    


def bleu(gts, res):
    try:
        scorer = Bleu(n=4)
        score, scores = scorer.compute_score(gts, res)
        # print('belu = %s' % score)
        return score[3]
    except Exception as e:
        print (e)
        return 0
        check_code = 11
        report_error_msg(error_msg[check_code],error_msg[check_code], out_path)
  

def cider(gts, res):
    try:
        scorer = Cider()
        (score, scores) = scorer.compute_score(gts, res)
        # print('cider = %s' % score)
        return score
    except Exception as e:
        print (e)
        return 0
        check_code = 12
        report_error_msg(error_msg[check_code],error_msg[check_code], out_path)


def meteor(gts, res):
    try:
        scorer = Meteor()
        score, scores = scorer.compute_score(gts, res)
        #print('meter = %s' % score)
        return score
    except Exception as e:
        print (e)
        return 0
        check_code = 13
        report_error_msg(error_msg[check_code],error_msg[check_code], out_path)

def rouge(gts, res):
    try:
        scorer = Rouge()
        score, scores = scorer.compute_score(gts, res)
        # print('rouge = %s' % score)
        return score
    except Exception as e:
        print (e)
        return 0
        check_code = 14
        report_error_msg(error_msg[check_code],error_msg[check_code], out_path)
 

# 辅助函数
##　解压
def unzip_file(zfile_path, unzip_dir):
    try:
        with zipfile.ZipFile(zfile_path) as zfile:
            zfile.extractall(path=unzip_dir)
    except zipfile.BadZipFile as e:
        print (zfile_path+" is a bad zip file ,please check!")
        check_code = 1
        report_error_msg(error_msg[check_code],error_msg[check_code], out_path)

# 检查关键字
def check_keyword(tmp):
    print('result' not in tmp.keys(),'333333333333')
    if('result' not in tmp.keys()):
        check_code = 3
        print('3333333222')
        report_error_msg(error_msg[check_code],error_msg[check_code], out_path)
    elif(not len(tmp['result'])):
        check_code = 4
        report_error_msg(error_msg[check_code],error_msg[check_code], out_path)
    else:
        resList = tmp['result']
        return convert2Obj(resList)

# list转为对应
def convert2Obj(resList):
    obj={}
    for i in range(len(resList)):
        obj[resList[i]['video_id']]=[resList[i]['caption']]
    return obj

def checkValid(submitObj, standardObj):
    flag = True
    for key in submitObj:
        if(key not in standardObj.keys()):
            flag = False
    return flag

if __name__=="__main__":
    '''
      online evaluation
      
    '''
    in_param_path = sys.argv[1]
    out_path = sys.argv[2]

    # 自定义变量区
    ## 假设答案命名为a.json, 用户提交为b.json，命名内容为zip包自定义，应该前期约束
    standard_file = 'standard.json'
    submit_file = 'result.json'


    # read submit and answer file from first parameter
    with open(in_param_path, 'r') as load_f:
        input_params = json.load(load_f)

    # 标准答案路径
    standard_path=input_params["fileData"]["standardFilePath"]
    print("Read standard from %s" % standard_path)

    # 选手提交的结果文件路径
    submit_path=input_params["fileData"]["userFilePath"]
    print("Read user submit file from %s" % submit_path)

    # 解压
    # unzip_file(standard_path, './')
    # unzip_file(submit_path, './')


    # 读取文件，是否是json
    try:
        with open(standard_file, 'r') as load_f:
            standard_tmp = json.load(load_f)
        with open(submit_file, 'r') as load_f:
            submit_tmp = json.load(load_f)
    except Exception as e:
        check_code = 2
        report_error_msg(error_msg[check_code],error_msg[check_code], out_path)

    standard_tmp1 = {
        "version": "VERSION 1.0",
        "result": [
            {
                "video_id": "cooking001",
                "caption": "It is in a kitchen. A person in black overall is standing by a counter.There are two broken eggs in a bowl.A ceramic bowl is full of sugar.And a smaller one is full of salt.Beside the person is towel.The person is mixing flour with eggs in a silver bowl.She is using a whisk to beat the mixture"
            },
            {
                "video_id": "cooking002",
                "caption": "It is a cooking show named cooking with dogs. There is a furry dog barking at the beginning.An Asian woman appears in a kitchen with the dog sitting by.The woman bows and starts talking .There are list of ingredients presented of the table.The kitchen has a microwave,kettle,oven and various equipments and ingredients"
            },
            {
                "video_id": "cooking003",
                "caption": "A chef is in a kitchen.He is in front of a white surface counter with a knife in hand.There is a piece of beef on the table which he is ready to cut.He slices two pieces and places them equally on top of each other.He then cleans the knife with a towel and put the knife down.There is a bowl of water beside him with ceramic plates."
            },
            {
                "video_id": "cooking004",
                "caption": "There is a wooden surface board.A person is mincing leaves.He is holding a sharp big knife.He is chopping very fast."
            },
            {
                "video_id": "cooking005",
                "caption": "A woman in a white and red shirt with green apron is standing by a big pan. She is adding some vegetables and spices to those already prepared noodles. She is using two spatulas,one in each hand,to get a uniform mix of noodles,vegetables and spices.Then she adds some soy,water and salt into the noodles and continues cooking"
            },
            {
                "video_id": "cooking006",
                "caption": "A sunny-side up egg is seen in a frying pan with oil.The egg is cooked partially and still frying.Spices has been sprinkled on the egg surface."
            },
            {
                "video_id": "cooking007",
                "caption": "A chef is adding some cooking oil and using a spatula to stir and fry squid and broccoli. The cuisine is served on a white plate.A pink flower is added to the cuisine as decoration.The plate is rotated to present this dish."
            },
            {
                "video_id": "cooking008",
                "caption": "Oil in a frying pan is boiling.Meat has been fried in the oil to a brownish state.The meat is keeping frying."
            },
            {
                "video_id": "cooking-batch2-001",
                "caption": "A stainless fry pot is placed on a small gas stove which is on a wooden surface table.A tablespoon of cooking oil is put into the pot as well as a bowl of chicken.A person then turns on the stove and starts to stir the chicken with a wooden spoon.The color of the chicken in the pot turns yellow as the person stirs.The video seems to be a cook show where instructions or guidelines are written on the left side of the video in English,Japanese and Chinese."
            },
            {
                "video_id": "education001",
                "caption": "Someone is typing on a laptop.The laptop is on a wooden surface.The left side of the laptop alone can be seen with the various slots."
            },
            {
                "video_id": "education002",
                "caption": "A man in a brown jacket and a black hat is standing in a library.There are various books available.A woman is sitting and holding a phone.He is reading a book.He later glances at the window.Someone is seen walking on the street by the window."
            },
            {
                "video_id": "education003",
                "caption": "A person holding a pencil is writing something.He makes a mistake and erases what he wrote using the eraser at the other end of the pencil.He then clears the residue."
            },
            {
                "video_id": "education004",
                "caption": "A teacher is cleaning a green backboard.She cleans it with a yellow duster.She is in a classroom.She is wearing a white and black dress."
            },
            {
                "video_id": "education005",
                "caption": "A girl with red finger nails presses the start button on her I-pod.She is reading a book written in English.She is sitting outside with sunshine penetrating the leaves.She turns the leaflet as she reads"
            },
            {
                "video_id": "education-batch2-001",
                "caption": "There is a classroom with a green chalkboard.  A middle-aged white male wears a blue suit,a blue tie and black spectacles.  He seems to be the teacher in the classroom.An Asian student in a sweater is standing in front of the class next to the board with a chalk in his hand.The teacher writes with his right hand a complex equation on the board for the student to solve.The student is however in quick response in solving and makes an attempt.On the front desk is a book and a mouse pad."
            },
            {
                "video_id": "gaming001",
                "caption": "The video is taken from PetTube.com.There is an old white man wearing a white top,blue shorts and white shoes.There is also a black dog. The man is playing basketball with the dog. There is a short basketball hoop.The dog bounces the ball and passes the ball to the man.The man then dribbles the ball for a while and passes the ball back to the dog.The dog catches the ball and dunks it into the rim.The man claps and celebrates,then they keep on playing."
            },
            {
                "video_id": "gaming002",
                "caption": "The video starts with two kids playing on a trampoline.The trampoline is in front of a house surrounded by trees.The video then goes to another scene where two kids are jumping on a trampoline with a safety net around it.The kid wearing a white top is doing flips continuously.The video ends with a young boy jumping backwards off the edge of a trailer loaded with sand and stones. "
            },
            {
                "video_id": "gaming003",
                "caption": "A woman wearing a blue dress is playing a jigsaw puzzle.The woman is in a room.The items are placed on a white table.There is a round mirror in the background and a grey sofa in the room."
            },
            {
                "video_id": "gaming004",
                "caption": "Two people are playing chess.There are two different sets of wooden chess pieces.One of the players makes a move and then stops his clock. The other writes something on a paper.One wears a black shirt and the other wears a brown top.The video has a white background."
            },
            {
                "video_id": "gaming005",
                "caption": "Three children are playing in a sandpit.One of the three children only wears a hat and shorts.He starts to climb a slide with two boomerangs in each hand.One boomerang falls as he climbing.An older kid snatch the boomerang,and runs away with another kid. The younger kid starts screaming.The older kid throws the boomerang back.The younger kid picks it up and runs away."
            },
            {
                "video_id": "gaming006",
                "caption": "A group of people are playing cards around a wood surface table in a room. They sit around the table with opened drinks placed around them.A pack of cards is placed in the middle.They wait in turns to pick a card for each to keep.  One is smoking. One has a tattoo on the hand."
            },
            {
                "video_id": "gaming-batch2-001",
                "caption": "The are several people having fun on rides. There are a few trees planted beside the rides. A group of peoples are sitting on a sling chair as it shoots in the air. The seat is connected by strings to two big trusses that secures it from falling.After they are shot in air,they rock back and forth. "
            },
            {
                "video_id": "movie001",
                "caption": "There is a country road with trees and fields at each side. A green and white bus which is full of passengers passes by.Some passengers are sitting on top of the bus beside their luggage.The bus is moving forward at a steady pace. The passengers on the bus seem to be singing and enjoying themselves."
            },
            {
                "video_id": "movie002",
                "caption": "A man is dressed in black with short hair. He holds a camera and keeps staring at the camera screen.He is taking video with both hands holding the camera steadily."
            },
            {
                "video_id": "movie003",
                "caption": "There is a Sony video recorder.Someone walks by in the background.Some people sit at a table.The recorder is set at steady position and recording."
            },
            {
                "video_id": "movie-batch2-001",
                "caption": "It is at on a mountain top.A movie crew is recording a movie scene on the top of a mountain.A group of cameramen are holding video recorders recording a video cautiously.There is a girl in a cream satin dress being recorded with her face covered by a satin veil.The wind is blowing against her dress as she moves.She is moving gracefully like a ballet dancer.  "
            },
            {
                "video_id": "music001",
                "caption": "There is an old man wearing a black suit.He is sitting by a white piano.He is playing the piano with his two hands.He is in a small room with white curtains.He sways his head and body as he plays the music."
            },
            {
                "video_id": "music002",
                "caption": "A young white male is in a room.There are books and some personal belongings on a shelf.He is playing a guitar while sitting on a chair.He is dressed casually.He changes the tune of the guitar as he plays.He is focused on the instrument."
            },
            {
                "video_id": "music003",
                "caption": "This is some form of reality show.A middle-aged Indian man is playing a brown violin.He wears a white top and has short dark hair.After he finishes a verse,a middle-aged Indian woman starts singing.She is holding a microphone with her right hand.The video is from Asia net."
            },
            {
                "video_id": "music004",
                "caption": "There is a number of chairs and tables in the room as well as a shelf filled with bottles and books.There are monitors on top of the tables.There is a suitcase and a sound box placed on the floor with a guitar on top of the sound box.A casually dressed young man walks in the room.He picks up the guitar and hangs it on him.He starts playing the guitar.He then places his right leg on the suitcase while playing."
            },
            {
                "video_id": "music005",
                "caption": "There is a rock show in a room.The room is bit dark.There are stage lights swiveling around to create a disco effect.A lot of people are dancing and clapping while listening to the music.There is a stage with musicians holding guitars and playing songs.There are also standing microphones which the musician are singing into.There room however is very crowded."
            },
            {
                "video_id": "music006",
                "caption": "A man is sitting down on a chair.He is casually dressed.He is playing a brown guitar which is on his lap.He changes the tune of the guitar as he plays."
            },
            {
                "video_id": "music007",
                "caption": "A white male is holding a pair of drum sticks in a dark room.He is playing drums.A group of people are sitting and standing around him,watching his performance.Among the audience is a small boy sitting on a chair.The boy later stands up."
            },
            {
                "video_id": "music008",
                "caption": "A middle-aged white man is standing on a big stage.He is singing into a standing microphone. His eyes are closed while singing.He seems to be playing a guitar while singing."
            },
            {
                "video_id": "music-batch2-001",
                "caption": "A young Asian girl with a ponytail in casual wear is sitting on a stool.She is wearing a pair of glasses,a black top,denim overall shorts and black socks with white dots.She is holding a set of drumsticks and playing a transparent drum set.Different individuals are sitting around listening to the music.Some are standing around on the stairs.And the girl is playing the drum set rhythmically with earphones in her ears.She sways her body and enjoys the music as she plays."
            },
            {
                "video_id": "sports001",
                "caption": "A lot of people in a room. They gather around circularly around a stage.A young white skin,red hair,casually dressed male is playing football.Everyone else’s eyes are on him and some are clapping.He is dribbling a ball with both legs.He is making nice moves."
            },
            {
                "video_id": "sports002",
                "caption": "A young white woman is in a bikini and a life jacket.She is water skiing on a water surface.There are hills surrounding the water.She is sliding right-angled on the surface.She is holding tightly to the rope for support."
            },
            {
                "video_id": "sports003",
                "caption": "A woman is riding a horse in a stable.She is wearing horse riding clothes.The stable is a wooden structure.The horse is a grey and white horse.There are sightings of trees and a few people standing around."
            },
            {
                "video_id": "sports004",
                "caption": "There is a view of the ocean.A couple of white males are sailing a catboat with blue sail.One of them is at the right edge of the boat holding on to a string attached to the boat.They are riding vehemently.The name of the boat is Sekonda.Buildings can be seen far away as they sail."
            },
            {
                "video_id": "sports005",
                "caption": "A woman is standing wearing gym clothes and doing exercise.She is holding a pair of dumbbells.She is alternately raising them once at a time.A man who seems to be the gym instructor is guiding her as she exercises.She is looking directly ahead with serious expression."
            },
            {
                "video_id": "sports006",
                "caption": "A middle-aged man is doing exercises at the gym.He is in gym clothes.He is doing press-ups with his hands.In the gym there are several punching bags,boxing rings and other gym gadgets. At the end of the video,the man had done six press-ups.There are other people seem doing other forms of exercise in the far distance."
            },
            {
                "video_id": "sports007",
                "caption": "A man dressed in black clothes and grey shoes is skating in the alleyway.He is on the road between buildings.He is skating with both legs.He later tried to maneuver the skating board with leg. He lands right back on it."
            },
            {
                "video_id": "sports008",
                "caption": "A bare-chested man with only shorts is doing exercise.Only his back is shown and he has short hair.He is doing pull-up exercises with a pull-up bar.He is grunting as he is doing the pull-ups.He has a tattoo on his back.The bar is between walls which hold the bar on each side."
            },
            {
                "video_id": "sports009",
                "caption": "It is at a basketball court.People are watching the game being played.There is a projector that projects what is happening.A man in sportswear dunks the basketball into the basket.He is in a green shirt and black shorts.The court has a blue background with lights around. A man is seen shooting videos in the projection."
            },
            {
                "video_id": "sports010",
                "caption": "There are green grasses and trees in the surroundings.A young guy in a white T-shirt is holding a football.He passes the football to another individual who is standing in a distance.They seem to be playing the game of throw and catch.The other guy is in a blue shirt with patterns on it.A car is parked in the distance with an individual running in the background."
            },
            {
                "video_id": "sports011",
                "caption": "There is a game on a basketball court.The court has a wooden tiled floor.Two teams are playing a heated match.One team is in red and the other in white jerseys. There are a lot of people watching the game.One player wearing a white jersey is using a crossover to pass Michael Jordan in a red jersey.He then shoots and scores.He put the ball into the net."
            },
            {
                "video_id": "sports-batch2-001",
                "caption": "There is a lighted well-equipped basketball court with wooden floors. The warriors’ symbol can be seen on the wall.The basketball court is in black and white paintings.Two males dressed in sport wear are playing a brown basketball.They are bouncing and throwing the basketball around.Another individual with a black hat is standing around watching the game.One player is wearing an ash top and the other is in a white top.They are playing one-on-one and they both make some shots."
            },
            {
                "video_id": "travel001",
                "caption": "It is at the beach side.,There are waves splashing in the ocean,Beach sand and stones can be seen,A shadow of a man appears slowly,The man walks by with only his lower body showing,He is wearing shorts and a shirt"
            },
            {
                "video_id": "travel002",
                "caption": "The video is a timeline,The wind is blowing tenderly,The clouds are moving,There are numerous buildings,There is a tall tower,The video is taken in a city"
            },
            {
                "video_id": "travel003",
                "caption": "Numerous people are sitting by the beach,Others are walking around,They are all in beach clothes,The video was taken at by a palm tree,The camera was angled under a branch of the tree,The video is vague,There are a few buildings around the beach with flags waving"
            },
            {
                "video_id": "travel004",
                "caption": "It is a sunny day,There are numerous tall buildings and houses around,There are numerous electrical transformer sited and green grass everywhere,There is a hill or mountain nearby with building at its edge,A red and white train runs from a distance,The train gradually pass by,It goes away leaving behind the quiet and peace,"
            },
            {
                "video_id": "travel005",
                "caption": "The video was taken at the top of a building,The waves in the ocean row calmly,With a tree in site and some cars moving,There are people moving around the beach,A light skin woman appears,holding a camera,She angles the camera and take pictures of the scenery,She leaves after taking some pictures"
            },
            {
                "video_id": "travel-batch2-001",
                "caption": "There is a rail road and a platform surrounded by many cherry trees with blooming pink flowers,A white and blue train is passing by with passengers in it,The train is moving slowly,There is a man in pink trousers squatting and taking a picture of the scene,The number on the train is MR-613,There is also green grass at each side of the railway"
            },
            {
                "video_id": "travel-batch2-002",
                "caption": "There are a group of deer galloping by in a park,There are several individuals surrounding the deer in the park,Some are taking pictures,some are feeding the deer and others are just watching,The deer seem to be indifferent to the individuals around them and are eating their meal,A group of deer are gathered around a man who is feeding them and hoping to get some of what he is offering"
            },
            {
                "video_id": "tvshows001",
                "caption": "This is a television show,The name of the television show is Animal Planet,A young lady holding a microphone is interviewing a parrot,The parrot is grey in color and is sitting on a wooden bar,The parrot starts talking,Three people consisting of two women and a man are seated behind a table watching the interview,After the parrot finishes talking they start laughing,The stage has been designed with stars and ornaments,The name of the woman wearing a black top is Melissa and the name of the man in the middle is Chris"
            },
            {
                "video_id": "tvshows002",
                "caption": "It is television show from natgeotv.com,It is snowing,A brown bear is with a cub in the wild,Two wolves are trying to attack the cub,The grown bear is protecting the cub and tries to scare the wolves off,The bear chases off the wolves and comes back to the cub"
            },
            {
                "video_id": "tvshows003",
                "caption": "This is an animation with green background,There is an animated man in black suits and white shirt with a camera around his neck,He is fair in complexion with short hair,He is wearing glasses and a necklace,He seems to be dancing with all of his body parts"
            },
            {
                "video_id": "tvshows004",
                "caption": "There are mountains in the background,A young white female is wearing a red top,a scarf,a white hat and a pair of jeans,She is holding a phone,She is smiling at the camera with open arms"
            },
            {
                "video_id": "tvshows005",
                "caption": "Two bikers in racing suits are riding their motorcycles,They ride at a very fast speed and then fly into the air at the same time,Both of them rotate 360 degrees and stretch their body with their bikes in the air synchronously and land on the ground successfully,There are ramps assisting the bikers to take off and land on ground"
            },
            {
                "video_id": "tvshows006",
                "caption": "A ballerina wearing a white tutu is dancing on the stage,She glides and turns as she dances,She spins around and then stands on her toes,She seems to be in a grand theater"
            },
            {
                "video_id": "tvshows007",
                "caption": "This is a television show from natgeotv.com,There is a large group of penguins moving together,They all seem to be similar to each other,They are all going at a slow and uniform pace,They are surrounded by snow and ice"
            },
            {
                "video_id": "tvshows-batch2-001",
                "caption": "A middle-aged man is sitting on a chair at a news studio,He is wearing a black patterned suit with a white shirt,a blue patterned tie and spectacles,His hands are folded together on his lap as he reads the news,The name of the channel is CCTV 13,The time in the video is 22:34,He seems to discuss a topic about nearsightedness,There is a microphone connected to his suit"
            },
            {
                "video_id": "vehicles001",
                "caption": "It is at a racing track,There are painted tires arranged alongside the tracks to guide the drivers,Two race cars drift by following each other,The stunt drivers are making a fast curve one after another and emitting plenty of smoke,A group of people are situated by on stands watching the race,A group on people are on the track assisting,"
            },
            {
                "video_id": "vehicles002",
                "caption": "It is at a car park along a seashore,A Mercedes Benz with plates ET 3633P s making a test drive,There are two people in the car,It is form of advertisement for a car,He make turns while driving,The car finally stops near the seashore with big GLA logo words"
            },
            {
                "video_id": "vehicles003",
                "caption": "There is a road leading to a light house,The road leads upward towards the hill,There are black and yellow polls along the road guiding drivers,A white van passes by and drives to the hill top"
            },
            {
                "video_id": "vehicles004",
                "caption": "A black BMW convertible car is parked on the road,There are grasses and trees along the road,The car top has been folded up,The folded car top is arranged back on top of the car"
            },
            {
                "video_id": "vehicles005",
                "caption": "It is a car park,There are a few decorative plants in the surrounding and white building,A man casually dressed is seen pushing a black saloon car alone from behind,He is in a white tee shirt,blue shorts and sneakers,He is using all his might,A young boy is seen sitting in the car directing it,He seems frustrated"
            },
            {
                "video_id": "vehicles-batch2-001",
                "caption": "There is a winding freeway in the middle of a rocky desert in sight,The road is surrounded by high rocks and trees,There is a red and black sports car on the road driving at a top speed,The video is advertising the car’s speed and other qualities"
            },
            {
                "video_id": "working001",
                "caption": "A man in green clothes is seen with a drilling machine,He is drilling nail into the wooden wall,He is guiding it with his hand,He is building a place for microwave oven,And it is finally built"
            },
            {
                "video_id": "working002",
                "caption": "There is a wooden surface table and a glass of drink on it,Someone is working on the table,He is typing something on keyboard and clicks with a mouse,He then writes something in a diary with a pen with his left hand,Then he takes a sip of the drink"
            },
            {
                "video_id": "working003",
                "caption": "A women sits on a cushion table with her legs crossed,There is a still green fun in the room,She is reading and typing something on a laptop,The laptop is on a table,She has a pillow at her back,After a while she relaxes"
            },
            {
                "video_id": "working004",
                "caption": "A woman with black polished nails is drawing on a wooden surface area,She is drawing a pro-type of a phone in a book pad,There is a glass of drink on the table,And an folded earpiece on the table,There are various markers on the table"
            },
            {
                "video_id": "working005",
                "caption": "A person measures a piece of wood to draw a line with a pencil,He then cuts with a cutting machine,He holds the piece of wood and cuts the part that is not needed"
            },
            {
                "video_id": "working006",
                "caption": "A person is standing by a white board,He is holding a black marker,He is writing or drawing something on the board,He is wearing a black jacket"
            },
            {
                "video_id": "working007",
                "caption": "It is a courtyard with grass and trees planted,An old man wears casual clothes and a pair of sandals,He is mowing the grass in front of a house"
            },
            {
                "video_id": "working-batch2-001",
                "caption": "A girl with black painted nails is typing on a black Samsung laptop with both hands,She seems to be wearing a furry gray top,She has a silver ring on her left ring finger,She then shows how to type in words with the left little finger pressing the ‘SHIFT’ button at the same time,There are some Chinese subtitles at the bottom of the video"
            }
        ],
        "external_data": {
            "used": "true",
            "details": ""
        }
    }
    submit_tmp2 = {
        "version": "VERSION 1.2",
        "result": [{
            "video_id": "cooking001",
            "caption": "a man is mixing eggs"
        }, {
            "video_id": "cooking002",
            "caption": "a woman is barking"
        }, {
            "video_id": "cooking003",
            "caption": "a man is slicing a piece of meat"
        }, {
            "video_id": "cooking004",
            "caption": "a person is chopping onions"
        }, {
            "video_id": "cooking005",
            "caption": "a man is cooking his kichen"
        }, {
            "video_id": "cooking006",
            "caption": "a person is cooking"
        }, {
            "video_id": "cooking007",
            "caption": "a man is cooking his kichen"
        }, {
            "video_id": "cooking008",
            "caption": "a person is cooking"
        }, {
            "video_id": "cooking-batch2-001",
            "caption": "a person is cooking"
        }, {
            "video_id": "education001",
            "caption": "a person is typing on a keyboard"
        }, {
            "video_id": "education002",
            "caption": "a man is looking at a window"
        }, {
            "video_id": "education003",
            "caption": "a person is slicing a potato"
        }, {
            "video_id": "education004",
            "caption": "a woman is swimming"
        }, {
            "video_id": "education005",
            "caption": "a man is playing a card"
        }, {
            "video_id": "education-batch2-001",
            "caption": "a man is writing on a chalkboard"
        }, {
            "video_id": "gaming001",
            "caption": "a dog is playing with a basketball"
        }, {
            "video_id": "gaming002",
            "caption": "a man is jumping on a trampoline"
        }, {
            "video_id": "gaming003",
            "caption": "a man is cutting a paper"
        }, {
            "video_id": "gaming004",
            "caption": "a man is playing with a toy"
        }, {
            "video_id": "gaming005",
            "caption": "a man is playing with a toy"
        }, {
            "video_id": "gaming006",
            "caption": "a person is playing a keyboard"
        }, {
            "video_id": "gaming-batch2-001",
            "caption": "a plane is flying in the sky"
        }, {
            "video_id": "movie001",
            "caption": "a group of people are riding a bus"
        }, {
            "video_id": "movie002",
            "caption": "a man is shooting a gun"
        }, {
            "video_id": "movie003",
            "caption": "a man is shooting a gun"
        }, {
            "video_id": "movie-batch2-001",
            "caption": "a man is shooting a gun"
        }, {
            "video_id": "music001",
            "caption": "a man is playing a piano"
        }, {
            "video_id": "music002",
            "caption": "a man is playing a guitar"
        }, {
            "video_id": "music003",
            "caption": "a man is playing a violin"
        }, {
            "video_id": "music004",
            "caption": "a man is playing a guitar"
        }, {
            "video_id": "music005",
            "caption": "a band is performing on stage"
        }, {
            "video_id": "music006",
            "caption": "a man is playing a guitar"
        }, {
            "video_id": "music007",
            "caption": "a man is playing the drums"
        }, {
            "video_id": "music008",
            "caption": "a man is singing into a microphone"
        }, {
            "video_id": "music-batch2-001",
            "caption": "a man is playing drums"
        }, {
            "video_id": "sports001",
            "caption": "a man is kicking a ball"
        }, {
            "video_id": "sports002",
            "caption": "a woman is riding a boat"
        }, {
            "video_id": "sports003",
            "caption": "a woman is riding a horse"
        }, {
            "video_id": "sports004",
            "caption": "a man is riding a boat"
        }, {
            "video_id": "sports005",
            "caption": "a woman is exercising"
        }, {
            "video_id": "sports006",
            "caption": "a man is doing exercise"
        }, {
            "video_id": "sports007",
            "caption": "a man is doing a skateboard"
        }, {
            "video_id": "sports008",
            "caption": "a man is doing exercise"
        }, {
            "video_id": "sports009",
            "caption": "a man is playing with a ball"
        }, {
            "video_id": "sports010",
            "caption": "a man is playing football"
        }, {
            "video_id": "sports011",
            "caption": "a group of men are playing basketball"
        }, {
            "video_id": "sports-batch2-001",
            "caption": "a man is playing with a ball"
        }, {
            "video_id": "travel001",
            "caption": "a boy is sitting on the beach"
        }, {
            "video_id": "travel002",
            "caption": "a man is skating"
        }, {
            "video_id": "travel003",
            "caption": "a man is walking on a street"
        }, {
            "video_id": "travel004",
            "caption": "a man is riding a horse"
        }, {
            "video_id": "travel005",
            "caption": "a woman is sitting on the water"
        }, {
            "video_id": "travel-batch2-001",
            "caption": "a bus is going on the road"
        }, {
            "video_id": "travel-batch2-002",
            "caption": "a group of people are walking"
        }, {
            "video_id": "tvshows001",
            "caption": "a man is talking into a microphone"
        }, {
            "video_id": "tvshows002",
            "caption": "a bear is chasing a bear"
        }, {
            "video_id": "tvshows003",
            "caption": "a man is speaking"
        }, {
            "video_id": "tvshows004",
            "caption": "a man is brushing a rope"
        }, {
            "video_id": "tvshows005",
            "caption": "a man is performing stunts on a motorcycle"
        }, {
            "video_id": "tvshows006",
            "caption": "a man is performing a rope"
        }, {
            "video_id": "tvshows007",
            "caption": "a group of penguins are walking"
        }, {
            "video_id": "tvshows-batch2-001",
            "caption": "a man is giving a speech"
        }, {
            "video_id": "vehicles001",
            "caption": "a man is riding a motorcycle"
        }, {
            "video_id": "vehicles002",
            "caption": "a man is driving a car"
        }, {
            "video_id": "vehicles003",
            "caption": "a man is running"
        }, {
            "video_id": "vehicles004",
            "caption": "a man is driving a car"
        }, {
            "video_id": "vehicles005",
            "caption": "a man is driving a car"
        }, {
            "video_id": "vehicles-batch2-001",
            "caption": "a man is riding a bike"
        }, {
            "video_id": "working001",
            "caption": "a man is cutting a piece of wood"
        }, {
            "video_id": "working002",
            "caption": "a man is slicing a keyboard"
        }, {
            "video_id": "working003",
            "caption": "a woman is cleaning a bed"
        }, {
            "video_id": "working004",
            "caption": "a person is slicing a carrot"
        }, {
            "video_id": "working005",
            "caption": "a man is cutting a piece of cucumber"
        }, {
            "video_id": "working006",
            "caption": "a man is stirring a cigarette"
        }, {
            "video_id": "working007",
            "caption": "a man is playing with a sword"
        }, {
            "video_id": "working-batch2-001",
            "caption": "a man is typing on a keyboard"
        }],
        "external_data": {
            "used": "true",
            "details": "First fully-connected layer from VGG-16 pre-trained on ILSVRC-2012 training set"
        }
    }

    # 检查格式，是否有对应字段
    standard_content = check_keyword(standard_tmp)
    submit_content = check_keyword(submit_tmp)

    # 进行算法
    totalS = 0
    totalS1 = 0
    totalS2 = 0
    totalS3 = 0
    totalS4 = 0
    count = len(standard_tmp['result'])
    try:
        if bool(submit_content):
            for key in submit_content:
                print(key)
                submit_text = submit_content[key]
                standard_text = standard_content[key]
                print(submit_text,'11111')
                print(standard_text,'222222')
                scores,extra = get_score(standard_content,submit_content)
                totalS = totalS + scores
                totalS1 = totalS1 + extra['bleu']
                totalS2 = totalS2 + extra['cider']
                totalS3 = totalS3 + extra['meteor']
                totalS4 = totalS4 + extra['rouge']
                print(key,round(scores,3),extra['bleu'],extra['cider'],extra['meteor'],extra['rouge'])
            scores = round(totalS / count,3)
            extra = {}
            extra['bleu']= round(totalS1 / count,3)
            extra['cider']= round(totalS2 / count,3)
            extra['meteor']= round(totalS3 / count,3)
            extra['rouge']= round(totalS4 / count,3)
            report_score(scores,extra, out_path)
    except Exception as e:
        report_error_msg(str(e),str(e), out_path)

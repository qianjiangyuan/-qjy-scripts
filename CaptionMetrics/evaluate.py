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
    14: 'rouge'
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
    s1  = bleu(gts, res)
    s2 = cider(gts, res)
    s3 = meteor(gts, res)
    s4 = rouge(gts, res)
    obj = {'s1':s1, 's2':s2, 's3':s3, 's4':s4}
    total = s1+s2+s3+s4
    return total, obj
    


def bleu(gts, res):
    try:
        scorer = Bleu(n=4)
        score, scores = scorer.compute_score(gts, res)
        # print('belu = %s' % score)
        return score[3]
    except Exception as e:
        print (e)
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
        check_code = 13
        report_error_msg(error_msg[check_code],error_msg[check_code], out_path)

def rouge(gts, res):
    try:
        scorer = Rouge()
        score, scores = scorer.compute_score(gts, res)
        #print('rouge = %s' % score)
        return score
    except Exception as e:
        print (e)
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
    if('result' not in tmp.keys()):
        check_code = 3
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


if __name__=="__main__":
    '''
      online evaluation
      
    '''
    in_param_path = sys.argv[1]
    out_path = sys.argv[2]

    # 自定义变量区
    ## 假设答案命名为a.json, 用户提交为b.json，命名内容为zip包自定义，应该前期约束
    standard_file = 'a.json'
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
    unzip_file(standard_path, './')
    unzip_file(submit_path, './')


    # 读取文件，是否是json
    try:
        with open(standard_file, 'r') as load_f:
            standard_tmp = json.load(load_f)
        with open(submit_file, 'r') as load_f:
            submit_tmp = json.load(load_f)
    except Exception as e:
        check_code = 2
        report_error_msg(error_msg[check_code],error_msg[check_code], out_path)


    # 检查格式，是否有对应字段
    standard_content = check_keyword(standard_tmp)
    submit_content = check_keyword(submit_tmp)

    # 进行算法
    try:
        scores,extra = get_score(standard_content,submit_content)
        report_score(scores,extra, out_path)
    except Exception as e:
        # NOTICE: 这个只是示例
        check_code = 1
        report_error_msg(error_msg[check_code],error_msg[check_code], out_path)

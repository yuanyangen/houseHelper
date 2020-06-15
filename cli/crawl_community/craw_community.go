package main

import (
	"fmt"
	"strings"
)

type Info struct {
	Section string
	SchoolName string
	Level string
}


func main () {
	tmpData := strings.Split(str, "\n")
	level := ""
	sectionName := ""
	ret := []*Info{}
	for _,v := range tmpData {
		if v == "" {
			continue
		}
		if strings.Contains(v, "-") {
			tt := strings.Split(v, " ")
			level = tt[0]
			if len(tt) == 2 {
				sectionName = tt[1]
			}
			continue
		}
		ret = append(ret, &Info{
			Section:    sectionName,
			SchoolName: v,
			Level:      level,
		})

	}
	for _,v := range ret {
		fmt.Printf("%v %v %v\n", v.Section,  v.Level, v.SchoolName)
	}
}










var str = `1-1 东城
府学小学
史家小学
景山学校
1-2
光明小学
和平里九小
和平里四小
黑芝麻胡同小学
一师附小
2-1
西中街小学
灯市口小学
崇文小学
培新小学
东交民巷
史家小学分校
板厂小学
景泰小学
2-2
地坛小学
东四九条小学
东师附（与和平里四小联合招生）
曙光小学
前门小学
美术馆后街小学
1-2 朝阳
朝阳实验小学
芳草地国际学校
2-1
陈经纶帝景
朝阳外国语学校
白家庄小学
呼家楼中心小学
朝师附小
陈经纶嘉铭分校
星河实验小学
花家地实验小学
2-2
安慧里中心小学
南湖东园小学
劲松四小
南湖中园小学
垂杨柳中心小学
八里庄中心小学
1-1 海淀
中关村第三小学
中关村第一小学
中关村第二小学
人大附小
人大附中实验小学
1-2
海淀实验小学
五一小学
上地实验小学
翠微小学
石油附小
北师大附小
北大附小
2-1
林大附小
科大附小
北理工附小
中关村第四小学
万泉小学
七一小学
北航附小
农科院附小
二里沟中心小学
北医附小
育英小学
育新学校小学部
海淀实验二小
清华附小
建华实验
海淀外国语
2-2
玉泉小学
羊坊店中心小学
羊坊店第四小学
太平路小学
育鹰小学
首师大附小
永泰小学
交大附小
育鸿学校
九一小学
花园村第二小学
双榆树中心小学
羊坊店第五小学
立新小学
图强第二小学
西苑小学
海淀民族小学
今典小学
双榆树第一小学
群英小学
红英小学
培英小学
北洼路小学
八里庄小学
彩和坊小学
西颐小学
学府苑小学
北外附校
海淀第三实验小学
海淀第四实验小学
巨山小学
田村小学
前进小学
六一小学
清华东路小学
1-1 西城
北京实验二小
育民小学
育翔小学
西师附小
三里河三小
1-2
五路通小学
北京小学本部
宏庙小学
北京实验一小
奋斗小学
黄城根小学
中古友谊小学
2-1
自忠小学
宣师一附小
复兴门外一小
白云路小学
康乐里小学
展览路一小
阜成门外一小
育才学校
力学小学
裕中小学
德外二小
北长街小学
顺城街一小
厂桥小学
玉桃园小学
西单小学
北礼士路一小
西什库小学
进步小学
四根柏小学
北京小学走读部
北京小学广外分校
2-2
鸦儿小学
半步桥小学
香厂路小学
华嘉小学
柳荫街小学
西四北四条小学
雷锋小学
中华路小学
文兴街小学
新街口东街小学
什刹海小学
`
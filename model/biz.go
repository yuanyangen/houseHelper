package model

type HouseInfo struct {
	*LianjiaBaseHouseInfo
	*LianjiaTradeInfo
	*StateInfo
	CrawlDate string
	ModifyTime int64
}

type HouseDealInfo struct {
	HouseId int64 //房子ID
	RawPrice float64 // 原始挂牌价格， 单位 万
	DealPrice float64 // 成交价格，单位 万
	DealTime int64 // 成交时间， unix时间戳
	ModifyTime int64 //
	Duration int64 // 成交周期， 单位天
	CrawlDate string
}

const WarnSupplyTypeSelf = 1
const WarnSupplyTypeCentral = 2

const DecorateTypeGood = 1
const DecorateTypeNormal = 2
const DecorateTypeNone = 3

type LianjiaBaseHouseInfo struct {
	Id             int64
	Type           string // 2-2-1-1
	FloorNumber    string // low (15)
	BuildingType      string    // 1 板楼 2 塔楼
	WarnSupplyType int    //
	DecorateType   int
	HouseUsageType string // 房屋用途 普通住宅
	CommunityName string

	ForwardType  string  // 朝向
	Elevator     bool    // 是否有电梯
	BuildingSize float64 // 建筑面积 单位m^2
	RealSize     float64 // 套内面积 单位m^2
	StandardType string // 户型结构， 平层
	BuildingStruct string // 建筑结构， 刚混结构
	FloorToUseCount string // 梯户比例， 1梯2户
	StandardUrl  string  // 户型图地
}

const YearTypeLessOther = 0
const YearTypeGreatThanTwo = 2
const YearTypeGreatThanFive = 5


type BelongType int

const (
	Public  BelongType = 1 // 共有产权
	Private BelongType = 2 // 非共有产权
)

type StateInfo struct {
	WeekSeeCount  int64
	MonthSeeCount int64
	TotalSeeCount int64
	FavCount      int64
}

type LianjiaTradeInfo struct {
	TradeYearType  int
	OnlineTime     int64      // 挂牌时间
	LastTradeTime  int64      // 上一次交易时间
	TradeType           string  // 交易权属 商品房， 经济适用房
	BelongTo       BelongType //房屋产权
	Mortgage       bool       // 是否有抵押
	MortgageNumber int32      // 抵押数目， 单位是万
	OwnYearCountType int      // 产权年限

}

func NewHouseInfo() *HouseInfo {
	return &HouseInfo{LianjiaBaseHouseInfo: &LianjiaBaseHouseInfo{}, LianjiaTradeInfo: &LianjiaTradeInfo{}, StateInfo: &StateInfo{}}
}

package controllers

/*
购买这个操作需要单独提供一个服务来进行支持
使用时间驱动模型来解决库存问题

1. 库存锁定问题
2.
*/

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
)

//type ShopEvent struct {
//	//购物事件
//	User models.User
//	GoodDetail models.GoodDetail
//	Num int
//	// 当购物事件处理完成之后进行回调
//	Controller *controllers.IndexController
//}

type SecOrderResultchan struct {
	//回调时需要 这个指针
	Controller *IndexController
	Result bool
	Err error
}

type OrderResultInterface interface{
	/* 结果处理
		订单的状态会存储到channel 中
		当获取到每个channel后，
	*/

	Run(*IndexController, bool ,error)
}
type OrderResultStructBase struct {
	SecChannel chan SecOrderResultchan
}
func (result *OrderResultStructBase) Run(controller *IndexController, status bool, err error) {
	v := SecOrderResultchan{Controller: controller, Result: status, Err: err}
	result.SecChannel <- v
}


func ProduceOrderEvent(
	shop_event ShopEvent,  stock_opera StockOperation,
	orderinterface OrderInterface, result_handler OrderResultInterface) (status bool, err error){
	/*一个订单生成的流程的封装，所有的步骤应该在一个事务中去执行来保证一致性
	将库存的操作 和订单生成的操作抽象成两个步骤, 这样在秒杀和 普通活动中可以共用
	目前没有比较完整的事务的特定操作api， 错误处理和回滚操作比较简陋，
	TODO 作为挑战， 可将事务的操作 也进行抽象， 提供interface, 对整个购物流程进行抽象
	*/
	defer func() {
		result_handler.Run(shop_event.Controller, status, err)
	}()

	if ret, err := stock_opera.ReduceStock(shop_event.GoodDetail.Id, shop_event.Num); ret {
		if ret, err2 := orderinterface.ProduceOrder(shop_event); ret {
			return true, nil
		} else {
			// 注意在这里需要恢复库存量
			fmt.Println("\n\nrecover stock")
			stock_opera.RecoverStock(shop_event.GoodDetail.Id, shop_event.Num)
			return false, err2
		}
	} else {
		return false, err
	}
}

/* 事件队列的相关操作 */
type ShopEventQueueOperation interface{
	AddEvent(ShopEvent) (bool, error)
	GetEvent(num int )([]*ShopEvent, error)
	ErrorHanlder(error) //错误处理
}

type ShopEventQueue struct {
	 Num int
	 Limits int
	 Cache []*ShopEvent
	 Lock sync.Mutex
}

func (queue *ShopEventQueue) AddEvent(event *ShopEvent) (bool, error) {
	if queue.Num >= queue.Limits {
		return false, errors.New("事件队列已经超过限制")
	}
	// 在这里加锁是因为每个请求都会忘队列加数据，每个请求是goroutine，因此会存在并发的问题
	queue.Lock.Lock()
	defer queue.Lock.Unlock()
	queue.Cache[queue.Num] = event
	queue.Num = queue.Num + 1
	return true, nil
}

// GetEvent 被限制在单独的的协程中，因此不用加锁读
func (que *ShopEventQueue) GetEvent(num int) ([]*ShopEvent, error){
	var count int = num
	//que.Lock.Lock()
	//defer que.Lock.Unlock()
	if num > que.Num {
		count = que.Num
	}
	if count == 0 {
		return nil, errors.New("队列已经清空了")
	}
	ret := que.Cache[:count]
	que.Cache = que.Cache[count:]
	que.Num = que.Num - count
	return ret, nil
}

func NewShopEventQueue(limits int) *ShopEventQueue{
	cache := make([]*ShopEvent, limits)
	queue := ShopEventQueue{Num: 0, Limits: limits, Cache: cache}
	return &queue
}
/* 事件队列操作结束 */


//秒杀服务和web服务部署在同一进程，那么在秒杀服务执行完一个消息，可以立马进行http_request的回调
// 如果是跨进程进行调用，那么模型
type TradeServer struct {
	// 秒杀活动的事件队列
	SecKillQueue *ShopEventQueue
	// queue队列的长度
	QueueLength int
	// 服务并发执行的限制
	Capacity int

	//普通购买的事件队列
	ComQueue *ShopEventQueue
	//普通队列的长度
	ComqueLength int
}

var (
	TradeSer *TradeServer;
	//为秒杀活动产生的缓存信息
	GoodCache *ShopCache
)

func init() {
	//should init from Config file
	//fmt.Println("\n\ntrade service init function ")
	TradeSer = NewTradeserver(5000, 20, 1000)
	GoodCache = InitShopCache(3)
}

func NewTradeserver(queue_length int, capacity int, com_length int) *TradeServer{
	v := TradeServer{QueueLength:queue_length, Capacity: capacity, ComqueLength: com_length}
	v.SecKillQueue = NewShopEventQueue(queue_length)
	v.ComQueue = NewShopEventQueue(com_length)
	return &v
}

// 对于用户的每一次秒杀操作 对应一个独立的消息
func (trade_server *TradeServer)AddSecEvent(event *ShopEvent) (bool, error) {
	return trade_server.SecKillQueue.AddEvent(event)
}

func (*TradeServer) Run (channel_server chan bool) {

	/*秒杀活动事件循环 */
	order_service := OrderService{}

	sec_channel := make(chan SecOrderResultchan, TradeSer.Capacity)
	SecOrderHandler := OrderResultStructBase{SecChannel: sec_channel}
	//fmt.Println("\ncheck good cached ", *GoodCache.GoodMap[1])
	for {

		if events, err := TradeSer.SecKillQueue.GetEvent(TradeSer.Capacity); err == nil {
			//fmt.Println("\n\nget event", len(events), events[0])
			for _, event := range events {
				//fmt.Println("zzzz", event, GoodCache, &order_service, &SecOrderHandler)
				go ProduceOrderEvent(*event, GoodCache, &order_service, &SecOrderHandler)
			}
			// 这段代码局限在 秒杀服务 和 web服务 在同一进程内
			for index := 0; index < len(events); index ++ {
				request_result := <-SecOrderHandler.SecChannel
				request_result.Controller.SetData("json", map[int]int{19: 12})
				runtime.Gosched()
				//fmt.Println("\n\n\n\nhandle event success")
				// request_result.Controller.ServeJSON()
			}

		} else {
			//fmt.Println("\n\nerror : ", err.Error())
			runtime.Gosched()
		}
	}
	channel_server <- true

}

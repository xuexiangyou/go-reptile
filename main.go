package main

import (
	"reptile/engine"
	"reptile/persist"
	"reptile/scheduler"
	"reptile/zhenai/parser"
)

func main() {
	var concurrentEngine = engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemSave:persist.ItemSave(),
	}
	concurrentEngine.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
}

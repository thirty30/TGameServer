package logicserver

func debugCmdCallBack(aCmd string) {
	switch aCmd {
	case "statistictime":
		{
			pMonitor := gServerSingleton.getMonitor()
			if pMonitor == nil {
				return
			}
			pMonitor.statisticTime()
		}
	}
}

package luaComponent

import (
	"encoding/json"
	"phoenixbuilder/omega/defines"
)

type Test struct {
	*defines.BasicComponent
	//通讯用通道
	ConnectChan chan ConnectPackage
	Monitor     *Monitor
}

func (b *Test) Init(cfg *defines.ComponentConfig, storage defines.StorageAndLogProvider) {
	m, _ := json.Marshal(cfg.Configs)
	err := json.Unmarshal(m, b)
	if err != nil {
		panic(err)
	}

}
func (b *Test) Inject(frame defines.MainFrame) {
	b.Frame = frame
	//注入frame等东西
	/*
		b.Frame.GetGameListener().SetOnTypedPacketCallBack(packet.IDAddItemActor, func(p packet.Packet) {
			fmt.Print("凋落物的包:", p, "\n")
		})
	*/
	// 创建一个新的Lua状态

	//初始化lua程序并且返回通讯通道
	b.BasicComponent.Inject(frame)
	b.Monitor = &Monitor{}

	b.Monitor.Start()
	if err := b.Monitor.CmdCenter("lua luas new test"); err != nil {
		printInfo(newPrintMsg("警告", err))
	}
	if err := b.Monitor.CmdCenter("lua start component test"); err != nil {
		printInfo(newPrintMsg("警告", err))
	}
}

// 启动交互器
func (b *Test) LuaFrameworkLauncher() {

}

// 启动go的交互器 与lua建立包的联系
func (b *Test) goFrameworkLauncher() {

}

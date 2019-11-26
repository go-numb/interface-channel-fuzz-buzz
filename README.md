# go-interface-channel-fuzz-buzz

### workers()
map[<識別子>]Worker{}などで辞書化
WorkerStructによって処理や呼び出しを変えるなどが可能  
各Worker処理を書かずとも、Struct単位で1コードで複数処理分岐が可能で楽  
WorkerStruct（対応項目など）が増える場合にも追記が楽  

### getChannels()
channelからInterfaceを満たすstructがくれば、型チェックして型による動的な処理が可能なことを確認   
定時やイベントチャンネルを受け取り、その送られてくるStructで処理を分ける場合などに便利

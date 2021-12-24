package watcher

func StartWatch() {
	ReadCluster()
	WatchCluster()
	WatchClusterNamespace()
}

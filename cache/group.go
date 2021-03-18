package cache
import "github.com/golang/groupcache"
var Group = groupcache.NewGroup("orders", 64<<20, groupcache.GetterFunc(
    func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
        dest.SetString("cached_" + key)
        return nil
    }))
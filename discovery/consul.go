package discovery
import "github.com/hashicorp/consul/api"
func Register() {
    c, _ := api.NewClient(api.DefaultConfig())
    c.Agent().ServiceRegister(&api.AgentServiceRegistration{Name: "order-gateway", Port: 8080})
}
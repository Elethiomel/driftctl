package repository

import (
	"fmt"

	"github.com/cloudskiff/driftctl/pkg/remote/cache"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/floatingips"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/routers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/groups"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/subnets"
	"github.com/sirupsen/logrus"
)

type NeutronRepository interface {
	ListAllPorts() ([]string, error)
	ListAllFloatingips() ([]string, error)
	ListAllNetworks() ([]string, error)
	ListAllRouterInterfaces() ([]string, error)
	ListAllRouters() ([]string, error)
	ListAllSecgroupRules() ([]string, error)
	ListAllSecgroups() ([]string, error)
	ListAllSubnets() ([]string, error)
}

type neutronRepository struct {
	client *gophercloud.ServiceClient
	cache  cache.Cache
}

func NewNeutronRepository(providerClient *gophercloud.ProviderClient, c cache.Cache) *neutronRepository {
	client, err := openstack.NewNetworkV2(providerClient, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Fatalf("Could not create Openstack Client for Neutron : %s", err)
	}

	logrus.Infof("Neutron %+v\n", client)
	return &neutronRepository{
		client,
		c,
	}
}

func (r *neutronRepository) ListAllPorts() ([]string, error) {
	if v := r.cache.Get("neutronListAllPorts"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := ports.List(r.client, ports.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allPorts, err := ports.ExtractPorts(allPages)
	if err != nil {
		panic(err)
	}

	for _, port := range allPorts {
		k = append(k, port.ID)
		logrus.Infof("port %s\n", port.ID)
		fmt.Printf("%+v\n", port)
	}

	//pager := ports.List(r.client, port.ListOpts{})

	// Define an anonymous function to be executed on each page's iteration

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no ports found")
	}

	r.cache.Put("neutronListAllPorts", k)
	return k, nil
}

func (r *neutronRepository) ListAllFloatingips() ([]string, error) {
	if v := r.cache.Get("neutronListAllFloatingips"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := floatingips.List(r.client, floatingips.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allFloatingips, err := floatingips.ExtractFloatingIPs(allPages)
	if err != nil {
		panic(err)
	}

	for _, floatingip := range allFloatingips {
		k = append(k, floatingip.ID)
		logrus.Infof("floatingip %s\n", floatingip.ID)
		fmt.Printf("%+v\n", floatingip)
	}

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no floatingips found")
	}

	r.cache.Put("neutronListAllFloatingips", k)
	return k, nil
}
func (r *neutronRepository) ListAllNetworks() ([]string, error) {
	if v := r.cache.Get("neutronListAllNetworks"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := networks.List(r.client, networks.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allNetworks, err := networks.ExtractNetworks(allPages)
	if err != nil {
		panic(err)
	}

	for _, network := range allNetworks {
		k = append(k, network.ID)
		logrus.Infof("network %s\n", network.ID)
		fmt.Printf("%+v\n", network)
	}

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no networks found")
	}

	r.cache.Put("neutronListAllNetworks", k)
	return k, nil
}
func (r *neutronRepository) ListAllRouterInterfaces() ([]string, error) {
	if v := r.cache.Get("neutronListAllRouterInterfaces"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	logrus.Warn("LISTALLROUTERINTERFACES is a STUB")
	//allPages, err := routerInterfaces.List(r.client, routerInterfaces.ListOpts{}).AllPages()
	//if err != nil {
	//	panic(err)
	//}

	//allRouterInterfaces, err := routerInterfaces.ExtractRouterInterfaces(allPages)
	//if err != nil {
	//	panic(err)
	//}

	//for _, routerInterface := range allRouterInterfaces {
	//	k = append(k, routerInterface.ID)
	//	logrus.Infof("routerInterface %s\n", routerInterface.ID)
	//	fmt.Printf("%+v\n", routerInterface)
	//}

	//if err != nil {
	//	logrus.Infof("Error 2 paging through objects : %s", err)
	//}

	//if len(k) == 0 {
	//	return k, fmt.Errorf("no routerInterfaces found")
	//}

	r.cache.Put("neutronListAllRouterInterfaces", k)
	return k, nil
}
func (r *neutronRepository) ListAllRouters() ([]string, error) {
	if v := r.cache.Get("neutronListAllRouters"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := routers.List(r.client, routers.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allRouters, err := routers.ExtractRouters(allPages)
	if err != nil {
		panic(err)
	}

	for _, router := range allRouters {
		k = append(k, router.ID)
		logrus.Infof("router %s\n", router.ID)
		fmt.Printf("%+v\n", router)
	}

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no routers found")
	}

	r.cache.Put("neutronListAllRouters", k)
	return k, nil
}
func (r *neutronRepository) ListAllSecgroupRules() ([]string, error) {
	if v := r.cache.Get("neutronListAllSecgroupRules"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := rules.List(r.client, rules.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allSecgroupRules, err := rules.ExtractRules(allPages)
	if err != nil {
		panic(err)
	}

	for _, secgroupRule := range allSecgroupRules {
		k = append(k, secgroupRule.ID)
		logrus.Infof("secgroupRule %s\n", secgroupRule.ID)
		fmt.Printf("%+v\n", secgroupRule)
	}

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no secgroupRules found")
	}

	r.cache.Put("neutronListAllSecgroupRules", k)
	return k, nil
}
func (r *neutronRepository) ListAllSecgroups() ([]string, error) {
	if v := r.cache.Get("neutronListAllSecgroups"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := groups.List(r.client, groups.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allSecgroups, err := groups.ExtractGroups(allPages)
	if err != nil {
		panic(err)
	}

	for _, secgroup := range allSecgroups {
		k = append(k, secgroup.ID)
		logrus.Infof("secgroup %s\n", secgroup.ID)
		fmt.Printf("%+v\n", secgroup)
	}

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no secgroups found")
	}

	r.cache.Put("neutronListAllSecgroups", k)
	return k, nil
}
func (r *neutronRepository) ListAllSubnets() ([]string, error) {
	if v := r.cache.Get("neutronListAllSubnets"); v != nil {
		return v.([]string), nil
	}

	k := make([]string, 0)

	allPages, err := subnets.List(r.client, subnets.ListOpts{}).AllPages()
	if err != nil {
		panic(err)
	}

	allSubnets, err := subnets.ExtractSubnets(allPages)
	if err != nil {
		panic(err)
	}

	for _, subnet := range allSubnets {
		k = append(k, subnet.ID)
		logrus.Infof("subnet %s\n", subnet.ID)
		fmt.Printf("%+v\n", subnet)
	}

	if err != nil {
		logrus.Infof("Error 2 paging through objects : %s", err)
	}

	if len(k) == 0 {
		return k, fmt.Errorf("no subnets found")
	}

	r.cache.Put("neutronListAllSubnets", k)
	return k, nil
}

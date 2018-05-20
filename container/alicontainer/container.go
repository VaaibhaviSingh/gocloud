package alicontainer

import (
	"github.com/cloudlibz/gocloud/aliauth"
	"strconv"
	"net/http"
)

//TODO
func (alicontainer *Alicontainer) Createcluster(request interface{}) (resp interface{}, err error) {
	var options CreateCluster

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	var regionID string

	for key, value := range param {
		switch key {
		case "region_id":
			regionID, _ = value.(string)
		case "name":
			name, _ := value.(string)
			options.Name = name
		case "size":
			switch value.(type) {
			case int:
				options.Size = int64(value.(int))
			case int64:
				options.Size = value.(int64)
			case string:
				options.Size, _ = strconv.ParseInt(value.(string), 10, 64)
			}
		case "instance_type":
			instanceType, _ := value.(string)
			options.InstanceType = instanceType
		case "network_mode":
			networkMode, _ := value.(string)
			options.NetworkMode = networkMode
		case "subnet_cidr":
			subnetCIDR, _ := value.(string)
			options.SubnetCIDR = subnetCIDR
		case "vpc_id":
			vpcID, _ := value.(string)
			options.VPCID = vpcID
		case "vswitch_id":
			vSwitchID, _ := value.(string)
			options.VSwitchID = vSwitchID
		case "password":
			password, _ := value.(string)
			options.Password = password
		case "data_disk_category":
			dataDiskCategory, _ := value.(string)
			options.DataDiskCategory = dataDiskCategory
		case "data_disk_size":
			switch value.(type) {
			case int:
				options.Size = int64(value.(int))
			case int64:
				options.Size = value.(int64)
			case string:
				options.DataDiskSize, _ = strconv.ParseInt(value.(string), 10, 64)
			}
		case "ecs_image_id":
			ecsImageID, _ := value.(string)
			options.ECSImageID = ecsImageID
		case "io_optimized":
			IOOptimized, _ := value.(string)
			options.IOOptimized = IOOptimized
		case "need_slb":
			switch value.(type) {
			case bool:
				options.NeedSLB = value.(bool)
			case string:
				options.NeedSLB = value.(string) == "true" || value.(string) == "True"
			}
		case "release_eip_flag":
			switch value.(type) {
			case bool:
				options.ReleaseEipFlag = value.(bool)
			case string:
				options.ReleaseEipFlag = value.(string) == "true" || value.(string) == "True"
			}
		}
	}

	response := make(map[string]interface{})
	err = aliauth.ContainerSignAndDoRequest(regionID, http.MethodPost, "/clusters", nil, options, response)
	resp = response
	return resp, err
}

//TODO
func (alicontainer *Alicontainer) Deletecluster(request interface{}) (resp interface{}, err error) {
	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	var regionID string
	var clusterID string

	for key, value := range param {
		switch key {
		case "region_id":
			regionID, _ = value.(string)
		case "cluster_id":
			clusterID, _ = value.(string)
		}
	}

	response := make(map[string]interface{})
	err = aliauth.ContainerSignAndDoRequest(regionID, http.MethodDelete, "/clusters/"+clusterID, nil, nil, response)
	resp = response
	return resp, err
}

//TODO
func (alicontainer *Alicontainer) Createservice(request interface{}) (resp interface{}, err error) {
	return resp, err
}

//TODO
func (alicontainer *Alicontainer) Runtask(request interface{}) (resp interface{}, err error) {
	return resp, err
}

//TODO
func (alicontainer *Alicontainer) Starttask(request interface{}) (resp interface{}, err error) {
	var options StartTask

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "name":
			options.Name, _ = value.(string)
		case "description":
			options.Description, _ = value.(string)
		case "template":
			options.Template, _ = value.(string)
		case "version":
			options.Version, _ = value.(string)
		case "environment":
			options.Environment, _ = value.(map[string]string)
		case "latest_image":
			switch value.(type) {
			case bool:
				options.LatestImage = value.(bool)
			case string:
				options.LatestImage = value.(string) == "true" || value.(string) == "True"
			}
		}
	}

	response := make(map[string]interface{})
	err = aliauth.ContainerSignAndDoRequest("", http.MethodPost, "/projects/", nil, options, response)
	resp = response
	return resp, err
}

//TODO
func (alicontainer *Alicontainer) Deleteservice(request interface{}) (resp interface{}, err error) {
	return resp, err
}

//TODO
func (alicontainer *Alicontainer) Stoptask(request interface{}) (resp interface{}, err error) {
	return resp, err
}

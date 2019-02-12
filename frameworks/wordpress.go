package main

func GetName() string {
	return "WordPress"
}

func GetInstance() interface{} {
	return &WordPress{
		RootPath:    "/",
		CorePath:    "/",
		ContentPath: "/wp-content",
		VendorPath:  "/vendor",
	}
}

type WordPress struct {
	RootPath    string `json:"root_path"`
	CorePath    string `json:"core_path"`
	ContentPath string `json:"content_path"`
	VendorPath  string `json:"vendor_path"`
}

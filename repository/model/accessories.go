package model

type UserLogin struct {
	Name string
	Pass string
}

type Accout struct {
	Id    int `gorm:"primaryKey"`
	Name  string
	Pass  string
	Email string
}

type Card struct {
	Id               int `gorm:"primaryKey"`
	GPU              string
	Core_speed       string
	Boost_speed      string
	CUDA_cores       string
	Video_memory     string
	Memory_type      string
	Memory_speed     string
	Memory_bus_width string
	Memory_bandwidth string
	Image            string
	PSUP             string
}

type Cpu struct {
	Id             int `gorm:"primaryKey"`
	Cpu            string
	Clock_speed    string
	Bus_speed      string
	Register_width string
	Cache          string
	Socket         string
	Cores_threads  string
	Tdp            string
	Image          string
	Psu            string
}

type HeatDissipation struct {
	Id               int `gorm:"primaryKey"`
	Heat_dissipation string
	Type             string
	Psus             string
	Image            string
}

type Main struct {
	Id           int `gorm:"primaryKey"`
	Main         string
	Size         string
	Socket       string
	Chipset      string
	Hdmi         string
	Display_port string
	Vga          string
	Divi         string
	Wifi         string
	Rgb          string
	Image        string
}

type Psup struct {
	Id               int `gorm:"primaryKey"`
	Psup             string
	Model            string
	Input_voltage    string
	Maximum_capacity string
	Image            string
}

type Ram struct {
	Id           int `gorm:"primaryKey"`
	Ram          string
	Type         string
	Ddr          string
	Dual_channel string
	Bus_speed    string
	Bandwidth    string
	Image        string
}

type Rom struct {
	Id          int `gorm:"primaryKey"`
	Rom         string
	Type        string
	Capacity    string
	Read_speed  string
	Write_speed string
	Image       string
}

type Build struct {
	Id                int `gorm:"PrimaryKey"`
	Name              string
	CardId            int
	MainId            int
	CpuId             int
	RamId             int
	RomId             int
	PsupId            int
	HeatDissipationId int
	UserId            int
	User              Accout          `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE;OnDelete:CACADE"`
	Card              Card            `gorm:"foreignKey:CardId;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	Main              Main            `gorm:"foreignKey:MainId;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	Cpu               Cpu             `gorm:"foreignKey:CpuId;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	Ram               Ram             `gorm:"foreignKey:RamId;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	Rom               Rom             `gorm:"foreignKey:RomId;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	Psup              Psup            `gorm:"foreignKey:PsupId;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
	HeatDissipation   HeatDissipation `gorm:"foreignKey:HeatDissipationId;constraint:OnUpdate:CASCADE;OnDelete:CASCADE"`
}

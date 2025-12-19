package cfg

type Web struct {
	IP   string `ini:"ip" json:"ip" yaml:"ip" comment:"Web 服务监听地址，IP 地址"`
	Port int    `ini:"port" json:"port" yaml:"port" comment:"Web 服务监听端口"`
	Mode string `ini:"mode" json:"mode" yaml:"mode" comment:"Web 服务模式，debug 调试模式，release 发布模式"`
}

type SQLite struct {
	DbName string `ini:"db_name" json:"db_name" yaml:"db_name" comment:"数据库名称，用于指定 SQLite 数据库文件名"`
}

type MySQL struct {
	Host     string `ini:"host" json:"host" yaml:"host" comment:"MySQL 服务器地址，支持域名或 IP"`
	Port     int    `ini:"port" json:"port" yaml:"port" comment:"MySQL 服务器端口号"`
	Username string `ini:"username" json:"username" yaml:"username" comment:"连接 MySQL 的用户名"`
	Password string `ini:"password" json:"password" yaml:"password" comment:"连接 MySQL 的密码"`
	DbName   string `ini:"db_name" json:"db_name" yaml:"db_name" comment:"MySQL 数据库名称"`
}

type Redis struct {
	Host     string `ini:"host" json:"host" yaml:"host" comment:"Redis 服务器地址，支持域名或 IP"`
	Port     int    `ini:"port" json:"port" yaml:"port" comment:"Redis 服务器端口"`
	Username string `ini:"username" json:"username" yaml:"username" comment:"Redis 用户名，若无用户名则为空"`
	Password string `ini:"password" json:"password" yaml:"password" comment:"Redis 连接密码，若无密码则为空"`
	DB       int    `ini:"db" json:"db" yaml:"db" comment:"Redis 数据库索引，默认为 0"`
}

type Email struct {
	SMTPHost string `ini:"smtp_host" json:"smtp_host" yaml:"smtp_host" comment:"邮件服务器地址(smtp.qq.com)"`
	SMTPPort int    `ini:"smtp_port" json:"smtp_port" yaml:"smtp_port" comment:"邮件服务器端口"`
	Username string `ini:"username" json:"username" yaml:"username" comment:"邮件地址"`
	Password string `ini:"password" json:"password" yaml:"password" comment:"SMTP服务，在第三方客户端登录时，需要使用授权密码"`
	From     string `ini:"from" json:"from" yaml:"from" comment:"发件人地址，通常与用户名相同"`
}

type Logger struct {
	Level      string `ini:"level" json:"level" yaml:"level" comment:"数据库日志级别 info > warn > error > silent  silent 不记录任何日志，相当于disabled\n; 系统日志级别   trace > debug > info > warn > error > fatal > panic > no > disabled\n; trace\t\t细粒度最高，最大量日志\n; debug\t\t调试日志\n; info\t\t常规运行状态日志\n; warn\t\t警告，非致命异常\n; error\t\t错误日志，功能异常\n; fatal\t\t致命错误，程序即将终止\n; panic\t\t更严重，触发 panic 行为\n; no\t\t没有级别，适合特殊用途\n; disabled\t禁止所有日志"`
	MaxSize    int    `ini:"max_size" json:"max_size" yaml:"max_size" comment:"单个日志文件最大尺寸，单位为 MB，超过该大小将触发日志切割"`
	LogPath    string `ini:"log_path" json:"log_path" yaml:"log_path" comment:"日志文件存放路径"`
	FormatJSON bool   `ini:"format_json" json:"format_json" yaml:"format_json" comment:"文件日志输出格式，true 表示结构化 JSON，false 表示纯文本"`
	Compress   bool   `ini:"compress" json:"compress" yaml:"compress" comment:"是否压缩旧日志文件，开启后使用 gzip 格式压缩"`
	MaxAge     int    `ini:"max_age" json:"max_age" yaml:"max_age" comment:"日志文件最大保存天数，超过该天数的日志文件将被删除"`
	MaxBackups int    `ini:"max_backups" json:"max_backups" yaml:"max_backups" comment:"保留旧日志文件的最大数量，超过时自动删除最早的日志"`
}

// Captcha 验证码配置
type Captcha struct {
	Length     int      `ini:"length" json:"length" yaml:"length" comment:"验证码长度"`
	NoiseCount int      `ini:"noise_count" json:"noise_count" yaml:"noise_count" comment:"噪点数量（图片上加多少个干扰点）"`
	Type       []string `ini:"type" json:"type" yaml:"type" comment:"验证码类型\n; digit\t\t数字\n; alphaNum\t数字字母\n; chinese\t汉字\n; math\t\t算术运算"`
	Fonts      []string `ini:"fonts" json:"fonts" yaml:"fonts" comment:"验证码字体名称，以下是系统字体，只能选择，不可更换\n; 3Dumb,actionj,ApothecaryFont,chromohv,Comismsh,DeborahFancyDress,DENNEthree-dee,Flim-Flam,RitaSmith,wqy-microhei"`
	ShowLine   []int    `ini:"show_line" json:"show_line" yaml:"show_line" comment:"配置干扰线条样式\n; 0\t\t无干扰线\n; 2\t\t空心线（中空的曲线，像钢笔划出来的线条）\n; 4\t\t粘稠线（细长黏糊糊的线，比较流动性）\n; 8\t\t正弦波线（类似波浪线，波形起伏）\n; 2|4\t空心线+粘稠线，两种线叠加，增强干扰\n; 2|8\t空心线+正弦波线，中空曲线和波浪线叠加\n; 4|8\t粘稠线+正弦波线，黏糊线+波浪线，视觉干扰力较强\n; 2|4|8\t空心线+粘稠线+正弦波线，全体上线，满满的视觉干扰"`
}

type Other struct {
	IsEmail      bool   `ini:"is_email" json:"is_email" yaml:"is_email" comment:"是否启用邮箱"`
	DbType       string `ini:"db_type" json:"db_type" yaml:"db_type" comment:"所使用的数据库类型，支持 SQLite 或 MySQL"`
	IsInitialize bool   `ini:"is_initialize" json:"is_initialize" yaml:"is_initialize" comment:"是否初始化，false 表示未初始化，true 表示已初始化"`
	CacheType    string `ini:"cache_type" json:"cache_type" yaml:"cache_type" comment:"所使用的数据库类型，支持 Memory 或 Redis"`
	DataPath     string `ini:"data_path" json:"data_path" yaml:"data_path" comment:"数据存储路径"`
}

type Config struct {
	SQLite  SQLite  `ini:"SQLite" json:"SQLite" yaml:"SQLite"`
	MySQL   MySQL   `ini:"MySQL" json:"MySQL" yaml:"MySQL"`
	Redis   Redis   `ini:"Redis" json:"Redis" yaml:"Redis"`
	Email   Email   `ini:"Email" json:"Email" yaml:"Email"`
	Web     Web     `ini:"Web" json:"Web" yaml:"Web"`
	Logger  Logger  `ini:"Logger" json:"Logger" yaml:"Logger"`
	Captcha Captcha `ini:"Captcha" json:"Captcha" yaml:"Captcha"`
	Other   Other   `ini:"Other" json:"Other" yaml:"Other"`
}

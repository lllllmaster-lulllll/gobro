package factory

// todo: 简单工厂模式
type IRuleConfigParser interface {
	Parse(data []byte)
}

type jsonRuleConfigParser struct {
}

func (j *jsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type yamlRuleConfigParser struct {
}

func (y *yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return &jsonRuleConfigParser{}
	case "yaml":
		return &yamlRuleConfigParser{}
	}
	return nil
}

// todo: 工厂方法
// todo:IRuleConfigParserFactory 工厂方法接口
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// todo:yamlRuleConfigParserFactory yamlRuleConfigParser 的工厂类
type yamlRuleConfigParserFactory struct {
}

// todo:CreateParser CreateParser
func (y *yamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &yamlRuleConfigParser{}
}

// TODO:jsonRuleConfigParserFactory jsonRuleConfigParser 的工厂类
type jsonRuleConfigParserFactory struct {
}

// todo:CreateParser CreateParser
func (y *jsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return &jsonRuleConfigParser{}
}

// TODO: NewIRuleConfigParserFactory 用一个简单工厂封装工厂方法
func NewIRuleConfigParserFactory(t string) IRuleConfigParserFactory {
	switch t {
	case "json":
		return &jsonRuleConfigParserFactory{}
	case "yaml":
		return &yamlRuleConfigParserFactory{}
	}
	return nil
}

// todo: 抽象工厂模式
type ISystemConfigParser interface {
	ParseSystem(data []byte)
}

// ISystemConfigParser
type jsonSystemConfigParser struct {
}

func (j *jsonSystemConfigParser) ParseSystem(data []byte) {
	panic("implement me")
}

type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}
type jsonConfigParserFactory struct {
}

func (j *jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return &jsonRuleConfigParser{}
}
func (j *jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return &jsonSystemConfigParser{}
}
